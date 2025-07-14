package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"server/config"
	"server/dto"
	"server/models"
	"server/repositories"
	"server/utils"
	"time"

	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type AuthService interface {
	ResendOTP(email string) error
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

func (s *authService) ResendOTP(email string) error {
	otpKey := "ticket:otp_data:" + email

	otpDataStr, err := config.RedisClient.Get(config.Ctx, otpKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return response.NewNotFound("Registration session not found. Please register again")
		}
		return response.NewInternalServerError("Failed to retrieve OTP data", err)
	}

	var otpData map[string]interface{}
	if err := json.Unmarshal([]byte(otpDataStr), &otpData); err != nil {
		return response.NewInternalServerError("Invalid OTP data", err)
	}

	limitKey := "ticket:otp_resend_limit:" + email
	count, _ := config.RedisClient.Get(config.Ctx, limitKey).Int()
	if count >= 3 {
		return response.NewTooManyRequests("Too many OTP requests. Try again later")
	}

	newOTP := utils.GenerateOTP(6)

	otpData["otp"] = newOTP
	otpData["last_sent"] = time.Now().Unix()

	updatedData, _ := json.Marshal(otpData)
	if err := config.RedisClient.Set(config.Ctx, otpKey, updatedData, 5*time.Minute).Err(); err != nil {
		return response.NewInternalServerError("Failed to save OTP to Redis", err)
	}

	subject := "Your New OTP Code"
	body := fmt.Sprintf("Your new OTP is %s", newOTP)
	if err := utils.SendEmail(subject, email, newOTP, body); err != nil {
		return response.NewInternalServerError("Failed to send OTP email", err)
	}

	config.RedisClient.Incr(config.Ctx, limitKey)
	config.RedisClient.Expire(config.Ctx, limitKey, 30*time.Minute)

	return nil
}
func (s *authService) VerifyOTP(email, otp string) (*dto.AuthResponse, error) {
	savedOtp, err := config.RedisClient.Get(config.Ctx, "ticket:otp:"+email).Result()
	if err != nil || savedOtp != otp {
		return nil, response.NewUnauthorized("OTP is invalid or has expired")
	}

	config.RedisClient.Del(config.Ctx, "ticket:otp:"+email)

	val, err := config.RedisClient.Get(config.Ctx, "ticket:otp_data:"+email).Result()
	if err != nil {
		return nil, response.NewUnauthorized("Session has expired")
	}
	var temp map[string]string
	if err := json.Unmarshal([]byte(val), &temp); err != nil {
		return nil, response.NewInternalServerError("Failed to parse user data", err)
	}

	user := models.User{
		ID:        uuid.New(),
		Email:     email,
		Fullname:  temp["fullname"],
		Password:  temp["password"],
		AvatarURL: utils.RandomUserAvatar(temp["fullname"]),
	}

	if err := s.user.CreateUser(&user); err != nil {
		return nil, response.NewConflict("Email is already registered")
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, response.NewInternalServerError("Failed to generate access token", err)
	}
	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, response.NewInternalServerError("Failed to generate refresh token", err)
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
		return nil, response.NewTooManyRequests("Too many login attempts, please try again later")
	}

	user, err := s.user.GetUserByEmail(req.Email)
	if err != nil || user == nil || !utils.CheckPasswordHash(req.Password, user.Password) {
		config.RedisClient.Incr(config.Ctx, redisKey)
		config.RedisClient.Expire(config.Ctx, redisKey, 30*time.Minute)
		return nil, response.NewUnauthorized("Invalid email or password")
	}

	config.RedisClient.Del(config.Ctx, redisKey)

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, response.NewInternalServerError("Failed to generate access token", err)
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, response.NewInternalServerError("Failed to generate refresh token", err)
	}

	userResponse := dto.ProfileResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		Fullname: user.Fullname,
		Balance:  user.Balance,
		Avatar:   user.AvatarURL,
		Role:     user.Role,
		JoinedAt: user.CreatedAt,
	}

	return &dto.AuthResponse{
		User:         userResponse,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (s *authService) Register(req *dto.RegisterRequest) error {
	user, err := s.user.GetUserByEmail(req.Email)
	if err != nil {
		return response.NewInternalServerError("Failed to check user existence", err)
	}

	if user != nil {
		return response.NewConflict("Email is already registered")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return response.NewInternalServerError("Failed to hash password", err)
	}

	otp := utils.GenerateOTP(6)
	subject := "One-Time Password (OTP)"
	body := fmt.Sprintf("Your OTP code is %s", otp)

	if err := utils.SendEmail(subject, req.Email, otp, body); err != nil {
		return response.NewInternalServerError("Failed to send OTP email", err)
	}

	if err := config.RedisClient.Set(config.Ctx, "ticket:otp:"+req.Email, otp, 5*time.Minute).Err(); err != nil {
		return response.NewInternalServerError("Failed to save OTP to Redis", err)
	}

	tempData := map[string]string{
		"fullname": req.Fullname,
		"password": hashedPassword,
		"email":    req.Email,
	}
	jsonStr, err := json.Marshal(tempData)
	if err != nil {
		return response.NewInternalServerError("Failed to marshal user data", err)
	}
	if err := config.RedisClient.Set(config.Ctx, "ticket:otp_data:"+req.Email, jsonStr, 30*time.Minute).Err(); err != nil {
		return response.NewInternalServerError("Failed to save OTP data to Redis", err)
	}

	return nil
}

func (s *authService) RefreshToken(c *gin.Context, refreshToken string) (string, error) {

	userID, err := utils.DecodeRefreshToken(refreshToken)
	if err != nil {
		return "", response.NewUnauthorized("Invalid refresh token")
	}

	user, err := s.user.GetUserByID(userID)
	if err != nil || user == nil {
		return "", response.NewNotFound("User not found").WithContext("userID", userID)
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return "", response.NewInternalServerError("Failed to generate access token", err)
	}

	return accessToken, nil
}
