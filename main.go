package main

import (
	"example.com/m/v2/controllers"
	"example.com/m/v2/initializers"
	"github.com/gin-gonic/gin"
)

//init function runs just before main

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	initializers.MigrateDatabase()
}

func main() {
	router := gin.Default()
	router.POST("/posts", controllers.PostCreate)
	router.GET("/posts", controllers.RetrievePosts)
	router.GET("/posts/:id", controllers.RetrievePost)
	router.PATCH("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.Run()

}
