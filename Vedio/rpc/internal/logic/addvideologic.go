package logic

import (
	"context"
	"onlineEudcation/Vedio/Dao"
	"onlineEudcation/Vedio/Etity"

	"onlineEudcation/Vedio/rpc/internal/svc"
	"onlineEudcation/Vedio/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var videoDao Dao.VideoDaoInterface = new(Dao.VideoDao)

func NewAddVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVideoLogic {
	return &AddVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddVideoLogic) AddVideo(in *rpc.AddVideoRequest) (*rpc.AddVideoResponse, error) {
	videoList := make([]Etity.Video, 0)
	videoList = append(videoList, Etity.Video{
		CourseId:     in.CourseId,
		VideoFileUrl: in.VideoFileUrl,
		VideoPlaySum: 0,
		VideoName:    in.VideoName,
		VideoSeq:     in.VideoSeq,
	})
	videoDao.AddVideo(videoList)
	return &rpc.AddVideoResponse{
		Message: "添加成功",
		Code:    200,
	}, nil
}
