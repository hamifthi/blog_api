package entity

// Tag model representing the blog tags
type Tag struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:255;not null"`
	Blogs []Blog `gorm:"many2many:blog_tags;"`
}
