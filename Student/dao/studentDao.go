package dao

import (
	"onlineEudcation/Student/etity"
	"onlineEudcation/Tools/scripts"
)

var DB = scripts.GetDatabaseConnection()
var studentList []etity.Student

type StudentDao struct{}

func (StudentDao) SelectAll() []etity.Student {
	DB.Find(&studentList)
	return studentList
}

func (StudentDao) SelectByID(id int) etity.Student {
	var student etity.Student
	DB.Where("id = ?", id).First(&student)
	return student
}

func (StudentDao) InsertStudent(stu etity.Student) bool {
	DB.Create(&stu)
	return true
}

func (StudentDao) UpdateStudent(stu etity.Student) bool {
	DB.Model(&etity.Student{}).Where("id = ?", stu.ID).Updates(stu)
	return true
}

func (StudentDao) DeleteStudent(id int) bool {
	DB.Where("id = ?", id).Delete(&etity.Student{})
	return true
}

func (d StudentDao) SelectStudentByName(name string) []etity.Student {
	DB.Where("s_name LIKE ?", "%"+name+"%").Find(&studentList)
	return studentList
}

func (d StudentDao) FindStudentByName(name string) (etity.Student, error) {
	var student etity.Student
	db := DB.Where("s_name = ?", name).First(&student)
	if db.Error != nil {
		return student, db.Error
	} else {
		return student, nil
	}
}
