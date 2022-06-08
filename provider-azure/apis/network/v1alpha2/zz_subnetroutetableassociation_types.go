/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SubnetRouteTableAssociationObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SubnetRouteTableAssociationParameters struct {

	// +kubebuilder:validation:Required
	RouteTableID *string `json:"routeTableId" tf:"route_table_id,omitempty"`

	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/official-providers/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// SubnetRouteTableAssociationSpec defines the desired state of SubnetRouteTableAssociation
type SubnetRouteTableAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetRouteTableAssociationParameters `json:"forProvider"`
}

// SubnetRouteTableAssociationStatus defines the observed state of SubnetRouteTableAssociation.
type SubnetRouteTableAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetRouteTableAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetRouteTableAssociation is the Schema for the SubnetRouteTableAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SubnetRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetRouteTableAssociationSpec   `json:"spec"`
	Status            SubnetRouteTableAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetRouteTableAssociationList contains a list of SubnetRouteTableAssociations
type SubnetRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetRouteTableAssociation `json:"items"`
}

// Repository type metadata.
var (
	SubnetRouteTableAssociation_Kind             = "SubnetRouteTableAssociation"
	SubnetRouteTableAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetRouteTableAssociation_Kind}.String()
	SubnetRouteTableAssociation_KindAPIVersion   = SubnetRouteTableAssociation_Kind + "." + CRDGroupVersion.String()
	SubnetRouteTableAssociation_GroupVersionKind = CRDGroupVersion.WithKind(SubnetRouteTableAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetRouteTableAssociation{}, &SubnetRouteTableAssociationList{})
}