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
	panic("method not implemented")
}

func (ctx *realHttpContext) readBody([]byte) (int, error) {
	panic("method not implemented")
}

func (ctx *realHttpContext) json(int, map[string]interface{}) {
	panic("method not implemented")
}
