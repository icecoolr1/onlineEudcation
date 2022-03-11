package etity

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Password string `gorm:"column:t_password"`
	Email    string `gorm:"column:t_email"`
	Name     string `gorm:"column:t_name"`
}
