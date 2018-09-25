package main

import (
	tacitDb "tacit-api/db"
	"tacit-api/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreatePostReadsBody(t *testing.T) {

	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &httpContextMock{
		bindJSONIsCalled: false,
	}
	db := &tacitDb.TacitDBMock{}

	//execution
	createPost(c, db, mockLogger)

	//assertions
	if !c.bindJSONIsCalled {
		t.Error("bindJSON is never called and should be called at least once")
	}

}

func TestCreatePostHapyPath(t *testing.T) {

	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &httpContextMock{
		jsonCode:          0,
		timesJSONisCalled: 0,
	}
	db := &tacitDb.TacitDBMock{}

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
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &httpContextMock{
		jsonCode:          0,
		timesJSONisCalled: 0,
		bindJSONDoesError: true,
	}

	db := &tacitDb.TacitDBMock{}

	mockLogger.EXPECT().Error(gomock.Any()).AnyTimes()
	mockLogger.EXPECT().Errorf("There was an error binding to aPost: %v", gomock.Any())

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
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &httpContextMock{}
	db := &tacitDb.TacitDBMock{TimesCreateWasCalled: 0}
	expectedDbCreates := 1

	//execution
	createPost(c, db, mockLogger)

	//assertions
	if db.TimesCreateWasCalled != expectedDbCreates {
		t.Errorf("db.create is expected to be called %v time(s) but instead was called %v time(s)", expectedDbCreates, db.TimesCreateWasCalled)
	}
}

func TestCreatePostBindJSONFailureLogsError(t *testing.T) {

	//setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &httpContextMock{
		bindJSONDoesError: true,
	}
	db := &tacitDb.TacitDBMock{}

	expectedMessage := "There was an error binding to aPost: %v"
	mockLogger.EXPECT().Error("There was no body provided")
	mockLogger.EXPECT().Errorf(expectedMessage, gomock.Any())

	//execution
	createPost(c, db, mockLogger)

	//assertions
	// taken care of through gomock
}

func TestListPostsHappyPath(t *testing.T) {

	// setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &httpContextMock{
		bindJSONDoesError: true,
	}
	db := &tacitDb.TacitDBMock{}

	listPosts(c, db, mockLogger)

	//assertions
	if c.jsonCode != 200 {
		t.Errorf("The expected http status code is 200 for happy path. The current status code was %v", c.jsonCode)
	}

	if c.timesJSONisCalled != 1 {
		t.Errorf("json should be called on the context exactly once but instead was called %v Times", c.timesJSONisCalled)
	}
}

func TestListPostsReadsFromDB(t *testing.T) {

	// setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &httpContextMock{}
	db := &tacitDb.TacitDBMock{}

	listPosts(c, db, mockLogger)

	//assertions
	if db.TimesFindWasCalled != 1 {
		t.Errorf("Find was not called the expected number of times: 1")
	}
	if db.TableSearched != "posts" {
		t.Errorf("The expected table was not searched: posts")
	}
}

func TestListPostsWillLogAnErrorInCaseOfDBFailure(t *testing.T) {

	// setup
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockLogger := mocks.NewMockFieldLogger(mockCtrl)

	c := &httpContextMock{}
	db := &tacitDb.TacitDBMock{HasError: true}

	mockLogger.EXPECT().Errorln("An error has occured fetching posts: ", gomock.Any())

	listPosts(c, db, mockLogger)

	//assertions
	// taken care of through gomock

}
