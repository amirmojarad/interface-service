package api

import (
	"interface_project/api/middlewares"

	"github.com/gin-gonic/gin"
)

func (api *API) userGroup(path string) {
	userGroup := api.Engine.Group(path)
	userGroup.GET("/", middlewares.CheckAuth(), api.deleteUser())
}

func (api *API) deleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
