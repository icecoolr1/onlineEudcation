package main

import (
	"onlineEudcation/Teacher/etity"
	"onlineEudcation/Tools/scripts"
)

func main() {
	conn := scripts.GetDatabaseConnection()
	conn.Create(&etity.Teacher{
		Password: "test",
		Email:    "test",
		Name:     "汪溢",
	})
}
