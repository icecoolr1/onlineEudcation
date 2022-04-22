package logic

import (
	"context"
	"onlineEudcation/Vedio/rpc/rpc"

	"onlineEudcation/Vedio/api/internal/svc"
	"onlineEudcation/Vedio/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddVideoLogic {
	return AddVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddVideoLogic) AddVideo(req types.AddVideoReq) (resp *types.AddVideoRes, err error) {
	_, err = l.svcCtx.VideoRpc.AddVideo(l.ctx, &rpc.AddVideoRequest{
		CourseId:     req.CourseId,
		VideoFileUrl: req.VideoFileUrl,
		VideoName:    req.VideoName,
		VideoSeq:     req.VideoSeq,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddVideoRes{
		Code:    200,
		Message: "success",
		Res:     true,
	}, nil
}
