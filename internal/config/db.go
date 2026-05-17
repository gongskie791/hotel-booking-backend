package config

import (
	"booking-system/lvl1/internal/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewDB() *DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL environment variable is not set")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database Connected")
	db.AutoMigrate(&model.User{}, &model.PersonalDetails{}, &model.Hotel{}, &model.Room{}, &model.HotelImage{})
	return &DB{db}
}
