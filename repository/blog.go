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

func (r *BlogRepository) Create(b *entity.Blog) error {
	return r.db.Create(b).Error
}

func (r *BlogRepository) GetByID() (entity.Blog, error) {
	var blog entity.Blog
	if err := r.db.First(&blog).Error; err != nil {
		return blog, err
	}
	return blog, nil
}

func (r *BlogRepository) GetAll() ([]entity.Blog, error) {
	var blogs []entity.Blog
	if err := r.db.Find(&blogs).Error; err != nil {
		return blogs, err
	}
	return blogs, nil
}

func (r *BlogRepository) Update(b *entity.Blog) error {
	return r.db.Save(b).Error
}

func (r *BlogRepository) Delete(b *entity.Blog) error {
	return r.db.Delete(b).Error
}
