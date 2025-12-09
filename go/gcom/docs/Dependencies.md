# Dependencies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrafanaVersion** | Pointer to **string** | (Deprecated) Required Grafana version for this plugin, e.g. &#x60;6.x.x 7.x.x&#x60; to denote plugin requires Grafana v6.x.x or v7.x.x. | [optional] 
**GrafanaDependency** | **string** | Required Grafana version for this plugin. Validated using https://github.com/npm/node-semver. | 
**Plugins** | Pointer to [**[]PluginsInner**](PluginsInner.md) | An array of required plugins on which this plugin depends. Only non-core (that is, external plugins) have to be specified in this list. | [optional] 
**Extensions** | Pointer to [**Extensions**](Extensions.md) |  | [optional] 

## Methods

### NewDependencies

`func NewDependencies(grafanaDependency string, ) *Dependencies`

NewDependencies instantiates a new Dependencies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDependenciesWithDefaults

`func NewDependenciesWithDefaults() *Dependencies`

NewDependenciesWithDefaults instantiates a new Dependencies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrafanaVersion

`func (o *Dependencies) GetGrafanaVersion() string`

GetGrafanaVersion returns the GrafanaVersion field if non-nil, zero value otherwise.

### GetGrafanaVersionOk

`func (o *Dependencies) GetGrafanaVersionOk() (*string, bool)`

GetGrafanaVersionOk returns a tuple with the GrafanaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaVersion

`func (o *Dependencies) SetGrafanaVersion(v string)`

SetGrafanaVersion sets GrafanaVersion field to given value.

### HasGrafanaVersion

`func (o *Dependencies) HasGrafanaVersion() bool`

HasGrafanaVersion returns a boolean if a field has been set.

### GetGrafanaDependency

`func (o *Dependencies) GetGrafanaDependency() string`

GetGrafanaDependency returns the GrafanaDependency field if non-nil, zero value otherwise.

### GetGrafanaDependencyOk

`func (o *Dependencies) GetGrafanaDependencyOk() (*string, bool)`

GetGrafanaDependencyOk returns a tuple with the GrafanaDependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaDependency

`func (o *Dependencies) SetGrafanaDependency(v string)`

SetGrafanaDependency sets GrafanaDependency field to given value.


### GetPlugins

`func (o *Dependencies) GetPlugins() []PluginsInner`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *Dependencies) GetPluginsOk() (*[]PluginsInner, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *Dependencies) SetPlugins(v []PluginsInner)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *Dependencies) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### GetExtensions

`func (o *Dependencies) GetExtensions() Extensions`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Dependencies) GetExtensionsOk() (*Extensions, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Dependencies) SetExtensions(v Extensions)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Dependencies) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


