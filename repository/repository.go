package repository

import "personal_blog/entity"

type CommonBehavior[T any] interface {
	GetByID(id string) (T, error)
	GetAll() ([]T, error)
	Save(ent T) error
	Update(ent T) error
	Delete(id string) error
}

type Blog interface {
	*CommonBehavior[entity.Blog]
}
