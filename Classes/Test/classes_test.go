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
		CourseId:  10,
		StudentId: 7,
	}

	classesDao.AddClass(class)
}

func TestDelete(t *testing.T) {
	classesDao.DelClass(7, 10)
}

func TestFindAll(t *testing.T) {
	classes := classesDao.FindAllClasses(7)
	fmt.Println(classes)
}

func TestIsFav(t *testing.T) {
	fav := classesDao.IsFav(7, 10)
	fmt.Println(fav)

}
