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

type UserAssignedIdentityObservation struct {

	// The ID of the app associated with the Identity.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// The ID of the User Assigned Identity.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Service Principal object associated with the created Identity.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The ID of the Tenant which the Identity belongs to.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type UserAssignedIdentityParameters struct {

	// The Azure Region where the User Assigned Identity should exist. Changing this forces a new User Assigned Identity to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Specifies the name of this User Assigned Identity. Changing this forces a new User Assigned Identity to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the name of the Resource Group within which this User Assigned Identity should exist. Changing this forces a new User Assigned Identity to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the User Assigned Identity.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// UserAssignedIdentitySpec defines the desired state of UserAssignedIdentity
type UserAssignedIdentitySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserAssignedIdentityParameters `json:"forProvider"`
}

// UserAssignedIdentityStatus defines the observed state of UserAssignedIdentity.
type UserAssignedIdentityStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserAssignedIdentityObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// UserAssignedIdentity is the Schema for the UserAssignedIdentitys API. Manages a User Assigned Identity.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type UserAssignedIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserAssignedIdentitySpec   `json:"spec"`
	Status            UserAssignedIdentityStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserAssignedIdentityList contains a list of UserAssignedIdentitys
type UserAssignedIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserAssignedIdentity `json:"items"`
}

// Repository type metadata.
var (
	UserAssignedIdentity_Kind             = "UserAssignedIdentity"
	UserAssignedIdentity_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserAssignedIdentity_Kind}.String()
	UserAssignedIdentity_KindAPIVersion   = UserAssignedIdentity_Kind + "." + CRDGroupVersion.String()
	UserAssignedIdentity_GroupVersionKind = CRDGroupVersion.WithKind(UserAssignedIdentity_Kind)
)

func init() {
	SchemeBuilder.Register(&UserAssignedIdentity{}, &UserAssignedIdentityList{})
}
