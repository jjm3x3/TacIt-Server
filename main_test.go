package main

import (
	"testing"
)

func TestNothing(t *testing.T) {
	// result := main()
	if 1 != 2 {
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
