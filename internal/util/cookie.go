package util

import "github.com/gin-gonic/gin"

const (
	refreshTokenCookieName = "refresh_token"
	refreshTokenMaxAge     = 7 * 24 * 60 * 60
)

func SetRefreshTokenCookie(ctx *gin.Context, token string) {
	ctx.SetCookie(
		refreshTokenCookieName,
		token,
		refreshTokenMaxAge,
		"/",
		"",    // domain — empty for localhost
		false, // secure — set true in production
		true,  // httpOnly
	)
}

func ClearRefreshTokenCookie(ctx *gin.Context) {
	ctx.SetCookie(
		refreshTokenCookieName,
		"",
		-1, // expire immediately
		"/",
		"",
		false,
		true,
	)
}
