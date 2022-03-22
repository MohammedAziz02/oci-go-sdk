// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
)

// SummarizeHostInsightResourceUtilizationInsightAggregation Insights response containing current/projected groups for CPU or memory.
type SummarizeHostInsightResourceUtilizationInsightAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (CPU, Physical Memory, Logical Memory)
	ResourceMetric SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	ProjectedUtilization *ResourceInsightProjectedUtilization `mandatory:"true" json:"projectedUtilization"`

	CurrentUtilization *ResourceInsightCurrentUtilization `mandatory:"true" json:"currentUtilization"`
}

func (m SummarizeHostInsightResourceUtilizationInsightAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeHostInsightResourceUtilizationInsightAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum Enum with underlying type: string
type SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum
const (
	SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricCpu           SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum = "CPU"
	SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricMemory        SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum = "MEMORY"
	SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricLogicalMemory SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum = "LOGICAL_MEMORY"
)

var mappingSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum = map[string]SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum{
	"CPU":            SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricCpu,
	"MEMORY":         SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricMemory,
	"LOGICAL_MEMORY": SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricLogicalMemory,
}

var mappingSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnumLowerCase = map[string]SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum{
	"cpu":            SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricCpu,
	"memory":         SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricMemory,
	"logical_memory": SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricLogicalMemory,
}

// GetSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnumValues Enumerates the set of values for SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum
func GetSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnumValues() []SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum {
	values := make([]SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum
func GetSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"MEMORY",
		"LOGICAL_MEMORY",
	}
}

// GetMappingSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum(val string) (SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
