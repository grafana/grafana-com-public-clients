# AuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**AccessPolicyId** | **string** |  | 
**Name** | **string** |  | 
**DisplayName** | Pointer to **string** | Will be set to &#x60;name&#x60; if not provided. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Token does not expire if not provided. | [optional] 
**FirstUsedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**LastUsedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewAuthToken

`func NewAuthToken(accessPolicyId string, name string, ) *AuthToken`

NewAuthToken instantiates a new AuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenWithDefaults

`func NewAuthTokenWithDefaults() *AuthToken`

NewAuthTokenWithDefaults instantiates a new AuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccessPolicyId

`func (o *AuthToken) GetAccessPolicyId() string`

GetAccessPolicyId returns the AccessPolicyId field if non-nil, zero value otherwise.

### GetAccessPolicyIdOk

`func (o *AuthToken) GetAccessPolicyIdOk() (*string, bool)`

GetAccessPolicyIdOk returns a tuple with the AccessPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyId

`func (o *AuthToken) SetAccessPolicyId(v string)`

SetAccessPolicyId sets AccessPolicyId field to given value.


### GetName

`func (o *AuthToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthToken) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *AuthToken) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AuthToken) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AuthToken) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AuthToken) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AuthToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AuthToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AuthToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AuthToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFirstUsedAt

`func (o *AuthToken) GetFirstUsedAt() time.Time`

GetFirstUsedAt returns the FirstUsedAt field if non-nil, zero value otherwise.

### GetFirstUsedAtOk

`func (o *AuthToken) GetFirstUsedAtOk() (*time.Time, bool)`

GetFirstUsedAtOk returns a tuple with the FirstUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUsedAt

`func (o *AuthToken) SetFirstUsedAt(v time.Time)`

SetFirstUsedAt sets FirstUsedAt field to given value.

### HasFirstUsedAt

`func (o *AuthToken) HasFirstUsedAt() bool`

HasFirstUsedAt returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *AuthToken) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *AuthToken) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *AuthToken) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *AuthToken) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthToken) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthToken) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthToken) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthToken) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOrgId

`func (o *AuthToken) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AuthToken) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AuthToken) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AuthToken) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


