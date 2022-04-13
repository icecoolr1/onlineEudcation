package Dao

import (
	"fmt"
	"onlineEudcation/Courses/Etity"
	"onlineEudcation/Tools/scripts"
)

type CourseDao struct{}

var DB = scripts.GetDatabaseConnection()
var courseList []Etity.Course

func (CourseDao) AddCourse(courses []Etity.Course) bool {
	for _, e := range courses {
		fmt.Printf("%v\n", e)
		DB.Create(&e)
	}
	return true
}

func (CourseDao) DelCourse(courseId []int) bool {
	tx := DB.Where("id in (?)", courseId).Delete(&Etity.Course{})
	if tx.Error != nil {
		return false
	} else {
		return true
	}
}

func (CourseDao) UpdateCourse(course Etity.Course) bool {
	updates := DB.Model(&Etity.Course{}).Where("id = ? ", course.ID).Updates(course)
	if updates.Error != nil {
		return false
	} else {
		return true
	}
}

func (CourseDao) FindCourseByCourseId(courseId int) Etity.Course {
	course := Etity.Course{}
	DB.Where("id = ?", courseId).First(&course)
	return course
}

func (CourseDao) FindCourseByCourseName(courseName string) []Etity.Course {
	DB.Where("course_name LIKE ?", "%"+courseName+"%").Find(&courseList)
	return courseList
}

func (CourseDao) FindAllCourses() []Etity.Course {
	DB.Find(&courseList)
	return courseList
}

func (d CourseDao) FindCourseByTeacherId(teacherId int) []Etity.Course {
	DB.Where("t_id = ?", teacherId).Find(&courseList)
	return courseList
}
