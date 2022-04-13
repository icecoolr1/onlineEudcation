// Code generated by goctl. DO NOT EDIT!
// Source: courseRpc.proto

package courserpc

import (
	"context"

	"onlineEudcation/Courses/Rpc/Rpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CourseAddReq = Rpc.CourseAddReq
	CourseResp   = Rpc.CourseResp

	CourseRpc interface {
		AddCourse(ctx context.Context, in *CourseAddReq, opts ...grpc.CallOption) (*CourseResp, error)
	}

	defaultCourseRpc struct {
		cli zrpc.Client
	}
)

func NewCourseRpc(cli zrpc.Client) CourseRpc {
	return &defaultCourseRpc{
		cli: cli,
	}
}

func (m *defaultCourseRpc) AddCourse(ctx context.Context, in *CourseAddReq, opts ...grpc.CallOption) (*CourseResp, error) {
	client := Rpc.NewCourseRpcClient(m.cli.Conn())
	return client.AddCourse(ctx, in, opts...)
}