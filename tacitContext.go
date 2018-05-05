package main

import (
	"github.com/gin-gonic/gin"
)



type tacitContext interface {
	bindJSON(obj interface{}) error
	readBody([]byte) (int, error)
	json(int, map[string]interface{})
}

type realTacitContext struct {
	ginCtx *gin.Context
}

func (ctx *realTacitContext) bindJSON(obj interface{}) error {
	panic("method not implemented")
}

func (ctx *realTacitContext) readBody([]byte) (int, error) {
	panic("method not implemented")
}

func (ctx *realTacitContext) json(int, map[string]interface{}) {
	panic("method not implemented")
}
