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

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	api.InjectUser(route, userService)

	postRepository := repositories.NewPostRepository(db)
	postService := services.NewPostService(postRepository, userRepository)
	api.InjectPost(route, postService)

	//running
	route.Run()
}
