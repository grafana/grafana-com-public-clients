# FormattedApiInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertCnt** | **float32** |  | 
**AlertQuota** | **float32** |  | 
**AmInstanceGeneratorUrl** | **string** |  | 
**AmInstanceGeneratorUrlDatasource** | **string** |  | 
**AmInstanceId** | **float32** |  | 
**AmInstanceName** | **string** |  | 
**AmInstanceStatus** | **string** |  | 
**AmInstanceUrl** | **string** |  | 
**AmInstanceClusterId** | **float32** |  | 
**BillingActiveUsers** | **float32** |  | 
**BillingEndDate** | **NullableString** |  | 
**BillingGrafanaActiveUsers** | **float32** |  | 
**BillingOnCallActiveUsers** | **float32** |  | 
**BillingStartDate** | **string** |  | 
**ClusterId** | **float32** |  | 
**ClusterName** | **string** |  | 
**ClusterSlug** | **string** |  | 
**CreatedAt** | **string** |  | 
**CreatedBy** | **string** |  | 
**CurrentActiveAdminUsers** | **float32** |  | 
**CurrentActiveEditorUsers** | **float32** |  | 
**CurrentActiveUsers** | **float32** |  | 
**CurrentActiveViewerUsers** | **float32** |  | 
**CustomAuth** | **bool** |  | 
**CustomDomain** | **bool** |  | 
**DailyAdminCnt** | **float32** |  | 
**DailyEditorCnt** | **float32** |  | 
**DailyUserCnt** | **float32** |  | 
**DailyViewerCnt** | **float32** |  | 
**DashboardCnt** | **float32** |  | 
**DashboardQuota** | **float32** |  | 
**DatasourceCnts** | **map[string]interface{}** |  | 
**Description** | **string** |  | 
**Gateway** | **string** |  | 
**HlInstanceCurrentUsage** | **float32** |  | 
**HlInstanceBillingUsage** | **float32** |  | 
**HlInstanceId** | **float32** |  | 
**HlInstanceName** | **string** |  | 
**HlInstanceStatus** | **string** |  | 
**HlInstanceUrl** | **string** |  | 
**HlInstanceClusterId** | **float32** |  | 
**HmInstanceGraphiteCurrentUsage** | **float32** |  | 
**HmInstanceGraphiteBillingUsage** | **float32** |  | 
**HmInstanceGraphiteId** | **float32** |  | 
**HmInstanceGraphiteName** | **string** |  | 
**HmInstanceGraphiteStatus** | **string** |  | 
**HmInstanceGraphiteType** | **string** |  | 
**HmInstanceGraphiteUrl** | **string** |  | 
**HmInstancePromClusterId** | **float32** |  | 
**HmInstancePromCurrentActiveSeries** | **NullableFloat32** |  | 
**HmInstancePromCurrentUsage** | **float32** |  | 
**HmInstancePromBillingUsage** | **float32** |  | 
**HmInstancePromId** | **float32** |  | 
**HmInstancePromName** | **string** |  | 
**HmInstancePromStatus** | **string** |  | 
**HmInstancePromUrl** | **string** |  | 
**HmInstanceGraphiteClusterId** | **float32** |  | 
**HtInstanceId** | **float32** |  | 
**HtInstanceName** | **string** |  | 
**HtInstanceStatus** | **string** |  | 
**HtInstanceUrl** | **string** |  | 
**HtInstanceClusterId** | **float32** |  | 
**HtInstanceCurrentUsage** | **float32** |  | 
**HtInstanceBillingUsage** | **float32** |  | 
**HpInstanceId** | **float32** |  | 
**HpInstanceName** | **string** |  | 
**HpInstanceStatus** | **string** |  | 
**HpInstanceUrl** | **string** |  | 
**HpInstanceClusterId** | **float32** |  | 
**HpInstanceCurrentUsage** | **float32** |  | 
**HpInstanceBillingUsage** | **float32** |  | 
**Id** | **float32** |  | 
**Incident** | **float32** |  | 
**Labels** | Pointer to **map[string]interface{}** |  | [optional] 
**MachineLearning** | **float32** |  | 
**Name** | **string** |  | 
**OrgId** | **float32** |  | 
**OrgName** | **string** |  | 
**OrgSlug** | **string** |  | 
**Plan** | **string** |  | 
**PlanName** | **string** |  | 
**RegionId** | **float32** |  | 
**RegionSlug** | **string** |  | 
**RegionPublicName** | **string** |  | 
**Provider** | **string** |  | 
**ProviderRegion** | **string** |  | 
**RunningVersion** | **string** |  | 
**Slug** | **string** |  | 
**Ssl** | **bool** |  | 
**Status** | **string** |  | 
**Support** | **bool** |  | 
**Trial** | **float32** |  | 
**TrialExpiresAt** | **NullableString** |  | 
**UpdatedAt** | **NullableString** |  | 
**UpdatedBy** | **NullableString** |  | 
**Url** | **string** |  | 
**UserQuota** | **float32** |  | 
**Version** | **string** |  | 
**VersionIssueLink** | Pointer to **string** |  | [optional] 
**AgentManagementInstanceId** | **float32** |  | 
**AgentManagementInstanceUrl** | **string** |  | 
**AgentManagementInstanceName** | **string** |  | 
**AgentManagementInstanceStatus** | **string** |  | 
**AgentManagementInstanceClusterId** | **float32** |  | 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | **string** |  | 
**K6OrgId** | **NullableFloat32** |  | 
**MachineLearningLogsToken** | **string** |  | 
**UsageStatsId** | **string** |  | 
**RegionStackStateServiceUrl** | **string** |  | 
**RegionSyntheticMonitoringApiUrl** | **string** |  | 
**RegionInsightsApiUrl** | **string** |  | 
**RegionIntegrationsApiUrl** | **string** |  | 
**RegionHostedExportersApiUrl** | **string** |  | 
**RegionMachineLearningApiUrl** | **string** |  | 
**RegionLLMGatewayUrl** | **string** |  | 
**RegionAssistantUrl** | **string** |  | 
**Links** | [**[]LinksInner1**](LinksInner1.md) |  | 
**DeleteProtection** | **bool** |  | 

## Methods

### NewFormattedApiInstance

`func NewFormattedApiInstance(alertCnt float32, alertQuota float32, amInstanceGeneratorUrl string, amInstanceGeneratorUrlDatasource string, amInstanceId float32, amInstanceName string, amInstanceStatus string, amInstanceUrl string, amInstanceClusterId float32, billingActiveUsers float32, billingEndDate NullableString, billingGrafanaActiveUsers float32, billingOnCallActiveUsers float32, billingStartDate string, clusterId float32, clusterName string, clusterSlug string, createdAt string, createdBy string, currentActiveAdminUsers float32, currentActiveEditorUsers float32, currentActiveUsers float32, currentActiveViewerUsers float32, customAuth bool, customDomain bool, dailyAdminCnt float32, dailyEditorCnt float32, dailyUserCnt float32, dailyViewerCnt float32, dashboardCnt float32, dashboardQuota float32, datasourceCnts map[string]interface{}, description string, gateway string, hlInstanceCurrentUsage float32, hlInstanceBillingUsage float32, hlInstanceId float32, hlInstanceName string, hlInstanceStatus string, hlInstanceUrl string, hlInstanceClusterId float32, hmInstanceGraphiteCurrentUsage float32, hmInstanceGraphiteBillingUsage float32, hmInstanceGraphiteId float32, hmInstanceGraphiteName string, hmInstanceGraphiteStatus string, hmInstanceGraphiteType string, hmInstanceGraphiteUrl string, hmInstancePromClusterId float32, hmInstancePromCurrentActiveSeries NullableFloat32, hmInstancePromCurrentUsage float32, hmInstancePromBillingUsage float32, hmInstancePromId float32, hmInstancePromName string, hmInstancePromStatus string, hmInstancePromUrl string, hmInstanceGraphiteClusterId float32, htInstanceId float32, htInstanceName string, htInstanceStatus string, htInstanceUrl string, htInstanceClusterId float32, htInstanceCurrentUsage float32, htInstanceBillingUsage float32, hpInstanceId float32, hpInstanceName string, hpInstanceStatus string, hpInstanceUrl string, hpInstanceClusterId float32, hpInstanceCurrentUsage float32, hpInstanceBillingUsage float32, id float32, incident float32, machineLearning float32, name string, orgId float32, orgName string, orgSlug string, plan string, planName string, regionId float32, regionSlug string, regionPublicName string, provider string, providerRegion string, runningVersion string, slug string, ssl bool, status string, support bool, trial float32, trialExpiresAt NullableString, updatedAt NullableString, updatedBy NullableString, url string, userQuota float32, version string, agentManagementInstanceId float32, agentManagementInstanceUrl string, agentManagementInstanceName string, agentManagementInstanceStatus string, agentManagementInstanceClusterId float32, type_ string, k6OrgId NullableFloat32, machineLearningLogsToken string, usageStatsId string, regionStackStateServiceUrl string, regionSyntheticMonitoringApiUrl string, regionInsightsApiUrl string, regionIntegrationsApiUrl string, regionHostedExportersApiUrl string, regionMachineLearningApiUrl string, regionLLMGatewayUrl string, regionAssistantUrl string, links []LinksInner1, deleteProtection bool, ) *FormattedApiInstance`

