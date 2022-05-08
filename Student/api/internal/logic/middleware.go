package logic

import (
	"onlineEudcation/Student/dao"
	"onlineEudcation/Tools/scripts"
)

var studentDao dao.StudentDaoInterface = new(dao.StudentDao)
var redis = scripts.GetRedis(1)
var redis2 = scripts.GetRedis(2)
