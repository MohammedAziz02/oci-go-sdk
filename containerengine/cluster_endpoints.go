// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Engine for Kubernetes API
//
// API for the Container Engine for Kubernetes service. Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Container Engine for Kubernetes (https://docs.cloud.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"github.com/oracle/oci-go-sdk/v42/common"
)

// ClusterEndpoints The properties that define endpoints for a cluster.
type ClusterEndpoints struct {

	// The non-native networking Kubernetes API server endpoint.
	Kubernetes *string `mandatory:"false" json:"kubernetes"`

	// The public native networking Kubernetes API server endpoint, if one was requested.
	PublicEndpoint *string `mandatory:"false" json:"publicEndpoint"`

	// The private native networking Kubernetes API server endpoint.
	PrivateEndpoint *string `mandatory:"false" json:"privateEndpoint"`
}

func (m ClusterEndpoints) String() string {
	return common.PointerString(m)
}
