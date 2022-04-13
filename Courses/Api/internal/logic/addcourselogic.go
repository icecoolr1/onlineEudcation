package logic

import (
	"context"
	"onlineEudcation/Courses/Rpc/Rpc"

	"onlineEudcation/Courses/Api/internal/svc"
	"onlineEudcation/Courses/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddCourseLogic {
	return AddCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCourseLogic) AddCourse(req types.CourseAddrequest) (resp *types.Response, err error) {
	course, err := l.svcCtx.CourseRpc.AddCourse(l.ctx, &Rpc.CourseAddReq{
		CourseName: req.CourseName,
		TeacherId:  req.TeacherId,
		ImgUrl:     req.ImgUrl,
		CourseType: req.CourseTag,
	})
	if course.Code == 200 {
		return &types.Response{
			Code:    200,
			Message: "success",
			Res:     true,
		}, nil
	} else {
		return &types.Response{
			Code:    500,
			Message: "fail",
			Res:     false,
		}, nil
	}
}
