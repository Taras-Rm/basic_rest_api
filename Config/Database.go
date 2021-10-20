package Config

import (
	"fmt"

	"github.com/Taras-Rm/basic_rest_api/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// import (
// 	"fmt"

// 	"github.com/Taras-Rm/basic_rest_api/Models"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jinzhu/gorm"
// )

// var DB *gorm.DB
// var err error

// func DBConnect() error {
// 	InitConfig()

// 	DbURL := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", Config.App.UserName, Config.App.Pass, Config.App.Host, Config.App.DbName)

// 	// open DB
// 	DB, err = gorm.Open("mysql", DbURL)
// 	// catch errors
// 	if err != nil {
// 		fmt.Println("Status:", err)
// 		return err
// 	}

// 	DB.AutoMigrate(&Models.User{}, &Models.Post{})

// 	return nil
// }

var DB *gorm.DB
var err error

func DBConnect() error {
	InitConfig()

	//DbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=goapi   sslmode=disable", , , )
	DbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s   sslmode=disable", Config.App.Host, Config.App.UserName, Config.App.Pass, Config.App.DbName)
	fmt.Println(DbURL)
	// open DB
	DB, err = gorm.Open(postgres.Open(DbURL), &gorm.Config{})

	// catch error
	if err != nil {
		fmt.Println("Status:", err)
		return err
	}

	err = DB.AutoMigrate(&Models.User{}, &Models.Post{})
	if err != nil {
		fmt.Println("Status:", err)
		return err
	}

	return nil
}
