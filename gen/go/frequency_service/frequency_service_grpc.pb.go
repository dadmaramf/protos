// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.1
// source: frequency_service/frequency_service.proto

package frequency_servicev1

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
	StreamingService_NewConnect_FullMethodName = "/rand.StreamingService/NewConnect"
)

// StreamingServiceClient is the client API for StreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingServiceClient interface {
	NewConnect(ctx context.Context, in *FrequencyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FrequencyResponse], error)
}

type streamingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingServiceClient(cc grpc.ClientConnInterface) StreamingServiceClient {
	return &streamingServiceClient{cc}
}

func (c *streamingServiceClient) NewConnect(ctx context.Context, in *FrequencyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[FrequencyResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &StreamingService_ServiceDesc.Streams[0], StreamingService_NewConnect_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[FrequencyRequest, FrequencyResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StreamingService_NewConnectClient = grpc.ServerStreamingClient[FrequencyResponse]

// StreamingServiceServer is the server API for StreamingService service.
// All implementations must embed UnimplementedStreamingServiceServer
// for forward compatibility.
type StreamingServiceServer interface {
	NewConnect(*FrequencyRequest, grpc.ServerStreamingServer[FrequencyResponse]) error
	mustEmbedUnimplementedStreamingServiceServer()
}

// UnimplementedStreamingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStreamingServiceServer struct{}

func (UnimplementedStreamingServiceServer) NewConnect(*FrequencyRequest, grpc.ServerStreamingServer[FrequencyResponse]) error {
	return status.Errorf(codes.Unimplemented, "method NewConnect not implemented")
}
func (UnimplementedStreamingServiceServer) mustEmbedUnimplementedStreamingServiceServer() {}
func (UnimplementedStreamingServiceServer) testEmbeddedByValue()                          {}

// UnsafeStreamingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingServiceServer will
// result in compilation errors.
type UnsafeStreamingServiceServer interface {
	mustEmbedUnimplementedStreamingServiceServer()
}

func RegisterStreamingServiceServer(s grpc.ServiceRegistrar, srv StreamingServiceServer) {
	// If the following call pancis, it indicates UnimplementedStreamingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StreamingService_ServiceDesc, srv)
}

func _StreamingService_NewConnect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FrequencyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamingServiceServer).NewConnect(m, &grpc.GenericServerStream[FrequencyRequest, FrequencyResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type StreamingService_NewConnectServer = grpc.ServerStreamingServer[FrequencyResponse]

// StreamingService_ServiceDesc is the grpc.ServiceDesc for StreamingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rand.StreamingService",
	HandlerType: (*StreamingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NewConnect",
			Handler:       _StreamingService_NewConnect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "frequency_service/frequency_service.proto",
}
