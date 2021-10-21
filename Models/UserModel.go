package Models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Posts   []Post `json:"posts" gorm:"foreignKey:UserRef"`
}

// it`s func that return table name "users" for structure User
func (b *User) TableName() string {
	return "users"
}
