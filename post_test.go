package main

import (
	"testing"
)

func TestCreatePost(t *testing.T) {

	//setup
	c := &tacitContextMock{
		bindJSONIsCalled: false,
	}
	db := &tacitDBMock{}

	//execution
	createPost(c, db)

	//assertions
	if !c.bindJSONIsCalled {
		t.Error("bindJSON is never called and should be called at least once")
	}
}

type tacitContextMock struct {
	bindJSONIsCalled bool
	jsonCode         int
}

func (ctx *tacitContextMock) bindJSON(obj interface{}) error {
	ctx.bindJSONIsCalled = true
	return nil
}

func (ctx *tacitContextMock) readBody([]byte) (int, error) {
	panic("method not implemented")
}

func (ctx *tacitContextMock) json(code int, obj map[string]interface{}) {
	ctx.jsonCode = code
}
