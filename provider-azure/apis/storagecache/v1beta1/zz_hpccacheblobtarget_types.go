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

type HPCCacheBlobTargetObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type HPCCacheBlobTargetParameters struct {

	// +kubebuilder:validation:Optional
	AccessPolicyName *string `json:"accessPolicyName,omitempty" tf:"access_policy_name,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storagecache/v1beta1.HPCCache
	// +kubebuilder:validation:Optional
	CacheName *string `json:"cacheName,omitempty" tf:"cache_name,omitempty"`

	// +kubebuilder:validation:Optional
	CacheNameRef *v1.Reference `json:"cacheNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	CacheNameSelector *v1.Selector `json:"cacheNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	NamespacePath *string `json:"namespacePath" tf:"namespace_path,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1.Container
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("resource_manager_id",true)
	// +kubebuilder:validation:Optional
	StorageContainerID *string `json:"storageContainerId,omitempty" tf:"storage_container_id,omitempty"`

	// +kubebuilder:validation:Optional
	StorageContainerIDRef *v1.Reference `json:"storageContainerIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	StorageContainerIDSelector *v1.Selector `json:"storageContainerIdSelector,omitempty" tf:"-"`
}

// HPCCacheBlobTargetSpec defines the desired state of HPCCacheBlobTarget
type HPCCacheBlobTargetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HPCCacheBlobTargetParameters `json:"forProvider"`
}

// HPCCacheBlobTargetStatus defines the observed state of HPCCacheBlobTarget.
type HPCCacheBlobTargetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HPCCacheBlobTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCacheBlobTarget is the Schema for the HPCCacheBlobTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type HPCCacheBlobTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HPCCacheBlobTargetSpec   `json:"spec"`
	Status            HPCCacheBlobTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HPCCacheBlobTargetList contains a list of HPCCacheBlobTargets
type HPCCacheBlobTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HPCCacheBlobTarget `json:"items"`
}

// Repository type metadata.
var (
	HPCCacheBlobTarget_Kind             = "HPCCacheBlobTarget"
	HPCCacheBlobTarget_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HPCCacheBlobTarget_Kind}.String()
	HPCCacheBlobTarget_KindAPIVersion   = HPCCacheBlobTarget_Kind + "." + CRDGroupVersion.String()
	HPCCacheBlobTarget_GroupVersionKind = CRDGroupVersion.WithKind(HPCCacheBlobTarget_Kind)
)

func init() {
	SchemeBuilder.Register(&HPCCacheBlobTarget{}, &HPCCacheBlobTargetList{})
}
