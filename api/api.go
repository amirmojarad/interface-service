package api

import (
	"context"
	"interface_project/api/routers"
	"interface_project/database/crud"
	"interface_project/ent"

	"github.com/gin-gonic/gin"
)

type API struct {
	Crud   *crud.Crud
	Router *routers.Router
}

func RunAPI(ctx *context.Context, client *ent.Client) {
	interfaceAPI := API{
		Crud: &crud.Crud{
			Ctx:    ctx,
			Client: client,
		},
		Router: &routers.Router{
			Engine:    gin.Default(),
			UserPath:  "/users",
			MoviePath: "/movie",
		},
	}
	interfaceAPI.Router.UsersRouter()
	interfaceAPI.Router.Engine.Run("localhost:8080")
}
