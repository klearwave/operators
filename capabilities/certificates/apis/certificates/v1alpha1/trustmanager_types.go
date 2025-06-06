/*
Copyright 2025.

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
	"errors"

	"github.com/nukleros/operator-builder-tools/pkg/controller/workload"
	"github.com/nukleros/operator-builder-tools/pkg/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var ErrUnableToConvertTrustManager = errors.New("unable to convert to TrustManager")

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TrustManagerSpec defines the desired state of TrustManager.
type TrustManagerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:default="klearwave-certificates-system"
	// +kubebuilder:validation:Optional
	// (Default: "klearwave-certificates-system")
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Controller TrustManagerSpecController `json:"controller,omitempty"`
}

type TrustManagerSpecController struct {
	// +kubebuilder:default=1
	// +kubebuilder:validation:Optional
	// (Default: 1)
	//
	//	Number of replicas to use for the trust-manager controller deployment.
	Replicas int `json:"replicas,omitempty"`

	// +kubebuilder:default="quay.io/jetstack/trust-manager:v0.17.1"
	// +kubebuilder:validation:Optional
	// (Default: "quay.io/jetstack/trust-manager:v0.17.1")
	//
	//	Image to use for trust-manager controller deployment.
	Image string `json:"image,omitempty"`

	// +kubebuilder:validation:Optional
	Resources TrustManagerSpecControllerResources `json:"resources,omitempty"`
}

type TrustManagerSpecControllerResources struct {
	// +kubebuilder:validation:Optional
	Requests TrustManagerSpecControllerResourcesRequests `json:"requests,omitempty"`

	// +kubebuilder:validation:Optional
	Limits TrustManagerSpecControllerResourcesLimits `json:"limits,omitempty"`
}

type TrustManagerSpecControllerResourcesRequests struct {
	// +kubebuilder:default="25m"
	// +kubebuilder:validation:Optional
	// (Default: "25m")
	//
	//	CPU requests to use for trust-manager controller deployment.
	Cpu string `json:"cpu,omitempty"`

	// +kubebuilder:default="32Mi"
	// +kubebuilder:validation:Optional
	// (Default: "32Mi")
	//
	//	Memory requests to use for trust-manager controller deployment.
	Memory string `json:"memory,omitempty"`
}

type TrustManagerSpecControllerResourcesLimits struct {
	// +kubebuilder:default="64Mi"
	// +kubebuilder:validation:Optional
	// (Default: "64Mi")
	//
	//	Memory limits to use for trust-manager controller deployment.
	Memory string `json:"memory,omitempty"`
}

// TrustManagerStatus defines the observed state of TrustManager.
type TrustManagerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Created               bool                     `json:"created,omitempty"`
	DependenciesSatisfied bool                     `json:"dependenciesSatisfied,omitempty"`
	Conditions            []*status.PhaseCondition `json:"conditions,omitempty"`
	Resources             []*status.ChildResource  `json:"resources,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// TrustManager is the Schema for the trustmanagers API.
type TrustManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrustManagerSpec   `json:"spec,omitempty"`
	Status            TrustManagerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrustManagerList contains a list of TrustManager.
type TrustManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrustManager `json:"items"`
}

// interface methods

// GetReadyStatus returns the ready status for a component.
func (component *TrustManager) GetReadyStatus() bool {
	return component.Status.Created
}

// SetReadyStatus sets the ready status for a component.
func (component *TrustManager) SetReadyStatus(ready bool) {
	component.Status.Created = ready
}

// GetDependencyStatus returns the dependency status for a component.
func (component *TrustManager) GetDependencyStatus() bool {
	return component.Status.DependenciesSatisfied
}

// SetDependencyStatus sets the dependency status for a component.
func (component *TrustManager) SetDependencyStatus(dependencyStatus bool) {
	component.Status.DependenciesSatisfied = dependencyStatus
}

// GetPhaseConditions returns the phase conditions for a component.
func (component *TrustManager) GetPhaseConditions() []*status.PhaseCondition {
	return component.Status.Conditions
}

// SetPhaseCondition sets the phase conditions for a component.
func (component *TrustManager) SetPhaseCondition(condition *status.PhaseCondition) {
	for i, currentCondition := range component.GetPhaseConditions() {
		if currentCondition.Phase == condition.Phase {
			component.Status.Conditions[i] = condition

			return
		}
	}

	// phase not found, lets add it to the list.
	component.Status.Conditions = append(component.Status.Conditions, condition)
}

// GetResources returns the child resource status for a component.
func (component *TrustManager) GetChildResourceConditions() []*status.ChildResource {
	return component.Status.Resources
}

// SetResources sets the phase conditions for a component.
func (component *TrustManager) SetChildResourceCondition(resource *status.ChildResource) {
	for i, currentResource := range component.GetChildResourceConditions() {
		if currentResource.Group == resource.Group && currentResource.Version == resource.Version && currentResource.Kind == resource.Kind {
			if currentResource.Name == resource.Name && currentResource.Namespace == resource.Namespace {
				component.Status.Resources[i] = resource

				return
			}
		}
	}

	// phase not found, lets add it to the collection
	component.Status.Resources = append(component.Status.Resources, resource)
}

// GetDependencies returns the dependencies for a component.
func (*TrustManager) GetDependencies() []workload.Workload {
	return []workload.Workload{}
}

// GetComponentGVK returns a GVK object for the component.
func (*TrustManager) GetWorkloadGVK() schema.GroupVersionKind {
	return GroupVersion.WithKind("TrustManager")
}

func init() {
	SchemeBuilder.Register(&TrustManager{}, &TrustManagerList{})
}
