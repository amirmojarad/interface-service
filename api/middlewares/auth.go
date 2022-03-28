package middlewares

import (
	"interface_project/api/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// CheckAuth get token from payload and
// check is it vald or not.
func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtService := auth.JWTAuthService()
		header := ctx.Request.Header
		token := strings.Split(header["Authorization"][0], " ")[1]
		if _, err := jwtService.ValidateToken(token); err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		} else {
			ctx.Next()
		}
	}
}

func IsSuperUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtService := auth.JWTAuthService()
		header := ctx.Request.Header
		token := strings.Split(header["Authorization"][0], " ")[1]
		validatedToken, _ := jwtService.ValidateToken(token)
		jwtClaims := jwtService.GetMapClaims(validatedToken)
		if jwtClaims["isAdmin"] == true {
			ctx.Set("isAdmin", true)
		} else {
			ctx.Set("isAdmin", false)
		}
		ctx.Next()
	}
}
