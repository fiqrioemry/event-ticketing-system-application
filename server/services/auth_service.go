package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/fiqrioemry/event_ticketing_system_app/server/config"
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/models"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

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
	RefreshToken(c *gin.Context, refreshToken string) (*dto.ProfileResponse, string, error)
}

type authService struct {
	user repositories.UserRepository
}

func NewAuthService(user repositories.UserRepository) AuthService {
	return &authService{user: user}
}

func (s *authService) Register(req *dto.RegisterRequest) error {
	// Check if user already exists
	user, err := s.user.GetUserByEmail(req.Email)
	if err != nil {
		return response.NewInternalServerError("Failed to check user existence", err)
	}

	if user != nil {
		return response.NewConflict("Email is already registered")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return response.NewInternalServerError("Failed to hash password", err)
	}

	// Generate OTP
	otp := utils.GenerateOTP(6)

	// Store OTP and user data in Redis with consistent TTL
	otpKey := "ticket:otp:" + req.Email
	otpDataKey := "ticket:otp_data:" + req.Email

	// Set OTP expiration time (5 minutes)
	otpTTL := 5 * time.Minute

	// Store OTP
	if err := config.RedisClient.Set(config.Ctx, otpKey, otp, otpTTL).Err(); err != nil {
		return response.NewInternalServerError("Failed to save OTP to Redis", err)
	}

	// Store user registration data
	otpData := map[string]interface{}{
		"fullname":   req.Fullname,
		"password":   hashedPassword,
		"email":      req.Email,
		"otp":        otp,
		"created_at": time.Now().Unix(),
		"last_sent":  time.Now().Unix(),
	}

	jsonData, err := json.Marshal(otpData)
	if err != nil {
		return response.NewInternalServerError("Failed to marshal user data", err)
	}

	// Store with longer TTL for user data (30 minutes)
	if err := config.RedisClient.Set(config.Ctx, otpDataKey, jsonData, 30*time.Minute).Err(); err != nil {
		return response.NewInternalServerError("Failed to save user data to Redis", err)
	}

	// Send OTP email
	subject := "One-Time Password (OTP)"
	body := fmt.Sprintf("Your OTP code is %s. This code will expire in 5 minutes.", otp)

	if err := utils.SendEmail(subject, req.Email, otp, body); err != nil {
		// Clean up Redis data if email fails
		config.RedisClient.Del(config.Ctx, otpKey)
		config.RedisClient.Del(config.Ctx, otpDataKey)
		return response.NewInternalServerError("Failed to send OTP email", err)
	}

	return nil
}

func (s *authService) ResendOTP(email string) error {
	otpDataKey := "ticket:otp_data:" + email
	otpKey := "ticket:otp:" + email

	// Get existing registration data
	otpDataStr, err := config.RedisClient.Get(config.Ctx, otpDataKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return response.NewNotFound("Registration session not found. Please register again")
		}
		return response.NewInternalServerError("Failed to retrieve registration data", err)
	}

	// Check resend limit
	limitKey := "ticket:otp_resend_limit:" + email
	count, _ := config.RedisClient.Get(config.Ctx, limitKey).Int()
	if count >= 3 {
		return response.NewTooManyRequests("Too many OTP requests. Try again later")
	}

	// Parse existing data
	var otpData map[string]interface{}
	if err := json.Unmarshal([]byte(otpDataStr), &otpData); err != nil {
		return response.NewInternalServerError("Invalid registration data", err)
	}

	// Generate new OTP
	newOTP := utils.GenerateOTP(6)

	// Update OTP data
	otpData["otp"] = newOTP
	otpData["last_sent"] = time.Now().Unix()

	// Save updated data
	updatedData, _ := json.Marshal(otpData)
	if err := config.RedisClient.Set(config.Ctx, otpDataKey, updatedData, 30*time.Minute).Err(); err != nil {
		return response.NewInternalServerError("Failed to update registration data", err)
	}

	// Update OTP key with new OTP and reset TTL
	if err := config.RedisClient.Set(config.Ctx, otpKey, newOTP, 5*time.Minute).Err(); err != nil {
		return response.NewInternalServerError("Failed to save new OTP", err)
	}

	// Send new OTP email
	subject := "Your New OTP Code"
	body := fmt.Sprintf("Your new OTP code is %s. This code will expire in 5 minutes.", newOTP)
	if err := utils.SendEmail(subject, email, newOTP, body); err != nil {
		return response.NewInternalServerError("Failed to send OTP email", err)
	}

	// Update resend limit
	config.RedisClient.Incr(config.Ctx, limitKey)
	config.RedisClient.Expire(config.Ctx, limitKey, 30*time.Minute)

	return nil
}

