package logic

import (
	"context"

	"onlineEudcation/Student/api/internal/svc"
	"onlineEudcation/Student/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendChosenTagsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendChosenTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) SendChosenTagsLogic {
	return SendChosenTagsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendChosenTagsLogic) SendChosenTags(req types.SendChosenTags) (resp *types.Response, err error) {
	if len(req.Tags) == 0 {
		return &types.Response{
			Code:    200,
			Message: "没有选择任何标签",
		}, nil
	} else {
		for _, tag := range req.Tags {
			redis2.LPush(req.StudentName, tag)
		}
		return &types.Response{
			Code:    200,
			Message: "标签添加成功",
		}, nil
	}

}
