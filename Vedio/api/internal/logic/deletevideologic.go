package logic

import (
	"context"
	"onlineEudcation/Vedio/rpc/rpc"

	"onlineEudcation/Vedio/api/internal/svc"
	"onlineEudcation/Vedio/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteVideoLogic {
	return DeleteVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteVideoLogic) DeleteVideo(req types.DeleteVideoReq) (resp *types.FindCourseVideosRes, err error) {
	_, err = l.svcCtx.VideoRpc.DelVideo(l.ctx, &rpc.DelVideoReq{VideoId: int64(req.VideoId)})
	if err != nil {
		return nil, err
	} else {
		videoList := make([]types.Video, 0)
		videos := videoDao.FindVideoByCourseId(int(req.CourseId))
		for _, video := range videos {
			videoList = append(videoList, types.Video{
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
