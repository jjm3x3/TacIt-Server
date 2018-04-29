package main

import "github.com/jinzhu/gorm"

type tacitDB interface {
	autoMigrate(values ...interface{})
	create(value interface{})
	first(out interface{}, where ...interface{})
	where(query interface{}, args ...interface{}) tacitDB
}

type realTacitDB struct {
	gormDB *gorm.DB
}

func (db *realTacitDB) autoMigrate(values ...interface{}) {
	db.gormDB.AutoMigrate(values)
}

func (db *realTacitDB) create(value interface{}) {
	db.gormDB.Create(value)
}

func (db *realTacitDB) first(out interface{}, where ...interface{}) {
	db.gormDB.First(out, where)
}

func (db *realTacitDB) where(query interface{}, args ...interface{}) tacitDB {
	db.gormDB = db.gormDB.Where(query, args)
	return db
}
