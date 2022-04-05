package main

import (
	"onlineEudcation/Tools/scripts"
	"onlineEudcation/Vedio/Etity"
)

func main() {
	conn := scripts.GetDatabaseConnection()

	err := conn.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(new(Etity.Video))
	if err != nil {
		panic(err)
	}
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
