package logic

import (
	"context"
	"onlineEudcation/Classes/rpc/rpc"

	"onlineEudcation/Classes/api/internal/svc"
	"onlineEudcation/Classes/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddClassLogic {
	return AddClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddClassLogic) AddClass(req types.AddClassRequest) (resp *types.Response, err error) {
	res, err := l.svcCtx.ClassesRpc.AddClass(l.ctx, &rpc.AddClassRequest{
		StudentId: int32(req.StudentId),
		CourseId:  int32(req.CourseId),
	})

	if err != nil {
		return nil, err
	} else {

		if res.GetCode() == 300 {
			return &types.Response{
				Code:    300,
				Message: "您已收藏该课程",
			}, nil
		} else {
			return &types.Response{
				Code:    200,
				Message: "ok",
			}, nil
		}
	}

}
