package logic

import (
	"context"
	"onlineEudcation/Courses/Dao"
	"onlineEudcation/Courses/Etity"

	"onlineEudcation/Courses/Rpc/Rpc"
	"onlineEudcation/Courses/Rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

var CourseDao Dao.CourseDaoInterface = new(Dao.CourseDao)

type AddCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCourseLogic {
	return &AddCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCourseLogic) AddCourse(in *Rpc.CourseAddReq) (*Rpc.CourseResp, error) {
	var courseList []Etity.Course
	courseList = append(courseList, Etity.Course{
		CourseName:        in.CourseName,
		TeacherId:         in.TeacherId,
		CourseTag:         in.CourseType,
		CoursePlaySum:     0,
		CourseVideoNumBer: 0,
		CourseStatus:      false,
		CourseImg:         in.ImgUrl,
	})
	CourseDao.AddCourse(courseList)

	return &Rpc.CourseResp{
		Msg:  "成功",
		Code: 200,
		Res:  true,
	}, nil
}
