package repository

import (
	"myfinnplan/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user entity.User) (entity.User, error)
	Edit(user entity.User) (entity.User, error)
	FindById(id int) ([]entity.User, error)
	FindByUserName(userName string) ([]entity.User, error)

	FindAll() ([]entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) Edit(user entity.User) (entity.User, error) {
	err := r.db.Save(user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindById(id int) ([]entity.User, error) {
	var user []entity.User

	err := r.db.Where("id = ? AND deleted_by = ?", id, "").Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
func (r *userRepository) FindByUserName(userName string) ([]entity.User, error) {
	var user []entity.User

	err := r.db.Where("user_name = ? AND deleted_by = ?", userName, "").Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User

	err := r.db.Where("deleted_by = ?", "").Find(&users).Error

	if err != nil {
		return users, err
	}

	return users, nil
}
