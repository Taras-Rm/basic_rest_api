package main

import (
	"fmt"

	"github.com/Taras-Rm/basic_rest_api/Config"
	"github.com/Taras-Rm/basic_rest_api/Models"

	"github.com/Taras-Rm/basic_rest_api/Controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	// open DB
	Config.DB, err = gorm.Open("mysql", Config.DbURL)
	// catch errors
	if err != nil {
		fmt.Println("Status:", err)
	}
	// close DB
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.User{}, &Models.Post{})

	route := gin.Default()

	grp1 := route.Group("/users")
	{
		grp1.GET("/", Controllers.GetUsers)
		grp1.POST("/", Controllers.CreateUser)
		grp1.GET("/:id", Controllers.GetUserByID)
		grp1.PUT("/:id", Controllers.UpdateUser)
		grp1.DELETE("/:id", Controllers.DeleteUser)
	}

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
