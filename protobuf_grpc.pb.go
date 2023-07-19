// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: protobuf.proto

package grpc

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

// CodeExecutionServiceClient is the client API for CodeExecutionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CodeExecutionServiceClient interface {
	ExecuteCode(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeResponse, error)
}

type codeExecutionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeExecutionServiceClient(cc grpc.ClientConnInterface) CodeExecutionServiceClient {
	return &codeExecutionServiceClient{cc}
}

func (c *codeExecutionServiceClient) ExecuteCode(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*CodeResponse, error) {
	out := new(CodeResponse)
	err := c.cc.Invoke(ctx, "/grpc.CodeExecutionService/ExecuteCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeExecutionServiceServer is the server API for CodeExecutionService service.
// All implementations must embed UnimplementedCodeExecutionServiceServer
// for forward compatibility
type CodeExecutionServiceServer interface {
	ExecuteCode(context.Context, *CodeRequest) (*CodeResponse, error)
	mustEmbedUnimplementedCodeExecutionServiceServer()
}

// UnimplementedCodeExecutionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCodeExecutionServiceServer struct {
}

func (UnimplementedCodeExecutionServiceServer) ExecuteCode(context.Context, *CodeRequest) (*CodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCode not implemented")
}
func (UnimplementedCodeExecutionServiceServer) mustEmbedUnimplementedCodeExecutionServiceServer() {}

// UnsafeCodeExecutionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodeExecutionServiceServer will
// result in compilation errors.
type UnsafeCodeExecutionServiceServer interface {
	mustEmbedUnimplementedCodeExecutionServiceServer()
}

func RegisterCodeExecutionServiceServer(s grpc.ServiceRegistrar, srv CodeExecutionServiceServer) {
	s.RegisterService(&CodeExecutionService_ServiceDesc, srv)
}

func _CodeExecutionService_ExecuteCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeExecutionServiceServer).ExecuteCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.CodeExecutionService/ExecuteCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeExecutionServiceServer).ExecuteCode(ctx, req.(*CodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CodeExecutionService_ServiceDesc is the grpc.ServiceDesc for CodeExecutionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeExecutionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.CodeExecutionService",
	HandlerType: (*CodeExecutionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteCode",
			Handler:    _CodeExecutionService_ExecuteCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf.proto",
}
