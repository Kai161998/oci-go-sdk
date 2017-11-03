// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// GetConsoleHistoryRequest wrapper for the GetConsoleHistory operation
type GetConsoleHistoryRequest struct {

	// The OCID of the console history.
	InstanceConsoleHistoryID *string `mandatory:"true" contributesTo:"path" name:"instanceConsoleHistoryId"`
}

func (request GetConsoleHistoryRequest) String() string {
	return common.PointerString(request)
}

// GetConsoleHistoryResponse wrapper for the GetConsoleHistory operation
type GetConsoleHistoryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ConsoleHistory instance
	ConsoleHistory `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetConsoleHistoryResponse) String() string {
	return common.PointerString(response)
}
