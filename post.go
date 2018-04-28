package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type post struct {
	gorm.Model
	Title string `json:"title"`
	Body  string `json:"body"`
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