NewFormattedApiInstance instantiates a new FormattedApiInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedApiInstanceWithDefaults

`func NewFormattedApiInstanceWithDefaults() *FormattedApiInstance`

NewFormattedApiInstanceWithDefaults instantiates a new FormattedApiInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertCnt

`func (o *FormattedApiInstance) GetAlertCnt() float32`

GetAlertCnt returns the AlertCnt field if non-nil, zero value otherwise.

### GetAlertCntOk

`func (o *FormattedApiInstance) GetAlertCntOk() (*float32, bool)`

GetAlertCntOk returns a tuple with the AlertCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCnt

`func (o *FormattedApiInstance) SetAlertCnt(v float32)`

SetAlertCnt sets AlertCnt field to given value.


### GetAlertQuota

`func (o *FormattedApiInstance) GetAlertQuota() float32`

GetAlertQuota returns the AlertQuota field if non-nil, zero value otherwise.

### GetAlertQuotaOk

`func (o *FormattedApiInstance) GetAlertQuotaOk() (*float32, bool)`

GetAlertQuotaOk returns a tuple with the AlertQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertQuota

`func (o *FormattedApiInstance) SetAlertQuota(v float32)`

SetAlertQuota sets AlertQuota field to given value.


### GetAmInstanceGeneratorUrl

`func (o *FormattedApiInstance) GetAmInstanceGeneratorUrl() string`

GetAmInstanceGeneratorUrl returns the AmInstanceGeneratorUrl field if non-nil, zero value otherwise.

### GetAmInstanceGeneratorUrlOk

`func (o *FormattedApiInstance) GetAmInstanceGeneratorUrlOk() (*string, bool)`

GetAmInstanceGeneratorUrlOk returns a tuple with the AmInstanceGeneratorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmInstanceGeneratorUrl

`func (o *FormattedApiInstance) SetAmInstanceGeneratorUrl(v string)`

SetAmInstanceGeneratorUrl sets AmInstanceGeneratorUrl field to given value.


### GetAmInstanceGeneratorUrlDatasource

`func (o *FormattedApiInstance) GetAmInstanceGeneratorUrlDatasource() string`

GetAmInstanceGeneratorUrlDatasource returns the AmInstanceGeneratorUrlDatasource field if non-nil, zero value otherwise.

### GetAmInstanceGeneratorUrlDatasourceOk

`func (o *FormattedApiInstance) GetAmInstanceGeneratorUrlDatasourceOk() (*string, bool)`

GetAmInstanceGeneratorUrlDatasourceOk returns a tuple with the AmInstanceGeneratorUrlDatasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmInstanceGeneratorUrlDatasource

`func (o *FormattedApiInstance) SetAmInstanceGeneratorUrlDatasource(v string)`

SetAmInstanceGeneratorUrlDatasource sets AmInstanceGeneratorUrlDatasource field to given value.


### GetAmInstanceId

`func (o *FormattedApiInstance) GetAmInstanceId() float32`

GetAmInstanceId returns the AmInstanceId field if non-nil, zero value otherwise.

### GetAmInstanceIdOk

`func (o *FormattedApiInstance) GetAmInstanceIdOk() (*float32, bool)`

GetAmInstanceIdOk returns a tuple with the AmInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmInstanceId

`func (o *FormattedApiInstance) SetAmInstanceId(v float32)`

SetAmInstanceId sets AmInstanceId field to given value.


### GetAmInstanceName

`func (o *FormattedApiInstance) GetAmInstanceName() string`

GetAmInstanceName returns the AmInstanceName field if non-nil, zero value otherwise.

### GetAmInstanceNameOk

`func (o *FormattedApiInstance) GetAmInstanceNameOk() (*string, bool)`

GetAmInstanceNameOk returns a tuple with the AmInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmInstanceName

`func (o *FormattedApiInstance) SetAmInstanceName(v string)`

SetAmInstanceName sets AmInstanceName field to given value.


### GetAmInstanceStatus

`func (o *FormattedApiInstance) GetAmInstanceStatus() string`

GetAmInstanceStatus returns the AmInstanceStatus field if non-nil, zero value otherwise.

### GetAmInstanceStatusOk

`func (o *FormattedApiInstance) GetAmInstanceStatusOk() (*string, bool)`

GetAmInstanceStatusOk returns a tuple with the AmInstanceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmInstanceStatus

`func (o *FormattedApiInstance) SetAmInstanceStatus(v string)`

SetAmInstanceStatus sets AmInstanceStatus field to given value.


### GetAmInstanceUrl

`func (o *FormattedApiInstance) GetAmInstanceUrl() string`

GetAmInstanceUrl returns the AmInstanceUrl field if non-nil, zero value otherwise.

### GetAmInstanceUrlOk

`func (o *FormattedApiInstance) GetAmInstanceUrlOk() (*string, bool)`

GetAmInstanceUrlOk returns a tuple with the AmInstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmInstanceUrl

`func (o *FormattedApiInstance) SetAmInstanceUrl(v string)`

SetAmInstanceUrl sets AmInstanceUrl field to given value.


### GetAmInstanceClusterId

`func (o *FormattedApiInstance) GetAmInstanceClusterId() float32`

GetAmInstanceClusterId returns the AmInstanceClusterId field if non-nil, zero value otherwise.

### GetAmInstanceClusterIdOk

`func (o *FormattedApiInstance) GetAmInstanceClusterIdOk() (*float32, bool)`

GetAmInstanceClusterIdOk returns a tuple with the AmInstanceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmInstanceClusterId

`func (o *FormattedApiInstance) SetAmInstanceClusterId(v float32)`

SetAmInstanceClusterId sets AmInstanceClusterId field to given value.


### GetBillingActiveUsers

`func (o *FormattedApiInstance) GetBillingActiveUsers() float32`

GetBillingActiveUsers returns the BillingActiveUsers field if non-nil, zero value otherwise.

### GetBillingActiveUsersOk

`func (o *FormattedApiInstance) GetBillingActiveUsersOk() (*float32, bool)`

GetBillingActiveUsersOk returns a tuple with the BillingActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingActiveUsers

`func (o *FormattedApiInstance) SetBillingActiveUsers(v float32)`

SetBillingActiveUsers sets BillingActiveUsers field to given value.


### GetBillingEndDate

`func (o *FormattedApiInstance) GetBillingEndDate() string`

GetBillingEndDate returns the BillingEndDate field if non-nil, zero value otherwise.

### GetBillingEndDateOk

`func (o *FormattedApiInstance) GetBillingEndDateOk() (*string, bool)`

GetBillingEndDateOk returns a tuple with the BillingEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingEndDate

`func (o *FormattedApiInstance) SetBillingEndDate(v string)`

SetBillingEndDate sets BillingEndDate field to given value.


### SetBillingEndDateNil

`func (o *FormattedApiInstance) SetBillingEndDateNil(b bool)`

 SetBillingEndDateNil sets the value for BillingEndDate to be an explicit nil

### UnsetBillingEndDate
`func (o *FormattedApiInstance) UnsetBillingEndDate()`

UnsetBillingEndDate ensures that no value is present for BillingEndDate, not even an explicit nil
### GetBillingGrafanaActiveUsers

`func (o *FormattedApiInstance) GetBillingGrafanaActiveUsers() float32`

GetBillingGrafanaActiveUsers returns the BillingGrafanaActiveUsers field if non-nil, zero value otherwise.

### GetBillingGrafanaActiveUsersOk

`func (o *FormattedApiInstance) GetBillingGrafanaActiveUsersOk() (*float32, bool)`

GetBillingGrafanaActiveUsersOk returns a tuple with the BillingGrafanaActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGrafanaActiveUsers

`func (o *FormattedApiInstance) SetBillingGrafanaActiveUsers(v float32)`

SetBillingGrafanaActiveUsers sets BillingGrafanaActiveUsers field to given value.


### GetBillingOnCallActiveUsers

`func (o *FormattedApiInstance) GetBillingOnCallActiveUsers() float32`

GetBillingOnCallActiveUsers returns the BillingOnCallActiveUsers field if non-nil, zero value otherwise.

### GetBillingOnCallActiveUsersOk

`func (o *FormattedApiInstance) GetBillingOnCallActiveUsersOk() (*float32, bool)`

GetBillingOnCallActiveUsersOk returns a tuple with the BillingOnCallActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingOnCallActiveUsers

`func (o *FormattedApiInstance) SetBillingOnCallActiveUsers(v float32)`

SetBillingOnCallActiveUsers sets BillingOnCallActiveUsers field to given value.


### GetBillingStartDate

