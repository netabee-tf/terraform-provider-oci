// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"strings"
)

// FilterPush The information about a filter operator. The filter operator lets you select certain attributes from the inbound port to continue downstream to the outbound port.
type FilterPush struct {

	// The filter condition.
	FilterCondition *string `mandatory:"false" json:"filterCondition"`
}

func (m FilterPush) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FilterPush) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FilterPush) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFilterPush FilterPush
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeFilterPush
	}{
		"FILTER",
		(MarshalTypeFilterPush)(m),
	}

	return json.Marshal(&s)
}
