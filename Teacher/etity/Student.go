package etity

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	StudentPassword string `gorm:"column:s_password"`
	StudentName     string `gorm:"column:s_name"`
	StudentEmail    string `gorm:"column:s_email"`
}
