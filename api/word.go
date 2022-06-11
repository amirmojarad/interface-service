package api

import (
	"github.com/gin-gonic/gin"
	"interface_project/api/middlewares"

	"interface_project/api/dto"
	"net/http"
	"strconv"
)

func (api API) wordGroup(path string) {
	group := api.Engine.Group(path, middlewares.CheckAuth())
	group.GET("/", api.getFileWords())
	group.POST("/", api.getAllFileWordsByTitle())
}

func (api API) getAllFileWordsByTitle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var requestSchema dto.GetWordSentencesRequest
		if err := ctx.BindJSON(&requestSchema); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "invalid json schema",
				"error":   err.Error(),
			},
			)
		}
		words, err := api.Crud.GetAllWordsByTitle(requestSchema.FileID, requestSchema.Title)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "error occurred while fetching words from database",
				"error":   err.Error(),
			})
			return
		}
		ctx.IndentedJSON(http.StatusOK, words)
	}
}

func (api API) getFileWords() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Request.URL.Query().Get("file_id"))
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "request does not contain any id in url",
				"error":   err.Error(),
			})
		}
		if words, err := api.Crud.GetAllWordsByFileID(id); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "while fetching words from database error occurred",
				"error":   err.Error(),
			})
		} else {
			ctx.IndentedJSON(http.StatusOK, words)
		}
	}
}
