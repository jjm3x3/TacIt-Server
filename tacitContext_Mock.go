package main

import (
	"fmt"
)

type tacitContextMock struct {
	bindJSONIsCalled      bool
	jsonCode              int
	timesJSONisCalled     int
	bindJSONDoesError     bool
	bindJSONResultWebUser *webUser
}

func (ctx *tacitContextMock) bindJSON(obj interface{}) error {
	ctx.bindJSONIsCalled = true
	if ctx.bindJSONDoesError {
		return fmt.Errorf("error")
	}
	wobj, k := obj.(*webUser)
	if k {
		wobj.Username = ctx.bindJSONResultWebUser.Username
		wobj.Password = ctx.bindJSONResultWebUser.Password
	}
	return nil

}

func (ctx *tacitContextMock) readBody([]byte) (int, error) {
	return 0, nil
}

func (ctx *tacitContextMock) json(code int, obj map[string]interface{}) {
	ctx.jsonCode = code
	ctx.timesJSONisCalled++
}
