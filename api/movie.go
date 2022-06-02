package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"interface_project/api/dto"
	"interface_project/api/middlewares"
	"interface_project/ent"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (api *API) movieGroup(path string) {
	movieGroup := api.Engine.Group(path, middlewares.CheckAuth())
	movieGroup.POST("/search", api.searchMovies())
}

// searchMovies first search movies from local database, else there are no related movie,
// second search movies with given title from imdb api.
func (api API) searchMovies() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		searchMovieSchema := dto.SearchMovieSchema{}
		if err := ctx.BindJSON(&searchMovieSchema); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "invalid json schema",
				"error":   err.Error(),
			})
			return
		}
		// adding searched keyword to database
		email := ctx.GetString("email")
		_, err := api.Crud.AddSearchKeywordToUser(email, searchMovieSchema.Title)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "error while adding keywords database.",
				"error":   err.Error(),
			})
			return
		}
		if searchedMovies, err := api.Crud.SearchMovieSortByID(searchMovieSchema); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "error while fetching movies from database.",
				"error":   err.Error(),
			})
			return
		} else {
			// database has no movie with given movie title, so query movies from imdb api and add them to database.
			if len(searchedMovies) == 0 {
				if movieCreateSchemas, err := queryMovieFromIMDB(searchMovieSchema.Title); err != nil {
					ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
						"message": "error while fetching movies from IMDB API.",
						"error":   err.Error(),
					})
					return
				} else {
					if addedMovies, err := addSearchedImdbMoviesToDatabase(movieCreateSchemas, &api); err != nil {
						ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
							"message": "error while adding movies to database.",
							"error":   err.Error(),
						})
						return
					} else {
						ctx.IndentedJSON(http.StatusCreated, addedMovies)
						return
					}
				}
			} else {
				ctx.IndentedJSON(http.StatusOK, searchedMovies)
				return
			}
		}
	}
}

// func (api *API) queryMovies() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {

// 		movieTitleSchema := dto.SearchMovieSchema{}
// 		if err := ctx.BindJSON(&movieTitleSchema); err != nil {
// 			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
// 				"error":   err.Error(),
// 				"message": "invalid json schema",
// 			})
// 			return
// 		}
// 		log.Println(movieTitleSchema.Title)
// 		if movies, err := api.Crud.SearchMovie(movieTitleSchema.Title); err != nil {
// 			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
// 				"message": "some error!",
// 				"error":   err.Error(),
// 			})
// 		} else {
// 			if len(movies) == 0 {
// 				log.Println(len(movies))
// 				if imdbMovies, err := queryMovieFromIMDB(movieTitleSchema.Title); err != nil {
// 					ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
// 						"message": "error occured when fetching data from imdb",
// 					})
// 				} else {
// 					log.Printf("%+v\n", imdbMovies[0])
// 					movies := make([]*ent.MovieCreate, len(imdbMovies))
// 					for i, imdbMovie := range imdbMovies {
// 						movies[i] = api.Crud.Client.Movie.Create().
// 							SetGenres(imdbMovie.Genres).
// 							SetImDbRating(imdbMovie.ImDbRating).
// 							SetImageURL(imdbMovie.Image).
// 							SetPlot(imdbMovie.Plot).
// 							SetStars(imdbMovie.Stars).
// 							SetRuntimeStr(imdbMovie.RuntimeStr).
// 							SetTitle(imdbMovie.Title).
// 							SetYear(imdbMovie.Description).
// 							SetMetacriticRating(imdbMovie.MetacriticRating)
// 					}
// 					if newMovies, err := api.Crud.AddMovies(movies); err != nil {
// 						ctx.IndentedJSON(http.StatusInternalServerError, err)
// 					} else {
// 						ctx.IndentedJSON(http.StatusOK, newMovies)
// 					}
// 				}
// 			} else {
// 				ctx.IndentedJSON(http.StatusOK, movies)
// 			}
// 		}
// 	}
// }

