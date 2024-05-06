package controllers

import (
	"net/http"
	"vnexpress/initializers"
	"vnexpress/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Title   string
		Image   string
		Content string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Image: body.Image, Content: body.Content}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {

	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Title   string
		Image   string
		Content string
	}

	c.Bind(&body)
	var post models.Post
	initializers.DB.First(&post, id)
	initializers.DB.Model(&post).Updates(models.Post{
		Title:   body.Title,
		Image:   body.Image,
		Content: body.Content,
	})
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {

	id := c.Param("id")
	initializers.DB.Delete(&models.Post{}, id)
	c.Status(http.StatusOK)
}
