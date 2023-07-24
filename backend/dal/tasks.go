package dal

import (
	"context"
	"log"

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

func (t *TaskDB) UpdateName(ctx context.Context, taskID, newName string) error {
	log.Println("asdasda")
	return t.DB.WithContext(ctx).Model(&Tasks{ID: taskID}).Update("name", newName).Error
}

func (t *TaskDB) UpdateCompleted(ctx context.Context, taskID string) error {
	var task Tasks
    err := t.DB.WithContext(ctx).Where("id = ?", taskID).First(&task).Error
    if err != nil {
        return err
    }

    task.Completed = !task.Completed

    err = t.DB.WithContext(ctx).Save(&task).Error
    if err != nil {
        return err
    }

    return nil
}