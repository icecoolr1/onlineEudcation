package logic

import (
	"context"

	"onlineEudcation/Admin/api/internal/svc"
	"onlineEudcation/Admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchNeededCoursesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchNeededCoursesLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchNeededCoursesLogic {
	return SearchNeededCoursesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchNeededCoursesLogic) SearchNeededCourses(req types.SearchCoursesReq) (resp *types.SearchCoursesRes, err error) {
	resList := make([]types.Course, 0)
	courses := adminDao.GetAllCensorCourses()

	for _, course := range courses {
		resList = append(resList, types.Course{
			CourseId:          int32(course.ID),
			CourseName:        course.CourseName,
			CourseImg:         course.CourseImg,
			CourseType:        course.CourseTag,
			CourseStatus:      course.CourseStatus,
			CoursePlaySum:     course.CoursePlaySum,
			CourseVideoNumBer: int32(course.CourseVideoNumBer),
			TeacherId:         int(course.TeacherId),
		})

	}
	return &types.SearchCoursesRes{
		Courses: resList,
	}, nil
}
