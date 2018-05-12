package main

import (
	tacitDb "tacit-api/db"
	"testing"
)

func TestLoginReadsBody(t *testing.T) {
	aWebUser := &webUser{
		Username: "Username",
		Password: "Password",
	}

	aDBUser := &tacitDb.DbUser{}

	c := &httpContextMock{
		bindJSONIsCalled:      false,
		bindJSONResultWebUser: aWebUser,
	}

	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
	}

	crypt := &tacitDb.TacitCryptMock{}

	login(c, db, crypt)

	if !c.bindJSONIsCalled {
		t.Error("bindJSON is never called and should be called at least once.")
	}

}
func TestLoginHappyPath(t *testing.T) {
	aWebUser := &webUser{
		Username: "Username",
		Password: "Password",
	}

	aDBUser := &tacitDb.DbUser{
		Username: aWebUser.Username,
		Password: aWebUser.Password,
	}

	c := &httpContextMock{
		jsonCode:              0,
		timesJSONisCalled:     0,
		bindJSONResultWebUser: aWebUser,
	}

	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
		NoRecordFound:     false,
	}
	crypt := &tacitDb.TacitCryptMock{}

	login(c, db, crypt)

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

	aDBUser := &tacitDb.DbUser{}

	c := &httpContextMock{
		jsonCode:              0,
		timesJSONisCalled:     0,
		bindJSONResultWebUser: aWebUser,
	}

	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
	}
	crypt := &tacitDb.TacitCryptMock{}

	login(c, db, crypt)
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

	aDBUser := &tacitDb.DbUser{
		Username: aWebUser.Username,
		Password: aWebUser.Password + "d",
	}

	c := &httpContextMock{
		jsonCode:              0,
		timesJSONisCalled:     0,
		bindJSONResultWebUser: aWebUser,
	}

	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
	}
	crypt := &tacitDb.TacitCryptMock{}

	login(c, db, crypt)
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

	aDBUser := &tacitDb.DbUser{}
	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
	}

	crypt := &tacitDb.TacitCryptMock{}

	login(c, db, crypt)
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

	aDBUser := &tacitDb.DbUser{
		Username: aWebUser.Username,
		Password: aWebUser.Password,
	}

	c := &httpContextMock{
		jsonCode:              0,
		timesJSONisCalled:     0,
		bindJSONResultWebUser: aWebUser,
	}

	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
		NoRecordFound:     true,
	}

	crypt := &tacitDb.TacitCryptMock{}

	login(c, db, crypt)

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

	db := &tacitDb.TacitDBMock{}

	crypt := &tacitDb.TacitCryptMock{}

	createUser(c, db, crypt)

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

	db := &tacitDb.TacitDBMock{
		NoRecordFound: true,
	}

	crypt := &tacitDb.TacitCryptMock{}

	createUser(c, db, crypt)

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

	aDBUser := &tacitDb.DbUser{}
	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
	}

	crypt := &tacitDb.TacitCryptMock{}

	createUser(c, db, crypt)
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
	db := &tacitDb.TacitDBMock{
		TimesCreateWasCalled: 0,
		NoRecordFound:        true,
	}

	crypt := &tacitDb.TacitCryptMock{}

	expectedDbCreates := 1

	//execution
	createUser(c, db, crypt)

	//assertions
	if db.TimesCreateWasCalled != expectedDbCreates {
		t.Errorf("db.create is expected to be called %v time(s) but instead was called %v time(s)", expectedDbCreates, db.TimesCreateWasCalled)
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
	db := &tacitDb.TacitDBMock{}
	crypt := &tacitDb.TacitCryptMock{}

	createUser(c, db, crypt)
	if aWebUser.Password == db.StoredPassword {
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
	db := &tacitDb.TacitDBMock{
		HasError:      true,
		NoRecordFound: true,
	}
	crypt := &tacitDb.TacitCryptMock{}

	createUser(c, db, crypt)
	if c.jsonCode != 500 {
		t.Errorf("The expected http status code is 500 for user database creation error. The current status code was %v", c.jsonCode)
	}
	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.timesJSONisCalled)
	}

}

func TestCreateUserGeneratePasswordError(t *testing.T) {
	aWebUser := &webUser{
		Username: "Username",
		Password: "Password",
	}
	c := &httpContextMock{
		bindJSONResultWebUser: aWebUser,
	}
	db := &tacitDb.TacitDBMock{}
	crypt := &tacitDb.TacitCryptMock{
		HasGeneratePasswordError: true,
	}
	createUser(c, db, crypt)
	if c.jsonCode != 500 {
		t.Errorf("The expected http status code is 500 for user database createion error. The current status code was %v", c.jsonCode)
	}
	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.timesJSONisCalled)
	}

}
