package logic

import (
	"context"

	"onlineEudcation/Admin/api/internal/svc"
	"onlineEudcation/Admin/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCensorListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCensorListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetCensorListLogic {
	return GetCensorListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCensorListLogic) GetCensorList(req types.GetReq) (resp *types.GetRes, err error) {
	resList := make([]types.CensorItem, 0)
	result := adminDao.GetResult(req.TeacherId)
	for _, censorship := range result {
		resList = append(resList, types.CensorItem{
			ID:           int(censorship.ID),
			CourseId:     censorship.CourseId,
			TeacherId:    censorship.TeacherId,
			CensorStatus: censorship.Result,
			Message:      censorship.Message,
		})
	}
	return &types.GetRes{
		CensorList: resList,
	}, nil
}
