// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: songService.proto

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

// SongServiceClient is the client API for SongService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SongServiceClient interface {
	CreateSong(ctx context.Context, opts ...grpc.CallOption) (SongService_CreateSongClient, error)
	UpdateSong(ctx context.Context, in *UpdatedSong, opts ...grpc.CallOption) (*Response, error)
	GetSongById(ctx context.Context, in *SongId, opts ...grpc.CallOption) (*Song, error)
	SongFilter(ctx context.Context, in *SongMetadata, opts ...grpc.CallOption) (SongService_SongFilterClient, error)
	RemoveSongById(ctx context.Context, in *SongId, opts ...grpc.CallOption) (*Response, error)
}

type songServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSongServiceClient(cc grpc.ClientConnInterface) SongServiceClient {
	return &songServiceClient{cc}
}

func (c *songServiceClient) CreateSong(ctx context.Context, opts ...grpc.CallOption) (SongService_CreateSongClient, error) {
	stream, err := c.cc.NewStream(ctx, &SongService_ServiceDesc.Streams[0], "/songs.SongService/CreateSong", opts...)
	if err != nil {
		return nil, err
	}
	x := &songServiceCreateSongClient{stream}
	return x, nil
}

type SongService_CreateSongClient interface {
	Send(*RawSong) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type songServiceCreateSongClient struct {
	grpc.ClientStream
}

func (x *songServiceCreateSongClient) Send(m *RawSong) error {
	return x.ClientStream.SendMsg(m)
}

func (x *songServiceCreateSongClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *songServiceClient) UpdateSong(ctx context.Context, in *UpdatedSong, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/songs.SongService/UpdateSong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) GetSongById(ctx context.Context, in *SongId, opts ...grpc.CallOption) (*Song, error) {
	out := new(Song)
	err := c.cc.Invoke(ctx, "/songs.SongService/GetSongById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *songServiceClient) SongFilter(ctx context.Context, in *SongMetadata, opts ...grpc.CallOption) (SongService_SongFilterClient, error) {
	stream, err := c.cc.NewStream(ctx, &SongService_ServiceDesc.Streams[1], "/songs.SongService/SongFilter", opts...)
	if err != nil {
		return nil, err
	}
	x := &songServiceSongFilterClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SongService_SongFilterClient interface {
	Recv() (*Song, error)
	grpc.ClientStream
}

type songServiceSongFilterClient struct {
	grpc.ClientStream
}

func (x *songServiceSongFilterClient) Recv() (*Song, error) {
	m := new(Song)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *songServiceClient) RemoveSongById(ctx context.Context, in *SongId, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/songs.SongService/RemoveSongById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SongServiceServer is the server API for SongService service.
// All implementations must embed UnimplementedSongServiceServer
// for forward compatibility
type SongServiceServer interface {
	CreateSong(SongService_CreateSongServer) error
	UpdateSong(context.Context, *UpdatedSong) (*Response, error)
	GetSongById(context.Context, *SongId) (*Song, error)
	SongFilter(*SongMetadata, SongService_SongFilterServer) error
	RemoveSongById(context.Context, *SongId) (*Response, error)
	mustEmbedUnimplementedSongServiceServer()
}

// UnimplementedSongServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSongServiceServer struct {
}

func (UnimplementedSongServiceServer) CreateSong(SongService_CreateSongServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateSong not implemented")
}
func (UnimplementedSongServiceServer) UpdateSong(context.Context, *UpdatedSong) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSong not implemented")
}
func (UnimplementedSongServiceServer) GetSongById(context.Context, *SongId) (*Song, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSongById not implemented")
}
func (UnimplementedSongServiceServer) SongFilter(*SongMetadata, SongService_SongFilterServer) error {
	return status.Errorf(codes.Unimplemented, "method SongFilter not implemented")
}
func (UnimplementedSongServiceServer) RemoveSongById(context.Context, *SongId) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSongById not implemented")
}
func (UnimplementedSongServiceServer) mustEmbedUnimplementedSongServiceServer() {}

// UnsafeSongServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SongServiceServer will
// result in compilation errors.
type UnsafeSongServiceServer interface {
	mustEmbedUnimplementedSongServiceServer()
}

func RegisterSongServiceServer(s grpc.ServiceRegistrar, srv SongServiceServer) {
	s.RegisterService(&SongService_ServiceDesc, srv)
}

func _SongService_CreateSong_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SongServiceServer).CreateSong(&songServiceCreateSongServer{stream})
}

type SongService_CreateSongServer interface {
	SendAndClose(*Response) error
	Recv() (*RawSong, error)
	grpc.ServerStream
}

type songServiceCreateSongServer struct {
	grpc.ServerStream
}

func (x *songServiceCreateSongServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *songServiceCreateSongServer) Recv() (*RawSong, error) {
	m := new(RawSong)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SongService_UpdateSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatedSong)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).UpdateSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/songs.SongService/UpdateSong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).UpdateSong(ctx, req.(*UpdatedSong))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_GetSongById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).GetSongById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/songs.SongService/GetSongById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).GetSongById(ctx, req.(*SongId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SongService_SongFilter_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SongMetadata)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SongServiceServer).SongFilter(m, &songServiceSongFilterServer{stream})
}

type SongService_SongFilterServer interface {
	Send(*Song) error
	grpc.ServerStream
}

type songServiceSongFilterServer struct {
	grpc.ServerStream
}

func (x *songServiceSongFilterServer) Send(m *Song) error {
	return x.ServerStream.SendMsg(m)
}

func _SongService_RemoveSongById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SongServiceServer).RemoveSongById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/songs.SongService/RemoveSongById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SongServiceServer).RemoveSongById(ctx, req.(*SongId))
	}
	return interceptor(ctx, in, info, handler)
}

// SongService_ServiceDesc is the grpc.ServiceDesc for SongService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SongService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "songs.SongService",
	HandlerType: (*SongServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateSong",
			Handler:    _SongService_UpdateSong_Handler,
		},
		{
			MethodName: "GetSongById",
			Handler:    _SongService_GetSongById_Handler,
		},
		{
			MethodName: "RemoveSongById",
			Handler:    _SongService_RemoveSongById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateSong",
			Handler:       _SongService_CreateSong_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SongFilter",
			Handler:       _SongService_SongFilter_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "songService.proto",
}
