package Dao

import "onlineEudcation/Classes/Etity"

type ClassesDaoInterface interface {
	AddClass(class Etity.Classes) bool
	DelClass(studentId int, courseId int) bool
	UpdateClass(classes Etity.Classes) bool
	FindClass(classId int) Etity.Classes
	FindAllClasses(studentId int) []Etity.Classes
	IsFav(studentId int, courseId int) bool
}
