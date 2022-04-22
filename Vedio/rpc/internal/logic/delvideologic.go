package logic

import (
	"context"

	"onlineEudcation/Vedio/rpc/internal/svc"
	"onlineEudcation/Vedio/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelVideoLogic {
	return &DelVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelVideoLogic) DelVideo(in *rpc.DelVideoReq) (*rpc.AddVideoResponse, error) {
	videoDao.DelVideo([]int{int(in.VideoId)})

	return &rpc.AddVideoResponse{
		Message: "删除成功",
		Code:    200,
	}, nil
}
