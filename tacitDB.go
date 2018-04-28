package main

import "github.com/jinzhu/gorm"

type tacitDB interface {
	autoMigrate(values ...interface{})
	create(value interface{})
}

type realTacitDB struct {
	gormDB *gorm.DB
}

func (db *realTacitDB) autoMigrate(values ...interface{}) {
	db.gormDB.AutoMigrate(values)
}

func (db *realTacitDB) create(value interface{}) {
	panic("method not implemented")
}
