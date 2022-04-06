package api

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (api *API) fileGroup(path string) {
	fileGroup := api.Engine.Group(path)
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

func (api API) upload() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file, header, err := ctx.Request.FormFile("file")
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		fileName := header.Filename
		if !strings.Contains(fileName, ".srt") {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "uploaded file is not an srt file.",
			})
			return
		}
		out, err := os.Create("subs/" + fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		written, err := io.Copy(out, file)
		log.Println(written)
		if err != nil {
			log.Fatal(err)
		}
		filePath := "http://localhost:8080/file/" + fileName
		ctx.IndentedJSON(http.StatusCreated, gin.H{
			"filePath": filePath,
		})
	}
}
