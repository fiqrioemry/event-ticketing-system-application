package services

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/models"
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"

	"github.com/fiqrioemry/go-api-toolkit/response"
)

type UserService interface {
	GetUserProfile(userID string) (*dto.ProfileResponse, error)
	GetUserDetail(id string) (*dto.UserDetailResponse, error)
	GetAllUsers(params dto.UserQueryParams) ([]dto.UserListResponse, int, error)
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
