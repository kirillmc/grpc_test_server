// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: program_v3.proto

package program_v3

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

// ProgramV3Client is the client API for ProgramV3 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProgramV3Client interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*TrainPrograms, error)
	Create(ctx context.Context, in *TrainPrograms, opts ...grpc.CallOption) (*Response, error)
	Update(ctx context.Context, in *TrainPrograms, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Response, error)
}

type programV3Client struct {
	cc grpc.ClientConnInterface
}

func NewProgramV3Client(cc grpc.ClientConnInterface) ProgramV3Client {
	return &programV3Client{cc}
}

func (c *programV3Client) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*TrainPrograms, error) {
	out := new(TrainPrograms)
	err := c.cc.Invoke(ctx, "/program_v3.ProgramV3/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programV3Client) Create(ctx context.Context, in *TrainPrograms, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/program_v3.ProgramV3/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programV3Client) Update(ctx context.Context, in *TrainPrograms, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/program_v3.ProgramV3/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *programV3Client) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/program_v3.ProgramV3/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProgramV3Server is the server API for ProgramV3 service.
// All implementations must embed UnimplementedProgramV3Server
// for forward compatibility
type ProgramV3Server interface {
	Get(context.Context, *GetRequest) (*TrainPrograms, error)
	Create(context.Context, *TrainPrograms) (*Response, error)
	Update(context.Context, *TrainPrograms) (*Response, error)
	Delete(context.Context, *DeleteRequest) (*Response, error)
	mustEmbedUnimplementedProgramV3Server()
}

// UnimplementedProgramV3Server must be embedded to have forward compatible implementations.
type UnimplementedProgramV3Server struct {
}

func (UnimplementedProgramV3Server) Get(context.Context, *GetRequest) (*TrainPrograms, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedProgramV3Server) Create(context.Context, *TrainPrograms) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProgramV3Server) Update(context.Context, *TrainPrograms) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProgramV3Server) Delete(context.Context, *DeleteRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProgramV3Server) mustEmbedUnimplementedProgramV3Server() {}

// UnsafeProgramV3Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProgramV3Server will
// result in compilation errors.
type UnsafeProgramV3Server interface {
	mustEmbedUnimplementedProgramV3Server()
}

func RegisterProgramV3Server(s grpc.ServiceRegistrar, srv ProgramV3Server) {
	s.RegisterService(&ProgramV3_ServiceDesc, srv)
}

func _ProgramV3_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramV3Server).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/program_v3.ProgramV3/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramV3Server).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgramV3_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainPrograms)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramV3Server).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/program_v3.ProgramV3/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramV3Server).Create(ctx, req.(*TrainPrograms))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgramV3_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainPrograms)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramV3Server).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/program_v3.ProgramV3/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramV3Server).Update(ctx, req.(*TrainPrograms))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgramV3_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgramV3Server).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/program_v3.ProgramV3/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgramV3Server).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProgramV3_ServiceDesc is the grpc.ServiceDesc for ProgramV3 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProgramV3_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "program_v3.ProgramV3",
	HandlerType: (*ProgramV3Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ProgramV3_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ProgramV3_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProgramV3_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProgramV3_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "program_v3.proto",
}
