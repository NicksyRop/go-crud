package main

import (
	"example.com/m/v2/controllers"
	"example.com/m/v2/initializers"
	"example.com/m/v2/middleware"
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
	router.POST("/posts", middleware.RequireAuth, controllers.PostCreate)
	router.GET("/posts", middleware.RequireAuth, controllers.RetrievePosts)
	router.GET("/posts/:id", middleware.RequireAuth, controllers.RetrievePost)
	router.PATCH("/posts/:id", middleware.RequireAuth, controllers.UpdatePost)
	router.DELETE("/posts/:id", middleware.RequireAuth, controllers.DeletePost)
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.Run()

}
