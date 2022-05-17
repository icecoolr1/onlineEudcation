package test

import (
	"fmt"
	"onlineEudcation/Admin/dao"
	"onlineEudcation/Admin/etity"
	"testing"
)

var adminDao dao.AdminDaoInterface = new(dao.AdminDao)

func TestCourses(t *testing.T) {
	courses := adminDao.GetAllCensorCourses()
	for _, course := range courses {
		fmt.Println(course.TeacherId)
	}
}

func TestChange(t *testing.T) {
	adminDao.ChangeStatus(15)
}

func TestAdd(t *testing.T) {
	adminDao.AddCensorship(etity.Censorship{
		CourseId:  15,
		TeacherId: 38,
		Result:    true,
		Message:   "通过",
	})

}

func TestG(t *testing.T) {
	result := adminDao.GetResult(38)
	fmt.Println(result)
}
