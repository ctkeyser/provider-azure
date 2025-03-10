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

type AuthSettingsActiveDirectoryObservation struct {
}

type AuthSettingsActiveDirectoryParameters struct {

	// Allowed audience values to consider when validating JWTs issued by Azure Active Directory.
	// +kubebuilder:validation:Optional
	AllowedAudiences []*string `json:"allowedAudiences,omitempty" tf:"allowed_audiences,omitempty"`

	// The OAuth 2.0 client ID that was created for the app used for authentication.
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// The OAuth 2.0 client secret that was created for the app used for authentication.
	// +kubebuilder:validation:Optional
	ClientSecretSecretRef *v1.SecretKeySelector `json:"clientSecretSecretRef,omitempty" tf:"-"`
}

type AuthSettingsFacebookObservation struct {
}

type AuthSettingsFacebookParameters struct {

	// The App ID of the Facebook app used for login
	// +kubebuilder:validation:Required
	AppID *string `json:"appId" tf:"app_id,omitempty"`

	// The App Secret of the Facebook app used for Facebook login.
	// +kubebuilder:validation:Required
	AppSecretSecretRef v1.SecretKeySelector `json:"appSecretSecretRef" tf:"-"`

	// The OAuth 2.0 scopes that will be requested as part of Microsoft Account authentication. https://msdn.microsoft.com/en-us/library/dn631845.aspx
	// +kubebuilder:validation:Optional
	OauthScopes []*string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
}

type AuthSettingsGoogleObservation struct {
}

type AuthSettingsGoogleParameters struct {

	// The OAuth 2.0 client ID that was created for the app used for authentication.
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// The OAuth 2.0 client secret that was created for the app used for authentication.
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// The OAuth 2.0 scopes that will be requested as part of Microsoft Account authentication. https://msdn.microsoft.com/en-us/library/dn631845.aspx
	// +kubebuilder:validation:Optional
	OauthScopes []*string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
}

type AuthSettingsMicrosoftObservation struct {
}

type AuthSettingsMicrosoftParameters struct {

	// The OAuth 2.0 client ID that was created for the app used for authentication.
	// +kubebuilder:validation:Required
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// The OAuth 2.0 client secret that was created for the app used for authentication.
	// +kubebuilder:validation:Required
	ClientSecretSecretRef v1.SecretKeySelector `json:"clientSecretSecretRef" tf:"-"`

	// The OAuth 2.0 scopes that will be requested as part of Microsoft Account authentication. https://msdn.microsoft.com/en-us/library/dn631845.aspx
	// +kubebuilder:validation:Optional
	OauthScopes []*string `json:"oauthScopes,omitempty" tf:"oauth_scopes,omitempty"`
}

type AuthSettingsTwitterObservation struct {
}

type AuthSettingsTwitterParameters struct {

	// The OAuth 1.0a consumer key of the Twitter application used for sign-in.
	// +kubebuilder:validation:Required
	ConsumerKey *string `json:"consumerKey" tf:"consumer_key,omitempty"`

	// The OAuth 1.0a consumer secret of the Twitter application used for sign-in.
	// +kubebuilder:validation:Required
	ConsumerSecretSecretRef v1.SecretKeySelector `json:"consumerSecretSecretRef" tf:"-"`
}

type FunctionAppSlotAuthSettingsObservation struct {
}