// func (api *API) searchMovie() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		movieTitle := ctx.Query("title")
// 		email := fmt.Sprint(ctx.MustGet("email"))
// 		api.Crud.AddSearchKeywordToUser(email, movieTitle)
// 		log.Println("MOVIE TITLE: ", movieTitle)
// 		if movies, err := api.Crud.SearchMovie(movieTitle); err != nil {
// 			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
// 				"message": "some error!",
// 				"error":   err.Error(),
// 			})
// 		} else {
// 			ctx.IndentedJSON(http.StatusOK, movies)
// 		}
// 	}
// }

// func (api *API) getAllMovies() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		if movies, err := api.Crud.GetAllMovies(); err != nil {
// 			ctx.IndentedJSON(http.StatusInternalServerError, err)
// 		} else {
// 			ctx.IndentedJSON(http.StatusOK, movies)
// 		}
// 	}

// }

// func (api *API) addMovies() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		title := ctx.Param("title")
// 		if imdbMovies, err := queryMovieFromIMDB(title); err != nil {
// 			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
// 				"message": "error occured when fetching data from imdb",
// 			})
// 		} else {
// 			log.Printf("%+v\n", imdbMovies[0])
// 			movies := make([]*ent.MovieCreate, len(imdbMovies))
// 			for i, imdbMovie := range imdbMovies {
// 				movies[i] = api.Crud.Client.Movie.Create().
// 					SetGenres(imdbMovie.Genres).
// 					SetImDbRating(imdbMovie.ImDbRating).
// 					SetImageURL(imdbMovie.Image).
// 					SetPlot(imdbMovie.Plot).
// 					SetStars(imdbMovie.Stars).
// 					SetRuntimeStr(imdbMovie.RuntimeStr).
// 					SetTitle(imdbMovie.Title).
// 					SetYear(imdbMovie.Description).
// 					SetMetacriticRating(imdbMovie.MetacriticRating)
// 			}
// 			if newMovies, err := api.Crud.AddMovies(movies); err != nil {
// 				ctx.IndentedJSON(http.StatusInternalServerError, err)
// 			} else {
// 				ctx.IndentedJSON(http.StatusOK, newMovies)
// 			}
// 		}

// 	}

// }

// ------- helpers

const (
	imdbPathURL = "https://imdb-api.com/API/AdvancedSearch"
)

type imdbMovie struct {
	Image            string
	Title            string
	Description      string
	RuntimeStr       string
	Genres           string
	ImDbRating       string
	MetacriticRating string
	Plot             string
	Stars            string
}

type QueryResult struct {
	Results []*imdbMovie
}

func addSearchedImdbMoviesToDatabase(searchedMovies []*imdbMovie, api *API) ([]*ent.Movie, error) {
	movies := make([]*ent.MovieCreate, len(searchedMovies))
	for i, imdbMovie := range searchedMovies {
		movies[i] = api.Crud.Client.Movie.Create().
			SetGenres(imdbMovie.Genres).
			SetImDbRating(imdbMovie.ImDbRating).
			SetImageURL(imdbMovie.Image).
			SetPlot(imdbMovie.Plot).
			SetStars(imdbMovie.Stars).
			SetRuntimeStr(imdbMovie.RuntimeStr).
			SetTitle(imdbMovie.Title).
			SetYear(imdbMovie.Description).
			SetMetacriticRating(imdbMovie.MetacriticRating)
	}
	if addedMovies, err := api.Crud.AddMovies(movies); err != nil {
		return nil, errors.New("error while adding MovieCreate slices to database")
	} else {
		return addedMovies, nil
	}
}

func queryMovieFromIMDB(title string) ([]*imdbMovie, error) {
	apiKey := os.Getenv("API_KEY")
	if response, err := http.Get(fmt.Sprintf("%s/%s?title=%s&count=100", imdbPathURL, apiKey, title)); err != nil {
		return nil, err
	} else {
		log.Println(fmt.Sprintf("%s/%s?title=%s&count=250", imdbPathURL, apiKey, title))
		log.Println(response.StatusCode)
		log.Println(fmt.Sprintf("%+v", response.Body))
		defer response.Body.Close()
		decoder := json.NewDecoder(response.Body)
		var queryResult QueryResult
		err = decoder.Decode(&queryResult)
		if err != nil {
			return nil, err
		} else {
			return queryResult.Results, nil
		}
	}
}
