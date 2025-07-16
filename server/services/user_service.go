package services

import (
	"context"
	"fmt"
	"time"

	"github.com/fiqrioemry/event_ticketing_system_app/server/config"
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/models"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"

	"github.com/fiqrioemry/go-api-toolkit/response"
)

type UserService interface {
	GetUserProfile(userID string) (*dto.ProfileResponse, error)
	GetUserDetail(id string) (*dto.UserDetailResponse, error)
	GetAllUsers(params dto.UserQueryParams) ([]dto.UserListResponse, int, error)
	UpdateUserDetail(userID string, req *dto.UpdateProfileRequest) (*models.User, error)

	// change password features
	ChangePassword(userID string, req *dto.ChangePasswordRequest) error

	// password reset features
	ForgotPassword(c *gin.Context, req *dto.ForgotPasswordRequest) error
	ValidateToken(token string) (string, error)
	ResetPassword(req *dto.ResetPasswordRequest) error

	// Google OAuth features
	GetGoogleOAuthURL() string
	GoogleSignIn(tokenId string) (*dto.AuthResponse, error)
	HandleGoogleOAuthCallback(code string) (*dto.AuthResponse, error)
}

type userService struct {
	user repositories.UserRepository
}

func NewUserService(user repositories.UserRepository) UserService {
	return &userService{user: user}
}

func (s *userService) GetUserProfile(userID string) (*dto.ProfileResponse, error) {
	user, err := s.user.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, response.NewNotFound("user not found")
	}

	profile := &dto.ProfileResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		Fullname: user.Fullname,
		Avatar:   user.AvatarURL,
		Balance:  user.Balance,
		JoinedAt: user.CreatedAt,
		Role:     user.Role,
	}
	return profile, nil
}

func (s *userService) UpdateUserDetail(userID string, req *dto.UpdateProfileRequest) (*models.User, error) {
	user, err := s.user.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, response.NewNotFound("user not found")
	}

	user.Fullname = req.Fullname
	user.AvatarURL = req.AvatarURL

	if err := s.user.UpdateUser(user); err != nil {
		return nil, response.NewInternalServerError("failed to update user profile", err)
	}

	return user, nil
}

func (s *userService) GetAllUsers(params dto.UserQueryParams) ([]dto.UserListResponse, int, error) {
	users, total, err := s.user.GetAllUsers(params)
	if err != nil {
		return nil, 0, response.NewInternalServerError("failed to fetch user list", err)
	}

	var results []dto.UserListResponse
	for _, u := range users {
		results = append(results, dto.UserListResponse{
			ID:       u.ID.String(),
			Email:    u.Email,
			Role:     u.Role,
			Avatar:   u.AvatarURL,
			Fullname: u.Fullname,
			JoinedAt: u.CreatedAt,
		})
	}

	return results, int(total), nil
}

func (s *userService) GetUserDetail(id string) (*dto.UserDetailResponse, error) {
	user, err := s.user.GetUserByID(id)
	if err != nil || user == nil {
		return nil, response.NewNotFound("user not found").WithContext("userID", id)
	}

	res := &dto.UserDetailResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		Role:     user.Role,
		Avatar:   user.AvatarURL,
		Fullname: user.Fullname,
		JoinedAt: user.CreatedAt,
	}

	return res, nil
}

func (s *userService) ChangePassword(userID string, req *dto.ChangePasswordRequest) error {
	// Validate confirm password
	if req.NewPassword != req.ConfirmPassword {
		return response.NewBadRequest("New password and confirm password don't match")
	}

	// check if user exists
	user, err := s.user.GetUserByID(userID)
	if err != nil {
		return response.NewNotFound("User not found")
	}

	// Verify current password
	if !utils.CheckPasswordHash(req.CurrentPassword, user.Password) {
		return response.NewBadRequest("Current password is incorrect")
	}

	// Hash new password
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return response.NewInternalServerError("Failed to hash password", err)
	}

	// Update password
	user.Password = hashedPassword
	if err := s.user.UpdateUser(user); err != nil {
		return response.NewInternalServerError("Failed to update password", err)
	}

	return nil
}

func (s *userService) ForgotPassword(c *gin.Context, req *dto.ForgotPasswordRequest) error {
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

func (s *userService) ResetPassword(req *dto.ResetPasswordRequest) error {

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

func (s *userService) ValidateToken(token string) (string, error) {

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

func (s *userService) GoogleSignIn(tokenId string) (*dto.AuthResponse, error) {
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
		user = &models.User{
			Email:     email,
			AvatarURL: utils.RandomUserAvatar(name),
			Fullname:  name,
			Password:  "-",
		}

		if err := s.user.CreateUser(user); err != nil {
			return nil, err
		}

		if user.ID == uuid.Nil {
			return nil, response.NewInternalServerError("Failed to assign UUID to user", err)
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

func (s *userService) GetGoogleOAuthURL() string {
	return config.GoogleOAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func (s *userService) HandleGoogleOAuthCallback(code string) (*dto.AuthResponse, error) {
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
