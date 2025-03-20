package repositories

import (
	"errors"
	"log"

	"github.com/NguyenAnhQuan-Dev/NKbook-API/internal/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(userName string, email string, password string, roleName string) (models.User, error)
	UpdateUser(id string, payload map[string]any) (models.User, error)
	GetAllUser() ([]models.User, error)
	GetUserByID(id string) (models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	IsEmailExist(email string) (bool, error)
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
func (u *userRepository) CreateUser(userName string, email string, password string, roleName string) (models.User, error) {
	tx := u.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Println("Recovered in CreateUserWithRole:", r)
		}
	}()
	newUser := models.User{
		UserName: userName,
		Email:    email,
		Password: password,
	}
	err := tx.Create(&newUser).Error
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	var role models.Role
	if err := tx.Where("role_name = ?", roleName).First(&role).Error; err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	if err := tx.Model(&newUser).Association("Roles").Append(&role); err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	tx.Commit()

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

// GetuserByEmail implements IUserRepository.
func (u *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	err := u.db.Where("email = ?", email).First(&user).Error

	return &user, err
}

// IsEmailExist implements IUserRepository.
func (u *userRepository) IsEmailExist(email string) (bool, error) {
	var user models.User
	err := u.db.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
