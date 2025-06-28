package handlers

import (
	"net/http"
	"server/dto"
	"server/services"
	"server/utils"

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

	response, err := h.service.GetUserProfile(userID)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := utils.MustGetUserID(c)
	var req dto.UpdateProfileRequest

	if !utils.BindAndValidateForm(c, &req) {
		return
	}

	avatarURL, err := utils.UploadImageWithValidation(req.Avatar)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	req.AvatarURL = avatarURL

	if err := h.service.UpdateUserDetail(userID, &req); err != nil {
		utils.CleanupImageOnError(avatarURL)
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var params dto.UserQueryParams
	if !utils.BindAndValidateForm(c, &params) {
		return
	}

	users, pagination, err := h.service.GetAllUsers(params)
	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":       users,
		"pagination": pagination,
	})
}

func (h *UserHandler) GetUserDetail(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUserDetail(id)

	if err != nil {
		utils.HandleServiceError(c, err, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}
