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
		ctx.JSON(http.StatusAccepted, ShortenUrl{})
	}
}

func (h Handler) GetShortUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		shorten := ctx.Param("id")
		if len(shorten) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{})
		}
		fmt.Println(shorten)
		ctx.JSON(http.StatusAccepted, ShortenUrl{})
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

func (h Handler) SetUrlAccessibility() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		q, ok := ctx.GetQuery("access")
		if ok && q == "true" {
			// TODO set url active to true
		} else {
			// TODO set url active to false
		}
	}
}

func (h Handler) SetUserAccessibility() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		q, ok := ctx.GetQuery("access")
		if ok && q == "true" {
			// TODO set user active to true
		} else {
			// TODO set user active to false
		}
	}
}
