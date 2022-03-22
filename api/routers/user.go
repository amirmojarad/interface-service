package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router Router) UsersRouter() {
	// localhost:8080/users
	user := router.Engine.Group(router.UserPath)
	user.GET("/", router.getDefaultPath())
}

func (router Router) getDefaultPath() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"status": "user router",
		})
	}
}
