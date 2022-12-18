// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: vacancy.proto

package pb

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

// VacancyServiceClient is the client API for VacancyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VacancyServiceClient interface {
	CreateVacancy(ctx context.Context, in *CreateVacancyReq, opts ...grpc.CallOption) (*CreateVacancyResponse, error)
	ReadVacancy(ctx context.Context, in *ReadVacancyRequest, opts ...grpc.CallOption) (*ReadVacancyResponse, error)
	UpdateVacancy(ctx context.Context, in *UpdateVacancyReq, opts ...grpc.CallOption) (*UpdateVacancyResponse, error)
	DeleteVacancy(ctx context.Context, in *DeleteVacancyRequest, opts ...grpc.CallOption) (*DeleteVacancyResponse, error)
	ListVacancies(ctx context.Context, in *ListVacancyRequest, opts ...grpc.CallOption) (*ListVacancyResponse, error)
}

type vacancyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVacancyServiceClient(cc grpc.ClientConnInterface) VacancyServiceClient {
	return &vacancyServiceClient{cc}
}

func (c *vacancyServiceClient) CreateVacancy(ctx context.Context, in *CreateVacancyReq, opts ...grpc.CallOption) (*CreateVacancyResponse, error) {
	out := new(CreateVacancyResponse)
	err := c.cc.Invoke(ctx, "/pb.VacancyService/CreateVacancy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) ReadVacancy(ctx context.Context, in *ReadVacancyRequest, opts ...grpc.CallOption) (*ReadVacancyResponse, error) {
	out := new(ReadVacancyResponse)
	err := c.cc.Invoke(ctx, "/pb.VacancyService/ReadVacancy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) UpdateVacancy(ctx context.Context, in *UpdateVacancyReq, opts ...grpc.CallOption) (*UpdateVacancyResponse, error) {
	out := new(UpdateVacancyResponse)
	err := c.cc.Invoke(ctx, "/pb.VacancyService/UpdateVacancy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) DeleteVacancy(ctx context.Context, in *DeleteVacancyRequest, opts ...grpc.CallOption) (*DeleteVacancyResponse, error) {
	out := new(DeleteVacancyResponse)
	err := c.cc.Invoke(ctx, "/pb.VacancyService/DeleteVacancy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacancyServiceClient) ListVacancies(ctx context.Context, in *ListVacancyRequest, opts ...grpc.CallOption) (*ListVacancyResponse, error) {
	out := new(ListVacancyResponse)
	err := c.cc.Invoke(ctx, "/pb.VacancyService/ListVacancies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VacancyServiceServer is the server API for VacancyService service.
// All implementations should embed UnimplementedVacancyServiceServer
// for forward compatibility
type VacancyServiceServer interface {
	CreateVacancy(context.Context, *CreateVacancyReq) (*CreateVacancyResponse, error)
	ReadVacancy(context.Context, *ReadVacancyRequest) (*ReadVacancyResponse, error)
	UpdateVacancy(context.Context, *UpdateVacancyReq) (*UpdateVacancyResponse, error)
	DeleteVacancy(context.Context, *DeleteVacancyRequest) (*DeleteVacancyResponse, error)
	ListVacancies(context.Context, *ListVacancyRequest) (*ListVacancyResponse, error)
}

// UnimplementedVacancyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedVacancyServiceServer struct {
}

func (UnimplementedVacancyServiceServer) CreateVacancy(context.Context, *CreateVacancyReq) (*CreateVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVacancy not implemented")
}
func (UnimplementedVacancyServiceServer) ReadVacancy(context.Context, *ReadVacancyRequest) (*ReadVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadVacancy not implemented")
}
func (UnimplementedVacancyServiceServer) UpdateVacancy(context.Context, *UpdateVacancyReq) (*UpdateVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVacancy not implemented")
}
func (UnimplementedVacancyServiceServer) DeleteVacancy(context.Context, *DeleteVacancyRequest) (*DeleteVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVacancy not implemented")
}
func (UnimplementedVacancyServiceServer) ListVacancies(context.Context, *ListVacancyRequest) (*ListVacancyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVacancies not implemented")
}

// UnsafeVacancyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VacancyServiceServer will
// result in compilation errors.
type UnsafeVacancyServiceServer interface {
	mustEmbedUnimplementedVacancyServiceServer()
}

func RegisterVacancyServiceServer(s grpc.ServiceRegistrar, srv VacancyServiceServer) {
	s.RegisterService(&VacancyService_ServiceDesc, srv)
}

func _VacancyService_CreateVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVacancyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).CreateVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VacancyService/CreateVacancy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).CreateVacancy(ctx, req.(*CreateVacancyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_ReadVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).ReadVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VacancyService/ReadVacancy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).ReadVacancy(ctx, req.(*ReadVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_UpdateVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateVacancyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).UpdateVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VacancyService/UpdateVacancy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).UpdateVacancy(ctx, req.(*UpdateVacancyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_DeleteVacancy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).DeleteVacancy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VacancyService/DeleteVacancy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).DeleteVacancy(ctx, req.(*DeleteVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacancyService_ListVacancies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVacancyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacancyServiceServer).ListVacancies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VacancyService/ListVacancies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacancyServiceServer).ListVacancies(ctx, req.(*ListVacancyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VacancyService_ServiceDesc is the grpc.ServiceDesc for VacancyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VacancyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.VacancyService",
	HandlerType: (*VacancyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVacancy",
			Handler:    _VacancyService_CreateVacancy_Handler,
		},
		{
			MethodName: "ReadVacancy",
			Handler:    _VacancyService_ReadVacancy_Handler,
		},
		{
			MethodName: "UpdateVacancy",
			Handler:    _VacancyService_UpdateVacancy_Handler,
		},
		{
			MethodName: "DeleteVacancy",
			Handler:    _VacancyService_DeleteVacancy_Handler,
		},
		{
			MethodName: "ListVacancies",
			Handler:    _VacancyService_ListVacancies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vacancy.proto",
}
