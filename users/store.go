package main

import (
	"context"

	common "github.com/tuananh9201/commons"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db}
}

func (s *sqlStore) CreateNewUser(ctx context.Context, data *User) error {
	user := &User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
	}
	if err := s.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) GetUsers(ctx context.Context, paging *common.Paging, filter *UserFilter) ([]*User, *common.Paging, error) {
	var users []*User
	db := s.db
	if filter == nil {
		filter = &UserFilter{}
	}
	if filter.Email != "" {
		db = db.Where("email LIKE ?", "%"+filter.Email+"%")
	}
	paging.Process()
	offset := (paging.Page - 1) * paging.Limit
	limit := paging.Limit
	var total int64
	if err := db.Model(&User{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}
	db = db.Offset(offset).Limit(limit)
	if err := db.Find(&users).Error; err != nil {
		return nil, nil, err
	}
	paging.Total = int(total)
	return users, paging, nil
}

func (s *sqlStore) GetUserByID(ctx context.Context, id int) (*User, error) {
	user := &User{}
	if err := s.db.First(user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *sqlStore) UpdateUser(ctx context.Context, id int, data *User) error {
	user := &User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
	}
	_, err := s.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	if err := s.db.Model(&User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) DeleteUser(ctx context.Context, id int) error {
	_, err := s.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	if err := s.db.Model(&User{}).Where("id = ?", id).Updates(User{
		DeleteFlag: true,
	}).Error; err != nil {
		return err
	}
	return nil
}
