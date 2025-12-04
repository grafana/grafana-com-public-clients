# AddedComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | **[]string** | The extension point ids your plugin registers the extension to, e.g. &#x60;[\&quot;grafana/user/profile/tab\&quot;]&#x60; | 
**Title** | **string** | The title of your component extension. | 
**Description** | Pointer to **string** | Additional information about your component extension. | [optional] 

## Methods

### NewAddedComponentsInner

`func NewAddedComponentsInner(targets []string, title string, ) *AddedComponentsInner`

NewAddedComponentsInner instantiates a new AddedComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedComponentsInnerWithDefaults

`func NewAddedComponentsInnerWithDefaults() *AddedComponentsInner`

NewAddedComponentsInnerWithDefaults instantiates a new AddedComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *AddedComponentsInner) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AddedComponentsInner) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AddedComponentsInner) SetTargets(v []string)`

SetTargets sets Targets field to given value.


### GetTitle

`func (o *AddedComponentsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AddedComponentsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AddedComponentsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *AddedComponentsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddedComponentsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddedComponentsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddedComponentsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


