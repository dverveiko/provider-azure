/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1beta1 "github.com/upbound/official-providers/provider-azure/apis/apimanagement/v1beta1"
	v1beta1authorization "github.com/upbound/official-providers/provider-azure/apis/authorization/v1beta1"
	v1beta1azure "github.com/upbound/official-providers/provider-azure/apis/azure/v1beta1"
	v1beta1cache "github.com/upbound/official-providers/provider-azure/apis/cache/v1beta1"
	v1beta1compute "github.com/upbound/official-providers/provider-azure/apis/compute/v1beta1"
	v1beta1containerservice "github.com/upbound/official-providers/provider-azure/apis/containerservice/v1beta1"
	v1beta1cosmosdb "github.com/upbound/official-providers/provider-azure/apis/cosmosdb/v1beta1"
	v1beta1datashare "github.com/upbound/official-providers/provider-azure/apis/datashare/v1beta1"
	v1beta1dbformariadb "github.com/upbound/official-providers/provider-azure/apis/dbformariadb/v1beta1"
	v1beta1dbforpostgresql "github.com/upbound/official-providers/provider-azure/apis/dbforpostgresql/v1beta1"
	v1beta1devices "github.com/upbound/official-providers/provider-azure/apis/devices/v1beta1"
	v1beta1eventhub "github.com/upbound/official-providers/provider-azure/apis/eventhub/v1beta1"
	v1beta1insights "github.com/upbound/official-providers/provider-azure/apis/insights/v1beta1"
	v1beta1keyvault "github.com/upbound/official-providers/provider-azure/apis/keyvault/v1beta1"
	v1beta1logic "github.com/upbound/official-providers/provider-azure/apis/logic/v1beta1"
	v1beta1management "github.com/upbound/official-providers/provider-azure/apis/management/v1beta1"
	v1beta1marketplaceordering "github.com/upbound/official-providers/provider-azure/apis/marketplaceordering/v1beta1"
	v1beta1network "github.com/upbound/official-providers/provider-azure/apis/network/v1beta1"
	v1beta1notificationhubs "github.com/upbound/official-providers/provider-azure/apis/notificationhubs/v1beta1"
	v1beta1operationalinsights "github.com/upbound/official-providers/provider-azure/apis/operationalinsights/v1beta1"
	v1beta1resources "github.com/upbound/official-providers/provider-azure/apis/resources/v1beta1"
	v1beta1security "github.com/upbound/official-providers/provider-azure/apis/security/v1beta1"
	v1beta1sql "github.com/upbound/official-providers/provider-azure/apis/sql/v1beta1"
	v1beta1storage "github.com/upbound/official-providers/provider-azure/apis/storage/v1beta1"
	v1beta1storagecache "github.com/upbound/official-providers/provider-azure/apis/storagecache/v1beta1"
	v1beta1storagesync "github.com/upbound/official-providers/provider-azure/apis/storagesync/v1beta1"
	v1alpha1 "github.com/upbound/official-providers/provider-azure/apis/v1alpha1"
	v1beta1apis "github.com/upbound/official-providers/provider-azure/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1beta1.SchemeBuilder.AddToScheme,
		v1beta1authorization.SchemeBuilder.AddToScheme,
		v1beta1azure.SchemeBuilder.AddToScheme,
		v1beta1cache.SchemeBuilder.AddToScheme,
		v1beta1compute.SchemeBuilder.AddToScheme,
		v1beta1containerservice.SchemeBuilder.AddToScheme,
		v1beta1cosmosdb.SchemeBuilder.AddToScheme,
		v1beta1datashare.SchemeBuilder.AddToScheme,
		v1beta1dbformariadb.SchemeBuilder.AddToScheme,
		v1beta1dbforpostgresql.SchemeBuilder.AddToScheme,
		v1beta1devices.SchemeBuilder.AddToScheme,
		v1beta1eventhub.SchemeBuilder.AddToScheme,
		v1beta1insights.SchemeBuilder.AddToScheme,
		v1beta1keyvault.SchemeBuilder.AddToScheme,
		v1beta1logic.SchemeBuilder.AddToScheme,
		v1beta1management.SchemeBuilder.AddToScheme,
		v1beta1marketplaceordering.SchemeBuilder.AddToScheme,
		v1beta1network.SchemeBuilder.AddToScheme,
		v1beta1notificationhubs.SchemeBuilder.AddToScheme,
		v1beta1operationalinsights.SchemeBuilder.AddToScheme,
		v1beta1resources.SchemeBuilder.AddToScheme,
		v1beta1security.SchemeBuilder.AddToScheme,
		v1beta1sql.SchemeBuilder.AddToScheme,
		v1beta1storage.SchemeBuilder.AddToScheme,
		v1beta1storagecache.SchemeBuilder.AddToScheme,
		v1beta1storagesync.SchemeBuilder.AddToScheme,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1beta1apis.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
