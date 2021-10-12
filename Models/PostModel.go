package Models

type Post struct {
	Id uint `json:"id"`
	//gorm.Model
	Title   string `json:"title"`
	Text    string `json:"text"`
	UserRef uint
}

func (b *Post) TableName() string {
	return "posts"
}
