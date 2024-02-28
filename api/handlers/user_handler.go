package handlers

import (
	"net/http"
	"strconv"

	"test_backend_developer_go/models"
	"test_backend_developer_go/repository"

	"github.com/gin-gonic/gin"
)

// UserHandler adalah struct untuk menangani permintaan terkait pengguna
type UserHandler struct {
	UserRepository *repository.UserRepository
}

// NewUserHandler membuat instance UserHandler baru
func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: userRepo,
	}
}

// CreateUser membuat pengguna baru
func (uh *UserHandler) CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uh.UserRepository.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat pengguna baru"})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

// GetUsers mengambil semua pengguna
func (uh *UserHandler) GetUsers(c *gin.Context) {
	users, err := uh.UserRepository.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil pengguna"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUser mengambil detail pengguna berdasarkan ID
func (uh *UserHandler) GetUser(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID pengguna tidak valid"})
		return
	}

	// Panggil repository untuk mendapatkan detail user
	userDetail, err := uh.UserRepository.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, userDetail)
}

// UpdateUser memperbarui informasi pengguna berdasarkan ID
func (uh *UserHandler) UpdateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID pengguna tidak valid"})
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser.ID = int(userID) // set ID user dengan ID yang diterima

	if err := uh.UserRepository.UpdateUser(&updatedUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui pengguna"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser menghapus pengguna berdasarkan ID
func (uh *UserHandler) DeleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID pengguna tidak valid"})
		return
	}

	user := &models.User{ID: int(userID)}

	if err := uh.UserRepository.DeleteUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus pengguna"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pengguna berhasil dihapus"})
}
