package logic

import (
	"context"
	"onlineEudcation/Teacher/etity"

	"onlineEudcation/Teacher/Rpc/Rpc"
	"onlineEudcation/Teacher/Rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
	"onlineEudcation/Tools/scripts"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Register 教师注册功能实现
func (l *RegisterLogic) Register(in *Rpc.TeacherInfo) (*Rpc.Response, error) {
	var mysqlDB = scripts.GetDatabaseConnection()
	result := mysqlDB.Create(&etity.Teacher{
		Password: in.Password,
		Email:    in.TeacherEmail,
		Name:     in.TeacherName,
	})
	if result.Error != nil {
		return &Rpc.Response{
			Code:    400,
			Res:     false,
			Message: result.Error.Error(),
		}, nil
	} else {
		return &Rpc.Response{
			Code:    200,
			Res:     true,
			Message: "OK",
		}, nil
	}
}
