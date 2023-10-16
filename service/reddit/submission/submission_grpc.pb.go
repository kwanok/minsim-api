// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: service/reddit/submission/submission.proto

package submission

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

// SubmissionClient is the client API for Submission service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubmissionClient interface {
	NewSubmission(ctx context.Context, in *NewSubmissionRequest, opts ...grpc.CallOption) (*NewSubmissionResponse, error)
}

type submissionClient struct {
	cc grpc.ClientConnInterface
}

func NewSubmissionClient(cc grpc.ClientConnInterface) SubmissionClient {
	return &submissionClient{cc}
}

func (c *submissionClient) NewSubmission(ctx context.Context, in *NewSubmissionRequest, opts ...grpc.CallOption) (*NewSubmissionResponse, error) {
	out := new(NewSubmissionResponse)
	err := c.cc.Invoke(ctx, "/server.post.Submission/NewSubmission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubmissionServer is the server API for Submission service.
// All implementations must embed UnimplementedSubmissionServer
// for forward compatibility
type SubmissionServer interface {
	NewSubmission(context.Context, *NewSubmissionRequest) (*NewSubmissionResponse, error)
	mustEmbedUnimplementedSubmissionServer()
}

// UnimplementedSubmissionServer must be embedded to have forward compatible implementations.
type UnimplementedSubmissionServer struct {
}

func (UnimplementedSubmissionServer) NewSubmission(context.Context, *NewSubmissionRequest) (*NewSubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewSubmission not implemented")
}
func (UnimplementedSubmissionServer) mustEmbedUnimplementedSubmissionServer() {}

// UnsafeSubmissionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubmissionServer will
// result in compilation errors.
type UnsafeSubmissionServer interface {
	mustEmbedUnimplementedSubmissionServer()
}

func RegisterSubmissionServer(s grpc.ServiceRegistrar, srv SubmissionServer) {
	s.RegisterService(&Submission_ServiceDesc, srv)
}

func _Submission_NewSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubmissionServer).NewSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.post.Submission/NewSubmission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubmissionServer).NewSubmission(ctx, req.(*NewSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Submission_ServiceDesc is the grpc.ServiceDesc for Submission service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Submission_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "server.post.Submission",
	HandlerType: (*SubmissionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewSubmission",
			Handler:    _Submission_NewSubmission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/reddit/submission/submission.proto",
}