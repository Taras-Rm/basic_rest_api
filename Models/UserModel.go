package Models

type User struct {
	Id uint `json:"id"`
	//gorm.Model
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

//{"name": "Iryna", "email": "ira3@gmail.com", "phone": "0215478541", "address": "fdfffd3" }
