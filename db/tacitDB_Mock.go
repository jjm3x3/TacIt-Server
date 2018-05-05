package db

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type TacitDBMock struct {
	//Status Check
	TimesAutoMigrateWasCalled int
	TimesCreateWasCalled      int
	TimesWhereWasCalled       int
	TimesFirstWasCalled       int
	StoredPassword            string

	//Behavioral Setup
	FirstResultDBUser *DbUser
	HasError          bool
	NoRecordFound     bool
}

func (db *TacitDBMock) AutoMigrate(values ...interface{}) {
	db.TimesAutoMigrateWasCalled++
}

func (db *TacitDBMock) Create(value interface{}) TacitDB {
	db.TimesCreateWasCalled++
	cvalue, k := value.(*DbUser)
	if k {
		db.StoredPassword = cvalue.Password
	}

	return db
}

func (db *TacitDBMock) Where(query interface{}, args ...interface{}) TacitDB {
	db.TimesWhereWasCalled++
	return db
}

func (db *TacitDBMock) First(out interface{}, where ...interface{}) {
	db.TimesFirstWasCalled++
	wout, k := out.(*DbUser)
	if k {
		wout.Username = db.FirstResultDBUser.Username
		hashword, pearr := bcrypt.GenerateFromPassword([]byte(db.FirstResultDBUser.Password), 10)
		if pearr != nil {
			fmt.Printf("Hashing error: %v\n", pearr)
		} else {
			wout.Password = string(hashword)
		}
	}
}
func (db *TacitDBMock) Error() error {
	if db.HasError {
		return fmt.Errorf("___GENERIC_DATABASE_ERROR___")
	}
	return nil
}

//Mocks gorm NoRecordFound
func (db *TacitDBMock) RecordNotFound() bool {
	return db.NoRecordFound
}
