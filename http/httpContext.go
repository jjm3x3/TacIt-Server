package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type HttpContext interface {
	// methods defined to act as a gin.Context
	BindJSON(obj interface{}) error
	ReadBody([]byte) (int, error)
	JSON(int, map[string]interface{})
	GetHeader(string) string
	Set(string, interface{})

	// New Methods defined here
	IsAuthed() bool
}

type RealHttpContext struct {
	ginCtx *gin.Context
}

func NewContext(c *gin.Context) HttpContext {
	return &RealHttpContext{c}
}

func (ctx *RealHttpContext) BindJSON(obj interface{}) error {
	return ctx.ginCtx.BindJSON(obj)
}

func (ctx *RealHttpContext) ReadBody(outbytes []byte) (int, error) {
	return ctx.ginCtx.Request.Body.Read(outbytes)
}

func (ctx *RealHttpContext) JSON(code int, jsonResponse map[string]interface{}) {
	ctx.ginCtx.JSON(code, jsonResponse)
}

func (ctx *RealHttpContext) GetHeader(key string) string {
	return ctx.ginCtx.GetHeader(key)
}

func (ctx *RealHttpContext) Set(key string, value interface{}) {
	ctx.ginCtx.Set(key, value)
}

func (ctx *RealHttpContext) IsAuthed() bool {
	if isAuthed := ctx.ginCtx.GetBool("authed"); !isAuthed {
		ctx.ginCtx.JSON(http.StatusUnauthorized, gin.H{"result": "unauthorized"})
		return false
	}
	return true
}
