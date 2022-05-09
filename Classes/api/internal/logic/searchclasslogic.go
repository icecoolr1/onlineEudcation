package logic

import (
	"context"

	"onlineEudcation/Classes/api/internal/svc"
	"onlineEudcation/Classes/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchClassLogic {
	return SearchClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchClassLogic) SearchClass(req types.SearchClassRequest) (resp *types.SearchClassResponse, err error) {
	//先查询选课表获得选课信息
	classes := classDao.FindAllClasses(req.StudentId)
	classesList := make([]types.Course, 0)
	//通过选课表获取课程id 查询课程信息
	for _, class := range classes {
		course := courseDao.FindCourseByCourseId(int(class.CourseId))
		tname := teacherDao.FindTeacherNameById(int(course.TeacherId))
		classesList = append(classesList, types.Course{
			CourseId:      int32(course.ID),
			CourseName:    course.CourseName,
			CourseImg:     course.CourseImg,
			CourseType:    course.CourseTag,
			CoursePlaySum: course.CoursePlaySum,
			TeacherName:   tname,
		})
	}
	return &types.SearchClassResponse{
		Code:    200,
		Courses: classesList,
	}, nil
}
