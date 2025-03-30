/*
Copyright 2022 The Crossplane Authors.

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

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// WebhookParameters are the configurable fields of a Webhook.
type WebhookParameters struct {
	// Repository is the
	Repository *string

	// Owner is the owner or organization of the repository
	Owner *string

	// TODO: move the token to a secret
	// Token is the github token used to register the webhook
	Token *string

	// URL is the target of the webhook
	URL *string `json:"url"`
}

// WebhookObservation represents a project hook.
type WebhookObservation struct {
	// ID of the project hook at github
	ID int `json:"id,omitempty"`

	// CreatedAt specifies the time the project hook was created
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
}

// A WebhookSpec defines the desired state of a Webhook.
type WebhookSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WebhookParameters `json:"forProvider"`
}

// A WebhookStatus represents the observed state of a Webhook.
type WebhookStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A Webhook is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,provider-github}
type Webhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WebhookSpec   `json:"spec"`
	Status WebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WebhookList contains a list of Webhook
type WebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Webhook `json:"items"`
}
