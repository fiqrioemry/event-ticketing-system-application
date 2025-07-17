package handlers

import (
	"fmt"
	"log"

	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"

	"github.com/fiqrioemry/event_ticketing_system_app/server/services"

	"github.com/fiqrioemry/go-api-toolkit/pagination"
	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetMyProfile(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	userProfile, err := h.service.GetUserProfile(userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Profile retrieved successfully", userProfile)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	var req dto.UpdateProfileRequest

	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	log.Println("Updating profile for user:", req.Avatar)
	fmt.Println("Updating profile for user:", req.Avatar)
	if req.Avatar != nil && req.Avatar.Filename != "" {
		imageURL, err := utils.UploadImageWithValidation(req.Avatar)
		if err != nil {
			response.Error(c, err)
			return
		}
		req.AvatarURL = imageURL
	}

	updatedProfile, err := h.service.UpdateUserDetail(userID, &req)
	if err != nil {
		utils.CleanupImageOnError(req.AvatarURL)
		response.Error(c, err)
		return
	}
	response.OK(c, "Profile updated successfully", updatedProfile)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var params dto.UserQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	if err := pagination.BindAndSetDefaults(c, &params); err != nil {
		response.Error(c, response.BadRequest(err.Error()))
		return
	}

	users, total, err := h.service.GetAllUsers(params)
	if err != nil {
		response.Error(c, err)
		return
	}

	pagination := pagination.Build(params.Page, params.Limit, total)

	response.OKWithPagination(c, "Users retrieved successfully", users, pagination)
}

func (h *UserHandler) GetUserDetail(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUserDetail(id)

	if err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "User details retrieved successfully", user)
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	userID := utils.MustGetUserID(c)

	var req dto.ChangePasswordRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.ChangePassword(userID, &req); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Password changed successfully", nil)
}
