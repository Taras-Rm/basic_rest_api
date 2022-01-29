package Models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Text    string `json:"text"`
	UserRef uint
}

func (b *Post) TableName() string {
	return "posts"
}
