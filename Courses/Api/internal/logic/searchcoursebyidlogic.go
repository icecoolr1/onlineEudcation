package logic

import (
	"context"
	"onlineEudcation/Courses/Dao"

	"onlineEudcation/Courses/Api/internal/svc"
	"onlineEudcation/Courses/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCourseByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

var courseDao Dao.CourseDaoInterface = new(Dao.CourseDao)

func NewSearchCourseByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) SearchCourseByIdLogic {
	return SearchCourseByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchCourseByIdLogic) SearchCourseById(req types.SearchCourseByIdReq) (resp *types.SearchCourseByIdRes, err error) {
	resList := make([]types.SearchCourseByIdItem, 0)
	courseList := courseDao.FindCourseByTeacherId(int(req.TeacherId))
	for _, res := range courseList {
		item := types.SearchCourseByIdItem{
			CourseId:          int32(res.ID),
			CourseName:        res.CourseName,
			CourseImg:         res.CourseImg,
			CourseType:        res.CourseTag,
			CourseStatus:      res.CourseStatus,
			TeacherId:         res.TeacherId,
			CoursePlaySum:     res.CoursePlaySum,
			CourseVideoNumBer: int32(res.CourseVideoNumBer),
		}
		resList = append(resList, item)
	}

	return &types.SearchCourseByIdRes{
		CourseList: resList,
	}, nil
}
