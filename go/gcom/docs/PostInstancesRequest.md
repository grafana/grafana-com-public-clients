# PostInstancesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminUserInstance** | Pointer to **bool** |  | [optional] 
**Alerts** | Pointer to **bool** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**DeleteProtection** | Pointer to **bool** |  | [optional] [default to false]
**Description** | Pointer to **string** |  | [optional] 
**Graphite** | Pointer to **bool** |  | [optional] 
**HlInstanceId** | Pointer to **int32** |  | [optional] 
**Hosted** | Pointer to **bool** |  | [optional] 
**K6OrgId** | Pointer to **int32** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Logs** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Org** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Plugins** | Pointer to **[]string** |  | [optional] 
**Prometheus** | Pointer to **bool** |  | [optional] 
**PublicInstance** | Pointer to **bool** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UsernameOrEmail** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**WaitForReadiness** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewPostInstancesRequest

`func NewPostInstancesRequest(name string, ) *PostInstancesRequest`

NewPostInstancesRequest instantiates a new PostInstancesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostInstancesRequestWithDefaults

`func NewPostInstancesRequestWithDefaults() *PostInstancesRequest`

NewPostInstancesRequestWithDefaults instantiates a new PostInstancesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminUserInstance

`func (o *PostInstancesRequest) GetAdminUserInstance() bool`

GetAdminUserInstance returns the AdminUserInstance field if non-nil, zero value otherwise.

### GetAdminUserInstanceOk

`func (o *PostInstancesRequest) GetAdminUserInstanceOk() (*bool, bool)`

GetAdminUserInstanceOk returns a tuple with the AdminUserInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUserInstance

`func (o *PostInstancesRequest) SetAdminUserInstance(v bool)`

SetAdminUserInstance sets AdminUserInstance field to given value.

### HasAdminUserInstance

`func (o *PostInstancesRequest) HasAdminUserInstance() bool`

HasAdminUserInstance returns a boolean if a field has been set.

### GetAlerts

`func (o *PostInstancesRequest) GetAlerts() bool`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *PostInstancesRequest) GetAlertsOk() (*bool, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *PostInstancesRequest) SetAlerts(v bool)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *PostInstancesRequest) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetCluster

`func (o *PostInstancesRequest) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *PostInstancesRequest) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *PostInstancesRequest) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *PostInstancesRequest) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *PostInstancesRequest) GetDeleteProtection() bool`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *PostInstancesRequest) GetDeleteProtectionOk() (*bool, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *PostInstancesRequest) SetDeleteProtection(v bool)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *PostInstancesRequest) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *PostInstancesRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostInstancesRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostInstancesRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostInstancesRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGraphite

`func (o *PostInstancesRequest) GetGraphite() bool`

GetGraphite returns the Graphite field if non-nil, zero value otherwise.

### GetGraphiteOk

`func (o *PostInstancesRequest) GetGraphiteOk() (*bool, bool)`

GetGraphiteOk returns a tuple with the Graphite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphite

`func (o *PostInstancesRequest) SetGraphite(v bool)`

SetGraphite sets Graphite field to given value.

### HasGraphite

`func (o *PostInstancesRequest) HasGraphite() bool`

HasGraphite returns a boolean if a field has been set.

### GetHlInstanceId

`func (o *PostInstancesRequest) GetHlInstanceId() int32`

GetHlInstanceId returns the HlInstanceId field if non-nil, zero value otherwise.

### GetHlInstanceIdOk

`func (o *PostInstancesRequest) GetHlInstanceIdOk() (*int32, bool)`

GetHlInstanceIdOk returns a tuple with the HlInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlInstanceId

`func (o *PostInstancesRequest) SetHlInstanceId(v int32)`

SetHlInstanceId sets HlInstanceId field to given value.

### HasHlInstanceId

`func (o *PostInstancesRequest) HasHlInstanceId() bool`

HasHlInstanceId returns a boolean if a field has been set.

### GetHosted

`func (o *PostInstancesRequest) GetHosted() bool`

GetHosted returns the Hosted field if non-nil, zero value otherwise.

### GetHostedOk

`func (o *PostInstancesRequest) GetHostedOk() (*bool, bool)`

GetHostedOk returns a tuple with the Hosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosted

`func (o *PostInstancesRequest) SetHosted(v bool)`

SetHosted sets Hosted field to given value.

### HasHosted

`func (o *PostInstancesRequest) HasHosted() bool`

HasHosted returns a boolean if a field has been set.

### GetK6OrgId

`func (o *PostInstancesRequest) GetK6OrgId() int32`

