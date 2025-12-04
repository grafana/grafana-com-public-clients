# TokenAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL to fetch the authentication token. | [optional] 
**Scopes** | Pointer to **[]string** | The list of scopes that your application should be granted access to. | [optional] 
**Params** | Pointer to [**Params**](Params.md) |  | [optional] 

## Methods

### NewTokenAuth

`func NewTokenAuth() *TokenAuth`

NewTokenAuth instantiates a new TokenAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenAuthWithDefaults

`func NewTokenAuthWithDefaults() *TokenAuth`

NewTokenAuthWithDefaults instantiates a new TokenAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *TokenAuth) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TokenAuth) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TokenAuth) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *TokenAuth) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetScopes

`func (o *TokenAuth) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenAuth) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenAuth) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenAuth) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetParams

`func (o *TokenAuth) GetParams() Params`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *TokenAuth) GetParamsOk() (*Params, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *TokenAuth) SetParams(v Params)`

SetParams sets Params field to given value.

### HasParams

`func (o *TokenAuth) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


