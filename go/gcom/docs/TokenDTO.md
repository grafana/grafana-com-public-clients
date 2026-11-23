# TokenDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **NullableTime** | when the token was created (RFC3339) | 
**Expiration** | **NullableTime** | when the token expires (RFC3339) | 
**HasExpired** | **bool** | whether the token has expired | 
**Id** | **int64** | service account token id | 
**IsRevoked** | **NullableBool** | whether the token has been revoked | 
**LastUsedAt** | **NullableTime** | when the token was last used (RFC3339) | 
**Name** | **string** | service account token name | 
**SecondsUntilExpiration** | **NullableFloat64** | seconds until the token expires | 

## Methods

### NewTokenDTO

`func NewTokenDTO(created NullableTime, expiration NullableTime, hasExpired bool, id int64, isRevoked NullableBool, lastUsedAt NullableTime, name string, secondsUntilExpiration NullableFloat64, ) *TokenDTO`

NewTokenDTO instantiates a new TokenDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenDTOWithDefaults

`func NewTokenDTOWithDefaults() *TokenDTO`

NewTokenDTOWithDefaults instantiates a new TokenDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *TokenDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TokenDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TokenDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### SetCreatedNil

`func (o *TokenDTO) SetCreatedNil(b bool)`

 SetCreatedNil sets the value for Created to be an explicit nil

### UnsetCreated
`func (o *TokenDTO) UnsetCreated()`

UnsetCreated ensures that no value is present for Created, not even an explicit nil
### GetExpiration

`func (o *TokenDTO) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *TokenDTO) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *TokenDTO) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.


### SetExpirationNil

`func (o *TokenDTO) SetExpirationNil(b bool)`

 SetExpirationNil sets the value for Expiration to be an explicit nil

### UnsetExpiration
`func (o *TokenDTO) UnsetExpiration()`

UnsetExpiration ensures that no value is present for Expiration, not even an explicit nil
### GetHasExpired

`func (o *TokenDTO) GetHasExpired() bool`

GetHasExpired returns the HasExpired field if non-nil, zero value otherwise.

### GetHasExpiredOk

`func (o *TokenDTO) GetHasExpiredOk() (*bool, bool)`

GetHasExpiredOk returns a tuple with the HasExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExpired

`func (o *TokenDTO) SetHasExpired(v bool)`

SetHasExpired sets HasExpired field to given value.


### GetId

`func (o *TokenDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenDTO) SetId(v int64)`

SetId sets Id field to given value.


### GetIsRevoked

`func (o *TokenDTO) GetIsRevoked() bool`

GetIsRevoked returns the IsRevoked field if non-nil, zero value otherwise.

### GetIsRevokedOk

`func (o *TokenDTO) GetIsRevokedOk() (*bool, bool)`

GetIsRevokedOk returns a tuple with the IsRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRevoked

`func (o *TokenDTO) SetIsRevoked(v bool)`

SetIsRevoked sets IsRevoked field to given value.


### SetIsRevokedNil

`func (o *TokenDTO) SetIsRevokedNil(b bool)`

 SetIsRevokedNil sets the value for IsRevoked to be an explicit nil

### UnsetIsRevoked
`func (o *TokenDTO) UnsetIsRevoked()`

UnsetIsRevoked ensures that no value is present for IsRevoked, not even an explicit nil
### GetLastUsedAt

`func (o *TokenDTO) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *TokenDTO) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *TokenDTO) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.


### SetLastUsedAtNil

`func (o *TokenDTO) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *TokenDTO) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetName

`func (o *TokenDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenDTO) SetName(v string)`

SetName sets Name field to given value.


### GetSecondsUntilExpiration

`func (o *TokenDTO) GetSecondsUntilExpiration() float64`

GetSecondsUntilExpiration returns the SecondsUntilExpiration field if non-nil, zero value otherwise.

### GetSecondsUntilExpirationOk

`func (o *TokenDTO) GetSecondsUntilExpirationOk() (*float64, bool)`

GetSecondsUntilExpirationOk returns a tuple with the SecondsUntilExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsUntilExpiration

`func (o *TokenDTO) SetSecondsUntilExpiration(v float64)`

SetSecondsUntilExpiration sets SecondsUntilExpiration field to given value.


### SetSecondsUntilExpirationNil

`func (o *TokenDTO) SetSecondsUntilExpirationNil(b bool)`

 SetSecondsUntilExpirationNil sets the value for SecondsUntilExpiration to be an explicit nil

### UnsetSecondsUntilExpiration
`func (o *TokenDTO) UnsetSecondsUntilExpiration()`

UnsetSecondsUntilExpiration ensures that no value is present for SecondsUntilExpiration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


