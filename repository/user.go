package repository

import (
	"gorm.io/gorm"
	"personal_blog/entity"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *entity.User) (uint, error) {
	result := r.db.Create(user)
	if err := result.Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (r *UserRepository) GetByID(id uint) (*entity.User, error) {
	var user entity.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(user *entity.User) error {
	return r.db.Delete(user).Error
}
