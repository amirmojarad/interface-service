package api

import (
	"interface_project/ent"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) userGroup(path string) {
	userGroup := api.Engine.Group(path)
	userGroup.POST("/signup", api.signUp())
	userGroup.POST("/login", api.login())
	userGroup.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "ASDASD")
	})
}

func (api *API) signUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userSchema := &ent.User{}
		ctx.BindJSON(&userSchema)
		log.Printf("UserSchema: %+v", userSchema)
		if newUser, err := api.Crud.AddUser(userSchema); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		} else {
			ctx.IndentedJSON(http.StatusCreated, newUser)
		}
	}
}

func (api *API) login() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
