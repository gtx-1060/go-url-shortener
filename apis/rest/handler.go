package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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

		if result, err := h.Service.MakeShortUrl(url); err != nil {
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