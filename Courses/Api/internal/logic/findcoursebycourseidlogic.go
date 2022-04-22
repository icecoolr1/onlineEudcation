package logic

import (
	"context"

	"onlineEudcation/Courses/Api/internal/svc"
	"onlineEudcation/Courses/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindCourseByCourseIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindCourseByCourseIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindCourseByCourseIdLogic {
	return FindCourseByCourseIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindCourseByCourseIdLogic) FindCourseByCourseId(req types.FindCourseByCourseId) (resp *types.SearchCourseByIdItem, err error) {
	course := courseDao.FindCourseByCourseId(req.CourseId)

	return &types.SearchCourseByIdItem{
		CourseId:          int32(course.ID),
		CourseName:        course.CourseName,
		CourseImg:         course.CourseImg,
		CourseType:        course.CourseTag,
		CourseStatus:      course.CourseStatus,
		TeacherId:         course.TeacherId,
		CoursePlaySum:     course.CoursePlaySum,
		CourseVideoNumBer: int32(course.CourseVideoNumBer),
	}, nil
}