`func (o *FormattedApiInstance) GetBillingStartDate() string`

GetBillingStartDate returns the BillingStartDate field if non-nil, zero value otherwise.

### GetBillingStartDateOk

`func (o *FormattedApiInstance) GetBillingStartDateOk() (*string, bool)`

GetBillingStartDateOk returns a tuple with the BillingStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingStartDate

`func (o *FormattedApiInstance) SetBillingStartDate(v string)`

SetBillingStartDate sets BillingStartDate field to given value.


### GetClusterId

`func (o *FormattedApiInstance) GetClusterId() float32`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *FormattedApiInstance) GetClusterIdOk() (*float32, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *FormattedApiInstance) SetClusterId(v float32)`

SetClusterId sets ClusterId field to given value.


### GetClusterName

`func (o *FormattedApiInstance) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *FormattedApiInstance) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *FormattedApiInstance) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetClusterSlug

`func (o *FormattedApiInstance) GetClusterSlug() string`

GetClusterSlug returns the ClusterSlug field if non-nil, zero value otherwise.

### GetClusterSlugOk

`func (o *FormattedApiInstance) GetClusterSlugOk() (*string, bool)`

GetClusterSlugOk returns a tuple with the ClusterSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSlug

`func (o *FormattedApiInstance) SetClusterSlug(v string)`

SetClusterSlug sets ClusterSlug field to given value.


### GetCreatedAt

`func (o *FormattedApiInstance) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormattedApiInstance) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormattedApiInstance) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *FormattedApiInstance) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FormattedApiInstance) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FormattedApiInstance) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCurrentActiveAdminUsers

`func (o *FormattedApiInstance) GetCurrentActiveAdminUsers() float32`

GetCurrentActiveAdminUsers returns the CurrentActiveAdminUsers field if non-nil, zero value otherwise.

### GetCurrentActiveAdminUsersOk

`func (o *FormattedApiInstance) GetCurrentActiveAdminUsersOk() (*float32, bool)`

GetCurrentActiveAdminUsersOk returns a tuple with the CurrentActiveAdminUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentActiveAdminUsers

`func (o *FormattedApiInstance) SetCurrentActiveAdminUsers(v float32)`

SetCurrentActiveAdminUsers sets CurrentActiveAdminUsers field to given value.


### GetCurrentActiveEditorUsers

`func (o *FormattedApiInstance) GetCurrentActiveEditorUsers() float32`

GetCurrentActiveEditorUsers returns the CurrentActiveEditorUsers field if non-nil, zero value otherwise.

### GetCurrentActiveEditorUsersOk

`func (o *FormattedApiInstance) GetCurrentActiveEditorUsersOk() (*float32, bool)`

GetCurrentActiveEditorUsersOk returns a tuple with the CurrentActiveEditorUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentActiveEditorUsers

`func (o *FormattedApiInstance) SetCurrentActiveEditorUsers(v float32)`

SetCurrentActiveEditorUsers sets CurrentActiveEditorUsers field to given value.


### GetCurrentActiveUsers

`func (o *FormattedApiInstance) GetCurrentActiveUsers() float32`

GetCurrentActiveUsers returns the CurrentActiveUsers field if non-nil, zero value otherwise.

### GetCurrentActiveUsersOk

`func (o *FormattedApiInstance) GetCurrentActiveUsersOk() (*float32, bool)`

GetCurrentActiveUsersOk returns a tuple with the CurrentActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentActiveUsers

`func (o *FormattedApiInstance) SetCurrentActiveUsers(v float32)`

SetCurrentActiveUsers sets CurrentActiveUsers field to given value.


### GetCurrentActiveViewerUsers

`func (o *FormattedApiInstance) GetCurrentActiveViewerUsers() float32`

GetCurrentActiveViewerUsers returns the CurrentActiveViewerUsers field if non-nil, zero value otherwise.

### GetCurrentActiveViewerUsersOk

`func (o *FormattedApiInstance) GetCurrentActiveViewerUsersOk() (*float32, bool)`

GetCurrentActiveViewerUsersOk returns a tuple with the CurrentActiveViewerUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentActiveViewerUsers

`func (o *FormattedApiInstance) SetCurrentActiveViewerUsers(v float32)`

SetCurrentActiveViewerUsers sets CurrentActiveViewerUsers field to given value.


### GetCustomAuth

`func (o *FormattedApiInstance) GetCustomAuth() bool`

GetCustomAuth returns the CustomAuth field if non-nil, zero value otherwise.

### GetCustomAuthOk

`func (o *FormattedApiInstance) GetCustomAuthOk() (*bool, bool)`

GetCustomAuthOk returns a tuple with the CustomAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAuth

`func (o *FormattedApiInstance) SetCustomAuth(v bool)`

SetCustomAuth sets CustomAuth field to given value.


### GetCustomDomain

`func (o *FormattedApiInstance) GetCustomDomain() bool`

GetCustomDomain returns the CustomDomain field if non-nil, zero value otherwise.

### GetCustomDomainOk

`func (o *FormattedApiInstance) GetCustomDomainOk() (*bool, bool)`

GetCustomDomainOk returns a tuple with the CustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomain

`func (o *FormattedApiInstance) SetCustomDomain(v bool)`

SetCustomDomain sets CustomDomain field to given value.


### GetDailyAdminCnt

`func (o *FormattedApiInstance) GetDailyAdminCnt() float32`

GetDailyAdminCnt returns the DailyAdminCnt field if non-nil, zero value otherwise.

### GetDailyAdminCntOk

`func (o *FormattedApiInstance) GetDailyAdminCntOk() (*float32, bool)`

GetDailyAdminCntOk returns a tuple with the DailyAdminCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyAdminCnt

`func (o *FormattedApiInstance) SetDailyAdminCnt(v float32)`

SetDailyAdminCnt sets DailyAdminCnt field to given value.


### GetDailyEditorCnt

`func (o *FormattedApiInstance) GetDailyEditorCnt() float32`

GetDailyEditorCnt returns the DailyEditorCnt field if non-nil, zero value otherwise.

### GetDailyEditorCntOk

`func (o *FormattedApiInstance) GetDailyEditorCntOk() (*float32, bool)`

GetDailyEditorCntOk returns a tuple with the DailyEditorCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyEditorCnt

`func (o *FormattedApiInstance) SetDailyEditorCnt(v float32)`

SetDailyEditorCnt sets DailyEditorCnt field to given value.


### GetDailyUserCnt

`func (o *FormattedApiInstance) GetDailyUserCnt() float32`

GetDailyUserCnt returns the DailyUserCnt field if non-nil, zero value otherwise.

### GetDailyUserCntOk

`func (o *FormattedApiInstance) GetDailyUserCntOk() (*float32, bool)`

GetDailyUserCntOk returns a tuple with the DailyUserCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyUserCnt

`func (o *FormattedApiInstance) SetDailyUserCnt(v float32)`

SetDailyUserCnt sets DailyUserCnt field to given value.


### GetDailyViewerCnt

`func (o *FormattedApiInstance) GetDailyViewerCnt() float32`

GetDailyViewerCnt returns the DailyViewerCnt field if non-nil, zero value otherwise.

### GetDailyViewerCntOk

`func (o *FormattedApiInstance) GetDailyViewerCntOk() (*float32, bool)`

GetDailyViewerCntOk returns a tuple with the DailyViewerCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyViewerCnt

`func (o *FormattedApiInstance) SetDailyViewerCnt(v float32)`

SetDailyViewerCnt sets DailyViewerCnt field to given value.


### GetDashboardCnt

`func (o *FormattedApiInstance) GetDashboardCnt() float32`

GetDashboardCnt returns the DashboardCnt field if non-nil, zero value otherwise.

### GetDashboardCntOk

`func (o *FormattedApiInstance) GetDashboardCntOk() (*float32, bool)`

GetDashboardCntOk returns a tuple with the DashboardCnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardCnt

`func (o *FormattedApiInstance) SetDashboardCnt(v float32)`

SetDashboardCnt sets DashboardCnt field to given value.


### GetDashboardQuota

`func (o *FormattedApiInstance) GetDashboardQuota() float32`

GetDashboardQuota returns the DashboardQuota field if non-nil, zero value otherwise.

### GetDashboardQuotaOk

`func (o *FormattedApiInstance) GetDashboardQuotaOk() (*float32, bool)`

GetDashboardQuotaOk returns a tuple with the DashboardQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardQuota

`func (o *FormattedApiInstance) SetDashboardQuota(v float32)`

SetDashboardQuota sets DashboardQuota field to given value.


### GetDatasourceCnts

`func (o *FormattedApiInstance) GetDatasourceCnts() map[string]interface{}`

GetDatasourceCnts returns the DatasourceCnts field if non-nil, zero value otherwise.

### GetDatasourceCntsOk

`func (o *FormattedApiInstance) GetDatasourceCntsOk() (*map[string]interface{}, bool)`

GetDatasourceCntsOk returns a tuple with the DatasourceCnts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceCnts

`func (o *FormattedApiInstance) SetDatasourceCnts(v map[string]interface{})`

SetDatasourceCnts sets DatasourceCnts field to given value.


### GetDescription

`func (o *FormattedApiInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FormattedApiInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FormattedApiInstance) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGateway

