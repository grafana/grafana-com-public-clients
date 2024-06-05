# AuthTokenWithSecret

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
**Token** | Pointer to **string** | This token is auto generated and will be shown only once. | [optional] 

## Methods

### NewAuthTokenWithSecret

`func NewAuthTokenWithSecret(accessPolicyId string, name string, ) *AuthTokenWithSecret`

NewAuthTokenWithSecret instantiates a new AuthTokenWithSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenWithSecretWithDefaults

`func NewAuthTokenWithSecretWithDefaults() *AuthTokenWithSecret`

NewAuthTokenWithSecretWithDefaults instantiates a new AuthTokenWithSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthTokenWithSecret) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthTokenWithSecret) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthTokenWithSecret) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthTokenWithSecret) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccessPolicyId

`func (o *AuthTokenWithSecret) GetAccessPolicyId() string`

GetAccessPolicyId returns the AccessPolicyId field if non-nil, zero value otherwise.

### GetAccessPolicyIdOk

`func (o *AuthTokenWithSecret) GetAccessPolicyIdOk() (*string, bool)`

GetAccessPolicyIdOk returns a tuple with the AccessPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPolicyId

`func (o *AuthTokenWithSecret) SetAccessPolicyId(v string)`

SetAccessPolicyId sets AccessPolicyId field to given value.


### GetName

`func (o *AuthTokenWithSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthTokenWithSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthTokenWithSecret) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *AuthTokenWithSecret) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AuthTokenWithSecret) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AuthTokenWithSecret) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AuthTokenWithSecret) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExpiresAt

`func (o *AuthTokenWithSecret) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *AuthTokenWithSecret) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *AuthTokenWithSecret) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *AuthTokenWithSecret) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetFirstUsedAt

`func (o *AuthTokenWithSecret) GetFirstUsedAt() time.Time`

GetFirstUsedAt returns the FirstUsedAt field if non-nil, zero value otherwise.

### GetFirstUsedAtOk

`func (o *AuthTokenWithSecret) GetFirstUsedAtOk() (*time.Time, bool)`

GetFirstUsedAtOk returns a tuple with the FirstUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstUsedAt

`func (o *AuthTokenWithSecret) SetFirstUsedAt(v time.Time)`

SetFirstUsedAt sets FirstUsedAt field to given value.

### HasFirstUsedAt

`func (o *AuthTokenWithSecret) HasFirstUsedAt() bool`

HasFirstUsedAt returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *AuthTokenWithSecret) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *AuthTokenWithSecret) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *AuthTokenWithSecret) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *AuthTokenWithSecret) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AuthTokenWithSecret) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthTokenWithSecret) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthTokenWithSecret) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthTokenWithSecret) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthTokenWithSecret) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthTokenWithSecret) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthTokenWithSecret) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthTokenWithSecret) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetOrgId

`func (o *AuthTokenWithSecret) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AuthTokenWithSecret) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AuthTokenWithSecret) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AuthTokenWithSecret) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetToken

`func (o *AuthTokenWithSecret) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthTokenWithSecret) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthTokenWithSecret) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthTokenWithSecret) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


