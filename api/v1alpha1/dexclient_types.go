/*
Copyright 2021.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DexClientSpec defines the desired state of DexClient
type DexClientSpec struct {
	// +kubebuilder:validation:MinLength=4
	// The name of the oidc config
	ClientID string `json:"clientID,omitempty"`
	// +kubebuilder:validation:Required
	// The shared oidc secret
	ClientSecretRef corev1.SecretReference `json:"clientSecretRef,omitempty"`
	// +optional
	// Sets the public flag
	Public bool `json:"public,omitempty"`
	// Redirect URIs
	RedirectURIs []string `json:"redirectURIs,omitempty"`
	// +optional
	// Trusted Peers
	TrustedPeers []string `json:"trustedPeers,omitempty"`
	// +optional
	// LogoURL
	LogoURL string `json:"logoURL,omitempty"`
}

const (
	DexClientConditionTypeApplied             string = "Applied"
	DexClientConditionTypeOAuth2ClientCreated string = "OAuth2ClientCreated"
)

// DexClientStatus defines the observed state of DexClient
type DexClientStatus struct {
	// +optional
	RelatedObjects []RelatedObjectReference `json:"relatedObjects,omitempty"`
	// Conditions contains the different condition statuses for this DexClient.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DexClient is the Schema for the dexclients API
type DexClient struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DexClientSpec   `json:"spec,omitempty"`
	Status DexClientStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DexClientList contains a list of DexClient
type DexClientList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DexClient `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DexClient{}, &DexClientList{})
}
