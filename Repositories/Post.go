package repositories

import (
	"github.com/Taras-Rm/basic_rest_api/Models"
	"gorm.io/gorm"
)

type PostRepository interface {
	CreatePost(user *Models.User, post *Models.Post) (*Models.Post, error)
	GetPostsByUserId(id uint) ([]Models.Post, error)
	UpdatePost(id uint, post *Models.Post) (*Models.Post, error)
	DeletePost(id uint) error
	GetPostById(id uint) (*Models.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) CreatePost(user *Models.User, post *Models.Post) (*Models.Post, error) {
	res := r.db.Model(&user).Association("Posts").Append(post)
	return post, res
} //+

func (r *postRepository) GetPostsByUserId(id uint) ([]Models.Post, error) {
	var posts []Models.Post
	res := r.db.Where("user_ref = ?", id).Find(&posts)
	return posts, res.Error
}

func (r *postRepository) UpdatePost(id uint, post *Models.Post) (*Models.Post, error) {
	res := r.db.Model(post).Where("id = ?", id).Updates(map[string]interface{}{
		"title": post.Title,
		"text":  post.Text,
	})
	return post, res.Error
}

func (r *postRepository) DeletePost(id uint) error {
	res := r.db.Where("id = ?", id).Delete(&Models.Post{})
	return res.Error
}

func (r *postRepository) GetPostById(id uint) (*Models.Post, error) {
	var post *Models.Post
	res := r.db.Where("id = ?", id).First(&post)
	return post, res.Error
}
