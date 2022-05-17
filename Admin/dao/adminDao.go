package dao

import (
	"onlineEudcation/Admin/etity"
	"onlineEudcation/Courses/Etity"
	"onlineEudcation/Tools/scripts"
)

var DB = scripts.GetDatabaseConnection()
var courseList []Etity.Course
var resList []etity.Censorship

type AdminDao struct{}

func (AdminDao) GetAllCensorCourses() []Etity.Course {
	DB.Where("course_status = ?", 0).Order("UpdatedAt desc").Find(&courseList)
	return courseList
}

func (AdminDao) AddCensorship(c etity.Censorship) bool {
	DB.Create(&etity.Censorship{
		CourseId:  c.CourseId,
		TeacherId: c.TeacherId,
		Result:    c.Result,
		Message:   c.Message,
	})
	return true
}

func (AdminDao) ChangeStatus(courseId int) bool {
	DB.Model(&Etity.Course{}).Where("ID = ?", courseId).Update("course_status", 1)
	return true
}

func (AdminDao) GetResult(teacherId int) []etity.Censorship {
	DB.Where("teacher_id = ?", teacherId).Order("UpdatedAt desc").Find(&resList)
	return resList
}