`func (o *FormattedApiInstance) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *FormattedApiInstance) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *FormattedApiInstance) SetGateway(v string)`

SetGateway sets Gateway field to given value.


### GetHlInstanceCurrentUsage

`func (o *FormattedApiInstance) GetHlInstanceCurrentUsage() float32`

GetHlInstanceCurrentUsage returns the HlInstanceCurrentUsage field if non-nil, zero value otherwise.

### GetHlInstanceCurrentUsageOk

`func (o *FormattedApiInstance) GetHlInstanceCurrentUsageOk() (*float32, bool)`

GetHlInstanceCurrentUsageOk returns a tuple with the HlInstanceCurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlInstanceCurrentUsage

`func (o *FormattedApiInstance) SetHlInstanceCurrentUsage(v float32)`

SetHlInstanceCurrentUsage sets HlInstanceCurrentUsage field to given value.


### GetHlInstanceBillingUsage

`func (o *FormattedApiInstance) GetHlInstanceBillingUsage() float32`

GetHlInstanceBillingUsage returns the HlInstanceBillingUsage field if non-nil, zero value otherwise.

### GetHlInstanceBillingUsageOk

`func (o *FormattedApiInstance) GetHlInstanceBillingUsageOk() (*float32, bool)`

GetHlInstanceBillingUsageOk returns a tuple with the HlInstanceBillingUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlInstanceBillingUsage

`func (o *FormattedApiInstance) SetHlInstanceBillingUsage(v float32)`

SetHlInstanceBillingUsage sets HlInstanceBillingUsage field to given value.


### GetHlInstanceId

`func (o *FormattedApiInstance) GetHlInstanceId() float32`

GetHlInstanceId returns the HlInstanceId field if non-nil, zero value otherwise.

### GetHlInstanceIdOk

`func (o *FormattedApiInstance) GetHlInstanceIdOk() (*float32, bool)`

GetHlInstanceIdOk returns a tuple with the HlInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlInstanceId

`func (o *FormattedApiInstance) SetHlInstanceId(v float32)`

SetHlInstanceId sets HlInstanceId field to given value.


### GetHlInstanceName

`func (o *FormattedApiInstance) GetHlInstanceName() string`

GetHlInstanceName returns the HlInstanceName field if non-nil, zero value otherwise.

### GetHlInstanceNameOk

`func (o *FormattedApiInstance) GetHlInstanceNameOk() (*string, bool)`

GetHlInstanceNameOk returns a tuple with the HlInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlInstanceName

`func (o *FormattedApiInstance) SetHlInstanceName(v string)`

SetHlInstanceName sets HlInstanceName field to given value.


### GetHlInstanceStatus

`func (o *FormattedApiInstance) GetHlInstanceStatus() string`

GetHlInstanceStatus returns the HlInstanceStatus field if non-nil, zero value otherwise.

### GetHlInstanceStatusOk

`func (o *FormattedApiInstance) GetHlInstanceStatusOk() (*string, bool)`

GetHlInstanceStatusOk returns a tuple with the HlInstanceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlInstanceStatus

`func (o *FormattedApiInstance) SetHlInstanceStatus(v string)`

SetHlInstanceStatus sets HlInstanceStatus field to given value.


### GetHlInstanceUrl

`func (o *FormattedApiInstance) GetHlInstanceUrl() string`

GetHlInstanceUrl returns the HlInstanceUrl field if non-nil, zero value otherwise.

### GetHlInstanceUrlOk

`func (o *FormattedApiInstance) GetHlInstanceUrlOk() (*string, bool)`

GetHlInstanceUrlOk returns a tuple with the HlInstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlInstanceUrl

`func (o *FormattedApiInstance) SetHlInstanceUrl(v string)`

SetHlInstanceUrl sets HlInstanceUrl field to given value.


### GetHlInstanceClusterId

`func (o *FormattedApiInstance) GetHlInstanceClusterId() float32`

GetHlInstanceClusterId returns the HlInstanceClusterId field if non-nil, zero value otherwise.

### GetHlInstanceClusterIdOk

`func (o *FormattedApiInstance) GetHlInstanceClusterIdOk() (*float32, bool)`

GetHlInstanceClusterIdOk returns a tuple with the HlInstanceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHlInstanceClusterId

`func (o *FormattedApiInstance) SetHlInstanceClusterId(v float32)`

SetHlInstanceClusterId sets HlInstanceClusterId field to given value.


### GetHmInstanceGraphiteCurrentUsage

`func (o *FormattedApiInstance) GetHmInstanceGraphiteCurrentUsage() float32`

GetHmInstanceGraphiteCurrentUsage returns the HmInstanceGraphiteCurrentUsage field if non-nil, zero value otherwise.

### GetHmInstanceGraphiteCurrentUsageOk

`func (o *FormattedApiInstance) GetHmInstanceGraphiteCurrentUsageOk() (*float32, bool)`

GetHmInstanceGraphiteCurrentUsageOk returns a tuple with the HmInstanceGraphiteCurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstanceGraphiteCurrentUsage

`func (o *FormattedApiInstance) SetHmInstanceGraphiteCurrentUsage(v float32)`

SetHmInstanceGraphiteCurrentUsage sets HmInstanceGraphiteCurrentUsage field to given value.


### GetHmInstanceGraphiteBillingUsage

`func (o *FormattedApiInstance) GetHmInstanceGraphiteBillingUsage() float32`

GetHmInstanceGraphiteBillingUsage returns the HmInstanceGraphiteBillingUsage field if non-nil, zero value otherwise.

### GetHmInstanceGraphiteBillingUsageOk

`func (o *FormattedApiInstance) GetHmInstanceGraphiteBillingUsageOk() (*float32, bool)`

GetHmInstanceGraphiteBillingUsageOk returns a tuple with the HmInstanceGraphiteBillingUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstanceGraphiteBillingUsage

`func (o *FormattedApiInstance) SetHmInstanceGraphiteBillingUsage(v float32)`

SetHmInstanceGraphiteBillingUsage sets HmInstanceGraphiteBillingUsage field to given value.


### GetHmInstanceGraphiteId

`func (o *FormattedApiInstance) GetHmInstanceGraphiteId() float32`

GetHmInstanceGraphiteId returns the HmInstanceGraphiteId field if non-nil, zero value otherwise.

### GetHmInstanceGraphiteIdOk

`func (o *FormattedApiInstance) GetHmInstanceGraphiteIdOk() (*float32, bool)`

GetHmInstanceGraphiteIdOk returns a tuple with the HmInstanceGraphiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstanceGraphiteId

`func (o *FormattedApiInstance) SetHmInstanceGraphiteId(v float32)`

SetHmInstanceGraphiteId sets HmInstanceGraphiteId field to given value.


### GetHmInstanceGraphiteName

`func (o *FormattedApiInstance) GetHmInstanceGraphiteName() string`

GetHmInstanceGraphiteName returns the HmInstanceGraphiteName field if non-nil, zero value otherwise.

### GetHmInstanceGraphiteNameOk

`func (o *FormattedApiInstance) GetHmInstanceGraphiteNameOk() (*string, bool)`

GetHmInstanceGraphiteNameOk returns a tuple with the HmInstanceGraphiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstanceGraphiteName

`func (o *FormattedApiInstance) SetHmInstanceGraphiteName(v string)`

SetHmInstanceGraphiteName sets HmInstanceGraphiteName field to given value.


### GetHmInstanceGraphiteStatus

`func (o *FormattedApiInstance) GetHmInstanceGraphiteStatus() string`

GetHmInstanceGraphiteStatus returns the HmInstanceGraphiteStatus field if non-nil, zero value otherwise.

### GetHmInstanceGraphiteStatusOk

`func (o *FormattedApiInstance) GetHmInstanceGraphiteStatusOk() (*string, bool)`

GetHmInstanceGraphiteStatusOk returns a tuple with the HmInstanceGraphiteStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstanceGraphiteStatus

`func (o *FormattedApiInstance) SetHmInstanceGraphiteStatus(v string)`

SetHmInstanceGraphiteStatus sets HmInstanceGraphiteStatus field to given value.


### GetHmInstanceGraphiteType

`func (o *FormattedApiInstance) GetHmInstanceGraphiteType() string`

GetHmInstanceGraphiteType returns the HmInstanceGraphiteType field if non-nil, zero value otherwise.

### GetHmInstanceGraphiteTypeOk

`func (o *FormattedApiInstance) GetHmInstanceGraphiteTypeOk() (*string, bool)`

GetHmInstanceGraphiteTypeOk returns a tuple with the HmInstanceGraphiteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstanceGraphiteType

`func (o *FormattedApiInstance) SetHmInstanceGraphiteType(v string)`

