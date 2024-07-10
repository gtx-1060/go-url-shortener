package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"url-shortener/dtos"
	. "url-shortener/services"
)

type Handler struct {
	Service *Service
}

func (h Handler) MakeShortUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url := dtos.UrlToShort{}
		if ok := ctx.ShouldBindJSON(&url); ok != nil {
			ctx.JSON(http.StatusBadRequest, dtos.ErrorResponse("there aren't needed args"))
			return
		}
		if time.Now().After(url.Expiration) {
			ctx.JSON(http.StatusBadRequest, dtos.ErrorResponse("incorrect expiration time"))
			return
		}

		if result, err := h.Service.MakeShortUrl(ctx, url); err != nil {
			ctx.JSON(http.StatusBadRequest, dtos.ErrorResponse(err.Error()))
		} else {
			ctx.JSON(http.StatusAccepted, result)
		}
	}
}

func (h Handler) GetShortUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		shorten, ok := ctx.GetQuery("id")
		if !ok {
			ctx.JSON(http.StatusBadRequest, dtos.ErrorResponse("url is empty"))
			return
		}
		if result, err := h.Service.GetUrlDataByShort(shorten); err != nil {
			ctx.JSON(http.StatusBadRequest, dtos.ErrorResponse(err.Error()))
		} else {
			ctx.JSON(http.StatusAccepted, result)
		}
	}
}

func (h Handler) FollowShortUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		shorten := ctx.Param("id")
		if len(shorten) == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{})
		}
		if url, err := h.Service.GetUrlByShort(shorten); err != nil {
			ctx.JSON(http.StatusNotFound, dtos.ErrorResponse(err.Error()))
		} else {
			fmt.Println(url)
			ctx.Redirect(http.StatusFound, url)
		}
	}
}
