package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// NewDBConnection membuat koneksi baru ke database
func NewDBConnection() (*gorm.DB, error) {
	// Ganti dengan konfigurasi koneksi database Anda
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=test_backend_go_db sslmode=disable password=admin")
	if err != nil {
		return nil, err
	}

	// Aktifkan mode log SQL (Opsional)
	db.LogMode(true)

	return db, nil
}
