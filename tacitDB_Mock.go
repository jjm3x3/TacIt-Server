package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type tacitDBMock struct {
	//Status Check
	timesAutoMigrateWasCalled int
	timesCreateWasCalled      int
	timesWhereWasCalled       int
	timesFirstWasCalled       int
	storedPassword            string

	//Behavioral Setup
	firstResultDBUser *dbUser
	hasError          bool
	noRecordFound     bool
}

func (db *tacitDBMock) autoMigrate(values ...interface{}) {
	db.timesAutoMigrateWasCalled++
}

func (db *tacitDBMock) create(value interface{}) tacitDB {
	db.timesCreateWasCalled++
	cvalue, k := value.(*dbUser)
	if k {
		db.storedPassword = cvalue.Password
	}

	return db
}

func (db *tacitDBMock) where(query interface{}, args ...interface{}) tacitDB {
	db.timesWhereWasCalled++
	return db
}

func (db *tacitDBMock) first(out interface{}, where ...interface{}) {
	db.timesFirstWasCalled++
	wout, k := out.(*dbUser)
	if k {
		wout.Username = db.firstResultDBUser.Username
		hashword, pearr := bcrypt.GenerateFromPassword([]byte(db.firstResultDBUser.Password), 10)
		if pearr != nil {
			fmt.Printf("Hashing error: %v\n", pearr)
		} else {
			wout.Password = string(hashword)
		}
	}
}
func (db *tacitDBMock) error() error {
	if db.hasError {
		return fmt.Errorf("___GENERIC_DATABASE_ERROR___")
	}
	return nil
}

//Mocks gorm NoRecordFound
func (db *tacitDBMock) recordNotFound() bool {
	return db.noRecordFound
}
