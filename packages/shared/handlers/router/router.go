package router

import (
	"go-webview-wasm/packages/shared/handlers/health"

	"github.com/gin-gonic/gin"
)

func GetEngine(appName string) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", health.Health(appName))
	return r
}
