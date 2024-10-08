// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: courieritem.proto

package genproto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CourierLocationService_CreateCourierLocation_FullMethodName = "/ecommerce.CourierLocationService/CreateCourierLocation"
	CourierLocationService_GetCourierLocation_FullMethodName    = "/ecommerce.CourierLocationService/GetCourierLocation"
	CourierLocationService_UpdateCourierLocation_FullMethodName = "/ecommerce.CourierLocationService/UpdateCourierLocation"
	CourierLocationService_DeleteCourierLocation_FullMethodName = "/ecommerce.CourierLocationService/DeleteCourierLocation"
	CourierLocationService_ListCourierLocations_FullMethodName  = "/ecommerce.CourierLocationService/ListCourierLocations"
)

// CourierLocationServiceClient is the client API for CourierLocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourierLocationServiceClient interface {
	CreateCourierLocation(ctx context.Context, in *CreateCourierLocationRequest, opts ...grpc.CallOption) (*Empty, error)
	GetCourierLocation(ctx context.Context, in *GetCourierLocationRequest, opts ...grpc.CallOption) (*CourierLocation, error)
	UpdateCourierLocation(ctx context.Context, in *UpdateCourierLocationRequest, opts ...grpc.CallOption) (*Empty, error)
	DeleteCourierLocation(ctx context.Context, in *DeleteCourierLocationRequest, opts ...grpc.CallOption) (*Empty, error)
	ListCourierLocations(ctx context.Context, in *GetAllCourierLocationsRequest, opts ...grpc.CallOption) (*CourierLocationList, error)
}

type courierLocationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourierLocationServiceClient(cc grpc.ClientConnInterface) CourierLocationServiceClient {
	return &courierLocationServiceClient{cc}
}

func (c *courierLocationServiceClient) CreateCourierLocation(ctx context.Context, in *CreateCourierLocationRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, CourierLocationService_CreateCourierLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierLocationServiceClient) GetCourierLocation(ctx context.Context, in *GetCourierLocationRequest, opts ...grpc.CallOption) (*CourierLocation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierLocation)
	err := c.cc.Invoke(ctx, CourierLocationService_GetCourierLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierLocationServiceClient) UpdateCourierLocation(ctx context.Context, in *UpdateCourierLocationRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, CourierLocationService_UpdateCourierLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierLocationServiceClient) DeleteCourierLocation(ctx context.Context, in *DeleteCourierLocationRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, CourierLocationService_DeleteCourierLocation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierLocationServiceClient) ListCourierLocations(ctx context.Context, in *GetAllCourierLocationsRequest, opts ...grpc.CallOption) (*CourierLocationList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CourierLocationList)
	err := c.cc.Invoke(ctx, CourierLocationService_ListCourierLocations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourierLocationServiceServer is the server API for CourierLocationService service.
// All implementations must embed UnimplementedCourierLocationServiceServer
// for forward compatibility.
type CourierLocationServiceServer interface {
	CreateCourierLocation(context.Context, *CreateCourierLocationRequest) (*Empty, error)
	GetCourierLocation(context.Context, *GetCourierLocationRequest) (*CourierLocation, error)
	UpdateCourierLocation(context.Context, *UpdateCourierLocationRequest) (*Empty, error)
	DeleteCourierLocation(context.Context, *DeleteCourierLocationRequest) (*Empty, error)
	ListCourierLocations(context.Context, *GetAllCourierLocationsRequest) (*CourierLocationList, error)
	mustEmbedUnimplementedCourierLocationServiceServer()
}

// UnimplementedCourierLocationServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCourierLocationServiceServer struct{}

func (UnimplementedCourierLocationServiceServer) CreateCourierLocation(context.Context, *CreateCourierLocationRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourierLocation not implemented")
}
func (UnimplementedCourierLocationServiceServer) GetCourierLocation(context.Context, *GetCourierLocationRequest) (*CourierLocation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourierLocation not implemented")
}
func (UnimplementedCourierLocationServiceServer) UpdateCourierLocation(context.Context, *UpdateCourierLocationRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCourierLocation not implemented")
}
func (UnimplementedCourierLocationServiceServer) DeleteCourierLocation(context.Context, *DeleteCourierLocationRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCourierLocation not implemented")
}
func (UnimplementedCourierLocationServiceServer) ListCourierLocations(context.Context, *GetAllCourierLocationsRequest) (*CourierLocationList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCourierLocations not implemented")
}
func (UnimplementedCourierLocationServiceServer) mustEmbedUnimplementedCourierLocationServiceServer() {
}
func (UnimplementedCourierLocationServiceServer) testEmbeddedByValue() {}

// UnsafeCourierLocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourierLocationServiceServer will
// result in compilation errors.
type UnsafeCourierLocationServiceServer interface {
	mustEmbedUnimplementedCourierLocationServiceServer()
}

func RegisterCourierLocationServiceServer(s grpc.ServiceRegistrar, srv CourierLocationServiceServer) {
	// If the following call pancis, it indicates UnimplementedCourierLocationServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CourierLocationService_ServiceDesc, srv)
}

func _CourierLocationService_CreateCourierLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourierLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierLocationServiceServer).CreateCourierLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierLocationService_CreateCourierLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierLocationServiceServer).CreateCourierLocation(ctx, req.(*CreateCourierLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierLocationService_GetCourierLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCourierLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierLocationServiceServer).GetCourierLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierLocationService_GetCourierLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierLocationServiceServer).GetCourierLocation(ctx, req.(*GetCourierLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierLocationService_UpdateCourierLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourierLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierLocationServiceServer).UpdateCourierLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierLocationService_UpdateCourierLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierLocationServiceServer).UpdateCourierLocation(ctx, req.(*UpdateCourierLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierLocationService_DeleteCourierLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCourierLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierLocationServiceServer).DeleteCourierLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierLocationService_DeleteCourierLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierLocationServiceServer).DeleteCourierLocation(ctx, req.(*DeleteCourierLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierLocationService_ListCourierLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCourierLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierLocationServiceServer).ListCourierLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CourierLocationService_ListCourierLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierLocationServiceServer).ListCourierLocations(ctx, req.(*GetAllCourierLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CourierLocationService_ServiceDesc is the grpc.ServiceDesc for CourierLocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CourierLocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.CourierLocationService",
	HandlerType: (*CourierLocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCourierLocation",
			Handler:    _CourierLocationService_CreateCourierLocation_Handler,
		},
		{
			MethodName: "GetCourierLocation",
			Handler:    _CourierLocationService_GetCourierLocation_Handler,
		},
		{
			MethodName: "UpdateCourierLocation",
			Handler:    _CourierLocationService_UpdateCourierLocation_Handler,
		},
		{
			MethodName: "DeleteCourierLocation",
			Handler:    _CourierLocationService_DeleteCourierLocation_Handler,
		},
		{
			MethodName: "ListCourierLocations",
			Handler:    _CourierLocationService_ListCourierLocations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "courieritem.proto",
}
