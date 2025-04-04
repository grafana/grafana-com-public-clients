# PostInstancePluginsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoRestart** | Pointer to **bool** |  | [optional] 
**Plugin** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewPostInstancePluginsRequest

`func NewPostInstancePluginsRequest(plugin string, ) *PostInstancePluginsRequest`

NewPostInstancePluginsRequest instantiates a new PostInstancePluginsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostInstancePluginsRequestWithDefaults

`func NewPostInstancePluginsRequestWithDefaults() *PostInstancePluginsRequest`

NewPostInstancePluginsRequestWithDefaults instantiates a new PostInstancePluginsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoRestart

`func (o *PostInstancePluginsRequest) GetNoRestart() bool`

GetNoRestart returns the NoRestart field if non-nil, zero value otherwise.

### GetNoRestartOk

`func (o *PostInstancePluginsRequest) GetNoRestartOk() (*bool, bool)`

GetNoRestartOk returns a tuple with the NoRestart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoRestart

`func (o *PostInstancePluginsRequest) SetNoRestart(v bool)`

SetNoRestart sets NoRestart field to given value.

### HasNoRestart

`func (o *PostInstancePluginsRequest) HasNoRestart() bool`

HasNoRestart returns a boolean if a field has been set.

### GetPlugin

`func (o *PostInstancePluginsRequest) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *PostInstancePluginsRequest) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *PostInstancePluginsRequest) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.


### GetVersion

`func (o *PostInstancePluginsRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *PostInstancePluginsRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *PostInstancePluginsRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *PostInstancePluginsRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


