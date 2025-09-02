package controllers

import (
	"log"

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

func RetrievePosts(c *gin.Context) {
	//nb you can get the user fetching post

	user, exists := c.Get("user")
	//convert user to User model
	if exists {
		log.Printf("User with ID %d retrieving posts", user.(models.User).ID)
	}

	//get posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//return posts
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func RetrievePost(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	//get post
	var post models.Post
	initializers.DB.First(&post, id)

	//return post
	c.JSON(200, gin.H{
		"post": post,
	})

}

func UpdatePost(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	//Get data of request body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	//get post in db
	var post models.Post
	initializers.DB.First(&post, id)

	//update the post
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	//return post
	c.JSON(200, gin.H{
		"post": post,
	})

}

func DeletePost(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	//get post in db
	var post models.Post
	initializers.DB.Delete(&post, id)

	c.Status(200)

}
