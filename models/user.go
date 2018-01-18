package models

import "github.com/jinzhu/gorm"

//User usuario de la api
type User struct {
	gorm.Model
	Username        string `json:"username" gorm:"not null;unique"`
	Email           string `json:"email" gorm:"not null; unique"`
	Namefull        string `json:"namefull" gorm:"not null"`
	Password        string `json:"password,omitempty" gorm:"not null;type:varchar(256)"`
	ConfirmPassword string `json:"confirmPassword,omitempty" gorm:"-"`
	Perfil          string `json:"perfil" gorm:"default:'Usuario'"`
	Avatar          string `json:"avatar"`
}
