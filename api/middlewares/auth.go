package middlewares

import (
	"interface_project/api/auth"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwtService := auth.JWTAuthService()
		header := ctx.Request.Header
		token := strings.Split(header["Authorization"][0], " ")[1]
		log.Println(token)
		if _, err := jwtService.ValidateToken(token); err != nil {
			log.Println("ERROR ", err)
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		} else {
			ctx.Next()
		}
	}
}
