package api

import (
	"net/http"
	"strconv"

	"github.com/Taras-Rm/basic_rest_api/Models"
	services "github.com/Taras-Rm/basic_rest_api/Services"
	"github.com/gin-gonic/gin"
)

func InjectPost(gr *gin.Engine, postService services.PostService) {
	handler := gr.Group("posts")
	handler.POST("/:id", postCreate(postService))
	handler.GET("/:id", postGetByUserId(postService))
	handler.PUT("/:id", postUpdate(postService))
	handler.DELETE("/:id", postDelete(postService))
}

func postCreate(postService services.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post *Models.Post

		id := c.Params.ByName("id")
		userId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		err = c.BindJSON(&post)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		res, err := postService.CreatePost(uint(userId), post)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Server error", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func postGetByUserId(postService services.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")

		userId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		posts, err := postService.GetPostsByUserId(uint(userId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, posts)
	}
}

func postUpdate(postService services.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post *Models.Post

		id := c.Params.ByName("id")
		postId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		err = c.BindJSON(&post)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		err = postService.UpdatePost(uint(postId), post)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Server error", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Post is updated"})
	}
}

func postDelete(postService services.PostService) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Params.ByName("id")

		postId, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
			return
		}

		err = postService.DeletePost(uint(postId))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Post with that id is missing"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
