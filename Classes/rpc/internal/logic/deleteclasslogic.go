package logic

import (
	"context"

	"onlineEudcation/Classes/rpc/internal/svc"
	"onlineEudcation/Classes/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteClassLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteClassLogic {
	return &DeleteClassLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteClassLogic) DeleteClass(in *rpc.AddClassRequest) (*rpc.Request, error) {
	classesDao.DelClass(int(in.StudentId), int(in.CourseId))
	return &rpc.Request{
		Code:    200,
		Message: "ok",
	}, nil
}
