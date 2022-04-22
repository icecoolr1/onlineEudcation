package logic

import (
	"context"
	"gorm.io/gorm"
	"onlineEudcation/Vedio/Etity"

	"onlineEudcation/Vedio/rpc/internal/svc"
	"onlineEudcation/Vedio/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateVideoLogic {
	return &UpdateVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateVideoLogic) UpdateVideo(in *rpc.UpdateVideoRequest) (*rpc.AddVideoResponse, error) {
	videoDao.UpdateVideo(Etity.Video{
		Model:        gorm.Model{ID: uint(in.VideoId)},
		VideoFileUrl: in.VideoFileUrl,
		VideoName:    in.VideoName,
		VideoSeq:     in.VideoSeq,
	})

	return &rpc.AddVideoResponse{
		Message: "update success",
		Code:    200,
	}, nil
}
