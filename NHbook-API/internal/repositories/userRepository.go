package repositories

import (
	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(userName string, email string, password string) (models.User, error)
	UpdateUser(id string, payload map[string]any) (models.User, error)
	GetAllUser() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{
		db: db,
	}
}

// CreateUser implements IUserRepository.
func (u *userRepository) CreateUser(userName string, email string, password string) (models.User, error) {
	newUser := models.User{
		UserName: userName,
		Email:    email,
		Password: password,
	}
	err := u.db.Create(&newUser).Error
	return newUser, err
}

// GetAllUser implements IUserRepository.
func (u *userRepository) GetAllUser() ([]models.User, error) {
	var users []models.User
	err := u.db.Find(&users).Error
	return users, err
}

// GetUserByID implements IUserRepository.
func (u *userRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	err := u.db.Where("user_id = ?", id).First(&user).Error
	return user, err
}

// UpdateUser implements IUserRepository.
func (u *userRepository) UpdateUser(id string, payload map[string]any) (models.User, error) {
	var userUpdated models.User
	err := u.db.Model(&models.User{}).Updates(payload).Error
	if err != nil {
		return userUpdated, err
	}

	errFind := u.db.Where("user_id = ?", id).Find(&userUpdated).Error

	return userUpdated, errFind
}
