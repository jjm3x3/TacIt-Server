package db

import (
	"github.com/jinzhu/gorm"
)

type DbUser struct {
	gorm.Model
	Username string
	Password string
}

type Post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

type TacitDB interface {
	AutoMigrate(values ...interface{})
	Create(value interface{}) TacitDB
	First(out interface{}, where ...interface{})
	Where(query interface{}, args ...interface{}) TacitDB
	Error() error
	RecordNotFound() bool
}



type RealTacitDB struct {
	GormDB *gorm.DB
}



func (db *RealTacitDB) AutoMigrate(values ...interface{}) {
	db.GormDB.AutoMigrate(values)
}

func (db *RealTacitDB) Create(value interface{}) TacitDB {
	db.GormDB = db.GormDB.Create(value)
	return db
}

func (db *RealTacitDB) First(out interface{}, where ...interface{}) {
	db.GormDB.First(out, where)
}

func (db *RealTacitDB) Where(query interface{}, args ...interface{}) TacitDB {
	db.GormDB = db.GormDB.Where(query, args)
	return db
}

func (db *RealTacitDB) Error() error {
	return db.GormDB.Error
}

//Uses gorm RecordNotFound
func (db *RealTacitDB) RecordNotFound() bool {
	notFound := db.GormDB.RecordNotFound()
	return notFound
}

func RunMigration(db TacitDB) {
	// probably doesn't need to happen every time
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&DbUser{})
}

