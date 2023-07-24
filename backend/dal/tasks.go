package dal

import (
	"context"

	"gorm.io/gorm"
)

func (Tasks) TableName() string {
	return "tasks"
}

type TaskDB struct {
	DB *gorm.DB
}

func NewTaskDB() TaskDB {
	return TaskDB{
		DB: newDB(),
	}
}

func (t *TaskDB) GetByUserID(ctx context.Context, userID int) ([]*Tasks, error) {
	tasks := []*Tasks{}
	err := t.DB.WithContext(ctx).Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskDB) Add(ctx context.Context, newTask *Tasks) error {
	return t.DB.WithContext(ctx).Create(&newTask).Error
}

func (t *TaskDB) Del(ctx context.Context, taskID string) error {
	return t.DB.WithContext(ctx).Delete(&Tasks{ID: taskID}).Error
}