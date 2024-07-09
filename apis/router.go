package apis

import (
	"github.com/gin-gonic/gin"
	. "url-shortener/apis/rest"
	. "url-shortener/services"
)

func InitRouter(service *Service) *gin.Engine {
	handler := Handler{service}
	engine := gin.Default()
	engine.Use(CORSMiddleware("*"))
	engine.POST("/url", handler.MakeShortUrl())
	engine.GET("/url", handler.GetShortUrl())
	engine.GET("/s/:id", handler.FollowShortUrl())
	return engine
}
