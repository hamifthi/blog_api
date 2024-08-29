package repository

import (
	"gorm.io/gorm"
	"personal_blog/entity"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{db}
}

func (r *BlogRepository) Create(blog *entity.Blog) (uint, error) {
	result := r.db.Create(blog)
	if err := result.Error; err != nil {
		return 0, err
	}
	return blog.ID, nil
}

func (r *BlogRepository) GetByID(id uint) (*entity.Blog, error) {
	var blog entity.Blog
	if err := r.db.Preload("Author").First(&blog, id).Error; err != nil {
		return nil, err
	}
	return &blog, nil
}

func (r *BlogRepository) Update(blog *entity.Blog) error {
	return r.db.Save(blog).Error
}

func (r *BlogRepository) Delete(blog *entity.Blog) error {
	return r.db.Delete(blog).Error
}
