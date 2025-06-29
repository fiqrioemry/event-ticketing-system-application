package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"server/config"
	"server/dto"
	customErr "server/errors"
	"server/models"
	"server/repositories"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService interface {
	SendOTP(email string) error
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.AuthResponse, error)
	VerifyOTP(email, otp string) (*dto.AuthResponse, error)
	RefreshToken(c *gin.Context, refreshToken string) (string, error)
}

type authService struct {
	user repositories.UserRepository
}

func NewAuthService(user repositories.UserRepository) AuthService {
	return &authService{user: user}
}

func (s *authService) SendOTP(email string) error {
	// get user data from Redis
	_, err := config.RedisClient.Get(config.Ctx, "otp_data:"+email).Result()
	if err != nil {
		return customErr.NewNotFound("OTP data not found")
	}

	// check if the user has requested OTP too many times
	limitKey := "otp_resend_limit:" + email
	count, _ := config.RedisClient.Get(config.Ctx, limitKey).Int()
	if count >= 3 {
		return customErr.NewTooManyRequest("Too many OTP requests")
	}
	config.RedisClient.Incr(config.Ctx, limitKey)
	config.RedisClient.Expire(config.Ctx, limitKey, 30*time.Minute)

	// generate OTP and send email
	otp := utils.GenerateOTP(6)
	if err := config.RedisClient.Set(config.Ctx, "otp:"+email, otp, 5*time.Minute).Err(); err != nil {
		return customErr.NewInternal("Failed to store OTP", err)
	}

	subject := "Your New OTP Code"
	body := fmt.Sprintf("Your new OTP is %s", otp)
	if err := utils.SendEmail(subject, email, otp, body); err != nil {
		return customErr.NewInternal("Failed to send OTP email", err)
	}

	return nil
}

func (s *authService) VerifyOTP(email, otp string) (*dto.AuthResponse, error) {
	savedOtp, err := config.RedisClient.Get(config.Ctx, "otp:"+email).Result()
	if err != nil || savedOtp != otp {
		return nil, customErr.NewUnauthorized("OTP is invalid or has expired")
	}
	config.RedisClient.Del(config.Ctx, "otp:"+email)

	val, err := config.RedisClient.Get(config.Ctx, "otp_data:"+email).Result()
	if err != nil {
		return nil, customErr.NewUnauthorized("Session has expired")
	}
	var temp map[string]string
	if err := json.Unmarshal([]byte(val), &temp); err != nil {
		return nil, customErr.NewInternal("Failed to parse user data", err)
	}

	user := models.User{
		ID:        uuid.New(),
		Email:     email,
		Fullname:  temp["fullname"],
		Password:  temp["password"],
		AvatarURL: utils.RandomUserAvatar(temp["fullname"]),
	}

	if err := s.user.CreateUser(&user); err != nil {
		return nil, customErr.NewConflict("Email is already registered")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, customErr.ErrTokenGeneration
	}
	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, customErr.ErrTokenGeneration
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	redisKey := fmt.Sprintf("login:attempt:%s", req.Email)
	attempts, _ := config.RedisClient.Get(config.Ctx, redisKey).Int()
	if attempts >= 5 {
		return nil, customErr.NewTooManyRequest("Too many request, please try again in 30 minutes")
	}

	user, err := s.user.GetUserByEmail(req.Email)
	if err != nil || user == nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		config.RedisClient.Incr(config.Ctx, redisKey)
		config.RedisClient.Expire(config.Ctx, redisKey, 30*time.Minute)
		return nil, customErr.NewUnauthorized("Invalid email or password")
	}

	config.RedisClient.Del(config.Ctx, redisKey)

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, customErr.ErrTokenGeneration
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, customErr.ErrTokenGeneration
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (s *authService) Register(req *dto.RegisterRequest) error {
	user, err := s.user.GetUserByEmail(req.Email)
	if err == nil && user != nil {
		// Email sudah terdaftar
		return customErr.NewAlreadyExist("Email already registered")
	}

	// Jika error bukan karena record not found, maka itu error internal
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return customErr.ErrInternalServer
	}

	// Generate password hash
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return customErr.ErrInternalServer
	}

	// Kirim OTP
	otp := utils.GenerateOTP(6)
	subject := "One-Time Password (OTP)"
	body := fmt.Sprintf("Your OTP code is %s", otp)

	if err := utils.SendEmail(subject, req.Email, otp, body); err != nil {
		return customErr.ErrInternalServer
	}

	// Simpan OTP ke Redis
	if err := config.RedisClient.Set(config.Ctx, "otp:"+req.Email, otp, 5*time.Minute).Err(); err != nil {
		return customErr.ErrInternalServer
	}

	// Simpan data user sementara di Redis
	tempData := map[string]string{
		"fullname": req.Fullname,
		"password": hashedPassword,
		"email":    req.Email,
	}
	jsonStr, err := json.Marshal(tempData)
	if err != nil {
		return customErr.ErrInternalServer
	}
	if err := config.RedisClient.Set(config.Ctx, "otp_data:"+req.Email, jsonStr, 30*time.Minute).Err(); err != nil {
		return customErr.ErrInternalServer
	}

	return nil
}

func (s *authService) RefreshToken(c *gin.Context, refreshToken string) (string, error) {

	// check if refresh token is provided
	userID, err := utils.DecodeRefreshToken(refreshToken)
	if err != nil {
		return "", customErr.ErrUnauthorized
	}

	user, err := s.user.GetUserByID(userID)
	if err != nil {
		return "", customErr.NewInternal("failed to find user", err)
	}
	if user == nil {
		return "", customErr.NewNotFound("User not found")
	}

	// regenerate access token
	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return "", customErr.ErrTokenGeneration
	}

	return accessToken, nil
}
