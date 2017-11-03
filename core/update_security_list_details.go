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

type UpdateSecurityListDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// Rules for allowing egress IP packets.
	EgressSecurityRules *[]EgressSecurityRule `mandatory:"false" json:"egressSecurityRules,omitempty"`

	// Rules for allowing ingress IP packets.
	IngressSecurityRules *[]IngressSecurityRule `mandatory:"false" json:"ingressSecurityRules,omitempty"`
}

func (model UpdateSecurityListDetails) String() string {
	return common.PointerString(model)
}
