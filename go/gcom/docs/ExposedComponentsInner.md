# ExposedComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for your exposed component. This is used to reference the component in other plugins. It must be in the following format: &#x60;{PLUGIN_ID}/name-of-component/v1&#x60;. It is recommended to add a version suffix to prevent future breaking changes. E.g.: &#x60;myorg-extensions-app/exposed-component/v1&#x60;. | 
**Title** | Pointer to **string** | The title of your exposed component. | [optional] 
**Description** | Pointer to **string** | Additional information about your exposed component. | [optional] 

## Methods

### NewExposedComponentsInner

`func NewExposedComponentsInner(id string, ) *ExposedComponentsInner`

NewExposedComponentsInner instantiates a new ExposedComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExposedComponentsInnerWithDefaults

`func NewExposedComponentsInnerWithDefaults() *ExposedComponentsInner`

NewExposedComponentsInnerWithDefaults instantiates a new ExposedComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExposedComponentsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExposedComponentsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExposedComponentsInner) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ExposedComponentsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ExposedComponentsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ExposedComponentsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ExposedComponentsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ExposedComponentsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExposedComponentsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExposedComponentsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExposedComponentsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


