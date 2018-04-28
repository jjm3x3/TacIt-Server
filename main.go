package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
// db *gorm.DB
)

type webUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type dbUser struct {
	gorm.Model
	Username string
	Password string
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

	runMigration(dbHandle)

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

func runMigration(db *gorm.DB) {
	// probably doesn't need to happen every time
	db.AutoMigrate(&post{})
	db.AutoMigrate(&dbUser{})
}

func login(c *gin.Context, db *gorm.DB) {
	var login webUser
	err := c.BindJSON(&login)
	if err != nil {
		fmt.Println("There was an error parsing login: ", err)
	}
	fmt.Println("Here is the user info used to login: ", login)

	var theDbUser dbUser
	db.Where("username = ?", login.Username).First(&theDbUser)

	fmt.Println("Found this user from db: ", theDbUser)

	pwBytes := []byte(login.Password)
	err = bcrypt.CompareHashAndPassword([]byte(theDbUser.Password), pwBytes)
	if err != nil {
		fmt.Println("There was something very wrong when logging in!")
		fmt.Println("err: ", err)
		c.JSON(403, gin.H{"Error": "either username or passowrd do not match"})
	} else {
		fmt.Println("Login successful")
		c.JSON(200, gin.H{"status": "login successful"})
	}
}

func createUser(c *gin.Context, db *gorm.DB) {
	var aUser webUser
	err := c.BindJSON(&aUser)
	if err != nil {
		fmt.Println("There was an error parsing User: ", err)
	}
	fmt.Println("Here is the user to create: ", aUser)

	theUser := dbUser{Username: aUser.Username}

	pwBytes := []byte(aUser.Password)
	pwHashBytes, err := bcrypt.GenerateFromPassword(pwBytes, 10)
	if err != nil {
		fmt.Println("There was and error: ", err)
		c.JSON(500, gin.H{"Error": "There was and error with creating your passowrd"})
	}
	theUser.Password = string(pwHashBytes)

	fmt.Println("Here is the user That will be created: ", theUser)

	err = db.Create(&theUser).Error
	if err != nil {
		fmt.Println("There was an issue creating user: ", err)
	}
	c.JSON(200, gin.H{"status": "success"})
}

func createPost(c *gin.Context, db *gorm.DB) {
	var aPost post
	err := c.BindJSON(&aPost)
	if err != nil {
		// fmt.Println("has headers: ", c.GetHeader("Content-Type"))
		var body []byte
		num, err := c.Request.Body.Read(body)
		if num <= 0 { // not sure if this is really an error
			fmt.Println("There was no body provided")
		} else if err != nil {
			fmt.Println("There was an error reading the body: ", err)
		}
		fmt.Println("There was an error binding to aPost: ", body)
		c.JSON(400, gin.H{"Error": "There was an error with what you provided"})
		return
	}
	// fmt.Printf("Here is the result: '%v'\n", aPost)
	db.Create(&aPost)
	c.JSON(200, gin.H{"status": "success"})
}
