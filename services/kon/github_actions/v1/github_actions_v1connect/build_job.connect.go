// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: services/kon/github_actions/v1/build_job.proto

package github_actions_v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/containerish/OpenRegistry/services/kon/github_actions/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// GithubActionsBuildServiceName is the fully-qualified name of the GithubActionsBuildService
	// service.
	GithubActionsBuildServiceName = "services.kon.github_actions.v1.GithubActionsBuildService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GithubActionsBuildServiceStoreJobProcedure is the fully-qualified name of the
	// GithubActionsBuildService's StoreJob RPC.
	GithubActionsBuildServiceStoreJobProcedure = "/services.kon.github_actions.v1.GithubActionsBuildService/StoreJob"
	// GithubActionsBuildServiceGetBuildJobProcedure is the fully-qualified name of the
	// GithubActionsBuildService's GetBuildJob RPC.
	GithubActionsBuildServiceGetBuildJobProcedure = "/services.kon.github_actions.v1.GithubActionsBuildService/GetBuildJob"
	// GithubActionsBuildServiceTriggerBuildProcedure is the fully-qualified name of the
	// GithubActionsBuildService's TriggerBuild RPC.
	GithubActionsBuildServiceTriggerBuildProcedure = "/services.kon.github_actions.v1.GithubActionsBuildService/TriggerBuild"
	// GithubActionsBuildServiceCancelBuildProcedure is the fully-qualified name of the
	// GithubActionsBuildService's CancelBuild RPC.
	GithubActionsBuildServiceCancelBuildProcedure = "/services.kon.github_actions.v1.GithubActionsBuildService/CancelBuild"
	// GithubActionsBuildServiceDeleteJobProcedure is the fully-qualified name of the
	// GithubActionsBuildService's DeleteJob RPC.
	GithubActionsBuildServiceDeleteJobProcedure = "/services.kon.github_actions.v1.GithubActionsBuildService/DeleteJob"
	// GithubActionsBuildServiceListBuildJobsProcedure is the fully-qualified name of the
	// GithubActionsBuildService's ListBuildJobs RPC.
	GithubActionsBuildServiceListBuildJobsProcedure = "/services.kon.github_actions.v1.GithubActionsBuildService/ListBuildJobs"
	// GithubActionsBuildServiceBulkDeleteBuildJobsProcedure is the fully-qualified name of the
	// GithubActionsBuildService's BulkDeleteBuildJobs RPC.
	GithubActionsBuildServiceBulkDeleteBuildJobsProcedure = "/services.kon.github_actions.v1.GithubActionsBuildService/BulkDeleteBuildJobs"
)

// GithubActionsBuildServiceClient is a client for the
// services.kon.github_actions.v1.GithubActionsBuildService service.
type GithubActionsBuildServiceClient interface {
	StoreJob(context.Context, *connect_go.Request[v1.StoreJobRequest]) (*connect_go.Response[v1.StoreJobResponse], error)
	GetBuildJob(context.Context, *connect_go.Request[v1.GetBuildJobRequest]) (*connect_go.Response[v1.GetBuildJobResponse], error)
	TriggerBuild(context.Context, *connect_go.Request[v1.TriggerBuildRequest]) (*connect_go.Response[v1.TriggerBuildResponse], error)
	CancelBuild(context.Context, *connect_go.Request[v1.CancelBuildRequest]) (*connect_go.Response[v1.CancelBuildResponse], error)
	DeleteJob(context.Context, *connect_go.Request[v1.DeleteJobRequest]) (*connect_go.Response[v1.DeleteJobResponse], error)
	ListBuildJobs(context.Context, *connect_go.Request[v1.ListBuildJobsRequest]) (*connect_go.Response[v1.ListBuildJobsResponse], error)
	BulkDeleteBuildJobs(context.Context, *connect_go.Request[v1.BulkDeleteBuildJobsRequest]) (*connect_go.Response[v1.BulkDeleteBuildJobsResponse], error)
}

