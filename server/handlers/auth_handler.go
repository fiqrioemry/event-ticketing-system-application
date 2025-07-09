package handlers

import (
	"net/http"
	"server/dto"
	"server/services"
	"server/utils"

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
		utils.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent to email"})
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.Register(&req); err != nil {
		utils.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent to your email"})
}

func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	tokens, err := h.service.VerifyOTP(req.Email, req.OTP)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.SetAccessTokenCookie(c, tokens.AccessToken)
	utils.SetRefreshTokenCookie(c, tokens.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	response, err := h.service.Login(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}

	utils.SetAccessTokenCookie(c, response.AccessToken)
	utils.SetRefreshTokenCookie(c, response.RefreshToken)

	c.JSON(http.StatusOK, gin.H{"message": "Login successfully", "user": response.User})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	utils.ClearAccessTokenCookie(c)
	utils.ClearRefreshTokenCookie(c)
	c.JSON(http.StatusOK, gin.H{"message": "Logout successfully"})
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

	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})
}
