package handlers

import (
	"net/http"

	"github.com/fiqrioemry/event_ticketing_system_app/server/config"
	"github.com/fiqrioemry/event_ticketing_system_app/server/dto"
	"github.com/fiqrioemry/event_ticketing_system_app/server/services"
	"github.com/fiqrioemry/event_ticketing_system_app/server/utils"
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
	resp, err := h.service.VerifyOTP(req.Email, req.OTP)
	if err != nil {
		response.Error(c, err)
		return
	}
	// set cookies as httpOnly
	utils.SetAccessTokenCookie(c, resp.AccessToken)
	utils.SetRefreshTokenCookie(c, resp.RefreshToken)

	response.OK(c, "OTP verified successfully", resp.User)
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

	// set cookies as httpOnly
	utils.SetAccessTokenCookie(c, result.AccessToken)
	utils.SetRefreshTokenCookie(c, result.RefreshToken)

	response.OK(c, "Login successfully", result.User)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	// clear cookies
	utils.ClearAccessTokenCookie(c)
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
	user, token, err := h.service.RefreshToken(c, refreshToken)
	if err != nil {
		response.Error(c, err)
		return
	}

	// set access cookie
	utils.SetAccessTokenCookie(c, token)

	response.OK(c, "Token refreshed successfully", user)
}

// step 1 : User requests password reset
func (h *AuthHandler) ForgotPassword(c *gin.Context) {

	var req dto.ForgotPasswordRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.ForgotPassword(c, &req); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Password reset link sent successfully", nil)

}

// step 2 : validate reset token and reset password
func (h *AuthHandler) ValidateResetToken(c *gin.Context) {
	token := c.Query("token")

	email, err := h.service.ValidateToken(token)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Reset token is valid", email)
}

// step 3 : reset password
func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req dto.ResetPasswordRequest
	if !utils.BindAndValidateJSON(c, &req) {
		return
	}

	if err := h.service.ResetPassword(&req); err != nil {
		response.Error(c, err)
		return
	}

	response.OK(c, "Password has been reset successfully", nil)
}

func (h *AuthHandler) GoogleOAuthRedirect(c *gin.Context) {
	url := h.service.GetGoogleOAuthURL()
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) GoogleOAuthCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		response.Error(c, response.NewBadRequest("Authorization code is missing"))
		return
	}

	tokens, err := h.service.HandleGoogleOAuthCallback(code)
	if err != nil {
		response.Error(c, err)
		return
	}

	utils.SetAccessTokenCookie(c, tokens.AccessToken)

	utils.SetRefreshTokenCookie(c, tokens.RefreshToken)

	c.Redirect(http.StatusTemporaryRedirect, config.AppConfig.FrontendRedirectURL)
}
