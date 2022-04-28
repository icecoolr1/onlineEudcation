package logic

import (
	"context"

	"onlineEudcation/Student/api/internal/svc"
	"onlineEudcation/Student/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudentLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) StudentLoginLogic {
	return StudentLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudentLoginLogic) StudentLogin(req types.StudentLoginReq) (resp *types.StudentLoginRes, err error) {
	// todo: add your logic here and delete this line

	return
}
