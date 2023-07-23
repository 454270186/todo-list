package service

import (
	"context"
	"errors"
	"log"

	"todolist/dal"
)

type ServiceHandler struct {
	userDB dal.UserDB
}

func NewServiceHandler() *ServiceHandler {
	return &ServiceHandler{
		userDB: dal.NewUserDB(),
	}
}

func (s *ServiceHandler) Login(username, password string) (int, error) {
	log.Println(username, password)
	user, err := s.userDB.FindUser(context.Background(), username, password)
	if err != nil {
		return 0, err
	}
	log.Println(user)

	return user.ID, nil
}

func (s *ServiceHandler) Register(username, password string) (int, error) {
	isExist, err := s.userDB.IsUserExist(context.Background(), username)
	if err != nil {
		return 0, errors.New("database error")
	}
	if isExist {
		return 0, errors.New("user is exist")
	}

	newUserID, err := s.userDB.CreateUser(context.Background(), username, password)
	if err != nil {
		return 0, errors.New("database error")
	}

	return newUserID, nil
}