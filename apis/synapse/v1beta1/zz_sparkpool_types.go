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

type AutoPauseObservation struct {
}

type AutoPauseParameters struct {

	// Number of minutes of idle time before the Spark Pool is automatically paused. Must be between 5 and 10080.
	// +kubebuilder:validation:Required
	DelayInMinutes *float64 `json:"delayInMinutes" tf:"delay_in_minutes,omitempty"`
}

type AutoScaleObservation struct {
}

type AutoScaleParameters struct {

	// The maximum number of nodes the Spark Pool can support. Must be between 3 and 200.
	// +kubebuilder:validation:Required
	MaxNodeCount *float64 `json:"maxNodeCount" tf:"max_node_count,omitempty"`

	// The minimum number of nodes the Spark Pool can support. Must be between 3 and 200.
	// +kubebuilder:validation:Required
	MinNodeCount *float64 `json:"minNodeCount" tf:"min_node_count,omitempty"`
}

type LibraryRequirementObservation struct {
}

type LibraryRequirementParameters struct {

	// The content of library requirements.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// The name of the library requirements file.
	// +kubebuilder:validation:Required
	Filename *string `json:"filename" tf:"filename,omitempty"`
}

type SparkConfigObservation struct {
}

type SparkConfigParameters struct {

	// The contents of a spark configuration.
	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// The name of the file where the spark configuration content will be stored.
	// +kubebuilder:validation:Required
	Filename *string `json:"filename" tf:"filename,omitempty"`
}

type SparkPoolObservation struct {

	// The ID of the Synapse Spark Pool.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SparkPoolParameters struct {

	// An auto_pause block as defined below.
	// +kubebuilder:validation:Optional
	AutoPause []AutoPauseParameters `json:"autoPause,omitempty" tf:"auto_pause,omitempty"`

	// An auto_scale block as defined below. Exactly one of node_count or auto_scale must be specified.
	// +kubebuilder:validation:Optional
	AutoScale []AutoScaleParameters `json:"autoScale,omitempty" tf:"auto_scale,omitempty"`

	// The cache size in the Spark Pool.
	// +kubebuilder:validation:Optional
	CacheSize *float64 `json:"cacheSize,omitempty" tf:"cache_size,omitempty"`

	// Indicates whether compute isolation is enabled or not. Defaults to false.
	// +kubebuilder:validation:Optional
	ComputeIsolationEnabled *bool `json:"computeIsolationEnabled,omitempty" tf:"compute_isolation_enabled,omitempty"`

	// Indicates whether Dynamic Executor Allocation is enabled or not. Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicExecutorAllocationEnabled *bool `json:"dynamicExecutorAllocationEnabled,omitempty" tf:"dynamic_executor_allocation_enabled,omitempty"`

	// A library_requirement block as defined below.
	// +kubebuilder:validation:Optional
	LibraryRequirement []LibraryRequirementParameters `json:"libraryRequirement,omitempty" tf:"library_requirement,omitempty"`

	// The maximum number of executors allocated only when dynamic_executor_allocation_enabled set to true.
	// +kubebuilder:validation:Optional
	MaxExecutors *float64 `json:"maxExecutors,omitempty" tf:"max_executors,omitempty"`

	// The minimum number of executors allocated only when dynamic_executor_allocation_enabled set to true.
	// +kubebuilder:validation:Optional
	MinExecutors *float64 `json:"minExecutors,omitempty" tf:"min_executors,omitempty"`

	// The number of nodes in the Spark Pool. Exactly one of node_count or auto_scale must be specified.
	// +kubebuilder:validation:Optional
	NodeCount *float64 `json:"nodeCount,omitempty" tf:"node_count,omitempty"`

	// The level of node in the Spark Pool. Possible values are Small, Medium, Large, None, XLarge, XXLarge and XXXLarge.
	// +kubebuilder:validation:Required
	NodeSize *string `json:"nodeSize" tf:"node_size,omitempty"`

	// The kind of nodes that the Spark Pool provides. Possible values are MemoryOptimized and None.
	// +kubebuilder:validation:Required
	NodeSizeFamily *string `json:"nodeSizeFamily" tf:"node_size_family,omitempty"`

	// Indicates whether session level packages are enabled or not. Defaults to false.
	// +kubebuilder:validation:Optional
	SessionLevelPackagesEnabled *bool `json:"sessionLevelPackagesEnabled,omitempty" tf:"session_level_packages_enabled,omitempty"`

	// A spark_config block as defined below.
	// +kubebuilder:validation:Optional
	SparkConfig []SparkConfigParameters `json:"sparkConfig,omitempty" tf:"spark_config,omitempty"`

	// The Spark events folder. Defaults to /events.
	// +kubebuilder:validation:Optional
	SparkEventsFolder *string `json:"sparkEventsFolder,omitempty" tf:"spark_events_folder,omitempty"`

	// The default folder where Spark logs will be written. Defaults to /logs.
	// +kubebuilder:validation:Optional
	SparkLogFolder *string `json:"sparkLogFolder,omitempty" tf:"spark_log_folder,omitempty"`

	// The Apache Spark version. Possible values are 2.4 , 3.1 , 3.2 and 3.3. Defaults to 2.4.
	// +kubebuilder:validation:Optional
	SparkVersion *string `json:"sparkVersion,omitempty" tf:"spark_version,omitempty"`

	// The ID of the Synapse Workspace where the Synapse Spark Pool should exist. Changing this forces a new Synapse Spark Pool to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.Workspace
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SynapseWorkspaceID *string `json:"synapseWorkspaceId,omitempty" tf:"synapse_workspace_id,omitempty"`

	// Reference to a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDRef *v1.Reference `json:"synapseWorkspaceIdRef,omitempty" tf:"-"`

	// Selector for a Workspace in synapse to populate synapseWorkspaceId.
	// +kubebuilder:validation:Optional
	SynapseWorkspaceIDSelector *v1.Selector `json:"synapseWorkspaceIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Synapse Spark Pool.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// SparkPoolSpec defines the desired state of SparkPool
type SparkPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SparkPoolParameters `json:"forProvider"`
}

// SparkPoolStatus defines the observed state of SparkPool.
type SparkPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SparkPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SparkPool is the Schema for the SparkPools API. Manages a Synapse Spark Pool.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SparkPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SparkPoolSpec   `json:"spec"`
	Status            SparkPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SparkPoolList contains a list of SparkPools
type SparkPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SparkPool `json:"items"`
}

// Repository type metadata.
var (
	SparkPool_Kind             = "SparkPool"
	SparkPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SparkPool_Kind}.String()
	SparkPool_KindAPIVersion   = SparkPool_Kind + "." + CRDGroupVersion.String()
	SparkPool_GroupVersionKind = CRDGroupVersion.WithKind(SparkPool_Kind)
)

func init() {
	SchemeBuilder.Register(&SparkPool{}, &SparkPoolList{})
}
