/*
Package api contains routers and groups.
it has API structure that contains:
	- Crud instance (a package to manage database crud actions)
	- Engine instance that manage Gin HTTP web framework.
	- jwtService that manages jwt actions.
*/

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

// startEngine call all api groups and html templates.
func (api *API) startEngine() {
	api.groups()
	api.Engine.LoadHTMLGlob("template/*")
	api.Engine.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})
}

// groups call all api groups (routers)
func (api *API) groups() {
	api.Engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	api.userGroup("/users")
	api.movieGroup("/movies")
	api.fileGroup("/file")
	api.authGroup("/")
}

// RunAPI gives context and client to run api and start program.
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
