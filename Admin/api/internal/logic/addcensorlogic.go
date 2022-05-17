package logic

import (
	"context"
	"onlineEudcation/Admin/etity"

	"onlineEudcation/Admin/api/internal/svc"
	"onlineEudcation/Admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCensorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCensorLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddCensorLogic {
	return AddCensorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCensorLogic) AddCensor(req types.AddCensorReq) (resp *types.AddCensorRes, err error) {
	adminDao.AddCensorship(etity.Censorship{
		CourseId:  req.CourseId,
		TeacherId: req.TeacherId,
		Result:    req.CensorStatus,
		Message:   req.Message,
	})

	return
}
