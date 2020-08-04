package store

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/nguyenbt456/todolist-go/database"
	"github.com/nguyenbt456/todolist-go/model"
)

// UserStore is middleware between controller and database
type UserStore struct {
	db *gorm.DB
}

// NewUserStore create new UserStore with custome DB.
// Default value is local DB
func NewUserStore(customeDB ...*gorm.DB) *UserStore {
	userStore := &UserStore{db: customeDB[0]}
	if customeDB == nil {
		userStore.db = database.GetDB()
	}

	return userStore
}

// Create create new user
func (s *UserStore) Create(name, username, password, email string) (*model.User, error) {
	if name == "" {
		return nil, errors.New("Name is invalid")
	}
	if username == "" {
		return nil, errors.New("Username is invalid")
	}
	if password == "" {
		return nil, errors.New("Username is invalid")
	}

	user := &model.User{
		Name:     name,
		Username: username,
		Password: password,
		Email:    email,
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// FindByID find user by ID
func (s *UserStore) FindByID(id string) (*model.User, error) {
	user := &model.User{}

	if err := s.db.Where("id = ?", id).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// FindByUser find users by a user
func (s *UserStore) FindByUser(user model.User) (*[]model.User, error) {
	users := &[]model.User{}

	if err := s.db.Where(&user).Find(users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return users, nil
}

// FindByUsernameAndPassword find user by username and password
func (s *UserStore) FindByUsernameAndPassword(username, password string) (*model.User, error) {
	user := &model.User{}

	if err := s.db.Where("username = ? AND password = ?", username, password).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
