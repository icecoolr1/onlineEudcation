package test

import (
	"fmt"
	"gorm.io/gorm"
	"onlineEudcation/Student/dao"
	"onlineEudcation/Student/etity"
	"onlineEudcation/Tools/scripts"
	"testing"
)

var studentDao dao.StudentDaoInterface = new(dao.StudentDao)
var redis = scripts.GetRedis(1)
var redis2 = scripts.GetRedis(2)

func TestAdd(t *testing.T) {
	student := etity.Student{
		StudentPassword: "123",
		StudentName:     "我是猪",
		StudentEmail:    "dqdqad",
	}
	studentDao.InsertStudent(student)
}

func TestFindAll(t *testing.T) {
	students := studentDao.SelectAll()
	for _, student := range students {
		fmt.Println(student)
	}
}

func TestFindByID(t *testing.T) {
	student := studentDao.SelectByID(1)
	fmt.Println(student)
}

func TestUpdate(t *testing.T) {
	student := etity.Student{
		Model:        gorm.Model{ID: 1},
		StudentEmail: "testUpdate",
	}
	studentDao.UpdateStudent(student)
}

func TestDelete(t *testing.T) {
	studentDao.DeleteStudent(2)
}

func TestFindByName(t *testing.T) {
	student, _ := studentDao.FindStudentByName("我是猪")
	fmt.Println(student)
}

func TestRedis(t *testing.T) {
	exists := redis.Exists("哈哈")
	result, _ := exists.Result()
	fmt.Println(result)
	//redis.Set("哈哈", true, 0)

}

func TestList(t *testing.T) {
	//redis2.LPush("test", "3")
	lLen := redis2.LLen("test")

	fmt.Println(lLen.Val())

	lRange := redis2.LRange("test", 0, lLen.Val()).Val()
	fmt.Println(lRange)
}
