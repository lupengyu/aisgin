package helper

import (
	"context"
	"github.com/gin-gonic/gin"
)

const (
	RequestID  = "REQUEST_ID"
	MethodKey  = "K_METHOD"
	UrlPathKey = "URL_PATH"
)

func GenCtxFromGin(c *gin.Context) context.Context {
	ctx := context.WithValue(context.Background(), UrlPathKey, c.Request.URL.Path)
	ctx = context.WithValue(ctx, MethodKey, c.GetString(MethodKey))
	return ctx
}