SetHmInstanceGraphiteType sets HmInstanceGraphiteType field to given value.


### GetHmInstanceGraphiteUrl

`func (o *FormattedApiInstance) GetHmInstanceGraphiteUrl() string`

GetHmInstanceGraphiteUrl returns the HmInstanceGraphiteUrl field if non-nil, zero value otherwise.

### GetHmInstanceGraphiteUrlOk

`func (o *FormattedApiInstance) GetHmInstanceGraphiteUrlOk() (*string, bool)`

GetHmInstanceGraphiteUrlOk returns a tuple with the HmInstanceGraphiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstanceGraphiteUrl

`func (o *FormattedApiInstance) SetHmInstanceGraphiteUrl(v string)`

SetHmInstanceGraphiteUrl sets HmInstanceGraphiteUrl field to given value.


### GetHmInstancePromClusterId

`func (o *FormattedApiInstance) GetHmInstancePromClusterId() float32`

GetHmInstancePromClusterId returns the HmInstancePromClusterId field if non-nil, zero value otherwise.

### GetHmInstancePromClusterIdOk

`func (o *FormattedApiInstance) GetHmInstancePromClusterIdOk() (*float32, bool)`

GetHmInstancePromClusterIdOk returns a tuple with the HmInstancePromClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstancePromClusterId

`func (o *FormattedApiInstance) SetHmInstancePromClusterId(v float32)`

SetHmInstancePromClusterId sets HmInstancePromClusterId field to given value.


### GetHmInstancePromCurrentActiveSeries

`func (o *FormattedApiInstance) GetHmInstancePromCurrentActiveSeries() float32`

GetHmInstancePromCurrentActiveSeries returns the HmInstancePromCurrentActiveSeries field if non-nil, zero value otherwise.

### GetHmInstancePromCurrentActiveSeriesOk

`func (o *FormattedApiInstance) GetHmInstancePromCurrentActiveSeriesOk() (*float32, bool)`

GetHmInstancePromCurrentActiveSeriesOk returns a tuple with the HmInstancePromCurrentActiveSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstancePromCurrentActiveSeries

`func (o *FormattedApiInstance) SetHmInstancePromCurrentActiveSeries(v float32)`

SetHmInstancePromCurrentActiveSeries sets HmInstancePromCurrentActiveSeries field to given value.


### SetHmInstancePromCurrentActiveSeriesNil

`func (o *FormattedApiInstance) SetHmInstancePromCurrentActiveSeriesNil(b bool)`

 SetHmInstancePromCurrentActiveSeriesNil sets the value for HmInstancePromCurrentActiveSeries to be an explicit nil

### UnsetHmInstancePromCurrentActiveSeries
`func (o *FormattedApiInstance) UnsetHmInstancePromCurrentActiveSeries()`

UnsetHmInstancePromCurrentActiveSeries ensures that no value is present for HmInstancePromCurrentActiveSeries, not even an explicit nil
### GetHmInstancePromCurrentUsage

`func (o *FormattedApiInstance) GetHmInstancePromCurrentUsage() float32`

GetHmInstancePromCurrentUsage returns the HmInstancePromCurrentUsage field if non-nil, zero value otherwise.

### GetHmInstancePromCurrentUsageOk

`func (o *FormattedApiInstance) GetHmInstancePromCurrentUsageOk() (*float32, bool)`

GetHmInstancePromCurrentUsageOk returns a tuple with the HmInstancePromCurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstancePromCurrentUsage

`func (o *FormattedApiInstance) SetHmInstancePromCurrentUsage(v float32)`

SetHmInstancePromCurrentUsage sets HmInstancePromCurrentUsage field to given value.


### GetHmInstancePromBillingUsage

`func (o *FormattedApiInstance) GetHmInstancePromBillingUsage() float32`

GetHmInstancePromBillingUsage returns the HmInstancePromBillingUsage field if non-nil, zero value otherwise.

### GetHmInstancePromBillingUsageOk

`func (o *FormattedApiInstance) GetHmInstancePromBillingUsageOk() (*float32, bool)`

GetHmInstancePromBillingUsageOk returns a tuple with the HmInstancePromBillingUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstancePromBillingUsage

`func (o *FormattedApiInstance) SetHmInstancePromBillingUsage(v float32)`

SetHmInstancePromBillingUsage sets HmInstancePromBillingUsage field to given value.


### GetHmInstancePromId

`func (o *FormattedApiInstance) GetHmInstancePromId() float32`

GetHmInstancePromId returns the HmInstancePromId field if non-nil, zero value otherwise.

### GetHmInstancePromIdOk

`func (o *FormattedApiInstance) GetHmInstancePromIdOk() (*float32, bool)`

GetHmInstancePromIdOk returns a tuple with the HmInstancePromId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstancePromId

`func (o *FormattedApiInstance) SetHmInstancePromId(v float32)`

SetHmInstancePromId sets HmInstancePromId field to given value.


### GetHmInstancePromName

`func (o *FormattedApiInstance) GetHmInstancePromName() string`

GetHmInstancePromName returns the HmInstancePromName field if non-nil, zero value otherwise.

### GetHmInstancePromNameOk

`func (o *FormattedApiInstance) GetHmInstancePromNameOk() (*string, bool)`

GetHmInstancePromNameOk returns a tuple with the HmInstancePromName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstancePromName

`func (o *FormattedApiInstance) SetHmInstancePromName(v string)`

SetHmInstancePromName sets HmInstancePromName field to given value.


### GetHmInstancePromStatus

`func (o *FormattedApiInstance) GetHmInstancePromStatus() string`

GetHmInstancePromStatus returns the HmInstancePromStatus field if non-nil, zero value otherwise.

### GetHmInstancePromStatusOk

`func (o *FormattedApiInstance) GetHmInstancePromStatusOk() (*string, bool)`

GetHmInstancePromStatusOk returns a tuple with the HmInstancePromStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstancePromStatus

`func (o *FormattedApiInstance) SetHmInstancePromStatus(v string)`

SetHmInstancePromStatus sets HmInstancePromStatus field to given value.


### GetHmInstancePromUrl

`func (o *FormattedApiInstance) GetHmInstancePromUrl() string`

GetHmInstancePromUrl returns the HmInstancePromUrl field if non-nil, zero value otherwise.

### GetHmInstancePromUrlOk

`func (o *FormattedApiInstance) GetHmInstancePromUrlOk() (*string, bool)`

GetHmInstancePromUrlOk returns a tuple with the HmInstancePromUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstancePromUrl

`func (o *FormattedApiInstance) SetHmInstancePromUrl(v string)`

SetHmInstancePromUrl sets HmInstancePromUrl field to given value.


### GetHmInstanceGraphiteClusterId

`func (o *FormattedApiInstance) GetHmInstanceGraphiteClusterId() float32`

GetHmInstanceGraphiteClusterId returns the HmInstanceGraphiteClusterId field if non-nil, zero value otherwise.

### GetHmInstanceGraphiteClusterIdOk

`func (o *FormattedApiInstance) GetHmInstanceGraphiteClusterIdOk() (*float32, bool)`

GetHmInstanceGraphiteClusterIdOk returns a tuple with the HmInstanceGraphiteClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHmInstanceGraphiteClusterId

`func (o *FormattedApiInstance) SetHmInstanceGraphiteClusterId(v float32)`

SetHmInstanceGraphiteClusterId sets HmInstanceGraphiteClusterId field to given value.


### GetHtInstanceId

`func (o *FormattedApiInstance) GetHtInstanceId() float32`

GetHtInstanceId returns the HtInstanceId field if non-nil, zero value otherwise.

### GetHtInstanceIdOk

`func (o *FormattedApiInstance) GetHtInstanceIdOk() (*float32, bool)`

GetHtInstanceIdOk returns a tuple with the HtInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtInstanceId

`func (o *FormattedApiInstance) SetHtInstanceId(v float32)`

SetHtInstanceId sets HtInstanceId field to given value.


### GetHtInstanceName

`func (o *FormattedApiInstance) GetHtInstanceName() string`

GetHtInstanceName returns the HtInstanceName field if non-nil, zero value otherwise.

### GetHtInstanceNameOk

`func (o *FormattedApiInstance) GetHtInstanceNameOk() (*string, bool)`

GetHtInstanceNameOk returns a tuple with the HtInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtInstanceName

`func (o *FormattedApiInstance) SetHtInstanceName(v string)`

SetHtInstanceName sets HtInstanceName field to given value.


### GetHtInstanceStatus

`func (o *FormattedApiInstance) GetHtInstanceStatus() string`

GetHtInstanceStatus returns the HtInstanceStatus field if non-nil, zero value otherwise.

### GetHtInstanceStatusOk

`func (o *FormattedApiInstance) GetHtInstanceStatusOk() (*string, bool)`

GetHtInstanceStatusOk returns a tuple with the HtInstanceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtInstanceStatus

`func (o *FormattedApiInstance) SetHtInstanceStatus(v string)`

SetHtInstanceStatus sets HtInstanceStatus field to given value.


