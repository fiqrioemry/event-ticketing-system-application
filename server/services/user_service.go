package services

import (
	"server/dto"
	customErr "server/errors"
	"server/repositories"
	"server/utils"
)

type UserService interface {
	GetUserProfile(userID string) (*dto.ProfileResponse, error)
	GetUserDetail(id string) (*dto.UserDetailResponse, error)
	UpdateUserDetail(userID string, request *dto.UpdateProfileRequest) error
	GetAllUsers(params dto.UserQueryParams) ([]dto.UserListResponse, *dto.PaginationResponse, error)
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
		return nil, customErr.NewNotFound("user not found").WithContext("userID", userID)
	}

	return &dto.ProfileResponse{
		ID:       user.ID.String(),
		Email:    user.Email,
		Fullname: user.Fullname,
		Avatar:   user.AvatarURL,
		Balance:  user.Balance,
		JoinedAt: user.CreatedAt,
		Role:     user.Role,
	}, nil
}

func (s *userService) UpdateUserDetail(userID string, req *dto.UpdateProfileRequest) error {
	user, err := s.user.GetUserByID(userID)
	if err != nil || user == nil {
		return customErr.NewNotFound("user not found").WithContext("userID", userID)
	}

	user.Fullname = req.Fullname
	user.AvatarURL = req.AvatarURL

	if err := s.user.UpdateUser(user); err != nil {
		return customErr.NewInternalServerError("failed to update user profile", err)
	}

	return nil
}

func (s *userService) GetAllUsers(params dto.UserQueryParams) ([]dto.UserListResponse, *dto.PaginationResponse, error) {
	users, total, err := s.user.GetAllUsers(params)
	if err != nil {
		return nil, nil, customErr.NewInternalServerError("failed to fetch user list", err)
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

	pagination := utils.Paginate(total, params.Page, params.Limit)
	return results, pagination, nil
}

func (s *userService) GetUserDetail(id string) (*dto.UserDetailResponse, error) {
	user, err := s.user.GetUserByID(id)
	if err != nil || user == nil {
		return nil, customErr.NewNotFound("user not found").WithContext("userID", id)
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
