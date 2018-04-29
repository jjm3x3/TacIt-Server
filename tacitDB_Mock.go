package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

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
	wout, k := out.(*dbUser)
	if k {
		wout.Username = "Username"
		hashword, pearr := bcrypt.GenerateFromPassword([]byte("Password"), 10)
		if pearr != nil {
			fmt.Printf("Hashing error: %v\n", pearr)
		} else {
			wout.Password = string(hashword)
		}
	}
}
