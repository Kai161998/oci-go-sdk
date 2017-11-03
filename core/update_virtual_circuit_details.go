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

type UpdateVirtualCircuitDetails struct {

	// The provisioned data rate of the connection. To get a list of the
	// available bandwidth levels (i.e., shapes), see
	// ListFastConnectProviderVirtualCircuitBandwidthShapes.
	// To be updated only by the customer who owns the virtual circuit.
	BandwidthShapeName *string `mandatory:"false" json:"bandwidthShapeName,omitempty"`

	// An array of mappings, each containing properties for a cross-connect or
	// cross-connect group associated with this virtual circuit.
	// The customer and provider can update different properties in the mapping
	// depending on the situation. See the description of the
	// CrossConnectMapping.
	CrossConnectMappings *[]CrossConnectMapping `mandatory:"false" json:"crossConnectMappings,omitempty"`

	// The BGP ASN of the network at the other end of the BGP
	// session from Oracle.
	// If the BGP session is from the customer's edge router to Oracle, the
	// required value is the customer's ASN, and it can be updated only
	// by the customer.
	// If the BGP session is from the provider's edge router to Oracle, the
	// required value is the provider's ASN, and it can be updated only
	// by the provider.
	CustomerBgpAsn *int `mandatory:"false" json:"customerBgpAsn,omitempty"`

	// A user-friendly name. Does not have to be unique.
	// Avoid entering confidential information.
	// To be updated only by the customer who owns the virtual circuit.
	DisplayName *string `mandatory:"false" json:"displayName,omitempty"`

	// The OCID of the Drg
	// that this virtual circuit uses.
	// To be updated only by the customer who owns the virtual circuit.
	GatewayID *string `mandatory:"false" json:"gatewayId,omitempty"`

	// The provider's state in relation to this virtual circuit. Relevant only
	// if the customer is using FastConnect via a provider.  ACTIVE
	// means the provider has provisioned the virtual circuit from their
	// end. INACTIVE means the provider has not yet provisioned the virtual
	// circuit, or has de-provisioned it.
	// To be updated only by the provider.
	ProviderState UpdateVirtualCircuitDetailsProviderStateEnum `mandatory:"false" json:"providerState,omitempty"`

	// Provider-supplied reference information about this virtual circuit.
	// Relevant only if the customer is using FastConnect via a provider.
	// To be updated only by the provider.
	ReferenceComment *string `mandatory:"false" json:"referenceComment,omitempty"`
}

func (model UpdateVirtualCircuitDetails) String() string {
	return common.PointerString(model)
}

type UpdateVirtualCircuitDetailsProviderStateEnum string
type UpdateVirtualCircuitDetailsProviderState struct{}

const (
	UPDATE_VIRTUAL_CIRCUIT_DETAILS_PROVIDER_STATE_ACTIVE   UpdateVirtualCircuitDetailsProviderStateEnum = "ACTIVE"
	UPDATE_VIRTUAL_CIRCUIT_DETAILS_PROVIDER_STATE_INACTIVE UpdateVirtualCircuitDetailsProviderStateEnum = "INACTIVE"
	UPDATE_VIRTUAL_CIRCUIT_DETAILS_PROVIDER_STATE_UNKNOWN  UpdateVirtualCircuitDetailsProviderStateEnum = "UNKNOWN"
)

var mapping_updatevirtualcircuitdetails_providerState = map[string]UpdateVirtualCircuitDetailsProviderStateEnum{
	"ACTIVE":   UPDATE_VIRTUAL_CIRCUIT_DETAILS_PROVIDER_STATE_ACTIVE,
	"INACTIVE": UPDATE_VIRTUAL_CIRCUIT_DETAILS_PROVIDER_STATE_INACTIVE,
	"UNKNOWN":  UPDATE_VIRTUAL_CIRCUIT_DETAILS_PROVIDER_STATE_UNKNOWN,
}

func (receiver UpdateVirtualCircuitDetailsProviderState) Values() []UpdateVirtualCircuitDetailsProviderStateEnum {
	values := make([]UpdateVirtualCircuitDetailsProviderStateEnum, 0)
	for _, v := range mapping_updatevirtualcircuitdetails_providerState {
		if v != UPDATE_VIRTUAL_CIRCUIT_DETAILS_PROVIDER_STATE_UNKNOWN {
			values = append(values, v)
		}
	}
	return values
}

func (receiver UpdateVirtualCircuitDetailsProviderState) IsValid(toBeChecked string) bool {
	for _, v := range receiver.Values() {
		if UpdateVirtualCircuitDetailsProviderStateEnum(toBeChecked) == v {
			return true
		}
	}
	return false
}

func (receiver UpdateVirtualCircuitDetailsProviderState) From(toBeConverted string) UpdateVirtualCircuitDetailsProviderStateEnum {
	if val, ok := mapping_updatevirtualcircuitdetails_providerState[toBeConverted]; ok {
		return val
	}
	return UPDATE_VIRTUAL_CIRCUIT_DETAILS_PROVIDER_STATE_UNKNOWN
}
