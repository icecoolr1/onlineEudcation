package Test

import (
	"fmt"
	"onlineEudcation/Classes/Dao"
	"onlineEudcation/Classes/Etity"
	"testing"
)

var classesDao Dao.ClassesDaoInterface = new(Dao.ClassesDao)
var classesList []Etity.Classes

func TestAdd(t *testing.T) {
	class := Etity.Classes{
		CourseId:  32,
		StudentId: 3,
	}
	classesList = append(classesList, class)
	classesDao.AddClass(classesList)
}

func TestDelete(t *testing.T) {
	classesDao.DelClass([]int{1})
}

func TestFindAll(t *testing.T) {
	classes := classesDao.FindAllClasses()
	fmt.Println(classes)
}
