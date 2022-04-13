package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"onlineEudcation/Courses/Api/internal/config"
	"onlineEudcation/Courses/Rpc/courserpc"
)

type ServiceContext struct {
	Config    config.Config
	CourseRpc courserpc.CourseRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		CourseRpc: courserpc.NewCourseRpc(zrpc.MustNewClient(c.CourseRpc)),
	}
}
