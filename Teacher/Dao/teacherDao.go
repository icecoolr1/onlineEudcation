package dao

import (
	"fmt"
	"onlineEudcation/Teacher/etity"
	"onlineEudcation/Tools/scripts"
)

type TeacherDao struct{}

var users []etity.Teacher
var DB = scripts.GetDatabaseConnection()

func (t TeacherDao) AddTeacher(teacher []etity.Teacher) bool {
	for _, e := range teacher {
		fmt.Printf("%v\n", e)
		DB.Create(&e)
	}
	return true
}

func (t TeacherDao) DelTeacher(id []int) bool {
	DB.Where("id in (?)", id).Delete(&etity.Teacher{})
	return true
}

func (t TeacherDao) UpdateTeacher(teacher etity.Teacher) bool {
	DB.Model(&etity.Teacher{}).Where("id = ? ", teacher.ID).Updates(teacher)
	return true
}

func (t TeacherDao) FindTeacherById(id int) etity.Teacher {
	teacher := etity.Teacher{}
	DB.Where("id = ?", id).First(&teacher)
	return teacher
}

func (t TeacherDao) FindTeacherByName(name string) []etity.Teacher {
	DB.Where("t_name LIKE ?", "%"+name+"%").Find(&users)
	return users
}

func (t TeacherDao) FindAllTeachers() []etity.Teacher {
	DB.Find(&users)
	return users
}
