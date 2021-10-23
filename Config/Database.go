package Config

import (
	"fmt"

	"github.com/Taras-Rm/basic_rest_api/Models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func DBConnect() (*gorm.DB, error) {
	InitConfig()

	DbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s   sslmode=disable", Config.App.Host, Config.App.UserName, Config.App.Pass, Config.App.DbName)

	// open DB
	DB, err = gorm.Open(postgres.Open(DbURL), &gorm.Config{})

	// catch error
	if err != nil {
		fmt.Println("Status:", err)
		return nil, err
	}

	err = DB.AutoMigrate(&Models.User{}, &Models.Post{})
	if err != nil {
		fmt.Println("Status:", err)
		return nil, err
	}

	return DB, nil
}
