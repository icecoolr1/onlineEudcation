package logic

import (
	"context"
	"onlineEudcation/Teacher/Api/internal/svc"
	"onlineEudcation/Teacher/Api/internal/types"
	"onlineEudcation/Teacher/Rpc/Rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) TLoginLogic {
	return TLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// TLogin 教师登录实现
func (l *TLoginLogic) TLogin(req types.TeacherInfoForLogin) (resp *types.UserInfo, err error) {
	login, err := l.svcCtx.TeacherRpc.Login(l.ctx, &Rpc.TeacherLoginInfo{TeacherEmail: req.Email, TeacherPassword: req.Password})
	if err != nil {
		return &types.UserInfo{
			TeacherRes:  types.TeacherResponse{Code: 400, Res: false, Message: err.Error()},
			Id:          0,
			TeacherInfo: types.TeacherInfoForRegister{},
		}, nil
	}
	//成功
	if login.Code == 200 {
		return &types.UserInfo{
			TeacherRes: types.TeacherResponse{Code: 200, Res: true, Message: "OK"},
			Id:         login.Teacher.Id,
			TeacherInfo: types.TeacherInfoForRegister{
				Username: login.Teacher.Name,
				Password: login.Teacher.Password,
				Email:    login.Teacher.Email,
			},
		}, nil
	} else {
		return &types.UserInfo{
			TeacherRes:  types.TeacherResponse{Code: 400, Res: false, Message: login.Message},
			Id:          0,
			TeacherInfo: types.TeacherInfoForRegister{},
		}, nil

	}

}
