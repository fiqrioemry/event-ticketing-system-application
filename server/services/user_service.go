package services

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/models"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"github.com/fiqrioemry/go-api-toolkit/response"
)

type UserService interface {
	GetUserProfile(userID string) (*dto.ProfileResponse, error)
	ChangePassword(userID string, req *dto.ChangePasswordRequest) error
	UpdateUserDetail(userID string, req *dto.UpdateProfileRequest) (*models.User, error)
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
		Avatar:   user.Avatar,
		Role:     user.Role,
		Balance:  user.Balance,
		JoinedAt: user.CreatedAt,
	}
	return profile, nil
}

func (s *userService) UpdateUserDetail(userID string, req *dto.UpdateProfileRequest) (*models.User, error) {
	user, err := s.user.GetUserByID(userID)
	if err != nil || user == nil {
		return nil, response.NewNotFound("user not found")
	}

	user.Fullname = req.Fullname

	if req.AvatarURL != "" {
		user.Avatar = req.AvatarURL
	}

	if err := s.user.UpdateUser(user); err != nil {
		return nil, response.NewInternalServerError("failed to update user profile", err)
	}

	return user, nil
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
		Avatar:   user.Avatar,
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
