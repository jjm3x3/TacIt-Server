package main

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
