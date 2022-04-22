package logic

import (
	"context"
	"onlineEudcation/Courses/Rpc/Rpc"
	"onlineEudcation/Courses/Rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCourseLogic {
	return &DeleteCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCourseLogic) DeleteCourse(in *Rpc.DeleteCourseReq) (*Rpc.DeleteCourseResp, error) {

	CourseDao.DelCourse([]int{int(in.CourseId)})

	return &Rpc.DeleteCourseResp{
		Msg:  "删除成功",
		Code: 200,
		Res:  true,
	}, nil
}
