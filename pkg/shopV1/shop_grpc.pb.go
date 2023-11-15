// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: shop.proto

package shopV1

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ShopClient is the client API for Shop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShopClient interface {
	GetProduct(ctx context.Context, in *GetProduct_Request, opts ...grpc.CallOption) (*GetProduct_Response, error)
	GetAllProduct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllProducts_Response, error)
	AddProductToBasket(ctx context.Context, in *AddProductToBasket_Request, opts ...grpc.CallOption) (*empty.Empty, error)
	UpdateBasket(ctx context.Context, in *UpdateBasket_Request, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteBasket(ctx context.Context, in *DeleteBasket_Request, opts ...grpc.CallOption) (*empty.Empty, error)
	GetBasket(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetBasket_Response, error)
	CreateOrder(ctx context.Context, in *Order_Request, opts ...grpc.CallOption) (*Order_Response, error)
}

type shopClient struct {
	cc grpc.ClientConnInterface
}

func NewShopClient(cc grpc.ClientConnInterface) ShopClient {
	return &shopClient{cc}
}

func (c *shopClient) GetProduct(ctx context.Context, in *GetProduct_Request, opts ...grpc.CallOption) (*GetProduct_Response, error) {
	out := new(GetProduct_Response)
	err := c.cc.Invoke(ctx, "/Shop/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) GetAllProduct(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAllProducts_Response, error) {
	out := new(GetAllProducts_Response)
	err := c.cc.Invoke(ctx, "/Shop/GetAllProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) AddProductToBasket(ctx context.Context, in *AddProductToBasket_Request, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Shop/AddProductToBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) UpdateBasket(ctx context.Context, in *UpdateBasket_Request, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Shop/UpdateBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) DeleteBasket(ctx context.Context, in *DeleteBasket_Request, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Shop/DeleteBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) GetBasket(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetBasket_Response, error) {
	out := new(GetBasket_Response)
	err := c.cc.Invoke(ctx, "/Shop/GetBasket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shopClient) CreateOrder(ctx context.Context, in *Order_Request, opts ...grpc.CallOption) (*Order_Response, error) {
	out := new(Order_Response)
	err := c.cc.Invoke(ctx, "/Shop/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShopServer is the server API for Shop service.
// All implementations must embed UnimplementedShopServer
// for forward compatibility
type ShopServer interface {
	GetProduct(context.Context, *GetProduct_Request) (*GetProduct_Response, error)
	GetAllProduct(context.Context, *empty.Empty) (*GetAllProducts_Response, error)
	AddProductToBasket(context.Context, *AddProductToBasket_Request) (*empty.Empty, error)
	UpdateBasket(context.Context, *UpdateBasket_Request) (*empty.Empty, error)
	DeleteBasket(context.Context, *DeleteBasket_Request) (*empty.Empty, error)
	GetBasket(context.Context, *empty.Empty) (*GetBasket_Response, error)
	CreateOrder(context.Context, *Order_Request) (*Order_Response, error)
	mustEmbedUnimplementedShopServer()
}

// UnimplementedShopServer must be embedded to have forward compatible implementations.
type UnimplementedShopServer struct {
}

func (UnimplementedShopServer) GetProduct(context.Context, *GetProduct_Request) (*GetProduct_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedShopServer) GetAllProduct(context.Context, *empty.Empty) (*GetAllProducts_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProduct not implemented")
}
func (UnimplementedShopServer) AddProductToBasket(context.Context, *AddProductToBasket_Request) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProductToBasket not implemented")
}
func (UnimplementedShopServer) UpdateBasket(context.Context, *UpdateBasket_Request) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBasket not implemented")
}
func (UnimplementedShopServer) DeleteBasket(context.Context, *DeleteBasket_Request) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBasket not implemented")
}
func (UnimplementedShopServer) GetBasket(context.Context, *empty.Empty) (*GetBasket_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBasket not implemented")
}
func (UnimplementedShopServer) CreateOrder(context.Context, *Order_Request) (*Order_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedShopServer) mustEmbedUnimplementedShopServer() {}

// UnsafeShopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShopServer will
// result in compilation errors.
type UnsafeShopServer interface {
	mustEmbedUnimplementedShopServer()
}

func RegisterShopServer(s grpc.ServiceRegistrar, srv ShopServer) {
	s.RegisterService(&Shop_ServiceDesc, srv)
}

func _Shop_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProduct_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shop/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).GetProduct(ctx, req.(*GetProduct_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_GetAllProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).GetAllProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shop/GetAllProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).GetAllProduct(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_AddProductToBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductToBasket_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).AddProductToBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shop/AddProductToBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).AddProductToBasket(ctx, req.(*AddProductToBasket_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_UpdateBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBasket_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).UpdateBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shop/UpdateBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).UpdateBasket(ctx, req.(*UpdateBasket_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_DeleteBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBasket_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).DeleteBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shop/DeleteBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).DeleteBasket(ctx, req.(*DeleteBasket_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_GetBasket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).GetBasket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shop/GetBasket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).GetBasket(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shop_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShopServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shop/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShopServer).CreateOrder(ctx, req.(*Order_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Shop_ServiceDesc is the grpc.ServiceDesc for Shop service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Shop_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Shop",
	HandlerType: (*ShopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _Shop_GetProduct_Handler,
		},
		{
			MethodName: "GetAllProduct",
			Handler:    _Shop_GetAllProduct_Handler,
		},
		{
			MethodName: "AddProductToBasket",
			Handler:    _Shop_AddProductToBasket_Handler,
		},
		{
			MethodName: "UpdateBasket",
			Handler:    _Shop_UpdateBasket_Handler,
		},
		{
			MethodName: "DeleteBasket",
			Handler:    _Shop_DeleteBasket_Handler,
		},
		{
			MethodName: "GetBasket",
			Handler:    _Shop_GetBasket_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _Shop_CreateOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shop.proto",
}
