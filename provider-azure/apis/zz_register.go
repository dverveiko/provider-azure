/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha2 "github.com/upbound/official-providers/provider-azure/apis/authorization/v1alpha2"
	v1alpha2azure "github.com/upbound/official-providers/provider-azure/apis/azure/v1alpha2"
	v1alpha2cache "github.com/upbound/official-providers/provider-azure/apis/cache/v1alpha2"
	v1beta1 "github.com/upbound/official-providers/provider-azure/apis/compute/v1beta1"
	v1alpha2containerservice "github.com/upbound/official-providers/provider-azure/apis/containerservice/v1alpha2"
	v1alpha1 "github.com/upbound/official-providers/provider-azure/apis/cosmosdb/v1alpha1"
	v1alpha2cosmosdb "github.com/upbound/official-providers/provider-azure/apis/cosmosdb/v1alpha2"
	v1alpha2dbformariadb "github.com/upbound/official-providers/provider-azure/apis/dbformariadb/v1alpha2"
	v1alpha2dbforpostgresql "github.com/upbound/official-providers/provider-azure/apis/dbforpostgresql/v1alpha2"
	v1alpha1devices "github.com/upbound/official-providers/provider-azure/apis/devices/v1alpha1"
	v1alpha2devices "github.com/upbound/official-providers/provider-azure/apis/devices/v1alpha2"
	v1alpha2eventhub "github.com/upbound/official-providers/provider-azure/apis/eventhub/v1alpha2"
	v1alpha2insights "github.com/upbound/official-providers/provider-azure/apis/insights/v1alpha2"
	v1alpha1iothub "github.com/upbound/official-providers/provider-azure/apis/iothub/v1alpha1"
	v1alpha2keyvault "github.com/upbound/official-providers/provider-azure/apis/keyvault/v1alpha2"
	v1alpha2loganalytics "github.com/upbound/official-providers/provider-azure/apis/loganalytics/v1alpha2"
	v1alpha2network "github.com/upbound/official-providers/provider-azure/apis/network/v1alpha2"
	v1alpha1resource "github.com/upbound/official-providers/provider-azure/apis/resource/v1alpha1"
	v1alpha2resources "github.com/upbound/official-providers/provider-azure/apis/resources/v1alpha2"
	v1alpha2sql "github.com/upbound/official-providers/provider-azure/apis/sql/v1alpha2"
	v1alpha2storage "github.com/upbound/official-providers/provider-azure/apis/storage/v1alpha2"
	v1alpha1apis "github.com/upbound/official-providers/provider-azure/apis/v1alpha1"
	v1alpha1virtual "github.com/upbound/official-providers/provider-azure/apis/virtual/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha2.SchemeBuilder.AddToScheme,
		v1alpha2azure.SchemeBuilder.AddToScheme,
		v1alpha2cache.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha2containerservice.SchemeBuilder.AddToScheme,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha2cosmosdb.SchemeBuilder.AddToScheme,
		v1alpha2dbformariadb.SchemeBuilder.AddToScheme,
		v1alpha2dbforpostgresql.SchemeBuilder.AddToScheme,
		v1alpha1devices.SchemeBuilder.AddToScheme,
		v1alpha2devices.SchemeBuilder.AddToScheme,
		v1alpha2eventhub.SchemeBuilder.AddToScheme,
		v1alpha2insights.SchemeBuilder.AddToScheme,
		v1alpha1iothub.SchemeBuilder.AddToScheme,
		v1alpha2keyvault.SchemeBuilder.AddToScheme,
		v1alpha2loganalytics.SchemeBuilder.AddToScheme,
		v1alpha2network.SchemeBuilder.AddToScheme,
		v1alpha1resource.SchemeBuilder.AddToScheme,
		v1alpha2resources.SchemeBuilder.AddToScheme,
		v1alpha2sql.SchemeBuilder.AddToScheme,
		v1alpha2storage.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1alpha1virtual.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}