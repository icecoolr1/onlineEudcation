package logic

import (
	"context"
	"onlineEudcation/Student/etity"

	"onlineEudcation/Student/rpc/internal/svc"
	"onlineEudcation/Student/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStudentLogic {
	return &AddStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddStudentLogic) AddStudent(in *rpc.StudentAddReq) (*rpc.StudentAddRes, error) {
	studentDao.InsertStudent(etity.Student{
		StudentPassword: in.Student.SPassword,
		StudentName:     in.Student.SName,
		StudentEmail:    in.Student.SEmail,
	})
	return &rpc.StudentAddRes{
		Code: 200,
	}, nil
}
