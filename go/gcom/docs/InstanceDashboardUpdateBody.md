# InstanceDashboardUpdateBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DashboardJson** | Pointer to **interface{}** | the dashboard definition | [optional] 
**DisableDelete** | Pointer to **bool** | whether the dashboard is protected from deletion | [optional] 
**Folder** | Pointer to **string** | folder to place the dashboard in | [optional] 

## Methods

### NewInstanceDashboardUpdateBody

`func NewInstanceDashboardUpdateBody() *InstanceDashboardUpdateBody`

NewInstanceDashboardUpdateBody instantiates a new InstanceDashboardUpdateBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceDashboardUpdateBodyWithDefaults

`func NewInstanceDashboardUpdateBodyWithDefaults() *InstanceDashboardUpdateBody`

NewInstanceDashboardUpdateBodyWithDefaults instantiates a new InstanceDashboardUpdateBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboardJson

`func (o *InstanceDashboardUpdateBody) GetDashboardJson() interface{}`

GetDashboardJson returns the DashboardJson field if non-nil, zero value otherwise.

### GetDashboardJsonOk

`func (o *InstanceDashboardUpdateBody) GetDashboardJsonOk() (*interface{}, bool)`

GetDashboardJsonOk returns a tuple with the DashboardJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardJson

`func (o *InstanceDashboardUpdateBody) SetDashboardJson(v interface{})`

SetDashboardJson sets DashboardJson field to given value.

### HasDashboardJson

`func (o *InstanceDashboardUpdateBody) HasDashboardJson() bool`

HasDashboardJson returns a boolean if a field has been set.

### SetDashboardJsonNil

`func (o *InstanceDashboardUpdateBody) SetDashboardJsonNil(b bool)`

 SetDashboardJsonNil sets the value for DashboardJson to be an explicit nil

### UnsetDashboardJson
`func (o *InstanceDashboardUpdateBody) UnsetDashboardJson()`

UnsetDashboardJson ensures that no value is present for DashboardJson, not even an explicit nil
### GetDisableDelete

`func (o *InstanceDashboardUpdateBody) GetDisableDelete() bool`

GetDisableDelete returns the DisableDelete field if non-nil, zero value otherwise.

### GetDisableDeleteOk

`func (o *InstanceDashboardUpdateBody) GetDisableDeleteOk() (*bool, bool)`

GetDisableDeleteOk returns a tuple with the DisableDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDelete

`func (o *InstanceDashboardUpdateBody) SetDisableDelete(v bool)`

SetDisableDelete sets DisableDelete field to given value.

### HasDisableDelete

`func (o *InstanceDashboardUpdateBody) HasDisableDelete() bool`

HasDisableDelete returns a boolean if a field has been set.

### GetFolder

`func (o *InstanceDashboardUpdateBody) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *InstanceDashboardUpdateBody) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *InstanceDashboardUpdateBody) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *InstanceDashboardUpdateBody) HasFolder() bool`

HasFolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


