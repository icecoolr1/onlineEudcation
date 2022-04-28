package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"onlineEudcation/Student/api/internal/config"
	"onlineEudcation/Student/rpc/studentrpc"
)

type ServiceContext struct {
	Config     config.Config
	StudentRpc studentrpc.StudentRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		StudentRpc: studentrpc.NewStudentRpc(zrpc.MustNewClient(c.StudentRpc)),
	}
}
