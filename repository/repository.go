package repository

import "personal_blog/entity"

type CommonBehavior[T any] interface {
	GetByID(id uint) (*T, error)
	Create(ent *T) (uint, error)
	Update(ent *T) error
	Delete(ent *T) error
}

type Blog interface {
	CommonBehavior[entity.Blog]
}

type User interface {
	CommonBehavior[entity.User]
}
