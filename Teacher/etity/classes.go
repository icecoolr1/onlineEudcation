package etity

import "gorm.io/gorm"

type Classes struct {
	gorm.Model
	CourseId  int32 `gorm:"column:c_id"`
	StudentId int32 `gorm:"column:s_id"`
}
