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

type RouteObservation_2 struct {

	// The Route ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type RouteParameters_2 struct {

	// The destination to which the route applies. Can be CIDR (such as 10.1.0.0/16) or Azure Service Tag (such as ApiManagement, AzureBackup or AzureMonitor) format.
	// +kubebuilder:validation:Required
	AddressPrefix *string `json:"addressPrefix" tf:"address_prefix,omitempty"`

	// Contains the IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
	// +kubebuilder:validation:Optional
	NextHopInIPAddress *string `json:"nextHopInIpAddress,omitempty" tf:"next_hop_in_ip_address,omitempty"`

	// The type of Azure hop the packet should be sent to. Possible values are VirtualNetworkGateway, VnetLocal, Internet, VirtualAppliance and None.
	// +kubebuilder:validation:Required
	NextHopType *string `json:"nextHopType" tf:"next_hop_type,omitempty"`

	// The name of the resource group in which to create the route. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the route table within which create the route. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.RouteTable
	// +kubebuilder:validation:Optional
	RouteTableName *string `json:"routeTableName,omitempty" tf:"route_table_name,omitempty"`

	// Reference to a RouteTable in network to populate routeTableName.
	// +kubebuilder:validation:Optional
	RouteTableNameRef *v1.Reference `json:"routeTableNameRef,omitempty" tf:"-"`

	// Selector for a RouteTable in network to populate routeTableName.
	// +kubebuilder:validation:Optional
	RouteTableNameSelector *v1.Selector `json:"routeTableNameSelector,omitempty" tf:"-"`
}

// RouteSpec defines the desired state of Route
type RouteSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteParameters_2 `json:"forProvider"`
}

// RouteStatus defines the observed state of Route.
type RouteStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Route is the Schema for the Routes API. Manages a Route within a Route Table.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Route struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteSpec   `json:"spec"`
	Status            RouteStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteList contains a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Route `json:"items"`
}

// Repository type metadata.
var (
	Route_Kind             = "Route"
	Route_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Route_Kind}.String()
	Route_KindAPIVersion   = Route_Kind + "." + CRDGroupVersion.String()
	Route_GroupVersionKind = CRDGroupVersion.WithKind(Route_Kind)
)

func init() {
	SchemeBuilder.Register(&Route{}, &RouteList{})
}