func (s *authService) VerifyOTP(email, otp string) (*dto.AuthResponse, error) {
	otpKey := "ticket:otp:" + email
	otpDataKey := "ticket:otp_data:" + email

	// Verify OTP from Redis
	savedOtp, err := config.RedisClient.Get(config.Ctx, otpKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, response.NewUnauthorized("OTP has expired or does not exist")
		}
		return nil, response.NewInternalServerError("Failed to retrieve OTP", err)
	}

	// Check if OTP matches
	if savedOtp != otp {
		return nil, response.NewUnauthorized("Invalid OTP code")
	}

	// Get user registration data
	userDataStr, err := config.RedisClient.Get(config.Ctx, otpDataKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, response.NewUnauthorized("Registration session has expired")
		}
		return nil, response.NewInternalServerError("Failed to retrieve registration data", err)
	}

	// Parse user data
	var userData map[string]interface{}
	if err := json.Unmarshal([]byte(userDataStr), &userData); err != nil {
		return nil, response.NewInternalServerError("Failed to parse registration data", err)
	}

	// Extract user information with type assertions
	fullname, ok := userData["fullname"].(string)
	if !ok {
		return nil, response.NewInternalServerError("Invalid fullname in registration data", nil)
	}

	password, ok := userData["password"].(string)
	if !ok {
		return nil, response.NewInternalServerError("Invalid password in registration data", nil)
	}

	// Create user
	user := models.User{
		ID:        uuid.New(),
		Email:     email,
		Fullname:  fullname,
		Password:  password,
		AvatarURL: utils.RandomUserAvatar(fullname),
	}

	// Save user to database
	if err := s.user.CreateUser(&user); err != nil {
		return nil, response.NewConflict("Email is already registered")
	}

	// Generate tokens
	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, response.NewInternalServerError("Failed to generate access token", err)
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, response.NewInternalServerError("Failed to generate refresh token", err)
	}

	// Clean up Redis data after successful verification
	config.RedisClient.Del(config.Ctx, otpKey)
	config.RedisClient.Del(config.Ctx, otpDataKey)
	config.RedisClient.Del(config.Ctx, "ticket:otp_resend_limit:"+email)

	// Prepare response
	userResponse := dto.ProfileResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		Fullname: user.Fullname,
		Avatar:   user.AvatarURL,
	}

	return &dto.AuthResponse{
		User:         userResponse,
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
		return nil, response.NewBadRequest("Invalid email or password")
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

func (s *authService) RefreshToken(c *gin.Context, refreshToken string) (*dto.ProfileResponse, string, error) {

	userID, err := utils.DecodeRefreshToken(refreshToken)
	if err != nil {
		return nil, "", response.NewUnauthorized("Invalid refresh token")
	}

	user, err := s.user.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, "", response.NewNotFound("User not found").WithContext("userID", userID)
	}

	userResponse := dto.ProfileResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		Fullname: user.Fullname,
		Avatar:   user.AvatarURL,
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, "", response.NewInternalServerError("Failed to generate access token", err)
	}

	return &userResponse, accessToken, nil
}
