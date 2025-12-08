# JwtTokenAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL to fetch the JWT token. | [optional] 
**Scopes** | Pointer to **[]string** | The list of scopes that your application should be granted access to. | [optional] 
**Params** | Pointer to [**Params1**](Params1.md) |  | [optional] 

## Methods

### NewJwtTokenAuth

`func NewJwtTokenAuth() *JwtTokenAuth`

NewJwtTokenAuth instantiates a new JwtTokenAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJwtTokenAuthWithDefaults

`func NewJwtTokenAuthWithDefaults() *JwtTokenAuth`

NewJwtTokenAuthWithDefaults instantiates a new JwtTokenAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *JwtTokenAuth) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *JwtTokenAuth) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *JwtTokenAuth) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *JwtTokenAuth) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetScopes

`func (o *JwtTokenAuth) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *JwtTokenAuth) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *JwtTokenAuth) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *JwtTokenAuth) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetParams

`func (o *JwtTokenAuth) GetParams() Params1`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *JwtTokenAuth) GetParamsOk() (*Params1, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *JwtTokenAuth) SetParams(v Params1)`

SetParams sets Params field to given value.

### HasParams

`func (o *JwtTokenAuth) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


