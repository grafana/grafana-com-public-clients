# InstanceDashboardCreateBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DashboardJson** | **interface{}** | the dashboard definition | 
**DisableDelete** | **bool** | whether the dashboard is protected from deletion | 
**Folder** | Pointer to **string** | folder to place the dashboard in | [optional] 
**GrafanaOrgId** | Pointer to **int64** | grafana org id to create the dashboard in | [optional] [default to 1]
**Name** | **string** | dashboard name | 

## Methods

### NewInstanceDashboardCreateBody

`func NewInstanceDashboardCreateBody(dashboardJson interface{}, disableDelete bool, name string, ) *InstanceDashboardCreateBody`

NewInstanceDashboardCreateBody instantiates a new InstanceDashboardCreateBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceDashboardCreateBodyWithDefaults

`func NewInstanceDashboardCreateBodyWithDefaults() *InstanceDashboardCreateBody`

NewInstanceDashboardCreateBodyWithDefaults instantiates a new InstanceDashboardCreateBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboardJson

`func (o *InstanceDashboardCreateBody) GetDashboardJson() interface{}`

GetDashboardJson returns the DashboardJson field if non-nil, zero value otherwise.

### GetDashboardJsonOk

`func (o *InstanceDashboardCreateBody) GetDashboardJsonOk() (*interface{}, bool)`

GetDashboardJsonOk returns a tuple with the DashboardJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardJson

`func (o *InstanceDashboardCreateBody) SetDashboardJson(v interface{})`

SetDashboardJson sets DashboardJson field to given value.


### SetDashboardJsonNil

`func (o *InstanceDashboardCreateBody) SetDashboardJsonNil(b bool)`

 SetDashboardJsonNil sets the value for DashboardJson to be an explicit nil

### UnsetDashboardJson
`func (o *InstanceDashboardCreateBody) UnsetDashboardJson()`

UnsetDashboardJson ensures that no value is present for DashboardJson, not even an explicit nil
### GetDisableDelete

`func (o *InstanceDashboardCreateBody) GetDisableDelete() bool`

GetDisableDelete returns the DisableDelete field if non-nil, zero value otherwise.

### GetDisableDeleteOk

`func (o *InstanceDashboardCreateBody) GetDisableDeleteOk() (*bool, bool)`

GetDisableDeleteOk returns a tuple with the DisableDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDelete

`func (o *InstanceDashboardCreateBody) SetDisableDelete(v bool)`

SetDisableDelete sets DisableDelete field to given value.


### GetFolder

`func (o *InstanceDashboardCreateBody) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *InstanceDashboardCreateBody) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *InstanceDashboardCreateBody) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *InstanceDashboardCreateBody) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetGrafanaOrgId

`func (o *InstanceDashboardCreateBody) GetGrafanaOrgId() int64`

GetGrafanaOrgId returns the GrafanaOrgId field if non-nil, zero value otherwise.

### GetGrafanaOrgIdOk

`func (o *InstanceDashboardCreateBody) GetGrafanaOrgIdOk() (*int64, bool)`

GetGrafanaOrgIdOk returns a tuple with the GrafanaOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaOrgId

`func (o *InstanceDashboardCreateBody) SetGrafanaOrgId(v int64)`

SetGrafanaOrgId sets GrafanaOrgId field to given value.

### HasGrafanaOrgId

`func (o *InstanceDashboardCreateBody) HasGrafanaOrgId() bool`

HasGrafanaOrgId returns a boolean if a field has been set.

### GetName

`func (o *InstanceDashboardCreateBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceDashboardCreateBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceDashboardCreateBody) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


