package http

import (
	"fmt"
)

type HttpContextMock struct {
	BindJSONIsCalled      bool
	JSONCode              int
	TimesJSONisCalled     int
	BindJSONDoesError     bool
	BindJSONResultWebUser *WebUser
}

func (ctx *HttpContextMock) BindJSON(obj interface{}) error {
	ctx.BindJSONIsCalled = true
	if ctx.BindJSONDoesError {
		return fmt.Errorf("error")
	}
	wobj, k := obj.(*WebUser)
	if k {
		wobj.Username = ctx.BindJSONResultWebUser.Username
		wobj.Password = ctx.BindJSONResultWebUser.Password
	}
	return nil

}

func (ctx *HttpContextMock) ReadBody([]byte) (int, error) {
	return 0, nil
}

func (ctx *HttpContextMock) JSON(code int, obj map[string]interface{}) {
	ctx.JSONCode = code
	ctx.TimesJSONisCalled++
}

func (ctx *HttpContextMock) GetHeader(key string) string {
	panic("NOT IMPLEMENTED")
}

func (ctx *HttpContextMock) Set(key string, value interface{}) {
	panic("NOT IMPLEMENTED")
}
