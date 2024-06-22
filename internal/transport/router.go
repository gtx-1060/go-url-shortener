package transport

import (
	"github.com/gin-gonic/gin"
	"url-shortener/internal/transport/rest"
)

func InitRouter() *gin.Engine {
	context := rest.HandlerContext{}
	engine := gin.Default()
	engine.Use(rest.CORSMiddleware())
	engine.POST("/url", context.MakeShortUrl())
	engine.GET("/s/:id", context.GetShortUrl())
	return engine
}
