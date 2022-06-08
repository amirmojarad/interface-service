package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (api API) getAllWordnodes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if wordnodes, err := api.Crud.GetAllWordNodes(); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "error occured while fetching data from database",
				"error":   err.Error(),
			})
		} else {
			ctx.IndentedJSON(http.StatusOK, wordnodes)
		}
	}
}

func (api API) WordNodesGroup(path string) {
	router := api.Engine.Group(path)
	router.POST("/all", api.getFileWordNodes())
	// router.POST("/all/sort", api.GetMovieWordsNodeOrderBy()s)
}

func (api API) getFileWordNodes() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if fileIDasString := ctx.Request.URL.Query().Get("file_id"); fileIDasString == "" {
			return
		} else {
			fileID, err := strconv.Atoi(fileIDasString)
			if err != nil {
				ctx.IndentedJSON(http.StatusBadRequest, gin.H{
					"message": "please send a valid file id",
					"error":   err.Error(),
				})
				return
			}
			wordnodes, err := api.Crud.GetFileWordNodes(fileID)
			if err != nil {
				ctx.IndentedJSON(http.StatusNotFound, gin.H{
					"message": "error while fetching file word nodes from database",
					"error":   err.Error(),
				})
				return
			}
			ctx.IndentedJSON(http.StatusOK, wordnodes)
		}
	}
}
