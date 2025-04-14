# PostAccessPolicyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**NullablePostAccessPoliciesRequestAttributes**](PostAccessPoliciesRequestAttributes.md) |  | [optional] 
**Conditions** | Pointer to [**NullablePostAccessPoliciesRequestConditions**](PostAccessPoliciesRequestConditions.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Realms** | Pointer to [**[]PostAccessPoliciesRequestRealmsInner**](PostAccessPoliciesRequestRealmsInner.md) |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewPostAccessPolicyRequest

`func NewPostAccessPolicyRequest() *PostAccessPolicyRequest`

NewPostAccessPolicyRequest instantiates a new PostAccessPolicyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAccessPolicyRequestWithDefaults

`func NewPostAccessPolicyRequestWithDefaults() *PostAccessPolicyRequest`

NewPostAccessPolicyRequestWithDefaults instantiates a new PostAccessPolicyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *PostAccessPolicyRequest) GetAttributes() PostAccessPoliciesRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostAccessPolicyRequest) GetAttributesOk() (*PostAccessPoliciesRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostAccessPolicyRequest) SetAttributes(v PostAccessPoliciesRequestAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PostAccessPolicyRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *PostAccessPolicyRequest) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *PostAccessPolicyRequest) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetConditions

`func (o *PostAccessPolicyRequest) GetConditions() PostAccessPoliciesRequestConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PostAccessPolicyRequest) GetConditionsOk() (*PostAccessPoliciesRequestConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PostAccessPolicyRequest) SetConditions(v PostAccessPoliciesRequestConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PostAccessPolicyRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *PostAccessPolicyRequest) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *PostAccessPolicyRequest) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetDisplayName

`func (o *PostAccessPolicyRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PostAccessPolicyRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PostAccessPolicyRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PostAccessPolicyRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRealms

`func (o *PostAccessPolicyRequest) GetRealms() []PostAccessPoliciesRequestRealmsInner`

GetRealms returns the Realms field if non-nil, zero value otherwise.

### GetRealmsOk

`func (o *PostAccessPolicyRequest) GetRealmsOk() (*[]PostAccessPoliciesRequestRealmsInner, bool)`

GetRealmsOk returns a tuple with the Realms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealms

`func (o *PostAccessPolicyRequest) SetRealms(v []PostAccessPoliciesRequestRealmsInner)`

SetRealms sets Realms field to given value.

### HasRealms

`func (o *PostAccessPolicyRequest) HasRealms() bool`

HasRealms returns a boolean if a field has been set.

### GetScopes

`func (o *PostAccessPolicyRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PostAccessPolicyRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PostAccessPolicyRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PostAccessPolicyRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetStatus

`func (o *PostAccessPolicyRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostAccessPolicyRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostAccessPolicyRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostAccessPolicyRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


