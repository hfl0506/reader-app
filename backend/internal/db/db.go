package db

import (
	"github.com/hfl0506/reader-app/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.Book{}, &model.User{})

	return db, nil
}
