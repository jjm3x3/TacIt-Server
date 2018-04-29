package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type dbUser struct {
	gorm.Model
	Username string
	Password string
}

func login(c tacitContext, db tacitDB) {
	var login webUser
	err := c.bindJSON(&login)
	if err != nil {
		fmt.Println("There was an error parsing login: ", err)
	}
	fmt.Println("Here is the user info used to login: ", login)

	var theDbUser dbUser
	db.where("username = ?", login.Username).first(&theDbUser)

	fmt.Println("Found this user from db: ", theDbUser)

	pwBytes := []byte(login.Password)
	err = bcrypt.CompareHashAndPassword([]byte(theDbUser.Password), pwBytes)
	if err != nil {
		fmt.Println("There was something very wrong when logging in!")
		fmt.Println("err: ", err)
		c.json(401, gin.H{"Error": "either username or password do not match"})
	} else {
		fmt.Println("Login successful")
		c.json(200, gin.H{"status": "login successful"})
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
