package Controllers

import (
	"net/http"

	"github.com/Taras-Rm/basic_rest_api/Config"
	"github.com/Taras-Rm/basic_rest_api/Models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post Models.Post
	var user Models.User

	id := c.Params.ByName("id")
	c.BindJSON(&post)
	var err = Config.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User is missing"})
	}

	Config.DB.Model(&user).Association("posts").Append(&post)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, post)
	}
}

func GetPostsByUserId(c *gin.Context) {
	var posts []Models.Post

	id := c.Params.ByName("id")

	err := Config.DB.Where("user_ref = ?", id).Find(&posts).Error
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, posts)
	}
}

func UpdatePost(c *gin.Context) {
	var post Models.Post

	id := c.Params.ByName("id")

	var err = Config.DB.Where("id = ?", id).First(&post).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post is missing"})
	}
	c.BindJSON(&post)
	Config.DB.Save(post)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, post)
	}
}

func DeletePost(c *gin.Context) {
	var post Models.Post

	id := c.Params.ByName("id")
	err := Config.DB.Where("id = ?", id).Delete(&post).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Post with that id is missing"})
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
