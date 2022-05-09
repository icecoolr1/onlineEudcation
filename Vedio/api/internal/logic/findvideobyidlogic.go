package logic

import (
	"context"
	"fmt"
	"onlineEudcation/Vedio/Etity"

	"onlineEudcation/Vedio/api/internal/svc"
	"onlineEudcation/Vedio/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

var video Etity.Video

type FindVideoByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindVideoByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) FindVideoByIdLogic {
	return FindVideoByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindVideoByIdLogic) FindVideoById(req types.FindVideoByIdReq) (resp *types.FindVideoByIdRes, err error) {

	video = videoDao.FindVideoByVideoId(int(req.VideoId))
	fmt.Println(int(req.VideoId), "hhh")
	fmt.Println(video, "tttt")
	return &types.FindVideoByIdRes{
		Code: 200,
		VideoInfo: types.Video{
			VideoName:    video.VideoName,
			VideoFileUrl: video.VideoFileUrl,
			VideoSeq:     video.VideoSeq,
			VedioId:      int32(video.ID),
		},
	}, nil
}
