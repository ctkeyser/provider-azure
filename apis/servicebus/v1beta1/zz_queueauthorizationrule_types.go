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

type QueueAuthorizationRuleObservation struct {

	// The ID of the Authorization Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type QueueAuthorizationRuleParameters struct {

	// Does this Authorization Rule have Listen permissions to the ServiceBus Queue? Defaults to false.
	// +kubebuilder:validation:Optional
	Listen *bool `json:"listen,omitempty" tf:"listen,omitempty"`

	// Does this Authorization Rule have Manage permissions to the ServiceBus Queue? When this property is true - both listen and send must be too. Defaults to false.
	// +kubebuilder:validation:Optional
	Manage *bool `json:"manage,omitempty" tf:"manage,omitempty"`

	// Specifies the ID of the ServiceBus Queue. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/servicebus/v1beta1.Queue
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	QueueID *string `json:"queueId,omitempty" tf:"queue_id,omitempty"`

	// Reference to a Queue in servicebus to populate queueId.
	// +kubebuilder:validation:Optional
	QueueIDRef *v1.Reference `json:"queueIdRef,omitempty" tf:"-"`

	// Selector for a Queue in servicebus to populate queueId.
	// +kubebuilder:validation:Optional
	QueueIDSelector *v1.Selector `json:"queueIdSelector,omitempty" tf:"-"`

	// Does this Authorization Rule have Send permissions to the ServiceBus Queue? Defaults to false.
	// +kubebuilder:validation:Optional
	Send *bool `json:"send,omitempty" tf:"send,omitempty"`
}

// QueueAuthorizationRuleSpec defines the desired state of QueueAuthorizationRule
type QueueAuthorizationRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueueAuthorizationRuleParameters `json:"forProvider"`
}

// QueueAuthorizationRuleStatus defines the observed state of QueueAuthorizationRule.
type QueueAuthorizationRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueueAuthorizationRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// QueueAuthorizationRule is the Schema for the QueueAuthorizationRules API. Manages an Authorization Rule for a ServiceBus Queue.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type QueueAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QueueAuthorizationRuleSpec   `json:"spec"`
	Status            QueueAuthorizationRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueueAuthorizationRuleList contains a list of QueueAuthorizationRules
type QueueAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []QueueAuthorizationRule `json:"items"`
}

// Repository type metadata.
var (
	QueueAuthorizationRule_Kind             = "QueueAuthorizationRule"
	QueueAuthorizationRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: QueueAuthorizationRule_Kind}.String()
	QueueAuthorizationRule_KindAPIVersion   = QueueAuthorizationRule_Kind + "." + CRDGroupVersion.String()
	QueueAuthorizationRule_GroupVersionKind = CRDGroupVersion.WithKind(QueueAuthorizationRule_Kind)
)

func init() {
	SchemeBuilder.Register(&QueueAuthorizationRule{}, &QueueAuthorizationRuleList{})
}
