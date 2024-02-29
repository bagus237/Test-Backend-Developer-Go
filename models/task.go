package models

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type Task struct {
	ID          int       `gorm:"primary_key" json:"id"`
	User        User      `json:"-" gorm:"foreignKey:user_id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title" validate:"max=255"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"default:'pending'" validate:"max=50"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// untuk melakukan validasi tipe data
func (req Task) Validate() (errArr []error, err error) {
	v := validator.New()
	err = v.Struct(req)
	if err == nil {
		return nil, nil
	}
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			errArr = append(errArr, errors.New(e.Field()))
			err = errors.New("invalid input")
		case "max", "number", "alphanum", "base64", "email":
			errArr = append(errArr, errors.New(e.Field()))
			err = errors.New("invalid input")
		}
	}
	return errArr, err
}
