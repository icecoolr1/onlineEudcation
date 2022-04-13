package Dao

import "onlineEudcation/Courses/Etity"

type CourseDaoInterface interface {
	AddCourse(course []Etity.Course) bool
	DelCourse(courseId []int) bool
	UpdateCourse(course Etity.Course) bool
	FindCourseByCourseId(courseId int) Etity.Course
	FindCourseByCourseName(courseName string) []Etity.Course
	FindAllCourses() []Etity.Course
}
