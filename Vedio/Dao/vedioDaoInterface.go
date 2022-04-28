package Dao

import "onlineEudcation/Vedio/Etity"

type VideoDaoInterface interface {
	// AddVideo 添加视频
	AddVideo(videos []Etity.Video) bool
	// DelVideo 删除视频
	DelVideo(videoId []int) bool
	// UpdateVideo 更新视频
	UpdateVideo(video Etity.Video) bool
	// FindVideo 查找视频 (模糊查询)
	FindVideo(videoName string) []Etity.Video
	// FindAllVideo 查询所有视频
	FindAllVideo() []Etity.Video

	FindVideoByCourseId(courseId int) []Etity.Video
	FindVideoByVideoId(videoId int) Etity.Video
}
