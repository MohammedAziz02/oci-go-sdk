// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// UpdatePipelineDetails Properties used in pipeline update operations
type UpdatePipelineDetails struct {

	// Generated key that can be used in API calls to identify pipeline. On scenarios where reference to the pipeline is needed, a value can be passed in create.
	Key *string `mandatory:"true" json:"key"`

	// The type of the object.
	ModelType *string `mandatory:"true" json:"modelType"`

	// This is used by the service for optimistic locking of the object, to prevent multiple users from simultaneously updating the object.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// This is a version number that is used by the service to upgrade objects if needed through releases of the service.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// A list of nodes attached to the pipeline
	Nodes []FlowNode `mandatory:"false" json:"nodes"`

	// A list of additional parameters required in pipeline.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	FlowConfigValues *ConfigValues `mandatory:"false" json:"flowConfigValues"`

	// The list of variables required in pipeline.
	Variables []Variable `mandatory:"false" json:"variables"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
}

func (m UpdatePipelineDetails) String() string {
	return common.PointerString(m)
}