### GetHtInstanceUrl

`func (o *FormattedApiInstance) GetHtInstanceUrl() string`

GetHtInstanceUrl returns the HtInstanceUrl field if non-nil, zero value otherwise.

### GetHtInstanceUrlOk

`func (o *FormattedApiInstance) GetHtInstanceUrlOk() (*string, bool)`

GetHtInstanceUrlOk returns a tuple with the HtInstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtInstanceUrl

`func (o *FormattedApiInstance) SetHtInstanceUrl(v string)`

SetHtInstanceUrl sets HtInstanceUrl field to given value.


### GetHtInstanceClusterId

`func (o *FormattedApiInstance) GetHtInstanceClusterId() float32`

GetHtInstanceClusterId returns the HtInstanceClusterId field if non-nil, zero value otherwise.

### GetHtInstanceClusterIdOk

`func (o *FormattedApiInstance) GetHtInstanceClusterIdOk() (*float32, bool)`

GetHtInstanceClusterIdOk returns a tuple with the HtInstanceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtInstanceClusterId

`func (o *FormattedApiInstance) SetHtInstanceClusterId(v float32)`

SetHtInstanceClusterId sets HtInstanceClusterId field to given value.


### GetHtInstanceCurrentUsage

`func (o *FormattedApiInstance) GetHtInstanceCurrentUsage() float32`

GetHtInstanceCurrentUsage returns the HtInstanceCurrentUsage field if non-nil, zero value otherwise.

### GetHtInstanceCurrentUsageOk

`func (o *FormattedApiInstance) GetHtInstanceCurrentUsageOk() (*float32, bool)`

GetHtInstanceCurrentUsageOk returns a tuple with the HtInstanceCurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtInstanceCurrentUsage

`func (o *FormattedApiInstance) SetHtInstanceCurrentUsage(v float32)`

SetHtInstanceCurrentUsage sets HtInstanceCurrentUsage field to given value.


### GetHtInstanceBillingUsage

`func (o *FormattedApiInstance) GetHtInstanceBillingUsage() float32`

GetHtInstanceBillingUsage returns the HtInstanceBillingUsage field if non-nil, zero value otherwise.

### GetHtInstanceBillingUsageOk

`func (o *FormattedApiInstance) GetHtInstanceBillingUsageOk() (*float32, bool)`

GetHtInstanceBillingUsageOk returns a tuple with the HtInstanceBillingUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtInstanceBillingUsage

`func (o *FormattedApiInstance) SetHtInstanceBillingUsage(v float32)`

SetHtInstanceBillingUsage sets HtInstanceBillingUsage field to given value.


### GetHpInstanceId

`func (o *FormattedApiInstance) GetHpInstanceId() float32`

GetHpInstanceId returns the HpInstanceId field if non-nil, zero value otherwise.

### GetHpInstanceIdOk

`func (o *FormattedApiInstance) GetHpInstanceIdOk() (*float32, bool)`

GetHpInstanceIdOk returns a tuple with the HpInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpInstanceId

`func (o *FormattedApiInstance) SetHpInstanceId(v float32)`

SetHpInstanceId sets HpInstanceId field to given value.


### GetHpInstanceName

`func (o *FormattedApiInstance) GetHpInstanceName() string`

GetHpInstanceName returns the HpInstanceName field if non-nil, zero value otherwise.

### GetHpInstanceNameOk

`func (o *FormattedApiInstance) GetHpInstanceNameOk() (*string, bool)`

GetHpInstanceNameOk returns a tuple with the HpInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpInstanceName

`func (o *FormattedApiInstance) SetHpInstanceName(v string)`

SetHpInstanceName sets HpInstanceName field to given value.


### GetHpInstanceStatus

`func (o *FormattedApiInstance) GetHpInstanceStatus() string`

GetHpInstanceStatus returns the HpInstanceStatus field if non-nil, zero value otherwise.

### GetHpInstanceStatusOk

`func (o *FormattedApiInstance) GetHpInstanceStatusOk() (*string, bool)`

GetHpInstanceStatusOk returns a tuple with the HpInstanceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpInstanceStatus

`func (o *FormattedApiInstance) SetHpInstanceStatus(v string)`

SetHpInstanceStatus sets HpInstanceStatus field to given value.


### GetHpInstanceUrl

`func (o *FormattedApiInstance) GetHpInstanceUrl() string`

GetHpInstanceUrl returns the HpInstanceUrl field if non-nil, zero value otherwise.

### GetHpInstanceUrlOk

`func (o *FormattedApiInstance) GetHpInstanceUrlOk() (*string, bool)`

GetHpInstanceUrlOk returns a tuple with the HpInstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpInstanceUrl

`func (o *FormattedApiInstance) SetHpInstanceUrl(v string)`

SetHpInstanceUrl sets HpInstanceUrl field to given value.


### GetHpInstanceClusterId

`func (o *FormattedApiInstance) GetHpInstanceClusterId() float32`

GetHpInstanceClusterId returns the HpInstanceClusterId field if non-nil, zero value otherwise.

### GetHpInstanceClusterIdOk

`func (o *FormattedApiInstance) GetHpInstanceClusterIdOk() (*float32, bool)`

GetHpInstanceClusterIdOk returns a tuple with the HpInstanceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpInstanceClusterId

`func (o *FormattedApiInstance) SetHpInstanceClusterId(v float32)`

SetHpInstanceClusterId sets HpInstanceClusterId field to given value.


### GetHpInstanceCurrentUsage

`func (o *FormattedApiInstance) GetHpInstanceCurrentUsage() float32`

GetHpInstanceCurrentUsage returns the HpInstanceCurrentUsage field if non-nil, zero value otherwise.

### GetHpInstanceCurrentUsageOk

`func (o *FormattedApiInstance) GetHpInstanceCurrentUsageOk() (*float32, bool)`

GetHpInstanceCurrentUsageOk returns a tuple with the HpInstanceCurrentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpInstanceCurrentUsage

`func (o *FormattedApiInstance) SetHpInstanceCurrentUsage(v float32)`

SetHpInstanceCurrentUsage sets HpInstanceCurrentUsage field to given value.


### GetHpInstanceBillingUsage

`func (o *FormattedApiInstance) GetHpInstanceBillingUsage() float32`

GetHpInstanceBillingUsage returns the HpInstanceBillingUsage field if non-nil, zero value otherwise.

### GetHpInstanceBillingUsageOk

`func (o *FormattedApiInstance) GetHpInstanceBillingUsageOk() (*float32, bool)`

GetHpInstanceBillingUsageOk returns a tuple with the HpInstanceBillingUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHpInstanceBillingUsage

`func (o *FormattedApiInstance) SetHpInstanceBillingUsage(v float32)`

SetHpInstanceBillingUsage sets HpInstanceBillingUsage field to given value.


### GetId

`func (o *FormattedApiInstance) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormattedApiInstance) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormattedApiInstance) SetId(v float32)`

SetId sets Id field to given value.


### GetIncident

`func (o *FormattedApiInstance) GetIncident() float32`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *FormattedApiInstance) GetIncidentOk() (*float32, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *FormattedApiInstance) SetIncident(v float32)`

SetIncident sets Incident field to given value.


### GetLabels

`func (o *FormattedApiInstance) GetLabels() map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *FormattedApiInstance) GetLabelsOk() (*map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *FormattedApiInstance) SetLabels(v map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *FormattedApiInstance) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMachineLearning

`func (o *FormattedApiInstance) GetMachineLearning() float32`

GetMachineLearning returns the MachineLearning field if non-nil, zero value otherwise.

### GetMachineLearningOk

`func (o *FormattedApiInstance) GetMachineLearningOk() (*float32, bool)`

GetMachineLearningOk returns a tuple with the MachineLearning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLearning

`func (o *FormattedApiInstance) SetMachineLearning(v float32)`

SetMachineLearning sets MachineLearning field to given value.


### GetName

`func (o *FormattedApiInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormattedApiInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormattedApiInstance) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *FormattedApiInstance) GetOrgId() float32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *FormattedApiInstance) GetOrgIdOk() (*float32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *FormattedApiInstance) SetOrgId(v float32)`

SetOrgId sets OrgId field to given value.


### GetOrgName

`func (o *FormattedApiInstance) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *FormattedApiInstance) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *FormattedApiInstance) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.


### GetOrgSlug

`func (o *FormattedApiInstance) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *FormattedApiInstance) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *FormattedApiInstance) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.


### GetPlan

`func (o *FormattedApiInstance) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *FormattedApiInstance) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *FormattedApiInstance) SetPlan(v string)`

SetPlan sets Plan field to given value.


### GetPlanName

