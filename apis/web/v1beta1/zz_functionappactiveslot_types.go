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

type FunctionAppActiveSlotObservation struct {

	// The ID of the Function App Active Slot
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The timestamp of the last successful swap with Production
	// The timestamp of the last successful swap with `Production`
	LastSuccessfulSwap *string `json:"lastSuccessfulSwap,omitempty" tf:"last_successful_swap,omitempty"`
}

type FunctionAppActiveSlotParameters struct {

	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to true. Changing this forces a new resource to be created.
	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to `true`.
	// +kubebuilder:validation:Optional
	OverwriteNetworkConfig *bool `json:"overwriteNetworkConfig,omitempty" tf:"overwrite_network_config,omitempty"`

	// The ID of the Slot to swap with Production.
	// The ID of the Slot to swap with `Production`.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/web/v1beta1.WindowsFunctionAppSlot
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SlotID *string `json:"slotId,omitempty" tf:"slot_id,omitempty"`

	// Reference to a WindowsFunctionAppSlot in web to populate slotId.
	// +kubebuilder:validation:Optional
	SlotIDRef *v1.Reference `json:"slotIdRef,omitempty" tf:"-"`

	// Selector for a WindowsFunctionAppSlot in web to populate slotId.
	// +kubebuilder:validation:Optional
	SlotIDSelector *v1.Selector `json:"slotIdSelector,omitempty" tf:"-"`
}

// FunctionAppActiveSlotSpec defines the desired state of FunctionAppActiveSlot
type FunctionAppActiveSlotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionAppActiveSlotParameters `json:"forProvider"`
}

// FunctionAppActiveSlotStatus defines the observed state of FunctionAppActiveSlot.
type FunctionAppActiveSlotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionAppActiveSlotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionAppActiveSlot is the Schema for the FunctionAppActiveSlots API. Manages a Function App Active Slot.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FunctionAppActiveSlot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionAppActiveSlotSpec   `json:"spec"`
	Status            FunctionAppActiveSlotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionAppActiveSlotList contains a list of FunctionAppActiveSlots
type FunctionAppActiveSlotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunctionAppActiveSlot `json:"items"`
}

// Repository type metadata.
var (
	FunctionAppActiveSlot_Kind             = "FunctionAppActiveSlot"
	FunctionAppActiveSlot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionAppActiveSlot_Kind}.String()
	FunctionAppActiveSlot_KindAPIVersion   = FunctionAppActiveSlot_Kind + "." + CRDGroupVersion.String()
	FunctionAppActiveSlot_GroupVersionKind = CRDGroupVersion.WithKind(FunctionAppActiveSlot_Kind)
)

func init() {
	SchemeBuilder.Register(&FunctionAppActiveSlot{}, &FunctionAppActiveSlotList{})
}
