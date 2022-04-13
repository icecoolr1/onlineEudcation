package logic

import (
	"context"
	"fmt"

	"onlineEudcation/Courses/Api/internal/svc"
	"onlineEudcation/Courses/Api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCourseImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCourseImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddCourseImgLogic {
	return AddCourseImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCourseImgLogic) AddCourseImg(req types.CourseImg) (resp *types.Response, err error) {
	fmt.Println(req.File, "图片参数?")

	return &types.Response{
		Code:    0,
		Message: "success",
		Res:     true,
	}, nil
}
