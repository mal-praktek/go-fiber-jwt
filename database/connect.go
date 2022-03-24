package database

import (
	"fmt"
	"github.com/tegarsubkhan236/go-fiber-project/config"
	"github.com/tegarsubkhan236/go-fiber-project/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
)

func ConnectDB() {
	var err error
	port, _ := strconv.ParseUint(config.Config("DB_PORT"), 10, 32)
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
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
