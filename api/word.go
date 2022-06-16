package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"interface_project/api/middlewares"

	"interface_project/api/dto"
	"net/http"
	"strconv"
)

func (api API) wordGroup(path string) {
	group := api.Engine.Group(path, middlewares.CheckAuth())
	group.GET("/", api.getFileWordsPage())
	group.POST("/", api.getAllFileWordsByTitle())
}

const PER_PAGE = 24

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

func (api API) getFileWordsPage() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pageNumber, err := strconv.Atoi(ctx.Request.URL.Query().Get("page_number"))
		id, err := strconv.Atoi(ctx.Request.URL.Query().Get("file_id"))
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "request does not contain any id in url",
				"error":   err.Error(),
			})
		}
		if err != nil {
			pageNumber = 1
		}
		if words, err := api.Crud.GetAllWordsByFileID(id); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "while fetching words from database error occurred",
				"error":   err.Error(),
			})
		} else {
			limit := pageNumber * PER_PAGE
			start := limit - PER_PAGE
			if start < 0 {
				start = 0
			}
			if limit > len(words) {
				limit = len(words) - 1
			}
			if pageNumber > len(words)/PER_PAGE {
				pageNumber = len(words) / PER_PAGE
			}
			response := dto.WordSentencesResponse{
				Page:       pageNumber,
				PageCount:  len(words[start:limit]),
				PerPage:    PER_PAGE,
				TotalCount: len(words),
				Links: map[string]interface{}{
					"self":  fmt.Sprintf("/words/?file_id=%d&page=%d", id, pageNumber),
					"first": fmt.Sprintf("/words/?file_id=%d&page=1", id),
					"last":  fmt.Sprintf("/words/?file_id=%d&page=%d", id, len(words)/PER_PAGE),
					"next":  fmt.Sprintf("/words/?file_id=%d&page=%d", id, pageNumber+1),
					"prev":  fmt.Sprintf("/words/?file_id=%d&page=%d", id, pageNumber-1),
				},
				Records: words[start:limit],
			}

			ctx.IndentedJSON(http.StatusOK, response)
		}
	}
}
