package api

import (
	"fmt"
	"interface_project/api/middlewares"
	"interface_project/ent"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *API) userGroup(path string) {
	userGroup := api.Engine.Group(path, middlewares.CheckAuth())
	userGroup.DELETE("/", middlewares.IsSuperUser(), api.deleteUser())
	userGroup.PATCH("/", api.changeUser())
	userGroup.GET("/", middlewares.IsSuperUser(), api.getAllUsers())
	userGroup.POST("/favoriteMovies", api.addMoviesToFavorites())
	userGroup.GET("/favoriteMovies", api.getFavoritesMovies())
	userGroup.GET("/favoriteMovies/:id", api.getFavoriteMovie())
	userGroup.DELETE("/favoriteMovies", api.deleteMovieFromFavorites())
	userGroup.GET("/searchKeywords", api.getSearchKeywords())
	userGroup.GET("/upload", func(ctx *gin.Context) {
		location := url.URL{Path: "/file"}
		ctx.Redirect(http.StatusFound, location.RequestURI())
	})
}

func (api API) userIndex() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// ctx.HTML(http.StatusOK, )
	}
}

func (api API) getSearchKeywords() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := fmt.Sprint(ctx.MustGet("email"))
		if keywords, err := api.Crud.GetUserSearchKeyword(email); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "some error happend", "error": err.Error()})
		} else {
			ctx.IndentedJSON(http.StatusOK, keywords)
		}
	}
}

func (api *API) deleteMovieFromFavorites() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userEmail := fmt.Sprint(ctx.Query("email"))
		var movieIDs []int
		ctx.BindJSON(&movieIDs)
		if movies, err := api.Crud.DeleteMovieFromFavorites(userEmail, movieIDs); err != nil {
			ctx.IndentedJSON(http.StatusServiceUnavailable, gin.H{
				"message": "could not delete movies.",
				"error":   err.Error(),
			})
		} else {
			ctx.IndentedJSON(http.StatusAccepted, movies)
		}
	}
}

func (api *API) getFavoritesMovies() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userEmail := fmt.Sprint(ctx.MustGet("email"))
		if movies, err := api.Crud.GetFavoriteMovies(userEmail); err != nil {
			ctx.IndentedJSON(http.StatusServiceUnavailable, gin.H{
				"message": "could not get movies.",
				"error":   err.Error(),
			})
		} else {
			ctx.IndentedJSON(http.StatusOK, movies)
		}
	}
}

func (api *API) getFavoriteMovie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userEmail := ctx.MustGet("email")
		movieID, _ := strconv.Atoi(ctx.Param("id"))
		if movie, err := api.Crud.GetFavoriteMovie(fmt.Sprint(userEmail), movieID); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("movie with id %d not found", movieID), "error": err.Error()})
		} else {
			ctx.IndentedJSON(http.StatusCreated, movie)
		}
	}
}

func (api *API) addMoviesToFavorites() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userEmail := ctx.MustGet("email")
		var movieIDs []int
		ctx.BindJSON(&movieIDs)
		if movies, err := api.Crud.AddMoviesToUser(movieIDs, fmt.Sprint(userEmail)); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "could not add movies to user.", "error": err.Error(),
			})
		} else {
			ctx.IndentedJSON(http.StatusCreated, movies)
		}
	}
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
