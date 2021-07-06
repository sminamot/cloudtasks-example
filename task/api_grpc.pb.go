// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package task

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TaskAPIClient is the client API for TaskAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskAPIClient interface {
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	HandleTask(ctx context.Context, in *HandleTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type taskAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskAPIClient(cc grpc.ClientConnInterface) TaskAPIClient {
	return &taskAPIClient{cc}
}

func (c *taskAPIClient) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/task.TaskAPI/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskAPIClient) HandleTask(ctx context.Context, in *HandleTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/task.TaskAPI/HandleTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskAPIServer is the server API for TaskAPI service.
// All implementations must embed UnimplementedTaskAPIServer
// for forward compatibility
type TaskAPIServer interface {
	CreateTask(context.Context, *CreateTaskRequest) (*emptypb.Empty, error)
	HandleTask(context.Context, *HandleTaskRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedTaskAPIServer()
}

// UnimplementedTaskAPIServer must be embedded to have forward compatible implementations.
type UnimplementedTaskAPIServer struct {
}

func (UnimplementedTaskAPIServer) CreateTask(context.Context, *CreateTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedTaskAPIServer) HandleTask(context.Context, *HandleTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleTask not implemented")
}
func (UnimplementedTaskAPIServer) mustEmbedUnimplementedTaskAPIServer() {}

// UnsafeTaskAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskAPIServer will
// result in compilation errors.
type UnsafeTaskAPIServer interface {
	mustEmbedUnimplementedTaskAPIServer()
}

func RegisterTaskAPIServer(s grpc.ServiceRegistrar, srv TaskAPIServer) {
	s.RegisterService(&TaskAPI_ServiceDesc, srv)
}

func _TaskAPI_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskAPIServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskAPI/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskAPIServer).CreateTask(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskAPI_HandleTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskAPIServer).HandleTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/task.TaskAPI/HandleTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskAPIServer).HandleTask(ctx, req.(*HandleTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskAPI_ServiceDesc is the grpc.ServiceDesc for TaskAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "task.TaskAPI",
	HandlerType: (*TaskAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTask",
			Handler:    _TaskAPI_CreateTask_Handler,
		},
		{
			MethodName: "HandleTask",
			Handler:    _TaskAPI_HandleTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}