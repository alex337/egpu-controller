/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EGPU is a specification for a EGPU resource
type EGPU struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EGPUSpec   `json:"spec"`
	Status EGPUStatus `json:"status"`
}

// EGPUSpec is the spec for a EGPU resource
type EGPUSpec struct {
	NodeName string `json:"nodeName"`
	GPU []string `json:"gpu"`
	Resources EGPUResource `json:"resources"`
}

type EGPUResource struct {
	Capacity EGPUCapacity `json:"capacity"`
}

type EGPUCapacity struct {
	QGPUCore string `json:"QGPUCore"`
	QGPUMemory string `json:"QGPUMemory"`
}

// EGPUStatus is the status for a EGPU resource
type EGPUStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EGPUList is a list of EGPU resources
type EGPUList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EGPU `json:"items"`
}
