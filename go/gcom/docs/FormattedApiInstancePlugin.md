# FormattedApiInstancePlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**InstanceId** | **float32** |  | 
**InstanceUrl** | **string** |  | 
**InstanceSlug** | **string** |  | 
**PluginId** | **float32** |  | 
**PluginSlug** | **string** |  | 
**PluginName** | **string** |  | 
**Version** | **string** |  | 
**LatestVersion** | **string** |  | 
**CreatedAt** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Links** | [**[]LinksInner**](LinksInner.md) |  | 

## Methods

### NewFormattedApiInstancePlugin

`func NewFormattedApiInstancePlugin(id float32, instanceId float32, instanceUrl string, instanceSlug string, pluginId float32, pluginSlug string, pluginName string, version string, latestVersion string, createdAt string, links []LinksInner, ) *FormattedApiInstancePlugin`

NewFormattedApiInstancePlugin instantiates a new FormattedApiInstancePlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedApiInstancePluginWithDefaults

`func NewFormattedApiInstancePluginWithDefaults() *FormattedApiInstancePlugin`

NewFormattedApiInstancePluginWithDefaults instantiates a new FormattedApiInstancePlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FormattedApiInstancePlugin) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormattedApiInstancePlugin) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormattedApiInstancePlugin) SetId(v float32)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *FormattedApiInstancePlugin) GetInstanceId() float32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FormattedApiInstancePlugin) GetInstanceIdOk() (*float32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FormattedApiInstancePlugin) SetInstanceId(v float32)`

SetInstanceId sets InstanceId field to given value.


### GetInstanceUrl

`func (o *FormattedApiInstancePlugin) GetInstanceUrl() string`

GetInstanceUrl returns the InstanceUrl field if non-nil, zero value otherwise.

### GetInstanceUrlOk

`func (o *FormattedApiInstancePlugin) GetInstanceUrlOk() (*string, bool)`

GetInstanceUrlOk returns a tuple with the InstanceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUrl

`func (o *FormattedApiInstancePlugin) SetInstanceUrl(v string)`

SetInstanceUrl sets InstanceUrl field to given value.


### GetInstanceSlug

`func (o *FormattedApiInstancePlugin) GetInstanceSlug() string`

GetInstanceSlug returns the InstanceSlug field if non-nil, zero value otherwise.

### GetInstanceSlugOk

`func (o *FormattedApiInstancePlugin) GetInstanceSlugOk() (*string, bool)`

GetInstanceSlugOk returns a tuple with the InstanceSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSlug

`func (o *FormattedApiInstancePlugin) SetInstanceSlug(v string)`

SetInstanceSlug sets InstanceSlug field to given value.


### GetPluginId

`func (o *FormattedApiInstancePlugin) GetPluginId() float32`

GetPluginId returns the PluginId field if non-nil, zero value otherwise.

### GetPluginIdOk

`func (o *FormattedApiInstancePlugin) GetPluginIdOk() (*float32, bool)`

GetPluginIdOk returns a tuple with the PluginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginId

`func (o *FormattedApiInstancePlugin) SetPluginId(v float32)`

SetPluginId sets PluginId field to given value.


### GetPluginSlug

`func (o *FormattedApiInstancePlugin) GetPluginSlug() string`

GetPluginSlug returns the PluginSlug field if non-nil, zero value otherwise.

### GetPluginSlugOk

`func (o *FormattedApiInstancePlugin) GetPluginSlugOk() (*string, bool)`

GetPluginSlugOk returns a tuple with the PluginSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginSlug

`func (o *FormattedApiInstancePlugin) SetPluginSlug(v string)`

SetPluginSlug sets PluginSlug field to given value.


### GetPluginName

`func (o *FormattedApiInstancePlugin) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *FormattedApiInstancePlugin) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *FormattedApiInstancePlugin) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.


### GetVersion

`func (o *FormattedApiInstancePlugin) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FormattedApiInstancePlugin) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FormattedApiInstancePlugin) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetLatestVersion

`func (o *FormattedApiInstancePlugin) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *FormattedApiInstancePlugin) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *FormattedApiInstancePlugin) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.


### GetCreatedAt

`func (o *FormattedApiInstancePlugin) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormattedApiInstancePlugin) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormattedApiInstancePlugin) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *FormattedApiInstancePlugin) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormattedApiInstancePlugin) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormattedApiInstancePlugin) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FormattedApiInstancePlugin) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetLinks

`func (o *FormattedApiInstancePlugin) GetLinks() []LinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FormattedApiInstancePlugin) GetLinksOk() (*[]LinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FormattedApiInstancePlugin) SetLinks(v []LinksInner)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


