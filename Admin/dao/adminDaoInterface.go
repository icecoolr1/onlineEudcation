package dao

import (
	"onlineEudcation/Admin/etity"
	"onlineEudcation/Courses/Etity"
)

type AdminDaoInterface interface {
	// GetAllCensorCourses 查询所有待审核课程
	GetAllCensorCourses() []Etity.Course
	// AddCensorship 添加审核结果
	AddCensorship(censorship etity.Censorship) bool
	// ChangeStatus 更改课程状态
	ChangeStatus(courseId int) bool
	// GetResult 教师获取课程信息
	GetResult(teacherId int) []etity.Censorship
}
