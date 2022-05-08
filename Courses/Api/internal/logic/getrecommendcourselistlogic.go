package logic

import (
	"context"
	"onlineEudcation/Student/dao"
	"onlineEudcation/Tools/scripts"

	"onlineEudcation/Courses/Api/internal/svc"
	"onlineEudcation/Courses/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var redis2 = scripts.GetRedis(2)
var studentDao dao.StudentDaoInterface = new(dao.StudentDao)

type GetRecommendCourseListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRecommendCourseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRecommendCourseListLogic {
	return GetRecommendCourseListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRecommendCourseListLogic) GetRecommendCourseList(req types.GetRecommendCourseListReq) (resp *types.GetRecommendCourseListRes, err error) {
	recommendList := make([]types.SearchCourseByIdItemByStudent, 0)
	if req.StudentId == 0 {
		list := courseDao.GetRecommendCourseList()
		for _, course := range list {
			recommendList = append(recommendList, types.SearchCourseByIdItemByStudent{
				CourseId:      int32(course.ID),
				CourseName:    course.CourseName,
				CourseImg:     course.CourseImg,
				CourseType:    course.CourseTag,
				CoursePlaySum: course.CoursePlaySum,
				TeacherName:   teacherDao.FindTeacherNameById(int(course.TeacherId)),
			})
		}
	} else {
		studentName := studentDao.SelectByID(req.StudentId).StudentName
		len := redis2.LLen(studentName)
		//获得用户选择的课程标签
		tags := redis2.LRange(studentName, 0, len.Val()).Val()
		//推荐算法
		//实现思路，从tags里获得标签，每个标签查询两个当下该标签最热门的课程，然后封装到切边中推荐给用户
		for _, tag := range tags {
			courses := courseDao.GetRecommendCourseListPersonally(tag)
			for _, course := range courses {
				recommendList = append(recommendList, types.SearchCourseByIdItemByStudent{
					CourseId:      int32(course.ID),
					CourseName:    course.CourseName,
					CourseImg:     course.CourseImg,
					CourseType:    course.CourseTag,
					CoursePlaySum: course.CoursePlaySum,
					TeacherName:   teacherDao.FindTeacherNameById(int(course.TeacherId)),
				})
			}
		}
	}
	return &types.GetRecommendCourseListRes{
		Code:       200,
		CourseList: recommendList,
	}, nil
}
