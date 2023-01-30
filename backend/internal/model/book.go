package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Book struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;"`
	Title      string         `json:"title" gorm:"not null"`
	Author     string         `json:"author" gorm:"not null"`
	Uri        string         `json:"uri" gorm:"not null"`
	Categories pq.StringArray `json:"categories" gorm:"type:text[];not null;default:'{}'"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return
}
