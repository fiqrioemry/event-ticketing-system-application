package handlers

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/repositories"
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"

	"github.com/fiqrioemry/event_ticketing_system_app/server/services"

	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service    services.UserService
	repository repositories.AuditLogRepository
}

func NewUserHandler(service services.UserService, repository repositories.AuditLogRepository) *UserHandler {
	return &UserHandler{service, repository}
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

	// record audit log
	auditLog := utils.BuildAuditLog(c, userID, "update", "profile", updatedProfile)

	go h.repository.Create(c.Request.Context(), auditLog)

	response.OK(c, "Profile updated successfully", updatedProfile)
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
