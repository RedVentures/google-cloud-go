// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package gsuiteaddons

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	gsuiteaddonspb "google.golang.org/genproto/googleapis/cloud/gsuiteaddons/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	GetAuthorization    []gax.CallOption
	CreateDeployment    []gax.CallOption
	ReplaceDeployment   []gax.CallOption
	GetDeployment       []gax.CallOption
	ListDeployments     []gax.CallOption
	DeleteDeployment    []gax.CallOption
	InstallDeployment   []gax.CallOption
	UninstallDeployment []gax.CallOption
	GetInstallStatus    []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("gsuiteaddons.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("gsuiteaddons.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://gsuiteaddons.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		GetAuthorization: []gax.CallOption{},
		CreateDeployment: []gax.CallOption{},
		ReplaceDeployment: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Unknown,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetDeployment: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Unknown,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListDeployments: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Unknown,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteDeployment: []gax.CallOption{},
		InstallDeployment: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Unknown,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UninstallDeployment: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Unknown,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetInstallStatus: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.Unknown,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalClient is an interface that defines the methods availaible from Google Workspace Add-ons API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetAuthorization(context.Context, *gsuiteaddonspb.GetAuthorizationRequest, ...gax.CallOption) (*gsuiteaddonspb.Authorization, error)
	CreateDeployment(context.Context, *gsuiteaddonspb.CreateDeploymentRequest, ...gax.CallOption) (*gsuiteaddonspb.Deployment, error)
	ReplaceDeployment(context.Context, *gsuiteaddonspb.ReplaceDeploymentRequest, ...gax.CallOption) (*gsuiteaddonspb.Deployment, error)
	GetDeployment(context.Context, *gsuiteaddonspb.GetDeploymentRequest, ...gax.CallOption) (*gsuiteaddonspb.Deployment, error)
	ListDeployments(context.Context, *gsuiteaddonspb.ListDeploymentsRequest, ...gax.CallOption) *DeploymentIterator
	DeleteDeployment(context.Context, *gsuiteaddonspb.DeleteDeploymentRequest, ...gax.CallOption) error
	InstallDeployment(context.Context, *gsuiteaddonspb.InstallDeploymentRequest, ...gax.CallOption) error
	UninstallDeployment(context.Context, *gsuiteaddonspb.UninstallDeploymentRequest, ...gax.CallOption) error
	GetInstallStatus(context.Context, *gsuiteaddonspb.GetInstallStatusRequest, ...gax.CallOption) (*gsuiteaddonspb.InstallStatus, error)
}

// Client is a client for interacting with Google Workspace Add-ons API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// A service for managing Google Workspace Add-ons deployments.
//
// A Google Workspace Add-on is a third-party embedded component that can be
// installed in Google Workspace Applications like Gmail, Calendar, Drive, and
// the Google Docs, Sheets, and Slides editors. Google Workspace Add-ons can
// display UI cards, receive contextual information from the host application,
// and perform actions in the host application (See:
// https://developers.google.com/gsuite/add-ons/overview (at https://developers.google.com/gsuite/add-ons/overview) for more information).
//
// A Google Workspace Add-on deployment resource specifies metadata about the
// add-on, including a specification of the entry points in the host application
// that trigger add-on executions (see:
// https://developers.google.com/gsuite/add-ons/concepts/gsuite-manifests (at https://developers.google.com/gsuite/add-ons/concepts/gsuite-manifests)).
// Add-on deployments defined via the Google Workspace Add-ons API define their
// entrypoints using HTTPS URLs (See:
// https://developers.google.com/gsuite/add-ons/guides/alternate-runtimes (at https://developers.google.com/gsuite/add-ons/guides/alternate-runtimes)),
//
// A Google Workspace Add-on deployment can be installed in developer mode,
// which allows an add-on developer to test the experience an end-user would see
// when installing and running the add-on in their G Suite applications.  When
// running in developer mode, more detailed error messages are exposed in the
// add-on UI to aid in debugging.
//
// A Google Workspace Add-on deployment can be published to Google Workspace
// Marketplace, which allows other Google Workspace users to discover and
// install the add-on.  See:
// https://developers.google.com/gsuite/add-ons/how-tos/publish-add-on-overview (at https://developers.google.com/gsuite/add-ons/how-tos/publish-add-on-overview)
// for details.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetAuthorization gets the authorization information for deployments in a given project.
func (c *Client) GetAuthorization(ctx context.Context, req *gsuiteaddonspb.GetAuthorizationRequest, opts ...gax.CallOption) (*gsuiteaddonspb.Authorization, error) {
	return c.internalClient.GetAuthorization(ctx, req, opts...)
}

