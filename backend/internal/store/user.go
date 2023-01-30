package store

import (
	"github.com/google/uuid"
	"github.com/hfl0506/reader-app/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreateUserParam struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) GetUserById(id uuid.UUID) (*model.User, error) {
	var user *model.User

	if err := us.db.Find(&user).Where(id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserStore) GetUserByEmail(email string) (*model.User, error) {
	var user *model.User

	if err := us.db.Where("Email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserStore) CreateUser(payload *model.User) (*model.User, error) {
	result := us.db.Clauses(clause.Returning{}).Select("ID", "Name", "Email", "Password").Create(&payload)

	if result.Error != nil {
		return nil, result.Error
	}

	return payload, nil
}

func (us *UserStore) UpdateUser(id uuid.UUID, payload *model.User) error {
	return us.db.Updates(&payload).Error
}

func (us *UserStore) DeleteUser(id uuid.UUID) error {
	var user *model.User

	if err := us.db.Find(&user).Where("id = ?", id).Error; err != nil {
		return err
	}

	return us.db.Delete(&user).Error
}