`func (o *FormattedApiInstance) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *FormattedApiInstance) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *FormattedApiInstance) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.


### GetRegionId

`func (o *FormattedApiInstance) GetRegionId() float32`

GetRegionId returns the RegionId field if non-nil, zero value otherwise.

### GetRegionIdOk

`func (o *FormattedApiInstance) GetRegionIdOk() (*float32, bool)`

GetRegionIdOk returns a tuple with the RegionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionId

`func (o *FormattedApiInstance) SetRegionId(v float32)`

SetRegionId sets RegionId field to given value.


### GetRegionSlug

`func (o *FormattedApiInstance) GetRegionSlug() string`

GetRegionSlug returns the RegionSlug field if non-nil, zero value otherwise.

### GetRegionSlugOk

`func (o *FormattedApiInstance) GetRegionSlugOk() (*string, bool)`

GetRegionSlugOk returns a tuple with the RegionSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionSlug

`func (o *FormattedApiInstance) SetRegionSlug(v string)`

SetRegionSlug sets RegionSlug field to given value.


### GetRegionPublicName

`func (o *FormattedApiInstance) GetRegionPublicName() string`

GetRegionPublicName returns the RegionPublicName field if non-nil, zero value otherwise.

### GetRegionPublicNameOk

`func (o *FormattedApiInstance) GetRegionPublicNameOk() (*string, bool)`

GetRegionPublicNameOk returns a tuple with the RegionPublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionPublicName

`func (o *FormattedApiInstance) SetRegionPublicName(v string)`

SetRegionPublicName sets RegionPublicName field to given value.


### GetProvider

`func (o *FormattedApiInstance) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *FormattedApiInstance) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *FormattedApiInstance) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetProviderRegion

`func (o *FormattedApiInstance) GetProviderRegion() string`

GetProviderRegion returns the ProviderRegion field if non-nil, zero value otherwise.

### GetProviderRegionOk

`func (o *FormattedApiInstance) GetProviderRegionOk() (*string, bool)`

GetProviderRegionOk returns a tuple with the ProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderRegion

`func (o *FormattedApiInstance) SetProviderRegion(v string)`

SetProviderRegion sets ProviderRegion field to given value.


### GetRunningVersion

`func (o *FormattedApiInstance) GetRunningVersion() string`

GetRunningVersion returns the RunningVersion field if non-nil, zero value otherwise.

### GetRunningVersionOk

`func (o *FormattedApiInstance) GetRunningVersionOk() (*string, bool)`

GetRunningVersionOk returns a tuple with the RunningVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningVersion

`func (o *FormattedApiInstance) SetRunningVersion(v string)`

SetRunningVersion sets RunningVersion field to given value.


### GetSlug

`func (o *FormattedApiInstance) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *FormattedApiInstance) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *FormattedApiInstance) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetSsl

`func (o *FormattedApiInstance) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *FormattedApiInstance) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *FormattedApiInstance) SetSsl(v bool)`

SetSsl sets Ssl field to given value.


### GetStatus

`func (o *FormattedApiInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FormattedApiInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FormattedApiInstance) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSupport

`func (o *FormattedApiInstance) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *FormattedApiInstance) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *FormattedApiInstance) SetSupport(v bool)`

SetSupport sets Support field to given value.


### GetTrial

`func (o *FormattedApiInstance) GetTrial() float32`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *FormattedApiInstance) GetTrialOk() (*float32, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *FormattedApiInstance) SetTrial(v float32)`

SetTrial sets Trial field to given value.


### GetTrialExpiresAt

`func (o *FormattedApiInstance) GetTrialExpiresAt() string`

GetTrialExpiresAt returns the TrialExpiresAt field if non-nil, zero value otherwise.

### GetTrialExpiresAtOk

`func (o *FormattedApiInstance) GetTrialExpiresAtOk() (*string, bool)`

GetTrialExpiresAtOk returns a tuple with the TrialExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialExpiresAt

`func (o *FormattedApiInstance) SetTrialExpiresAt(v string)`

SetTrialExpiresAt sets TrialExpiresAt field to given value.


### SetTrialExpiresAtNil

`func (o *FormattedApiInstance) SetTrialExpiresAtNil(b bool)`

 SetTrialExpiresAtNil sets the value for TrialExpiresAt to be an explicit nil

### UnsetTrialExpiresAt
`func (o *FormattedApiInstance) UnsetTrialExpiresAt()`

UnsetTrialExpiresAt ensures that no value is present for TrialExpiresAt, not even an explicit nil
### GetUpdatedAt

`func (o *FormattedApiInstance) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormattedApiInstance) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormattedApiInstance) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *FormattedApiInstance) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *FormattedApiInstance) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUpdatedBy

`func (o *FormattedApiInstance) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FormattedApiInstance) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FormattedApiInstance) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### SetUpdatedByNil

`func (o *FormattedApiInstance) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *FormattedApiInstance) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil
### GetUrl

`func (o *FormattedApiInstance) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FormattedApiInstance) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FormattedApiInstance) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUserQuota

`func (o *FormattedApiInstance) GetUserQuota() float32`

GetUserQuota returns the UserQuota field if non-nil, zero value otherwise.

### GetUserQuotaOk

`func (o *FormattedApiInstance) GetUserQuotaOk() (*float32, bool)`

GetUserQuotaOk returns a tuple with the UserQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserQuota

`func (o *FormattedApiInstance) SetUserQuota(v float32)`

SetUserQuota sets UserQuota field to given value.


### GetVersion

`func (o *FormattedApiInstance) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FormattedApiInstance) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FormattedApiInstance) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVersionIssueLink

`func (o *FormattedApiInstance) GetVersionIssueLink() string`

GetVersionIssueLink returns the VersionIssueLink field if non-nil, zero value otherwise.

### GetVersionIssueLinkOk

`func (o *FormattedApiInstance) GetVersionIssueLinkOk() (*string, bool)`

GetVersionIssueLinkOk returns a tuple with the VersionIssueLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionIssueLink

`func (o *FormattedApiInstance) SetVersionIssueLink(v string)`

SetVersionIssueLink sets VersionIssueLink field to given value.

### HasVersionIssueLink

`func (o *FormattedApiInstance) HasVersionIssueLink() bool`

HasVersionIssueLink returns a boolean if a field has been set.

### GetAgentManagementInstanceId

`func (o *FormattedApiInstance) GetAgentManagementInstanceId() float32`

GetAgentManagementInstanceId returns the AgentManagementInstanceId field if non-nil, zero value otherwise.

### GetAgentManagementInstanceIdOk

`func (o *FormattedApiInstance) GetAgentManagementInstanceIdOk() (*float32, bool)`

GetAgentManagementInstanceIdOk returns a tuple with the AgentManagementInstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentManagementInstanceId

`func (o *FormattedApiInstance) SetAgentManagementInstanceId(v float32)`

SetAgentManagementInstanceId sets AgentManagementInstanceId field to given value.


### GetAgentManagementInstanceUrl

`func (o *FormattedApiInstance) GetAgentManagementInstanceUrl() string`

GetAgentManagementInstanceUrl returns the AgentManagementInstanceUrl field if non-nil, zero value otherwise.

### GetAgentManagementInstanceUrlOk

`func (o *FormattedApiInstance) GetAgentManagementInstanceUrlOk() (*string, bool)`

GetAgentManagementInstanceUrlOk returns a tuple with the AgentManagementInstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentManagementInstanceUrl

`func (o *FormattedApiInstance) SetAgentManagementInstanceUrl(v string)`

SetAgentManagementInstanceUrl sets AgentManagementInstanceUrl field to given value.


### GetAgentManagementInstanceName

`func (o *FormattedApiInstance) GetAgentManagementInstanceName() string`

GetAgentManagementInstanceName returns the AgentManagementInstanceName field if non-nil, zero value otherwise.

### GetAgentManagementInstanceNameOk

`func (o *FormattedApiInstance) GetAgentManagementInstanceNameOk() (*string, bool)`

GetAgentManagementInstanceNameOk returns a tuple with the AgentManagementInstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentManagementInstanceName

`func (o *FormattedApiInstance) SetAgentManagementInstanceName(v string)`

SetAgentManagementInstanceName sets AgentManagementInstanceName field to given value.


### GetAgentManagementInstanceStatus

`func (o *FormattedApiInstance) GetAgentManagementInstanceStatus() string`

GetAgentManagementInstanceStatus returns the AgentManagementInstanceStatus field if non-nil, zero value otherwise.

### GetAgentManagementInstanceStatusOk

`func (o *FormattedApiInstance) GetAgentManagementInstanceStatusOk() (*string, bool)`

GetAgentManagementInstanceStatusOk returns a tuple with the AgentManagementInstanceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentManagementInstanceStatus

`func (o *FormattedApiInstance) SetAgentManagementInstanceStatus(v string)`

SetAgentManagementInstanceStatus sets AgentManagementInstanceStatus field to given value.


### GetAgentManagementInstanceClusterId

`func (o *FormattedApiInstance) GetAgentManagementInstanceClusterId() float32`

GetAgentManagementInstanceClusterId returns the AgentManagementInstanceClusterId field if non-nil, zero value otherwise.

### GetAgentManagementInstanceClusterIdOk

