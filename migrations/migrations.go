package migrations

import (
	"test_backend_developer_go/models"

	"github.com/jinzhu/gorm"
)

// AutoMigrate berfungsi untuk melakukan migrasi otomatis pada tabel-tabel
func AutoMigrate(db *gorm.DB) error {
	// Migrate model-model Anda di sini
	db.Model(&models.Task{}).AddForeignKey("usuer_id", "users(id)", "RESTRICT", "RESTRICT")
	if err := db.AutoMigrate(&models.User{}, &models.Task{}).Error; err != nil {
		return err
	}
	return nil
}
