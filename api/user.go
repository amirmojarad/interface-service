package api

import (
	"interface_project/api/auth"
	"interface_project/api/middlewares"
	"interface_project/ent"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) userGroup(path string) {
	userGroup := api.Engine.Group(path)
	userGroup.POST("/signup", api.signUp())
	userGroup.POST("/login", api.login())
	// userGroup.GET("/", api.allUsers())
	userGroup.GET("/", middlewares.CheckAuth(), func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "ASDASD")
	})
}

func (api *API) allUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if users, err := api.Crud.GetAllUsers(); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, err)
			return
		} else {
			ctx.IndentedJSON(http.StatusOK, users)
			return
		}
	}
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
		userSchema := &ent.User{}
		ctx.BindJSON(&userSchema)
		if _, err := api.Crud.GetUserByEmail(userSchema.Email); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "invalid credentials",
				"error":   err.Error(),
			})
			return
		} else {
			jwtService := auth.JWTAuthService()
			tokenString := jwtService.GenerateToken(userSchema.Email, true)
			ctx.IndentedJSON(http.StatusOK, gin.H{
				"token": tokenString,
			})
			return
		}
	}
}
