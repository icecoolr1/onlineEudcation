// Code generated by goctl. DO NOT EDIT!
// Source: studentrpc.proto

package studentrpc

import (
	"context"

	"onlineEudcation/Student/rpc/rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Student       = rpc.Student
	StudentAddReq = rpc.StudentAddReq
	StudentAddRes = rpc.StudentAddRes

	StudentRpc interface {
		AddStudent(ctx context.Context, in *StudentAddReq, opts ...grpc.CallOption) (*StudentAddRes, error)
	}

	defaultStudentRpc struct {
		cli zrpc.Client
	}
)

func NewStudentRpc(cli zrpc.Client) StudentRpc {
	return &defaultStudentRpc{
		cli: cli,
	}
}

func (m *defaultStudentRpc) AddStudent(ctx context.Context, in *StudentAddReq, opts ...grpc.CallOption) (*StudentAddRes, error) {
	client := rpc.NewStudentRpcClient(m.cli.Conn())
	return client.AddStudent(ctx, in, opts...)
}
