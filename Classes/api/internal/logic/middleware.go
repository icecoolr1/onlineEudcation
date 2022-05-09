package logic

import (
	"onlineEudcation/Classes/Dao"
	Dao2 "onlineEudcation/Courses/Dao"
	dao "onlineEudcation/Teacher/Dao"
)

var classDao Dao.ClassesDaoInterface = new(Dao.ClassesDao)
var courseDao Dao2.CourseDaoInterface = new(Dao2.CourseDao)
var teacherDao dao.TeacherInterface = new(dao.TeacherDao)