GetK6OrgId returns the K6OrgId field if non-nil, zero value otherwise.

### GetK6OrgIdOk

`func (o *PostInstancesRequest) GetK6OrgIdOk() (*int32, bool)`

GetK6OrgIdOk returns a tuple with the K6OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6OrgId

`func (o *PostInstancesRequest) SetK6OrgId(v int32)`

SetK6OrgId sets K6OrgId field to given value.

### HasK6OrgId

`func (o *PostInstancesRequest) HasK6OrgId() bool`

HasK6OrgId returns a boolean if a field has been set.

### GetLabels

`func (o *PostInstancesRequest) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PostInstancesRequest) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PostInstancesRequest) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PostInstancesRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLogs

`func (o *PostInstancesRequest) GetLogs() bool`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PostInstancesRequest) GetLogsOk() (*bool, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PostInstancesRequest) SetLogs(v bool)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *PostInstancesRequest) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetName

`func (o *PostInstancesRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostInstancesRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostInstancesRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOrg

`func (o *PostInstancesRequest) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PostInstancesRequest) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PostInstancesRequest) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PostInstancesRequest) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPlan

`func (o *PostInstancesRequest) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PostInstancesRequest) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PostInstancesRequest) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *PostInstancesRequest) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPlugins

`func (o *PostInstancesRequest) GetPlugins() []string`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *PostInstancesRequest) GetPluginsOk() (*[]string, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *PostInstancesRequest) SetPlugins(v []string)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *PostInstancesRequest) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetPrometheus

`func (o *PostInstancesRequest) GetPrometheus() bool`

GetPrometheus returns the Prometheus field if non-nil, zero value otherwise.

### GetPrometheusOk

`func (o *PostInstancesRequest) GetPrometheusOk() (*bool, bool)`

GetPrometheusOk returns a tuple with the Prometheus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheus

`func (o *PostInstancesRequest) SetPrometheus(v bool)`

SetPrometheus sets Prometheus field to given value.

### HasPrometheus

`func (o *PostInstancesRequest) HasPrometheus() bool`

HasPrometheus returns a boolean if a field has been set.

### GetPublicInstance

`func (o *PostInstancesRequest) GetPublicInstance() bool`

GetPublicInstance returns the PublicInstance field if non-nil, zero value otherwise.

### GetPublicInstanceOk

`func (o *PostInstancesRequest) GetPublicInstanceOk() (*bool, bool)`

GetPublicInstanceOk returns a tuple with the PublicInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicInstance

`func (o *PostInstancesRequest) SetPublicInstance(v bool)`

SetPublicInstance sets PublicInstance field to given value.

### HasPublicInstance

`func (o *PostInstancesRequest) HasPublicInstance() bool`

HasPublicInstance returns a boolean if a field has been set.

### GetRegion

`func (o *PostInstancesRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostInstancesRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostInstancesRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PostInstancesRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSlug

`func (o *PostInstancesRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PostInstancesRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PostInstancesRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PostInstancesRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetUrl

`func (o *PostInstancesRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PostInstancesRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PostInstancesRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PostInstancesRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsernameOrEmail

`func (o *PostInstancesRequest) GetUsernameOrEmail() string`

GetUsernameOrEmail returns the UsernameOrEmail field if non-nil, zero value otherwise.

### GetUsernameOrEmailOk

`func (o *PostInstancesRequest) GetUsernameOrEmailOk() (*string, bool)`

GetUsernameOrEmailOk returns a tuple with the UsernameOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameOrEmail

`func (o *PostInstancesRequest) SetUsernameOrEmail(v string)`

SetUsernameOrEmail sets UsernameOrEmail field to given value.

### HasUsernameOrEmail

`func (o *PostInstancesRequest) HasUsernameOrEmail() bool`

HasUsernameOrEmail returns a boolean if a field has been set.

### GetVersion

`func (o *PostInstancesRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PostInstancesRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PostInstancesRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PostInstancesRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWaitForReadiness

`func (o *PostInstancesRequest) GetWaitForReadiness() bool`

GetWaitForReadiness returns the WaitForReadiness field if non-nil, zero value otherwise.

### GetWaitForReadinessOk

`func (o *PostInstancesRequest) GetWaitForReadinessOk() (*bool, bool)`

GetWaitForReadinessOk returns a tuple with the WaitForReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitForReadiness

`func (o *PostInstancesRequest) SetWaitForReadiness(v bool)`

SetWaitForReadiness sets WaitForReadiness field to given value.

### HasWaitForReadiness

`func (o *PostInstancesRequest) HasWaitForReadiness() bool`

HasWaitForReadiness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


