package controllers

import (
	"example.com/m/v2/initializers"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	//get data from request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)
	//check if we have error
	if result.Error != nil {
		c.Status(400)
		return
	}
	//return post
	c.JSON(200, gin.H{
		"post": post,
	})
}
