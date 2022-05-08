package logic

import (
	"context"
	"onlineEudcation/Courses/Rpc/Rpc"

	"onlineEudcation/Courses/Api/internal/svc"
	"onlineEudcation/Courses/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CourseHitsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCourseHitsLogic(ctx context.Context, svcCtx *svc.ServiceContext) CourseHitsLogic {
	return CourseHitsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CourseHitsLogic) CourseHits(req types.CourseHitsReq) (resp *types.Response, err error) {
	_, err = l.svcCtx.CourseRpc.CourseHits(l.ctx, &Rpc.CourseHitsReq{CourseId: int32(req.CourseId)})
	if err != nil {
		return nil, err
	} else {
		return &types.Response{
			Code:    200,
			Message: "ok",
			Res:     true,
		}, nil
	}

}
