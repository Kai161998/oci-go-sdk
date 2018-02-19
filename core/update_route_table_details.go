// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateRouteTableDetails The representation of UpdateRouteTableDetails
type UpdateRouteTableDetails struct {

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags ({{DOC_SERVER_URL}}/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags ({{DOC_SERVER_URL}}/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The collection of rules used for routing destination IPs to network devices.
	RouteRules []RouteRule `mandatory:"false" json:"routeRules"`
}

func (m UpdateRouteTableDetails) String() string {
	return common.PointerString(m)
}
