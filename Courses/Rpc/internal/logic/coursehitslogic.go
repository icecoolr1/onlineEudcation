package logic

import (
	"context"

	"onlineEudcation/Courses/Rpc/Rpc"
	"onlineEudcation/Courses/Rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CourseHitsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCourseHitsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseHitsLogic {
	return &CourseHitsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CourseHitsLogic) CourseHits(in *Rpc.CourseHitsReq) (*Rpc.CourseResp, error) {
	CourseDao.CourseHits(int(in.CourseId))

	return &Rpc.CourseResp{
		Msg:  "ok",
		Code: 200,
		Res:  true,
	}, nil
}
