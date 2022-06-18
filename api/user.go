package api

import (
	"fmt"
	"interface_project/api/middlewares"
	"interface_project/ent"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api *API) userGroup(path string) {
	userGroup := api.Engine.Group(path, middlewares.CheckAuth())
	// Favorite Movies
	userGroup.POST("/favoriteMovies", api.addMoviesToFavorites())
	userGroup.GET("/favoriteMovies", api.getFavoritesMovies())
	userGroup.GET("/favoriteMovies/:id", api.getFavoriteMovie())
	userGroup.DELETE("/favoriteMovies", api.deleteMovieFromFavorites())

	// Crud user
	userGroup.PATCH("/", api.changeUser())
	userGroup.POST("/", api.updateUser())
	userGroup.GET("/", api.getUserByID())

	// Super User
	userGroup.DELETE("/admin", middlewares.IsSuperUser(), api.deleteUser())
	userGroup.GET("/admin", middlewares.IsSuperUser(), api.getAllUsers())

	// Keywords
	userGroup.GET("/searchKeywords", api.getSearchKeywords())

	// Upload subtitle
	userGroup.POST("/upload", api.sendSubtitleText())

	// FavoriteWords
	userGroup.POST("/favorite_words", api.addWordsToUser())

	// User Files
	userGroup.GET("/files/all", api.getAllFiles())

}

func (api API) addWordsToUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := fmt.Sprint(ctx.MustGet("email"))
		words := []string{}
		if err := ctx.BindJSON(&words); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "invalid json schema",
				"error":   err.Error(),
			})
			return
		}
		favoriteWords, err := api.Crud.AddFavoriteWordsToUser(words, email)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "error while adding words to user",
				"error":   err.Error(),
			})
			return
		}
		ctx.IndentedJSON(http.StatusOK, favoriteWords)
	}
}

func (api API) updateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user ent.User
		ctx.BindJSON(&user)
		log.Println(user)
		if fetchedUser, err := api.Crud.GetUserByID(user.ID); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "error",
				"error":   err.Error(),
				"user":    user,
			})
			return
		} else {
			email := fmt.Sprint(ctx.MustGet("email"))
			if email == fetchedUser.Email {
				if updatedUser, err := api.Crud.UpdateUser(user.ID, &user); err != nil {
					ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
						"message": "error while updating user",
						"error":   err.Error(),
					})

				} else {
					ctx.IndentedJSON(http.StatusOK, updatedUser)
				}
			} else {
				ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
					"message": fmt.Sprintf("user with id %d does not accessible", user.ID),
				})
			}
			return
		}
	}
}

func (api API) getUserByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if id, err := strconv.Atoi(ctx.Request.URL.Query().Get("id")); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "request does not contain any id in url",
				"error":   err.Error(),
			})
		} else {
			if fetchedUser, err := api.Crud.GetUserByID(id); err != nil {

				ctx.IndentedJSON(http.StatusNotFound, gin.H{
					"message": "cannot find user by given id",
					"error":   err.Error(),
				})
			} else {
				email := fmt.Sprint(ctx.MustGet("email"))
				if fetchedUser.Email != email {
					ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
						"message": "given token does not match to given id",
					})
					return
				}
				ctx.IndentedJSON(http.StatusOK, gin.H{
					"user": fetchedUser,
				})
			}
		}
	}
}

func (api API) getSearchKeywords() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := fmt.Sprint(ctx.MustGet("email"))
		if keywords, err := api.Crud.GetUserSearchKeyword(email); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "some error happend", "error": err.Error()})
		} else {
			searchTitles := []string{}
			for _, keyword := range keywords {
				searchTitles = append(searchTitles, keyword.Title)
			}
			ctx.IndentedJSON(http.StatusOK, searchTitles)
		}
	}
}

func (api *API) deleteMovieFromFavorites() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userEmail := fmt.Sprint(ctx.MustGet("email"))
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
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("movie with id %d not found", movieID), "error": err.Error()})
		} else {
			ctx.IndentedJSON(http.StatusOK, movie)
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
