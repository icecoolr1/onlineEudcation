// Code generated by goctl. DO NOT EDIT!
// Source: videoRpc.proto

package videoservice

import (
	"context"

	"onlineEudcation/Vedio/rpc/rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddVideoRequest    = rpc.AddVideoRequest
	AddVideoResponse   = rpc.AddVideoResponse
	DelVideoReq        = rpc.DelVideoReq
	UpdateVideoRequest = rpc.UpdateVideoRequest

	VideoService interface {
		AddVideo(ctx context.Context, in *AddVideoRequest, opts ...grpc.CallOption) (*AddVideoResponse, error)
		DelVideo(ctx context.Context, in *DelVideoReq, opts ...grpc.CallOption) (*AddVideoResponse, error)
		UpdateVideo(ctx context.Context, in *UpdateVideoRequest, opts ...grpc.CallOption) (*AddVideoResponse, error)
	}

	defaultVideoService struct {
		cli zrpc.Client
	}
)

func NewVideoService(cli zrpc.Client) VideoService {
	return &defaultVideoService{
		cli: cli,
	}
}

func (m *defaultVideoService) AddVideo(ctx context.Context, in *AddVideoRequest, opts ...grpc.CallOption) (*AddVideoResponse, error) {
	client := rpc.NewVideoServiceClient(m.cli.Conn())
	return client.AddVideo(ctx, in, opts...)
}

func (m *defaultVideoService) DelVideo(ctx context.Context, in *DelVideoReq, opts ...grpc.CallOption) (*AddVideoResponse, error) {
	client := rpc.NewVideoServiceClient(m.cli.Conn())
	return client.DelVideo(ctx, in, opts...)
}

func (m *defaultVideoService) UpdateVideo(ctx context.Context, in *UpdateVideoRequest, opts ...grpc.CallOption) (*AddVideoResponse, error) {
	client := rpc.NewVideoServiceClient(m.cli.Conn())
	return client.UpdateVideo(ctx, in, opts...)
}
