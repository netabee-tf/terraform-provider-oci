// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v33/common"
	"net/http"
)

// DeleteTaskRunRequest wrapper for the DeleteTaskRun operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/DeleteTaskRun.go.html to see an example of how to use DeleteTaskRunRequest.
type DeleteTaskRunRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The application key.
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// The task run key.
	TaskRunKey *string `mandatory:"true" contributesTo:"path" name:"taskRunKey"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match` parameter to the value of the `etag` from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the `etag` you provide matches the resource's current `etag` value.
	// When 'if-match' is provided and its value does not exactly match the 'etag' of the resource on the server, the request fails with the 412 response code.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteTaskRunRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteTaskRunRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteTaskRunRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DeleteTaskRunResponse wrapper for the DeleteTaskRun operation
type DeleteTaskRunResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteTaskRunResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteTaskRunResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}