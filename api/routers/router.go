package routers

import "github.com/gin-gonic/gin"

type Router struct {
	Engine    *gin.Engine
	UserPath  string
	MoviePath string
}
