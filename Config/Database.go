package Config

import (
	"fmt"

	"github.com/Taras-Rm/basic_rest_api/Models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

func DBConnect() error {
	InitConfig()

	DbURL := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", Config.App.UserName, Config.App.Pass, Config.App.Host, Config.App.DbName)

	// open DB
	DB, err = gorm.Open("mysql", DbURL)
	// catch errors
	if err != nil {
		fmt.Println("Status:", err)
		return err
	}

	DB.AutoMigrate(&Models.User{}, &Models.Post{})

	return nil
}
