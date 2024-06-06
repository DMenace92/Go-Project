package main

import (
	"github.com/dennisenwiya/go_CRUD_API/controllers"
	"github.com/dennisenwiya/go_CRUD_API/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostFetch)

	r.GET("/posts/:id", controllers.PostSingleFetch)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)

	r.Run() // listen and serve on 0.0.0.0:8080
}
