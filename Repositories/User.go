package repositories

import (
	"github.com/Taras-Rm/basic_rest_api/Models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers() ([]Models.User, error)
	CreateUser(user *Models.User) (*Models.User, error)
	GetUserByID(id uint) (*Models.User, error)
	UpdateUser(id uint, user *Models.User) error
	DeleteUser(id uint) error
	DeleteUserPosts(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers() ([]Models.User, error) {
	var users []Models.User
	res := r.db.Preload("Posts").Find(&users)
	return users, res.Error
}

func (r *userRepository) CreateUser(user *Models.User) (*Models.User, error) {
	res := r.db.Create(&user)
	return user, res.Error
}

func (r *userRepository) GetUserByID(id uint) (*Models.User, error) {
	var user *Models.User
	res := r.db.Preload("Posts").Where("id = ?", id).First(&user)
	return user, res.Error
}

func (r *userRepository) UpdateUser(id uint, user *Models.User) error {
	res := r.db.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{
		"Name":    user.Name,
		"Email":   user.Email,
		"Phone":   user.Phone,
		"Address": user.Address,
	})
	return res.Error
}

func (r *userRepository) DeleteUser(id uint) error {
	res := r.db.Where("id = ?", id).Delete(&Models.User{})
	return res.Error
}

func (r *userRepository) DeleteUserPosts(id uint) error {
	res := r.db.Where("user_ref = ?", id).Delete(&Models.Post{})
	return res.Error
}
