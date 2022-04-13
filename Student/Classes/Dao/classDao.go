package Dao

import (
	"fmt"
	"onlineEudcation/Classes/Etity"
	"onlineEudcation/Tools/scripts"
)

type ClassesDao struct{}

var DB = scripts.GetDatabaseConnection()
var classesList []Etity.Classes

func (ClassesDao) AddClass(classes []Etity.Classes) bool {
	for _, e := range classes {
		fmt.Printf("%v\n", e)
		DB.Create(&e)
	}
	return true
}

func (ClassesDao) DelClass(classId []int) bool {
	DB.Where("id in (?)", classId).Delete(&Etity.Classes{})
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

func (ClassesDao) FindAllClasses() []Etity.Classes {
	DB.Find(&classesList)
	return classesList
}
