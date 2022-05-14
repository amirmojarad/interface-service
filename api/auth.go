package api

import (
	"interface_project/api/auth"
	"interface_project/usecases/generators"

	"interface_project/ent"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) authGroup(path string) {
	userGroup := api.Engine.Group(path)
	userGroup.POST("/signup", api.signUp())
	userGroup.POST("/login", api.login())
}

func (api *API) signUp() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userSchema := &ent.User{}
		ctx.BindJSON(&userSchema)
		log.Printf("UserSchema: %+v", userSchema)
		if newUser, err := api.Crud.AddUser(userSchema); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, err.Error())
		} else {
			token := api.jwtService.GenerateToken(userSchema.Email, userSchema.IsAdmin)
			ctx.IndentedJSON(http.StatusCreated, gin.H{
				"user":  newUser,
				"token": token,
			})
		}
	}
}

func (api *API) login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userSchema := &ent.User{}
		ctx.BindJSON(&userSchema)
		if fetchedUser, err := api.Crud.GetUserByEmail(userSchema.Email); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "invalid credentials",
				"error":   err.Error(),
			})
			return
		} else {
			if !generators.CheckPasswordHash(userSchema.Password, fetchedUser.Password) {
				ctx.IndentedJSON(http.StatusBadRequest, gin.H{
					"message": "wrong password",
				})
				return
			}
			jwtService := auth.JWTAuthService()
			tokenString := jwtService.GenerateToken(userSchema.Email, userSchema.IsAdmin)
			ctx.IndentedJSON(http.StatusOK, gin.H{
				"token": tokenString,
			})
			return
		}
	}
}
