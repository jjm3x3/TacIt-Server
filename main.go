package main

import (
	"os"

	"tacit-api/crypt"
	"tacit-api/db"
	tacitHttp "tacit-api/http"
	"tacit-api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

type env struct {
	ourDB    db.TacitDB
	logger   logrus.FieldLogger
	ourCrypt crypt.TacitCrypt
}

func (e *env) doCreateUser(c *gin.Context) {
	ctx := tacitHttp.NewContext(c)
	createUser(ctx, e.ourDB, e.ourCrypt, e.logger)
}

func (e *env) doLogin(c *gin.Context) {
	ctx := tacitHttp.NewContext(c)
	login(ctx, e.ourDB, e.ourCrypt, e.logger)
}

func (e *env) doCreatePost(c *gin.Context) {
	ctx := tacitHttp.NewContext(c)
	createPost(ctx, e.ourDB, e.logger)
}

func (e *env) doListPosts(c *gin.Context) {
	ctx := tacitHttp.NewContext(c)
	listPosts(ctx, e.ourDB, e.logger)
}

func main() {
	aLogger := logrus.New()
	aLogger.Info("Tacit-api has started")
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
	dbPort := os.Getenv("DB_PORT")
	if len(dbPort) == 0 {
		dbPort = "3306"
	}

	connectionString := dbUser + ":" + dbPassword + "@tcp(127.0.0.1:" + dbPort + ")/" + defaultDb + "?charset=utf8&parseTime=True&loc=Local"
	// connectionString := "host="+defaultHost+" port="+defaultPort+" user="+defaultUser+" dbname="+defaultDb+" sslmode=disable"
	dbHandle, err := gorm.Open("mysql", connectionString) // TODO:: enable ssl
	defer dbHandle.Close()

	if err != nil {
		aLogger.Errorln("There was an error opeing the db: ", err)
		// TODO :: should exit right away
	}

	aRealTacitDB := &db.RealTacitDB{
		GormDB: dbHandle,
	}

	// db.RunMigration(aRealTacitDB)

	r := gin.Default()

	r.Use(middleware.JwtValidation(aLogger))

	r.GET("/ping", func(c *gin.Context) {
		isAuthed := c.GetBool("authed")
		if isAuthed {
			c.JSON(200, gin.H{"message": "pong: authed"})
		} else {
			c.JSON(200, gin.H{"message": "pong: unauthed"})
		}
	})
	anEnv := &env{ourDB: aRealTacitDB, logger: aLogger, ourCrypt: &crypt.RealTacitCrypt{}}

	r.POST("/user", anEnv.doCreateUser)

	r.POST("/login", anEnv.doLogin)

	r.POST("/note", anEnv.doCreatePost)

	r.GET("/note", anEnv.doListPosts)

	r.Run()
}
