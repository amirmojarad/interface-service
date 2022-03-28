package api

import (
	"interface_project/api/middlewares"
	"interface_project/ent"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// -------- helper middlewares

func (api *API) userGroup(path string) {
	userGroup := api.Engine.Group(path, middlewares.CheckAuth())
	userGroup.DELETE("/", middlewares.IsSuperUser(), api.deleteUser())
	userGroup.PATCH("/", api.changeUser())
	userGroup.POST("/", middlewares.IsSuperUser(), api.getAllUsers())
	userGroup.POST("/favoriteMovies", api.addMovieToFavorites())
	userGroup.GET("/favoriteMovies", api.getFavoritesMovies())
	userGroup.GET("/favoriteMovies/:id", api.getFavoriteMovie())
	userGroup.DELETE("/favoriteMovies", api.deleteMovieFromFavorites())
	userGroup.POST("/search", api.searchMovie())
}

func (api *API) searchMovie() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
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
		userSchema := &ent.User{}
		header := ctx.Request.Header
		token := strings.Split(header["Authorization"][0], " ")[1]
		validatedToken, _ := api.jwtService.ValidateToken(token)
		jwtClaims := api.jwtService.GetMapClaims(validatedToken)
		if userSchema.Email == jwtClaims["email"] && jwtClaims["isAdmin"] == true {
			if deletedUser, err := api.Crud.DeleteUserByEmail(userSchema.Email); err != nil {
				ctx.IndentedJSON(http.StatusInternalServerError, err)
			} else {
				ctx.IndentedJSON(http.StatusAccepted, deletedUser)
			}
		} else {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "user not authorizated",
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
