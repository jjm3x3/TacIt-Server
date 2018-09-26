package main

import (
	tacitCrypt "tacit-api/crypt"
	tacitDb "tacit-api/db"
	tacitHttp "tacit-api/http"

	"testing"

	"tacit-api/mocks"

	"github.com/golang/mock/gomock"
)

func TestLoginReadsBody(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	aWebUser := &tacitHttp.WebUser{
		Username: "Username",
		Password: "Password",
	}

	aDBUser := &tacitDb.DbUser{}

	c := &tacitHttp.HttpContextMock{
		BindJSONIsCalled:      false,
		BindJSONResultWebUser: aWebUser,
	}

	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
	}

	crypt := &tacitCrypt.TacitCryptMock{}

	mockLogger.EXPECT().Infof("User %v Logging in", gomock.Any())
	mockLogger.EXPECT().Infof("Found this user from db: %v", gomock.Any())
	mockLogger.EXPECT().Errorf("Error when logging in: %v\n", gomock.Any())

	login(c, db, crypt, mockLogger)

	if !c.BindJSONIsCalled {
		t.Error("bindJSON is never called and should be called at least once.")
	}

}
func TestLoginHappyPath(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	aWebUser := &tacitHttp.WebUser{
		Username: "Username",
		Password: "Password",
	}

	aDBUser := &tacitDb.DbUser{
		Username: aWebUser.Username,
		Password: aWebUser.Password,
	}

	c := &tacitHttp.HttpContextMock{
		JSONCode:              0,
		TimesJSONisCalled:     0,
		BindJSONResultWebUser: aWebUser,
	}

	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
		NoRecordFound:     false,
	}
	crypt := &tacitCrypt.TacitCryptMock{}

	mockLogger.EXPECT().Infof("User %v Logging in", gomock.Any())
	mockLogger.EXPECT().Infof("Found this user from db: %v", gomock.Any())

	login(c, db, crypt, mockLogger)

	if c.JSONCode != 200 {
		t.Errorf("The expected http status code is 200 for happy path. The current status code was %v", c.JSONCode)
	}
	if c.TimesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.TimesJSONisCalled)
	}
}

func TestLoginWrongUsernameRightPassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	aWebUser := &tacitHttp.WebUser{
		Username: "Usernam",
		Password: "Password",
	}

	aDBUser := &tacitDb.DbUser{}

	c := &tacitHttp.HttpContextMock{
		JSONCode:              0,
		TimesJSONisCalled:     0,
		BindJSONResultWebUser: aWebUser,
	}

	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
	}
	crypt := &tacitCrypt.TacitCryptMock{}

	mockLogger.EXPECT().Infof("User %v Logging in", gomock.Any())
	mockLogger.EXPECT().Infof("Found this user from db: %v", gomock.Any())
	mockLogger.EXPECT().Errorf("Error when logging in: %v\n", gomock.Any())

	login(c, db, crypt, mockLogger)

	if c.JSONCode != 401 {
		t.Errorf("The expected http status code is 401 for sad path. The current status code was %v", c.JSONCode)
	}
	if c.TimesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.TimesJSONisCalled)
	}
}

func TestLoginRightUsernameWrongPassword(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	aWebUser := &tacitHttp.WebUser{
		Username: "Username",
		Password: "Passwor",
	}

	aDBUser := &tacitDb.DbUser{
		Username: aWebUser.Username,
		Password: aWebUser.Password + "d",
	}

	c := &tacitHttp.HttpContextMock{
		JSONCode:              0,
		TimesJSONisCalled:     0,
		BindJSONResultWebUser: aWebUser,
	}

	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
	}
	crypt := &tacitCrypt.TacitCryptMock{}

	mockLogger.EXPECT().Infof("User %v Logging in", gomock.Any())
	mockLogger.EXPECT().Infof("Found this user from db: %v", gomock.Any())
	mockLogger.EXPECT().Errorf("Error when logging in: %v\n", gomock.Any())

	login(c, db, crypt, mockLogger)

	if c.JSONCode != 401 {
		t.Errorf("The expected http status code is 401 for sad path. The current status code was %v", c.JSONCode)
	}
	if c.TimesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.TimesJSONisCalled)
	}
}

