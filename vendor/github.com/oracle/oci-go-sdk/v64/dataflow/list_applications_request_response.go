// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"net/http"
	"strings"
)

// ListApplicationsRequest wrapper for the ListApplications operation
type ListApplicationsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request. If provided, the returned request ID will include this value.
	// Otherwise, a random request ID will be generated by the service.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of results to return in a paginated `List` call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` or `opc-prev-page` response header from the last `List` call
	// to sent back to server for getting the next page of results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field used to sort the results. Multiple fields are not supported.
	SortBy ListApplicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The ordering of results in ascending or descending order.
	SortOrder ListApplicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The query parameter for the Spark application name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID of the user who created the resource.
	OwnerPrincipalId *string `mandatory:"false" contributesTo:"query" name:"ownerPrincipalId"`

	// The displayName prefix.
	DisplayNameStartsWith *string `mandatory:"false" contributesTo:"query" name:"displayNameStartsWith"`

	// The Spark version utilized to run the application.
	SparkVersion *string `mandatory:"false" contributesTo:"query" name:"sparkVersion"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListApplicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListApplicationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListApplicationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListApplicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListApplicationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListApplicationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListApplicationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListApplicationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListApplicationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListApplicationsResponse wrapper for the ListApplications operation
type ListApplicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ApplicationSummary instances
	Items []ApplicationSummary `presentIn:"body"`

	// Retrieves the previous page of results.
	// When this header appears in the response, previous pages of results exist.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Retrieves the next page of results. When this header appears in the response,
	// additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListApplicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListApplicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListApplicationsSortByEnum Enum with underlying type: string
type ListApplicationsSortByEnum string

// Set of constants representing the allowable values for ListApplicationsSortByEnum
const (
	ListApplicationsSortByTimecreated ListApplicationsSortByEnum = "timeCreated"
	ListApplicationsSortByDisplayname ListApplicationsSortByEnum = "displayName"
	ListApplicationsSortByLanguage    ListApplicationsSortByEnum = "language"
)

var mappingListApplicationsSortByEnum = map[string]ListApplicationsSortByEnum{
	"timeCreated": ListApplicationsSortByTimecreated,
	"displayName": ListApplicationsSortByDisplayname,
	"language":    ListApplicationsSortByLanguage,
}

var mappingListApplicationsSortByEnumLowerCase = map[string]ListApplicationsSortByEnum{
	"timecreated": ListApplicationsSortByTimecreated,
	"displayname": ListApplicationsSortByDisplayname,
	"language":    ListApplicationsSortByLanguage,
}

// GetListApplicationsSortByEnumValues Enumerates the set of values for ListApplicationsSortByEnum
func GetListApplicationsSortByEnumValues() []ListApplicationsSortByEnum {
	values := make([]ListApplicationsSortByEnum, 0)
	for _, v := range mappingListApplicationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplicationsSortByEnumStringValues Enumerates the set of values in String for ListApplicationsSortByEnum
func GetListApplicationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"language",
	}
}

// GetMappingListApplicationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplicationsSortByEnum(val string) (ListApplicationsSortByEnum, bool) {
	enum, ok := mappingListApplicationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListApplicationsSortOrderEnum Enum with underlying type: string
type ListApplicationsSortOrderEnum string

// Set of constants representing the allowable values for ListApplicationsSortOrderEnum
const (
	ListApplicationsSortOrderAsc  ListApplicationsSortOrderEnum = "ASC"
	ListApplicationsSortOrderDesc ListApplicationsSortOrderEnum = "DESC"
)

var mappingListApplicationsSortOrderEnum = map[string]ListApplicationsSortOrderEnum{
	"ASC":  ListApplicationsSortOrderAsc,
	"DESC": ListApplicationsSortOrderDesc,
}

var mappingListApplicationsSortOrderEnumLowerCase = map[string]ListApplicationsSortOrderEnum{
	"asc":  ListApplicationsSortOrderAsc,
	"desc": ListApplicationsSortOrderDesc,
}

// GetListApplicationsSortOrderEnumValues Enumerates the set of values for ListApplicationsSortOrderEnum
func GetListApplicationsSortOrderEnumValues() []ListApplicationsSortOrderEnum {
	values := make([]ListApplicationsSortOrderEnum, 0)
	for _, v := range mappingListApplicationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListApplicationsSortOrderEnumStringValues Enumerates the set of values in String for ListApplicationsSortOrderEnum
func GetListApplicationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListApplicationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListApplicationsSortOrderEnum(val string) (ListApplicationsSortOrderEnum, bool) {
	enum, ok := mappingListApplicationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
