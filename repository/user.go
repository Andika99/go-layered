package repository

import (
	"go-layered/models"
	"gorm.io/gorm"
)

// Interface
type UserRepository interface {
	FindAll() ([]models.User, error)
	FindById(id string) (models.User, error)
	Create(user models.User) (models.User, error)
	Update(user models.User) (models.User, error)
	Patch(id string, updates map[string]interface{}) (models.User, error)
	Delete(id string) error
}

// Struct
type userRepository struct {
	db *gorm.DB
}

// Contructor
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// Methods
func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User

	err := r.db.Find(&users).Error
	
	return users, err
}

func (r *userRepository) FindById(id string) (models.User, error) {
	var user models.User
	
	err := r.db.First(&user, id).Error
	
	return user, err
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) Update(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) Patch(id string, updates map[string]interface{}) (models.User, error) {
	var user models.User

	err := r.db.Model(&user).Where("id = ?", id).Updates(updates).Error

	return user, err	
}

func (r *userRepository) Delete(id string) error {
	err := r.db.Delete(&models.User{}, id).Error
	return err
}
