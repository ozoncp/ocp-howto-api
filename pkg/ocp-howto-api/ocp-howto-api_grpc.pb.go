// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ocp_howto_api

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

// OcpHowtoApiClient is the client API for OcpHowtoApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OcpHowtoApiClient interface {
	// Создает новую сущность howto и возвращает её идентификатор
	CreateHowtoV1(ctx context.Context, in *CreateHowtoV1Request, opts ...grpc.CallOption) (*CreateHowtoV1Response, error)
	// Создает несколько новых сущностей howto
	MultiCreateHowtoV1(ctx context.Context, in *MultiCreateHowtoV1Request, opts ...grpc.CallOption) (*MultiCreateHowtoV1Response, error)
	// Обновляет сущность howto
	UpdateHowtoV1(ctx context.Context, in *UpdateHowtoV1Request, opts ...grpc.CallOption) (*UpdateHowtoV1Response, error)
	// Возвращает полное описание сущности howto по её идентификатору
	DescribeHowtoV1(ctx context.Context, in *DescribeHowtoV1Request, opts ...grpc.CallOption) (*DescribeHowtoV1Response, error)
	// Возвращает список сущностей howto
	ListHowtosV1(ctx context.Context, in *ListHowtosV1Request, opts ...grpc.CallOption) (*ListHowtosV1Response, error)
	// Удаляет сущность howto по её идентификатору
	RemoveHowtoV1(ctx context.Context, in *RemoveHowtoV1Request, opts ...grpc.CallOption) (*RemoveHowtoV1Response, error)
}

type ocpHowtoApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOcpHowtoApiClient(cc grpc.ClientConnInterface) OcpHowtoApiClient {
	return &ocpHowtoApiClient{cc}
}

