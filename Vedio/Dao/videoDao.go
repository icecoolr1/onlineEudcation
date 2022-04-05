package Dao

import (
	"fmt"
	"onlineEudcation/Tools/scripts"
	"onlineEudcation/Vedio/Etity"
)

type VideoDao struct{}

var videoList []Etity.Video
var DB = scripts.GetDatabaseConnection()

func (VideoDao) AddVideo(videos []Etity.Video) bool {
	for _, e := range videos {
		fmt.Printf("%v\n", e)
		DB.Create(&e)
	}
	return true
}

func (VideoDao) DelVideo(videoId []int) bool {
	DB.Where("id in (?)", videoId).Delete(&Etity.Video{})
	return true
}

func (VideoDao) UpdateVideo(video Etity.Video) bool {
	DB.Model(&Etity.Video{}).Where("id = ? ", video.Model.ID).Updates(video)
	return true
}

func (VideoDao) FindVideo(videoName string) []Etity.Video {
	DB.Where("video_name LIKE ?", "%"+videoName+"%").Find(&videoList)
	return videoList
}

func (VideoDao) FindAllVideo() []Etity.Video {
	DB.Find(&videoList)
	return videoList
}