// NewGithubActionsBuildServiceClient constructs a client for the
// services.kon.github_actions.v1.GithubActionsBuildService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGithubActionsBuildServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) GithubActionsBuildServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &githubActionsBuildServiceClient{
		storeJob: connect_go.NewClient[v1.StoreJobRequest, v1.StoreJobResponse](
			httpClient,
			baseURL+GithubActionsBuildServiceStoreJobProcedure,
			opts...,
		),
		getBuildJob: connect_go.NewClient[v1.GetBuildJobRequest, v1.GetBuildJobResponse](
			httpClient,
			baseURL+GithubActionsBuildServiceGetBuildJobProcedure,
			opts...,
		),
		triggerBuild: connect_go.NewClient[v1.TriggerBuildRequest, v1.TriggerBuildResponse](
			httpClient,
			baseURL+GithubActionsBuildServiceTriggerBuildProcedure,
			opts...,
		),
		cancelBuild: connect_go.NewClient[v1.CancelBuildRequest, v1.CancelBuildResponse](
			httpClient,
			baseURL+GithubActionsBuildServiceCancelBuildProcedure,
			opts...,
		),
		deleteJob: connect_go.NewClient[v1.DeleteJobRequest, v1.DeleteJobResponse](
			httpClient,
			baseURL+GithubActionsBuildServiceDeleteJobProcedure,
			opts...,
		),
		listBuildJobs: connect_go.NewClient[v1.ListBuildJobsRequest, v1.ListBuildJobsResponse](
			httpClient,
			baseURL+GithubActionsBuildServiceListBuildJobsProcedure,
			opts...,
		),
		bulkDeleteBuildJobs: connect_go.NewClient[v1.BulkDeleteBuildJobsRequest, v1.BulkDeleteBuildJobsResponse](
			httpClient,
			baseURL+GithubActionsBuildServiceBulkDeleteBuildJobsProcedure,
			opts...,
		),
	}
}

// githubActionsBuildServiceClient implements GithubActionsBuildServiceClient.
type githubActionsBuildServiceClient struct {
	storeJob            *connect_go.Client[v1.StoreJobRequest, v1.StoreJobResponse]
	getBuildJob         *connect_go.Client[v1.GetBuildJobRequest, v1.GetBuildJobResponse]
	triggerBuild        *connect_go.Client[v1.TriggerBuildRequest, v1.TriggerBuildResponse]
	cancelBuild         *connect_go.Client[v1.CancelBuildRequest, v1.CancelBuildResponse]
	deleteJob           *connect_go.Client[v1.DeleteJobRequest, v1.DeleteJobResponse]
	listBuildJobs       *connect_go.Client[v1.ListBuildJobsRequest, v1.ListBuildJobsResponse]
	bulkDeleteBuildJobs *connect_go.Client[v1.BulkDeleteBuildJobsRequest, v1.BulkDeleteBuildJobsResponse]
}

// StoreJob calls services.kon.github_actions.v1.GithubActionsBuildService.StoreJob.
func (c *githubActionsBuildServiceClient) StoreJob(ctx context.Context, req *connect_go.Request[v1.StoreJobRequest]) (*connect_go.Response[v1.StoreJobResponse], error) {
	return c.storeJob.CallUnary(ctx, req)
}

// GetBuildJob calls services.kon.github_actions.v1.GithubActionsBuildService.GetBuildJob.
func (c *githubActionsBuildServiceClient) GetBuildJob(ctx context.Context, req *connect_go.Request[v1.GetBuildJobRequest]) (*connect_go.Response[v1.GetBuildJobResponse], error) {
	return c.getBuildJob.CallUnary(ctx, req)
}

// TriggerBuild calls services.kon.github_actions.v1.GithubActionsBuildService.TriggerBuild.
func (c *githubActionsBuildServiceClient) TriggerBuild(ctx context.Context, req *connect_go.Request[v1.TriggerBuildRequest]) (*connect_go.Response[v1.TriggerBuildResponse], error) {
	return c.triggerBuild.CallUnary(ctx, req)
}

