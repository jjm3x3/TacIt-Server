package main

import (
	"fmt"
)

type httpContextMock struct {
	bindJSONIsCalled      bool
	jsonCode              int
	timesJSONisCalled     int
	bindJSONDoesError     bool
	bindJSONResultWebUser *webUser
}

func (ctx *httpContextMock) bindJSON(obj interface{}) error {
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

func (ctx *httpContextMock) readBody([]byte) (int, error) {
	return 0, nil
}

func (ctx *httpContextMock) json(code int, obj map[string]interface{}) {
	ctx.jsonCode = code
	ctx.timesJSONisCalled++
}
