package db

import (
	"testing"
)

func TestNothing(t *testing.T) {
	if 1 != 1 {
		t.Fail()
	}
}

func TestRunMigration(t *testing.T) {

	// setup
	db := &TacitDBMock{TimesAutoMigrateWasCalled: 0}
	expectedTimes := 2

	// execution
	RunMigration(db)

	//assertions
	if db.TimesAutoMigrateWasCalled == 0 {
		t.Error("autoMigrate was not called on tacitDB")
		t.FailNow()
	}

	if db.TimesAutoMigrateWasCalled != expectedTimes {
		t.Errorf("autoMigrate was not called the expected number of times %v instead it was called %v times", expectedTimes, db.TimesAutoMigrateWasCalled)
	}

}
