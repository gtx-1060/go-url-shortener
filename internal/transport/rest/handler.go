package rest

import (
	"github.com/gin-gonic/gin"
)

type HandlerContext struct {
}

func (hCtx HandlerContext) MakeShortUrl() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func (hCtx HandlerContext) GetShortUrl() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
