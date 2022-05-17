package etity

import "gorm.io/gorm"

type Censorship struct {
	gorm.Model
	CourseId  int    `gorm:"column:course_id"`
	TeacherId int    `gorm:"column:teacher_id"`
	Result    bool   `gorm:"column:result"`
	Message   string `gorm:"column:message"`
}
