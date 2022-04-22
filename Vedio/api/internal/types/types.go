// Code generated by goctl. DO NOT EDIT.
package types

type AddVideoReq struct {
	Video
}

type AddVideoRes struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Res     bool   `json:"res"`
}

type Video struct {
	CourseId     int32  `json:"courseId"`
	VideoFileUrl string `json:"videoFileUrl"`
	VideoPlaySum int32  `json:"videoPlaySum"`
	VideoName    string `json:"videoName"`
	VideoSeq     int32  `json:"videoSeq"`
	VedioId      int32  `json:"vedioId"`
}

type FindCourseVideosReq struct {
	CourseId int32 `json:"courseId"`
}

type FindCourseVideosRes struct {
	Code   int32   `json:"code"`
	Videos []Video `json:"videos"`
}

type DeleteVideoReq struct {
	VideoId int32 `json:"videoId"`
}