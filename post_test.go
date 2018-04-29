package main

import (
	"testing"
	"fmt"
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
		timesJSONisCalled: 0,
	}
	db := &tacitDBMock{}

	//execution
	createPost(c, db)

	//assertions
	if c.jsonCode != 200 {
		t.Errorf("The expected http status code is 200 for happy path. The current status code was %v", c.jsonCode)
	}

	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.timesJSONisCalled)
	}

}

func TestCreatePostSadPath(t *testing.T) {

	//setup
	c := &tacitContextMock{
		jsonCode: 	0,
		timesJSONisCalled: 0,
		bindJSONDoesError: true,
	}

	db := &tacitDBMock{}

	//execution
	createPost(c, db)

	//assertions
	if c.jsonCode != 400 {
		t.Errorf("The expected http status code is 400 for sad path. The current status code is %v", c.jsonCode)
	}
	if c.timesJSONisCalled !=1 {
		t.Errorf("json should be called on teh context exactly once but instead was called %v times", c.timesJSONisCalled)
	}
	
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