func (c *ocpHowtoApiClient) CreateHowtoV1(ctx context.Context, in *CreateHowtoV1Request, opts ...grpc.CallOption) (*CreateHowtoV1Response, error) {
	out := new(CreateHowtoV1Response)
	err := c.cc.Invoke(ctx, "/ocp.howto.api.OcpHowtoApi/CreateHowtoV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpHowtoApiClient) MultiCreateHowtoV1(ctx context.Context, in *MultiCreateHowtoV1Request, opts ...grpc.CallOption) (*MultiCreateHowtoV1Response, error) {
	out := new(MultiCreateHowtoV1Response)
	err := c.cc.Invoke(ctx, "/ocp.howto.api.OcpHowtoApi/MultiCreateHowtoV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpHowtoApiClient) UpdateHowtoV1(ctx context.Context, in *UpdateHowtoV1Request, opts ...grpc.CallOption) (*UpdateHowtoV1Response, error) {
	out := new(UpdateHowtoV1Response)
	err := c.cc.Invoke(ctx, "/ocp.howto.api.OcpHowtoApi/UpdateHowtoV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpHowtoApiClient) DescribeHowtoV1(ctx context.Context, in *DescribeHowtoV1Request, opts ...grpc.CallOption) (*DescribeHowtoV1Response, error) {
	out := new(DescribeHowtoV1Response)
	err := c.cc.Invoke(ctx, "/ocp.howto.api.OcpHowtoApi/DescribeHowtoV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpHowtoApiClient) ListHowtosV1(ctx context.Context, in *ListHowtosV1Request, opts ...grpc.CallOption) (*ListHowtosV1Response, error) {
	out := new(ListHowtosV1Response)
	err := c.cc.Invoke(ctx, "/ocp.howto.api.OcpHowtoApi/ListHowtosV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ocpHowtoApiClient) RemoveHowtoV1(ctx context.Context, in *RemoveHowtoV1Request, opts ...grpc.CallOption) (*RemoveHowtoV1Response, error) {
	out := new(RemoveHowtoV1Response)
	err := c.cc.Invoke(ctx, "/ocp.howto.api.OcpHowtoApi/RemoveHowtoV1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OcpHowtoApiServer is the server API for OcpHowtoApi service.
// All implementations must embed UnimplementedOcpHowtoApiServer
// for forward compatibility
type OcpHowtoApiServer interface {
	// Создает новую сущность howto и возвращает её идентификатор
	CreateHowtoV1(context.Context, *CreateHowtoV1Request) (*CreateHowtoV1Response, error)
	// Создает несколько новых сущностей howto
	MultiCreateHowtoV1(context.Context, *MultiCreateHowtoV1Request) (*MultiCreateHowtoV1Response, error)
	// Обновляет сущность howto
	UpdateHowtoV1(context.Context, *UpdateHowtoV1Request) (*UpdateHowtoV1Response, error)
	// Возвращает полное описание сущности howto по её идентификатору
	DescribeHowtoV1(context.Context, *DescribeHowtoV1Request) (*DescribeHowtoV1Response, error)
	// Возвращает список сущностей howto
	ListHowtosV1(context.Context, *ListHowtosV1Request) (*ListHowtosV1Response, error)
	// Удаляет сущность howto по её идентификатору
	RemoveHowtoV1(context.Context, *RemoveHowtoV1Request) (*RemoveHowtoV1Response, error)
	mustEmbedUnimplementedOcpHowtoApiServer()
}

// UnimplementedOcpHowtoApiServer must be embedded to have forward compatible implementations.
type UnimplementedOcpHowtoApiServer struct {
}

func (UnimplementedOcpHowtoApiServer) CreateHowtoV1(context.Context, *CreateHowtoV1Request) (*CreateHowtoV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHowtoV1 not implemented")
}
func (UnimplementedOcpHowtoApiServer) MultiCreateHowtoV1(context.Context, *MultiCreateHowtoV1Request) (*MultiCreateHowtoV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiCreateHowtoV1 not implemented")
}
func (UnimplementedOcpHowtoApiServer) UpdateHowtoV1(context.Context, *UpdateHowtoV1Request) (*UpdateHowtoV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHowtoV1 not implemented")
}
func (UnimplementedOcpHowtoApiServer) DescribeHowtoV1(context.Context, *DescribeHowtoV1Request) (*DescribeHowtoV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeHowtoV1 not implemented")
}
func (UnimplementedOcpHowtoApiServer) ListHowtosV1(context.Context, *ListHowtosV1Request) (*ListHowtosV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListHowtosV1 not implemented")
}
func (UnimplementedOcpHowtoApiServer) RemoveHowtoV1(context.Context, *RemoveHowtoV1Request) (*RemoveHowtoV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveHowtoV1 not implemented")
}
func (UnimplementedOcpHowtoApiServer) mustEmbedUnimplementedOcpHowtoApiServer() {}

// UnsafeOcpHowtoApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OcpHowtoApiServer will
// result in compilation errors.
type UnsafeOcpHowtoApiServer interface {
	mustEmbedUnimplementedOcpHowtoApiServer()
}

func RegisterOcpHowtoApiServer(s grpc.ServiceRegistrar, srv OcpHowtoApiServer) {
	s.RegisterService(&OcpHowtoApi_ServiceDesc, srv)
}

func _OcpHowtoApi_CreateHowtoV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHowtoV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpHowtoApiServer).CreateHowtoV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.howto.api.OcpHowtoApi/CreateHowtoV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpHowtoApiServer).CreateHowtoV1(ctx, req.(*CreateHowtoV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpHowtoApi_MultiCreateHowtoV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MultiCreateHowtoV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpHowtoApiServer).MultiCreateHowtoV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.howto.api.OcpHowtoApi/MultiCreateHowtoV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpHowtoApiServer).MultiCreateHowtoV1(ctx, req.(*MultiCreateHowtoV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpHowtoApi_UpdateHowtoV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHowtoV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpHowtoApiServer).UpdateHowtoV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.howto.api.OcpHowtoApi/UpdateHowtoV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpHowtoApiServer).UpdateHowtoV1(ctx, req.(*UpdateHowtoV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpHowtoApi_DescribeHowtoV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeHowtoV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpHowtoApiServer).DescribeHowtoV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.howto.api.OcpHowtoApi/DescribeHowtoV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpHowtoApiServer).DescribeHowtoV1(ctx, req.(*DescribeHowtoV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpHowtoApi_ListHowtosV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListHowtosV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpHowtoApiServer).ListHowtosV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.howto.api.OcpHowtoApi/ListHowtosV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpHowtoApiServer).ListHowtosV1(ctx, req.(*ListHowtosV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _OcpHowtoApi_RemoveHowtoV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveHowtoV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OcpHowtoApiServer).RemoveHowtoV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocp.howto.api.OcpHowtoApi/RemoveHowtoV1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OcpHowtoApiServer).RemoveHowtoV1(ctx, req.(*RemoveHowtoV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// OcpHowtoApi_ServiceDesc is the grpc.ServiceDesc for OcpHowtoApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OcpHowtoApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ocp.howto.api.OcpHowtoApi",
	HandlerType: (*OcpHowtoApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHowtoV1",
			Handler:    _OcpHowtoApi_CreateHowtoV1_Handler,
		},
		{
			MethodName: "MultiCreateHowtoV1",
			Handler:    _OcpHowtoApi_MultiCreateHowtoV1_Handler,
		},
		{
			MethodName: "UpdateHowtoV1",
			Handler:    _OcpHowtoApi_UpdateHowtoV1_Handler,
		},
		{
			MethodName: "DescribeHowtoV1",
			Handler:    _OcpHowtoApi_DescribeHowtoV1_Handler,
		},
		{
			MethodName: "ListHowtosV1",
			Handler:    _OcpHowtoApi_ListHowtosV1_Handler,
		},
		{
			MethodName: "RemoveHowtoV1",
			Handler:    _OcpHowtoApi_RemoveHowtoV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ocp-howto-api.proto",
}
