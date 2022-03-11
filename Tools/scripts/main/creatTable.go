package main

import (
	"onlineEudcation/Teacher/etity"
	"onlineEudcation/Tools/scripts"
)

func main() {
	conn := scripts.GetDatabaseConnection()

	conn.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(new(etity.Admin))
	//conn.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(new(etity.Course))
	//conn.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(new(etity.Video))
	//course := etity.Course{
	//	CourseName:        "哈哈",
	//	TeacherId:         1,
	//	CourseTag:         "java",
	//	CoursePlaySum:     0,
	//	CourseVideoNumBer: 0,
	//	CourseStatus:      false,
	//}
	//conn.Create(&course)

}
