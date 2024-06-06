package main

import (
	"github.com/dennisenwiya/go_CRUD_API/initializers"

	"github.com/dennisenwiya/go_CRUD_API/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