// CancelBuild calls services.kon.github_actions.v1.GithubActionsBuildService.CancelBuild.
func (c *githubActionsBuildServiceClient) CancelBuild(ctx context.Context, req *connect_go.Request[v1.CancelBuildRequest]) (*connect_go.Response[v1.CancelBuildResponse], error) {
	return c.cancelBuild.CallUnary(ctx, req)
}

// DeleteJob calls services.kon.github_actions.v1.GithubActionsBuildService.DeleteJob.
func (c *githubActionsBuildServiceClient) DeleteJob(ctx context.Context, req *connect_go.Request[v1.DeleteJobRequest]) (*connect_go.Response[v1.DeleteJobResponse], error) {
	return c.deleteJob.CallUnary(ctx, req)
}

// ListBuildJobs calls services.kon.github_actions.v1.GithubActionsBuildService.ListBuildJobs.
func (c *githubActionsBuildServiceClient) ListBuildJobs(ctx context.Context, req *connect_go.Request[v1.ListBuildJobsRequest]) (*connect_go.Response[v1.ListBuildJobsResponse], error) {
	return c.listBuildJobs.CallUnary(ctx, req)
}

// BulkDeleteBuildJobs calls
// services.kon.github_actions.v1.GithubActionsBuildService.BulkDeleteBuildJobs.
func (c *githubActionsBuildServiceClient) BulkDeleteBuildJobs(ctx context.Context, req *connect_go.Request[v1.BulkDeleteBuildJobsRequest]) (*connect_go.Response[v1.BulkDeleteBuildJobsResponse], error) {
	return c.bulkDeleteBuildJobs.CallUnary(ctx, req)
}

// GithubActionsBuildServiceHandler is an implementation of the
// services.kon.github_actions.v1.GithubActionsBuildService service.
type GithubActionsBuildServiceHandler interface {
	StoreJob(context.Context, *connect_go.Request[v1.StoreJobRequest]) (*connect_go.Response[v1.StoreJobResponse], error)
	GetBuildJob(context.Context, *connect_go.Request[v1.GetBuildJobRequest]) (*connect_go.Response[v1.GetBuildJobResponse], error)
	TriggerBuild(context.Context, *connect_go.Request[v1.TriggerBuildRequest]) (*connect_go.Response[v1.TriggerBuildResponse], error)
	CancelBuild(context.Context, *connect_go.Request[v1.CancelBuildRequest]) (*connect_go.Response[v1.CancelBuildResponse], error)
	DeleteJob(context.Context, *connect_go.Request[v1.DeleteJobRequest]) (*connect_go.Response[v1.DeleteJobResponse], error)
	ListBuildJobs(context.Context, *connect_go.Request[v1.ListBuildJobsRequest]) (*connect_go.Response[v1.ListBuildJobsResponse], error)
	BulkDeleteBuildJobs(context.Context, *connect_go.Request[v1.BulkDeleteBuildJobsRequest]) (*connect_go.Response[v1.BulkDeleteBuildJobsResponse], error)
}

// NewGithubActionsBuildServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGithubActionsBuildServiceHandler(svc GithubActionsBuildServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	githubActionsBuildServiceStoreJobHandler := connect_go.NewUnaryHandler(
		GithubActionsBuildServiceStoreJobProcedure,
		svc.StoreJob,
		opts...,
	)
	githubActionsBuildServiceGetBuildJobHandler := connect_go.NewUnaryHandler(
		GithubActionsBuildServiceGetBuildJobProcedure,
		svc.GetBuildJob,
		opts...,
	)
	githubActionsBuildServiceTriggerBuildHandler := connect_go.NewUnaryHandler(
		GithubActionsBuildServiceTriggerBuildProcedure,
		svc.TriggerBuild,
		opts...,
	)
	githubActionsBuildServiceCancelBuildHandler := connect_go.NewUnaryHandler(
		GithubActionsBuildServiceCancelBuildProcedure,
		svc.CancelBuild,
		opts...,
	)
	githubActionsBuildServiceDeleteJobHandler := connect_go.NewUnaryHandler(
		GithubActionsBuildServiceDeleteJobProcedure,
		svc.DeleteJob,
		opts...,
	)
	githubActionsBuildServiceListBuildJobsHandler := connect_go.NewUnaryHandler(
		GithubActionsBuildServiceListBuildJobsProcedure,
		svc.ListBuildJobs,
		opts...,
	)
	githubActionsBuildServiceBulkDeleteBuildJobsHandler := connect_go.NewUnaryHandler(
		GithubActionsBuildServiceBulkDeleteBuildJobsProcedure,
		svc.BulkDeleteBuildJobs,
		opts...,
	)
	return "/services.kon.github_actions.v1.GithubActionsBuildService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GithubActionsBuildServiceStoreJobProcedure:
			githubActionsBuildServiceStoreJobHandler.ServeHTTP(w, r)
		case GithubActionsBuildServiceGetBuildJobProcedure:
			githubActionsBuildServiceGetBuildJobHandler.ServeHTTP(w, r)
		case GithubActionsBuildServiceTriggerBuildProcedure:
			githubActionsBuildServiceTriggerBuildHandler.ServeHTTP(w, r)
		case GithubActionsBuildServiceCancelBuildProcedure:
			githubActionsBuildServiceCancelBuildHandler.ServeHTTP(w, r)
		case GithubActionsBuildServiceDeleteJobProcedure:
			githubActionsBuildServiceDeleteJobHandler.ServeHTTP(w, r)
		case GithubActionsBuildServiceListBuildJobsProcedure:
			githubActionsBuildServiceListBuildJobsHandler.ServeHTTP(w, r)
		case GithubActionsBuildServiceBulkDeleteBuildJobsProcedure:
			githubActionsBuildServiceBulkDeleteBuildJobsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGithubActionsBuildServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGithubActionsBuildServiceHandler struct{}

func (UnimplementedGithubActionsBuildServiceHandler) StoreJob(context.Context, *connect_go.Request[v1.StoreJobRequest]) (*connect_go.Response[v1.StoreJobResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.kon.github_actions.v1.GithubActionsBuildService.StoreJob is not implemented"))
}

func (UnimplementedGithubActionsBuildServiceHandler) GetBuildJob(context.Context, *connect_go.Request[v1.GetBuildJobRequest]) (*connect_go.Response[v1.GetBuildJobResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.kon.github_actions.v1.GithubActionsBuildService.GetBuildJob is not implemented"))
}

func (UnimplementedGithubActionsBuildServiceHandler) TriggerBuild(context.Context, *connect_go.Request[v1.TriggerBuildRequest]) (*connect_go.Response[v1.TriggerBuildResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.kon.github_actions.v1.GithubActionsBuildService.TriggerBuild is not implemented"))
}

func (UnimplementedGithubActionsBuildServiceHandler) CancelBuild(context.Context, *connect_go.Request[v1.CancelBuildRequest]) (*connect_go.Response[v1.CancelBuildResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.kon.github_actions.v1.GithubActionsBuildService.CancelBuild is not implemented"))
}

func (UnimplementedGithubActionsBuildServiceHandler) DeleteJob(context.Context, *connect_go.Request[v1.DeleteJobRequest]) (*connect_go.Response[v1.DeleteJobResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.kon.github_actions.v1.GithubActionsBuildService.DeleteJob is not implemented"))
}

func (UnimplementedGithubActionsBuildServiceHandler) ListBuildJobs(context.Context, *connect_go.Request[v1.ListBuildJobsRequest]) (*connect_go.Response[v1.ListBuildJobsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.kon.github_actions.v1.GithubActionsBuildService.ListBuildJobs is not implemented"))
}

func (UnimplementedGithubActionsBuildServiceHandler) BulkDeleteBuildJobs(context.Context, *connect_go.Request[v1.BulkDeleteBuildJobsRequest]) (*connect_go.Response[v1.BulkDeleteBuildJobsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.kon.github_actions.v1.GithubActionsBuildService.BulkDeleteBuildJobs is not implemented"))
}
