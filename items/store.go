package main

import (
	"context"

	"github.com/tuananh9201/commons"
	"gorm.io/gorm"
)

type sqlStore struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *sqlStore {
	return &sqlStore{db}
}

func (s *sqlStore) CreateItem(ctx context.Context, data *ItemCreate) error {
	item := &Item{
		Name:  data.Name,
		Price: data.Price,
	}
	if err := s.db.Create(*item).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) GetItem(ctx context.Context, paging *common.Paging, filter *ItemFilter) ([]*Item, *common.Paging, error) {
	var items []*Item
	db := s.db
	if filter == nil {
		filter = &ItemFilter{}
	}
	if filter.Name != "" {
		db = db.Where("name LIKE ?", "%"+filter.Name+"%")
	}
	paging.Process()
	offset := (paging.Page - 1) * paging.Limit
	limit := paging.Limit
	var total int64
	if err := db.Model(&Item{}).Count(&total).Error; err != nil {
		return nil, nil, err
	}
	db = db.Offset(offset).Limit(limit)
	if err := db.Find(&items).Error; err != nil {
		return nil, nil, err
	}
	paging.Total = int(total)
	return items, paging, nil
}

func (s *sqlStore) GetItemByID(ctx context.Context, id int) (*Item, error) {
	item := &Item{}
	if err := s.db.First(item, id).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (s *sqlStore) UpdateItem(ctx context.Context, id int, data *ItemCreate) error {
	item := &Item{
		Name:  data.Name,
		Price: data.Price,
	}
	if err := s.db.Model(&Item{}).Where("id = ?", id).Updates(item).Error; err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) DeleteItem(ctx context.Context, id int) error {
	if err := s.db.Delete(&Item{}, id).Error; err != nil {
		return err
	}
	return nil
}
