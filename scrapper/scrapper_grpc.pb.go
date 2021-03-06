// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: scrapper/scrapper.proto

package scrapperpc

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

// TikTokScrapperServiceClient is the client API for TikTokScrapperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TikTokScrapperServiceClient interface {
	GetTrackInfo(ctx context.Context, in *TrackInfoRequest, opts ...grpc.CallOption) (*TrackInfoResponse, error)
}

type tikTokScrapperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTikTokScrapperServiceClient(cc grpc.ClientConnInterface) TikTokScrapperServiceClient {
	return &tikTokScrapperServiceClient{cc}
}

func (c *tikTokScrapperServiceClient) GetTrackInfo(ctx context.Context, in *TrackInfoRequest, opts ...grpc.CallOption) (*TrackInfoResponse, error) {
	out := new(TrackInfoResponse)
	err := c.cc.Invoke(ctx, "/scrapper.TikTokScrapperService/GetTrackInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TikTokScrapperServiceServer is the server API for TikTokScrapperService service.
// All implementations must embed UnimplementedTikTokScrapperServiceServer
// for forward compatibility
type TikTokScrapperServiceServer interface {
	GetTrackInfo(context.Context, *TrackInfoRequest) (*TrackInfoResponse, error)
	mustEmbedUnimplementedTikTokScrapperServiceServer()
}

// UnimplementedTikTokScrapperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTikTokScrapperServiceServer struct {
}

func (UnimplementedTikTokScrapperServiceServer) GetTrackInfo(context.Context, *TrackInfoRequest) (*TrackInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrackInfo not implemented")
}
func (UnimplementedTikTokScrapperServiceServer) mustEmbedUnimplementedTikTokScrapperServiceServer() {}

// UnsafeTikTokScrapperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TikTokScrapperServiceServer will
// result in compilation errors.
type UnsafeTikTokScrapperServiceServer interface {
	mustEmbedUnimplementedTikTokScrapperServiceServer()
}

func RegisterTikTokScrapperServiceServer(s grpc.ServiceRegistrar, srv TikTokScrapperServiceServer) {
	s.RegisterService(&TikTokScrapperService_ServiceDesc, srv)
}

func _TikTokScrapperService_GetTrackInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrackInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TikTokScrapperServiceServer).GetTrackInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scrapper.TikTokScrapperService/GetTrackInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TikTokScrapperServiceServer).GetTrackInfo(ctx, req.(*TrackInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TikTokScrapperService_ServiceDesc is the grpc.ServiceDesc for TikTokScrapperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TikTokScrapperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scrapper.TikTokScrapperService",
	HandlerType: (*TikTokScrapperServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTrackInfo",
			Handler:    _TikTokScrapperService_GetTrackInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scrapper/scrapper.proto",
}
