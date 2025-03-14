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

type LiveEventOutputObservation struct {

	// The ID of the Live Output.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type LiveEventOutputParameters struct {

	// ISO 8601 time between 1 minute to 25 hours to indicate the maximum content length that can be archived in the asset for this live output. This also sets the maximum content length for the rewind window. For example, use PT1H30M to indicate 1 hour and 30 minutes of archive window. Changing this forces a new Live Output to be created.
	// +kubebuilder:validation:Required
	ArchiveWindowDuration *string `json:"archiveWindowDuration" tf:"archive_window_duration,omitempty"`

	// The asset that the live output will write to. Changing this forces a new Live Output to be created.
	// +crossplane:generate:reference:type=Asset
	// +kubebuilder:validation:Optional
	AssetName *string `json:"assetName,omitempty" tf:"asset_name,omitempty"`

	// Reference to a Asset to populate assetName.
	// +kubebuilder:validation:Optional
	AssetNameRef *v1.Reference `json:"assetNameRef,omitempty" tf:"-"`

	// Selector for a Asset to populate assetName.
	// +kubebuilder:validation:Optional
	AssetNameSelector *v1.Selector `json:"assetNameSelector,omitempty" tf:"-"`

	// The description of the live output. Changing this forces a new Live Output to be created.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The number of fragments in an HTTP Live Streaming (HLS) TS segment in the output of the live event. This value does not affect the packing ratio for HLS CMAF output. Changing this forces a new Live Output to be created.
	// +kubebuilder:validation:Optional
	HlsFragmentsPerTSSegment *float64 `json:"hlsFragmentsPerTsSegment,omitempty" tf:"hls_fragments_per_ts_segment,omitempty"`

	// The id of the live event. Changing this forces a new Live Output to be created.
	// +crossplane:generate:reference:type=LiveEvent
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LiveEventID *string `json:"liveEventId,omitempty" tf:"live_event_id,omitempty"`

	// Reference to a LiveEvent to populate liveEventId.
	// +kubebuilder:validation:Optional
	LiveEventIDRef *v1.Reference `json:"liveEventIdRef,omitempty" tf:"-"`

	// Selector for a LiveEvent to populate liveEventId.
	// +kubebuilder:validation:Optional
	LiveEventIDSelector *v1.Selector `json:"liveEventIdSelector,omitempty" tf:"-"`

	// The manifest file name. If not provided, the service will generate one automatically. Changing this forces a new Live Output to be created.
	// +kubebuilder:validation:Optional
	ManifestName *string `json:"manifestName,omitempty" tf:"manifest_name,omitempty"`

	// The initial timestamp that the live output will start at, any content before this value will not be archived. Changing this forces a new Live Output to be created.
	// +kubebuilder:validation:Optional
	OutputSnapTimeInSeconds *float64 `json:"outputSnapTimeInSeconds,omitempty" tf:"output_snap_time_in_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	RewindWindowDuration *string `json:"rewindWindowDuration,omitempty" tf:"rewind_window_duration,omitempty"`
}

// LiveEventOutputSpec defines the desired state of LiveEventOutput
type LiveEventOutputSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LiveEventOutputParameters `json:"forProvider"`
}

// LiveEventOutputStatus defines the observed state of LiveEventOutput.
type LiveEventOutputStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LiveEventOutputObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LiveEventOutput is the Schema for the LiveEventOutputs API. Manages an Azure Media Live Event Output.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LiveEventOutput struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LiveEventOutputSpec   `json:"spec"`
	Status            LiveEventOutputStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LiveEventOutputList contains a list of LiveEventOutputs
type LiveEventOutputList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LiveEventOutput `json:"items"`
}

// Repository type metadata.
var (
	LiveEventOutput_Kind             = "LiveEventOutput"
	LiveEventOutput_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LiveEventOutput_Kind}.String()
	LiveEventOutput_KindAPIVersion   = LiveEventOutput_Kind + "." + CRDGroupVersion.String()
	LiveEventOutput_GroupVersionKind = CRDGroupVersion.WithKind(LiveEventOutput_Kind)
)

func init() {
	SchemeBuilder.Register(&LiveEventOutput{}, &LiveEventOutputList{})
}
