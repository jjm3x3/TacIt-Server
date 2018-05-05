package main

import (
	tacitDB "tacit-api/db"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func createPost(c httpContext, db tacitDB.TacitDB, logger logrus.FieldLogger) {
	var aPost tacitDB.Post
	err := c.bindJSON(&aPost)
	if err != nil {
		var body []byte
		num, err := c.readBody(body)
		if num <= 0 { // not sure if this is really an error
			logger.Error("There was no body provided")
		} else if err != nil {
			logger.Errorln("There was an error reading the body: ", err)
		}
		logger.Errorf("There was an error binding to aPost: %v", body)
		c.json(400, gin.H{"Error": "There was an error with what you provided"})
		return
	}
	db.Create(&aPost)
	c.json(200, gin.H{"status": "success"})
}
