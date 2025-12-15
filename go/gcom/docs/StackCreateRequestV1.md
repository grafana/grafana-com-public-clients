# StackCreateRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **bool** | protection active or not for the stack | [optional] [default to true]
**Description** | Pointer to **NullableString** | description of the stack | [optional] 
**Labels** | Pointer to **map[string]string** | labels to add to the stack | [optional] 
**Name** | **string** | of the stack | 
**Org** | **string** | org owning the stack | 
**Plugins** | Pointer to **[]string** | plugins to install to the stack | [optional] 
**Region** | **string** | region of the stack | 
**Slug** | **string** | slug of the stack | 
**Url** | Pointer to **NullableString** | url for the stack | [optional] 

## Methods

### NewStackCreateRequestV1

`func NewStackCreateRequestV1(name string, org string, region string, slug string, ) *StackCreateRequestV1`

NewStackCreateRequestV1 instantiates a new StackCreateRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackCreateRequestV1WithDefaults

`func NewStackCreateRequestV1WithDefaults() *StackCreateRequestV1`

NewStackCreateRequestV1WithDefaults instantiates a new StackCreateRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *StackCreateRequestV1) GetDeleteProtection() bool`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *StackCreateRequestV1) GetDeleteProtectionOk() (*bool, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *StackCreateRequestV1) SetDeleteProtection(v bool)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *StackCreateRequestV1) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *StackCreateRequestV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StackCreateRequestV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StackCreateRequestV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StackCreateRequestV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *StackCreateRequestV1) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StackCreateRequestV1) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *StackCreateRequestV1) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StackCreateRequestV1) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StackCreateRequestV1) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *StackCreateRequestV1) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *StackCreateRequestV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackCreateRequestV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackCreateRequestV1) SetName(v string)`

SetName sets Name field to given value.


### GetOrg

`func (o *StackCreateRequestV1) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *StackCreateRequestV1) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *StackCreateRequestV1) SetOrg(v string)`

SetOrg sets Org field to given value.


### GetPlugins

`func (o *StackCreateRequestV1) GetPlugins() []string`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *StackCreateRequestV1) GetPluginsOk() (*[]string, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *StackCreateRequestV1) SetPlugins(v []string)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *StackCreateRequestV1) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### SetPluginsNil

`func (o *StackCreateRequestV1) SetPluginsNil(b bool)`

 SetPluginsNil sets the value for Plugins to be an explicit nil

### UnsetPlugins
`func (o *StackCreateRequestV1) UnsetPlugins()`

UnsetPlugins ensures that no value is present for Plugins, not even an explicit nil
### GetRegion

`func (o *StackCreateRequestV1) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StackCreateRequestV1) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StackCreateRequestV1) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSlug

`func (o *StackCreateRequestV1) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *StackCreateRequestV1) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *StackCreateRequestV1) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetUrl

`func (o *StackCreateRequestV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StackCreateRequestV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StackCreateRequestV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StackCreateRequestV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StackCreateRequestV1) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StackCreateRequestV1) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


