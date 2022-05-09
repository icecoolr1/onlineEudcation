package logic

import (
	"context"

	"onlineEudcation/Courses/Api/internal/svc"
	"onlineEudcation/Courses/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCourseByNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchCourseByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchCourseByNameLogic {
	return SearchCourseByNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchCourseByNameLogic) SearchCourseByName(req types.SearchCourseByNameReq) (resp *types.SearchCourseByNameRes, err error) {
	courselist := make([]types.SearchCourseByIdItem, 0)
	name := courseDao.FindCourseByCourseName(req.Text)
	if len(name) == 0 {
		return &types.SearchCourseByNameRes{
			Code:       201,
			CourseList: nil,
		}, nil
	} else {
		for _, v := range name {
			courselist = append(courselist, types.SearchCourseByIdItem{
				CourseId:      int32(v.ID),
				CourseName:    v.CourseName,
				CourseImg:     v.CourseImg,
				CourseType:    v.CourseTag,
				CourseStatus:  v.CourseStatus,
				TeacherId:     v.TeacherId,
				CoursePlaySum: v.CoursePlaySum,
			})
		}
		return &types.SearchCourseByNameRes{
			Code:       200,
			CourseList: courselist,
		}, nil
	}

}
