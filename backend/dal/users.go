package dal

import (
	"context"
	"errors"
	"log"

	"gorm.io/gorm"
)

func (Users) TableName() string {
	return "users"
}

type UserDB struct {
	DB *gorm.DB
}

func NewUserDB() UserDB {
	return UserDB{
		DB: newDB(),
	}
}

func (u *UserDB) CreateUser(ctx context.Context, username, password string) (int ,error) {
	newUser := Users{
		Username: username,
		Password: password,
	}

	err := DB.WithContext(ctx).Create(&newUser).Error
	if err != nil {
		log.Println(err)
		return 0, errors.New("database error")
	}

	return newUser.ID, nil
}

func (u *UserDB) FindUser(ctx context.Context, username, password string) (*Users, error) {
	user := Users{
		Username: username,
		Password: password,
	}
	err := DB.WithContext(ctx).Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user is not exist")
		}
		return nil, errors.New("database error")
	}

	return &user, nil
}

func (u *UserDB) IsUserExist(ctx context.Context, username string) (bool, error) {
	var user Users
	err := DB.WithContext(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}