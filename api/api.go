package api

import (
	"context"
	"interface_project/api/auth"
	"interface_project/database/crud"
	"interface_project/ent"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	Crud       *crud.Crud
	Engine     *gin.Engine
	jwtService auth.JWTService
}

func (api *API) startEngine() {
	api.groups()
	api.Engine.LoadHTMLGlob("template/*")
	api.Engine.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
}

func (api *API) groups() {
	api.userGroup("/users")
	api.movieGroup("/movies")
	api.fileGroup("/file")
	api.authGroup("/")
}

func RunAPI(ctx *context.Context, client *ent.Client) {
	interfaceAPI := &API{
		Engine: gin.Default(),
		Crud: &crud.Crud{
			Ctx:    ctx,
			Client: client,
		},
		jwtService: auth.JWTAuthService(),
	}
	interfaceAPI.startEngine()
	interfaceAPI.Engine.Run(":8080")
}
