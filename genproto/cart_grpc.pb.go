// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: cart.proto

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
	CartService_CreateCart_FullMethodName  = "/ecommerce.CartService/CreateCart"
	CartService_GetCart_FullMethodName     = "/ecommerce.CartService/GetCart"
	CartService_GetAllCarts_FullMethodName = "/ecommerce.CartService/GetAllCarts"
	CartService_UpdateCart_FullMethodName  = "/ecommerce.CartService/UpdateCart"
	CartService_DeleteCart_FullMethodName  = "/ecommerce.CartService/DeleteCart"
	CartService_SearchCarts_FullMethodName = "/ecommerce.CartService/SearchCarts"
)

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	CreateCart(ctx context.Context, in *CreateCartReq, opts ...grpc.CallOption) (*CartEmpty, error)
	GetCart(ctx context.Context, in *GetByIdCartRequest, opts ...grpc.CallOption) (*Cart, error)
	GetAllCarts(ctx context.Context, in *GetAllCartsReq, opts ...grpc.CallOption) (*GetAllCartsRes, error)
	UpdateCart(ctx context.Context, in *UpdateCartReq, opts ...grpc.CallOption) (*UpdateCartRes, error)
	DeleteCart(ctx context.Context, in *DeleteCartRequest, opts ...grpc.CallOption) (*DeleteCartResp, error)
	SearchCarts(ctx context.Context, in *SearchCartsReq, opts ...grpc.CallOption) (*SearchCartsRes, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) CreateCart(ctx context.Context, in *CreateCartReq, opts ...grpc.CallOption) (*CartEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CartEmpty)
	err := c.cc.Invoke(ctx, CartService_CreateCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCart(ctx context.Context, in *GetByIdCartRequest, opts ...grpc.CallOption) (*Cart, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Cart)
	err := c.cc.Invoke(ctx, CartService_GetCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetAllCarts(ctx context.Context, in *GetAllCartsReq, opts ...grpc.CallOption) (*GetAllCartsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllCartsRes)
	err := c.cc.Invoke(ctx, CartService_GetAllCarts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) UpdateCart(ctx context.Context, in *UpdateCartReq, opts ...grpc.CallOption) (*UpdateCartRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCartRes)
	err := c.cc.Invoke(ctx, CartService_UpdateCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) DeleteCart(ctx context.Context, in *DeleteCartRequest, opts ...grpc.CallOption) (*DeleteCartResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCartResp)
	err := c.cc.Invoke(ctx, CartService_DeleteCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) SearchCarts(ctx context.Context, in *SearchCartsReq, opts ...grpc.CallOption) (*SearchCartsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchCartsRes)
	err := c.cc.Invoke(ctx, CartService_SearchCarts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility.
type CartServiceServer interface {
	CreateCart(context.Context, *CreateCartReq) (*CartEmpty, error)
	GetCart(context.Context, *GetByIdCartRequest) (*Cart, error)
	GetAllCarts(context.Context, *GetAllCartsReq) (*GetAllCartsRes, error)
	UpdateCart(context.Context, *UpdateCartReq) (*UpdateCartRes, error)
	DeleteCart(context.Context, *DeleteCartRequest) (*DeleteCartResp, error)
	SearchCarts(context.Context, *SearchCartsReq) (*SearchCartsRes, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCartServiceServer struct{}

func (UnimplementedCartServiceServer) CreateCart(context.Context, *CreateCartReq) (*CartEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCart not implemented")
}
func (UnimplementedCartServiceServer) GetCart(context.Context, *GetByIdCartRequest) (*Cart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedCartServiceServer) GetAllCarts(context.Context, *GetAllCartsReq) (*GetAllCartsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCarts not implemented")
}
func (UnimplementedCartServiceServer) UpdateCart(context.Context, *UpdateCartReq) (*UpdateCartRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCart not implemented")
}
func (UnimplementedCartServiceServer) DeleteCart(context.Context, *DeleteCartRequest) (*DeleteCartResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCart not implemented")
}
func (UnimplementedCartServiceServer) SearchCarts(context.Context, *SearchCartsReq) (*SearchCartsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchCarts not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}
func (UnimplementedCartServiceServer) testEmbeddedByValue()                     {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	// If the following call pancis, it indicates UnimplementedCartServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_CreateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CreateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_CreateCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CreateCart(ctx, req.(*CreateCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByIdCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_GetCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCart(ctx, req.(*GetByIdCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetAllCarts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCartsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetAllCarts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_GetAllCarts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetAllCarts(ctx, req.(*GetAllCartsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_UpdateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).UpdateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_UpdateCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).UpdateCart(ctx, req.(*UpdateCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_DeleteCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).DeleteCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_DeleteCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).DeleteCart(ctx, req.(*DeleteCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_SearchCarts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchCartsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).SearchCarts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CartService_SearchCarts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).SearchCarts(ctx, req.(*SearchCartsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecommerce.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCart",
			Handler:    _CartService_CreateCart_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _CartService_GetCart_Handler,
		},
		{
			MethodName: "GetAllCarts",
			Handler:    _CartService_GetAllCarts_Handler,
		},
		{
			MethodName: "UpdateCart",
			Handler:    _CartService_UpdateCart_Handler,
		},
		{
			MethodName: "DeleteCart",
			Handler:    _CartService_DeleteCart_Handler,
		},
		{
			MethodName: "SearchCarts",
			Handler:    _CartService_SearchCarts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cart.proto",
}
