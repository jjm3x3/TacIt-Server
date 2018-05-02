package main

import (
	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
}

func createPost(c httpContext, db tacitDB, logger logrus.FieldLogger) {
	var aPost post
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
	db.create(&aPost)
	c.json(200, gin.H{"status": "success"})
}
