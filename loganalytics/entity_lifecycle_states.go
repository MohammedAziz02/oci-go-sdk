// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

// EntityLifecycleStatesEnum Enum with underlying type: string
type EntityLifecycleStatesEnum string

// Set of constants representing the allowable values for EntityLifecycleStatesEnum
const (
	EntityLifecycleStatesActive  EntityLifecycleStatesEnum = "ACTIVE"
	EntityLifecycleStatesDeleted EntityLifecycleStatesEnum = "DELETED"
)

var mappingEntityLifecycleStates = map[string]EntityLifecycleStatesEnum{
	"ACTIVE":  EntityLifecycleStatesActive,
	"DELETED": EntityLifecycleStatesDeleted,
}

// GetEntityLifecycleStatesEnumValues Enumerates the set of values for EntityLifecycleStatesEnum
func GetEntityLifecycleStatesEnumValues() []EntityLifecycleStatesEnum {
	values := make([]EntityLifecycleStatesEnum, 0)
	for _, v := range mappingEntityLifecycleStates {
		values = append(values, v)
	}
	return values
}
