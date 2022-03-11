package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"onlineEudcation/Teacher/Rpc/Rpc"
	"onlineEudcation/Teacher/Rpc/internal/svc"
	"onlineEudcation/Teacher/etity"
	"onlineEudcation/Tools/scripts"
	"strconv"
	"strings"
	"time"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Login 教师登录功能实现 寻找邮箱是否存在，接着判断密码是否正确
func (l *LoginLogic) Login(in *Rpc.TeacherLoginInfo) (*Rpc.LoginResponse, error) {
	var mysqlDB = scripts.GetDatabaseConnection()
	// 寻找邮箱是否存在
	var teacher etity.Teacher
	var rMsg *Rpc.Teacher
	result := mysqlDB.Where("t_email = ? ", in.TeacherEmail).First(&teacher)
	//返回不为空
	if result.Error != nil {
		rMsg = nil
		return &Rpc.LoginResponse{Code: 400, Res: false, Message: "邮箱不存在", Teacher: rMsg}, nil
	} else {
		//正确
		teacherinfo := Rpc.Teacher{Id: int32(teacher.ID), Password: teacher.Password, Email: teacher.Email, Name: teacher.Name}
		if strings.Compare(teacher.Password, in.TeacherPassword) == 0 {
			redis := scripts.GetRedis(0)
			key := strconv.Itoa(int(teacher.ID))
			name := teacher.Name
			err := redis.Set(key, name, 3*time.Hour).Err()
			if err != nil {
				panic(err)
			}
			return &Rpc.LoginResponse{Res: true, Code: 200, Message: "OK", Teacher: &teacherinfo}, nil
		} else {
			rMsg = nil
			return &Rpc.LoginResponse{Code: 400, Res: false, Message: "密码错误", Teacher: rMsg}, nil
		}
	}
}
