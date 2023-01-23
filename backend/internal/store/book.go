package store

import (
	"github.com/google/uuid"
	"github.com/hfl0506/reader-app/internal/model"
	"gorm.io/gorm"
)

type BookStore struct {
	db *gorm.DB
}

func NewBookStore(db *gorm.DB) *BookStore {
	return &BookStore{
		db: db,
	}
}

func (bs *BookStore) GetAll() ([]*model.Book, error) {
	var books []*model.Book

	if err := bs.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (bs *BookStore) GetByTitle(title string) ([]*model.Book, error) {
	var books []*model.Book

	if err := bs.db.Where(&model.Book{ Title: title}).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (bs *BookStore) GetByAuthor(author string) ([]*model.Book, error) {
	var books []*model.Book

	if err := bs.db.Where(&model.Book{Author: author}).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (bs *BookStore) GetByCategories(category []string) ([]*model.Book, error) {
	var books []*model.Book

	if err := bs.db.Where("categories @> ?", category).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (bs *BookStore) GetById(id uuid.UUID) (*model.Book, error) {
	var book *model.Book

	if err := bs.db.Find(&model.Book{ ID: id}).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (bs *BookStore) CreateOne(payload )
