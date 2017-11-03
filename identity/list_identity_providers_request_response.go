// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package identity

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
	"net/http"
)

// ListIdentityProvidersRequest wrapper for the ListIdentityProviders operation
type ListIdentityProvidersRequest struct {

	// The protocol used for federation.
	Protocol ListIdentityProvidersProtocolEnum `mandatory:"true" contributesTo:"query" name:"protocol"`

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentID *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`
}

func (request ListIdentityProvidersRequest) String() string {
	return common.PointerString(request)
}

// ListIdentityProvidersResponse wrapper for the ListIdentityProviders operation
type ListIdentityProvidersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []IdentityProvider instance
	Items []IdentityProvider `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestID *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListIdentityProvidersResponse) String() string {
	return common.PointerString(response)
}

type ListIdentityProvidersProtocolEnum string
type ListIdentityProvidersProtocol struct{}

const (
	LIST_IDENTITY_PROVIDERS_PROTOCOL_SAML2   ListIdentityProvidersProtocolEnum = "SAML2"
	LIST_IDENTITY_PROVIDERS_PROTOCOL_UNKNOWN ListIdentityProvidersProtocolEnum = "UNKNOWN"
)

var mapping_Identity_protocol = map[string]ListIdentityProvidersProtocolEnum{
	"SAML2":   LIST_IDENTITY_PROVIDERS_PROTOCOL_SAML2,
	"UNKNOWN": LIST_IDENTITY_PROVIDERS_PROTOCOL_UNKNOWN,
}

func (receiver ListIdentityProvidersProtocol) Values() []ListIdentityProvidersProtocolEnum {
	values := make([]ListIdentityProvidersProtocolEnum, 0)
	for _, v := range mapping_Identity_protocol {
		if v != LIST_IDENTITY_PROVIDERS_PROTOCOL_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver ListIdentityProvidersProtocol) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if ListIdentityProvidersProtocolEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver ListIdentityProvidersProtocol) From(toBeConverted string) ListIdentityProvidersProtocolEnum {
	if val, ok := mapping_Identity_protocol[toBeConverted]; ok {
		return val
	}
	return LIST_IDENTITY_PROVIDERS_PROTOCOL_UNKNOWN
}
