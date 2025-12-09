# AccessPoliciesConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Realms** | [**[]RealmsInner**](RealmsInner.md) |  | 
**Scopes** | [**[]ScopesInner**](ScopesInner.md) |  | 
**Sets** | [**[]SetsInner**](SetsInner.md) |  | 

## Methods

### NewAccessPoliciesConfig

`func NewAccessPoliciesConfig(realms []RealmsInner, scopes []ScopesInner, sets []SetsInner, ) *AccessPoliciesConfig`

NewAccessPoliciesConfig instantiates a new AccessPoliciesConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPoliciesConfigWithDefaults

`func NewAccessPoliciesConfigWithDefaults() *AccessPoliciesConfig`

NewAccessPoliciesConfigWithDefaults instantiates a new AccessPoliciesConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRealms

`func (o *AccessPoliciesConfig) GetRealms() []RealmsInner`

GetRealms returns the Realms field if non-nil, zero value otherwise.

### GetRealmsOk

`func (o *AccessPoliciesConfig) GetRealmsOk() (*[]RealmsInner, bool)`

GetRealmsOk returns a tuple with the Realms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealms

`func (o *AccessPoliciesConfig) SetRealms(v []RealmsInner)`

SetRealms sets Realms field to given value.


### GetScopes

`func (o *AccessPoliciesConfig) GetScopes() []ScopesInner`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AccessPoliciesConfig) GetScopesOk() (*[]ScopesInner, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AccessPoliciesConfig) SetScopes(v []ScopesInner)`

SetScopes sets Scopes field to given value.


### GetSets

`func (o *AccessPoliciesConfig) GetSets() []SetsInner`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *AccessPoliciesConfig) GetSetsOk() (*[]SetsInner, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *AccessPoliciesConfig) SetSets(v []SetsInner)`

SetSets sets Sets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


