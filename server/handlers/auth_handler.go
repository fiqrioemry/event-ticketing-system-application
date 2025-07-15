package handlers

import (
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"

	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"

	"github.com/fiqrioemry/event_ticketing_system_app/server/services"

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
	// bind request data
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	// send OTP email
	if err := h.service.ResendOTP(req.Email); err != nil {
		response.Error(c, err)
		return
	}
	response.OK(c, "OTP sent to email successfully", nil)
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	// bind request data
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	// create user account
	if err := h.service.Register(&req); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "OTP sent to email successfully", nil)
}

func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	var req dto.VerifyOTPRequest
	// bind request data
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	// verify OTP code
	tokens, err := h.service.VerifyOTP(req.Email, req.OTP)
	if err != nil {
		response.Error(c, err)
		return
	}

	// set access cookie
	utils.SetAccessTokenCookie(c, tokens.AccessToken)

	// set refresh cookie
	utils.SetRefreshTokenCookie(c, tokens.RefreshToken)

	response.OK(c, "OTP verified successfully", tokens)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	// bind request data
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	// authenticate user credentials
	result, err := h.service.Login(&req)
	if err != nil {
		response.Error(c, err)
		return
	}

	// set access cookie
	utils.SetAccessTokenCookie(c, result.AccessToken)

	// set refresh cookie
	utils.SetRefreshTokenCookie(c, result.RefreshToken)

	response.OK(c, "Login successfully", result.User)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	// clear access cookie
	utils.ClearAccessTokenCookie(c)
	// clear refresh cookie
	utils.ClearRefreshTokenCookie(c)
	response.OK(c, "Logout successfully", nil)
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	// get refresh token
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		response.Error(c, err)
		return
	}

	// generate new token
	token, err := h.service.RefreshToken(c, refreshToken)
	if err != nil {
		response.Error(c, err)
		return
	}

	// set access cookie
	utils.SetAccessTokenCookie(c, token)

	accessToken := gin.H{
		"accessToken": token,
	}

	response.OK(c, "Token refreshed successfully", accessToken)
}