type FunctionAppSlotAuthSettingsParameters struct {

	// An active_directory block as defined below.
	// +kubebuilder:validation:Optional
	ActiveDirectory []AuthSettingsActiveDirectoryParameters `json:"activeDirectory,omitempty" tf:"active_directory,omitempty"`

	// login parameters to send to the OpenID Connect authorization endpoint when a user logs in. Each parameter must be in the form "key=value".
	// +kubebuilder:validation:Optional
	AdditionalLoginParams map[string]*string `json:"additionalLoginParams,omitempty" tf:"additional_login_params,omitempty"`

	// External URLs that can be redirected to as part of logging in or logging out of the app.
	// +kubebuilder:validation:Optional
	AllowedExternalRedirectUrls []*string `json:"allowedExternalRedirectUrls,omitempty" tf:"allowed_external_redirect_urls,omitempty"`

	// The default provider to use when multiple providers have been set up. Possible values are AzureActiveDirectory, Facebook, Google, MicrosoftAccount and Twitter.
	// +kubebuilder:validation:Optional
	DefaultProvider *string `json:"defaultProvider,omitempty" tf:"default_provider,omitempty"`

	// Is Authentication enabled?
	// +kubebuilder:validation:Required
	Enabled *bool `json:"enabled" tf:"enabled,omitempty"`

	// A facebook block as defined below.
	// +kubebuilder:validation:Optional
	Facebook []AuthSettingsFacebookParameters `json:"facebook,omitempty" tf:"facebook,omitempty"`

	// A google block as defined below.
	// +kubebuilder:validation:Optional
	Google []AuthSettingsGoogleParameters `json:"google,omitempty" tf:"google,omitempty"`

	// Issuer URI. When using Azure Active Directory, this value is the URI of the directory tenant, e.g. https://sts.windows.net/{tenant-guid}/.
	// +kubebuilder:validation:Optional
	Issuer *string `json:"issuer,omitempty" tf:"issuer,omitempty"`

	// A microsoft block as defined below.
	// +kubebuilder:validation:Optional
	Microsoft []AuthSettingsMicrosoftParameters `json:"microsoft,omitempty" tf:"microsoft,omitempty"`

	// The runtime version of the Authentication/Authorization module.
	// +kubebuilder:validation:Optional
	RuntimeVersion *string `json:"runtimeVersion,omitempty" tf:"runtime_version,omitempty"`

	// The number of hours after session token expiration that a session token can be used to call the token refresh API. Defaults to 72.
	// +kubebuilder:validation:Optional
	TokenRefreshExtensionHours *float64 `json:"tokenRefreshExtensionHours,omitempty" tf:"token_refresh_extension_hours,omitempty"`

	// If enabled the module will durably store platform-specific security tokens that are obtained during login flows. Defaults to false.
	// +kubebuilder:validation:Optional
	TokenStoreEnabled *bool `json:"tokenStoreEnabled,omitempty" tf:"token_store_enabled,omitempty"`

	// A twitter block as defined below.
	// +kubebuilder:validation:Optional
	Twitter []AuthSettingsTwitterParameters `json:"twitter,omitempty" tf:"twitter,omitempty"`

	// The action to take when an unauthenticated client attempts to access the app. Possible values are AllowAnonymous and RedirectToLoginPage.
	// +kubebuilder:validation:Optional
	UnauthenticatedClientAction *string `json:"unauthenticatedClientAction,omitempty" tf:"unauthenticated_client_action,omitempty"`
}

type FunctionAppSlotConnectionStringObservation struct {
}

