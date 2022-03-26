package service

import (
	"errors"
	"example/pkg/config"
	"example/pkg/model"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
)

func GetUserById(id string) (*model.User, error) {
	db := config.DB
	var user model.User
	if err := db.Find(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(e string) (*model.User, error) {
	db := config.DB
	var user model.User
	if err := db.Where(&model.User{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByUsername(u string) (*model.User, error) {
	db := config.DB
	var user model.User
	if err := db.Where(&model.User{Username: u}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

//func CreateUser(data struct{}) (*model.User, error) {
//	db := config.DB
//	if err := db.Create(&data).Error; err != nil {
//		return nil, err
//	}
//	return &data, nil
//}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ValidToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}
	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))
	if uid != n {
		return false
	}
	return true
}
