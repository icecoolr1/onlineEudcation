package test

import (
	"fmt"
	"gorm.io/gorm"
	"onlineEudcation/Student/dao"
	"onlineEudcation/Student/etity"
	"testing"
)

var studentDao dao.StudentDaoInterface = new(dao.StudentDao)

func TestAdd(t *testing.T) {
	student := etity.Student{
		StudentPassword: "123",
		StudentName:     "我是猪",
		StudentEmail:    "dqdqad",
	}
	studentDao.InsertStudent(student)
}

func TestFindAll(t *testing.T) {
	students := studentDao.SelectAll()
	for _, student := range students {
		fmt.Println(student)
	}
}

func TestFindByID(t *testing.T) {
	student := studentDao.SelectByID(1)
	fmt.Println(student)
}

func TestUpdate(t *testing.T) {
	student := etity.Student{
		Model:        gorm.Model{ID: 1},
		StudentEmail: "testUpdate",
	}
	studentDao.UpdateStudent(student)
}

func TestDelete(t *testing.T) {
	studentDao.DeleteStudent(1)
}

func TestFindByName(t *testing.T) {
	student, _ := studentDao.FindStudentByName("我是猪")
	fmt.Println(student)
}
