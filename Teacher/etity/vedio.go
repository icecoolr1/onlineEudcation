package etity

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	//课程id
	CourseId int32 `gorm:"column:c_id"`
	//视频文件系统url
	VideoFileUrl string `gorm:"column:video_file_url"`
	//视频播放量
	VideoPlaySum string `gorm:"column:video_play_sum"`
	//视频名称
	VideoName string `gorm:"column:video_name"`
	//视频图片文件系统位置
	VideoImagePath string `gorm:"column:video_imag_path"`
	//视频顺序
	VideoSeq int32 `gorm:"column:video_seq"`
}
