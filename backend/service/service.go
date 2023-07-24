package service

import (
	"context"
	"errors"
	"log"

	"todolist/dal"
	"todolist/dto"
)

type ServiceHandler struct {
	userDB dal.UserDB
	taskDB dal.TaskDB
}

func NewServiceHandler() *ServiceHandler {
	return &ServiceHandler{
		userDB: dal.NewUserDB(),
		taskDB: dal.NewTaskDB(),
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

func (s *ServiceHandler) GetTasksByID(userID int) ([]*dto.Task, error) {
	tasks, err := s.taskDB.GetByUserID(context.Background(), userID)
	if err != nil {
		return nil, errors.New("database error: " + err.Error())
	}

	dtoTasks := []*dto.Task{}
	for _, task := range tasks {
		dtoTask := dto.Task{
			ID: task.ID,
			UserID: task.UserID,
			Name: task.Name,
			Completed: task.Completed,
		}

		dtoTasks = append(dtoTasks, &dtoTask)
	}

	return dtoTasks, nil
}

func (s *ServiceHandler) AddNewTask(userID int, name string, taskID string) error {
	newTask := dal.Tasks{
		ID: taskID,
		UserID: userID,
		Name: name,
		Completed: false,
	}

	err := s.taskDB.Add(context.Background(), &newTask)
	if err != nil {
		return errors.New("database error: " + err.Error())
	}

	return nil
}

func (s *ServiceHandler) DelTaskByID(taskID string) error {
	return s.taskDB.Del(context.Background(), taskID)
}

func (s *ServiceHandler) ChangeName(taskID, newName string) error {
	return s.taskDB.UpdateName(context.Background(), taskID, newName)
}

func (s *ServiceHandler) ChangeCompleted(taskID string) error {
	return s.taskDB.UpdateCompleted(context.Background(), taskID)
}