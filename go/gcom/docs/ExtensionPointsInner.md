# ExtensionPointsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for your extension point. This is used to reference the extension point in other plugins. It must be in the following format: &#x60;{PLUGIN_ID}/name-of-my-extension-point/v1&#x60;. It is recommended to add a version suffix to prevent future breaking changes. E.g.: &#x60;myorg-extensions-app/extension-point/v1&#x60;. | 
**Title** | Pointer to **string** | The title of your extension point. | [optional] 
**Description** | Pointer to **string** | Additional information about your extension point. | [optional] 

## Methods

### NewExtensionPointsInner

`func NewExtensionPointsInner(id string, ) *ExtensionPointsInner`

NewExtensionPointsInner instantiates a new ExtensionPointsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionPointsInnerWithDefaults

`func NewExtensionPointsInnerWithDefaults() *ExtensionPointsInner`

NewExtensionPointsInnerWithDefaults instantiates a new ExtensionPointsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtensionPointsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtensionPointsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtensionPointsInner) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ExtensionPointsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExtensionPointsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExtensionPointsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ExtensionPointsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ExtensionPointsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExtensionPointsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExtensionPointsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExtensionPointsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


