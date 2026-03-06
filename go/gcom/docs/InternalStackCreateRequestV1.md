# InternalStackCreateRequestV1

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
**TrialExpiresAt** | Pointer to **NullableTime** | trial expiration date for the stack | [optional] 
**Url** | Pointer to **NullableString** | url for the stack | [optional] 

## Methods

### NewInternalStackCreateRequestV1

`func NewInternalStackCreateRequestV1(name string, org string, region string, slug string, ) *InternalStackCreateRequestV1`

NewInternalStackCreateRequestV1 instantiates a new InternalStackCreateRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalStackCreateRequestV1WithDefaults

`func NewInternalStackCreateRequestV1WithDefaults() *InternalStackCreateRequestV1`

NewInternalStackCreateRequestV1WithDefaults instantiates a new InternalStackCreateRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *InternalStackCreateRequestV1) GetDeleteProtection() bool`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *InternalStackCreateRequestV1) GetDeleteProtectionOk() (*bool, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *InternalStackCreateRequestV1) SetDeleteProtection(v bool)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *InternalStackCreateRequestV1) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *InternalStackCreateRequestV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalStackCreateRequestV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalStackCreateRequestV1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternalStackCreateRequestV1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InternalStackCreateRequestV1) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InternalStackCreateRequestV1) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *InternalStackCreateRequestV1) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InternalStackCreateRequestV1) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InternalStackCreateRequestV1) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InternalStackCreateRequestV1) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *InternalStackCreateRequestV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalStackCreateRequestV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalStackCreateRequestV1) SetName(v string)`

SetName sets Name field to given value.


### GetOrg

`func (o *InternalStackCreateRequestV1) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *InternalStackCreateRequestV1) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *InternalStackCreateRequestV1) SetOrg(v string)`

SetOrg sets Org field to given value.


### GetPlugins

`func (o *InternalStackCreateRequestV1) GetPlugins() []string`

GetPlugins returns the Plugins field if non-nil, zero value otherwise.

### GetPluginsOk

`func (o *InternalStackCreateRequestV1) GetPluginsOk() (*[]string, bool)`

GetPluginsOk returns a tuple with the Plugins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugins

`func (o *InternalStackCreateRequestV1) SetPlugins(v []string)`

SetPlugins sets Plugins field to given value.

### HasPlugins

`func (o *InternalStackCreateRequestV1) HasPlugins() bool`

HasPlugins returns a boolean if a field has been set.

### SetPluginsNil

`func (o *InternalStackCreateRequestV1) SetPluginsNil(b bool)`

 SetPluginsNil sets the value for Plugins to be an explicit nil

### UnsetPlugins
`func (o *InternalStackCreateRequestV1) UnsetPlugins()`

UnsetPlugins ensures that no value is present for Plugins, not even an explicit nil
### GetRegion

`func (o *InternalStackCreateRequestV1) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InternalStackCreateRequestV1) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InternalStackCreateRequestV1) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSlug

`func (o *InternalStackCreateRequestV1) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *InternalStackCreateRequestV1) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *InternalStackCreateRequestV1) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTrialExpiresAt

`func (o *InternalStackCreateRequestV1) GetTrialExpiresAt() time.Time`

GetTrialExpiresAt returns the TrialExpiresAt field if non-nil, zero value otherwise.

### GetTrialExpiresAtOk

`func (o *InternalStackCreateRequestV1) GetTrialExpiresAtOk() (*time.Time, bool)`

GetTrialExpiresAtOk returns a tuple with the TrialExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialExpiresAt

`func (o *InternalStackCreateRequestV1) SetTrialExpiresAt(v time.Time)`

SetTrialExpiresAt sets TrialExpiresAt field to given value.

### HasTrialExpiresAt

`func (o *InternalStackCreateRequestV1) HasTrialExpiresAt() bool`

HasTrialExpiresAt returns a boolean if a field has been set.

### SetTrialExpiresAtNil

`func (o *InternalStackCreateRequestV1) SetTrialExpiresAtNil(b bool)`

 SetTrialExpiresAtNil sets the value for TrialExpiresAt to be an explicit nil

### UnsetTrialExpiresAt
`func (o *InternalStackCreateRequestV1) UnsetTrialExpiresAt()`

UnsetTrialExpiresAt ensures that no value is present for TrialExpiresAt, not even an explicit nil
### GetUrl

`func (o *InternalStackCreateRequestV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InternalStackCreateRequestV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InternalStackCreateRequestV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *InternalStackCreateRequestV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *InternalStackCreateRequestV1) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *InternalStackCreateRequestV1) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


