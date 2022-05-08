package Test

import (
	"fmt"
	"gorm.io/gorm"
	"onlineEudcation/Courses/Dao"
	"onlineEudcation/Courses/Etity"
	"testing"
)

var courseDao Dao.CourseDaoInterface = new(Dao.CourseDao)

func TestAddCourse(t *testing.T) {
	courseList := make([]Etity.Course, 0)
	course := Etity.Course{
		CourseName:        "test",
		TeacherId:         1,
		CourseTag:         "编程",
		CoursePlaySum:     0,
		CourseVideoNumBer: 10,
		CourseStatus:      false,
	}
	// 添加至列表
	courseList = append(courseList, course)
	courseDao.AddCourse(courseList)
}

func TestDelCourse(t *testing.T) {
	courseDao.DelCourse([]int{1})
}

func TestUpdate(t *testing.T) {
	course := Etity.Course{
		Model:      gorm.Model{ID: 13},
		CourseName: "test",

		CourseTag: "数学",

		CourseStatus: false,
	}
	courseDao.UpdateCourse(course)
}

func TestFindAll(t *testing.T) {
	courseList := courseDao.FindAllCourses()
	fmt.Println("courseList:", courseList)
}

func TestFindBytID(t *testing.T) {
	courselist := courseDao.FindCourseByTeacherId(37)
	fmt.Println("courselist:", courselist)
}
func TestSearch(t *testing.T) {
	list := courseDao.GetCourseList("golang")
	fmt.Println("list:", list)
}

func TestRecommend(t *testing.T) {
	list := courseDao.GetRecommendCourseList()
	fmt.Println("list:", list)
	fmt.Println(len(list))
}

func TestHits(t *testing.T) {
	courseDao.CourseHits(10)

}
