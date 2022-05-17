package logic

import (
	"context"

	"onlineEudcation/Admin/api/internal/svc"
	"onlineEudcation/Admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeCourseStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeCourseStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChangeCourseStatusLogic {
	return ChangeCourseStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeCourseStatusLogic) ChangeCourseStatus(req types.ChangeStatusReq) (resp *types.ChangeStatusRes, err error) {
	adminDao.ChangeStatus(int(req.CourseId))

	return &types.ChangeStatusRes{
		Code: 200,
	}, nil
}
