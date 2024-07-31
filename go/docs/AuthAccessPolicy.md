# AuthAccessPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Name** | **string** |  | 
**DisplayName** | Pointer to **string** | Will be set to &#x60;name&#x60; if not provided. | [optional] 
**Source** | Pointer to **string** | Source of the Access Policy (requires system token). | [optional] 
**Scopes** | **[]string** |  | 
**Realms** | [**[]AuthAccessPolicyRealmsInner**](AuthAccessPolicyRealmsInner.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Conditions** | Pointer to [**AuthAccessPolicyConditions**](AuthAccessPolicyConditions.md) |  | [optional] 
**Attributes** | Pointer to [**AuthAccessPolicyAttributes**](AuthAccessPolicyAttributes.md) |  | [optional] 
**Status** | Pointer to **string** | The status of the access policy. | [optional] 

## Methods

### NewAuthAccessPolicy

`func NewAuthAccessPolicy(name string, scopes []string, realms []AuthAccessPolicyRealmsInner, ) *AuthAccessPolicy`

NewAuthAccessPolicy instantiates a new AuthAccessPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAccessPolicyWithDefaults

`func NewAuthAccessPolicyWithDefaults() *AuthAccessPolicy`

NewAuthAccessPolicyWithDefaults instantiates a new AuthAccessPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthAccessPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthAccessPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthAccessPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthAccessPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *AuthAccessPolicy) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *AuthAccessPolicy) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *AuthAccessPolicy) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *AuthAccessPolicy) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetName

`func (o *AuthAccessPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthAccessPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthAccessPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *AuthAccessPolicy) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AuthAccessPolicy) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AuthAccessPolicy) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AuthAccessPolicy) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSource

`func (o *AuthAccessPolicy) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AuthAccessPolicy) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AuthAccessPolicy) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *AuthAccessPolicy) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetScopes

`func (o *AuthAccessPolicy) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AuthAccessPolicy) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AuthAccessPolicy) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetRealms

`func (o *AuthAccessPolicy) GetRealms() []AuthAccessPolicyRealmsInner`

GetRealms returns the Realms field if non-nil, zero value otherwise.

### GetRealmsOk

`func (o *AuthAccessPolicy) GetRealmsOk() (*[]AuthAccessPolicyRealmsInner, bool)`

GetRealmsOk returns a tuple with the Realms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealms

`func (o *AuthAccessPolicy) SetRealms(v []AuthAccessPolicyRealmsInner)`

SetRealms sets Realms field to given value.


### GetCreatedAt

`func (o *AuthAccessPolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AuthAccessPolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AuthAccessPolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AuthAccessPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AuthAccessPolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AuthAccessPolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AuthAccessPolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AuthAccessPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetConditions

`func (o *AuthAccessPolicy) GetConditions() AuthAccessPolicyConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AuthAccessPolicy) GetConditionsOk() (*AuthAccessPolicyConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AuthAccessPolicy) SetConditions(v AuthAccessPolicyConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AuthAccessPolicy) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetAttributes

`func (o *AuthAccessPolicy) GetAttributes() AuthAccessPolicyAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AuthAccessPolicy) GetAttributesOk() (*AuthAccessPolicyAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AuthAccessPolicy) SetAttributes(v AuthAccessPolicyAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AuthAccessPolicy) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetStatus

`func (o *AuthAccessPolicy) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthAccessPolicy) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthAccessPolicy) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthAccessPolicy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


