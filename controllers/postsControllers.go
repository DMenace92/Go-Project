package controllers

import (
	"github.com/dennisenwiya/go_CRUD_API/initializers"
	"github.com/dennisenwiya/go_CRUD_API/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//get data off req body
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}
	//create a post

	//requtn it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostFetch(c *gin.Context) {
	//get the posts

	// Get all records
	var posts []models.Post
	initializers.DB.Find(&posts)
	//respond to them

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostSingleFetch(c *gin.Context) {
	//get id off url
	id := c.Param("id")
	//get the posts

	// Get all records
	var post models.Post
	initializers.DB.First(&post, id)
	//respond to them

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//Get the id off url
	id := c.Param("id")
	//get the data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	//Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	//update it

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	//respond with it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	//Get the id off url
	id := c.Param("id")
	// delete the post
	initializers.DB.Delete(&models.Post{}, id)
	//respond
	c.Status(200)
}