type FunctionAppSlotConnectionStringParameters struct {

	// The name of the Connection String.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The type of the Connection String. Possible values are APIHub, Custom, DocDb, EventHub, MySQL, NotificationHub, PostgreSQL, RedisCache, ServiceBus, SQLAzure and SQLServer.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// The value for the Connection String.
	// +kubebuilder:validation:Required
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

type FunctionAppSlotIdentityObservation struct {

	// The Principal ID for the Service Principal associated with the Managed Service Identity of this App Service.
	PrincipalID *string `json:"principalId,omitempty" tf:"principal_id,omitempty"`

	// The Tenant ID for the Service Principal associated with the Managed Service Identity of this App Service.
	TenantID *string `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`
}

type FunctionAppSlotIdentityParameters struct {

	// Specifies a list of user managed identity ids to be assigned. Required if type is UserAssigned.
	// +kubebuilder:validation:Optional
	IdentityIds []*string `json:"identityIds,omitempty" tf:"identity_ids,omitempty"`

	// Specifies the identity type of the Function App. Possible values are SystemAssigned (where Azure will generate a Service Principal for you), UserAssigned where you can specify the Service Principal IDs in the identity_ids field, and SystemAssigned, UserAssigned which assigns both a system managed identity as well as the specified user assigned identities.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type FunctionAppSlotObservation struct {

	// The default hostname associated with the Function App - such as mysite.azurewebsites.net
	DefaultHostName *string `json:"defaultHostname,omitempty" tf:"default_hostname,omitempty"`

	// The ID of the Function App Slot
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []FunctionAppSlotIdentityObservation `json:"identity,omitempty" tf:"identity,omitempty"`

	// The Function App kind - such as functionapp,linux,container
	Kind *string `json:"kind,omitempty" tf:"kind,omitempty"`

	// A comma separated list of outbound IP addresses - such as 52.23.25.3,52.143.43.12
	OutboundIPAddresses *string `json:"outboundIpAddresses,omitempty" tf:"outbound_ip_addresses,omitempty"`

	// A comma separated list of outbound IP addresses - such as 52.23.25.3,52.143.43.12,52.143.43.17 - not all of which are necessarily in use. Superset of outbound_ip_addresses.
	PossibleOutboundIPAddresses *string `json:"possibleOutboundIpAddresses,omitempty" tf:"possible_outbound_ip_addresses,omitempty"`

	// A site_credential block as defined below, which contains the site-level credentials used to publish to this Function App Slot.
	SiteCredential []FunctionAppSlotSiteCredentialObservation `json:"siteCredential,omitempty" tf:"site_credential,omitempty"`
}

type FunctionAppSlotParameters struct {

	// The ID of the App Service Plan within which to create this Function App Slot. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/web/v1beta1.AppServicePlan
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AppServicePlanID *string `json:"appServicePlanId,omitempty" tf:"app_service_plan_id,omitempty"`

	// Reference to a AppServicePlan in web to populate appServicePlanId.
	// +kubebuilder:validation:Optional
	AppServicePlanIDRef *v1.Reference `json:"appServicePlanIdRef,omitempty" tf:"-"`

	// Selector for a AppServicePlan in web to populate appServicePlanId.
	// +kubebuilder:validation:Optional
	AppServicePlanIDSelector *v1.Selector `json:"appServicePlanIdSelector,omitempty" tf:"-"`

	// A key-value pair of App Settings.
	// +kubebuilder:validation:Optional
	AppSettings map[string]*string `json:"appSettings,omitempty" tf:"app_settings,omitempty"`

	// An auth_settings block as defined below.
	// +kubebuilder:validation:Optional
	AuthSettings []FunctionAppSlotAuthSettingsParameters `json:"authSettings,omitempty" tf:"auth_settings,omitempty"`

	// A connection_string block as defined below.
	// +kubebuilder:validation:Optional
	ConnectionString []FunctionAppSlotConnectionStringParameters `json:"connectionString,omitempty" tf:"connection_string,omitempty"`

	// The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan.
	// +kubebuilder:validation:Optional
	DailyMemoryTimeQuota *float64 `json:"dailyMemoryTimeQuota,omitempty" tf:"daily_memory_time_quota,omitempty"`

	// Should the built-in logging of the Function App be enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	EnableBuiltinLogging *bool `json:"enableBuiltinLogging,omitempty" tf:"enable_builtin_logging,omitempty"`

	// Is the Function App enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the Function App within which to create the Function App Slot. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/web/v1beta1.FunctionApp
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("name",false)
	// +kubebuilder:validation:Optional
	FunctionAppName *string `json:"functionAppName,omitempty" tf:"function_app_name,omitempty"`

	// Reference to a FunctionApp in web to populate functionAppName.
	// +kubebuilder:validation:Optional
	FunctionAppNameRef *v1.Reference `json:"functionAppNameRef,omitempty" tf:"-"`

	// Selector for a FunctionApp in web to populate functionAppName.
	// +kubebuilder:validation:Optional
	FunctionAppNameSelector *v1.Selector `json:"functionAppNameSelector,omitempty" tf:"-"`

	// Can the Function App only be accessed via HTTPS? Defaults to false.
	// +kubebuilder:validation:Optional
	HTTPSOnly *bool `json:"httpsOnly,omitempty" tf:"https_only,omitempty"`

	// An identity block as defined below.
	// +kubebuilder:validation:Optional
	Identity []FunctionAppSlotIdentityParameters `json:"identity,omitempty" tf:"identity,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// A string indicating the Operating System type for this function app. The only possible value is linux. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// The name of the resource group in which to create the Function App Slot. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A site_config object as defined below.
	// +kubebuilder:validation:Optional
	SiteConfig []FunctionAppSlotSiteConfigParameters `json:"siteConfig,omitempty" tf:"site_config,omitempty"`

	// The access key which will be used to access the backend storage account for the Function App.
	// +kubebuilder:validation:Required
	StorageAccountAccessKeySecretRef v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef" tf:"-"`

	// The backend storage account name which will be used by the Function App (such as the dashboard, logs). Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +kubebuilder:validation:Optional
	StorageAccountName *string `json:"storageAccountName,omitempty" tf:"storage_account_name,omitempty"`

	// Reference to a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameRef *v1.Reference `json:"storageAccountNameRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageAccountName.
	// +kubebuilder:validation:Optional
	StorageAccountNameSelector *v1.Selector `json:"storageAccountNameSelector,omitempty" tf:"-"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The runtime version associated with the Function App. Defaults to ~1.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type FunctionAppSlotSiteConfigObservation struct {
}

type FunctionAppSlotSiteConfigParameters struct {

	// Should the Function App be loaded at all times? Defaults to false.
	// +kubebuilder:validation:Optional
	AlwaysOn *bool `json:"alwaysOn,omitempty" tf:"always_on,omitempty"`

	// The number of workers this function app can scale out to. Only applicable to apps on the Consumption and Premium plan.
	// +kubebuilder:validation:Optional
	AppScaleLimit *float64 `json:"appScaleLimit,omitempty" tf:"app_scale_limit,omitempty"`

	// The name of the slot to automatically swap to during deployment
	// +kubebuilder:validation:Optional
	AutoSwapSlotName *string `json:"autoSwapSlotName,omitempty" tf:"auto_swap_slot_name,omitempty"`

	// A cors block as defined below.
	// +kubebuilder:validation:Optional
	Cors []SiteConfigCorsParameters `json:"cors,omitempty" tf:"cors,omitempty"`

	// The version of the .NET framework's CLR used in this function app. Possible values are v4.0 (including .NET Core 2.1 and 3.1), v5.0 and v6.0. For more information on which .NET Framework version to use based on the runtime version you're targeting - please see this table. Defaults to v4.0.
	// +kubebuilder:validation:Optional
	DotnetFrameworkVersion *string `json:"dotnetFrameworkVersion,omitempty" tf:"dotnet_framework_version,omitempty"`

	// The number of minimum instances for this function app. Only applicable to apps on the Premium plan.
	// +kubebuilder:validation:Optional
	ElasticInstanceMinimum *float64 `json:"elasticInstanceMinimum,omitempty" tf:"elastic_instance_minimum,omitempty"`

	// State of FTP / FTPS service for this function app. Possible values include: AllAllowed, FtpsOnly and Disabled.
	// +kubebuilder:validation:Optional
	FtpsState *string `json:"ftpsState,omitempty" tf:"ftps_state,omitempty"`

	// Path which will be checked for this function app health.
	// +kubebuilder:validation:Optional
	HealthCheckPath *string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`

	// Specifies whether or not the HTTP2 protocol should be enabled. Defaults to false.
	// +kubebuilder:validation:Optional
	Http2Enabled *bool `json:"http2Enabled,omitempty" tf:"http2_enabled,omitempty"`

	// A List of objects representing IP restrictions as defined below.
	// +kubebuilder:validation:Optional
	IPRestriction []SiteConfigIPRestrictionParameters `json:"ipRestriction,omitempty" tf:"ip_restriction,omitempty"`

	// Java version hosted by the function app in Azure. Possible values are 1.8, 11 & 17 (In-Preview).
	// +kubebuilder:validation:Optional
	JavaVersion *string `json:"javaVersion,omitempty" tf:"java_version,omitempty"`

	// Linux App Framework and version for the AppService, e.g. DOCKER|(golang:latest).
	// +kubebuilder:validation:Optional
	LinuxFxVersion *string `json:"linuxFxVersion,omitempty" tf:"linux_fx_version,omitempty"`

	// The minimum supported TLS version for the function app. Possible values are 1.0, 1.1, and 1.2. Defaults to 1.2 for new function apps.
	// +kubebuilder:validation:Optional
	MinTLSVersion *string `json:"minTlsVersion,omitempty" tf:"min_tls_version,omitempty"`

	// The number of pre-warmed instances for this function app. Only affects apps on the Premium plan.
	// +kubebuilder:validation:Optional
	PreWarmedInstanceCount *float64 `json:"preWarmedInstanceCount,omitempty" tf:"pre_warmed_instance_count,omitempty"`

	// Should Runtime Scale Monitoring be enabled?. Only applicable to apps on the Premium plan. Defaults to false.
	// +kubebuilder:validation:Optional
	RuntimeScaleMonitoringEnabled *bool `json:"runtimeScaleMonitoringEnabled,omitempty" tf:"runtime_scale_monitoring_enabled,omitempty"`

	// A List of objects representing IP restrictions as defined below.
	// +kubebuilder:validation:Optional
	ScmIPRestriction []SiteConfigScmIPRestrictionParameters `json:"scmIpRestriction,omitempty" tf:"scm_ip_restriction,omitempty"`

	// The type of Source Control used by this function App. Valid values include: BitBucketGit, BitBucketHg, CodePlexGit, CodePlexHg, Dropbox, ExternalGit, ExternalHg, GitHub, LocalGit, None (default), OneDrive, Tfs, VSO, and VSTSRM.
	// +kubebuilder:validation:Optional
	ScmType *string `json:"scmType,omitempty" tf:"scm_type,omitempty"`

	// IP security restrictions for scm to use main. Defaults to false.
	// +kubebuilder:validation:Optional
	ScmUseMainIPRestriction *bool `json:"scmUseMainIpRestriction,omitempty" tf:"scm_use_main_ip_restriction,omitempty"`

	// Should the Function App run in 32 bit mode, rather than 64 bit mode? Defaults to true.
	// +kubebuilder:validation:Optional
	Use32BitWorkerProcess *bool `json:"use32BitWorkerProcess,omitempty" tf:"use_32_bit_worker_process,omitempty"`

	// Is the Function App enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	VnetRouteAllEnabled *bool `json:"vnetRouteAllEnabled,omitempty" tf:"vnet_route_all_enabled,omitempty"`

	// Should WebSockets be enabled?
	// +kubebuilder:validation:Optional
	WebsocketsEnabled *bool `json:"websocketsEnabled,omitempty" tf:"websockets_enabled,omitempty"`
}

type FunctionAppSlotSiteCredentialObservation struct {

	// The password associated with the username, which can be used to publish to this App Service.
	Password *string `json:"password,omitempty" tf:"password,omitempty"`

	// The username which can be used to publish to this App Service
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type FunctionAppSlotSiteCredentialParameters struct {
}

type IPRestrictionHeadersObservation struct {
}

type IPRestrictionHeadersParameters struct {

	// A list of allowed Azure FrontDoor IDs in UUID notation with a maximum of 8.
	// +kubebuilder:validation:Optional
	XAzureFdid []*string `json:"xAzureFdid,omitempty" tf:"x_azure_fdid"`

	// A list to allow the Azure FrontDoor health probe header. Only allowed value is "1".
	// +kubebuilder:validation:Optional
	XFdHealthProbe []*string `json:"xFdHealthProbe,omitempty" tf:"x_fd_health_probe"`

	// A list of allowed 'X-Forwarded-For' IPs in CIDR notation with a maximum of 8
	// +kubebuilder:validation:Optional
	XForwardedFor []*string `json:"xForwardedFor,omitempty" tf:"x_forwarded_for"`

	// A list of allowed 'X-Forwarded-Host' domains with a maximum of 8.
	// +kubebuilder:validation:Optional
	XForwardedHost []*string `json:"xForwardedHost,omitempty" tf:"x_forwarded_host"`
}

type SiteConfigCorsObservation struct {
}

type SiteConfigCorsParameters struct {

	// A list of origins which should be able to make cross-origin calls. * can be used to allow all calls.
	// +kubebuilder:validation:Required
	AllowedOrigins []*string `json:"allowedOrigins" tf:"allowed_origins,omitempty"`

	// Are credentials supported?
	// +kubebuilder:validation:Optional
	SupportCredentials *bool `json:"supportCredentials,omitempty" tf:"support_credentials,omitempty"`
}

type SiteConfigIPRestrictionObservation struct {
}

type SiteConfigIPRestrictionParameters struct {

	// Allow or Deny access for this IP range. Defaults to Allow.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action"`

	// The headers for this specific scm_ip_restriction as defined below.
	// +kubebuilder:validation:Optional
	Headers []IPRestrictionHeadersParameters `json:"headers,omitempty" tf:"headers"`

	// The IP Address used for this IP Restriction in CIDR notation.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`

	// The name for this IP Restriction.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// The priority for this IP Restriction. Restrictions are enforced in priority order. By default, priority is set to 65000 if not specified.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority"`

	// The Service Tag used for this IP Restriction.
	// +kubebuilder:validation:Optional
	ServiceTag *string `json:"serviceTag,omitempty" tf:"service_tag"`

	// The Virtual Network Subnet ID used for this IP Restriction.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetID *string `json:"virtualNetworkSubnetId,omitempty" tf:"virtual_network_subnet_id"`

	// Reference to a Subnet in network to populate virtualNetworkSubnetId.
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIDRef *v1.Reference `json:"virtualNetworkSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate virtualNetworkSubnetId.
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIDSelector *v1.Selector `json:"virtualNetworkSubnetIdSelector,omitempty" tf:"-"`
}

type SiteConfigScmIPRestrictionHeadersObservation struct {
}

type SiteConfigScmIPRestrictionHeadersParameters struct {

	// A list of allowed Azure FrontDoor IDs in UUID notation with a maximum of 8.
	// +kubebuilder:validation:Optional
	XAzureFdid []*string `json:"xAzureFdid,omitempty" tf:"x_azure_fdid"`

	// A list to allow the Azure FrontDoor health probe header. Only allowed value is "1".
	// +kubebuilder:validation:Optional
	XFdHealthProbe []*string `json:"xFdHealthProbe,omitempty" tf:"x_fd_health_probe"`

	// A list of allowed 'X-Forwarded-For' IPs in CIDR notation with a maximum of 8
	// +kubebuilder:validation:Optional
	XForwardedFor []*string `json:"xForwardedFor,omitempty" tf:"x_forwarded_for"`

	// A list of allowed 'X-Forwarded-Host' domains with a maximum of 8.
	// +kubebuilder:validation:Optional
	XForwardedHost []*string `json:"xForwardedHost,omitempty" tf:"x_forwarded_host"`
}

type SiteConfigScmIPRestrictionObservation struct {
}

type SiteConfigScmIPRestrictionParameters struct {

	// Allow or Deny access for this IP range. Defaults to Allow.
	// +kubebuilder:validation:Optional
	Action *string `json:"action,omitempty" tf:"action"`

	// The headers for this specific scm_ip_restriction as defined below.
	// +kubebuilder:validation:Optional
	Headers []SiteConfigScmIPRestrictionHeadersParameters `json:"headers,omitempty" tf:"headers"`

	// The IP Address used for this IP Restriction in CIDR notation.
	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`

	// The name for this IP Restriction.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// The priority for this IP Restriction. Restrictions are enforced in priority order. By default, priority is set to 65000 if not specified.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority"`

	// The Service Tag used for this IP Restriction.
	// +kubebuilder:validation:Optional
	ServiceTag *string `json:"serviceTag,omitempty" tf:"service_tag"`

	// The Virtual Network Subnet ID used for this IP Restriction.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetID *string `json:"virtualNetworkSubnetId,omitempty" tf:"virtual_network_subnet_id"`

	// Reference to a Subnet in network to populate virtualNetworkSubnetId.
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIDRef *v1.Reference `json:"virtualNetworkSubnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate virtualNetworkSubnetId.
	// +kubebuilder:validation:Optional
	VirtualNetworkSubnetIDSelector *v1.Selector `json:"virtualNetworkSubnetIdSelector,omitempty" tf:"-"`
}

// FunctionAppSlotSpec defines the desired state of FunctionAppSlot
type FunctionAppSlotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionAppSlotParameters `json:"forProvider"`
}

// FunctionAppSlotStatus defines the observed state of FunctionAppSlot.
type FunctionAppSlotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionAppSlotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionAppSlot is the Schema for the FunctionAppSlots API. Manages a Function App Deployment Slot.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type FunctionAppSlot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionAppSlotSpec   `json:"spec"`
	Status            FunctionAppSlotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionAppSlotList contains a list of FunctionAppSlots
type FunctionAppSlotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FunctionAppSlot `json:"items"`
}

// Repository type metadata.
var (
	FunctionAppSlot_Kind             = "FunctionAppSlot"
	FunctionAppSlot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FunctionAppSlot_Kind}.String()
	FunctionAppSlot_KindAPIVersion   = FunctionAppSlot_Kind + "." + CRDGroupVersion.String()
	FunctionAppSlot_GroupVersionKind = CRDGroupVersion.WithKind(FunctionAppSlot_Kind)
)

func init() {
	SchemeBuilder.Register(&FunctionAppSlot{}, &FunctionAppSlotList{})
}
