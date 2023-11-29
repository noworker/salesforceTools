package repositories

import (
	"github.com/noworker/salesforceTools/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *model.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(model *model.User) error {
	if err := ur.db.Create(model).Error; err != nil {
		return err
	}
	return nil
}
