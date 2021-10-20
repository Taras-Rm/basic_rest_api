package Controllers

import (
	"fmt"
	"net/http"

	"github.com/Taras-Rm/basic_rest_api/Config"
	"github.com/Taras-Rm/basic_rest_api/Models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post Models.Post
	var user Models.User

	id := c.Params.ByName("id")
	err := c.BindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	err = Config.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	fmt.Println(user)

	err = Config.DB.Model(&user).Association("Posts").Append(&post)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"messageAsoc": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func GetPostsByUserId(c *gin.Context) {
	var posts []Models.Post

	id := c.Params.ByName("id")

	err := Config.DB.Where("user_ref = ?", id).Find(&posts).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func UpdatePost(c *gin.Context) {
	var post Models.Post
	id := c.Params.ByName("id")

	err := Config.DB.Where("id = ?", id).First(&post).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
		return
	}
	err = c.BindJSON(&post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}

	err = Config.DB.Save(post).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context) {
	var post Models.Post

	id := c.Params.ByName("id")
	err := Config.DB.Where("id = ?", id).First(&post).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post not found"})
		return
	}

	err = Config.DB.Where("id = ?", id).Delete(&post).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post with that id is missing"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}
