# GrafanaTokenDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Expiration** | Pointer to **time.Time** |  | [optional] 
**HasExpired** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**IsRevoked** | Pointer to **bool** |  | [optional] 
**LastUsedAt** | Pointer to **time.Time** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SecondsUntilExpiration** | Pointer to **float64** |  | [optional] 

## Methods

### NewGrafanaTokenDTO

`func NewGrafanaTokenDTO() *GrafanaTokenDTO`

NewGrafanaTokenDTO instantiates a new GrafanaTokenDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrafanaTokenDTOWithDefaults

`func NewGrafanaTokenDTOWithDefaults() *GrafanaTokenDTO`

NewGrafanaTokenDTOWithDefaults instantiates a new GrafanaTokenDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GrafanaTokenDTO) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GrafanaTokenDTO) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GrafanaTokenDTO) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GrafanaTokenDTO) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpiration

`func (o *GrafanaTokenDTO) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *GrafanaTokenDTO) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *GrafanaTokenDTO) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *GrafanaTokenDTO) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetHasExpired

`func (o *GrafanaTokenDTO) GetHasExpired() bool`

GetHasExpired returns the HasExpired field if non-nil, zero value otherwise.

### GetHasExpiredOk

`func (o *GrafanaTokenDTO) GetHasExpiredOk() (*bool, bool)`

GetHasExpiredOk returns a tuple with the HasExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasExpired

`func (o *GrafanaTokenDTO) SetHasExpired(v bool)`

SetHasExpired sets HasExpired field to given value.

### HasHasExpired

`func (o *GrafanaTokenDTO) HasHasExpired() bool`

HasHasExpired returns a boolean if a field has been set.

### GetId

`func (o *GrafanaTokenDTO) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GrafanaTokenDTO) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GrafanaTokenDTO) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GrafanaTokenDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRevoked

`func (o *GrafanaTokenDTO) GetIsRevoked() bool`

GetIsRevoked returns the IsRevoked field if non-nil, zero value otherwise.

### GetIsRevokedOk

`func (o *GrafanaTokenDTO) GetIsRevokedOk() (*bool, bool)`

GetIsRevokedOk returns a tuple with the IsRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRevoked

`func (o *GrafanaTokenDTO) SetIsRevoked(v bool)`

SetIsRevoked sets IsRevoked field to given value.

### HasIsRevoked

`func (o *GrafanaTokenDTO) HasIsRevoked() bool`

HasIsRevoked returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *GrafanaTokenDTO) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *GrafanaTokenDTO) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *GrafanaTokenDTO) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *GrafanaTokenDTO) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetName

`func (o *GrafanaTokenDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GrafanaTokenDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GrafanaTokenDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GrafanaTokenDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecondsUntilExpiration

`func (o *GrafanaTokenDTO) GetSecondsUntilExpiration() float64`

GetSecondsUntilExpiration returns the SecondsUntilExpiration field if non-nil, zero value otherwise.

### GetSecondsUntilExpirationOk

`func (o *GrafanaTokenDTO) GetSecondsUntilExpirationOk() (*float64, bool)`

GetSecondsUntilExpirationOk returns a tuple with the SecondsUntilExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsUntilExpiration

`func (o *GrafanaTokenDTO) SetSecondsUntilExpiration(v float64)`

SetSecondsUntilExpiration sets SecondsUntilExpiration field to given value.

### HasSecondsUntilExpiration

`func (o *GrafanaTokenDTO) HasSecondsUntilExpiration() bool`

HasSecondsUntilExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


