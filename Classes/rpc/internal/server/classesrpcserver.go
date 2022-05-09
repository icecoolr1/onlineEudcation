// Code generated by goctl. DO NOT EDIT!
// Source: classesrpc.proto

package server

import (
	"context"

	"onlineEudcation/Classes/rpc/internal/logic"
	"onlineEudcation/Classes/rpc/internal/svc"
	"onlineEudcation/Classes/rpc/rpc"
)

type ClassesRpcServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedClassesRpcServer
}

func NewClassesRpcServer(svcCtx *svc.ServiceContext) *ClassesRpcServer {
	return &ClassesRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *ClassesRpcServer) AddClass(ctx context.Context, in *rpc.AddClassRequest) (*rpc.Request, error) {
	l := logic.NewAddClassLogic(ctx, s.svcCtx)
	return l.AddClass(in)
}

func (s *ClassesRpcServer) DeleteClass(ctx context.Context, in *rpc.AddClassRequest) (*rpc.Request, error) {
	l := logic.NewDeleteClassLogic(ctx, s.svcCtx)
	return l.DeleteClass(in)
}
