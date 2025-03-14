/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProviderObservation struct {

	// The URI of the Attestation Service.
	AttestationURI *string `json:"attestationUri,omitempty" tf:"attestation_uri,omitempty"`

	// The ID of the Attestation Provider.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Trust model used for the Attestation Service.
	TrustModel *string `json:"trustModel,omitempty" tf:"trust_model,omitempty"`
}

type ProviderParameters struct {

	// The Azure Region where the Attestation Provider should exist. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// A valid X.509 certificate (Section 4 of RFC4648). Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	PolicySigningCertificateData *string `json:"policySigningCertificateData,omitempty" tf:"policy_signing_certificate_data,omitempty"`

	// The name of the Resource Group where the attestation provider should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Attestation Provider.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// ProviderSpec defines the desired state of Provider
type ProviderSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProviderParameters `json:"forProvider"`
}

// ProviderStatus defines the observed state of Provider.
type ProviderStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Provider is the Schema for the Providers API. Manages a Attestation Provider.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Provider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProviderSpec   `json:"spec"`
	Status            ProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProviderList contains a list of Providers
type ProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Provider `json:"items"`
}

// Repository type metadata.
var (
	Provider_Kind             = "Provider"
	Provider_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Provider_Kind}.String()
	Provider_KindAPIVersion   = Provider_Kind + "." + CRDGroupVersion.String()
	Provider_GroupVersionKind = CRDGroupVersion.WithKind(Provider_Kind)
)

func init() {
	SchemeBuilder.Register(&Provider{}, &ProviderList{})
}
