package logic

import (
	"context"
	"fmt"
	"onlineEudcation/Teacher/Rpc/teacherrpc"

	"onlineEudcation/Teacher/Api/internal/svc"
	"onlineEudcation/Teacher/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) TRegisterLogic {
	return TRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TRegisterLogic) TRegister(req types.TeacherInfoForRegister) (resp *types.TeacherResponse, err error) {
	register, err := l.svcCtx.TeacherRpc.Register(l.ctx, &teacherrpc.TeacherInfo{
		TeacherName:  req.Username,
		Password:     req.Password,
		TeacherEmail: req.Email,
	})
	fmt.Println(register.GetRes())
	if register.GetRes() == true {
		return &types.TeacherResponse{
			Code:    200,
			Res:     true,
			Message: "OK",
		}, err
	} else {
		return &types.TeacherResponse{
			Code:    400,
			Res:     false,
			Message: "邮箱已被注册",
		}, err
	}
}
