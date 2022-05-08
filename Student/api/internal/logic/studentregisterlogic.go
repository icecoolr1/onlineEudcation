package logic

import (
	"context"
	"onlineEudcation/Student/rpc/rpc"

	"onlineEudcation/Student/api/internal/svc"
	"onlineEudcation/Student/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudentRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) StudentRegisterLogic {
	return StudentRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudentRegisterLogic) StudentRegister(req types.StudentRegisterReq) (resp *types.StudentRegisterRes, err error) {
	// 用户名不可重复
	_, err = studentDao.FindStudentByName(req.StudentName)
	if err != nil {
		_, err = l.svcCtx.StudentRpc.AddStudent(l.ctx, &rpc.StudentAddReq{
			Student: &rpc.Student{
				SName:     req.StudentName,
				SEmail:    req.StudentEmail,
				SPassword: req.StudentPassword,
			},
		})
		if err != nil {
			return nil, err
		}
		return &types.StudentRegisterRes{
			Code:    200,
			Message: "注册成功",
		}, err
	} else {
		return &types.StudentRegisterRes{
			Code:    201,
			Message: "用户名已存在",
		}, err
	}

}
