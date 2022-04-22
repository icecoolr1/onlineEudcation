package logic

import (
	"context"
	"onlineEudcation/Courses/Rpc/Rpc"

	"onlineEudcation/Courses/Api/internal/svc"
	"onlineEudcation/Courses/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCourseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateCourseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UpdateCourseInfoLogic {
	return UpdateCourseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCourseInfoLogic) UpdateCourseInfo(req types.UpdateCourseInfoReq) (resp *types.Response, err error) {
	_, err = l.svcCtx.CourseRpc.UpdateCourse(l.ctx, &Rpc.UpdateCourseReq{
		CourseId:   req.CourseId,
		CourseName: req.CourseName,
		ImgUrl:     req.ImgUrl,
		CourseType: req.CourseTag,
	})
	if err != nil {
		return nil, err
	} else {
		return &types.Response{
			Code:    200,
			Message: "更新成功",
			Res:     true,
		}, nil
	}
}
