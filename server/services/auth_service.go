package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/fiqrioemry/event_ticketing_system_app/server/config"
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/models"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"

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

	// password reset features
	ForgotPassword(c *gin.Context, req *dto.ForgotPasswordRequest) error
	ValidateToken(token string) (string, error)
	ResetPassword(req *dto.ResetPasswordRequest) error

	// Google OAuth features
	GetGoogleOAuthURL() string
	GoogleSignIn(tokenId string) (*dto.AuthResponse, error)
	HandleGoogleOAuthCallback(code string) (*dto.AuthResponse, error)
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
	otpData := map[string]any{
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

	// send OTP email
	if err := utils.SendOTPEmail(req.Email, req.Fullname, otp, 15*60*time.Second); err != nil {
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
	var otpData map[string]any
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
	var userData map[string]any
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
		ID:       uuid.New(),
		Email:    email,
		Fullname: fullname,
		Password: password,
		Role:     "user", // Default role for new users
		Avatar:   utils.RandomUserAvatar(fullname),
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
		Avatar:   user.Avatar,
		Role:     user.Role,
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
		Avatar:   user.Avatar,
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
		Balance:  user.Balance,
		JoinedAt: user.CreatedAt,
		Avatar:   user.Avatar,
		Role:     user.Role,
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, "", response.NewInternalServerError("Failed to generate access token", err)
	}

	return &userResponse, accessToken, nil
}

func (s *authService) ForgotPassword(c *gin.Context, req *dto.ForgotPasswordRequest) error {
	// Check rate limit attempts
	attemptsKey := "asset_app:forgot_password_attempts:" + c.ClientIP()
	if err := utils.CheckForgotPasswordAttempts(c.ClientIP(), 3); err != nil {
		return response.NewTooManyRequests("Too many forgot password attempts, please try again later")
	}

	// Check token existence
	existingTokenKey := "asset_app:reset_token:" + req.Email
	if utils.KeyExists(existingTokenKey) {
		return response.NewTooManyRequests("Password reset link has already been sent. Please check your email or wait before requesting again.")
	}

	// Check if email exists
	user, err := s.user.GetUserByEmail(req.Email)
	if err != nil {
		// Increment attempts
		utils.IncrementAttempts(attemptsKey)
		return nil // Don't reveal if email exists for security reasons
	}

	if user == nil {
		utils.IncrementAttempts(attemptsKey)
		return nil // Don't reveal if email exists
	}

	// Generate reset token
	resetToken, err := utils.GenerateResetToken()
	if err != nil {
		return response.NewInternalServerError("Failed to generate reset token", err)
	}

	// Prepare token data
	tokenData := map[string]any{
		"userId":    user.ID.String(),
		"email":     user.Email,
		"createdAt": time.Now().Unix(),
		"expiresAt": time.Now().Add(1 * time.Hour).Unix(),
	}

	// Store reset token data
	resetTokenKey := "asset_app:password_reset:" + resetToken
	if err := utils.AddKeys(resetTokenKey, tokenData, 1*time.Hour); err != nil {
		return response.NewInternalServerError("Failed to store reset token", err)
	}

	// Store email -
	emailTokenKey := "asset_app:reset_token:" + user.Email
	if err := utils.AddKeys(emailTokenKey, resetToken, 1*time.Hour); err != nil {
		// Clean up reset token
		utils.DeleteKeys(resetTokenKey)
		return response.NewInternalServerError("Failed to store email token mapping", err)
	}

	// Create reset link
	frontendURL := config.AppConfig.FrontendURL

	resetLink := fmt.Sprintf("%s/reset-password?token=%s", frontendURL, resetToken)

	// Send reset password email
	if err := utils.SendResetPasswordEmail(user.Email, user.Fullname, resetLink, 1*time.Hour); err != nil {
		// Clean up tokens
		utils.DeleteKeys(resetTokenKey, emailTokenKey)
		return response.NewInternalServerError("Failed to send reset password email", err)
	}

	// Increment attempts
	go utils.IncrementAttempts(attemptsKey)

	return nil
}

func (s *authService) ResetPassword(req *dto.ResetPasswordRequest) error {

	// check password match
	if req.NewPassword != req.ConfirmPassword {
		return response.NewBadRequest("New password and confirm password do not match")
	}

	resetTokenKey := "asset_app:password_reset:" + req.Token
	var tokenData map[string]any

	// get token from cache
	if err := utils.GetKey(resetTokenKey, &tokenData); err != nil {
		return response.NewBadRequest("Invalid or expired reset token")
	}

	email, ok := tokenData["email"].(string)
	if !ok {
		return response.NewBadRequest("Invalid reset token data")
	}

	userID, ok := tokenData["userId"].(string)
	if !ok {
		return response.NewBadRequest("Invalid reset token data")
	}

	// check user exists
	user, err := s.user.GetUserByID(userID)
	if err != nil {
		return response.NewNotFound("User not found")
	}

	// hash new password
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return response.NewInternalServerError("Failed to hash password", err)
	}

	// update user password
	user.Password = hashedPassword

	if err := s.user.UpdateUser(user); err != nil {
		return response.NewInternalServerError("Failed to update password", err)
	}

	// delete related cache keys
	go utils.DeleteKeys(resetTokenKey)
	go utils.DeleteKeys("asset_app:reset_token:" + email)
	go utils.DeleteKeys("asset_app:forgot_password_attempts:" + email)

	return nil
}

func (s *authService) ValidateToken(token string) (string, error) {

	resetTokenKey := "asset_app:password_reset:" + token
	var tokenData map[string]any
	// check token existence
	if err := utils.GetKey(resetTokenKey, &tokenData); err != nil {
		return "", response.NewBadRequest("Invalid or expired reset token")
	}

	// Check token age
	createdAt, ok := tokenData["createdAt"].(float64)
	if !ok {
		return "", response.NewBadRequest("Invalid token data")
	}

	tokenAge := time.Since(time.Unix(int64(createdAt), 0))
	if tokenAge > 1*time.Hour {
		utils.DeleteKeys(resetTokenKey)
		return "", response.NewBadRequest("Reset token has expired")
	}

	// email for display
	email, _ := tokenData["email"].(string)

	return email, nil

}

func (s *authService) GoogleSignIn(tokenId string) (*dto.AuthResponse, error) {
	payload, err := idtoken.Validate(context.Background(), tokenId, config.AppConfig.GoogleClientID)
	if err != nil {
		return nil, response.NewUnauthorized("Invalid Google ID token")
	}

	email, ok := payload.Claims["email"].(string)
	if !ok || email == "" {
		return nil, response.NewNotFound("Email not found in token")
	}

	name, _ := payload.Claims["name"].(string)

	user, err := s.user.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	// Jika user belum ada, buat user baru
	if user == nil {
		user = &models.User{
			Email:    email,
			Avatar:   utils.RandomUserAvatar(name),
			Fullname: name,
			Role:     "user",
			Password: "-", // placeholder karena pakai OAuth
		}

		if err := s.user.CreateUser(user); err != nil {
			return nil, err
		}
	}

	accessToken, err := utils.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, err
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authService) GetGoogleOAuthURL() string {
	return config.GoogleOAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func (s *authService) HandleGoogleOAuthCallback(code string) (*dto.AuthResponse, error) {
	token, err := config.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, response.NewUnauthorized("Failed to exchange Google OAuth code")
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, response.NewUnauthorized("ID token not found in Google OAuth response")
	}

	return s.GoogleSignIn(rawIDToken)
}
