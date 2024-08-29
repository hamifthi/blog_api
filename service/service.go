package service

import "personal_blog/entity"

type Service interface {
	FindByID(id string) (*entity.Blog, error)
	Create(entity *any) error
	Update(entity *any) error
	Delete(id string) error
}

type User interface {
	Service
}

type Blog interface {
	Service
}
