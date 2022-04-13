package logic

import (
	"context"
	dao "onlineEudcation/Teacher/Dao"
	"onlineEudcation/Teacher/etity"

	"onlineEudcation/Teacher/Rpc/Rpc"
	"onlineEudcation/Teacher/Rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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

var TeacherDao dao.TeacherInterface = new(dao.TeacherDao)
var teacherList []etity.Teacher = make([]etity.Teacher, 0)

// Register 教师注册功能实现
func (l *RegisterLogic) Register(in *Rpc.TeacherInfo) (*Rpc.Response, error) {
	//判断邮箱是否已经注册
	res, err := TeacherDao.FindTeacherByEmail(in.TeacherEmail)
	if res != nil {
		return &Rpc.Response{
			Code:    400,
			Message: "邮箱已被注册,请登录",
			Res:     false,
		}, err
	} else {
		// 添加教师逻辑
		teacher := etity.Teacher{
			Password: in.Password,
			Email:    in.TeacherEmail,
			Name:     in.TeacherName,
		}
		teacherList = append(teacherList, teacher)
		_, err := TeacherDao.AddTeacher(teacherList)
		if err != nil {
			return &Rpc.Response{
				Code:    400,
				Res:     false,
				Message: err.Error(),
			}, err
		} else {
			return &Rpc.Response{
				Code:    200,
				Res:     true,
				Message: "OK",
			}, nil
		}
	}

}
