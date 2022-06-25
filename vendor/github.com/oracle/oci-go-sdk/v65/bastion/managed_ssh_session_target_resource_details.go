// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Bastion API
//
// Use the Bastion API to provide restricted and time-limited access to target resources that don't have public endpoints. Bastions let authorized users connect from specific IP addresses to target resources using Secure Shell (SSH) sessions. For more information, see the Bastion documentation (https://docs.cloud.oracle.com/iaas/Content/Bastion/home.htm).
//

package bastion

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ManagedSshSessionTargetResourceDetails Details about a managed SSH session for a target resource.
type ManagedSshSessionTargetResourceDetails struct {

	// The name of the user on the target resource operating system that the session uses for the connection.
	TargetResourceOperatingSystemUserName *string `mandatory:"true" json:"targetResourceOperatingSystemUserName"`

	// The unique identifier (OCID) of the target resource (a Compute instance, for example) that the session connects to.
	TargetResourceId *string `mandatory:"true" json:"targetResourceId"`

	// The display name of the target Compute instance that the session connects to.
	TargetResourceDisplayName *string `mandatory:"true" json:"targetResourceDisplayName"`

	// The port number to connect to on the target resource.
	TargetResourcePort *int `mandatory:"true" json:"targetResourcePort"`

	// The private IP address of the target resource that the session connects to.
	TargetResourcePrivateIpAddress *string `mandatory:"false" json:"targetResourcePrivateIpAddress"`
}

func (m ManagedSshSessionTargetResourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedSshSessionTargetResourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ManagedSshSessionTargetResourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeManagedSshSessionTargetResourceDetails ManagedSshSessionTargetResourceDetails
	s := struct {
		DiscriminatorParam string `json:"sessionType"`
		MarshalTypeManagedSshSessionTargetResourceDetails
	}{
		"MANAGED_SSH",
		(MarshalTypeManagedSshSessionTargetResourceDetails)(m),
	}

	return json.Marshal(&s)
}
