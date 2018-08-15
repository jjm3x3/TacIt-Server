package main

import (
	"github.com/gin-gonic/gin"
)

type httpContext interface {
	bindJSON(obj interface{}) error
	readBody([]byte) (int, error)
	json(int, map[string]interface{})
}

type realHttpContext struct {
	ginCtx *gin.Context
}

func (ctx *realHttpContext) bindJSON(obj interface{}) error {
	return ctx.ginCtx.BindJSON(obj)
}

func (ctx *realHttpContext) readBody(outbytes []byte) (int, error) {
	return ctx.ginCtx.ReadBody(outbytes)
}

func (ctx *realHttpContext) json(code int, jsonResponse map[string]interface{}) {
	return ctx.ginCtx.JSON(code, jsonResponse)
}
