package model

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	Title string `json:"title" gorm:"not null"`
	Author string `json:"author" gorm:"not null"`
	Uri string `json:"uri" gorm:"not null"`
	Categories []string `json:"categories" gorm:"not null;default:'{}'"`
	CreatedAt time.Time `json:"created_at" gorm:"default:"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:"`
}