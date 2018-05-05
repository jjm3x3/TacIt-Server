package main

import (
	"fmt"

	tacitDb "tacit-api/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func login(c httpContext, db tacitDb.TacitDB) {
	var login webUser
	err := c.bindJSON(&login)
	if err != nil {
		fmt.Println("There was an error parsing login: ", err)
		c.json(400, gin.H{"Error": "Invalid login body"})
		return
	}
	fmt.Println("Here is the user info used to login: ", login)

	var theDbUser tacitDb.DbUser
	db.Where("username = ?", login.Username).First(&theDbUser)
	if db.RecordNotFound() {
		//Questionable Error return
		c.json(401, gin.H{"Error": "User does not exist"})
		return
	}

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

func createUser(c httpContext, db tacitDb.TacitDB) {
	var aUser webUser
	err := c.bindJSON(&aUser)
	if err != nil {
		fmt.Println("There was an error parsing User: ", err)
		c.json(400, gin.H{"Error": "Invalid create user body"})
		return
	}
	fmt.Println("Here is the user to create: ", aUser)

	theUser := tacitDb.DbUser{Username: aUser.Username}

	pwBytes := []byte(aUser.Password)
	pwHashBytes, err := bcrypt.GenerateFromPassword(pwBytes, 10)
	if err != nil {
		fmt.Println("There was and error: ", err)
		c.json(500, gin.H{"Error": "There was an error with creating your password"})
		return
	}
	theUser.Password = string(pwHashBytes)

	fmt.Println("Here is the user That will be created: ", theUser)

	err = db.Create(&theUser).Error()
	if err != nil {
		fmt.Println("There was an issue creating user: ", err)
		c.json(500, gin.H{"Error": "There was an error with creating your user"})
		return
	}
	c.json(200, gin.H{"status": "success"})
}
