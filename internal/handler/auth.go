package handler

import (
	"booking-system/lvl1/internal/model"
	"booking-system/lvl1/internal/service"
	"booking-system/lvl1/internal/util"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *service.AuthService
}

func NewAuthHandler(userService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	loginResponse, refreshToken, err := h.userService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	util.SetRefreshTokenCookie(c, refreshToken)
	c.JSON(200, loginResponse)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	util.ClearRefreshTokenCookie(c)
	c.JSON(200, gin.H{"message": "Logged out successfully"})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(401, gin.H{"error": "Missing refresh token"})
		return
	}

	accessToken, refreshToken, err := h.userService.RefreshToken(cookie)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	util.SetRefreshTokenCookie(c, refreshToken)
	c.JSON(200, gin.H{
		"access_token": accessToken,
	})
}
