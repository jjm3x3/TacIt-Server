package main

import (
	"testing"
)

func TestCreatePostReadsBody(t *testing.T) {

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

func TestCreatePostHapyPath(t *testing.T) {

	//setup
	c := &tacitContextMock{
		jsonCode:         0,
		timesJSONisCaled: 0,
	}
	db := &tacitDBMock{}

	//execution
	createPost(c, db)

	//assertions
	if c.jsonCode != 200 {
		t.Error("The expected http status code is 200 for happy path")
	}

	if c.timesJSONisCaled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v time(s)", c.timesJSONisCaled)
	}

}

func TestCreatePostReadsBody(t *testing.T) {

	//setup
	c := &tacitContextMock{
		jsonCode:         0,
		timesJSONisCaled: 0,
	}
	db := &tacitDBMock{}

	//execution
	createPost(c, db)

	//assertions

}
func TestCreatePostSavesPost(t *testing.T) {

	//setup
	c := &tacitContextMock{}
	db := &tacitDBMock{timesCreateWasCalled: 0}
	expectedDbCreates := 1

	//execution
	createPost(c, db)

	//assertions
	if db.timesCreateWasCalled != expectedDbCreates {
		t.Errorf("db.create is expected to be called %v time(s) but instead was called %v time(s)", expectedDbCreates, db.timesCreateWasCalled)
	}
}

type tacitContextMock struct {
	bindJSONIsCalled bool
	jsonCode         int
	timesJSONisCaled int
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
	ctx.timesJSONisCaled++
}
