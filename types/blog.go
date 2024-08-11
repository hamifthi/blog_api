package types

import (
	"gorm.io/gorm"
	"time"
)

type Status string

const (
	StatusDraft     Status = "draft"
	StatusPublished Status = "published"
	StatusArchived  Status = "archived"
)

type Blog struct {
	gorm.Model
	Title       string   `gorm:"size:255;not null"`
	Content     string   `gorm:"type:text;not null"`
	AuthorID    uint     `gorm:"not null"`
	Author      User     `gorm:"foreignKey:AuthorID"`
	CategoryID  uint     `gorm:"not null"`
	Category    Category `gorm:"foreignKey:CategoryID"`
	Tags        []Tag    `gorm:"many2many:blog_tags;"`
	PublishedAt *time.Time
	Status      Status `gorm:"size:50;default:'draft'"`
}

// User model representing the author
type User struct {
	gorm.Model
	Name  string `gorm:"size:255;not null"`
	Email string `gorm:"size:255;unique;not null"`
	Blogs []Blog `gorm:"foreignKey:AuthorID"`
}

// Category model representing the blog category
type Category struct {
	gorm.Model
	Name  string `gorm:"size:255;not null"`
	Blogs []Blog `gorm:"foreignKey:CategoryID"`
}

// Tag model representing the blog tags
type Tag struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255;not null"`
	Blogs []Blog `gorm:"many2many:blog_tags;"`
}
