package initializers

import "gorm.io/gorm"

var DB *gorm.DB

func ConnectDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(nil, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return DB, nil
}
