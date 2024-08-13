package entity

import "gorm.io/gorm"

// User model representing the author
type User struct {
	gorm.Model
	Name  string `gorm:"size:255;not null"`
	Email string `gorm:"size:255;unique;not null"`
	Blogs []Blog `gorm:"foreignKey:AuthorID"`
}
