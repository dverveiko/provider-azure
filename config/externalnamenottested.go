package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameNotTestedConfigs contains no-tested configurations for this
// provider.
var ExternalNameNotTestedConfigs = map[string]config.ExternalName{
	// aad
	//
	// Domain Services can be imported using the resource ID, together with the
	// Replica Set ID that you wish to designate as the initial replica set
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AAD/domainServices/instance1/initialReplicaSetId/00000000-0000-0000-0000-000000000000
	"azurerm_active_directory_domain_service": config.IdentifierFromProvider,
	// Domain Service Replica Sets can be imported using the resource ID of the
	// parent Domain Service and the Replica Set ID
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AAD/domainServices/instance1/replicaSets/00000000-0000-0000-0000-000000000000
	"azurerm_active_directory_domain_service_replica_set": config.IdentifierFromProvider,

	// aadb2c
	//
	// AAD B2C Directories can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.AzureActiveDirectory/b2cDirectories/directory-name
	"azurerm_aadb2c_directory": config.IdentifierFromProvider,

	// aadiam
	//
	// Monitor Azure Active Directory Diagnostic Settings can be imported using the resource id
	// /providers/Microsoft.AADIAM/diagnosticSettings/setting1
	"azurerm_monitor_aad_diagnostic_setting": config.TemplatedStringAsIdentifier("name", "/providers/Microsoft.AADIAM/diagnosticSettings/{{ .external_name }}"),

	// api
	//
	// API Connections can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.Web/connections/example-connection
	"azurerm_api_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/connections/{{ .external_name }}"),

	// apimanagement
	//
	// API Management Custom Domains can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/instance1/customDomains/default
	"azurerm_api_management_custom_domain": config.TemplatedStringAsIdentifier("", "{{ .parameters.api_management_id }}/customDomains/default"),
	// API Management Azure AD B2C Identity Providers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ApiManagement/service/service1/identityProviders/AadB2C
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_api_management_identity_provider_aadb2c": config.IdentifierFromProvider,

	// appconfiguration
	//
	// App Configurations can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1
	"azurerm_app_configuration": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AppConfiguration/configurationStores/{{ .external_name }}"),
	// There are two different syntaxes:
	// App Configuration Features can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationFeature/appConfFeature1/Label/label1
	// If you wish to import a key with an empty label then sustitute the label's name with %00, like this
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationFeature/appConfFeature1/Label/%00
	"azurerm_app_configuration_feature": config.IdentifierFromProvider,
	// There are two different syntaxes:
	// App Configuration Keys can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationKey/appConfKey1/Label/label1
	// If you wish to import a key with an empty label then sustitute the label's name with %00, like this
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resourceGroup1/providers/Microsoft.AppConfiguration/configurationStores/appConf1/AppConfigurationKey/appConfKey1/Label/%00
	"azurerm_app_configuration_key": config.IdentifierFromProvider,

	// attestation
	//
	// Attestation Providers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Attestation/attestationProviders/provider1
	"azurerm_attestation_provider": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Attestation/attestationProviders/{{ .external_name }}"),

	// kusto
	//
	// Customer Managed Keys for a Kusto Cluster can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1
	"azurerm_kusto_cluster_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.cluster_id }}"),
	// Kusto Scripts can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Kusto/Clusters/cluster1/Databases/database1/Scripts/script1
	"azurerm_kusto_script": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Kusto/Clusters/{{ .parameters.cluster_name }}/Databases/{{ .parameters.database_id }}/Scripts/{{ .external_name }}"),

	// devtest
	//
	// An existing Dev Test Global Shutdown Schedule can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-rg/providers/Microsoft.DevTestLab/schedules/shutdown-computevm-SampleVM
	// TODO: For now API is not normalized. While testing resource we can check and normalize the API.
	"azurerm_dev_test_global_vm_shutdown_schedule": config.IdentifierFromProvider,
	// Dev Test Labs can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1
	"azurerm_dev_test_lab": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DevTestLab/labs/{{ .external_name }}"),
	// Dev Test Linux Virtual Machines can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/virtualmachines/machine1
	"azurerm_dev_test_linux_virtual_machine": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DevTestLab/labs/{{ .parameters.lab_name }}/virtualmachines/{{ .external_name }}"),
	// Dev Test Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/policysets/default/policies/policy1
	"azurerm_dev_test_policy": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DevTestLab/labs/{{ .parameters.lab_name }}/policysets/{{ .parameters.policy_set_name }}/policies/{{ .external_name }}"),
	// DevTest Schedule's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DevTestLab/labs/myDevTestLab/schedules/labvmautostart
	"azurerm_dev_test_schedule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DevTestLab/labs/{{ .parameters.lab_name }}/schedules/{{ .external_name }}"),
	// DevTest Virtual Networks can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/virtualnetworks/network1
	"azurerm_dev_test_virtual_network": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DevTestLab/labs/{{ .parameters.lab_name }}/virtualnetworks/{{ .external_name }}"),
	// DevTest Windows Virtual Machines can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DevTestLab/labs/lab1/virtualmachines/machine1
	"azurerm_dev_test_windows_virtual_machine": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DevTestLab/labs/{{ .parameters.lab_name }}/virtualmachines/{{ .external_name }}"),

	// appplatform
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/apiPortals/apiPortal1
	"azurerm_spring_cloud_api_portal": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/apiPortals/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/deployments/deploy1
	"azurerm_spring_cloud_build_deployment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/deployments/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/buildServices/buildService1/builders/builder1/buildPackBindings/binding1
	"azurerm_spring_cloud_build_pack_binding": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_builder_id }}/buildPackBindings/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/buildServices/buildService1/builders/builder1
	"azurerm_spring_cloud_builder": config.IdentifierFromProvider,
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/configurationServices/configurationService1
	"azurerm_spring_cloud_configuration_service": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/configurationServices/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.AppPlatform/spring/spring1/apps/app1/deployments/deploy1
	"azurerm_spring_cloud_container_deployment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_app_id }}/deployments/{{ .external_name }}"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/gateways/gateway1
	"azurerm_spring_cloud_gateway": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_service_id }}/gateways/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/gateways/gateway1/domains/domain1
	"azurerm_spring_cloud_gateway_custom_domain": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_gateway_id }}/domains/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.AppPlatform/spring/service1/gateways/gateway1/routeConfigs/routeConfig1
	"azurerm_spring_cloud_gateway_route_config": config.TemplatedStringAsIdentifier("name", "{{ .parameters.spring_cloud_gateway_id }}/routeConfigs/{{ .external_name }}"),

	// avs
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AVS/privateClouds/privateCloud1/clusters/cluster1
	"azurerm_vmware_cluster": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vmware_cloud_id }}/clusters/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.AVS/privateClouds/privateCloud1/authorizations/authorization1
	"azurerm_vmware_express_route_authorization": config.TemplatedStringAsIdentifier("name", "{{ .parameters.vmware_cloud_id }}/clusters/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/PrivateClouds/privateCloud1
	"azurerm_vmware_private_cloud": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.AVS/PrivateClouds/{{ .external_name }}"),

	// blueprint
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Blueprint/blueprintAssignments/assignSimpleBlueprint
	"azurerm_blueprint_assignment": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Blueprint/blueprintAssignments/{{ .external_name }}"),

	// cdn
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Cdn/profiles/profile1/endpoints/endpoint1/customDomains/domain1
	"azurerm_cdn_endpoint_custom_domain": config.TemplatedStringAsIdentifier("name", "{{ .parameters.cdn_endpoint_id }}/customDomains/{{ .external_name }}"),

	// certificateregistration
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.CertificateRegistration/certificateOrders/certificateorder1
	"azurerm_app_service_certificate_order": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.CertificateRegistration/certificateOrders/{{ .external_name }}"),

	// cognitiveservices
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.CognitiveServices/accounts/account1
	"azurerm_cognitive_account_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .paramteres.cognitive_account_id }}"),

	// authorization
	//
	// /providers/Microsoft.Management/managementGroups/group1/providers/Microsoft.Authorization/policyAssignments/assignment1
	"azurerm_management_group_policy_assignment": config.TemplatedStringAsIdentifier("name", "{{ .parameters.management_group_id }}/providers/Microsoft.Authorization/policyAssignments/{{ .external_name }}"),
	// /providers/Microsoft.Management/managementGroups/group1/providers/Microsoft.Authorization/policyExemptions/exemption1
	"azurerm_management_group_policy_exemption": config.TemplatedStringAsIdentifier("name", "{{ .parameters.management_group_id }}/providers/Microsoft.Authorization/policyExemptions/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Authorization/policySetDefinitions/testPolicySet
	"azurerm_policy_set_definition": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.Authorization/policySetDefinitions/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Authorization/policyExemptions/exemption1
	"azurerm_resource_group_policy_exemption": config.TemplatedStringAsIdentifier("name", "{{ .parameters.resource_group_id }}/providers/Microsoft.Authorization/policyExemptions/{{ .external_name }}"),

	// automation
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/certificates/certificate1
	"azurerm_automation_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/certificates/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
	"azurerm_automation_connection": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/connections/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
	"azurerm_automation_connection_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/connections/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
	"azurerm_automation_connection_classic_certificate": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/connections/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/connections/conn1
	"azurerm_automation_connection_service_principal": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/connections/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/nodeConfigurations/configuration1
	"azurerm_automation_dsc_nodeconfiguration": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/nodeConfigurations/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/jobSchedules/10000000-1001-1001-1001-000000000001
	"azurerm_automation_job_schedule": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/runbooks/Get-AzureVMTutorial
	"azurerm_automation_runbook": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/runbooks/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/webHooks/TestRunbook_webhook
	"azurerm_automation_webhook": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/schedules/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/schedules/schedule1
	"azurerm_automation_schedule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Automation/automationAccounts/{{ .parameters.automation_account_name }}/schedules/{{ .external_name }}"),

	// botservice
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/DirectLineSpeechChannel
	"azurerm_bot_channel_direct_line_speech": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/DirectLineSpeechChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.BotService/botServices/example/channels/EmailChannel
	"azurerm_bot_channel_email": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/EmailChannel"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.BotService/botServices/botService1/channels/FacebookChannel
	"azurerm_bot_channel_facebook": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .parameters.bot_name }}/channels/FacebookChannel"),
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.BotService/botServices/botService1
	"azurerm_bot_service_azure_bot": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.BotService/botServices/{{ .external_name }}"),

	// compute
	//
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/example-resource-group/providers/Microsoft.StoragePool/diskPools/diskPoolValue/iscsiTargets/iscsiTargetValue/lun|/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Compute/disks/disk1
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_disk_pool_iscsi_target_lun": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StoragePool/diskPools/storagePool1/managedDisks|/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Compute/disks/disk1
	// TODO: For now API is not normalized. While testing resource we can check the actual ID and normalize the API.
	"azurerm_disk_pool_managed_disk_attachment": config.IdentifierFromProvider,
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1
	"azurerm_shared_image": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/galleries/{{ .parameters.gallery_name }}/images/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Compute/galleries/gallery1/images/image1/versions/1.2.3
	"azurerm_shared_image_version": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/galleries/{{ .parameters.gallery_name }}/images/{{ .parameters.image_name }}/versions/{{ .external_name }}"),
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Compute/SshPublicKeys/mySshPublicKeyName1
	"azurerm_ssh_public_key": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Compute/SshPublicKeys/{{ .external_name }}"),

	// customproviders
	//
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.CustomProviders/resourceProviders/example
	"azurerm_custom_provider": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.CustomProviders/resourceProviders/{{ .external_name }}"),

	// batch
	//
	// Batch Account can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Batch/batchAccounts/account1
	"azurerm_batch_account": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Batch/batchAccounts/{{ .externalName }}"),
	// Batch Application can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Batch/batchAccounts/exampleba/applications/example-batch-application
	"azurerm_batch_application": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Batch/batchAccounts/{{ .parameters.batch_account_name }}/applications/{{ .externalName }}"),
	// Batch Certificate can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example-rg/providers/Microsoft.Batch/batchAccounts/batch1/certificates/certificate1
	"azurerm_batch_certificate": config.TemplatedStringAsIdentifier("name", "/subscBriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Batch/batchAccounts/{{ .parameters.batch_account_name }}/certificates/{{ .externalName }}"),
	// Batch Pool can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.Batch/batchAccounts/myBatchAccount1/pools/myBatchPool1
	"azurerm_batch_pool": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Batch/batchAccounts/{{ .parameters.batch_account_name }}/pools/{{ .externalName }}"),
	// Batch job can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Batch/batchAccounts/account1/pools/pool1/jobs/job1
	"azurerm_batch_job": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Batch/batchAccounts/{{ .parameters.batch_account_name }}/pools/{{ .parameters.batch_pool_name }}/jobs/{{ .externalName }}"),

	// insights
	//
	// Activity log alerts can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/activityLogAlerts/myalertname
	"azurerm_monitor_activity_log_alert": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/activityLogAlerts/{{ .externalName }}"),
	// AutoScale Setting can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/autoScaleSettings/setting1
	"azurerm_monitor_autoscale_setting": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/autoScaleSettings/{{ .externalName }}"),
	// Diagnostic Settings can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.KeyVault/vaults/vault1|logMonitoring1
	"azurerm_monitor_diagnostic_setting": config.IdentifierFromProvider,
	// Scheduled Query Rule Alerts can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/scheduledQueryRules/myrulename
	"azurerm_monitor_scheduled_query_rules_alert": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/scheduledQueryRules/{{ .externalName }}"),
	// Scheduled Query Rule Log can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Insights/scheduledQueryRules/myrulename
	"azurerm_monitor_scheduled_query_rules_log": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Insights/scheduledQueryRules/{{ .externalName }}"),

	// iot
	//
	// Azure IoT Time Series Insights EventHub Event Source can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.TimeSeriesInsights/environments/environment1/eventSources/example
	"azurerm_time_series_insights_event_source_eventhub": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.TimeSeriesInsights/environments/{{ .parameters.environment_name }}/eventSources/{{ .externalName }}"),

	// linux
	//
	// Linux Function Apps can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1
	"azurerm_linux_function_app": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Web/sites/{{ .externalName }}"),
	// A Linux Function App Slot can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1
	"azurerm_linux_function_app_slot": config.TemplatedStringAsIdentifier("name", "{{ .parameters.function_app_id }}/slots/{{ .externalName }}"),
	// Linux Web Apps can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Web/sites/site1/slots/slot1
	"azurerm_linux_web_app_slot": config.TemplatedStringAsIdentifier("name", "{{ .parameters.function_app_id }}/slots/{{ .externalName }}"),

	// load
	//
	// An existing Load Test can be imported into Terraform using the resource id
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.LoadTestService/loadTests/{loadTestName}
	"azurerm_load_test": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.LoadTestService/loadTests/{{ .externalName }}"),

	// logic
	//
	// Logic App Custom Actions can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/actions/custom1
	"azurerm_logic_app_action_custom": config.TemplatedStringAsIdentifier("name", "{{ .parameters.logic_app_id }}/actions/{{ .externalName }}"),
	// Logic App HTTP Actions can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Logic/workflows/workflow1/actions/webhook1
	"azurerm_logic_app_action_http": config.TemplatedStringAsIdentifier("name", "{{ .parameters.logic_app_id }}/actions/{{ .externalName }}"),
	// Logic App Integration Account Agreements can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/agreements/agreement1
	"azurerm_logic_app_integration_account_agreement": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logic/integrationAccounts/{{ .parameters.integration_account_name }}/agreements/{{ .externalName }}"),
	// Logic App Integration Account Assemblies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/assemblies/assembly1
	"azurerm_logic_app_integration_account_assembly": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logic/integrationAccounts/{{ .parameters.integration_account_name }}/assemblies/{{ .externalName }}"),
	// Logic App Integration Account Maps can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/maps/map1
	"azurerm_logic_app_integration_account_map": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Logic/integrationAccounts/{{ .parameters.integration_account_name }}/maps/{{ .externalName }}"),

	// operationalinsights
	//
	// Log Analytics Clusters can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/clusters/cluster1
	"azurerm_log_analytics_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.OperationalInsights/clusters/{{ .external_name }}"),
	// Log Analytics Cluster Customer Managed Keys can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/clusters/cluster1
	"azurerm_log_analytics_cluster_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.log_analytics_cluster_id }}"),
	// Log Analytics Data Export Rule can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/dataExports/dataExport1
	"azurerm_log_analytics_data_export_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.workspace_resource_id }}/dataExports/{{ .external_name }}"),
	// Log Analytics Windows Event DataSources can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/dataSources/datasource1
	"azurerm_log_analytics_datasource_windows_event": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.OperationalInsights/workspaces/{{ .parameters.workspace_name }}/dataSources/{{ .external_name }}"),
	// Log Analytics Windows Performance Counter DataSources can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/dataSources/datasource1
	"azurerm_log_analytics_datasource_windows_performance_counter": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.OperationalInsights/workspaces/{{ .parameters.workspace_name }}/dataSources/{{ .external_name }}"),
	// Log Analytics Workspaces can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/linkedServices/Automation
	"azurerm_log_analytics_linked_service": config.TemplatedStringAsIdentifier("", "{{ .parameters.workspace_id }}/linkedServices/{{ .external_name }}"),
	// Log Analytics Linked Storage Accounts can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/linkedStorageAccounts/{dataSourceType}
	"azurerm_log_analytics_linked_storage_account": config.TemplatedStringAsIdentifier("", "{{ .parameters.workspace_resource_id }}/linkedStorageAccounts/{{ .external_name }}"),
	// Log Analytics Saved Searches can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/savedSearches/search1
	"azurerm_log_analytics_saved_search": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/savedSearches/{{ .external_name }}"),
	// Log Analytics Solutions can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.OperationsManagement/solutions/solution1
	"azurerm_log_analytics_solution": config.TemplatedStringAsIdentifier("solution_name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.OperationsManagement/solutions/{{ .external_name }}"),

	// portal
	//
	// Dashboards can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.Portal/dashboards/00000000-0000-0000-0000-000000000000
	"azurerm_portal_dashboard": config.IdentifierFromProvider,
	// Portal Tenant Configurations can be imported using the resource id
	// /providers/Microsoft.Portal/tenantConfigurations/default
	"azurerm_portal_tenant_configuration": config.IdentifierFromProvider,

	// streamanalytics
	//
	// Stream Analytics Outputs to Microsoft SQL Server Database can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
	"azurerm_stream_analytics_output_mssql": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourcegroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingjobs/{{ .parameters.stream_analytics_job_name }}/outputs/{{ .external_name }}"),
	// Stream Analytics Output ServiceBus Queue's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
	"azurerm_stream_analytics_output_servicebus_queue": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourcegroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingjobs/{{ .parameters.stream_analytics_job_name }}/outputs/{{ .external_name }}"),
	// Stream Analytics Output ServiceBus Topic's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/outputs/output1
	"azurerm_stream_analytics_output_servicebus_topic": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourcegroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingjobs/{{ .parameters.stream_analytics_job_name }}/outputs/{{ .external_name }}"),
	// Stream Analytics Reference Input Blob's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/group1/providers/Microsoft.StreamAnalytics/streamingjobs/job1/inputs/input1
	"azurerm_stream_analytics_reference_input_blob": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourcegroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingjobs/{{ .parameters.stream_analytics_job_name }}/inputs/{{ .external_name }}"),
	// Stream Analytics Stream Input Blob's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
	"azurerm_stream_analytics_stream_input_blob": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourcegroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingjobs/{{ .parameters.stream_analytics_job_name }}/inputs/{{ .external_name }}"),
	// Stream Analytics Stream Input EventHub's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
	"azurerm_stream_analytics_stream_input_eventhub": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourcegroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingjobs/{{ .parameters.stream_analytics_job_name }}/inputs/{{ .external_name }}"),
	// Stream Analytics Stream Input IoTHub's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.StreamAnalytics/streamingJobs/job1/inputs/input1
	"azurerm_stream_analytics_stream_input_iothub": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourcegroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StreamAnalytics/streamingjobs/{{ .parameters.stream_analytics_job_name }}/inputs/{{ .external_name }}"),

	// subscription
	//
	// Policy Remediations can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.PolicyInsights/remediations/remediation1
	"azurerm_subscription_policy_remediation": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/providers/Microsoft.PolicyInsights/remediations/{{ .external_name }}"),

	// synapse
	//
	// Synapse Firewall Rule can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourcegroup1/providers/Microsoft.Synapse/workspaces/workspace1/firewallRules/rule1
	"azurerm_synapse_firewall_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.synapse_workspace_id }}/firewallRules/{{ .external_name }}"),
	// Synapse Azure Integration Runtimes can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/integrationRuntimes/IntegrationRuntime1
	"azurerm_synapse_integration_runtime_azure": config.TemplatedStringAsIdentifier("name", "{{ .parameters.synapse_workspace_id }}/integrationRuntimes/{{ .external_name }}"),
	// Synapse Self-hosted Integration Runtimes can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/integrationRuntimes/IntegrationRuntime1
	"azurerm_synapse_integration_runtime_self_hosted": config.TemplatedStringAsIdentifier("name", "{{ .parameters.synapse_workspace_id }}/integrationRuntimes/{{ .external_name }}"),
	// Synapse Linked Services can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/linkedServices/linkedservice1
	"azurerm_synapse_linked_service": config.TemplatedStringAsIdentifier("name", "{{ .parameters.synapse_workspace_id }}/linkedServices/{{ .external_name }}"),
	// Synapse Managed Private Endpoint can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/managedVirtualNetworks/default/managedPrivateEndpoints/endpoint1
	"azurerm_synapse_managed_private_endpoint": config.IdentifierFromProvider,
	// Synapse Private Link Hub can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/privateLinkHubs/privateLinkHub1
	"azurerm_synapse_private_link_hub": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Synapse/privateLinkHubs/{{ .external_name }}"),
	// Synapse Role Assignment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1|000000000000
	"azurerm_synapse_role_assignment": config.IdentifierFromProvider,
	// Synapse Spark Pool can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/bigDataPools/sparkPool1
	"azurerm_synapse_spark_pool": config.TemplatedStringAsIdentifier("name", "{{ .parameters.synapse_workspace_id }}/bigDataPools/{{ .external_name }}"),
	// Synapse SQL Pool can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1
	"azurerm_synapse_sql_pool": config.TemplatedStringAsIdentifier("name", "{{ .parameters.synapse_workspace_id }}/sqlPools/{{ .external_name }}"),
	// Synapse SQL Pool Extended Auditing Policys can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/extendedAuditingSettings/default
	"azurerm_synapse_sql_pool_extended_auditing_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.sql_pool_id }}/extendedAuditingSettings/default"),
	// Synapse SQL Pool Security Alert Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/securityAlertPolicies/default
	"azurerm_synapse_sql_pool_security_alert_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.sql_pool_id }}/securityAlertPolicies/default"),
	// Synapse SQL Pool Vulnerability Assessment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/vulnerabilityAssessments/default
	"azurerm_synapse_sql_pool_vulnerability_assessment": config.TemplatedStringAsIdentifier("", "{{ .parameters.sql_pool_id }}/vulnerabilityAssessments/default"),
	// Synapse SQL Pool Vulnerability Assessment Rule Baselines can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/vulnerabilityAssessments/default/rules/rule1/baselines/baseline1
	"azurerm_synapse_sql_pool_vulnerability_assessment_baseline": config.TemplatedStringAsIdentifier("", "{{ .parameters.sql_pool_vulnerability_assessment_id }}/rules/{{ .parameters.rule_name }}/baselines/{{ .external_name }}"),
	// Synapse SQL Pool Workload Classifiers can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/workloadGroups/workloadGroup1/workloadClassifiers/workloadClassifier1
	"azurerm_synapse_sql_pool_workload_classifier": config.TemplatedStringAsIdentifier("name", "{{ .parameters.workload_group_id }}/workloadClassifiers/{{ .external_name }}"),
	// Synapse SQL Pool Workload Groups can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Synapse/workspaces/workspace1/sqlPools/sqlPool1/workloadGroups/workloadGroup1
	"azurerm_synapse_sql_pool_workload_group": config.TemplatedStringAsIdentifier("name", "{{ .parameters.sql_pool_id }}/workloadGroups/{{ .external_name }}"),
	// Synapse Workspace can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1
	"azurerm_synapse_workspace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Synapse/workspaces/{{ .external_name }}"),
	// Synapse Workspace Azure AD Administrator can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.Synapse/workspaces/workspace1/administrators/activeDirectory
	"azurerm_synapse_workspace_aad_admin": config.TemplatedStringAsIdentifier("", "{{ .parameters.synapse_workspace_id }}/administrators/{{ .external_name }}"),
	// Synapse Workspace Extended Auditing Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/extendedAuditingSettings/default
	"azurerm_synapse_workspace_extended_auditing_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.synapse_workspace_id }}/extendedAuditingSettings/default"),
	// Synapse Workspace Keys can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/keys/key1
	"azurerm_synapse_workspace_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.synapse_workspace_id }}/keys/{{ .external_name }}"),
	// Synapse Workspace Security Alert Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1/securityAlertPolicies/Default
	"azurerm_synapse_workspace_security_alert_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.synapse_workspace_id }}/securityAlertPolicies/Default"),

	// sql
	//
	// MS SQL Database Extended Auditing Policies can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Sql/servers/sqlServer1/databases/db1/extendedAuditingSettings/default
	"azurerm_mssql_database_extended_auditing_policy": config.TemplatedStringAsIdentifier("", "{{ .parameters.database_id }}/extendedAuditingSettings/default"),
	// Database Vulnerability Assessment Rule Baseline can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/databases/mysqldatabase/vulnerabilityAssessments/Default/rules/VA2065/baselines/master
	"azurerm_mssql_database_vulnerability_assessment_rule_baseline": config.TemplatedStringAsIdentifier("baseline_name", "{{ .parameters.server_vulnerability_assessment_id }}/rules/{{ .parameters.rule_id }}/baselines/{{ .external_name }}"),
	// SQL Elastic Pool can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/elasticPools/myelasticpoolname
	"azurerm_mssql_elasticpool": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/elasticPools/{{ .external_name }}"),
	// SQL Firewall Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/firewallRules/rule1
	"azurerm_mssql_firewall_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.server_id }}/firewallRules/{{ .external_name }}"),
	// Elastic Job Agents can be imported using the id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1
	"azurerm_mssql_job_agent": config.IdentifierFromProvider,
	// Elastic Job Credentials can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Sql/servers/myserver1/jobAgents/myjobagent1/credentials/credential1
	"azurerm_mssql_job_credential": config.TemplatedStringAsIdentifier("name", "{{ .parameters.job_agent_id }}/credentials/{{ .external_name }}"),
	// MS SQL Server Security Alert Policy can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/securityAlertPolicies/Default
	"azurerm_mssql_server_security_alert_policy": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/securityAlertPolicies/default"),
	// MS SQL Server Vulnerability Assessment can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/acceptanceTestResourceGroup1/providers/Microsoft.Sql/servers/mssqlserver/vulnerabilityAssessments/Default
	"azurerm_mssql_server_vulnerability_assessment": config.IdentifierFromProvider,
	// A SQL Active Directory Administrator can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/administrators/activeDirectory
	"azurerm_sql_active_directory_administrator": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/administrators/activeDirectory/default"),
	// SQL Databases can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/databases/database1
	"azurerm_sql_database": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/databases/{{ .external_name }}"),
	// SQL Elastic Pool's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/elasticPools/pool1
	"azurerm_sql_elasticpool": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/elasticPools/{{ .external_name }}"),
	// SQL Firewall Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/firewallRules/rule1
	"azurerm_sql_firewall_rule": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/servers/{{ .parameters.server_name }}/firewallRules/{{ .external_name }}"),
	// SQL Managed Databases can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/managedInstances/myserver/databases/mydatabase
	"azurerm_sql_managed_database": config.TemplatedStringAsIdentifier("name", "{{ .parameters.sql_managed_instance_id }}/databases/{{ .external_name }}"),
	// SQL Servers can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/managedInstances/myserver
	"azurerm_sql_managed_instance": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/managedInstances/{{ .external_name }}"),
	// A SQL Active Directory Administrator can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/managedInstances/mymanagedinstance/administrators/activeDirectory
	"azurerm_sql_managed_instance_active_directory_administrator": config.TemplatedStringAsIdentifier("", "{{ .parameters.sql_managed_instance_id }}/administrators/activeDirectory"),
	// SQL Instance Failover Groups can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Sql/locations/Location/instanceFailoverGroups/failoverGroup1
	"azurerm_sql_managed_instance_failover_group": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Sql/locations/{{ .parameters.location }}/instanceFailoverGroups/{{ .external_name }}"),

	// sqlvirtualmachine
	//
	// Microsoft SQL Virtual Machines can be imported using the resource id
	/// subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/example1
	"azurerm_mssql_virtual_machine": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/{{ .external_name }}"),

	// static
	// Static Site Custom Domains can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Web/staticSites/my-static-site1/customDomains/name.contoso.com
	"azurerm_static_site_custom_domain": config.IdentifierFromProvider,

	// storage
	//
	// Customer Managed Keys for a Storage Account can be imported using the resource id of the Storage Account
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
	"azurerm_storage_account_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.storage_account_id }}"),
	// Storage Account Network Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Storage/storageAccounts/myaccount
	"azurerm_storage_account_network_rules": config.TemplatedStringAsIdentifier("", "{{ .parameters.storage_account_id }}"),

	// databoxedge
	//
	// DEPRECATED
	// Databox Edge Orders can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/device1/orders/default
	"azurerm_databox_edge_order": config.TemplatedStringAsIdentifier("", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{ .parameters.device_name }}/orders/default"),

	// datafactory
	//
	// Data Factory Linked Service's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
	"azurerm_data_factory_linked_service_cosmosdb_mongoapi": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/linkedservices/{{ .external_name }}"),
	// Data Factory SQL Server Table Datasets can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
	"azurerm_data_factory_dataset_sql_server": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/datasets/{{ .external_name }}"),
	// Data Factory Azure Integration Runtimes can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationruntimes/example
	"azurerm_data_factory_integration_runtime_azure": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/integrationruntimes/{{ .external_name }}"),
	// Data Factory Azure-SSIS Integration Runtimes can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/integrationruntimes/example
	"azurerm_data_factory_integration_runtime_azure_ssis": config.TemplatedStringAsIdentifier("name", "{{ .parameters.data_factory_id }}/integrationruntimes/{{ .external_name }}"),

	// databoxedge
	//
	// Databox Edge Devices can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/device1
	"azurerm_databox_edge_device": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{{ .external_name }}"),

	// databricks
	//
	// Databrick Workspaces can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/workspaces/workspace1
	"azurerm_databricks_workspace": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Databricks/workspaces/{{ .external_name }}"),
	// Databricks Workspace Customer Managed Key can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Databricks/workspaces/workspace1
	"azurerm_databricks_workspace_customer_managed_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.workspace_id }}"),

	// securityinsights
	//
	// Sentinel Scheduled Alert Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
	"azurerm_sentinel_alert_rule_scheduled": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/alertRules/{{ .external_name }}"),
	// AWS CloudTrail Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_aws_cloud_trail": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Azure Active Directory Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_azure_active_directory": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Azure Advanced Threat Protection Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_azure_advanced_threat_protection": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Azure Security Center Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_azure_security_center": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Microsoft Cloud App Security Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_microsoft_cloud_app_security": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Office 365 Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_office_365": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Threat Intelligence Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_threat_intelligence": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),

	// sentinel
	//
	// Sentinel Automation Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/automationRules/rule1
	"azurerm_sentinel_automation_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/automationRules/{{ .external_name }}"),
	// AWS S3 Data Connectors can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/dataConnectors/dc1
	"azurerm_sentinel_data_connector_aws_s3": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/dataConnectors/{{ .external_name }}"),
	// Sentinel Watchlists can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/watchlists/list1
	"azurerm_sentinel_watchlist": config.TemplatedStringAsIdentifier("name", "{{ .parameters.log_analytics_workspace_id }}/providers/Microsoft.SecurityInsights/watchlists/{{ .external_name }}"),
	// Sentinel Watchlist Items can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/watchlists/list1/watchlistItems/item1
	"azurerm_sentinel_watchlist_item": config.TemplatedStringAsIdentifier("", "{{ .parameters.watchlist_id }}/watchlistItems/{{ .external_name }}"),

	// service
	//
	// Resource Groups can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resourceGroup1/providers/Microsoft.ServiceFabric/managedClusters/clusterName1
	"azurerm_service_fabric_managed_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ServiceFabric/managedClusters/{{ .external_name }}"),

	// servicebus
	//
	// ServiceBus Namespace authorization rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/authorizationRules/rule1
	"azurerm_servicebus_namespace_authorization_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.namespace_id }}/authorizationRules/{{ .external_name }}"),
	// Service Bus DR configs can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/disasterRecoveryConfigs/config1
	"azurerm_servicebus_namespace_disaster_recovery_config": config.TemplatedStringAsIdentifier("name", "{{ .parameters.namespace_id }}/disasterRecoveryConfigs/{{ .external_name }}"),
	// Service Bus Namespace can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/sbns1
	// TODO: Check documentation, it seems there is a bug
	"azurerm_servicebus_namespace_network_rule_set": config.IdentifierFromProvider,
	// Service Bus Queue can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1/queues/snqueue1
	"azurerm_servicebus_queue": config.TemplatedStringAsIdentifier("name", "{{ .parameters.namespace_id }}/queues/{{ .external_name }}"),
	// ServiceBus Queue Authorization Rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/queues/queue1/authorizationRules/rule1
	"azurerm_servicebus_queue_authorization_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.queue_id }}/authorizationRules/{{ .external_name }}"),
	// Service Bus Subscriptions can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1/topics/sntopic1/subscriptions/sbsub1
	"azurerm_servicebus_subscription": config.TemplatedStringAsIdentifier("name", "{{ .parameters.topic_id }}/subscriptions/{{ .external_name }}"),
	// Service Bus Subscription Rule can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1/topics/sntopic1/subscriptions/sbsub1/rules/sbrule1
	"azurerm_servicebus_subscription_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.subscription_id }}/rules/{{ .external_name }}"),
	// Service Bus Topics can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceBus/namespaces/sbns1/topics/sntopic1
	"azurerm_servicebus_topic": config.TemplatedStringAsIdentifier("name", "{{ .parameters.namespace_id }}/topics/{{ .external_name }}"),
	// ServiceBus Topic authorization rules can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ServiceBus/namespaces/namespace1/topics/topic1/authorizationRules/rule1
	"azurerm_servicebus_topic_authorization_rule": config.TemplatedStringAsIdentifier("name", "{{ .parameters.topic_id }}/authorizationRules/{{ .external_name }}"),

	// servicefabric
	//
	// Service Fabric Clusters can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceFabric/clusters/cluster1
	"azurerm_service_fabric_cluster": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.ServiceFabric/clusters/{{ .external_name }}"),

	// signalrservice
	//
	// SignalR services can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/terraform-signalr/providers/Microsoft.SignalRService/signalR/tfex-signalr
	"azurerm_signalr_service": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.SignalRService/signalR/{{ .external_name }}"),
	// Network ACLs for a SignalR service can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SignalRService/signalR/signalr1
	// TODO: Check documentation, it seems there is a bug
	"azurerm_signalr_service_network_acl": config.IdentifierFromProvider,

	// solutions
	//
	// Managed Application can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Solutions/applications/app1
	"azurerm_managed_application": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Solutions/applications/{{ .external_name }}"),
	// Managed Application Definition can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Solutions/applicationDefinitions/appDefinition1
	"azurerm_managed_application_definition": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.Solutions/applicationDefinitions/{{ .external_name }}"),

	// dbformysql
	//
	// A MySQL Active Directory Administrator can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.DBforMySQL/servers/myserver/administrators/activeDirectory
	"azurerm_mysql_active_directory_administrator": config.IdentifierFromProvider,
	// MySQL Database's can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.DBforMySQL/servers/server1/databases/database1
	"azurerm_mysql_database": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DBforMySQL/servers/{{ .parameters.server_name }}/databases/{{ .external_name }}"),
	// A MySQL Server Key can be imported using the resource id of the MySQL Server Key
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DBforMySQL/servers/server1/keys/keyvaultname_key-name_keyversion
	"azurerm_mysql_server_key": config.TemplatedStringAsIdentifier("", "{{ .parameters.server_id }}/keys/{{ .external_name }}"),

	// digitaltwins
	//
	// Digital Twins Eventgrid Endpoints can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/endpoints/ep1
	"azurerm_digital_twins_endpoint_eventgrid": config.TemplatedStringAsIdentifier("name", "{{ .parameters.digital_twins_id }}/endpoints/{{ .external_name }}"),
	// Digital Twins Eventhub Endpoints can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/endpoints/ep1
	"azurerm_digital_twins_endpoint_eventhub": config.TemplatedStringAsIdentifier("name", "{{ .parameters.digital_twins_id }}/endpoints/{{ .external_name }}"),
	// Digital Twins Service Bus Endpoints can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1/endpoints/ep1
	"azurerm_digital_twins_endpoint_servicebus": config.TemplatedStringAsIdentifier("name", "{{ .parameters.digital_twins_id }}/endpoints/{{ .external_name }}"),
	// Digital Twins instances can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DigitalTwins/digitalTwinsInstances/dt1
	"azurerm_digital_twins_instance": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.DigitalTwins/digitalTwinsInstances/{{ .external_name }}"),

	// disk
	//
	// Disk Pools can be imported using the resource id
	// /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.StoragePool/diskPools/diskPool1
	"azurerm_disk_pool": config.TemplatedStringAsIdentifier("name", "/subscriptions/{{ .setup.configuration.subscription_id }}/resourceGroups/{{ .parameters.resource_group_name }}/providers/Microsoft.StoragePool/diskPools/{{ .external_name }}"),
	// iSCSI Targets can be imported using the resource id
	// /subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.StoragePool/diskPools/pool1/iscsiTargets/iscsiTarget1
	// SKIP LIST: Azure are officially halting the preview of Azure Disk Pools, and it will not be made generally available.
	"azurerm_disk_pool_iscsi_target": config.TemplatedStringAsIdentifier("name", "{{ .parameters.disks_pool_id }}/iscsiTargets/{{ .external_name }}"),
}
