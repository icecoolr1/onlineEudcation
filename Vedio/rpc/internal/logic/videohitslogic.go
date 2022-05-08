package logic

import (
	"context"

	"onlineEudcation/Vedio/rpc/internal/svc"
	"onlineEudcation/Vedio/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoHitsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVideoHitsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VideoHitsLogic {
	return &VideoHitsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *VideoHitsLogic) VideoHits(in *rpc.VideoHitsReq) (*rpc.AddVideoResponse, error) {
	videoDao.VideoHits(int(in.VideoId))

	return &rpc.AddVideoResponse{
		Message: "ok",
		Code:    200,
	}, nil
}
