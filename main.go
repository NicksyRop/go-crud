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
}

func main() {

	router := gin.Default()
	router.POST("/posts", controllers.PostCreate)
	router.Run()

}
