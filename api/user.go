package api

import (
	"fmt"
	"interface_project/api/middlewares"
	"interface_project/ent"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// -------- helper middlewares

func (api *API) userGroup(path string) {
	userGroup := api.Engine.Group(path, middlewares.CheckAuth())
	userGroup.DELETE("/", middlewares.IsSuperUser(), api.deleteUser())
	userGroup.PATCH("/", api.changeUser())
	userGroup.GET("/", middlewares.IsSuperUser(), api.getAllUsers())
	userGroup.POST("/favoriteMovies", api.addMovieToFavorites())
	userGroup.GET("/favoriteMovies", api.getFavoritesMovies())
	userGroup.GET("/favoriteMovies/:id", api.getFavoriteMovie())
	userGroup.DELETE("/favoriteMovies", api.deleteMovieFromFavorites())
	userGroup.POST("/search", api.searchMovie())
}

func (api *API) searchMovie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	}
}

func (api *API) deleteMovieFromFavorites() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (api *API) getFavoritesMovies() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (api *API) getFavoriteMovie() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (api *API) addMovieToFavorites() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (api *API) changeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (api *API) deleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userSchema := ent.User{}
		ctx.BindJSON(&userSchema)
		isAdmin := ctx.MustGet("isAdmin")
		if isAdmin == true {
			if deletedUser, err := api.Crud.DeleteUserByEmail(userSchema.Email); err != nil {
				ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
					"error":   err.Error(),
					"message": fmt.Sprint("no user found with email", &userSchema.Email),
				})
			} else {
				log.Printf("DELETED USER: %+v", deletedUser)
				ctx.IndentedJSON(http.StatusAccepted, deletedUser)
			}
		} else {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "your not superuser.",
			})
		}
	}
}

func (api *API) getAllUsers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		isAdmin := ctx.MustGet("isAdmin")
		log.Println("IS ADMIIIIN: ", isAdmin)
		if isAdmin == true {
			if users, err := api.Crud.GetAllUsers(); err != nil {
				ctx.IndentedJSON(http.StatusInternalServerError, err)
			} else {
				ctx.IndentedJSON(http.StatusOK, users)
			}
		} else {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "your not superuser.",
			})
		}
	}
}
