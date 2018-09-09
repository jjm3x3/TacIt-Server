package main

import (
	// "fmt"
	"fmt"
	"os"

	"tacit-api/crypt"
	"tacit-api/db"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

type webUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type env struct {
	ourDB    db.TacitDB
	logger   logrus.FieldLogger
	ourCrypt crypt.TacitCrypt
}

func (e *env) doCreateUser(c *gin.Context) {
	ctx := &realHttpContext{ginCtx: c}
	createUser(ctx, e.ourDB, e.ourCrypt, e.logger)
}

func (e *env) doLogin(c *gin.Context) {
	ctx := &realHttpContext{ginCtx: c}
	login(ctx, e.ourDB, e.ourCrypt, e.logger)
}

func (e *env) doCreatePost(c *gin.Context) {
	ctx := &realHttpContext{ginCtx: c}
	createPost(ctx, e.ourDB, e.logger)
}

func (e *env) doListPosts(c *gin.Context) {
	ctx := &realHttpContext{ginCtx: c}
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

	r.Use(FUCKITIWRITEMYOWN())

	r.GET("/ping", func(c *gin.Context) {
		response := c.MustGet("thing").(string)
		c.JSON(200, gin.H{"message": "pong: " + response})
	})
	anEnv := &env{ourDB: aRealTacitDB, logger: aLogger, ourCrypt: &crypt.RealTacitCrypt{}}

	r.POST("/user", anEnv.doCreateUser)

	r.POST("/login", anEnv.doLogin)

	r.POST("/note", anEnv.doCreatePost)

	r.GET("/note", anEnv.doListPosts)

	r.Run()
}

func FUCKITIWRITEMYOWN() gin.HandlerFunc {
	fmt.Println("installing a middleware")
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return "", nil
		})
		if err != nil {
			fmt.Println("There was an issue parsing the JWT")
		}
		// Don't forget to validate the alg is what you expect:
		// token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsExample{}, func(token *jwt.Token) (interface{}, error) {
		// 	// since we only use the one private key to sign the tokens,
		// 	// we also only use its public counter part to verify
		// 	return verifyKey, nil
		// })
		c.Set("thing", tokenString)
		fmt.Printf("IN A MIDDLEWARE BITCH with: '%v'\n", tokenString)
	}
}
