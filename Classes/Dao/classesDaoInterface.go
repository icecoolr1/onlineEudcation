package Dao

import "onlineEudcation/Classes/Etity"

type ClassesDaoInterface interface {
	AddClass(classes []Etity.Classes) bool
	DelClass(classId []int) bool
	UpdateClass(classes Etity.Classes) bool
	FindClass(classId int) Etity.Classes
	FindAllClasses() []Etity.Classes
}
