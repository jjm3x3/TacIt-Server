package main

import (
	"TacIt-go/mocks"
	tacitDb "TacIt-go/db"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	// "fmt"
)

type loggerMock struct {
	errorMessage string
}

func (logger *loggerMock) WithField(key string, value interface{}) *logrus.Entry {
	return nil
}

func (logger *loggerMock) WithFields(fields logrus.Fields) *logrus.Entry {
	return nil
}

func (logger *loggerMock) WithError(err error) *logrus.Entry {
	return nil
}

func (logger *loggerMock) Debugf(format string, args ...interface{}) {
}
func (logger *loggerMock) Infof(format string, args ...interface{}) {
}
func (logger *loggerMock) Printf(format string, args ...interface{}) {
}
func (logger *loggerMock) Warnf(format string, args ...interface{}) {
}
func (logger *loggerMock) Warningf(format string, args ...interface{}) {
}
func (logger *loggerMock) Errorf(format string, args ...interface{}) {
	logger.errorMessage = format
}
func (logger *loggerMock) Fatalf(format string, args ...interface{}) {
}
func (logger *loggerMock) Panicf(format string, args ...interface{}) {
}

func (logger *loggerMock) Debug(args ...interface{}) {
}
func (logger *loggerMock) Info(args ...interface{}) {
}
func (logger *loggerMock) Print(args ...interface{}) {
}
func (logger *loggerMock) Warn(args ...interface{}) {
}
func (logger *loggerMock) Warning(args ...interface{}) {
}
func (logger *loggerMock) Error(args ...interface{}) {
}
func (logger *loggerMock) Fatal(args ...interface{}) {
}
func (logger *loggerMock) Panic(args ...interface{}) {
}

func (logger *loggerMock) Debugln(args ...interface{}) {
}
func (logger *loggerMock) Infoln(args ...interface{}) {
}
func (logger *loggerMock) Println(args ...interface{}) {
}
func (logger *loggerMock) Warnln(args ...interface{}) {
}
func (logger *loggerMock) Warningln(args ...interface{}) {
}
func (logger *loggerMock) Errorln(args ...interface{}) {
}
func (logger *loggerMock) Fatalln(args ...interface{}) {
}
func (logger *loggerMock) Panicln(args ...interface{}) {
}

func TestCreatePostReadsBody(t *testing.T) {

	//setup
	c := &httpContextMock{
		bindJSONIsCalled: false,
	}
	db := &tacitDb.TacitDBMock{}
	logger := &loggerMock{}

	//execution
	createPost(c, db, logger)

	//assertions
	if !c.bindJSONIsCalled {
		t.Error("bindJSON is never called and should be called at least once")
	}

}

func TestCreatePostHapyPath(t *testing.T) {

	//setup
	c := &httpContextMock{
		jsonCode:          0,
		timesJSONisCalled: 0,
	}
	db := &tacitDb.TacitDBMock{}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	// logger := &loggerMock{}

	//execution
	createPost(c, db, mockLogger)

	//assertions
	if c.jsonCode != 200 {
		t.Errorf("The expected http status code is 200 for happy path. The current status code was %v", c.jsonCode)
	}

	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v Times", c.timesJSONisCalled)
	}

}

func TestCreatePostSadPath(t *testing.T) {

	//setup
	c := &httpContextMock{
		jsonCode:          0,
		timesJSONisCalled: 0,
		bindJSONDoesError: true,
	}

	db := &tacitDb.TacitDBMock{}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	mockLogger.EXPECT().Error("There was no body provided")
	mockLogger.EXPECT().Errorf("There was an error binding to aPost: %v", gomock.Any())
	// logger := &loggerMock{}

	//execution
	createPost(c, db, mockLogger)

	//assertions
	if c.jsonCode != 400 {
		t.Errorf("The expected http status code is 400 for sad path. The current status code is %v", c.jsonCode)
	}
	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on teh context exactly once but instead was called %v times", c.timesJSONisCalled)
	}

}
func TestCreatePostSavesPost(t *testing.T) {

	//setup
	c := &httpContextMock{}
	db := &tacitDb.TacitDBMock{TimesCreateWasCalled: 0}
	expectedDbCreates := 1

	logger := &loggerMock{}

	//execution
	createPost(c, db, logger)

	//assertions
	if db.TimesCreateWasCalled != expectedDbCreates {
		t.Errorf("db.create is expected to be called %v time(s) but instead was called %v time(s)", expectedDbCreates, db.TimesCreateWasCalled)
	}
}

func TestCreatePostBindJSONFailureLogsError(t *testing.T) {

	//setup
	c := &httpContextMock{
		bindJSONDoesError: true,
	}
	db := &tacitDBMock{}

	logger := &loggerMock{}

	//execution
	createPost(c, db, logger)

	expectedMessage := "There was an error binding to aPost: %v"
	//assertions
	if logger.errorMessage != expectedMessage {
		t.Errorf("In the case of a bindJSON error the logger is expected to log the fallowing '%v' but instead logged the message: '%v'", expectedMessage, logger.errorMessage)
	}
}
