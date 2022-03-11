package etity

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	//课程名
	CourseName string `gorm:"column:course_name"`
	//所属教师ID
	TeacherId int32 `gorm:"column:t_id"`
	//课程类型
	CourseTag string `gorm:"column:course_tag"`
	//课程总播放量
	CoursePlaySum int32 `gorm:"column:course_play_sum"`
	//课程视频数
	CourseVideoNumBer int `gorm:"column:course_video_number"`
	//视频审核状态
	CourseStatus bool `gorm:"column:course_status"`
}
