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

type DataSetParquetAzureBlobStorageLocationObservation struct {
}

type DataSetParquetAzureBlobStorageLocationParameters struct {

	// The container on the Azure Blob Storage Account hosting the file.
	// +kubebuilder:validation:Required
	Container *string `json:"container" tf:"container,omitempty"`

	// Is the container using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicContainerEnabled *bool `json:"dynamicContainerEnabled,omitempty" tf:"dynamic_container_enabled,omitempty"`

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	// +kubebuilder:validation:Optional
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`
}

type DataSetParquetHTTPServerLocationObservation struct {
}

type DataSetParquetHTTPServerLocationParameters struct {

	// Is the filename using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicFilenameEnabled *bool `json:"dynamicFilenameEnabled,omitempty" tf:"dynamic_filename_enabled,omitempty"`

	// Is the path using dynamic expression, function or system variables? Defaults to false.
	// +kubebuilder:validation:Optional
	DynamicPathEnabled *bool `json:"dynamicPathEnabled,omitempty" tf:"dynamic_path_enabled,omitempty"`

	// The filename of the file on the web server.
	// +kubebuilder:validation:Required
	Filename *string `json:"filename" tf:"filename,omitempty"`

	// The folder path to the file on the web server.
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// The base URL to the web server hosting the file.
	// +kubebuilder:validation:Required
	RelativeURL *string `json:"relativeUrl" tf:"relative_url,omitempty"`
}

type DataSetParquetObservation struct {

	// The ID of the Data Factory Dataset.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type DataSetParquetParameters struct {

	// A map of additional properties to associate with the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// A azure_blob_storage_location block as defined below.
	// +kubebuilder:validation:Optional
	AzureBlobStorageLocation []DataSetParquetAzureBlobStorageLocationParameters `json:"azureBlobStorageLocation,omitempty" tf:"azure_blob_storage_location,omitempty"`

	// The compression codec used to read/write text files. Valid values are bzip2, gzip, deflate, ZipDeflate, TarGzip, Tar, snappy, or lz4. Please note these values are case-sensitive.
	// +kubebuilder:validation:Optional
	CompressionCodec *string `json:"compressionCodec,omitempty" tf:"compression_codec,omitempty"`

	// Specifies the compression level. Possible values are Optimal and Fastest,
	// +kubebuilder:validation:Optional
	CompressionLevel *string `json:"compressionLevel,omitempty" tf:"compression_level,omitempty"`

	// The Data Factory ID in which to associate the Dataset with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	// +kubebuilder:validation:Optional
	Folder *string `json:"folder,omitempty" tf:"folder,omitempty"`

	// A http_server_location block as defined below.
	// +kubebuilder:validation:Optional
	HTTPServerLocation []DataSetParquetHTTPServerLocationParameters `json:"httpServerLocation,omitempty" tf:"http_server_location,omitempty"`

	// The Data Factory Linked Service name in which to associate the Dataset with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.LinkedServiceWeb
	// +kubebuilder:validation:Optional
	LinkedServiceName *string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`

	// Reference to a LinkedServiceWeb in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameRef *v1.Reference `json:"linkedServiceNameRef,omitempty" tf:"-"`

	// Selector for a LinkedServiceWeb in datafactory to populate linkedServiceName.
	// +kubebuilder:validation:Optional
	LinkedServiceNameSelector *v1.Selector `json:"linkedServiceNameSelector,omitempty" tf:"-"`

	// A map of parameters to associate with the Data Factory Dataset.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// A schema_column block as defined below.
	// +kubebuilder:validation:Optional
	SchemaColumn []DataSetParquetSchemaColumnParameters `json:"schemaColumn,omitempty" tf:"schema_column,omitempty"`
}

type DataSetParquetSchemaColumnObservation struct {
}

type DataSetParquetSchemaColumnParameters struct {

	// The description of the column.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the column.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Type of the column. Valid values are Byte, Byte[], Boolean, Date, DateTime,DateTimeOffset, Decimal, Double, Guid, Int16, Int32, Int64, Single, String, TimeSpan. Please note these values are case sensitive.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// DataSetParquetSpec defines the desired state of DataSetParquet
type DataSetParquetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataSetParquetParameters `json:"forProvider"`
}

// DataSetParquetStatus defines the observed state of DataSetParquet.
type DataSetParquetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataSetParquetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetParquet is the Schema for the DataSetParquets API. Manages an Azure Parquet Dataset inside an Azure Data Factory.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type DataSetParquet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataSetParquetSpec   `json:"spec"`
	Status            DataSetParquetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataSetParquetList contains a list of DataSetParquets
type DataSetParquetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataSetParquet `json:"items"`
}

// Repository type metadata.
var (
	DataSetParquet_Kind             = "DataSetParquet"
	DataSetParquet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataSetParquet_Kind}.String()
	DataSetParquet_KindAPIVersion   = DataSetParquet_Kind + "." + CRDGroupVersion.String()
	DataSetParquet_GroupVersionKind = CRDGroupVersion.WithKind(DataSetParquet_Kind)
)

func init() {
	SchemeBuilder.Register(&DataSetParquet{}, &DataSetParquetList{})
}
