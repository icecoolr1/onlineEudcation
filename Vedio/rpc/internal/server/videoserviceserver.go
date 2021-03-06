// Code generated by goctl. DO NOT EDIT!
// Source: videoRpc.proto

package server

import (
	"context"

	"onlineEudcation/Vedio/rpc/internal/logic"
	"onlineEudcation/Vedio/rpc/internal/svc"
	"onlineEudcation/Vedio/rpc/rpc"
)

type VideoServiceServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedVideoServiceServer
}

func NewVideoServiceServer(svcCtx *svc.ServiceContext) *VideoServiceServer {
	return &VideoServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *VideoServiceServer) AddVideo(ctx context.Context, in *rpc.AddVideoRequest) (*rpc.AddVideoResponse, error) {
	l := logic.NewAddVideoLogic(ctx, s.svcCtx)
	return l.AddVideo(in)
}

func (s *VideoServiceServer) DelVideo(ctx context.Context, in *rpc.DelVideoReq) (*rpc.AddVideoResponse, error) {
	l := logic.NewDelVideoLogic(ctx, s.svcCtx)
	return l.DelVideo(in)
}

func (s *VideoServiceServer) UpdateVideo(ctx context.Context, in *rpc.UpdateVideoRequest) (*rpc.AddVideoResponse, error) {
	l := logic.NewUpdateVideoLogic(ctx, s.svcCtx)
	return l.UpdateVideo(in)
}

func (s *VideoServiceServer) VideoHits(ctx context.Context, in *rpc.VideoHitsReq) (*rpc.AddVideoResponse, error) {
	l := logic.NewVideoHitsLogic(ctx, s.svcCtx)
	return l.VideoHits(in)
}
