// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ActivityServiceClient is the client API for ActivityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityServiceClient interface {
	GetActivities(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error)
}

type activityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityServiceClient(cc grpc.ClientConnInterface) ActivityServiceClient {
	return &activityServiceClient{cc}
}

func (c *activityServiceClient) GetActivities(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error) {
	out := new(ActivityResponse)
	err := c.cc.Invoke(ctx, "/demo.account.ActivityService/GetActivities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServiceServer is the server API for ActivityService service.
// All implementations must embed UnimplementedActivityServiceServer
// for forward compatibility
type ActivityServiceServer interface {
	GetActivities(context.Context, *ActivityRequest) (*ActivityResponse, error)
	mustEmbedUnimplementedActivityServiceServer()
}

// UnimplementedActivityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedActivityServiceServer struct {
}

func (UnimplementedActivityServiceServer) GetActivities(context.Context, *ActivityRequest) (*ActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivities not implemented")
}
func (UnimplementedActivityServiceServer) mustEmbedUnimplementedActivityServiceServer() {}

// UnsafeActivityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServiceServer will
// result in compilation errors.
type UnsafeActivityServiceServer interface {
	mustEmbedUnimplementedActivityServiceServer()
}

func RegisterActivityServiceServer(s grpc.ServiceRegistrar, srv ActivityServiceServer) {
	s.RegisterService(&_ActivityService_serviceDesc, srv)
}

func _ActivityService_GetActivities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).GetActivities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.account.ActivityService/GetActivities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).GetActivities(ctx, req.(*ActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActivityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demo.account.ActivityService",
	HandlerType: (*ActivityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActivities",
			Handler:    _ActivityService_GetActivities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/activity.proto",
}
