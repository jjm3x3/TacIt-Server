package main

import (
	"testing"
)

func TestLoginReadsBody(t *testing.T) {
	aWebUser := &webUser{
		Username: "Username",
		Password: "Password",
	}

	aDBUser := &dbUser{}

	c := &httpContextMock{
		bindJSONIsCalled:      false,
		bindJSONResultWebUser: aWebUser,
	}

	db := &tacitDBMock{
		firstResultDBUser: aDBUser,
	}

	login(c, db)

	if !c.bindJSONIsCalled {
		t.Error("bindJSON is never called and should be called at least once.")
	}

}
func TestLoginHappyPath(t *testing.T) {
	aWebUser := &webUser{
		Username: "Username",
		Password: "Password",
	}

	aDBUser := &dbUser{
		Username: aWebUser.Username,
		Password: aWebUser.Password,
	}

	c := &httpContextMock{
		jsonCode:              0,
		timesJSONisCalled:     0,
		bindJSONResultWebUser: aWebUser,
	}

	db := &tacitDBMock{
		firstResultDBUser: aDBUser,
	}

	login(c, db)

	if c.jsonCode != 200 {
		t.Errorf("The expected http status code is 200 for happy path. The current status code was %v", c.jsonCode)
	}
	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.timesJSONisCalled)
	}
}

func TestLoginWrongUsernameRightPassword(t *testing.T) {

	aWebUser := &webUser{
		Username: "Usernam",
		Password: "Password",
	}

	aDBUser := &dbUser{}

	c := &httpContextMock{
		jsonCode:              0,
		timesJSONisCalled:     0,
		bindJSONResultWebUser: aWebUser,
	}

	db := &tacitDBMock{
		firstResultDBUser: aDBUser,
	}

	login(c, db)
	if c.jsonCode != 401 {
		t.Errorf("The expected http status code is 401 for sad path. The current status code was %v", c.jsonCode)
	}
	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.timesJSONisCalled)
	}
}

func TestLoginRightUsernameWrongPassword(t *testing.T) {

	aWebUser := &webUser{
		Username: "Username",
		Password: "Passwor",
	}

	aDBUser := &dbUser{
		Username: aWebUser.Username,
		Password: aWebUser.Password + "d",
	}

	c := &httpContextMock{
		jsonCode:              0,
		timesJSONisCalled:     0,
		bindJSONResultWebUser: aWebUser,
	}

	db := &tacitDBMock{
		firstResultDBUser: aDBUser,
	}

	login(c, db)
	if c.jsonCode != 401 {
		t.Errorf("The expected http status code is 401 for sad path. The current status code was %v", c.jsonCode)
	}
	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.timesJSONisCalled)
	}
}

func TestLoginBindError(t *testing.T) {
	c := &httpContextMock{
		jsonCode:          0,
		timesJSONisCalled: 0,
		bindJSONDoesError: true,
	}

	aDBUser := &dbUser{}
	db := &tacitDBMock{
		firstResultDBUser: aDBUser,
	}

	login(c, db)
	if c.jsonCode != 400 {
		t.Errorf("The expected http status code is 400 for sad path. The current status code was %v", c.jsonCode)
	}
	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.timesJSONisCalled)
	}

}
func TestLoginUserDoesNotExistError(t *testing.T) {
	aWebUser := &webUser{
		Username: "Username",
		Password: "Password",
	}

	aDBUser := &dbUser{
		Username: aWebUser.Username,
		Password: aWebUser.Password,
	}

	c := &httpContextMock{
		jsonCode:              0,
		timesJSONisCalled:     0,
		bindJSONResultWebUser: aWebUser,
	}

	db := &tacitDBMock{
		firstResultDBUser: aDBUser,
		noRecordFound:     true,
	}

	login(c, db)

	if c.jsonCode != 401 {
		t.Errorf("The expected http status code is 401 for this path. The current status code was %v", c.jsonCode)
	}
	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.timesJSONisCalled)
	}
}
func TestCreateUserReadsBody(t *testing.T) {
	aWebUser := &webUser{
		Username: "Username",
		Password: "Password",
	}

	c := &httpContextMock{
		bindJSONIsCalled:      false,
		bindJSONResultWebUser: aWebUser,
	}

	db := &tacitDBMock{}

	createUser(c, db)

	if !c.bindJSONIsCalled {
		t.Error("bindJSON is never called and should be called at least once.")
	}

}

func TestCreateUserHappyPath(t *testing.T) {
	aWebUser := &webUser{
		Username: "Username",
		Password: "Password",
	}

	c := &httpContextMock{
		jsonCode:              0,
		timesJSONisCalled:     0,
		bindJSONResultWebUser: aWebUser,
	}

	db := &tacitDBMock{}

	createUser(c, db)

	if c.jsonCode != 200 {
		t.Errorf("The expected http status code is 200 for happy path. The current status code was %v", c.jsonCode)
	}
	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.timesJSONisCalled)
	}
}

func TestCreateUserBindError(t *testing.T) {
	c := &httpContextMock{
		jsonCode:          0,
		timesJSONisCalled: 0,
		bindJSONDoesError: true,
	}

	aDBUser := &dbUser{}
	db := &tacitDBMock{
		firstResultDBUser: aDBUser,
	}

	createUser(c, db)
	if c.jsonCode != 400 {
		t.Errorf("The expected http status code is 400 for sad path. The current status code was %v", c.jsonCode)
	}
	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.timesJSONisCalled)
	}

}

func TestCreateUserSavesUser(t *testing.T) {
	aWebUser := &webUser{
		Username: "Username",
		Password: "Password",
	}
	//setup
	c := &httpContextMock{
		bindJSONResultWebUser: aWebUser,
	}
	db := &tacitDBMock{
		timesCreateWasCalled: 0,
	}
	expectedDbCreates := 1

	//execution
	createUser(c, db)

	//assertions
	if db.timesCreateWasCalled != expectedDbCreates {
		t.Errorf("db.create is expected to be called %v time(s) but instead was called %v time(s)", expectedDbCreates, db.timesCreateWasCalled)
	}

}

func TestCreateUserPasswordStoredProperly(t *testing.T) {
	aWebUser := &webUser{
		Username: "Username",
		Password: "Password",
	}

	c := &httpContextMock{
		bindJSONResultWebUser: aWebUser,
	}
	db := &tacitDBMock{}
	createUser(c, db)
	if aWebUser.Password == db.storedPassword {
		t.Errorf("Password stored in plain text.")
	}
}

func TestCreateUserDatabaseCreationError(t *testing.T) {
	aWebUser := &webUser{
		Username: "Username",
		Password: "Password",
	}

	c := &httpContextMock{
		bindJSONResultWebUser: aWebUser,
	}
	db := &tacitDBMock{
		hasError: true,
	}

	createUser(c, db)
	if c.jsonCode != 500 {
		t.Errorf("The expected http status code is 500 for user database creation error. The current status code was %v", c.jsonCode)
	}
	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.timesJSONisCalled)
	}

}
