// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.4
// source: protobuf.proto

package __

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
	Bidireccional(ctx context.Context, opts ...grpc.CallOption) (CodeExecutionService_BidireccionalClient, error)
}

type codeExecutionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeExecutionServiceClient(cc grpc.ClientConnInterface) CodeExecutionServiceClient {
	return &codeExecutionServiceClient{cc}
}

func (c *codeExecutionServiceClient) Bidireccional(ctx context.Context, opts ...grpc.CallOption) (CodeExecutionService_BidireccionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &CodeExecutionService_ServiceDesc.Streams[0], "/grpc.CodeExecutionService/Bidireccional", opts...)
	if err != nil {
		return nil, err
	}
	x := &codeExecutionServiceBidireccionalClient{stream}
	return x, nil
}

type CodeExecutionService_BidireccionalClient interface {
	Send(*CodeRequest) error
	Recv() (*CodeResponse, error)
	grpc.ClientStream
}

type codeExecutionServiceBidireccionalClient struct {
	grpc.ClientStream
}

func (x *codeExecutionServiceBidireccionalClient) Send(m *CodeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *codeExecutionServiceBidireccionalClient) Recv() (*CodeResponse, error) {
	m := new(CodeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CodeExecutionServiceServer is the server API for CodeExecutionService service.
// All implementations must embed UnimplementedCodeExecutionServiceServer
// for forward compatibility
type CodeExecutionServiceServer interface {
	Bidireccional(CodeExecutionService_BidireccionalServer) error
	mustEmbedUnimplementedCodeExecutionServiceServer()
}

// UnimplementedCodeExecutionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCodeExecutionServiceServer struct {
}

func (UnimplementedCodeExecutionServiceServer) Bidireccional(CodeExecutionService_BidireccionalServer) error {
	return status.Errorf(codes.Unimplemented, "method Bidireccional not implemented")
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

func _CodeExecutionService_Bidireccional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CodeExecutionServiceServer).Bidireccional(&codeExecutionServiceBidireccionalServer{stream})
}

type CodeExecutionService_BidireccionalServer interface {
	Send(*CodeResponse) error
	Recv() (*CodeRequest, error)
	grpc.ServerStream
}

type codeExecutionServiceBidireccionalServer struct {
	grpc.ServerStream
}

func (x *codeExecutionServiceBidireccionalServer) Send(m *CodeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *codeExecutionServiceBidireccionalServer) Recv() (*CodeRequest, error) {
	m := new(CodeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CodeExecutionService_ServiceDesc is the grpc.ServiceDesc for CodeExecutionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeExecutionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.CodeExecutionService",
	HandlerType: (*CodeExecutionServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Bidireccional",
			Handler:       _CodeExecutionService_Bidireccional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "protobuf.proto",
}
