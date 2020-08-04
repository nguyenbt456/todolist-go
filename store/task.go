package store

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/nguyenbt456/todolist-go/database"
	"github.com/nguyenbt456/todolist-go/model"
)

// TaskStore is middleware between controller and database
type TaskStore struct {
	db *gorm.DB
}

// NewTaskStore create new TaskStore with custome DB
func NewTaskStore() *TaskStore {
	return &TaskStore{db: database.GetDB()}
}

// Create create new task
func (s *TaskStore) Create(name string, status model.TaskStatusType, taskType model.TaskTypeType, userID string) (*model.Task, error) {
	task := &model.Task{
		Name:      name,
		Status:    status,
		Type:      taskType,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.db.Create(task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

// UpdateStatus update status of task which user want to change
func (s *TaskStore) UpdateStatus(userID string, taskID string, status model.TaskStatusType) error {
	task := &model.Task{}

	if err := s.db.Model(task).Where("id = ? AND user_id = ?", taskID, userID).Update("status", status).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("TaskID is invalid")
		}
		return err
	}

	return nil
}

// FindByID find task by ID
func (s *TaskStore) FindByID(id string) (*model.Task, error) {
	task := &model.Task{}

	if err := s.db.Where("id = ?", id).First(task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return task, nil
}

// FindByDate find tasks by date
func (s *TaskStore) FindByDate(date time.Time) (*[]model.Task, error) {
	tasks := &[]model.Task{}

	err := s.db.Where("created_at < ? && created_at > ?", time.Now(), time.Now().AddDate(0, 0, -1)).Find(tasks).Error
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// DeleteByID delete task by ID
func (s *TaskStore) DeleteByID(id string) error {
	task := &model.Task{}

	if err := s.db.Where("id = ?", id).Delete(task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("TaskID is invalid")
		}
		return err
	}

	return nil
}