// CreateDeployment creates a deployment with the specified name and configuration.
func (c *Client) CreateDeployment(ctx context.Context, req *gsuiteaddonspb.CreateDeploymentRequest, opts ...gax.CallOption) (*gsuiteaddonspb.Deployment, error) {
	return c.internalClient.CreateDeployment(ctx, req, opts...)
}

// ReplaceDeployment creates or replaces a deployment with the specified name.
func (c *Client) ReplaceDeployment(ctx context.Context, req *gsuiteaddonspb.ReplaceDeploymentRequest, opts ...gax.CallOption) (*gsuiteaddonspb.Deployment, error) {
	return c.internalClient.ReplaceDeployment(ctx, req, opts...)
}

// GetDeployment gets the deployment with the specified name.
func (c *Client) GetDeployment(ctx context.Context, req *gsuiteaddonspb.GetDeploymentRequest, opts ...gax.CallOption) (*gsuiteaddonspb.Deployment, error) {
	return c.internalClient.GetDeployment(ctx, req, opts...)
}

// ListDeployments lists all deployments in a particular project.
func (c *Client) ListDeployments(ctx context.Context, req *gsuiteaddonspb.ListDeploymentsRequest, opts ...gax.CallOption) *DeploymentIterator {
	return c.internalClient.ListDeployments(ctx, req, opts...)
}

// DeleteDeployment deletes the deployment with the given name.
func (c *Client) DeleteDeployment(ctx context.Context, req *gsuiteaddonspb.DeleteDeploymentRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteDeployment(ctx, req, opts...)
}

// InstallDeployment installs a deployment in developer mode.
// See:
// https://developers.google.com/gsuite/add-ons/how-tos/testing-gsuite-addons (at https://developers.google.com/gsuite/add-ons/how-tos/testing-gsuite-addons).
func (c *Client) InstallDeployment(ctx context.Context, req *gsuiteaddonspb.InstallDeploymentRequest, opts ...gax.CallOption) error {
	return c.internalClient.InstallDeployment(ctx, req, opts...)
}

// UninstallDeployment uninstalls a developer mode deployment.
// See:
// https://developers.google.com/gsuite/add-ons/how-tos/testing-gsuite-addons (at https://developers.google.com/gsuite/add-ons/how-tos/testing-gsuite-addons).
func (c *Client) UninstallDeployment(ctx context.Context, req *gsuiteaddonspb.UninstallDeploymentRequest, opts ...gax.CallOption) error {
	return c.internalClient.UninstallDeployment(ctx, req, opts...)
}

// GetInstallStatus fetches the install status of a developer mode deployment.
func (c *Client) GetInstallStatus(ctx context.Context, req *gsuiteaddonspb.GetInstallStatusRequest, opts ...gax.CallOption) (*gsuiteaddonspb.InstallStatus, error) {
	return c.internalClient.GetInstallStatus(ctx, req, opts...)
}

