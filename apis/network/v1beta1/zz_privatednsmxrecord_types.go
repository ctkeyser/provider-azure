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

type PrivateDNSMXRecordObservation struct {

	// The FQDN of the DNS MX Record.
	Fqdn *string `json:"fqdn,omitempty" tf:"fqdn,omitempty"`

	// The Private DNS MX Record ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type PrivateDNSMXRecordParameters struct {

	// One or more record blocks as defined below.
	// +kubebuilder:validation:Required
	Record []PrivateDNSMXRecordRecordParameters `json:"record" tf:"record,omitempty"`

	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The Time To Live (TTL) of the DNS record in seconds.
	// +kubebuilder:validation:Required
	TTL *float64 `json:"ttl" tf:"ttl,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the Private DNS Zone where the resource exists. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=PrivateDNSZone
	// +kubebuilder:validation:Optional
	ZoneName *string `json:"zoneName,omitempty" tf:"zone_name,omitempty"`

	// Reference to a PrivateDNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameRef *v1.Reference `json:"zoneNameRef,omitempty" tf:"-"`

	// Selector for a PrivateDNSZone to populate zoneName.
	// +kubebuilder:validation:Optional
	ZoneNameSelector *v1.Selector `json:"zoneNameSelector,omitempty" tf:"-"`
}

type PrivateDNSMXRecordRecordObservation struct {
}

type PrivateDNSMXRecordRecordParameters struct {

	// The FQDN of the exchange to MX record points to.
	// +kubebuilder:validation:Required
	Exchange *string `json:"exchange" tf:"exchange,omitempty"`

	// The preference of the MX record.
	// +kubebuilder:validation:Required
	Preference *float64 `json:"preference" tf:"preference,omitempty"`
}

// PrivateDNSMXRecordSpec defines the desired state of PrivateDNSMXRecord
type PrivateDNSMXRecordSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSMXRecordParameters `json:"forProvider"`
}

// PrivateDNSMXRecordStatus defines the observed state of PrivateDNSMXRecord.
type PrivateDNSMXRecordStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSMXRecordObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSMXRecord is the Schema for the PrivateDNSMXRecords API. Manages a Private DNS MX Record.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateDNSMXRecord struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDNSMXRecordSpec   `json:"spec"`
	Status            PrivateDNSMXRecordStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSMXRecordList contains a list of PrivateDNSMXRecords
type PrivateDNSMXRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSMXRecord `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSMXRecord_Kind             = "PrivateDNSMXRecord"
	PrivateDNSMXRecord_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSMXRecord_Kind}.String()
	PrivateDNSMXRecord_KindAPIVersion   = PrivateDNSMXRecord_Kind + "." + CRDGroupVersion.String()
	PrivateDNSMXRecord_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSMXRecord_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSMXRecord{}, &PrivateDNSMXRecordList{})
}
