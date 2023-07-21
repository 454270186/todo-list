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

func (s *ServiceHandler) Login(username, password string) error {
	log.Println(username, password)
	user, err := s.userDB.FindUser(context.Background(), username, password)
	if err != nil {
		return err
	}
	log.Println(user)

	return nil
}

func (s *ServiceHandler) Register(username, password string) error {
	isExist, err := s.userDB.IsUserExist(context.Background(), username)
	if err != nil {
		return errors.New("database error")
	}
	if isExist {
		return errors.New("user is exist")
	}

	err = s.userDB.CreateUser(context.Background(), username, password)
	if err != nil {
		return errors.New("database error")
	}

	return nil
}