package repositories

import (
	"github.com/noworker/salesforceTools/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByNameId(user *model.User, userId string) error
	CreateUser(user *model.User) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUserByNameId(user *model.User, userId string) error {
	err := ur.db.Where("name=?", userId).First(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) CreateUser(model *model.User) error {
	if err := ur.db.Create(model).Error; err != nil {
		return err
	}
	return nil
}
