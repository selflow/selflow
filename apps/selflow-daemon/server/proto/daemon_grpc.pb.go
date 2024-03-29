// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: apps/selflow-daemon/server/proto/daemon.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DaemonClient is the client API for Daemon service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DaemonClient interface {
	StartRun(ctx context.Context, in *StartRun_Request, opts ...grpc.CallOption) (*StartRun_Response, error)
	GetLogStream(ctx context.Context, in *GetLogStream_Request, opts ...grpc.CallOption) (Daemon_GetLogStreamClient, error)
	GetRunStatus(ctx context.Context, in *GetRunStatus_Request, opts ...grpc.CallOption) (*GetRunStatus_Response, error)
}

type daemonClient struct {
	cc grpc.ClientConnInterface
}

func NewDaemonClient(cc grpc.ClientConnInterface) DaemonClient {
	return &daemonClient{cc}
}

func (c *daemonClient) StartRun(ctx context.Context, in *StartRun_Request, opts ...grpc.CallOption) (*StartRun_Response, error) {
	out := new(StartRun_Response)
	err := c.cc.Invoke(ctx, "/Daemon/StartRun", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *daemonClient) GetLogStream(ctx context.Context, in *GetLogStream_Request, opts ...grpc.CallOption) (Daemon_GetLogStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Daemon_ServiceDesc.Streams[0], "/Daemon/GetLogStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &daemonGetLogStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Daemon_GetLogStreamClient interface {
	Recv() (*Log, error)
	grpc.ClientStream
}

type daemonGetLogStreamClient struct {
	grpc.ClientStream
}

func (x *daemonGetLogStreamClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *daemonClient) GetRunStatus(ctx context.Context, in *GetRunStatus_Request, opts ...grpc.CallOption) (*GetRunStatus_Response, error) {
	out := new(GetRunStatus_Response)
	err := c.cc.Invoke(ctx, "/Daemon/GetRunStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DaemonServer is the server API for Daemon service.
// All implementations must embed UnimplementedDaemonServer
// for forward compatibility
type DaemonServer interface {
	StartRun(context.Context, *StartRun_Request) (*StartRun_Response, error)
	GetLogStream(*GetLogStream_Request, Daemon_GetLogStreamServer) error
	GetRunStatus(context.Context, *GetRunStatus_Request) (*GetRunStatus_Response, error)
	mustEmbedUnimplementedDaemonServer()
}

// UnimplementedDaemonServer must be embedded to have forward compatible implementations.
type UnimplementedDaemonServer struct {
}

func (UnimplementedDaemonServer) StartRun(context.Context, *StartRun_Request) (*StartRun_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRun not implemented")
}
func (UnimplementedDaemonServer) GetLogStream(*GetLogStream_Request, Daemon_GetLogStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetLogStream not implemented")
}
func (UnimplementedDaemonServer) GetRunStatus(context.Context, *GetRunStatus_Request) (*GetRunStatus_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRunStatus not implemented")
}
func (UnimplementedDaemonServer) mustEmbedUnimplementedDaemonServer() {}

// UnsafeDaemonServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DaemonServer will
// result in compilation errors.
type UnsafeDaemonServer interface {
	mustEmbedUnimplementedDaemonServer()
}

func RegisterDaemonServer(s grpc.ServiceRegistrar, srv DaemonServer) {
	s.RegisterService(&Daemon_ServiceDesc, srv)
}

func _Daemon_StartRun_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRun_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).StartRun(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Daemon/StartRun",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).StartRun(ctx, req.(*StartRun_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Daemon_GetLogStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetLogStream_Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DaemonServer).GetLogStream(m, &daemonGetLogStreamServer{stream})
}

type Daemon_GetLogStreamServer interface {
	Send(*Log) error
	grpc.ServerStream
}

type daemonGetLogStreamServer struct {
	grpc.ServerStream
}

func (x *daemonGetLogStreamServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

func _Daemon_GetRunStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRunStatus_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DaemonServer).GetRunStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Daemon/GetRunStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DaemonServer).GetRunStatus(ctx, req.(*GetRunStatus_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Daemon_ServiceDesc is the grpc.ServiceDesc for Daemon service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Daemon_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Daemon",
	HandlerType: (*DaemonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartRun",
			Handler:    _Daemon_StartRun_Handler,
		},
		{
			MethodName: "GetRunStatus",
			Handler:    _Daemon_GetRunStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetLogStream",
			Handler:       _Daemon_GetLogStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "apps/selflow-daemon/server/proto/daemon.proto",
}