// gRPCClient is a client for interacting with Google Workspace Add-ons API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client gsuiteaddonspb.GSuiteAddOnsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new g suite add ons client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// A service for managing Google Workspace Add-ons deployments.
//
// A Google Workspace Add-on is a third-party embedded component that can be
// installed in Google Workspace Applications like Gmail, Calendar, Drive, and
// the Google Docs, Sheets, and Slides editors. Google Workspace Add-ons can
// display UI cards, receive contextual information from the host application,
// and perform actions in the host application (See:
// https://developers.google.com/gsuite/add-ons/overview (at https://developers.google.com/gsuite/add-ons/overview) for more information).
//
// A Google Workspace Add-on deployment resource specifies metadata about the
// add-on, including a specification of the entry points in the host application
// that trigger add-on executions (see:
// https://developers.google.com/gsuite/add-ons/concepts/gsuite-manifests (at https://developers.google.com/gsuite/add-ons/concepts/gsuite-manifests)).
// Add-on deployments defined via the Google Workspace Add-ons API define their
// entrypoints using HTTPS URLs (See:
// https://developers.google.com/gsuite/add-ons/guides/alternate-runtimes (at https://developers.google.com/gsuite/add-ons/guides/alternate-runtimes)),
//
// A Google Workspace Add-on deployment can be installed in developer mode,
// which allows an add-on developer to test the experience an end-user would see
// when installing and running the add-on in their G Suite applications.  When
// running in developer mode, more detailed error messages are exposed in the
// add-on UI to aid in debugging.
//
// A Google Workspace Add-on deployment can be published to Google Workspace
// Marketplace, which allows other Google Workspace users to discover and
// install the add-on.  See:
// https://developers.google.com/gsuite/add-ons/how-tos/publish-add-on-overview (at https://developers.google.com/gsuite/add-ons/how-tos/publish-add-on-overview)
// for details.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           gsuiteaddonspb.NewGSuiteAddOnsClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) GetAuthorization(ctx context.Context, req *gsuiteaddonspb.GetAuthorizationRequest, opts ...gax.CallOption) (*gsuiteaddonspb.Authorization, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 120000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetAuthorization[0:len((*c.CallOptions).GetAuthorization):len((*c.CallOptions).GetAuthorization)], opts...)
	var resp *gsuiteaddonspb.Authorization
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetAuthorization(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) CreateDeployment(ctx context.Context, req *gsuiteaddonspb.CreateDeploymentRequest, opts ...gax.CallOption) (*gsuiteaddonspb.Deployment, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateDeployment[0:len((*c.CallOptions).CreateDeployment):len((*c.CallOptions).CreateDeployment)], opts...)
	var resp *gsuiteaddonspb.Deployment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.CreateDeployment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ReplaceDeployment(ctx context.Context, req *gsuiteaddonspb.ReplaceDeploymentRequest, opts ...gax.CallOption) (*gsuiteaddonspb.Deployment, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "deployment.name", url.QueryEscape(req.GetDeployment().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ReplaceDeployment[0:len((*c.CallOptions).ReplaceDeployment):len((*c.CallOptions).ReplaceDeployment)], opts...)
	var resp *gsuiteaddonspb.Deployment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ReplaceDeployment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) GetDeployment(ctx context.Context, req *gsuiteaddonspb.GetDeploymentRequest, opts ...gax.CallOption) (*gsuiteaddonspb.Deployment, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetDeployment[0:len((*c.CallOptions).GetDeployment):len((*c.CallOptions).GetDeployment)], opts...)
	var resp *gsuiteaddonspb.Deployment
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetDeployment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ListDeployments(ctx context.Context, req *gsuiteaddonspb.ListDeploymentsRequest, opts ...gax.CallOption) *DeploymentIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListDeployments[0:len((*c.CallOptions).ListDeployments):len((*c.CallOptions).ListDeployments)], opts...)
	it := &DeploymentIterator{}
	req = proto.Clone(req).(*gsuiteaddonspb.ListDeploymentsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*gsuiteaddonspb.Deployment, string, error) {
		resp := &gsuiteaddonspb.ListDeploymentsResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.client.ListDeployments(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetDeployments(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *gRPCClient) DeleteDeployment(ctx context.Context, req *gsuiteaddonspb.DeleteDeploymentRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteDeployment[0:len((*c.CallOptions).DeleteDeployment):len((*c.CallOptions).DeleteDeployment)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.DeleteDeployment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) InstallDeployment(ctx context.Context, req *gsuiteaddonspb.InstallDeploymentRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).InstallDeployment[0:len((*c.CallOptions).InstallDeployment):len((*c.CallOptions).InstallDeployment)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.InstallDeployment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) UninstallDeployment(ctx context.Context, req *gsuiteaddonspb.UninstallDeploymentRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UninstallDeployment[0:len((*c.CallOptions).UninstallDeployment):len((*c.CallOptions).UninstallDeployment)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.client.UninstallDeployment(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *gRPCClient) GetInstallStatus(ctx context.Context, req *gsuiteaddonspb.GetInstallStatusRequest, opts ...gax.CallOption) (*gsuiteaddonspb.InstallStatus, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 10000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetInstallStatus[0:len((*c.CallOptions).GetInstallStatus):len((*c.CallOptions).GetInstallStatus)], opts...)
	var resp *gsuiteaddonspb.InstallStatus
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.GetInstallStatus(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeploymentIterator manages a stream of *gsuiteaddonspb.Deployment.
type DeploymentIterator struct {
	items    []*gsuiteaddonspb.Deployment
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*gsuiteaddonspb.Deployment, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *DeploymentIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *DeploymentIterator) Next() (*gsuiteaddonspb.Deployment, error) {
	var item *gsuiteaddonspb.Deployment
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *DeploymentIterator) bufLen() int {
	return len(it.items)
}

func (it *DeploymentIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
