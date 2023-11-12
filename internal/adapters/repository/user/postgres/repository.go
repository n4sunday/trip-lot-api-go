package repository

import (
	"trip-lot-api/internal/adapters/model"
	core "trip-lot-api/internal/core/user"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserPgRepository struct {
	db *gorm.DB
}

func NewUserPgRepository(db *gorm.DB) *UserPgRepository {
	return &UserPgRepository{
		db: db,
	}
}

func (r *UserPgRepository) FindAll() ([]core.User, error) {
	var users []model.User

	// Find from database
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err

	}

	var result []core.User
	// Copy from model to domain
	if err := copier.Copy(&result, &users); err != nil {
		return nil, err
	}

	return result, nil
}

func (r *UserPgRepository) FindByID(id string) (core.User, error) {
	var user model.User

	// Find from database
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return core.User{}, err
	}

	var result core.User
	// Copy from model to domain
	if err := copier.Copy(&result, &user); err != nil {
		return core.User{}, err
	}

	return result, nil
}

func (r *UserPgRepository) FindByEmail(email string) (core.User, error) {
	var user model.User

	// Find from database
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return core.User{}, err
	}

	var result core.User
	// Copy from model to domain
	if err := copier.Copy(&result, &user); err != nil {
		return core.User{}, err
	}

	return result, nil
}

func (r *UserPgRepository) IsUserExist(email string) (bool, error) {
	var count int64

	if err := r.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (r *UserPgRepository) CreateUser(payload core.User) error {
	var user model.User
	copier.Copy(&user, &payload)

	result := r.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserPgRepository) UpdateUser(user core.User) error {
	result := r.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserPgRepository) DeleteUser(id string) error {
	result := r.db.Where("id = ?", id).Delete(&model.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
