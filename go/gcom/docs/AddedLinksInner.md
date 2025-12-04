# AddedLinksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | **[]string** | The extension point ids your plugin registers the extension to, e.g. &#x60;[\&quot;grafana/dashboard/panel/menu\&quot;]&#x60; | 
**Title** | **string** | The title of your link extension. | 
**Description** | Pointer to **string** | Additional information about your link extension. | [optional] 

## Methods

### NewAddedLinksInner

`func NewAddedLinksInner(targets []string, title string, ) *AddedLinksInner`

NewAddedLinksInner instantiates a new AddedLinksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedLinksInnerWithDefaults

`func NewAddedLinksInnerWithDefaults() *AddedLinksInner`

NewAddedLinksInnerWithDefaults instantiates a new AddedLinksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *AddedLinksInner) GetTargets() []string`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *AddedLinksInner) GetTargetsOk() (*[]string, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *AddedLinksInner) SetTargets(v []string)`

SetTargets sets Targets field to given value.


### GetTitle

`func (o *AddedLinksInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AddedLinksInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AddedLinksInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *AddedLinksInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddedLinksInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddedLinksInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddedLinksInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


