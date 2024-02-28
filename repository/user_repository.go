package repository

import (
	"test_backend_developer_go/models"

	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository inisialisasi repository pengguna baru
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// CreateUser membuat pengguna baru dalam database
func (ur *UserRepository) CreateUser(user *models.User) error {
	return ur.DB.Create(user).Error
}

// GetAllUsers mengambil semua pengguna dari database
func (ur *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := ur.DB.Find(&users).Error
	return users, err
}

// GetUserByID mengambil pengguna berdasarkan ID dari database
func (ur *UserRepository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	err := ur.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser mengupdate pengguna dalam database
func (ur *UserRepository) UpdateUser(user *models.User) error {
	return ur.DB.Save(user).Error
}

// DeleteUser menghapus pengguna dari database
func (ur *UserRepository) DeleteUser(user *models.User) error {
	return ur.DB.Delete(user).Error
}
