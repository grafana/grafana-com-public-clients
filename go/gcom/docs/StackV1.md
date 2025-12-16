# StackV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | **bool** | protection active or not for the stack | 
**Description** | **NullableString** | description of the stack | 
**Id** | **int64** | id of the stack | 
**Labels** | **map[string]string** | labels to add to the stack | 
**Name** | **string** | of the stack | 
**OrgId** | **int64** | id of the org | 
**OrgSlug** | **string** | org owning the stack | 
**Region** | **string** | region of the stack | 
**Slug** | **string** | slug of the stack | 
**Url** | **string** | url for the stack | 

## Methods

### NewStackV1

`func NewStackV1(deleteProtection bool, description NullableString, id int64, labels map[string]string, name string, orgId int64, orgSlug string, region string, slug string, url string, ) *StackV1`

NewStackV1 instantiates a new StackV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackV1WithDefaults

`func NewStackV1WithDefaults() *StackV1`

NewStackV1WithDefaults instantiates a new StackV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *StackV1) GetDeleteProtection() bool`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *StackV1) GetDeleteProtectionOk() (*bool, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *StackV1) SetDeleteProtection(v bool)`

SetDeleteProtection sets DeleteProtection field to given value.


### GetDescription

`func (o *StackV1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StackV1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StackV1) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *StackV1) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *StackV1) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *StackV1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackV1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackV1) SetId(v int64)`

SetId sets Id field to given value.


### GetLabels

`func (o *StackV1) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *StackV1) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *StackV1) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetName

`func (o *StackV1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StackV1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StackV1) SetName(v string)`

SetName sets Name field to given value.


### GetOrgId

`func (o *StackV1) GetOrgId() int64`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *StackV1) GetOrgIdOk() (*int64, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *StackV1) SetOrgId(v int64)`

SetOrgId sets OrgId field to given value.


### GetOrgSlug

`func (o *StackV1) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *StackV1) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *StackV1) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.


### GetRegion

`func (o *StackV1) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *StackV1) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *StackV1) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSlug

`func (o *StackV1) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *StackV1) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *StackV1) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetUrl

`func (o *StackV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StackV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StackV1) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


