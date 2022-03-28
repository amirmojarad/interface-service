package api

import (
	"encoding/json"
	"fmt"
	"interface_project/ent"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func (api *API) movieGroup(path string) {
	movieGroup := api.Engine.Group(path)
	movieGroup.POST("/:title", api.addMovies())
	movieGroup.GET("/", api.getAllMovies())

}

func (api *API) getAllMovies() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if movies, err := api.Crud.GetAllMovies(); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, movies)
		}
	}

}

func (api *API) addMovies() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		title := ctx.Param("title")
		if imdbMovies, err := queryMovieFromIMDB(title); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "error occured when fetching data from imdb",
			})
		} else {
			movies := make([]*ent.MovieCreate, len(imdbMovies))
			for i, imdbMovie := range imdbMovies {
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
			if newMovies, err := api.Crud.AddMovies(movies); err != nil {
				ctx.IndentedJSON(http.StatusInternalServerError, err)
			} else {
				ctx.IndentedJSON(http.StatusCreated, newMovies)
			}
		}

	}

}

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

func queryMovieFromIMDB(title string) ([]*imdbMovie, error) {
	apiKey := os.Getenv("API_KEY")
	if response, err := http.Get(fmt.Sprintf("%s/%s?title=%s&count=100", imdbPathURL, apiKey, title)); err != nil {
		return nil, err
	} else {
		log.Println(fmt.Sprintf("%s/%s?title=%s&count=250", imdbPathURL, apiKey, title))
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
