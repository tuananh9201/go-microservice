package main

import (
	"context"

	common "github.com/tuananh9201/go-eco/common"
)

type CreateUserStorage interface {
	CreateNewUser(ctx context.Context, data *User) error
	GetUsers(ctx context.Context, paging *common.Paging, filter *UserFilter) ([]*User, *common.Paging, error)
	GetUserByID(ctx context.Context, id int) (*User, error)
	UpdateUser(ctx context.Context, id int, data *User) error
	DeleteUser(ctx context.Context, id int) error
}

type userUseCase struct {
	store CreateUserStorage
}

func NewUserUseCase(store CreateUserStorage) *userUseCase {
	return &userUseCase{store: store}
}

func (uc *userUseCase) CreateNewUser(ctx context.Context, data *User) error {
	if err := uc.store.CreateNewUser(ctx, data); err != nil {
		return err
	}
	return nil
}

func (uc *userUseCase) GetUsers(ctx context.Context, paging *common.Paging, filter *UserFilter) ([]*User, *common.Paging, error) {
	users, paging, err := uc.store.GetUsers(ctx, paging, filter)
	if err != nil {
		return nil, nil, err
	}
	return users, paging, nil
}

func (uc *userUseCase) GetUserByID(ctx context.Context, id int) (*User, error) {
	user, err := uc.store.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *userUseCase) UpdateUser(ctx context.Context, id int, data *User) error {
	if err := uc.store.UpdateUser(ctx, id, data); err != nil {
		return err
	}
	return nil
}

func (uc *userUseCase) DeleteUser(ctx context.Context, id int) error {
	if err := uc.store.DeleteUser(ctx, id); err != nil {
		return err
	}
	return nil
}
