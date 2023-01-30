package store

import (
	"github.com/google/uuid"
	"github.com/hfl0506/reader-app/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreateBookParam struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	Uri        string `json:"uri"`
	Categories string `json:"categories"`
}

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

	if err := bs.db.Where(&model.Book{Title: title}).Find(&books).Error; err != nil {
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

	if err := bs.db.Find(&model.Book{ID: id}).First(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (bs *BookStore) CreateOne(payload *model.Book) (*model.Book, error) {
	result := bs.db.Clauses(clause.Returning{}).Select("ID", "Title", "Author", "Uri", "Categories").Create(&payload)

	if result.Error != nil {
		return nil, result.Error
	}

	return payload, nil
}

func (bs *BookStore) UpdateOne(id uuid.UUID, payload *model.Book) error {
	return bs.db.Updates(&payload).Error
}

func (bs *BookStore) DeleteOne(id uuid.UUID) error {
	var book *model.Book

	if err := bs.db.Find(&book).Where("id = ?", id).Error; err != nil {
		return err
	}

	return bs.db.Delete(&book).Error
}
