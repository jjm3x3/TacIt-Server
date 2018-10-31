package main

import (
	"net/http"
	"os"

	"tacit-api/crypt"
	"tacit-api/db"
	tacitHttp "tacit-api/http"
	"tacit-api/middleware"
	pki "tacit-api/pki"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

type env struct {
	ourDB             db.TacitDB
	logger            logrus.FieldLogger
	ourCrypt          crypt.TacitCrypt
	publicKeyProvider pki.PublicKeyProvider
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

func (e *env) doJwtValidation(c *gin.Context) {
	ctx := tacitHttp.NewContext(c)
	middleware.JwtValidation(ctx, e.logger, e.publicKeyProvider)
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
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowHeaders = []string{"Origin", "Authorization"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	anEnv := &env{
		ourDB:             aRealTacitDB,
		logger:            aLogger,
		ourCrypt:          &crypt.RealTacitCrypt{},
		publicKeyProvider: pki.NewPublicKeyProvider(&http.Client{}, "https://tacit.auth0.com/.well-known/jwks.json"),
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "ok"})
	})

	r.POST("/user", anEnv.doCreateUser)

	r.POST("/login", anEnv.doLogin)

	/// expects controllers to perform callContext.IsAuthed() check
	/// like in the `/health/authed` route
	authedGroup := r.Group("/", anEnv.doJwtValidation)
	authedGroup.GET("/health/authed", func(c *gin.Context) {
		callContext := tacitHttp.NewContext(c)
		if !isAuthed(callContext) {
			return
		}

		callContext.JSON(http.StatusOK, gin.H{"result": "authorized"})
	})

	authedGroup.POST("/note", anEnv.doCreatePost)

	authedGroup.GET("/note", anEnv.doListPosts)

	r.Run()
}
