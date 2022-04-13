package Etity

import "gorm.io/gorm"

// Classes
/*
	由学生使用，记录学生所订阅课程
*/
type Classes struct {
	gorm.Model
	CourseId  int32 `gorm:"column:c_id"`
	StudentId int32 `gorm:"column:s_id"`
}
