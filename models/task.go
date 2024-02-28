package models

import (
	"time"
)

type Task struct {
	ID          int       `gorm:"primary_key" json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"default:'pending'"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
