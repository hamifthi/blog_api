package service

import "personal_blog/entity"

type Service[T any] interface {
	FindByID(id uint) (*T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(id uint) error
}

type User interface {
	Service[entity.User]
}

type Blog interface {
	Service[entity.Blog]
}
