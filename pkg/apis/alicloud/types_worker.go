// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package alicloud

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkerStatus contains information about created worker resources.
type WorkerStatus struct {
	metav1.TypeMeta

	// MachineImages is a list of machine images that have been used in this worker. Usually, the extension controller
	// gets the mapping from name/version to the provider-specific machine image data in its componentconfig. However, if
	// a version that is still in use gets removed from this componentconfig it cannot reconcile anymore existing `Worker`
	// resources that are still using this version. Hence, it stores the used versions in the provider status to ensure
	// reconciliation is possible.
	MachineImages []MachineImage
}

// MachineImage is a mapping from logical names and versions to provider-specific machine image data.
type MachineImage struct {
	// Name is the logical name of the machine image.
	Name string
	// Version is the logical version of the machine image.
	Version string
	// ID is the ID of the image.
	ID string
	// Encrypted is a flag to specify whether this image is encrypted or not
	Encrypted *bool
}
