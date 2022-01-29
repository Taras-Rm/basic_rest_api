package api

import (
	"net/http"
	"strconv"

	"github.com/Taras-Rm/basic_rest_api/Models"
	services "github.com/Taras-Rm/basic_rest_api/Services"
	"github.com/gin-gonic/gin"
)

func InjectUser(gr *gin.Engine, userService services.UserService) {
	handler := gr.Group("users")
	handler.GET("/", usersList(userService))
	handler.POST("/", userCreate(userService))
	handler.GET("/:id", userGet(userService))
	handler.PUT("/:id", userUpdate(userService))
	handler.DELETE("/:id", userDelete(userService))
}

func usersList(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := userService.GetUsers()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "No users"})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func userCreate(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *Models.User

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		res, err := userService.CreateUser(user)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Server error"})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func userGet(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")

		userId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
			return
		}

		res, err := userService.GetUserByID(uint(userId))

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Server error", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func userUpdate(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *Models.User

		id := c.Params.ByName("id")
		userId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
			return
		}

		err = c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
			return
		}

		err = userService.UpdateUser(uint(userId), user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User is updated"})
	}
}

func userDelete(userService services.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")

		userId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request", "error": err.Error()})
			return
		}

		err = userService.DeleteUser(uint(userId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
