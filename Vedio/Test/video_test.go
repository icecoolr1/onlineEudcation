package Test

import (
	"fmt"
	"gorm.io/gorm"
	"onlineEudcation/Vedio/Dao"
	"onlineEudcation/Vedio/Etity"
	"testing"
)

var videoDao Dao.VideoDaoInterface = new(Dao.VideoDao)

func TestAdd(t *testing.T) {
	videoList := make([]Etity.Video, 0)
	video := Etity.Video{
		CourseId:     1,
		VideoFileUrl: "test",
		VideoPlaySum: 0,
		VideoName:    "test",
		VideoSeq:     1,
	}
	video1 := Etity.Video{
		CourseId:     1,
		VideoFileUrl: "test2",
		VideoPlaySum: 0,
		VideoName:    "test2",
		VideoSeq:     1,
	}
	videoList = append(videoList, video)
	videoList = append(videoList, video1)
	videoDao.AddVideo(videoList)
}

func TestDelete(t *testing.T) {
	videoDao.DelVideo([]int{1})

}

func TestFind(t *testing.T) {
	video := videoDao.FindAllVideo()
	for _, e := range video {
		fmt.Println(e)
	}
}

func TestLike(t *testing.T) {
	videos := videoDao.FindVideo("test")
	for _, video := range videos {
		fmt.Println(video)
	}
}

func TestUpdate(t *testing.T) {
	video := Etity.Video{
		Model:        gorm.Model{ID: 2},
		CourseId:     1,
		VideoFileUrl: "update",
		VideoPlaySum: 0,
		VideoName:    "update",
		VideoSeq:     1,
	}
	videoDao.UpdateVideo(video)
}

func TestFindByCourseId(t *testing.T) {
	videos := videoDao.FindVideoByCourseId(10)
	for _, video := range videos {
		fmt.Println(video)
	}
}
