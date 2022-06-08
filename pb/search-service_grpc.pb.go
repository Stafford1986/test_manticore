// Code generated by protoc-gen-go-grpc_serv. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc_serv v1.2.0
// - protoc             v3.19.4
// source: search-service.proto

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

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchServiceClient interface {
	// services
	VacancySearch(ctx context.Context, in *VacancySearchEntity, opts ...grpc.CallOption) (*VacancySearchResponse, error)
	VacancyIndexUpdate(ctx context.Context, in *VacancyEntity, opts ...grpc.CallOption) (*CommentedResponse, error)
	VacancyIndexCreate(ctx context.Context, in *VacancyEntity, opts ...grpc.CallOption) (*CommentedResponse, error)
	ResumeSearch(ctx context.Context, in *ResumeSearchEntity, opts ...grpc.CallOption) (*ResumeSearchResponse, error)
	ResumeIndexUpdate(ctx context.Context, in *ResumeEntity, opts ...grpc.CallOption) (*CommentedResponse, error)
	ResumeIndexCreate(ctx context.Context, in *ResumeEntity, opts ...grpc.CallOption) (*CommentedResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) VacancySearch(ctx context.Context, in *VacancySearchEntity, opts ...grpc.CallOption) (*VacancySearchResponse, error) {
	out := new(VacancySearchResponse)
	err := c.cc.Invoke(ctx, "/search.SearchService/VacancySearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) VacancyIndexUpdate(ctx context.Context, in *VacancyEntity, opts ...grpc.CallOption) (*CommentedResponse, error) {
	out := new(CommentedResponse)
	err := c.cc.Invoke(ctx, "/search.SearchService/VacancyIndexUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) VacancyIndexCreate(ctx context.Context, in *VacancyEntity, opts ...grpc.CallOption) (*CommentedResponse, error) {
	out := new(CommentedResponse)
	err := c.cc.Invoke(ctx, "/search.SearchService/VacancyIndexCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) ResumeSearch(ctx context.Context, in *ResumeSearchEntity, opts ...grpc.CallOption) (*ResumeSearchResponse, error) {
	out := new(ResumeSearchResponse)
	err := c.cc.Invoke(ctx, "/search.SearchService/ResumeSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) ResumeIndexUpdate(ctx context.Context, in *ResumeEntity, opts ...grpc.CallOption) (*CommentedResponse, error) {
	out := new(CommentedResponse)
	err := c.cc.Invoke(ctx, "/search.SearchService/ResumeIndexUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *searchServiceClient) ResumeIndexCreate(ctx context.Context, in *ResumeEntity, opts ...grpc.CallOption) (*CommentedResponse, error) {
	out := new(CommentedResponse)
	err := c.cc.Invoke(ctx, "/search.SearchService/ResumeIndexCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
// All implementations should embed UnimplementedSearchServiceServer
// for forward compatibility
type SearchServiceServer interface {
	// services
	VacancySearch(context.Context, *VacancySearchEntity) (*VacancySearchResponse, error)
	VacancyIndexUpdate(context.Context, *VacancyEntity) (*CommentedResponse, error)
	VacancyIndexCreate(context.Context, *VacancyEntity) (*CommentedResponse, error)
	ResumeSearch(context.Context, *ResumeSearchEntity) (*ResumeSearchResponse, error)
	ResumeIndexUpdate(context.Context, *ResumeEntity) (*CommentedResponse, error)
	ResumeIndexCreate(context.Context, *ResumeEntity) (*CommentedResponse, error)
}

// UnimplementedSearchServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (UnimplementedSearchServiceServer) VacancySearch(context.Context, *VacancySearchEntity) (*VacancySearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VacancySearch not implemented")
}
func (UnimplementedSearchServiceServer) VacancyIndexUpdate(context.Context, *VacancyEntity) (*CommentedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VacancyIndexUpdate not implemented")
}
func (UnimplementedSearchServiceServer) VacancyIndexCreate(context.Context, *VacancyEntity) (*CommentedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VacancyIndexCreate not implemented")
}
func (UnimplementedSearchServiceServer) ResumeSearch(context.Context, *ResumeSearchEntity) (*ResumeSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeSearch not implemented")
}
func (UnimplementedSearchServiceServer) ResumeIndexUpdate(context.Context, *ResumeEntity) (*CommentedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeIndexUpdate not implemented")
}
func (UnimplementedSearchServiceServer) ResumeIndexCreate(context.Context, *ResumeEntity) (*CommentedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeIndexCreate not implemented")
}

// UnsafeSearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchServiceServer will
// result in compilation errors.
type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_VacancySearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacancySearchEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).VacancySearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/VacancySearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).VacancySearch(ctx, req.(*VacancySearchEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_VacancyIndexUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacancyEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).VacancyIndexUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/VacancyIndexUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).VacancyIndexUpdate(ctx, req.(*VacancyEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_VacancyIndexCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VacancyEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).VacancyIndexCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/VacancyIndexCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).VacancyIndexCreate(ctx, req.(*VacancyEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_ResumeSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeSearchEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).ResumeSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/ResumeSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).ResumeSearch(ctx, req.(*ResumeSearchEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_ResumeIndexUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).ResumeIndexUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/ResumeIndexUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).ResumeIndexUpdate(ctx, req.(*ResumeEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _SearchService_ResumeIndexCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).ResumeIndexCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/search.SearchService/ResumeIndexCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).ResumeIndexCreate(ctx, req.(*ResumeEntity))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchService_ServiceDesc is the grpc.ServiceDesc for SearchService service.
// It's only intended for direct use with grpc_serv.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VacancySearch",
			Handler:    _SearchService_VacancySearch_Handler,
		},
		{
			MethodName: "VacancyIndexUpdate",
			Handler:    _SearchService_VacancyIndexUpdate_Handler,
		},
		{
			MethodName: "VacancyIndexCreate",
			Handler:    _SearchService_VacancyIndexCreate_Handler,
		},
		{
			MethodName: "ResumeSearch",
			Handler:    _SearchService_ResumeSearch_Handler,
		},
		{
			MethodName: "ResumeIndexUpdate",
			Handler:    _SearchService_ResumeIndexUpdate_Handler,
		},
		{
			MethodName: "ResumeIndexCreate",
			Handler:    _SearchService_ResumeIndexCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search-service.proto",
}