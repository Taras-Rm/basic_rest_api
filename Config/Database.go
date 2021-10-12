package Config

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var DbURL = "root:1tab238B@tcp(localhost:3306)/goapp?charset=utf8&parseTime=True&loc=Local"
