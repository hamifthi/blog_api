package migrations

import "gorm.io/gorm"
import "personal_blog/entity"

type DBGenerator struct {
	db *gorm.DB
}

func ConnectDB(generator DBGenerator) error {
	var err error
	generator.db, err = gorm.Open(nil, &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func MigrateDB(generator DBGenerator) error {
	err := generator.db.AutoMigrate(
		&entity.User{},
		&entity.Tag{},
		&entity.Category{},
		&entity.Blog{},
	)
	if err != nil {
		return err
	}
	return nil
}
