package logic

import (
	"context"
	"onlineEudcation/Student/api/internal/svc"
	"onlineEudcation/Student/api/internal/types"
	"strconv"

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
	//获取用户名查询数据库
	student, err := studentDao.FindStudentByName(req.StudentName)
	if err != nil {
		l.Errorf("student login find student by name error: %s", err)
		return &types.StudentLoginRes{
			Code:    202,
			IsFirst: false,
			Message: "未查询到用户",
		}, nil
	} else {
		if student.StudentPassword == req.StudentPassword {
			//通过redis查询是否有登录记录
			exists := redis.Exists(student.StudentName)
			result, _ := exists.Result()
			if result == 0 {
				redis.Set(req.StudentName, true, 0)
				return &types.StudentLoginRes{
					Code:    200,
					IsFirst: true,
					Message: strconv.Itoa(int(student.ID)),
				}, nil
			} else {
				return &types.StudentLoginRes{
					Code:    200,
					IsFirst: false,
					Message: strconv.Itoa(int(student.ID)),
				}, nil
			}
		} else {
			return &types.StudentLoginRes{
				Code:    203,
				IsFirst: false,
				Message: "密码错误",
			}, nil
		}
	}
}
