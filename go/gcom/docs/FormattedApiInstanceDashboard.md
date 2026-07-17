# FormattedApiInstanceDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | when the dashboard was created | 
**DashboardJson** | **interface{}** | the dashboard definition | 
**DisableDelete** | **bool** | whether the dashboard is protected from deletion | 
**Folder** | **string** | dashboard folder | 
**GrafanaOrgId** | **int32** | grafana org id | 
**Id** | **int64** | dashboard id | 
**InstanceId** | **int32** | hosted grafana instance id | 
**Name** | **string** | dashboard name | 
**UpdatedAt** | **NullableTime** | when the dashboard was last updated | 

## Methods

### NewFormattedApiInstanceDashboard

`func NewFormattedApiInstanceDashboard(createdAt time.Time, dashboardJson interface{}, disableDelete bool, folder string, grafanaOrgId int32, id int64, instanceId int32, name string, updatedAt NullableTime, ) *FormattedApiInstanceDashboard`

NewFormattedApiInstanceDashboard instantiates a new FormattedApiInstanceDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedApiInstanceDashboardWithDefaults

`func NewFormattedApiInstanceDashboardWithDefaults() *FormattedApiInstanceDashboard`

NewFormattedApiInstanceDashboardWithDefaults instantiates a new FormattedApiInstanceDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FormattedApiInstanceDashboard) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormattedApiInstanceDashboard) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormattedApiInstanceDashboard) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDashboardJson

`func (o *FormattedApiInstanceDashboard) GetDashboardJson() interface{}`

GetDashboardJson returns the DashboardJson field if non-nil, zero value otherwise.

### GetDashboardJsonOk

`func (o *FormattedApiInstanceDashboard) GetDashboardJsonOk() (*interface{}, bool)`

GetDashboardJsonOk returns a tuple with the DashboardJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardJson

`func (o *FormattedApiInstanceDashboard) SetDashboardJson(v interface{})`

SetDashboardJson sets DashboardJson field to given value.


### SetDashboardJsonNil

`func (o *FormattedApiInstanceDashboard) SetDashboardJsonNil(b bool)`

 SetDashboardJsonNil sets the value for DashboardJson to be an explicit nil

### UnsetDashboardJson
`func (o *FormattedApiInstanceDashboard) UnsetDashboardJson()`

UnsetDashboardJson ensures that no value is present for DashboardJson, not even an explicit nil
### GetDisableDelete

`func (o *FormattedApiInstanceDashboard) GetDisableDelete() bool`

GetDisableDelete returns the DisableDelete field if non-nil, zero value otherwise.

### GetDisableDeleteOk

`func (o *FormattedApiInstanceDashboard) GetDisableDeleteOk() (*bool, bool)`

GetDisableDeleteOk returns a tuple with the DisableDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDelete

`func (o *FormattedApiInstanceDashboard) SetDisableDelete(v bool)`

SetDisableDelete sets DisableDelete field to given value.


### GetFolder

`func (o *FormattedApiInstanceDashboard) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *FormattedApiInstanceDashboard) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *FormattedApiInstanceDashboard) SetFolder(v string)`

SetFolder sets Folder field to given value.


### GetGrafanaOrgId

`func (o *FormattedApiInstanceDashboard) GetGrafanaOrgId() int32`

GetGrafanaOrgId returns the GrafanaOrgId field if non-nil, zero value otherwise.

### GetGrafanaOrgIdOk

`func (o *FormattedApiInstanceDashboard) GetGrafanaOrgIdOk() (*int32, bool)`

GetGrafanaOrgIdOk returns a tuple with the GrafanaOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaOrgId

`func (o *FormattedApiInstanceDashboard) SetGrafanaOrgId(v int32)`

SetGrafanaOrgId sets GrafanaOrgId field to given value.


### GetId

`func (o *FormattedApiInstanceDashboard) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormattedApiInstanceDashboard) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormattedApiInstanceDashboard) SetId(v int64)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *FormattedApiInstanceDashboard) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FormattedApiInstanceDashboard) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FormattedApiInstanceDashboard) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.


### GetName

`func (o *FormattedApiInstanceDashboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormattedApiInstanceDashboard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormattedApiInstanceDashboard) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *FormattedApiInstanceDashboard) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormattedApiInstanceDashboard) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormattedApiInstanceDashboard) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *FormattedApiInstanceDashboard) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *FormattedApiInstanceDashboard) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


