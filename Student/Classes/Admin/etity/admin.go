package etity

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	AdminName     string `gorm:"column:admin_name"`
	AdminPassword string `gorm:"column:admin_password"`
}
