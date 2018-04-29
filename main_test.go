package main

import (
	"testing"
)

func TestNothing(t *testing.T) {
	// result := main()
	if 1 != 1 {
		t.Fail()
	}
}

func TestRunMigration(t *testing.T) {

	// setup
	db := &tacitDBMock{timesAutoMigrateWasCalled: 0}
	expectedTimes := 2

	// execution
	runMigration(db)

	//assertions
	if db.timesAutoMigrateWasCalled == 0 {
		t.Error("autoMigrate was not called on tacitDB")
		t.FailNow()
	}

	if db.timesAutoMigrateWasCalled != expectedTimes {
		t.Errorf("autoMigrate was not called the expected number of times %v instead it was called %v times", expectedTimes, db.timesAutoMigrateWasCalled)
	}

}

type tacitDBMock struct {
	timesAutoMigrateWasCalled int
	timesCreateWasCalled      int
	timesWhereWasCalled       int
	timesFirstWasCalled       int
}

func (db *tacitDBMock) autoMigrate(values ...interface{}) {
	db.timesAutoMigrateWasCalled++
}

func (db *tacitDBMock) create(value interface{}) {
	db.timesCreateWasCalled++
}

func (db *tacitDBMock) where(query interface{}, args ...interface{}) tacitDB {
	db.timesWhereWasCalled++
	return db
}

func (db *tacitDBMock) first(out interface{}, where ...interface{}) {
	db.timesFirstWasCalled++
}
