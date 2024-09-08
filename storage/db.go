package storage

import "gorm.io/gorm"
import "personal_blog/entity"

type DBGenerator struct {
	DB *gorm.DB
}

func (generator *DBGenerator) ConnectDB(dialector gorm.Dialector) error {
	var err error
	generator.DB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func (generator *DBGenerator) MigrateDB() error {
	err := generator.DB.AutoMigrate(
		&entity.User{},
		&entity.Blog{},
	)
	if err != nil {
		return err
	}
	return nil
}
