package logic

import (
	"context"
	dao "onlineEudcation/Teacher/Dao"

	"onlineEudcation/Courses/Api/internal/svc"
	"onlineEudcation/Courses/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var teacherDao dao.TeacherInterface = new(dao.TeacherDao)

func NewGetCourseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCourseListLogic {
	return GetCourseListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCourseListLogic) GetCourseList(req types.GetCourseListReq) (resp *types.GetCourseListRes, err error) {
	list := courseDao.GetCourseList(req.Tag)
	courseList := make([]types.SearchCourseByIdItemByStudent, 0)
	if len(list) == 0 {
		return &types.GetCourseListRes{
			Code:       201,
			CourseList: nil,
		}, nil
	} else {
		for _, course := range list {
			tName := teacherDao.FindTeacherNameById(int(course.TeacherId))
			courseList = append(courseList, types.SearchCourseByIdItemByStudent{
				CourseId:      int32(course.ID),
				CourseName:    course.CourseName,
				CourseImg:     course.CourseImg,
				CourseType:    course.CourseTag,
				CourseStatus:  course.CourseStatus,
				TeacherName:   tName,
				CoursePlaySum: course.CoursePlaySum,
			})
		}

		return &types.GetCourseListRes{
			Code:       200,
			CourseList: courseList,
		}, nil
	}

}
