package controllers

import (
	"log"

	"example.com/m/v2/initializers"
	"example.com/m/v2/models"
	"github.com/gin-gonic/gin"
)

// PostCreate godoc
// @Summary      Create a new post
// @Description  Create a new blog post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        post  body      models.Post  true  "Post to create"
// @Success      200   {object}  models.Post
// @Failure      400   {object}  map[string]string
// @Router       /posts [post]
func PostCreate(c *gin.Context) {
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

// RetrievePosts godoc
// @Summary      Get all posts
// @Description  Retrieve all posts in the system
// @Tags         posts
// @Produce      json
// @Success      200  {array}   models.Post
// @Router       /posts [get]
func RetrievePosts(c *gin.Context) {
	user, exists := c.Get("user")
	if exists {
		log.Printf("User with ID %d retrieving posts", user.(models.User).ID)
	}

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

// RetrievePost godoc
// @Summary      Get a post by ID
// @Description  Retrieve a single post by its ID
// @Tags         posts
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Success      200  {object}  models.Post
// @Failure      404  {object}  map[string]string
// @Router       /posts/{id} [get]
func RetrievePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

// UpdatePost godoc
// @Summary      Update a post
// @Description  Update post details by ID
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        id    path      int         true  "Post ID"
// @Param        post  body      models.Post true  "Updated post"
// @Success      200   {object}  models.Post
// @Failure      400   {object}  map[string]string
// @Router       /posts/{id} [patch]
func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	c.JSON(200, gin.H{
		"post": post,
	})
}

// DeletePost godoc
// @Summary      Delete a post
// @Description  Delete a post by ID
// @Tags         posts
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Success      200  {string}  string  "ok"
// @Router       /posts/{id} [delete]
func DeletePost(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.Delete(&post, id)

	c.Status(200)
}
