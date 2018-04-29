package main

import (
	"testing"
)

func TestLoginReadsBody(t *testing.T) {

	c := &tacitContextMock{
		bindJSONIsCalled: false,
	}

	db := &tacitDBMock{}

	login(c, db)

	if !c.bindJSONIsCalled {
		t.Error("bindJSON is never called and should be called at least once.")
	}

}
func TestLoginHappyPath(t *testing.T) {

	c := &tacitContextMock{
		jsonCode:          0,
		timesJSONisCalled: 0,
	}

	db := &tacitDBMock{}

	login(c, db)

	if c.jsonCode != 200 {
		t.Errorf("The expected http status code is 200 for happy path. The current status code was %v", c.jsonCode)
	}
	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.timesJSONisCalled)
	}
}
func TestLoginSadPath(t *testing.T) {

	c := &tacitContextMock{
		jsonCode:          0,
		timesJSONisCalled: 0,
		bindJSONDoesError: true,
	}

	db := &tacitDBMock{}

	login(c, db)

	if c.jsonCode != 401 {
		t.Errorf("The expected http status code is 403 for happy path. The current status code was %v", c.jsonCode)
	}
	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.timesJSONisCalled)
	}
}
