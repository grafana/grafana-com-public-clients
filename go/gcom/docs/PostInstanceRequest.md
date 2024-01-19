# PostInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertQuota** | Pointer to **int32** |  | [optional] 
**Alerts** | Pointer to **bool** |  | [optional] 
**BillingEndDate** | Pointer to **time.Time** |  | [optional] 
**BillingStartDate** | Pointer to **time.Time** |  | [optional] 
**DashboardQuota** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Graphite** | Pointer to **bool** |  | [optional] 
**HlInstanceId** | Pointer to **int32** |  | [optional] 
**Hosted** | Pointer to **bool** |  | [optional] 
**Incident** | Pointer to **bool** |  | [optional] 
**IssueLink** | Pointer to **string** |  | [optional] 
**K6OrgId** | Pointer to **int32** |  | [optional] 
**Logs** | Pointer to **bool** |  | [optional] 
**MachineLearning** | Pointer to **bool** |  | [optional] 
**MachineLearningLogsToken** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NoRestart** | Pointer to **bool** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Prometheus** | Pointer to **bool** |  | [optional] 
**ReasonType** | Pointer to **string** |  | [optional] 
**SkipOrgConflictCheck** | Pointer to **bool** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Trial** | Pointer to **bool** |  | [optional] 
**TrialExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UserQuota** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewPostInstanceRequest

`func NewPostInstanceRequest() *PostInstanceRequest`

NewPostInstanceRequest instantiates a new PostInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostInstanceRequestWithDefaults

`func NewPostInstanceRequestWithDefaults() *PostInstanceRequest`

NewPostInstanceRequestWithDefaults instantiates a new PostInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertQuota

`func (o *PostInstanceRequest) GetAlertQuota() int32`

GetAlertQuota returns the AlertQuota field if non-nil, zero value otherwise.

### GetAlertQuotaOk

`func (o *PostInstanceRequest) GetAlertQuotaOk() (*int32, bool)`

GetAlertQuotaOk returns a tuple with the AlertQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertQuota

`func (o *PostInstanceRequest) SetAlertQuota(v int32)`

SetAlertQuota sets AlertQuota field to given value.

### HasAlertQuota

`func (o *PostInstanceRequest) HasAlertQuota() bool`

HasAlertQuota returns a boolean if a field has been set.

### GetAlerts

`func (o *PostInstanceRequest) GetAlerts() bool`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *PostInstanceRequest) GetAlertsOk() (*bool, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *PostInstanceRequest) SetAlerts(v bool)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *PostInstanceRequest) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetBillingEndDate

`func (o *PostInstanceRequest) GetBillingEndDate() time.Time`

GetBillingEndDate returns the BillingEndDate field if non-nil, zero value otherwise.

### GetBillingEndDateOk

`func (o *PostInstanceRequest) GetBillingEndDateOk() (*time.Time, bool)`

GetBillingEndDateOk returns a tuple with the BillingEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEndDate

`func (o *PostInstanceRequest) SetBillingEndDate(v time.Time)`

SetBillingEndDate sets BillingEndDate field to given value.

### HasBillingEndDate

`func (o *PostInstanceRequest) HasBillingEndDate() bool`

HasBillingEndDate returns a boolean if a field has been set.

### GetBillingStartDate

