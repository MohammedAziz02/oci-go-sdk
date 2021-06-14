// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicecatalog

import (
	"github.com/oracle/oci-go-sdk/v42/common"
	"net/http"
)

// GetPrivateApplicationPackageRequest wrapper for the GetPrivateApplicationPackage operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/GetPrivateApplicationPackage.go.html to see an example of how to use GetPrivateApplicationPackageRequest.
type GetPrivateApplicationPackageRequest struct {

	// The unique identifier for the private application package.
	PrivateApplicationPackageId *string `mandatory:"true" contributesTo:"path" name:"privateApplicationPackageId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetPrivateApplicationPackageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetPrivateApplicationPackageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetPrivateApplicationPackageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetPrivateApplicationPackageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetPrivateApplicationPackageResponse wrapper for the GetPrivateApplicationPackage operation
type GetPrivateApplicationPackageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The PrivateApplicationPackage instance
	PrivateApplicationPackage `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetPrivateApplicationPackageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetPrivateApplicationPackageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
