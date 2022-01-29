package main

import (
	"fmt"

	api "github.com/Taras-Rm/basic_rest_api/Api"
	"github.com/Taras-Rm/basic_rest_api/Config"
	repositories "github.com/Taras-Rm/basic_rest_api/Repositories"
	services "github.com/Taras-Rm/basic_rest_api/Services"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := Config.DBConnect()
	// catch errors
	if err != nil {
		fmt.Println("Status:", err)
	}

	route := gin.Default()

	route.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	api.InjectUser(route, userService)

	postRepository := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepository, userRepository)
	api.InjectPost(route, postService)

	//running
	route.Run()
}
