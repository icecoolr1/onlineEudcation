package logic

import (
	"context"
	"onlineEudcation/Vedio/rpc/rpc"

	"onlineEudcation/Vedio/api/internal/svc"
	"onlineEudcation/Vedio/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateVideoLogic {
	return UpdateVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateVideoLogic) UpdateVideo(req types.AddVideoReq) (resp *types.AddVideoRes, err error) {
	_, err = l.svcCtx.VideoRpc.UpdateVideo(l.ctx, &rpc.UpdateVideoRequest{
		VideoId:      int64(req.VedioId),
		VideoFileUrl: req.VideoFileUrl,
		VideoName:    req.VideoName,
		VideoSeq:     req.VideoSeq,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddVideoRes{
		Code:    0,
		Message: "success",
		Res:     true,
	}, nil
}
