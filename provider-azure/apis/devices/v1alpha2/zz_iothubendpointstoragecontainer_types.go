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

type IOTHubEndpointStorageContainerObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type IOTHubEndpointStorageContainerParameters struct {

	// +kubebuilder:validation:Optional
	AuthenticationType *string `json:"authenticationType,omitempty" tf:"authentication_type,omitempty"`

	// +kubebuilder:validation:Optional
	BatchFrequencyInSeconds *float64 `json:"batchFrequencyInSeconds,omitempty" tf:"batch_frequency_in_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	ConnectionStringSecretRef *v1.SecretKeySelector `json:"connectionStringSecretRef,omitempty" tf:"-"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/storage/v1alpha2.Container
	// +kubebuilder:validation:Optional
	ContainerName *string `json:"containerName,omitempty" tf:"container_name,omitempty"`

	// +kubebuilder:validation:Optional
	ContainerNameRef *v1.Reference `json:"containerNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ContainerNameSelector *v1.Selector `json:"containerNameSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Encoding *string `json:"encoding,omitempty" tf:"encoding,omitempty"`

	// +kubebuilder:validation:Optional
	EndpointURI *string `json:"endpointUri,omitempty" tf:"endpoint_uri,omitempty"`

	// +kubebuilder:validation:Optional
	FileNameFormat *string `json:"fileNameFormat,omitempty" tf:"file_name_format,omitempty"`

	// +kubebuilder:validation:Required
	IOTHubID *string `json:"iothubId" tf:"iothub_id,omitempty"`

	// +kubebuilder:validation:Optional
	IdentityID *string `json:"identityId,omitempty" tf:"identity_id,omitempty"`

	// +kubebuilder:validation:Optional
	MaxChunkSizeInBytes *float64 `json:"maxChunkSizeInBytes,omitempty" tf:"max_chunk_size_in_bytes,omitempty"`

	// +crossplane:generate:reference:type=github.com/upbound/official-providers/provider-azure/apis/azure/v1alpha2.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`
}

// IOTHubEndpointStorageContainerSpec defines the desired state of IOTHubEndpointStorageContainer
type IOTHubEndpointStorageContainerSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IOTHubEndpointStorageContainerParameters `json:"forProvider"`
}

// IOTHubEndpointStorageContainerStatus defines the observed state of IOTHubEndpointStorageContainer.
type IOTHubEndpointStorageContainerStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IOTHubEndpointStorageContainerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubEndpointStorageContainer is the Schema for the IOTHubEndpointStorageContainers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type IOTHubEndpointStorageContainer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IOTHubEndpointStorageContainerSpec   `json:"spec"`
	Status            IOTHubEndpointStorageContainerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IOTHubEndpointStorageContainerList contains a list of IOTHubEndpointStorageContainers
type IOTHubEndpointStorageContainerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IOTHubEndpointStorageContainer `json:"items"`
}

// Repository type metadata.
var (
	IOTHubEndpointStorageContainer_Kind             = "IOTHubEndpointStorageContainer"
	IOTHubEndpointStorageContainer_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IOTHubEndpointStorageContainer_Kind}.String()
	IOTHubEndpointStorageContainer_KindAPIVersion   = IOTHubEndpointStorageContainer_Kind + "." + CRDGroupVersion.String()
	IOTHubEndpointStorageContainer_GroupVersionKind = CRDGroupVersion.WithKind(IOTHubEndpointStorageContainer_Kind)
)

func init() {
	SchemeBuilder.Register(&IOTHubEndpointStorageContainer{}, &IOTHubEndpointStorageContainerList{})
}