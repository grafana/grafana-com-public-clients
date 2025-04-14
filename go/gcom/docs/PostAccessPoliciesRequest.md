# PostAccessPoliciesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to [**NullablePostAccessPoliciesRequestAttributes**](PostAccessPoliciesRequestAttributes.md) |  | [optional] 
**Conditions** | Pointer to [**NullablePostAccessPoliciesRequestConditions**](PostAccessPoliciesRequestConditions.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Realms** | [**[]PostAccessPoliciesRequestRealmsInner**](PostAccessPoliciesRequestRealmsInner.md) |  | 
**Scopes** | **[]string** |  | 

## Methods

### NewPostAccessPoliciesRequest

`func NewPostAccessPoliciesRequest(name string, realms []PostAccessPoliciesRequestRealmsInner, scopes []string, ) *PostAccessPoliciesRequest`

NewPostAccessPoliciesRequest instantiates a new PostAccessPoliciesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAccessPoliciesRequestWithDefaults

`func NewPostAccessPoliciesRequestWithDefaults() *PostAccessPoliciesRequest`

NewPostAccessPoliciesRequestWithDefaults instantiates a new PostAccessPoliciesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *PostAccessPoliciesRequest) GetAttributes() PostAccessPoliciesRequestAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *PostAccessPoliciesRequest) GetAttributesOk() (*PostAccessPoliciesRequestAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *PostAccessPoliciesRequest) SetAttributes(v PostAccessPoliciesRequestAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *PostAccessPoliciesRequest) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributesNil

`func (o *PostAccessPoliciesRequest) SetAttributesNil(b bool)`

 SetAttributesNil sets the value for Attributes to be an explicit nil

### UnsetAttributes
`func (o *PostAccessPoliciesRequest) UnsetAttributes()`

UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
### GetConditions

`func (o *PostAccessPoliciesRequest) GetConditions() PostAccessPoliciesRequestConditions`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *PostAccessPoliciesRequest) GetConditionsOk() (*PostAccessPoliciesRequestConditions, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *PostAccessPoliciesRequest) SetConditions(v PostAccessPoliciesRequestConditions)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *PostAccessPoliciesRequest) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### SetConditionsNil

`func (o *PostAccessPoliciesRequest) SetConditionsNil(b bool)`

 SetConditionsNil sets the value for Conditions to be an explicit nil

### UnsetConditions
`func (o *PostAccessPoliciesRequest) UnsetConditions()`

UnsetConditions ensures that no value is present for Conditions, not even an explicit nil
### GetDisplayName

`func (o *PostAccessPoliciesRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PostAccessPoliciesRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PostAccessPoliciesRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PostAccessPoliciesRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetName

`func (o *PostAccessPoliciesRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostAccessPoliciesRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostAccessPoliciesRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRealms

`func (o *PostAccessPoliciesRequest) GetRealms() []PostAccessPoliciesRequestRealmsInner`

GetRealms returns the Realms field if non-nil, zero value otherwise.

### GetRealmsOk

`func (o *PostAccessPoliciesRequest) GetRealmsOk() (*[]PostAccessPoliciesRequestRealmsInner, bool)`

GetRealmsOk returns a tuple with the Realms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealms

`func (o *PostAccessPoliciesRequest) SetRealms(v []PostAccessPoliciesRequestRealmsInner)`

SetRealms sets Realms field to given value.


### GetScopes

`func (o *PostAccessPoliciesRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PostAccessPoliciesRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PostAccessPoliciesRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


