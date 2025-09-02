// @title Posts API
// @version 1.0
// @description Posts and Authentication APIS
package main

import (
	"example.com/m/v2/controllers"
	_ "example.com/m/v2/docs" //Add docs module
	"example.com/m/v2/initializers"
	"example.com/m/v2/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//init function runs just before main

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	initializers.MigrateDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/posts", middleware.RequireAuth, controllers.PostCreate)
	router.GET("/posts", middleware.RequireAuth, controllers.RetrievePosts)
	router.GET("/posts/:id", middleware.RequireAuth, controllers.RetrievePost)
	router.PATCH("/posts/:id", middleware.RequireAuth, controllers.UpdatePost)
	router.DELETE("/posts/:id", middleware.RequireAuth, controllers.DeletePost)
	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.Run()

}
