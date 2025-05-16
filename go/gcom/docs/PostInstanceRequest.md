# PostInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to **bool** |  | [optional] 
**DeleteProtection** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Graphite** | Pointer to **bool** |  | [optional] 
**HlInstanceId** | Pointer to **int32** |  | [optional] 
**K6OrgId** | Pointer to **NullableInt32** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**Logs** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Plan** | Pointer to **string** |  | [optional] 
**Prometheus** | Pointer to **bool** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

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

### GetDeleteProtection

`func (o *PostInstanceRequest) GetDeleteProtection() bool`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *PostInstanceRequest) GetDeleteProtectionOk() (*bool, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *PostInstanceRequest) SetDeleteProtection(v bool)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *PostInstanceRequest) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

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

### SetK6OrgIdNil

`func (o *PostInstanceRequest) SetK6OrgIdNil(b bool)`

 SetK6OrgIdNil sets the value for K6OrgId to be an explicit nil

### UnsetK6OrgId
`func (o *PostInstanceRequest) UnsetK6OrgId()`

UnsetK6OrgId ensures that no value is present for K6OrgId, not even an explicit nil
### GetLabels

`func (o *PostInstanceRequest) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PostInstanceRequest) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PostInstanceRequest) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PostInstanceRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


