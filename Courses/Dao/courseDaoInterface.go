package Dao

import "onlineEudcation/Courses/Etity"

type CourseDaoInterface interface {
	AddCourse(course []Etity.Course) bool
	DelCourse(courseId []int) bool
	UpdateCourse(course Etity.Course) bool
	FindCourseByCourseId(courseId int) Etity.Course
	FindCourseByTeacherId(teacherId int) []Etity.Course
	FindCourseByCourseName(courseName string) []Etity.Course
	FindAllCourses() []Etity.Course
	GetCourseList(tag string) []Etity.Course
	GetRecommendCourseList() []Etity.Course
	GetRecommendCourseListPersonally(tag string) []Etity.Course
	CourseHits(courseId int) bool
}
