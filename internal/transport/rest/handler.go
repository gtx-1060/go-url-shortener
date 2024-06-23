package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func (h Handler) MakeShortUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url := &UrlToShort{}
		if ok := ctx.ShouldBindJSON(url); ok != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		ctx.JSON(http.StatusAccepted, ShortenUrl{Author: url.Author})
	}
}

func (h Handler) GetShortUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		shorten := ctx.Param("id")
		if len(shorten) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{})
		}
		fmt.Println(shorten)
		ctx.JSON(http.StatusAccepted, ShortenUrl{Author: shorten})
	}
}

func (h Handler) FollowShortUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		shorten := ctx.Param("id")
		if len(shorten) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{})
		}
		fmt.Println(shorten)
		ctx.Redirect(http.StatusMovedPermanently, "https://google.com")
	}
}
