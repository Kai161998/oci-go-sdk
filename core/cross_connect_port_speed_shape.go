// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"bitbucket.aka.lgl.grungy.us/golang-sdk2/common"
)

// CrossConnectPortSpeedShape. An individual port speed level for cross-connects.
type CrossConnectPortSpeedShape struct {

	// The name of the port speed shape.
	// Example: `10 Gbps`
	Name *string `mandatory:"true" json:"name,omitempty"`

	// The port speed in Gbps.
	// Example: `10`
	PortSpeedInGbps *int `mandatory:"true" json:"portSpeedInGbps,omitempty"`
}

func (model CrossConnectPortSpeedShape) String() string {
	return common.PointerString(model)
}
