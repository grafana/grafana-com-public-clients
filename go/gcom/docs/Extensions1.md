# Extensions1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedComponents** | Pointer to [**[]AddedComponentsInner**](AddedComponentsInner.md) | This list must contain all component extensions that your plugin registers to other extension points using [&#x60;.addComponent()&#x60;](https://grafana.com/developers/plugin-tools/reference/ui-extensions-reference/ui-extensions#addcomponent). **Components that are not listed here won&#39;t work.** | [optional] 
**AddedLinks** | Pointer to [**[]AddedLinksInner**](AddedLinksInner.md) | This list must contain all link extensions that your plugin registers to other extension points using [&#x60;.addLink()&#x60;](https://grafana.com/developers/plugin-tools/reference/ui-extensions-reference/ui-extensions#addlink). **Links that are not listed here won&#39;t work.** | [optional] 
**AddedFunctions** | Pointer to [**[]AddedFunctionsInner**](AddedFunctionsInner.md) | This list must contain all function extensions that your plugin registers to other extension points using [&#x60;.addedFunctions()&#x60;](https://grafana.com/developers/plugin-tools/reference/ui-extensions-reference/ui-extensions#addfunction). **Functions that are not listed here won&#39;t work.** | [optional] 
**ExposedComponents** | Pointer to [**[]ExposedComponentsInner**](ExposedComponentsInner.md) | This list must contain all components that your plugin exposes using [&#x60;.exposeComponent()&#x60;](https://grafana.com/developers/plugin-tools/reference/ui-extensions-reference/ui-extensions#exposecomponent). **Components that are not listed here won&#39;t work.** | [optional] 
**ExtensionPoints** | Pointer to [**[]ExtensionPointsInner**](ExtensionPointsInner.md) | This list must contain all extension points that your plugin defines using [&#x60;usePluginLinks()&#x60;](https://grafana.com/developers/plugin-tools/reference/ui-extensions-reference/ui-extensions#usepluginlinks) or [&#x60;usePluginComponents()&#x60;](https://grafana.com/developers/plugin-tools/reference/ui-extensions-reference/ui-extensions#useplugincomponents). **Extension points that are not listed in here won&#39;t work.** | [optional] 

## Methods

### NewExtensions1

`func NewExtensions1() *Extensions1`

NewExtensions1 instantiates a new Extensions1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensions1WithDefaults

`func NewExtensions1WithDefaults() *Extensions1`

NewExtensions1WithDefaults instantiates a new Extensions1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedComponents

`func (o *Extensions1) GetAddedComponents() []AddedComponentsInner`

GetAddedComponents returns the AddedComponents field if non-nil, zero value otherwise.

### GetAddedComponentsOk

`func (o *Extensions1) GetAddedComponentsOk() (*[]AddedComponentsInner, bool)`

GetAddedComponentsOk returns a tuple with the AddedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedComponents

`func (o *Extensions1) SetAddedComponents(v []AddedComponentsInner)`

SetAddedComponents sets AddedComponents field to given value.

### HasAddedComponents

`func (o *Extensions1) HasAddedComponents() bool`

HasAddedComponents returns a boolean if a field has been set.

### GetAddedLinks

`func (o *Extensions1) GetAddedLinks() []AddedLinksInner`

GetAddedLinks returns the AddedLinks field if non-nil, zero value otherwise.

### GetAddedLinksOk

`func (o *Extensions1) GetAddedLinksOk() (*[]AddedLinksInner, bool)`

GetAddedLinksOk returns a tuple with the AddedLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedLinks

`func (o *Extensions1) SetAddedLinks(v []AddedLinksInner)`

SetAddedLinks sets AddedLinks field to given value.

### HasAddedLinks

`func (o *Extensions1) HasAddedLinks() bool`

HasAddedLinks returns a boolean if a field has been set.

### GetAddedFunctions

`func (o *Extensions1) GetAddedFunctions() []AddedFunctionsInner`

GetAddedFunctions returns the AddedFunctions field if non-nil, zero value otherwise.

### GetAddedFunctionsOk

`func (o *Extensions1) GetAddedFunctionsOk() (*[]AddedFunctionsInner, bool)`

GetAddedFunctionsOk returns a tuple with the AddedFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedFunctions

`func (o *Extensions1) SetAddedFunctions(v []AddedFunctionsInner)`

SetAddedFunctions sets AddedFunctions field to given value.

### HasAddedFunctions

`func (o *Extensions1) HasAddedFunctions() bool`

HasAddedFunctions returns a boolean if a field has been set.

### GetExposedComponents

`func (o *Extensions1) GetExposedComponents() []ExposedComponentsInner`

GetExposedComponents returns the ExposedComponents field if non-nil, zero value otherwise.

### GetExposedComponentsOk

`func (o *Extensions1) GetExposedComponentsOk() (*[]ExposedComponentsInner, bool)`

GetExposedComponentsOk returns a tuple with the ExposedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedComponents

`func (o *Extensions1) SetExposedComponents(v []ExposedComponentsInner)`

SetExposedComponents sets ExposedComponents field to given value.

### HasExposedComponents

`func (o *Extensions1) HasExposedComponents() bool`

HasExposedComponents returns a boolean if a field has been set.

### GetExtensionPoints

`func (o *Extensions1) GetExtensionPoints() []ExtensionPointsInner`

GetExtensionPoints returns the ExtensionPoints field if non-nil, zero value otherwise.

### GetExtensionPointsOk

`func (o *Extensions1) GetExtensionPointsOk() (*[]ExtensionPointsInner, bool)`

GetExtensionPointsOk returns a tuple with the ExtensionPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionPoints

`func (o *Extensions1) SetExtensionPoints(v []ExtensionPointsInner)`

SetExtensionPoints sets ExtensionPoints field to given value.

### HasExtensionPoints

`func (o *Extensions1) HasExtensionPoints() bool`

HasExtensionPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


