package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type webUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type env struct {
	db *gorm.DB
}

func (e *env) doCreateUser(c *gin.Context) {
	createUser(c, e.db)
}

func (e *env) doLogin(c *gin.Context) {
	login(c, e.db)
}

func (e *env) doCreatePost(c *gin.Context) {
	createPost(c, e.db)
}

func main() {
	fmt.Println("Hello, World")
	// defaultHost := "localhost"
	// defaultPort := "5432"
	dbUser := os.Getenv("DB_USER")
	if len(dbUser) == 0 {
		dbUser = "gorm"
	}
	defaultDb := "tacit_db"
	dbPassword := os.Getenv("DB_PASSWORD")
	if len(dbPassword) == 0 {
		dbPassword = "@"
	}

	// var err error
	connectionString := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:3306)/" + defaultDb + "?charset=utf8&parseTime=True&loc=Local"
	// connectionString := "host="+defaultHost+" port="+defaultPort+" user="+defaultUser+" dbname="+defaultDb+" sslmode=disable"
	dbHandle, err := gorm.Open("mysql", connectionString) // TODO:: enable ssl
	defer dbHandle.Close()

	if err != nil {
		fmt.Println("There was an error opeing the db: ", err)
		// TODO :: should exit right away
	}

	runMigration(&realTacitDB{gormDB: dbHandle})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	anEnv := &env{db: dbHandle}

	r.POST("/user", anEnv.doCreateUser)

	r.POST("/login", anEnv.doLogin)

	r.POST("/note", anEnv.doCreatePost)

	r.Run()
}

func runMigration(db tacitDB) {
	// probably doesn't need to happen every time
	db.autoMigrate(&post{})
	db.autoMigrate(&dbUser{})
}

type tacitDB interface {
	autoMigrate(values ...interface{})
}

type realTacitDB struct {
	gormDB *gorm.DB
}

func (db *realTacitDB) autoMigrate(values ...interface{}) {
	db.gormDB.AutoMigrate(values)
}
