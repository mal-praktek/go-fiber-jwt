package database

import (
	"fmt"
	"github.com/tegarsubkhan236/go-fiber-project/config"
	model2 "github.com/tegarsubkhan236/go-fiber-project/src/model"
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
	err = DB.AutoMigrate(
		&model2.User{},
		&model2.Product{},
	)
	if err != nil {
		fmt.Println("Table failed to generate")
	}
}
