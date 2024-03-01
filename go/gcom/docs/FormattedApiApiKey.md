# FormattedApiApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **float32** |  | [optional] 
**OrgId** | Pointer to **float32** |  | [optional] 
**OrgSlug** | Pointer to **string** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **NullableFloat32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**FirstUsed** | Pointer to **NullableString** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Links** | Pointer to [**[]LinksInner**](LinksInner.md) |  | [optional] 

## Methods

### NewFormattedApiApiKey

`func NewFormattedApiApiKey() *FormattedApiApiKey`

NewFormattedApiApiKey instantiates a new FormattedApiApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedApiApiKeyWithDefaults

`func NewFormattedApiApiKeyWithDefaults() *FormattedApiApiKey`

NewFormattedApiApiKeyWithDefaults instantiates a new FormattedApiApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FormattedApiApiKey) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormattedApiApiKey) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormattedApiApiKey) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *FormattedApiApiKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *FormattedApiApiKey) GetOrgId() float32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *FormattedApiApiKey) GetOrgIdOk() (*float32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *FormattedApiApiKey) SetOrgId(v float32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *FormattedApiApiKey) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgSlug

`func (o *FormattedApiApiKey) GetOrgSlug() string`

GetOrgSlug returns the OrgSlug field if non-nil, zero value otherwise.

### GetOrgSlugOk

`func (o *FormattedApiApiKey) GetOrgSlugOk() (*string, bool)`

GetOrgSlugOk returns a tuple with the OrgSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgSlug

`func (o *FormattedApiApiKey) SetOrgSlug(v string)`

SetOrgSlug sets OrgSlug field to given value.

### HasOrgSlug

`func (o *FormattedApiApiKey) HasOrgSlug() bool`

HasOrgSlug returns a boolean if a field has been set.

### GetOrgName

`func (o *FormattedApiApiKey) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *FormattedApiApiKey) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *FormattedApiApiKey) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *FormattedApiApiKey) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetInstanceId

`func (o *FormattedApiApiKey) GetInstanceId() float32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FormattedApiApiKey) GetInstanceIdOk() (*float32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FormattedApiApiKey) SetInstanceId(v float32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *FormattedApiApiKey) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *FormattedApiApiKey) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *FormattedApiApiKey) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetName

`func (o *FormattedApiApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormattedApiApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormattedApiApiKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FormattedApiApiKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *FormattedApiApiKey) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *FormattedApiApiKey) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *FormattedApiApiKey) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *FormattedApiApiKey) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FormattedApiApiKey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormattedApiApiKey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormattedApiApiKey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FormattedApiApiKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FormattedApiApiKey) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormattedApiApiKey) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormattedApiApiKey) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FormattedApiApiKey) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *FormattedApiApiKey) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *FormattedApiApiKey) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetFirstUsed

`func (o *FormattedApiApiKey) GetFirstUsed() string`

GetFirstUsed returns the FirstUsed field if non-nil, zero value otherwise.

### GetFirstUsedOk

`func (o *FormattedApiApiKey) GetFirstUsedOk() (*string, bool)`

GetFirstUsedOk returns a tuple with the FirstUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUsed

`func (o *FormattedApiApiKey) SetFirstUsed(v string)`

SetFirstUsed sets FirstUsed field to given value.

### HasFirstUsed

`func (o *FormattedApiApiKey) HasFirstUsed() bool`

HasFirstUsed returns a boolean if a field has been set.

### SetFirstUsedNil

`func (o *FormattedApiApiKey) SetFirstUsedNil(b bool)`

 SetFirstUsedNil sets the value for FirstUsed to be an explicit nil

### UnsetFirstUsed
`func (o *FormattedApiApiKey) UnsetFirstUsed()`

UnsetFirstUsed ensures that no value is present for FirstUsed, not even an explicit nil
### GetToken

`func (o *FormattedApiApiKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FormattedApiApiKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FormattedApiApiKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FormattedApiApiKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetLinks

`func (o *FormattedApiApiKey) GetLinks() []LinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FormattedApiApiKey) GetLinksOk() (*[]LinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FormattedApiApiKey) SetLinks(v []LinksInner)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FormattedApiApiKey) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


