# Params

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | Pointer to **string** | OAuth grant type | [optional] 
**ClientId** | Pointer to **string** | OAuth client ID | [optional] 
**ClientSecret** | Pointer to **string** | OAuth client secret. Usually populated by decrypting the secret from the SecureJson blob. | [optional] 
**Resource** | Pointer to **string** | OAuth resource | [optional] 

## Methods

### NewParams

`func NewParams() *Params`

NewParams instantiates a new Params object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParamsWithDefaults

`func NewParamsWithDefaults() *Params`

NewParamsWithDefaults instantiates a new Params object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *Params) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *Params) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *Params) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *Params) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetClientId

`func (o *Params) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Params) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Params) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Params) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *Params) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Params) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Params) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *Params) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetResource

`func (o *Params) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *Params) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *Params) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *Params) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


