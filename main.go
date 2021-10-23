package main

import (
	"fmt"

	api "github.com/Taras-Rm/basic_rest_api/Api"
	"github.com/Taras-Rm/basic_rest_api/Config"
	repositories "github.com/Taras-Rm/basic_rest_api/Repositories"
	services "github.com/Taras-Rm/basic_rest_api/Services"

	"github.com/Taras-Rm/basic_rest_api/Controllers"
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

	grp2 := route.Group("/posts")
	{
		grp2.POST("/:id", Controllers.CreatePost)
		grp2.GET("/:id", Controllers.GetPostsByUserId)
		grp2.PUT("/:id", Controllers.UpdatePost)
		grp2.DELETE("/:id", Controllers.DeletePost)
	}

	//running
	route.Run()
}
