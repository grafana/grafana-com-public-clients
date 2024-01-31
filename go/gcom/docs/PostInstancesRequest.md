# PostInstancesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminUserInstance** | Pointer to **bool** |  | [optional] 
**AlertQuota** | Pointer to **int32** |  | [optional] 
**Alerts** | Pointer to **bool** |  | [optional] 
**AmCluster** | Pointer to **string** |  | [optional] 
**BillingEndDate** | Pointer to **time.Time** |  | [optional] 
**BillingStartDate** | Pointer to **time.Time** |  | [optional] 
**Cluster** | Pointer to **string** |  | [optional] 
**CreateTemporaryLicenseIfMissing** | Pointer to **bool** |  | [optional] 
**DashboardQuota** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Graphite** | Pointer to **bool** |  | [optional] 
**HlCluster** | Pointer to **string** |  | [optional] 
**HlInstanceId** | Pointer to **int32** |  | [optional] 
**HmGraphiteCluster** | Pointer to **string** |  | [optional] 
**HmGraphiteType** | Pointer to **string** |  | [optional] 
**HmPromCluster** | Pointer to **string** |  | [optional] 
**Hosted** | Pointer to **bool** |  | [optional] 
**HtCluster** | Pointer to **string** |  | [optional] 
**Incident** | Pointer to **bool** |  | [optional] 
**IssueLink** | Pointer to **string** |  | [optional] 
**K6OrgId** | Pointer to **int32** |  | [optional] 
**LlmIsOptIn** | Pointer to **bool** |  | [optional] 
**LlmOptInChangedBy** | Pointer to **string** |  | [optional] 
**Logs** | Pointer to **bool** |  | [optional] 
**MachineLearning** | Pointer to **bool** |  | [optional] 
**MachineLearningLogsToken** | Pointer to **string** |  | [optional] 
**Multitenant** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Org** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Plugins** | Pointer to **[]string** |  | [optional] 
**Prometheus** | Pointer to **bool** |  | [optional] 
**PublicInstance** | Pointer to **bool** |  | [optional] 
**ReasonType** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SkipOrgConflictCheck** | Pointer to **bool** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Trial** | Pointer to **bool** |  | [optional] 
**TrialExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UsernameOrEmail** | Pointer to **string** |  | [optional] 
**UserQuota** | Pointer to **int32** |  | [optional] 
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

### GetAlertQuota

`func (o *PostInstancesRequest) GetAlertQuota() int32`

GetAlertQuota returns the AlertQuota field if non-nil, zero value otherwise.

### GetAlertQuotaOk

`func (o *PostInstancesRequest) GetAlertQuotaOk() (*int32, bool)`

GetAlertQuotaOk returns a tuple with the AlertQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertQuota

`func (o *PostInstancesRequest) SetAlertQuota(v int32)`

SetAlertQuota sets AlertQuota field to given value.

### HasAlertQuota

`func (o *PostInstancesRequest) HasAlertQuota() bool`

HasAlertQuota returns a boolean if a field has been set.

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

### GetAmCluster

`func (o *PostInstancesRequest) GetAmCluster() string`

GetAmCluster returns the AmCluster field if non-nil, zero value otherwise.

### GetAmClusterOk

`func (o *PostInstancesRequest) GetAmClusterOk() (*string, bool)`

GetAmClusterOk returns a tuple with the AmCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmCluster

`func (o *PostInstancesRequest) SetAmCluster(v string)`

SetAmCluster sets AmCluster field to given value.

### HasAmCluster

`func (o *PostInstancesRequest) HasAmCluster() bool`

HasAmCluster returns a boolean if a field has been set.

### GetBillingEndDate

`func (o *PostInstancesRequest) GetBillingEndDate() time.Time`

GetBillingEndDate returns the BillingEndDate field if non-nil, zero value otherwise.

### GetBillingEndDateOk

`func (o *PostInstancesRequest) GetBillingEndDateOk() (*time.Time, bool)`

GetBillingEndDateOk returns a tuple with the BillingEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEndDate

`func (o *PostInstancesRequest) SetBillingEndDate(v time.Time)`

SetBillingEndDate sets BillingEndDate field to given value.

### HasBillingEndDate

`func (o *PostInstancesRequest) HasBillingEndDate() bool`

HasBillingEndDate returns a boolean if a field has been set.

### GetBillingStartDate

