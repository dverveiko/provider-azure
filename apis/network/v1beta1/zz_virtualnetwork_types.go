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

type DDOSProtectionPlanObservation struct {
}

type DDOSProtectionPlanParameters struct {

	// Enable/disable DDoS Protection Plan on Virtual Network.
	// +kubebuilder:validation:Required
	Enable *bool `json:"enable" tf:"enable,omitempty"`

	// The ID of DDoS Protection Plan.
	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`
}

type VirtualNetworkObservation struct {

	// The GUID of the virtual network.
	GUID *string `json:"guid,omitempty" tf:"guid,omitempty"`

	// The virtual NetworkConfiguration ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Can be specified multiple times to define multiple subnets. Each subnet block supports fields documented below.
	Subnet []VirtualNetworkSubnetObservation `json:"subnet,omitempty" tf:"subnet,omitempty"`
}

type VirtualNetworkParameters struct {

	// The address space that is used the virtual network. You can supply more than one address space.
	// +kubebuilder:validation:Required
	AddressSpace []*string `json:"addressSpace" tf:"address_space,omitempty"`

	// The BGP community attribute in format <as-number>:<community-value>.
	// +kubebuilder:validation:Optional
	BGPCommunity *string `json:"bgpCommunity,omitempty" tf:"bgp_community,omitempty"`

	// A ddos_protection_plan block as documented below.
	// +kubebuilder:validation:Optional
	DDOSProtectionPlan []DDOSProtectionPlanParameters `json:"ddosProtectionPlan,omitempty" tf:"ddos_protection_plan,omitempty"`

	// List of IP addresses of DNS servers
	// +kubebuilder:validation:Optional
	DNSServers []*string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`

	// Specifies the Edge Zone within the Azure Region where this Virtual Network should exist. Changing this forces a new Virtual Network to be created.
	// +kubebuilder:validation:Optional
	EdgeZone *string `json:"edgeZone,omitempty" tf:"edge_zone,omitempty"`

	// The flow timeout in minutes for the Virtual Network, which is used to enable connection tracking for intra-VM flows. Possible values are between 4 and 30 minutes.
	// +kubebuilder:validation:Optional
	FlowTimeoutInMinutes *float64 `json:"flowTimeoutInMinutes,omitempty" tf:"flow_timeout_in_minutes,omitempty"`

	// The location/region where the virtual network is created. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// The name of the resource group in which to create the virtual network. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VirtualNetworkSubnetObservation struct {

	// The address prefix to use for the subnet.
	AddressPrefix *string `json:"addressPrefix,omitempty" tf:"address_prefix,omitempty"`

	// The ID of this subnet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the subnet. Changing this forces a new resource to be created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The Network Security Group to associate with the subnet. (Referenced by id, ie. azurerm_network_security_group.example.id)
	SecurityGroup *string `json:"securityGroup,omitempty" tf:"security_group,omitempty"`
}

type VirtualNetworkSubnetParameters struct {
}

// VirtualNetworkSpec defines the desired state of VirtualNetwork
type VirtualNetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualNetworkParameters `json:"forProvider"`
}

// VirtualNetworkStatus defines the observed state of VirtualNetwork.
type VirtualNetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualNetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetwork is the Schema for the VirtualNetworks API. Manages a virtual network including any configured subnets. Each subnet can optionally be configured with a security group to be associated with the subnet.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualNetworkSpec   `json:"spec"`
	Status            VirtualNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualNetworkList contains a list of VirtualNetworks
type VirtualNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualNetwork `json:"items"`
}

// Repository type metadata.
var (
	VirtualNetwork_Kind             = "VirtualNetwork"
	VirtualNetwork_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualNetwork_Kind}.String()
	VirtualNetwork_KindAPIVersion   = VirtualNetwork_Kind + "." + CRDGroupVersion.String()
	VirtualNetwork_GroupVersionKind = CRDGroupVersion.WithKind(VirtualNetwork_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualNetwork{}, &VirtualNetworkList{})
}
