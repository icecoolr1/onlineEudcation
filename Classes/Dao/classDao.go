package Dao

import (
	"onlineEudcation/Classes/Etity"
	"onlineEudcation/Tools/scripts"
)

type ClassesDao struct{}

var DB = scripts.GetDatabaseConnection()
var classesList []Etity.Classes

func (ClassesDao) AddClass(class Etity.Classes) bool {
	DB.Create(&class)
	return true
}

func (ClassesDao) DelClass(studentId int, courseId int) bool {
	DB.Where("s_id = ? and c_id = ?", studentId, courseId).Delete(&Etity.Classes{})
	return true
}

func (ClassesDao) UpdateClass(classes Etity.Classes) bool {
	DB.Model(&Etity.Classes{}).Where("id = ? ", classes.Model.ID).Updates(classes)
	return true
}

func (ClassesDao) FindClass(classId int) Etity.Classes {
	class := Etity.Classes{}
	DB.Where("id = ?", classId).First(&class)
	return class
}

func (ClassesDao) FindAllClasses(studentId int) []Etity.Classes {
	DB.Where("s_id = ?", studentId).Find(&classesList)
	return classesList
}

func (d ClassesDao) IsFav(studentId int, courseId int) bool {
	first := DB.Where("s_id = ? and c_id = ?", studentId, courseId).First(&Etity.Classes{})
	if first.RowsAffected > 0 {
		// 已经被收藏
		return true
	} else {
		// 未被收藏
		return false
	}
}
