# StackUpdateRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **NullableBool** | protection active or not for the stack | [optional] 
**Description** | Pointer to **NullableString** | description of the stack | [optional] 
**Labels** | Pointer to **map[string]string** | labels to add to the stack | [optional] 
**Name** | Pointer to **NullableString** | of the stack | [optional] 
**Slug** | Pointer to **NullableString** | slug of the stack | [optional] 
**Url** | Pointer to **NullableString** | url for the stack | [optional] 

## Methods

### NewStackUpdateRequestV1

`func NewStackUpdateRequestV1() *StackUpdateRequestV1`

NewStackUpdateRequestV1 instantiates a new StackUpdateRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackUpdateRequestV1WithDefaults

`func NewStackUpdateRequestV1WithDefaults() *StackUpdateRequestV1`

NewStackUpdateRequestV1WithDefaults instantiates a new StackUpdateRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *StackUpdateRequestV1) GetDeleteProtection() bool`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *StackUpdateRequestV1) GetDeleteProtectionOk() (*bool, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *StackUpdateRequestV1) SetDeleteProtection(v bool)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *StackUpdateRequestV1) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### SetDeleteProtectionNil

`func (o *StackUpdateRequestV1) SetDeleteProtectionNil(b bool)`

 SetDeleteProtectionNil sets the value for DeleteProtection to be an explicit nil

### UnsetDeleteProtection
`func (o *StackUpdateRequestV1) UnsetDeleteProtection()`

UnsetDeleteProtection ensures that no value is present for DeleteProtection, not even an explicit nil
### GetDescription

`func (o *StackUpdateRequestV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StackUpdateRequestV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StackUpdateRequestV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StackUpdateRequestV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StackUpdateRequestV1) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StackUpdateRequestV1) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *StackUpdateRequestV1) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StackUpdateRequestV1) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StackUpdateRequestV1) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *StackUpdateRequestV1) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *StackUpdateRequestV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackUpdateRequestV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackUpdateRequestV1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StackUpdateRequestV1) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *StackUpdateRequestV1) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *StackUpdateRequestV1) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSlug

`func (o *StackUpdateRequestV1) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *StackUpdateRequestV1) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *StackUpdateRequestV1) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *StackUpdateRequestV1) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *StackUpdateRequestV1) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *StackUpdateRequestV1) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetUrl

`func (o *StackUpdateRequestV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StackUpdateRequestV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StackUpdateRequestV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StackUpdateRequestV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StackUpdateRequestV1) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StackUpdateRequestV1) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