`func (o *PostInstancesRequest) GetBillingStartDate() time.Time`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *PostInstancesRequest) GetBillingStartDateOk() (*time.Time, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *PostInstancesRequest) SetBillingStartDate(v time.Time)`

SetBillingStartDate sets BillingStartDate field to given value.

### HasBillingStartDate

`func (o *PostInstancesRequest) HasBillingStartDate() bool`

HasBillingStartDate returns a boolean if a field has been set.

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

### GetCreateTemporaryLicenseIfMissing

`func (o *PostInstancesRequest) GetCreateTemporaryLicenseIfMissing() bool`

GetCreateTemporaryLicenseIfMissing returns the CreateTemporaryLicenseIfMissing field if non-nil, zero value otherwise.

### GetCreateTemporaryLicenseIfMissingOk

`func (o *PostInstancesRequest) GetCreateTemporaryLicenseIfMissingOk() (*bool, bool)`

GetCreateTemporaryLicenseIfMissingOk returns a tuple with the CreateTemporaryLicenseIfMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTemporaryLicenseIfMissing

`func (o *PostInstancesRequest) SetCreateTemporaryLicenseIfMissing(v bool)`

SetCreateTemporaryLicenseIfMissing sets CreateTemporaryLicenseIfMissing field to given value.

### HasCreateTemporaryLicenseIfMissing

`func (o *PostInstancesRequest) HasCreateTemporaryLicenseIfMissing() bool`

HasCreateTemporaryLicenseIfMissing returns a boolean if a field has been set.

### GetDashboardQuota

`func (o *PostInstancesRequest) GetDashboardQuota() int32`

GetDashboardQuota returns the DashboardQuota field if non-nil, zero value otherwise.

### GetDashboardQuotaOk

`func (o *PostInstancesRequest) GetDashboardQuotaOk() (*int32, bool)`

GetDashboardQuotaOk returns a tuple with the DashboardQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardQuota

`func (o *PostInstancesRequest) SetDashboardQuota(v int32)`

SetDashboardQuota sets DashboardQuota field to given value.

### HasDashboardQuota

`func (o *PostInstancesRequest) HasDashboardQuota() bool`

HasDashboardQuota returns a boolean if a field has been set.

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

### GetHlCluster

`func (o *PostInstancesRequest) GetHlCluster() string`

GetHlCluster returns the HlCluster field if non-nil, zero value otherwise.

### GetHlClusterOk

`func (o *PostInstancesRequest) GetHlClusterOk() (*string, bool)`

GetHlClusterOk returns a tuple with the HlCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlCluster

`func (o *PostInstancesRequest) SetHlCluster(v string)`

SetHlCluster sets HlCluster field to given value.

### HasHlCluster

`func (o *PostInstancesRequest) HasHlCluster() bool`

HasHlCluster returns a boolean if a field has been set.

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

### GetHmGraphiteCluster

`func (o *PostInstancesRequest) GetHmGraphiteCluster() string`

GetHmGraphiteCluster returns the HmGraphiteCluster field if non-nil, zero value otherwise.

### GetHmGraphiteClusterOk

`func (o *PostInstancesRequest) GetHmGraphiteClusterOk() (*string, bool)`

GetHmGraphiteClusterOk returns a tuple with the HmGraphiteCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmGraphiteCluster

`func (o *PostInstancesRequest) SetHmGraphiteCluster(v string)`

SetHmGraphiteCluster sets HmGraphiteCluster field to given value.

### HasHmGraphiteCluster

`func (o *PostInstancesRequest) HasHmGraphiteCluster() bool`

HasHmGraphiteCluster returns a boolean if a field has been set.

### GetHmGraphiteType

`func (o *PostInstancesRequest) GetHmGraphiteType() string`

GetHmGraphiteType returns the HmGraphiteType field if non-nil, zero value otherwise.

### GetHmGraphiteTypeOk

`func (o *PostInstancesRequest) GetHmGraphiteTypeOk() (*string, bool)`

GetHmGraphiteTypeOk returns a tuple with the HmGraphiteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmGraphiteType

`func (o *PostInstancesRequest) SetHmGraphiteType(v string)`

SetHmGraphiteType sets HmGraphiteType field to given value.

### HasHmGraphiteType

`func (o *PostInstancesRequest) HasHmGraphiteType() bool`

HasHmGraphiteType returns a boolean if a field has been set.

### GetHmPromCluster

`func (o *PostInstancesRequest) GetHmPromCluster() string`

GetHmPromCluster returns the HmPromCluster field if non-nil, zero value otherwise.

### GetHmPromClusterOk

`func (o *PostInstancesRequest) GetHmPromClusterOk() (*string, bool)`

GetHmPromClusterOk returns a tuple with the HmPromCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmPromCluster

`func (o *PostInstancesRequest) SetHmPromCluster(v string)`

SetHmPromCluster sets HmPromCluster field to given value.

### HasHmPromCluster

`func (o *PostInstancesRequest) HasHmPromCluster() bool`

HasHmPromCluster returns a boolean if a field has been set.

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

### GetHtCluster

`func (o *PostInstancesRequest) GetHtCluster() string`

GetHtCluster returns the HtCluster field if non-nil, zero value otherwise.

### GetHtClusterOk

`func (o *PostInstancesRequest) GetHtClusterOk() (*string, bool)`

GetHtClusterOk returns a tuple with the HtCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtCluster

`func (o *PostInstancesRequest) SetHtCluster(v string)`

SetHtCluster sets HtCluster field to given value.

### HasHtCluster

`func (o *PostInstancesRequest) HasHtCluster() bool`

HasHtCluster returns a boolean if a field has been set.

### GetIncident

`func (o *PostInstancesRequest) GetIncident() bool`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *PostInstancesRequest) GetIncidentOk() (*bool, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *PostInstancesRequest) SetIncident(v bool)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *PostInstancesRequest) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetIssueLink

`func (o *PostInstancesRequest) GetIssueLink() string`

GetIssueLink returns the IssueLink field if non-nil, zero value otherwise.

### GetIssueLinkOk

`func (o *PostInstancesRequest) GetIssueLinkOk() (*string, bool)`

GetIssueLinkOk returns a tuple with the IssueLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueLink

`func (o *PostInstancesRequest) SetIssueLink(v string)`

SetIssueLink sets IssueLink field to given value.

### HasIssueLink

`func (o *PostInstancesRequest) HasIssueLink() bool`

HasIssueLink returns a boolean if a field has been set.

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

### GetLlmIsOptIn

`func (o *PostInstancesRequest) GetLlmIsOptIn() bool`

GetLlmIsOptIn returns the LlmIsOptIn field if non-nil, zero value otherwise.

### GetLlmIsOptInOk

`func (o *PostInstancesRequest) GetLlmIsOptInOk() (*bool, bool)`

GetLlmIsOptInOk returns a tuple with the LlmIsOptIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmIsOptIn

`func (o *PostInstancesRequest) SetLlmIsOptIn(v bool)`

SetLlmIsOptIn sets LlmIsOptIn field to given value.

### HasLlmIsOptIn

`func (o *PostInstancesRequest) HasLlmIsOptIn() bool`

HasLlmIsOptIn returns a boolean if a field has been set.

### GetLlmOptInChangedBy

`func (o *PostInstancesRequest) GetLlmOptInChangedBy() string`

GetLlmOptInChangedBy returns the LlmOptInChangedBy field if non-nil, zero value otherwise.

### GetLlmOptInChangedByOk

`func (o *PostInstancesRequest) GetLlmOptInChangedByOk() (*string, bool)`

GetLlmOptInChangedByOk returns a tuple with the LlmOptInChangedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmOptInChangedBy

`func (o *PostInstancesRequest) SetLlmOptInChangedBy(v string)`

SetLlmOptInChangedBy sets LlmOptInChangedBy field to given value.

### HasLlmOptInChangedBy

`func (o *PostInstancesRequest) HasLlmOptInChangedBy() bool`

HasLlmOptInChangedBy returns a boolean if a field has been set.

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

### GetMachineLearning

`func (o *PostInstancesRequest) GetMachineLearning() bool`

GetMachineLearning returns the MachineLearning field if non-nil, zero value otherwise.

### GetMachineLearningOk

`func (o *PostInstancesRequest) GetMachineLearningOk() (*bool, bool)`

GetMachineLearningOk returns a tuple with the MachineLearning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLearning

`func (o *PostInstancesRequest) SetMachineLearning(v bool)`

SetMachineLearning sets MachineLearning field to given value.

### HasMachineLearning

`func (o *PostInstancesRequest) HasMachineLearning() bool`

HasMachineLearning returns a boolean if a field has been set.

### GetMachineLearningLogsToken

`func (o *PostInstancesRequest) GetMachineLearningLogsToken() string`

GetMachineLearningLogsToken returns the MachineLearningLogsToken field if non-nil, zero value otherwise.

### GetMachineLearningLogsTokenOk

`func (o *PostInstancesRequest) GetMachineLearningLogsTokenOk() (*string, bool)`

GetMachineLearningLogsTokenOk returns a tuple with the MachineLearningLogsToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLearningLogsToken

`func (o *PostInstancesRequest) SetMachineLearningLogsToken(v string)`

SetMachineLearningLogsToken sets MachineLearningLogsToken field to given value.

### HasMachineLearningLogsToken

`func (o *PostInstancesRequest) HasMachineLearningLogsToken() bool`

HasMachineLearningLogsToken returns a boolean if a field has been set.

### GetMultitenant

`func (o *PostInstancesRequest) GetMultitenant() bool`

GetMultitenant returns the Multitenant field if non-nil, zero value otherwise.

### GetMultitenantOk

`func (o *PostInstancesRequest) GetMultitenantOk() (*bool, bool)`

GetMultitenantOk returns a tuple with the Multitenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenant

`func (o *PostInstancesRequest) SetMultitenant(v bool)`

SetMultitenant sets Multitenant field to given value.

### HasMultitenant

`func (o *PostInstancesRequest) HasMultitenant() bool`

HasMultitenant returns a boolean if a field has been set.

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

### GetReasonType

`func (o *PostInstancesRequest) GetReasonType() string`

GetReasonType returns the ReasonType field if non-nil, zero value otherwise.

### GetReasonTypeOk

`func (o *PostInstancesRequest) GetReasonTypeOk() (*string, bool)`

GetReasonTypeOk returns a tuple with the ReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonType

`func (o *PostInstancesRequest) SetReasonType(v string)`

SetReasonType sets ReasonType field to given value.

### HasReasonType

`func (o *PostInstancesRequest) HasReasonType() bool`

HasReasonType returns a boolean if a field has been set.

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

### GetSkipOrgConflictCheck

`func (o *PostInstancesRequest) GetSkipOrgConflictCheck() bool`

GetSkipOrgConflictCheck returns the SkipOrgConflictCheck field if non-nil, zero value otherwise.

### GetSkipOrgConflictCheckOk

`func (o *PostInstancesRequest) GetSkipOrgConflictCheckOk() (*bool, bool)`

GetSkipOrgConflictCheckOk returns a tuple with the SkipOrgConflictCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOrgConflictCheck

`func (o *PostInstancesRequest) SetSkipOrgConflictCheck(v bool)`

SetSkipOrgConflictCheck sets SkipOrgConflictCheck field to given value.

### HasSkipOrgConflictCheck

`func (o *PostInstancesRequest) HasSkipOrgConflictCheck() bool`

HasSkipOrgConflictCheck returns a boolean if a field has been set.

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

### GetTrial

`func (o *PostInstancesRequest) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *PostInstancesRequest) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *PostInstancesRequest) SetTrial(v bool)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *PostInstancesRequest) HasTrial() bool`

HasTrial returns a boolean if a field has been set.

### GetTrialExpiresAt

`func (o *PostInstancesRequest) GetTrialExpiresAt() time.Time`

GetTrialExpiresAt returns the TrialExpiresAt field if non-nil, zero value otherwise.

### GetTrialExpiresAtOk

`func (o *PostInstancesRequest) GetTrialExpiresAtOk() (*time.Time, bool)`

GetTrialExpiresAtOk returns a tuple with the TrialExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialExpiresAt

`func (o *PostInstancesRequest) SetTrialExpiresAt(v time.Time)`

SetTrialExpiresAt sets TrialExpiresAt field to given value.

### HasTrialExpiresAt

`func (o *PostInstancesRequest) HasTrialExpiresAt() bool`

HasTrialExpiresAt returns a boolean if a field has been set.

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

### GetUserQuota

`func (o *PostInstancesRequest) GetUserQuota() int32`

GetUserQuota returns the UserQuota field if non-nil, zero value otherwise.

### GetUserQuotaOk

`func (o *PostInstancesRequest) GetUserQuotaOk() (*int32, bool)`

GetUserQuotaOk returns a tuple with the UserQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuota

`func (o *PostInstancesRequest) SetUserQuota(v int32)`

SetUserQuota sets UserQuota field to given value.

### HasUserQuota

`func (o *PostInstancesRequest) HasUserQuota() bool`

HasUserQuota returns a boolean if a field has been set.

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


