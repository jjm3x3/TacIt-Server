package main

import (
	"fmt"

	tacitCrypt "tacit-api/crypt"
	tacitDb "tacit-api/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func login(c httpContext, db tacitDb.TacitDB, crypt tacitCrypt.TacitCrypt, logger logrus.FieldLogger) {
	var login webUser
	err := c.bindJSON(&login)
	if err != nil {
		logger.Errorf("There was an error parsing login: %v", err)
		c.json(400, gin.H{"Error": "Invalid login body"})
		return
	}
	logger.Infof("User %v Logging in", login.Username)

	var theDbUser tacitDb.DbUser
	db.Where("username = ?", login.Username).First(&theDbUser)
	if db.RecordNotFound() {
		logger.Errorf("The user %v tried to login but is not a user", login.Username)
		//Questionable Error return
		c.json(401, gin.H{"Error": "User does not exist"})
		return
	}

	logger.Infof("Found this user from db: %v", theDbUser)

	pwBytes := []byte(login.Password)
	err = crypt.CompareHashAndPassword([]byte(theDbUser.Password), pwBytes)

	if err != nil {
		logger.Errorf("Error when logging in: %v\n", err)
		c.json(401, gin.H{"Error": "either username or password do not match"})
	} else {
		c.json(200, gin.H{"status": "login successful"})
	}
}

func createUser(c httpContext, db tacitDb.TacitDB, crypt tacitCrypt.TacitCrypt) {
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
	pwHashBytes, err := crypt.GenerateFromPassword(pwBytes, 10)
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
