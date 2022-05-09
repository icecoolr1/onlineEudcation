package logic

import (
	"context"
	"onlineEudcation/Classes/rpc/rpc"

	"onlineEudcation/Classes/api/internal/svc"
	"onlineEudcation/Classes/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteClassLogic {
	return DeleteClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteClassLogic) DeleteClass(req types.DeleteClassRequest) (resp *types.DeleteClassResponse, err error) {
	_, err = l.svcCtx.ClassesRpc.DeleteClass(l.ctx, &rpc.AddClassRequest{StudentId: int32(req.StudentId), CourseId: int32(req.CourseId)})
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
	if err != nil {
		return nil, err
	} else {
		return &types.DeleteClassResponse{
			Code:    200,
			Courses: classesList,
		}, nil
	}

}
