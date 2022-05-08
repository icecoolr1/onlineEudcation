package Dao

import (
	"fmt"
	"gorm.io/gorm"
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

func (d CourseDao) GetCourseList(tag string) []Etity.Course {
	DB.Where("course_tag = ? and course_status = 1", tag).Order("course_play_sum desc").Find(&courseList)
	return courseList
}

func (d CourseDao) GetRecommendCourseList() []Etity.Course {
	DB.Where("course_status = 1").Order("course_play_sum desc").Limit(3).Find(&courseList)
	return courseList
}

func (d CourseDao) GetRecommendCourseListPersonally(tag string) []Etity.Course {
	DB.Where("course_tag = ? and course_status = 1", tag).Order("course_play_sum desc").Limit(2).Find(&courseList)
	return courseList
}

func (d CourseDao) CourseHits(courseId int) bool {
	course := Etity.Course{
		Model: gorm.Model{ID: uint(courseId)},
	}
	DB.Model(&course).Update("course_play_sum", gorm.Expr("course_play_sum + ?", 1))
	return true
}
