package dao

import "onlineEudcation/Teacher/etity"

type TeacherInterface interface {
	// AddTeacher 添加教师
	AddTeacher(teacher []etity.Teacher) (bool, error)
	// DelTeacher 按Id号删除教师
	DelTeacher(id []int) bool
	// UpdateTeacher 更新教师信息
	UpdateTeacher(teacher etity.Teacher) bool
	// FindTeacherById 通过教师ID查询
	FindTeacherById(id int) etity.Teacher
	// FindTeacherByName 通过教师名查询
	FindTeacherByName(name string) []etity.Teacher
	// FindAllTeachers 查询所有教师
	FindAllTeachers() []etity.Teacher
	// FindTeacherByEmail 通过邮箱查询教师信息
	FindTeacherByEmail(email string) (*etity.Teacher, error)
}
