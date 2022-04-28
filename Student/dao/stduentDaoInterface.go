package dao

import "onlineEudcation/Student/etity"

type StudentDaoInterface interface {
	// SelectAll 查询所有学生
	SelectAll() []etity.Student
	// SelectByID 根据ID查询学生
	SelectByID(id int) etity.Student
	// InsertStudent  插入学生信息
	InsertStudent(stu etity.Student) bool
	// UpdateStudent Update 更新学生信息
	UpdateStudent(stu etity.Student) bool
	// DeleteStudent Delete 删除学生信息
	DeleteStudent(id int) bool
	// SelectStudentByName SelectByName 根据姓名查询学生(模糊查询)
	SelectStudentByName(name string) []etity.Student
	// FindStudentByName 通过用户名查找学生
	FindStudentByName(name string) (etity.Student, error)
}
