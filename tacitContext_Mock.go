package main

import (
	"fmt"
)

type tacitContextMock struct {
	bindJSONIsCalled  bool
	jsonCode          int
	timesJSONisCalled int
	bindJSONDoesError bool
}

func (ctx *tacitContextMock) bindJSON(obj interface{}) error {
	ctx.bindJSONIsCalled = true
	if ctx.bindJSONDoesError {
		return fmt.Errorf("error")
	} else {
		return nil
	}
}

func (ctx *tacitContextMock) readBody([]byte) (int, error) {
	return 0, nil
	// panic("method not implemented")
}

func (ctx *tacitContextMock) json(code int, obj map[string]interface{}) {
	ctx.jsonCode = code
	ctx.timesJSONisCalled++
}
