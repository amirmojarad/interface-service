package api

import (
	"fmt"
	"interface_project/api/dto"
	"interface_project/api/middlewares"
	"interface_project/usecases/handlers/file_handler"
	"io"

	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (api *API) fileGroup(path string) {
	fileGroup := api.Engine.Group(path, middlewares.CheckAuth())
	fileGroup.GET("/", api.fileIndex())
	fileGroup.POST("/upload", api.upload())
	fileGroup.GET("/download", api.download())
}

func (api API) download() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO
	}
}

func (api API) fileIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "select_file.html", gin.H{})
	}
}

func (api API) sendSubtitleText() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// email := fmt.Sprint(ctx.MustGet("email"))
		var subtitleText dto.SubtitleText
		if err := ctx.BindJSON(&subtitleText); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "invalid json file",
				"error":   err.Error(),
			})
			return
		}
		if !strings.Contains(subtitleText.Title, ".srt") {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "send srt file only",
			})
			return
		}
		filePath := "subs/" + subtitleText.Title
		_, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		ctx.IndentedJSON(http.StatusCreated, gin.H{
			"message": fmt.Sprintf("file with path %s created successfully.", filePath),
		})
	}
}

func (api API) upload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := fmt.Sprint(ctx.MustGet("email"))
		file, header, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			log.Println(err.Error())
			return
		}
		fileName := header.Filename
		if !strings.Contains(fileName, ".srt") {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "uploaded file is not an srt file.",
			})
			log.Println("uploaded file is not an srt file.")
			return
		}
		filePath := "subs/" + fileName
		out, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		openedFile, _ := file_handler.Open(filePath)
		defer out.Close()
		written, err := io.Copy(out, file)
		log.Println(written)
		if err != nil {
			log.Fatal(err)
		}
		if user, err := api.Crud.GetUserByEmail(email); err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
		} else {
			createdFile, err := api.Crud.AddFileToUser(
				user, openedFile, filePath,
			)
			if err != nil {
				ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
					"message": "error occured when file entity created in database",
					"error":   err,
				})
			}
			ctx.IndentedJSON(http.StatusOK, gin.H{
				"message": "file downloaded successfuly",
				"file":    createdFile,
			})
		}
	}
}
