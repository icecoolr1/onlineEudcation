package logic

import (
	"context"
	"gorm.io/gorm"
	"onlineEudcation/Courses/Etity"

	"onlineEudcation/Courses/Rpc/Rpc"
	"onlineEudcation/Courses/Rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCourseLogic {
	return &UpdateCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCourseLogic) UpdateCourse(in *Rpc.UpdateCourseReq) (*Rpc.CourseResp, error) {

	CourseDao.UpdateCourse(Etity.Course{
		Model:      gorm.Model{ID: uint(in.CourseId)},
		CourseName: in.CourseName,
		CourseTag:  in.CourseType,
		CourseImg:  in.ImgUrl,
	})
	return &Rpc.CourseResp{
		Msg:  "更新成功",
		Code: 200,
		Res:  true,
	}, nil
}
