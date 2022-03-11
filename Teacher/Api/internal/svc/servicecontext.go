package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"onlineEudcation/Teacher/Api/internal/config"
	"onlineEudcation/Teacher/Rpc/teacherrpc"
)

type ServiceContext struct {
	Config     config.Config
	TeacherRpc teacherrpc.TeacherRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		TeacherRpc: teacherrpc.NewTeacherRpc(zrpc.MustNewClient(c.TeacherRpc)),
	}
}
