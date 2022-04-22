package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"onlineEudcation/Vedio/api/internal/config"
	"onlineEudcation/Vedio/rpc/videoservice"
)

type ServiceContext struct {
	Config   config.Config
	VideoRpc videoservice.VideoService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		VideoRpc: videoservice.NewVideoService(zrpc.MustNewClient(c.VideoRpc)),
	}
}
