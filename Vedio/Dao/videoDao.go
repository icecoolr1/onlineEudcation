package Dao

import (
	"fmt"
	"gorm.io/gorm"
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

func (d VideoDao) FindVideoByCourseId(courseId int) []Etity.Video {
	DB.Where("c_id = ?", courseId).Order("video_seq asc").Find(&videoList)
	return videoList
}

func (d VideoDao) FindVideoByVideoId(videoId int) Etity.Video {
	video := Etity.Video{}
	DB.Where("id = ?", videoId).First(&video)
	return video
}

func (d VideoDao) VideoHits(videoId int) bool {
	video := Etity.Video{Model: gorm.Model{ID: uint(videoId)}}
	DB.Model(&video).Update("video_play_sum", gorm.Expr("video_play_sum + ?", 1))
	return true
}
