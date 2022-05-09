package logic

import (
	"context"
	"onlineEudcation/Classes/Etity"

	"onlineEudcation/Classes/rpc/internal/svc"
	"onlineEudcation/Classes/rpc/rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddClassLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddClassLogic {
	return &AddClassLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddClassLogic) AddClass(in *rpc.AddClassRequest) (*rpc.Request, error) {
	fav := classesDao.IsFav(int(in.GetStudentId()), int(in.CourseId))
	if fav == true {
		return &rpc.Request{
			Code:    300,
			Message: "已收藏",
		}, nil
	} else {
		classesDao.AddClass(Etity.Classes{
			CourseId:  in.CourseId,
			StudentId: in.StudentId,
		})
		return &rpc.Request{
			Code:    200,
			Message: "ok",
		}, nil

	}

}
