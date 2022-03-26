package config

import (
	"example/pkg/model"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	port, _ := strconv.ParseUint(Config("DB_PORT"), 10, 32)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Config("DB_HOST"), port, Config("DB_USER"), Config("DB_PASSWORD"), Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can't connect to database")
	}
	fmt.Println("Connection success")
	DB.AutoMigrate(
		&model.User{},
		&model.Product{},
	)
}
