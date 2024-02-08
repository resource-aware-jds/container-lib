// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// WorkerNodeContainerReceiverClient is the client API for WorkerNodeContainerReceiver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerNodeContainerReceiverClient interface {
	SubmitSuccessTask(ctx context.Context, in *SubmitSuccessTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ReportTaskFailure(ctx context.Context, in *ReportTaskFailureRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetTaskFromQueue(ctx context.Context, in *GetTaskPayload, opts ...grpc.CallOption) (*Task, error)
}

type workerNodeContainerReceiverClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerNodeContainerReceiverClient(cc grpc.ClientConnInterface) WorkerNodeContainerReceiverClient {
	return &workerNodeContainerReceiverClient{cc}
}

func (c *workerNodeContainerReceiverClient) SubmitSuccessTask(ctx context.Context, in *SubmitSuccessTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/WorkerNode.WorkerNodeContainerReceiver/SubmitSuccessTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerNodeContainerReceiverClient) ReportTaskFailure(ctx context.Context, in *ReportTaskFailureRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/WorkerNode.WorkerNodeContainerReceiver/ReportTaskFailure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerNodeContainerReceiverClient) GetTaskFromQueue(ctx context.Context, in *GetTaskPayload, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/WorkerNode.WorkerNodeContainerReceiver/GetTaskFromQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerNodeContainerReceiverServer is the server API for WorkerNodeContainerReceiver service.
// All implementations must embed UnimplementedWorkerNodeContainerReceiverServer
// for forward compatibility
type WorkerNodeContainerReceiverServer interface {
	SubmitSuccessTask(context.Context, *SubmitSuccessTaskRequest) (*emptypb.Empty, error)
	ReportTaskFailure(context.Context, *ReportTaskFailureRequest) (*emptypb.Empty, error)
	GetTaskFromQueue(context.Context, *GetTaskPayload) (*Task, error)
	mustEmbedUnimplementedWorkerNodeContainerReceiverServer()
}

// UnimplementedWorkerNodeContainerReceiverServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerNodeContainerReceiverServer struct {
}

func (UnimplementedWorkerNodeContainerReceiverServer) SubmitSuccessTask(context.Context, *SubmitSuccessTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitSuccessTask not implemented")
}
func (UnimplementedWorkerNodeContainerReceiverServer) ReportTaskFailure(context.Context, *ReportTaskFailureRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportTaskFailure not implemented")
}
func (UnimplementedWorkerNodeContainerReceiverServer) GetTaskFromQueue(context.Context, *GetTaskPayload) (*Task, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskFromQueue not implemented")
}
func (UnimplementedWorkerNodeContainerReceiverServer) mustEmbedUnimplementedWorkerNodeContainerReceiverServer() {
}

// UnsafeWorkerNodeContainerReceiverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerNodeContainerReceiverServer will
// result in compilation errors.
type UnsafeWorkerNodeContainerReceiverServer interface {
	mustEmbedUnimplementedWorkerNodeContainerReceiverServer()
}

func RegisterWorkerNodeContainerReceiverServer(s grpc.ServiceRegistrar, srv WorkerNodeContainerReceiverServer) {
	s.RegisterService(&WorkerNodeContainerReceiver_ServiceDesc, srv)
}

func _WorkerNodeContainerReceiver_SubmitSuccessTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitSuccessTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerNodeContainerReceiverServer).SubmitSuccessTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkerNode.WorkerNodeContainerReceiver/SubmitSuccessTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerNodeContainerReceiverServer).SubmitSuccessTask(ctx, req.(*SubmitSuccessTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerNodeContainerReceiver_ReportTaskFailure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportTaskFailureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerNodeContainerReceiverServer).ReportTaskFailure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkerNode.WorkerNodeContainerReceiver/ReportTaskFailure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerNodeContainerReceiverServer).ReportTaskFailure(ctx, req.(*ReportTaskFailureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerNodeContainerReceiver_GetTaskFromQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskPayload)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerNodeContainerReceiverServer).GetTaskFromQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkerNode.WorkerNodeContainerReceiver/GetTaskFromQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerNodeContainerReceiverServer).GetTaskFromQueue(ctx, req.(*GetTaskPayload))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkerNodeContainerReceiver_ServiceDesc is the grpc.ServiceDesc for WorkerNodeContainerReceiver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkerNodeContainerReceiver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WorkerNode.WorkerNodeContainerReceiver",
	HandlerType: (*WorkerNodeContainerReceiverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitSuccessTask",
			Handler:    _WorkerNodeContainerReceiver_SubmitSuccessTask_Handler,
		},
		{
			MethodName: "ReportTaskFailure",
			Handler:    _WorkerNodeContainerReceiver_ReportTaskFailure_Handler,
		},
		{
			MethodName: "GetTaskFromQueue",
			Handler:    _WorkerNodeContainerReceiver_GetTaskFromQueue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "workernode.proto",
}
