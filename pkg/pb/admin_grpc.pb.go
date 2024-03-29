// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pkg/pb/admin.proto

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

const (
	AdminService_Feeds_FullMethodName               = "/admin.AdminService/Feeds"
	AdminService_PostDetails_FullMethodName         = "/admin.AdminService/PostDetails"
	AdminService_DeletePost_FullMethodName          = "/admin.AdminService/DeletePost"
	AdminService_CampaignRequestList_FullMethodName = "/admin.AdminService/CampaignRequestList"
	AdminService_CampaignDetails_FullMethodName     = "/admin.AdminService/CampaignDetails"
	AdminService_ApproveCampaign_FullMethodName     = "/admin.AdminService/ApproveCampaign"
	AdminService_RejectCampaign_FullMethodName      = "/admin.AdminService/RejectCampaign"
	AdminService_ReportedList_FullMethodName        = "/admin.AdminService/ReportedList"
	AdminService_ReportDetails_FullMethodName       = "/admin.AdminService/ReportDetails"
	AdminService_DeleteReport_FullMethodName        = "/admin.AdminService/DeleteReport"
	AdminService_CategoryList_FullMethodName        = "/admin.AdminService/CategoryList"
	AdminService_CategoryPosts_FullMethodName       = "/admin.AdminService/CategoryPosts"
	AdminService_NewCategory_FullMethodName         = "/admin.AdminService/NewCategory"
	AdminService_DeleteCategory_FullMethodName      = "/admin.AdminService/DeleteCategory"
	AdminService_AdminDashboard_FullMethodName      = "/admin.AdminService/AdminDashboard"
	AdminService_PostStats_FullMethodName           = "/admin.AdminService/PostStats"
	AdminService_CategoryStats_FullMethodName       = "/admin.AdminService/CategoryStats"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	Feeds(ctx context.Context, in *FeedsRequest, opts ...grpc.CallOption) (*FeedsResponse, error)
	PostDetails(ctx context.Context, in *PostDetailsRequest, opts ...grpc.CallOption) (*PostDetailsResponse, error)
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
	CampaignRequestList(ctx context.Context, in *CampaignRequestListRequest, opts ...grpc.CallOption) (*CampaignRequestListResponse, error)
	CampaignDetails(ctx context.Context, in *CampaignDetailsRequest, opts ...grpc.CallOption) (*CampaignDetailsResponse, error)
	ApproveCampaign(ctx context.Context, in *ApproveCampaignRequest, opts ...grpc.CallOption) (*ApproveCampaignResponse, error)
	RejectCampaign(ctx context.Context, in *RejectCampaignRequest, opts ...grpc.CallOption) (*RejectCampaignResponse, error)
	// //////////////////////////////////////////////////////////////////////////////
	ReportedList(ctx context.Context, in *ReportedListRequest, opts ...grpc.CallOption) (*ReportedListResponse, error)
	ReportDetails(ctx context.Context, in *ReportDetailsRequest, opts ...grpc.CallOption) (*ReportDetailsResponse, error)
	DeleteReport(ctx context.Context, in *DeleteReportRequest, opts ...grpc.CallOption) (*DeleteReportResponse, error)
	CategoryList(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*CategoryListResponse, error)
	CategoryPosts(ctx context.Context, in *CategoryPostsRequest, opts ...grpc.CallOption) (*CategoryPostsResponse, error)
	NewCategory(ctx context.Context, in *NewCategoryRequest, opts ...grpc.CallOption) (*NewCategoryResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error)
	AdminDashboard(ctx context.Context, in *AdminDashboardRequest, opts ...grpc.CallOption) (*AdminDashboardResponse, error)
	PostStats(ctx context.Context, in *PostStatsRequest, opts ...grpc.CallOption) (*PostStatsResponse, error)
	CategoryStats(ctx context.Context, in *CategoryStatsRequest, opts ...grpc.CallOption) (*CategoryStatsResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) Feeds(ctx context.Context, in *FeedsRequest, opts ...grpc.CallOption) (*FeedsResponse, error) {
	out := new(FeedsResponse)
	err := c.cc.Invoke(ctx, AdminService_Feeds_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) PostDetails(ctx context.Context, in *PostDetailsRequest, opts ...grpc.CallOption) (*PostDetailsResponse, error) {
	out := new(PostDetailsResponse)
	err := c.cc.Invoke(ctx, AdminService_PostDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, AdminService_DeletePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CampaignRequestList(ctx context.Context, in *CampaignRequestListRequest, opts ...grpc.CallOption) (*CampaignRequestListResponse, error) {
	out := new(CampaignRequestListResponse)
	err := c.cc.Invoke(ctx, AdminService_CampaignRequestList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CampaignDetails(ctx context.Context, in *CampaignDetailsRequest, opts ...grpc.CallOption) (*CampaignDetailsResponse, error) {
	out := new(CampaignDetailsResponse)
	err := c.cc.Invoke(ctx, AdminService_CampaignDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ApproveCampaign(ctx context.Context, in *ApproveCampaignRequest, opts ...grpc.CallOption) (*ApproveCampaignResponse, error) {
	out := new(ApproveCampaignResponse)
	err := c.cc.Invoke(ctx, AdminService_ApproveCampaign_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) RejectCampaign(ctx context.Context, in *RejectCampaignRequest, opts ...grpc.CallOption) (*RejectCampaignResponse, error) {
	out := new(RejectCampaignResponse)
	err := c.cc.Invoke(ctx, AdminService_RejectCampaign_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ReportedList(ctx context.Context, in *ReportedListRequest, opts ...grpc.CallOption) (*ReportedListResponse, error) {
	out := new(ReportedListResponse)
	err := c.cc.Invoke(ctx, AdminService_ReportedList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) ReportDetails(ctx context.Context, in *ReportDetailsRequest, opts ...grpc.CallOption) (*ReportDetailsResponse, error) {
	out := new(ReportDetailsResponse)
	err := c.cc.Invoke(ctx, AdminService_ReportDetails_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeleteReport(ctx context.Context, in *DeleteReportRequest, opts ...grpc.CallOption) (*DeleteReportResponse, error) {
	out := new(DeleteReportResponse)
	err := c.cc.Invoke(ctx, AdminService_DeleteReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CategoryList(ctx context.Context, in *CategoryListRequest, opts ...grpc.CallOption) (*CategoryListResponse, error) {
	out := new(CategoryListResponse)
	err := c.cc.Invoke(ctx, AdminService_CategoryList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CategoryPosts(ctx context.Context, in *CategoryPostsRequest, opts ...grpc.CallOption) (*CategoryPostsResponse, error) {
	out := new(CategoryPostsResponse)
	err := c.cc.Invoke(ctx, AdminService_CategoryPosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) NewCategory(ctx context.Context, in *NewCategoryRequest, opts ...grpc.CallOption) (*NewCategoryResponse, error) {
	out := new(NewCategoryResponse)
	err := c.cc.Invoke(ctx, AdminService_NewCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error) {
	out := new(DeleteCategoryResponse)
	err := c.cc.Invoke(ctx, AdminService_DeleteCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminDashboard(ctx context.Context, in *AdminDashboardRequest, opts ...grpc.CallOption) (*AdminDashboardResponse, error) {
	out := new(AdminDashboardResponse)
	err := c.cc.Invoke(ctx, AdminService_AdminDashboard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) PostStats(ctx context.Context, in *PostStatsRequest, opts ...grpc.CallOption) (*PostStatsResponse, error) {
	out := new(PostStatsResponse)
	err := c.cc.Invoke(ctx, AdminService_PostStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) CategoryStats(ctx context.Context, in *CategoryStatsRequest, opts ...grpc.CallOption) (*CategoryStatsResponse, error) {
	out := new(CategoryStatsResponse)
	err := c.cc.Invoke(ctx, AdminService_CategoryStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	Feeds(context.Context, *FeedsRequest) (*FeedsResponse, error)
	PostDetails(context.Context, *PostDetailsRequest) (*PostDetailsResponse, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	CampaignRequestList(context.Context, *CampaignRequestListRequest) (*CampaignRequestListResponse, error)
	CampaignDetails(context.Context, *CampaignDetailsRequest) (*CampaignDetailsResponse, error)
	ApproveCampaign(context.Context, *ApproveCampaignRequest) (*ApproveCampaignResponse, error)
	RejectCampaign(context.Context, *RejectCampaignRequest) (*RejectCampaignResponse, error)
	// //////////////////////////////////////////////////////////////////////////////
	ReportedList(context.Context, *ReportedListRequest) (*ReportedListResponse, error)
	ReportDetails(context.Context, *ReportDetailsRequest) (*ReportDetailsResponse, error)
	DeleteReport(context.Context, *DeleteReportRequest) (*DeleteReportResponse, error)
	CategoryList(context.Context, *CategoryListRequest) (*CategoryListResponse, error)
	CategoryPosts(context.Context, *CategoryPostsRequest) (*CategoryPostsResponse, error)
	NewCategory(context.Context, *NewCategoryRequest) (*NewCategoryResponse, error)
	DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error)
	AdminDashboard(context.Context, *AdminDashboardRequest) (*AdminDashboardResponse, error)
	PostStats(context.Context, *PostStatsRequest) (*PostStatsResponse, error)
	CategoryStats(context.Context, *CategoryStatsRequest) (*CategoryStatsResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) Feeds(context.Context, *FeedsRequest) (*FeedsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Feeds not implemented")
}
func (UnimplementedAdminServiceServer) PostDetails(context.Context, *PostDetailsRequest) (*PostDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostDetails not implemented")
}
func (UnimplementedAdminServiceServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedAdminServiceServer) CampaignRequestList(context.Context, *CampaignRequestListRequest) (*CampaignRequestListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CampaignRequestList not implemented")
}
func (UnimplementedAdminServiceServer) CampaignDetails(context.Context, *CampaignDetailsRequest) (*CampaignDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CampaignDetails not implemented")
}
func (UnimplementedAdminServiceServer) ApproveCampaign(context.Context, *ApproveCampaignRequest) (*ApproveCampaignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveCampaign not implemented")
}
func (UnimplementedAdminServiceServer) RejectCampaign(context.Context, *RejectCampaignRequest) (*RejectCampaignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectCampaign not implemented")
}
func (UnimplementedAdminServiceServer) ReportedList(context.Context, *ReportedListRequest) (*ReportedListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportedList not implemented")
}
func (UnimplementedAdminServiceServer) ReportDetails(context.Context, *ReportDetailsRequest) (*ReportDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportDetails not implemented")
}
func (UnimplementedAdminServiceServer) DeleteReport(context.Context, *DeleteReportRequest) (*DeleteReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReport not implemented")
}
func (UnimplementedAdminServiceServer) CategoryList(context.Context, *CategoryListRequest) (*CategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryList not implemented")
}
func (UnimplementedAdminServiceServer) CategoryPosts(context.Context, *CategoryPostsRequest) (*CategoryPostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryPosts not implemented")
}
func (UnimplementedAdminServiceServer) NewCategory(context.Context, *NewCategoryRequest) (*NewCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewCategory not implemented")
}
func (UnimplementedAdminServiceServer) DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedAdminServiceServer) AdminDashboard(context.Context, *AdminDashboardRequest) (*AdminDashboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminDashboard not implemented")
}
func (UnimplementedAdminServiceServer) PostStats(context.Context, *PostStatsRequest) (*PostStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostStats not implemented")
}
func (UnimplementedAdminServiceServer) CategoryStats(context.Context, *CategoryStatsRequest) (*CategoryStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryStats not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_Feeds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeedsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).Feeds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_Feeds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).Feeds(ctx, req.(*FeedsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_PostDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).PostDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_PostDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).PostDetails(ctx, req.(*PostDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CampaignRequestList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CampaignRequestListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CampaignRequestList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_CampaignRequestList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CampaignRequestList(ctx, req.(*CampaignRequestListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CampaignDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CampaignDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CampaignDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_CampaignDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CampaignDetails(ctx, req.(*CampaignDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ApproveCampaign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApproveCampaignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ApproveCampaign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_ApproveCampaign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ApproveCampaign(ctx, req.(*ApproveCampaignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_RejectCampaign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RejectCampaignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).RejectCampaign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_RejectCampaign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).RejectCampaign(ctx, req.(*RejectCampaignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ReportedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportedListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ReportedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_ReportedList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ReportedList(ctx, req.(*ReportedListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_ReportDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ReportDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_ReportDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ReportDetails(ctx, req.(*ReportDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeleteReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeleteReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_DeleteReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeleteReport(ctx, req.(*DeleteReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CategoryList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CategoryList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_CategoryList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CategoryList(ctx, req.(*CategoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CategoryPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CategoryPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_CategoryPosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CategoryPosts(ctx, req.(*CategoryPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_NewCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).NewCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_NewCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).NewCategory(ctx, req.(*NewCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeleteCategory(ctx, req.(*DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AdminDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminDashboard(ctx, req.(*AdminDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_PostStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).PostStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_PostStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).PostStats(ctx, req.(*PostStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_CategoryStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).CategoryStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_CategoryStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).CategoryStats(ctx, req.(*CategoryStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Feeds",
			Handler:    _AdminService_Feeds_Handler,
		},
		{
			MethodName: "PostDetails",
			Handler:    _AdminService_PostDetails_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _AdminService_DeletePost_Handler,
		},
		{
			MethodName: "CampaignRequestList",
			Handler:    _AdminService_CampaignRequestList_Handler,
		},
		{
			MethodName: "CampaignDetails",
			Handler:    _AdminService_CampaignDetails_Handler,
		},
		{
			MethodName: "ApproveCampaign",
			Handler:    _AdminService_ApproveCampaign_Handler,
		},
		{
			MethodName: "RejectCampaign",
			Handler:    _AdminService_RejectCampaign_Handler,
		},
		{
			MethodName: "ReportedList",
			Handler:    _AdminService_ReportedList_Handler,
		},
		{
			MethodName: "ReportDetails",
			Handler:    _AdminService_ReportDetails_Handler,
		},
		{
			MethodName: "DeleteReport",
			Handler:    _AdminService_DeleteReport_Handler,
		},
		{
			MethodName: "CategoryList",
			Handler:    _AdminService_CategoryList_Handler,
		},
		{
			MethodName: "CategoryPosts",
			Handler:    _AdminService_CategoryPosts_Handler,
		},
		{
			MethodName: "NewCategory",
			Handler:    _AdminService_NewCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _AdminService_DeleteCategory_Handler,
		},
		{
			MethodName: "AdminDashboard",
			Handler:    _AdminService_AdminDashboard_Handler,
		},
		{
			MethodName: "PostStats",
			Handler:    _AdminService_PostStats_Handler,
		},
		{
			MethodName: "CategoryStats",
			Handler:    _AdminService_CategoryStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/admin.proto",
}
