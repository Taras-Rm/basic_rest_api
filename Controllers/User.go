package Controllers

import (
	"net/http"

	"github.com/Taras-Rm/basic_rest_api/Config"
	"github.com/Taras-Rm/basic_rest_api/Models"
	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var users []Models.User
	// SELECT * FROM users
	err := Config.DB.Preload("Posts").Find(&users).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "No users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	// Create user
	err = Config.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Server error"})
		return
	}

	c.JSON(http.StatusOK, user)
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	// Get first matched record
	err := Config.DB.Preload("Posts").Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Config.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Can`t found user"})
		return
	}
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return
	}
	err = Config.DB.Save(user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Params.ByName("id")
	err := Config.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}
	err = Config.DB.Select("Posts").Delete(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	err = Config.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
}