func TestLoginBindError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &tacitHttp.HttpContextMock{
		JSONCode:          0,
		TimesJSONisCalled: 0,
		BindJSONDoesError: true,
	}

	aDBUser := &tacitDb.DbUser{}
	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
	}

	crypt := &tacitCrypt.TacitCryptMock{}

	mockLogger.EXPECT().Errorf("There was an error parsing login: %v", gomock.Any())

	login(c, db, crypt, mockLogger)

	if c.JSONCode != 400 {
		t.Errorf("The expected http status code is 400 for sad path. The current status code was %v", c.JSONCode)
	}
	if c.TimesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.TimesJSONisCalled)
	}

}
func TestLoginUserDoesNotExistError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	aWebUser := &tacitHttp.WebUser{
		Username: "Username",
		Password: "Password",
	}

	aDBUser := &tacitDb.DbUser{
		Username: aWebUser.Username,
		Password: aWebUser.Password,
	}

	c := &tacitHttp.HttpContextMock{
		JSONCode:              0,
		TimesJSONisCalled:     0,
		BindJSONResultWebUser: aWebUser,
	}

	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
		NoRecordFound:     true,
	}

	crypt := &tacitCrypt.TacitCryptMock{}

	mockLogger.EXPECT().Infof("User %v Logging in", gomock.Any())
	mockLogger.EXPECT().Errorf("The user %v tried to login but is not a user", gomock.Any())

	login(c, db, crypt, mockLogger)

	if c.JSONCode != 401 {
		t.Errorf("The expected http status code is 401 for this path. The current status code was %v", c.JSONCode)
	}
	if c.TimesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.TimesJSONisCalled)
	}
}
func TestCreateUserReadsBody(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	aWebUser := &tacitHttp.WebUser{
		Username: "Username",
		Password: "Password",
	}

	c := &tacitHttp.HttpContextMock{
		BindJSONIsCalled:      false,
		BindJSONResultWebUser: aWebUser,
	}

	db := &tacitDb.TacitDBMock{}

	crypt := &tacitCrypt.TacitCryptMock{}

	mockLogger.EXPECT().Infof("User %v being created", gomock.Any())
	mockLogger.EXPECT().Infof("User %v created", gomock.Any())

	createUser(c, db, crypt, mockLogger)

	if !c.BindJSONIsCalled {
		t.Error("bindJSON is never called and should be called at least once.")
	}

}

func TestCreateUserHappyPath(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	aWebUser := &tacitHttp.WebUser{
		Username: "Username",
		Password: "Password",
	}

	c := &tacitHttp.HttpContextMock{
		JSONCode:              0,
		TimesJSONisCalled:     0,
		BindJSONResultWebUser: aWebUser,
	}

	db := &tacitDb.TacitDBMock{
		NoRecordFound: true,
	}

	crypt := &tacitCrypt.TacitCryptMock{}

	mockLogger.EXPECT().Infof("User %v being created", gomock.Any())
	mockLogger.EXPECT().Infof("User %v created", gomock.Any())

	createUser(c, db, crypt, mockLogger)

	if c.JSONCode != 200 {
		t.Errorf("The expected http status code is 200 for happy path. The current status code was %v", c.JSONCode)
	}
	if c.TimesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.TimesJSONisCalled)
	}
}

