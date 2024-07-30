package main

import (
	"context"

	"github.com/tuananh9201/commons"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *ItemCreate) error
	GetItem(ctx context.Context, paging *common.Paging, filter *ItemFilter) ([]*Item, *common.Paging, error)
	GetItemByID(ctx context.Context, id int) (*Item, error)
	UpdateItem(ctx context.Context, id int, data *ItemCreate) error
	DeleteItem(ctx context.Context, id int) error
}

type itemUsecase struct {
	store CreateItemStorage
}

func NewItemUsecase(store CreateItemStorage) *itemUsecase {
	return &itemUsecase{store: store}
}
func (uc *itemUsecase) CreateNewItem(ctx context.Context, data *ItemCreate) error {
	// logic
	if err := uc.store.CreateItem(ctx, data); err != nil {
		return err
	}
	return nil
}

func (uc *itemUsecase) GetItems(ctx context.Context, paging *common.Paging, filter *ItemFilter) ([]*Item, *common.Paging, error) {
	items, paging, err := uc.store.GetItem(ctx, paging, filter)
	if err != nil {
		return nil, nil, err
	}
	return items, paging, nil
}

func (uc *itemUsecase) GetItemByID(ctx context.Context, id int) (*Item, error) {
	item, err := uc.store.GetItemByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func (uc *itemUsecase) UpdateItem(ctx context.Context, id int, data *ItemCreate) error {
	// logic
	if err := uc.store.UpdateItem(ctx, id, data); err != nil {
		return err
	}
	return nil
}

func (uc *itemUsecase) DeleteItem(ctx context.Context, id int) error {
	// logic
	if err := uc.store.DeleteItem(ctx, id); err != nil {
		return err
	}
	return nil
}
