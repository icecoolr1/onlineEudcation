package Test

import (
	"fmt"
	dao "onlineEudcation/Teacher/Dao"
	"onlineEudcation/Teacher/etity"
	"testing"
)

var teacherDao dao.TeacherInterface = new(dao.TeacherDao)

func TestDel(t *testing.T) {
	teacherDao.DelTeacher([]int{23})
}

func TestAdd(t *testing.T) {
	teacherlist := make([]etity.Teacher, 0)
	teacher1 := etity.Teacher{
		Password: "123",
		Email:    "123",
		Name:     "汪溢",
	}
	teacher2 := etity.Teacher{
		Password: "123",
		Email:    "123",
		Name:     "汪函",
	}

	teacherlist = append(teacherlist, teacher1)
	teacherlist = append(teacherlist, teacher2)
	teacherDao.AddTeacher(teacherlist)
}

func TestLike(t *testing.T) {
	teachers := teacherDao.FindTeacherByName("汪")
	for _, teacher := range teachers {
		fmt.Println(teacher)
	}
}

func TestEmail(t *testing.T) {
	teacher, err := teacherDao.FindTeacherByEmail("nil")
	if err != nil {
		fmt.Println("查询失败")
	} else {
		fmt.Println(*teacher)
	}
}

func TestId(t *testing.T) {
	teacher := teacherDao.FindTeacherNameById(36)
	fmt.Println(teacher)

}
