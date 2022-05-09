package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"onlineEudcation/Classes/api/internal/config"
	"onlineEudcation/Classes/rpc/classesrpc"
)

type ServiceContext struct {
	Config     config.Config
	ClassesRpc classesrpc.ClassesRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ClassesRpc: classesrpc.NewClassesRpc(zrpc.MustNewClient(c.ClassesRpc)),
	}
}
