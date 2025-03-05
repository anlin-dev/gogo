package router

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	rou "gogo/backend/router"
)

var (
	Router *gin.Engine
)

func Routes() *gin.Engine {

	Router = gin.Default()

	Router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		return
	})

	PublicGroup := Router.Group("")
	{
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
		PublicGroup.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	PrivateGroup := Router.Group("/api/v1")
	for _, router := range rou.RouterGroupApp {
		router.InitRouter(PrivateGroup)
	}
	return Router
}
