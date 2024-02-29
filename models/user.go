package models

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `json:"name" validate:"max=255"`
	Email     string    `json:"email" validate:"max=255" gorm:"unique"`
	Password  string    `json:"password" validate:"max=255"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// untuk melakukan validasi tipe data
func (req User) Validate() (errArr []error, err error) {
	v := validator.New()
	err = v.Struct(req)
	if err == nil {
		return nil, nil
	}
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			errArr = append(errArr, errors.New(e.Field()))
		case "max", "number", "alphanum", "base64", "email":
			errArr = append(errArr, errors.New(e.Field()))
		}
	}
	return errArr, err
}
