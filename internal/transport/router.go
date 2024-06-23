package transport

import (
	"github.com/gin-gonic/gin"
	"url-shortener/internal/transport/rest"
)

func InitRouter() *gin.Engine {
	handler := rest.Handler{}
	engine := gin.Default()
	engine.Use(rest.CORSMiddleware())
	engine.POST("/url", handler.MakeShortUrl())
	engine.GET("/url", handler.GetShortUrl())
	engine.GET("/s/:id", handler.FollowShortUrl())
	return engine
}
