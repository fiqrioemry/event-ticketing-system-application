package handlers

import (
	"server/dto"
	"server/services"
	"server/utils"

	"github.com/fiqrioemry/go-api-toolkit/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

func (h *AuthHandler) ResendOTP(c *gin.Context) {
	var req dto.ResendOTPRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.ResendOTP(req.Email); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "OTP sent to email successfully", nil)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.Register(&req); err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, "OTP sent to email successfully", nil)
}

func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	tokens, err := h.service.VerifyOTP(req.Email, req.OTP)
	if err != nil {
		response.Error(c, err)
		return
	}

	utils.SetAccessTokenCookie(c, tokens.AccessToken)
	utils.SetRefreshTokenCookie(c, tokens.RefreshToken)

	response.OK(c, "OTP verified successfully", tokens)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	result, err := h.service.Login(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.SetAccessTokenCookie(c, result.AccessToken)
	utils.SetRefreshTokenCookie(c, result.RefreshToken)

	response.OK(c, "Login successfully", result.User)

}

func (h *AuthHandler) Logout(c *gin.Context) {
	utils.ClearAccessTokenCookie(c)
	utils.ClearRefreshTokenCookie(c)
	response.OK(c, "Logout successfully", nil)
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	accessToken, err := h.service.RefreshToken(c, refreshToken)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.SetAccessTokenCookie(c, accessToken)

	response.OK(c, "Token refreshed successfully", accessToken)

}