func TestCreateUserBindError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &tacitHttp.HttpContextMock{
		JSONCode:          0,
		TimesJSONisCalled: 0,
		BindJSONDoesError: true,
	}

	aDBUser := &tacitDb.DbUser{}
	db := &tacitDb.TacitDBMock{
		FirstResultDBUser: aDBUser,
	}

	crypt := &tacitCrypt.TacitCryptMock{}

	mockLogger.EXPECT().Errorf("There was an error parsing login: %v", gomock.Any())

	createUser(c, db, crypt, mockLogger)
	if c.JSONCode != 400 {
		t.Errorf("The expected http status code is 400 for sad path. The current status code was %v", c.JSONCode)
	}
	if c.TimesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.TimesJSONisCalled)
	}

}

func TestCreateUserSavesUser(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	aWebUser := &tacitHttp.WebUser{
		Username: "Username",
		Password: "Password",
	}
	//setup
	c := &tacitHttp.HttpContextMock{
		BindJSONResultWebUser: aWebUser,
	}
	db := &tacitDb.TacitDBMock{
		TimesCreateWasCalled: 0,
		NoRecordFound:        true,
	}

	crypt := &tacitCrypt.TacitCryptMock{}

	expectedDbCreates := 1

	mockLogger.EXPECT().Infof("User %v being created", gomock.Any())
	mockLogger.EXPECT().Infof("User %v created", gomock.Any())

	//execution
	createUser(c, db, crypt, mockLogger)

	//assertions
	if db.TimesCreateWasCalled != expectedDbCreates {
		t.Errorf("db.create is expected to be called %v time(s) but instead was called %v time(s)", expectedDbCreates, db.TimesCreateWasCalled)
	}

}

func TestCreateUserPasswordStoredProperly(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	aWebUser := &tacitHttp.WebUser{
		Username: "Username",
		Password: "Password",
	}

	c := &tacitHttp.HttpContextMock{
		BindJSONResultWebUser: aWebUser,
	}
	db := &tacitDb.TacitDBMock{}
	crypt := &tacitCrypt.TacitCryptMock{}

	mockLogger.EXPECT().Infof("User %v being created", gomock.Any())
	mockLogger.EXPECT().Infof("User %v created", gomock.Any())

	createUser(c, db, crypt, mockLogger)
	if aWebUser.Password == db.StoredPassword {
		t.Errorf("Password stored in plain text.")
	}
}

func TestCreateUserDatabaseCreationError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	aWebUser := &tacitHttp.WebUser{
		Username: "Username",
		Password: "Password",
	}

	c := &tacitHttp.HttpContextMock{
		BindJSONResultWebUser: aWebUser,
	}
	db := &tacitDb.TacitDBMock{
		HasError:      true,
		NoRecordFound: true,
	}
	crypt := &tacitCrypt.TacitCryptMock{}

	mockLogger.EXPECT().Infof("User %v being created", gomock.Any())
	mockLogger.EXPECT().Errorf("There was an issue creating user: %v", gomock.Any())

	createUser(c, db, crypt, mockLogger)
	if c.JSONCode != 500 {
		t.Errorf("The expected http status code is 500 for user database creation error. The current status code was %v", c.JSONCode)
	}
	if c.TimesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.TimesJSONisCalled)
	}

}

func TestCreateUserGeneratePasswordError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	aWebUser := &tacitHttp.WebUser{
		Username: "Username",
		Password: "Password",
	}
	c := &tacitHttp.HttpContextMock{
		BindJSONResultWebUser: aWebUser,
	}
	db := &tacitDb.TacitDBMock{}
	crypt := &tacitCrypt.TacitCryptMock{
		HasGeneratePasswordError: true,
	}

	mockLogger.EXPECT().Infof("User %v being created", gomock.Any())
	mockLogger.EXPECT().Errorf("There was and error createing password: %v", gomock.Any())

	createUser(c, db, crypt, mockLogger)

	if c.JSONCode != 500 {
		t.Errorf("The expected http status code is 500 for user database createion error. The current status code was %v", c.JSONCode)
	}
	if c.TimesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v times", c.TimesJSONisCalled)
	}

}