`func (o *PostInstanceRequest) GetBillingStartDate() time.Time`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *PostInstanceRequest) GetBillingStartDateOk() (*time.Time, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *PostInstanceRequest) SetBillingStartDate(v time.Time)`

SetBillingStartDate sets BillingStartDate field to given value.

### HasBillingStartDate

`func (o *PostInstanceRequest) HasBillingStartDate() bool`

HasBillingStartDate returns a boolean if a field has been set.

### GetDashboardQuota

`func (o *PostInstanceRequest) GetDashboardQuota() int32`

GetDashboardQuota returns the DashboardQuota field if non-nil, zero value otherwise.

### GetDashboardQuotaOk

`func (o *PostInstanceRequest) GetDashboardQuotaOk() (*int32, bool)`

GetDashboardQuotaOk returns a tuple with the DashboardQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardQuota

`func (o *PostInstanceRequest) SetDashboardQuota(v int32)`

SetDashboardQuota sets DashboardQuota field to given value.

### HasDashboardQuota

`func (o *PostInstanceRequest) HasDashboardQuota() bool`

HasDashboardQuota returns a boolean if a field has been set.

### GetDescription

`func (o *PostInstanceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostInstanceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostInstanceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostInstanceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGateway

`func (o *PostInstanceRequest) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PostInstanceRequest) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PostInstanceRequest) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *PostInstanceRequest) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGraphite

`func (o *PostInstanceRequest) GetGraphite() bool`

GetGraphite returns the Graphite field if non-nil, zero value otherwise.

### GetGraphiteOk

`func (o *PostInstanceRequest) GetGraphiteOk() (*bool, bool)`

GetGraphiteOk returns a tuple with the Graphite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphite

`func (o *PostInstanceRequest) SetGraphite(v bool)`

SetGraphite sets Graphite field to given value.

### HasGraphite

`func (o *PostInstanceRequest) HasGraphite() bool`

HasGraphite returns a boolean if a field has been set.

### GetHlInstanceId

`func (o *PostInstanceRequest) GetHlInstanceId() int32`

GetHlInstanceId returns the HlInstanceId field if non-nil, zero value otherwise.

### GetHlInstanceIdOk

`func (o *PostInstanceRequest) GetHlInstanceIdOk() (*int32, bool)`

GetHlInstanceIdOk returns a tuple with the HlInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlInstanceId

`func (o *PostInstanceRequest) SetHlInstanceId(v int32)`

SetHlInstanceId sets HlInstanceId field to given value.

### HasHlInstanceId

`func (o *PostInstanceRequest) HasHlInstanceId() bool`

HasHlInstanceId returns a boolean if a field has been set.

### GetHosted

`func (o *PostInstanceRequest) GetHosted() bool`

GetHosted returns the Hosted field if non-nil, zero value otherwise.

### GetHostedOk

`func (o *PostInstanceRequest) GetHostedOk() (*bool, bool)`

GetHostedOk returns a tuple with the Hosted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosted

`func (o *PostInstanceRequest) SetHosted(v bool)`

SetHosted sets Hosted field to given value.

### HasHosted

`func (o *PostInstanceRequest) HasHosted() bool`

HasHosted returns a boolean if a field has been set.

### GetIncident

`func (o *PostInstanceRequest) GetIncident() bool`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *PostInstanceRequest) GetIncidentOk() (*bool, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *PostInstanceRequest) SetIncident(v bool)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *PostInstanceRequest) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetIssueLink

`func (o *PostInstanceRequest) GetIssueLink() string`

GetIssueLink returns the IssueLink field if non-nil, zero value otherwise.

### GetIssueLinkOk

`func (o *PostInstanceRequest) GetIssueLinkOk() (*string, bool)`

GetIssueLinkOk returns a tuple with the IssueLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueLink

`func (o *PostInstanceRequest) SetIssueLink(v string)`

SetIssueLink sets IssueLink field to given value.

### HasIssueLink

`func (o *PostInstanceRequest) HasIssueLink() bool`

HasIssueLink returns a boolean if a field has been set.

### GetK6OrgId

`func (o *PostInstanceRequest) GetK6OrgId() int32`

GetK6OrgId returns the K6OrgId field if non-nil, zero value otherwise.

### GetK6OrgIdOk

`func (o *PostInstanceRequest) GetK6OrgIdOk() (*int32, bool)`

GetK6OrgIdOk returns a tuple with the K6OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6OrgId

`func (o *PostInstanceRequest) SetK6OrgId(v int32)`

SetK6OrgId sets K6OrgId field to given value.

### HasK6OrgId

`func (o *PostInstanceRequest) HasK6OrgId() bool`

HasK6OrgId returns a boolean if a field has been set.

### GetLogs

`func (o *PostInstanceRequest) GetLogs() bool`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PostInstanceRequest) GetLogsOk() (*bool, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PostInstanceRequest) SetLogs(v bool)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *PostInstanceRequest) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMachineLearning

`func (o *PostInstanceRequest) GetMachineLearning() bool`

GetMachineLearning returns the MachineLearning field if non-nil, zero value otherwise.

### GetMachineLearningOk

`func (o *PostInstanceRequest) GetMachineLearningOk() (*bool, bool)`

GetMachineLearningOk returns a tuple with the MachineLearning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLearning

`func (o *PostInstanceRequest) SetMachineLearning(v bool)`

SetMachineLearning sets MachineLearning field to given value.

### HasMachineLearning

`func (o *PostInstanceRequest) HasMachineLearning() bool`

HasMachineLearning returns a boolean if a field has been set.

### GetMachineLearningLogsToken

`func (o *PostInstanceRequest) GetMachineLearningLogsToken() string`

GetMachineLearningLogsToken returns the MachineLearningLogsToken field if non-nil, zero value otherwise.

### GetMachineLearningLogsTokenOk

`func (o *PostInstanceRequest) GetMachineLearningLogsTokenOk() (*string, bool)`

GetMachineLearningLogsTokenOk returns a tuple with the MachineLearningLogsToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLearningLogsToken

`func (o *PostInstanceRequest) SetMachineLearningLogsToken(v string)`

SetMachineLearningLogsToken sets MachineLearningLogsToken field to given value.

### HasMachineLearningLogsToken

`func (o *PostInstanceRequest) HasMachineLearningLogsToken() bool`

HasMachineLearningLogsToken returns a boolean if a field has been set.

### GetName

`func (o *PostInstanceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostInstanceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostInstanceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostInstanceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNoRestart

`func (o *PostInstanceRequest) GetNoRestart() bool`

GetNoRestart returns the NoRestart field if non-nil, zero value otherwise.

### GetNoRestartOk

`func (o *PostInstanceRequest) GetNoRestartOk() (*bool, bool)`

GetNoRestartOk returns a tuple with the NoRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoRestart

`func (o *PostInstanceRequest) SetNoRestart(v bool)`

SetNoRestart sets NoRestart field to given value.

### HasNoRestart

`func (o *PostInstanceRequest) HasNoRestart() bool`

HasNoRestart returns a boolean if a field has been set.

### GetPlan

`func (o *PostInstanceRequest) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *PostInstanceRequest) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *PostInstanceRequest) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *PostInstanceRequest) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetPrometheus

`func (o *PostInstanceRequest) GetPrometheus() bool`

GetPrometheus returns the Prometheus field if non-nil, zero value otherwise.

### GetPrometheusOk

`func (o *PostInstanceRequest) GetPrometheusOk() (*bool, bool)`

GetPrometheusOk returns a tuple with the Prometheus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrometheus

`func (o *PostInstanceRequest) SetPrometheus(v bool)`

SetPrometheus sets Prometheus field to given value.

### HasPrometheus

`func (o *PostInstanceRequest) HasPrometheus() bool`

HasPrometheus returns a boolean if a field has been set.

### GetReasonType

`func (o *PostInstanceRequest) GetReasonType() string`

GetReasonType returns the ReasonType field if non-nil, zero value otherwise.

### GetReasonTypeOk

`func (o *PostInstanceRequest) GetReasonTypeOk() (*string, bool)`

GetReasonTypeOk returns a tuple with the ReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonType

`func (o *PostInstanceRequest) SetReasonType(v string)`

SetReasonType sets ReasonType field to given value.

### HasReasonType

`func (o *PostInstanceRequest) HasReasonType() bool`

HasReasonType returns a boolean if a field has been set.

### GetSkipOrgConflictCheck

`func (o *PostInstanceRequest) GetSkipOrgConflictCheck() bool`

GetSkipOrgConflictCheck returns the SkipOrgConflictCheck field if non-nil, zero value otherwise.

### GetSkipOrgConflictCheckOk

`func (o *PostInstanceRequest) GetSkipOrgConflictCheckOk() (*bool, bool)`

GetSkipOrgConflictCheckOk returns a tuple with the SkipOrgConflictCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipOrgConflictCheck

`func (o *PostInstanceRequest) SetSkipOrgConflictCheck(v bool)`

SetSkipOrgConflictCheck sets SkipOrgConflictCheck field to given value.

### HasSkipOrgConflictCheck

`func (o *PostInstanceRequest) HasSkipOrgConflictCheck() bool`

HasSkipOrgConflictCheck returns a boolean if a field has been set.

### GetSlug

`func (o *PostInstanceRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PostInstanceRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PostInstanceRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PostInstanceRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTrial

`func (o *PostInstanceRequest) GetTrial() bool`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *PostInstanceRequest) GetTrialOk() (*bool, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *PostInstanceRequest) SetTrial(v bool)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *PostInstanceRequest) HasTrial() bool`

HasTrial returns a boolean if a field has been set.

### GetTrialExpiresAt

`func (o *PostInstanceRequest) GetTrialExpiresAt() time.Time`

GetTrialExpiresAt returns the TrialExpiresAt field if non-nil, zero value otherwise.

### GetTrialExpiresAtOk

`func (o *PostInstanceRequest) GetTrialExpiresAtOk() (*time.Time, bool)`

GetTrialExpiresAtOk returns a tuple with the TrialExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialExpiresAt

`func (o *PostInstanceRequest) SetTrialExpiresAt(v time.Time)`

SetTrialExpiresAt sets TrialExpiresAt field to given value.

### HasTrialExpiresAt

`func (o *PostInstanceRequest) HasTrialExpiresAt() bool`

HasTrialExpiresAt returns a boolean if a field has been set.

### GetUrl

`func (o *PostInstanceRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PostInstanceRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PostInstanceRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PostInstanceRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserQuota

`func (o *PostInstanceRequest) GetUserQuota() int32`

GetUserQuota returns the UserQuota field if non-nil, zero value otherwise.

### GetUserQuotaOk

`func (o *PostInstanceRequest) GetUserQuotaOk() (*int32, bool)`

GetUserQuotaOk returns a tuple with the UserQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuota

`func (o *PostInstanceRequest) SetUserQuota(v int32)`

SetUserQuota sets UserQuota field to given value.

### HasUserQuota

`func (o *PostInstanceRequest) HasUserQuota() bool`

HasUserQuota returns a boolean if a field has been set.

### GetVersion

`func (o *PostInstanceRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PostInstanceRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PostInstanceRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PostInstanceRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


