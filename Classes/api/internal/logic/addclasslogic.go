package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return
}
