package main

import (
	"net/http"

	tacitDB "tacit-api/db"
	tacitHttp "tacit-api/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func createPost(c tacitHttp.HttpContext, db tacitDB.TacitDB, logger logrus.FieldLogger) {
	if !isAuthed(c) {
		return
	}

	var aPost tacitDB.Post
	err := c.BindJSON(&aPost)
	if err != nil {
		var body []byte
		num, err := c.ReadBody(body)
		if num <= 0 { // not sure if this is really an error
			logger.Error("There was no body provided")
		} else if err != nil {
			logger.Errorln("There was an error reading the body: ", err)
		}
		logger.Errorf("There was an error binding to aPost: %v", body)
		c.JSON(400, gin.H{"Error": "There was an error with what you provided"})
		return
	}
	db.Create(&aPost)
	c.JSON(200, gin.H{"status": "success"})
}

func listPosts(c tacitHttp.HttpContext, db tacitDB.TacitDB, logger logrus.FieldLogger) {
	if !isAuthed(c) {
		return
	}

	var somePosts []tacitDB.Post
	err := db.Table("posts").Find(&somePosts).Error()

	if err != nil {
		logger.Errorln("An error has occured fetching posts: ", err)
	}

	c.JSON(200, gin.H{"posts": somePosts})
}

func isAuthed(ctx tacitHttp.HttpContext) bool {
	if isAuthed := ctx.GetBool("authed"); !isAuthed {
		ctx.JSON(http.StatusUnauthorized, gin.H{"result": "unauthorized"})
		return false
	}
	return true
}
