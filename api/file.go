package api

import (
	"fmt"
	"interface_project/api/dto"
	"interface_project/api/middlewares"
	"interface_project/subs"

	"interface_project/usecases/handlers/sentences"
	"io"

	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func (api *API) fileGroup(path string) {
	fileGroup := api.Engine.Group(path)
	fileGroup.GET("/", middlewares.CheckAuth(), api.getAllFiles())
	fileGroup.POST("/upload", middlewares.CheckAuth(), api.uploadSubtitle())
	fileGroup.POST("/upload_profile", middlewares.CheckAuth(), api.uploadImage())
	fileGroup.DELETE("/", middlewares.CheckAuth(), api.deleteFile())
	fileGroup.GET("/download", api.download())
}

func (api API) getAllFiles() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := fmt.Sprint(ctx.MustGet("email"))
		user, err := api.Crud.GetUserByEmail(email)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "user not found in database",
				"error":   err.Error(),
			})
			return
		}
		files, err := api.Crud.GetAllFiles(user)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "error occurred while fetching files from database",
				"error":   err.Error(),
			})
			return
		}
		ctx.IndentedJSON(http.StatusOK, files)

	}
}

func (api API) download() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if image_url := ctx.Request.URL.Query().Get("image_url"); image_url == "" {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "request does not contain any image_url in url",
			})
		} else {
			log.Println(image_url)
			ctx.File(image_url)
		}
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

func (api API) uploadImage() gin.HandlerFunc {
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
		if !(strings.Contains(fileName, ".jpeg") || strings.Contains(fileName, ".png") || strings.Contains(fileName, ".jpg")) {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "uploaded file is not an image file.",
			})
			log.Println("uploaded file is not an image file.")
			return
		}
		user, err := api.Crud.GetUserByEmail(email)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "user not found",
				"error":   err.Error(),
			})
			return
		}
		folderPath := "images/" + email + "/"
		filePath := "images/" + email + "/" + fileName
		subs.MakeDir(folderPath)
		out, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		subs.Open(filePath)
		defer out.Close()
		written, err := io.Copy(out, file)
		log.Println(written)
		if err != nil {
			log.Fatal(err)
		}
		updatedUser, err := api.Crud.AddImageUrlToUser(user, filePath)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "error occured when file entity created in database",
				"error":   err,
			})
		}
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message":   "file uploaded successfuly",
			"image_url": updatedUser.ImageURL,
		})
	}
}

func (api API) deleteFile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var idList []int
		ctx.BindJSON(&idList)
		email := fmt.Sprint(ctx.MustGet("email"))
		user, err := api.Crud.GetUserByEmail(email)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{
				"message": "user not found in databse",
				"error":   err.Error(),
			})
			return
		}
		files, err := api.Crud.GetFiles(user, idList)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{
				"message": "files not found in databse",
				"error":   err.Error(),
			})
			return
		}
		if err := api.Crud.DeleteFiles(user, idList); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": "error while deleting file entity from user entity in databse",
				"error":   err.Error(),
			})
			return
		}
		for _, file := range files {
			if err := subs.DeleteFile(file.Path); err != nil {
				ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("error while deleting file with path %s from directory", file.Path),
					"error":   err.Error(),
				})
				return
			}
		}
		ctx.IndentedJSON(http.StatusOK, gin.H{
			"message": "all files deleted successfully",
			"files":   files,
		})
	}
}

func (api API) uploadSubtitle() gin.HandlerFunc {
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
		folderPath := "subs/" + email + "/"
		filePath := folderPath + fileName
		err = subs.MakeDir(folderPath)
		if err != nil {
			return
		}
		//check file with its description exists in database or not
		if api.Crud.CheckFileIsExists(email, filePath, fileName) {
			ctx.IndentedJSON(http.StatusConflict, gin.H{
				"message": "file with name " + fileName + " and path " + filePath + " and user with email " + email + " already is exists",
			})
			log.Println("file with name " + fileName + " and path " + filePath + " and user with email " + email + " already is exists")
			return
		}
		out, err := os.Create(filePath)
		if err != nil {
			log.Fatal(err)
		}
		openedFile, _ := subs.Open(filePath)
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
			wordCreateBulk := sentences.GetSentences(api.Crud.Client, openedFile, user, createdFile)
			bucketSize := len(wordCreateBulk) / 65535
			if bucketSize < 1 {
				_, err = api.Crud.CreateAllWords(wordCreateBulk)
				if err != nil {
					ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
						"error":   err.Error(),
						"message": "error while creating words to database",
					})
					log.Println(bucketSize)
					log.Println(len(wordCreateBulk))
					return
				}
			} else {
				_, err = api.Crud.CreateAllWords(wordCreateBulk[0:65535])
				if err != nil {
					ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
						"error":   err.Error(),
						"message": "error while creating words to database",
					})
					return
				}
				_, err = api.Crud.CreateAllWords(wordCreateBulk[65535:])
				if err != nil {
					ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
						"error":   err.Error(),
						"message": "error while creating words to database",
					})
					return
				}
			}
			ctx.IndentedJSON(http.StatusOK, gin.H{
				"message": "file downloaded successfully",
				"file":    createdFile,
			})
		}
	}

}
