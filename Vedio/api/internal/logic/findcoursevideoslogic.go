package logic

import (
	"context"
	"onlineEudcation/Vedio/Dao"

	"onlineEudcation/Vedio/api/internal/svc"
	"onlineEudcation/Vedio/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var videoDao Dao.VideoDaoInterface = new(Dao.VideoDao)

type FindCourseVideosLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindCourseVideosLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindCourseVideosLogic {
	return FindCourseVideosLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCourseVideosLogic) FindCourseVideos(req types.FindCourseVideosReq) (resp *types.FindCourseVideosRes, err error) {
	videos := videoDao.FindVideoByCourseId(int(req.CourseId))
	videoList := make([]types.Video, 0)
	if len(videos) == 0 {
		return &types.FindCourseVideosRes{
			Code:   201,
			Videos: nil,
		}, nil
	} else {
		for _, video := range videos {
			videoList = append(videoList, types.Video{
				CourseId:     video.CourseId,
				VideoFileUrl: video.VideoFileUrl,
				VideoPlaySum: video.VideoPlaySum,
				VideoName:    video.VideoName,
				VideoSeq:     video.VideoSeq,
				VedioId:      int32(video.Model.ID),
			})
		}
		return &types.FindCourseVideosRes{
			Code:   200,
			Videos: videoList,
		}, nil
	}

}
