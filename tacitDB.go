package main

import "github.com/jinzhu/gorm"

type tacitDB interface {
	autoMigrate(values ...interface{})
	create(value interface{}) tacitDB
	first(out interface{}, where ...interface{})
	where(query interface{}, args ...interface{}) tacitDB
	error() error
}

type realTacitDB struct {
	gormDB *gorm.DB
}

func (db *realTacitDB) autoMigrate(values ...interface{}) {
	db.gormDB.AutoMigrate(values)
}

func (db *realTacitDB) create(value interface{}) tacitDB {
	db.gormDB = db.gormDB.Create(value)
	return db
}

func (db *realTacitDB) first(out interface{}, where ...interface{}) {
	db.gormDB.First(out, where)
}

func (db *realTacitDB) where(query interface{}, args ...interface{}) tacitDB {
	db.gormDB = db.gormDB.Where(query, args)
	return db
}

func (db *realTacitDB) error() error {
	return db.gormDB.Error
}
