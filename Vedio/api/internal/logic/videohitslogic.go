package logic

import (
	"context"
	"onlineEudcation/Vedio/rpc/rpc"

	"onlineEudcation/Vedio/api/internal/svc"
	"onlineEudcation/Vedio/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VideoHitsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVideoHitsLogic(ctx context.Context, svcCtx *svc.ServiceContext) VideoHitsLogic {
	return VideoHitsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VideoHitsLogic) VideoHits(req types.VideoHitsReq) (resp *types.AddVideoRes, err error) {
	_, err = l.svcCtx.VideoRpc.VideoHits(l.ctx, &rpc.VideoHitsReq{VideoId: int64(req.VideoId)})
	if err != nil {
		return nil, err
	} else {
		return &types.AddVideoRes{
			Code:    200,
			Message: "ok",
			Res:     true,
		}, nil
	}
}