`func (o *FormattedApiInstance) GetAgentManagementInstanceClusterIdOk() (*float32, bool)`

GetAgentManagementInstanceClusterIdOk returns a tuple with the AgentManagementInstanceClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentManagementInstanceClusterId

`func (o *FormattedApiInstance) SetAgentManagementInstanceClusterId(v float32)`

SetAgentManagementInstanceClusterId sets AgentManagementInstanceClusterId field to given value.


### GetConfig

`func (o *FormattedApiInstance) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *FormattedApiInstance) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *FormattedApiInstance) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *FormattedApiInstance) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetType

`func (o *FormattedApiInstance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormattedApiInstance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormattedApiInstance) SetType(v string)`

SetType sets Type field to given value.


### GetK6OrgId

`func (o *FormattedApiInstance) GetK6OrgId() float32`

GetK6OrgId returns the K6OrgId field if non-nil, zero value otherwise.

### GetK6OrgIdOk

`func (o *FormattedApiInstance) GetK6OrgIdOk() (*float32, bool)`

GetK6OrgIdOk returns a tuple with the K6OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK6OrgId

`func (o *FormattedApiInstance) SetK6OrgId(v float32)`

SetK6OrgId sets K6OrgId field to given value.


### SetK6OrgIdNil

`func (o *FormattedApiInstance) SetK6OrgIdNil(b bool)`

 SetK6OrgIdNil sets the value for K6OrgId to be an explicit nil

### UnsetK6OrgId
`func (o *FormattedApiInstance) UnsetK6OrgId()`

UnsetK6OrgId ensures that no value is present for K6OrgId, not even an explicit nil
### GetMachineLearningLogsToken

`func (o *FormattedApiInstance) GetMachineLearningLogsToken() string`

GetMachineLearningLogsToken returns the MachineLearningLogsToken field if non-nil, zero value otherwise.

### GetMachineLearningLogsTokenOk

`func (o *FormattedApiInstance) GetMachineLearningLogsTokenOk() (*string, bool)`

GetMachineLearningLogsTokenOk returns a tuple with the MachineLearningLogsToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineLearningLogsToken

`func (o *FormattedApiInstance) SetMachineLearningLogsToken(v string)`

SetMachineLearningLogsToken sets MachineLearningLogsToken field to given value.


### GetUsageStatsId

`func (o *FormattedApiInstance) GetUsageStatsId() string`

GetUsageStatsId returns the UsageStatsId field if non-nil, zero value otherwise.

### GetUsageStatsIdOk

`func (o *FormattedApiInstance) GetUsageStatsIdOk() (*string, bool)`

GetUsageStatsIdOk returns a tuple with the UsageStatsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStatsId

`func (o *FormattedApiInstance) SetUsageStatsId(v string)`

SetUsageStatsId sets UsageStatsId field to given value.


### GetRegionStackStateServiceUrl

`func (o *FormattedApiInstance) GetRegionStackStateServiceUrl() string`

GetRegionStackStateServiceUrl returns the RegionStackStateServiceUrl field if non-nil, zero value otherwise.

### GetRegionStackStateServiceUrlOk

`func (o *FormattedApiInstance) GetRegionStackStateServiceUrlOk() (*string, bool)`

GetRegionStackStateServiceUrlOk returns a tuple with the RegionStackStateServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionStackStateServiceUrl

`func (o *FormattedApiInstance) SetRegionStackStateServiceUrl(v string)`

SetRegionStackStateServiceUrl sets RegionStackStateServiceUrl field to given value.


### GetRegionSyntheticMonitoringApiUrl

`func (o *FormattedApiInstance) GetRegionSyntheticMonitoringApiUrl() string`

GetRegionSyntheticMonitoringApiUrl returns the RegionSyntheticMonitoringApiUrl field if non-nil, zero value otherwise.

### GetRegionSyntheticMonitoringApiUrlOk

`func (o *FormattedApiInstance) GetRegionSyntheticMonitoringApiUrlOk() (*string, bool)`

GetRegionSyntheticMonitoringApiUrlOk returns a tuple with the RegionSyntheticMonitoringApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionSyntheticMonitoringApiUrl

`func (o *FormattedApiInstance) SetRegionSyntheticMonitoringApiUrl(v string)`

SetRegionSyntheticMonitoringApiUrl sets RegionSyntheticMonitoringApiUrl field to given value.


### GetRegionInsightsApiUrl

`func (o *FormattedApiInstance) GetRegionInsightsApiUrl() string`

GetRegionInsightsApiUrl returns the RegionInsightsApiUrl field if non-nil, zero value otherwise.

### GetRegionInsightsApiUrlOk

`func (o *FormattedApiInstance) GetRegionInsightsApiUrlOk() (*string, bool)`

GetRegionInsightsApiUrlOk returns a tuple with the RegionInsightsApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionInsightsApiUrl

`func (o *FormattedApiInstance) SetRegionInsightsApiUrl(v string)`

SetRegionInsightsApiUrl sets RegionInsightsApiUrl field to given value.


### GetRegionIntegrationsApiUrl

`func (o *FormattedApiInstance) GetRegionIntegrationsApiUrl() string`

GetRegionIntegrationsApiUrl returns the RegionIntegrationsApiUrl field if non-nil, zero value otherwise.

### GetRegionIntegrationsApiUrlOk

`func (o *FormattedApiInstance) GetRegionIntegrationsApiUrlOk() (*string, bool)`

GetRegionIntegrationsApiUrlOk returns a tuple with the RegionIntegrationsApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionIntegrationsApiUrl

`func (o *FormattedApiInstance) SetRegionIntegrationsApiUrl(v string)`

SetRegionIntegrationsApiUrl sets RegionIntegrationsApiUrl field to given value.


### GetRegionHostedExportersApiUrl

`func (o *FormattedApiInstance) GetRegionHostedExportersApiUrl() string`

GetRegionHostedExportersApiUrl returns the RegionHostedExportersApiUrl field if non-nil, zero value otherwise.

### GetRegionHostedExportersApiUrlOk

`func (o *FormattedApiInstance) GetRegionHostedExportersApiUrlOk() (*string, bool)`

GetRegionHostedExportersApiUrlOk returns a tuple with the RegionHostedExportersApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionHostedExportersApiUrl

`func (o *FormattedApiInstance) SetRegionHostedExportersApiUrl(v string)`

SetRegionHostedExportersApiUrl sets RegionHostedExportersApiUrl field to given value.


### GetRegionMachineLearningApiUrl

`func (o *FormattedApiInstance) GetRegionMachineLearningApiUrl() string`

GetRegionMachineLearningApiUrl returns the RegionMachineLearningApiUrl field if non-nil, zero value otherwise.

### GetRegionMachineLearningApiUrlOk

`func (o *FormattedApiInstance) GetRegionMachineLearningApiUrlOk() (*string, bool)`

GetRegionMachineLearningApiUrlOk returns a tuple with the RegionMachineLearningApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionMachineLearningApiUrl

`func (o *FormattedApiInstance) SetRegionMachineLearningApiUrl(v string)`

SetRegionMachineLearningApiUrl sets RegionMachineLearningApiUrl field to given value.


### GetRegionLLMGatewayUrl

`func (o *FormattedApiInstance) GetRegionLLMGatewayUrl() string`

GetRegionLLMGatewayUrl returns the RegionLLMGatewayUrl field if non-nil, zero value otherwise.

### GetRegionLLMGatewayUrlOk

`func (o *FormattedApiInstance) GetRegionLLMGatewayUrlOk() (*string, bool)`

GetRegionLLMGatewayUrlOk returns a tuple with the RegionLLMGatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionLLMGatewayUrl

`func (o *FormattedApiInstance) SetRegionLLMGatewayUrl(v string)`

SetRegionLLMGatewayUrl sets RegionLLMGatewayUrl field to given value.


### GetRegionAssistantUrl

`func (o *FormattedApiInstance) GetRegionAssistantUrl() string`

GetRegionAssistantUrl returns the RegionAssistantUrl field if non-nil, zero value otherwise.

### GetRegionAssistantUrlOk

`func (o *FormattedApiInstance) GetRegionAssistantUrlOk() (*string, bool)`

GetRegionAssistantUrlOk returns a tuple with the RegionAssistantUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionAssistantUrl

`func (o *FormattedApiInstance) SetRegionAssistantUrl(v string)`

SetRegionAssistantUrl sets RegionAssistantUrl field to given value.


### GetLinks

`func (o *FormattedApiInstance) GetLinks() []LinksInner1`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FormattedApiInstance) GetLinksOk() (*[]LinksInner1, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FormattedApiInstance) SetLinks(v []LinksInner1)`

SetLinks sets Links field to given value.


### GetDeleteProtection

`func (o *FormattedApiInstance) GetDeleteProtection() bool`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *FormattedApiInstance) GetDeleteProtectionOk() (*bool, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *FormattedApiInstance) SetDeleteProtection(v bool)`

SetDeleteProtection sets DeleteProtection field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


