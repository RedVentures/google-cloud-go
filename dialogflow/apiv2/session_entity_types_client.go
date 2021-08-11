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

package dialogflow

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
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newSessionEntityTypesClientHook clientHook

// SessionEntityTypesCallOptions contains the retry settings for each method of SessionEntityTypesClient.
type SessionEntityTypesCallOptions struct {
	ListSessionEntityTypes  []gax.CallOption
	GetSessionEntityType    []gax.CallOption
	CreateSessionEntityType []gax.CallOption
	UpdateSessionEntityType []gax.CallOption
	DeleteSessionEntityType []gax.CallOption
}

func defaultSessionEntityTypesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSessionEntityTypesCallOptions() *SessionEntityTypesCallOptions {
	return &SessionEntityTypesCallOptions{
		ListSessionEntityTypes: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetSessionEntityType: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateSessionEntityType: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateSessionEntityType: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteSessionEntityType: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalSessionEntityTypesClient is an interface that defines the methods availaible from Dialogflow API.
type internalSessionEntityTypesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListSessionEntityTypes(context.Context, *dialogflowpb.ListSessionEntityTypesRequest, ...gax.CallOption) *SessionEntityTypeIterator
	GetSessionEntityType(context.Context, *dialogflowpb.GetSessionEntityTypeRequest, ...gax.CallOption) (*dialogflowpb.SessionEntityType, error)
	CreateSessionEntityType(context.Context, *dialogflowpb.CreateSessionEntityTypeRequest, ...gax.CallOption) (*dialogflowpb.SessionEntityType, error)
	UpdateSessionEntityType(context.Context, *dialogflowpb.UpdateSessionEntityTypeRequest, ...gax.CallOption) (*dialogflowpb.SessionEntityType, error)
	DeleteSessionEntityType(context.Context, *dialogflowpb.DeleteSessionEntityTypeRequest, ...gax.CallOption) error
}

// SessionEntityTypesClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing SessionEntityTypes.
type SessionEntityTypesClient struct {
	// The internal transport-dependent client.
	internalClient internalSessionEntityTypesClient

	// The call options for this service.
	CallOptions *SessionEntityTypesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SessionEntityTypesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SessionEntityTypesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *SessionEntityTypesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListSessionEntityTypes returns the list of all session entity types in the specified session.
//
// This method doesn’t work with Google Assistant integration.
// Contact Dialogflow support if you need to use session entities
// with Google Assistant integration.
func (c *SessionEntityTypesClient) ListSessionEntityTypes(ctx context.Context, req *dialogflowpb.ListSessionEntityTypesRequest, opts ...gax.CallOption) *SessionEntityTypeIterator {
	return c.internalClient.ListSessionEntityTypes(ctx, req, opts...)
}

// GetSessionEntityType retrieves the specified session entity type.
//
// This method doesn’t work with Google Assistant integration.
// Contact Dialogflow support if you need to use session entities
// with Google Assistant integration.
func (c *SessionEntityTypesClient) GetSessionEntityType(ctx context.Context, req *dialogflowpb.GetSessionEntityTypeRequest, opts ...gax.CallOption) (*dialogflowpb.SessionEntityType, error) {
	return c.internalClient.GetSessionEntityType(ctx, req, opts...)
}

// CreateSessionEntityType creates a session entity type.
//
// If the specified session entity type already exists, overrides the session
// entity type.
//
// This method doesn’t work with Google Assistant integration.
// Contact Dialogflow support if you need to use session entities
// with Google Assistant integration.
func (c *SessionEntityTypesClient) CreateSessionEntityType(ctx context.Context, req *dialogflowpb.CreateSessionEntityTypeRequest, opts ...gax.CallOption) (*dialogflowpb.SessionEntityType, error) {
	return c.internalClient.CreateSessionEntityType(ctx, req, opts...)
}

// UpdateSessionEntityType updates the specified session entity type.
//
// This method doesn’t work with Google Assistant integration.
// Contact Dialogflow support if you need to use session entities
// with Google Assistant integration.
func (c *SessionEntityTypesClient) UpdateSessionEntityType(ctx context.Context, req *dialogflowpb.UpdateSessionEntityTypeRequest, opts ...gax.CallOption) (*dialogflowpb.SessionEntityType, error) {
	return c.internalClient.UpdateSessionEntityType(ctx, req, opts...)
}

// DeleteSessionEntityType deletes the specified session entity type.
//
// This method doesn’t work with Google Assistant integration.
// Contact Dialogflow support if you need to use session entities
// with Google Assistant integration.
func (c *SessionEntityTypesClient) DeleteSessionEntityType(ctx context.Context, req *dialogflowpb.DeleteSessionEntityTypeRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteSessionEntityType(ctx, req, opts...)
}

// sessionEntityTypesGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type sessionEntityTypesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing SessionEntityTypesClient
	CallOptions **SessionEntityTypesCallOptions

	// The gRPC API client.
	sessionEntityTypesClient dialogflowpb.SessionEntityTypesClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewSessionEntityTypesClient creates a new session entity types client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing SessionEntityTypes.
func NewSessionEntityTypesClient(ctx context.Context, opts ...option.ClientOption) (*SessionEntityTypesClient, error) {
	clientOpts := defaultSessionEntityTypesGRPCClientOptions()
	if newSessionEntityTypesClientHook != nil {
		hookOpts, err := newSessionEntityTypesClientHook(ctx, clientHookParams{})
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
	client := SessionEntityTypesClient{CallOptions: defaultSessionEntityTypesCallOptions()}

	c := &sessionEntityTypesGRPCClient{
		connPool:                 connPool,
		disableDeadlines:         disableDeadlines,
		sessionEntityTypesClient: dialogflowpb.NewSessionEntityTypesClient(connPool),
		CallOptions:              &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *sessionEntityTypesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *sessionEntityTypesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *sessionEntityTypesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *sessionEntityTypesGRPCClient) ListSessionEntityTypes(ctx context.Context, req *dialogflowpb.ListSessionEntityTypesRequest, opts ...gax.CallOption) *SessionEntityTypeIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListSessionEntityTypes[0:len((*c.CallOptions).ListSessionEntityTypes):len((*c.CallOptions).ListSessionEntityTypes)], opts...)
	it := &SessionEntityTypeIterator{}
	req = proto.Clone(req).(*dialogflowpb.ListSessionEntityTypesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dialogflowpb.SessionEntityType, string, error) {
		resp := &dialogflowpb.ListSessionEntityTypesResponse{}
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
			resp, err = c.sessionEntityTypesClient.ListSessionEntityTypes(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetSessionEntityTypes(), resp.GetNextPageToken(), nil
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

func (c *sessionEntityTypesGRPCClient) GetSessionEntityType(ctx context.Context, req *dialogflowpb.GetSessionEntityTypeRequest, opts ...gax.CallOption) (*dialogflowpb.SessionEntityType, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetSessionEntityType[0:len((*c.CallOptions).GetSessionEntityType):len((*c.CallOptions).GetSessionEntityType)], opts...)
	var resp *dialogflowpb.SessionEntityType
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.sessionEntityTypesClient.GetSessionEntityType(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *sessionEntityTypesGRPCClient) CreateSessionEntityType(ctx context.Context, req *dialogflowpb.CreateSessionEntityTypeRequest, opts ...gax.CallOption) (*dialogflowpb.SessionEntityType, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateSessionEntityType[0:len((*c.CallOptions).CreateSessionEntityType):len((*c.CallOptions).CreateSessionEntityType)], opts...)
	var resp *dialogflowpb.SessionEntityType
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.sessionEntityTypesClient.CreateSessionEntityType(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *sessionEntityTypesGRPCClient) UpdateSessionEntityType(ctx context.Context, req *dialogflowpb.UpdateSessionEntityTypeRequest, opts ...gax.CallOption) (*dialogflowpb.SessionEntityType, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session_entity_type.name", url.QueryEscape(req.GetSessionEntityType().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateSessionEntityType[0:len((*c.CallOptions).UpdateSessionEntityType):len((*c.CallOptions).UpdateSessionEntityType)], opts...)
	var resp *dialogflowpb.SessionEntityType
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.sessionEntityTypesClient.UpdateSessionEntityType(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *sessionEntityTypesGRPCClient) DeleteSessionEntityType(ctx context.Context, req *dialogflowpb.DeleteSessionEntityTypeRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteSessionEntityType[0:len((*c.CallOptions).DeleteSessionEntityType):len((*c.CallOptions).DeleteSessionEntityType)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.sessionEntityTypesClient.DeleteSessionEntityType(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// SessionEntityTypeIterator manages a stream of *dialogflowpb.SessionEntityType.
type SessionEntityTypeIterator struct {
	items    []*dialogflowpb.SessionEntityType
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
	InternalFetch func(pageSize int, pageToken string) (results []*dialogflowpb.SessionEntityType, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SessionEntityTypeIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SessionEntityTypeIterator) Next() (*dialogflowpb.SessionEntityType, error) {
	var item *dialogflowpb.SessionEntityType
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SessionEntityTypeIterator) bufLen() int {
	return len(it.items)
}

func (it *SessionEntityTypeIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
