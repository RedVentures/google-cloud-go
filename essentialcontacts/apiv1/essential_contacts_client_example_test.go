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

package essentialcontacts_test

import (
	"context"

	essentialcontacts "cloud.google.com/go/essentialcontacts/apiv1"
	"google.golang.org/api/iterator"
	essentialcontactspb "google.golang.org/genproto/googleapis/cloud/essentialcontacts/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := essentialcontacts.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_CreateContact() {
	ctx := context.Background()
	c, err := essentialcontacts.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &essentialcontactspb.CreateContactRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/essentialcontacts/v1#CreateContactRequest.
	}
	resp, err := c.CreateContact(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateContact() {
	ctx := context.Background()
	c, err := essentialcontacts.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &essentialcontactspb.UpdateContactRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/essentialcontacts/v1#UpdateContactRequest.
	}
	resp, err := c.UpdateContact(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListContacts() {
	ctx := context.Background()
	c, err := essentialcontacts.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &essentialcontactspb.ListContactsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/essentialcontacts/v1#ListContactsRequest.
	}
	it := c.ListContacts(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetContact() {
	ctx := context.Background()
	c, err := essentialcontacts.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &essentialcontactspb.GetContactRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/essentialcontacts/v1#GetContactRequest.
	}
	resp, err := c.GetContact(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteContact() {
	ctx := context.Background()
	c, err := essentialcontacts.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &essentialcontactspb.DeleteContactRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/essentialcontacts/v1#DeleteContactRequest.
	}
	err = c.DeleteContact(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ComputeContacts() {
	ctx := context.Background()
	c, err := essentialcontacts.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &essentialcontactspb.ComputeContactsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/essentialcontacts/v1#ComputeContactsRequest.
	}
	it := c.ComputeContacts(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_SendTestMessage() {
	ctx := context.Background()
	c, err := essentialcontacts.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &essentialcontactspb.SendTestMessageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/essentialcontacts/v1#SendTestMessageRequest.
	}
	err = c.SendTestMessage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
