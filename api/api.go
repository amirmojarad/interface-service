package api

import (
	"context"
	"interface_project/api/auth"
	"interface_project/database/crud"
	"interface_project/ent"

	"github.com/gin-gonic/gin"
)

type API struct {
	Crud       *crud.Crud
	Engine     *gin.Engine
	jwtService auth.JWTService
}

func (api *API) groups() {
	api.userGroup("/users")
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
	interfaceAPI.groups()
	interfaceAPI.Engine.Run(":8080")
}
