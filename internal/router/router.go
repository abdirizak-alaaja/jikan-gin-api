package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	AnimeRoutes(r)
	return r

}
