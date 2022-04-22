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

func (l *DeleteVideoLogic) DeleteVideo(req types.DeleteVideoReq) (resp *types.AddVideoRes, err error) {
	_, err = l.svcCtx.VideoRpc.DelVideo(l.ctx, &rpc.DelVideoReq{VideoId: int64(req.VideoId)})
	if err != nil {
		return nil, err
	} else {
		return &types.AddVideoRes{
			Code:    200,
			Message: "success",
			Res:     true,
		}, nil
	}

}
