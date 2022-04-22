package logic

import (
	"context"
	"onlineEudcation/Courses/Rpc/Rpc"

	"onlineEudcation/Courses/Api/internal/svc"
	"onlineEudcation/Courses/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) DelCourseLogic {
	return DelCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelCourseLogic) DelCourse(req types.DelCourseReq) (resp *types.DelCourseRes, err error) {
	resList := make([]types.SearchCourseByIdItem, 0)
	_, err = l.svcCtx.CourseRpc.DeleteCourse(l.ctx, &Rpc.DeleteCourseReq{
		CourseId: req.CourseId,
	})
	if err != nil {
		return nil, err
	}
	// 返回更新后的数据
	courses := courseDao.FindCourseByTeacherId(int(req.TeacherId))
	for _, course := range courses {
		resList = append(resList, types.SearchCourseByIdItem{
			CourseId:          int32(course.ID),
			CourseName:        course.CourseName,
			CourseImg:         course.CourseImg,
			CourseType:        course.CourseTag,
			CourseStatus:      course.CourseStatus,
			TeacherId:         course.TeacherId,
			CoursePlaySum:     course.CoursePlaySum,
			CourseVideoNumBer: int32(course.CourseVideoNumBer),
		})
	}
	return &types.DelCourseRes{
		Code:       200,
		CourseList: resList,
	}, nil
}
