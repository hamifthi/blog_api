package entity

import "gorm.io/gorm"

// Category model representing the blog category
type Category struct {
	gorm.Model
	Name  string `gorm:"size:255;not null"`
	Blogs []Blog `gorm:"foreignKey:CategoryID"`
}
