package http

import (
	"fmt"
)

type HttpContextMock struct {
	bindJSONIsCalled      bool
	jsonCode              int
	timesJSONisCalled     int
	bindJSONDoesError     bool
	bindJSONResultWebUser *WebUser
}

func (ctx *HttpContextMock) BindJSON(obj interface{}) error {
	ctx.bindJSONIsCalled = true
	if ctx.bindJSONDoesError {
		return fmt.Errorf("error")
	}
	wobj, k := obj.(*WebUser)
	if k {
		wobj.Username = ctx.bindJSONResultWebUser.Username
		wobj.Password = ctx.bindJSONResultWebUser.Password
	}
	return nil

}

func (ctx *HttpContextMock) ReadBody([]byte) (int, error) {
	return 0, nil
}

func (ctx *HttpContextMock) JSON(code int, obj map[string]interface{}) {
	ctx.jsonCode = code
	ctx.timesJSONisCalled++
}

func (ctx *HttpContextMock) GetHeader(key string) string {
	panic("NOT IMPLEMENTED")
}

func (ctx *HttpContextMock) Set(key string, value interface{}) {
	panic("NOT IMPLEMENTED")
}
