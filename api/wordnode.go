package api

import (
	"interface_project/api/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api API) WordNodesGroup(path string) {
	router := api.Engine.Group(path)
	router.POST("/all", api.GetMovieWordNodes())
	router.POST("/all/sort", api.GetMovieWordsNodeOrderBy())
}

func (api API) 	GetMovieWordNodes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var movieWordNodeSchema dto.MovieWordNode
		if err := ctx.BindJSON(&movieWordNodeSchema); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": "send valid json schema",
			})
			return
		}
		if wordNodes, err := api.Crud.GetMovieWordnodes(movieWordNodeSchema.MovieID); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"error": "error while fetching wordnodes from database",
			})

		} else {
			ctx.IndentedJSON(http.StatusBadRequest, wordNodes)
		}
	}
}

func (api API) GetMovieWordsNodeOrderBy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var movieWordNodeOrder dto.MovieWordNodeOrderBy
		if err := ctx.BindJSON(&movieWordNodeOrder); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "send valid json schema",
				"error":   err.Error(),
			})
		}
		if movieWordNodeOrder.SortByID {
			if wordNodes, err := api.Crud.SortByID(movieWordNodeOrder.MovieID); err != nil {
				ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
					"message": "error while fetching wordnodes from database in id sort",
					"error":   err.Error(),
				})
			} else {
				ctx.IndentedJSON(http.StatusOK, wordNodes)
			}
		} else if movieWordNodeOrder.SortByPreposition {
			if wordNodes, err := api.Crud.SortByPreposition(movieWordNodeOrder.MovieID); err != nil {
				ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
					"message": "error while fetching wordnodes from database in preposition sort",
					"error":   err.Error(),
				})
			} else {
				ctx.IndentedJSON(http.StatusInternalServerError, wordNodes)
			}
		} else if movieWordNodeOrder.SortByOccurence {
			if wordNodes, err := api.Crud.SortByPreposition(movieWordNodeOrder.MovieID); err != nil {
				ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
					"message": "error while fetching wordnodes from database in occurence sort",
					"error":   err.Error(),
				})
			} else {
				ctx.IndentedJSON(http.StatusInternalServerError, wordNodes)
			}
		}
	}
}
