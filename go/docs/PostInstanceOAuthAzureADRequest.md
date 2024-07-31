# PostInstanceOAuthAzureADRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedDomains** | Pointer to **[]string** |  | [optional] 
**AllowedGroups** | Pointer to **[]string** |  | [optional] 
**AuthUrl** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecret** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**TokenUrl** | **string** |  | 

## Methods

### NewPostInstanceOAuthAzureADRequest

`func NewPostInstanceOAuthAzureADRequest(authUrl string, clientId string, tokenUrl string, ) *PostInstanceOAuthAzureADRequest`

NewPostInstanceOAuthAzureADRequest instantiates a new PostInstanceOAuthAzureADRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostInstanceOAuthAzureADRequestWithDefaults

`func NewPostInstanceOAuthAzureADRequestWithDefaults() *PostInstanceOAuthAzureADRequest`

NewPostInstanceOAuthAzureADRequestWithDefaults instantiates a new PostInstanceOAuthAzureADRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedDomains

`func (o *PostInstanceOAuthAzureADRequest) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *PostInstanceOAuthAzureADRequest) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *PostInstanceOAuthAzureADRequest) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *PostInstanceOAuthAzureADRequest) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetAllowedGroups

`func (o *PostInstanceOAuthAzureADRequest) GetAllowedGroups() []string`

GetAllowedGroups returns the AllowedGroups field if non-nil, zero value otherwise.

### GetAllowedGroupsOk

`func (o *PostInstanceOAuthAzureADRequest) GetAllowedGroupsOk() (*[]string, bool)`

GetAllowedGroupsOk returns a tuple with the AllowedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGroups

`func (o *PostInstanceOAuthAzureADRequest) SetAllowedGroups(v []string)`

SetAllowedGroups sets AllowedGroups field to given value.

### HasAllowedGroups

`func (o *PostInstanceOAuthAzureADRequest) HasAllowedGroups() bool`

HasAllowedGroups returns a boolean if a field has been set.

### GetAuthUrl

`func (o *PostInstanceOAuthAzureADRequest) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *PostInstanceOAuthAzureADRequest) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *PostInstanceOAuthAzureADRequest) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.


### GetClientId

`func (o *PostInstanceOAuthAzureADRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PostInstanceOAuthAzureADRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PostInstanceOAuthAzureADRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *PostInstanceOAuthAzureADRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PostInstanceOAuthAzureADRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PostInstanceOAuthAzureADRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PostInstanceOAuthAzureADRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetScopes

`func (o *PostInstanceOAuthAzureADRequest) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PostInstanceOAuthAzureADRequest) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PostInstanceOAuthAzureADRequest) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PostInstanceOAuthAzureADRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetTokenUrl

`func (o *PostInstanceOAuthAzureADRequest) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *PostInstanceOAuthAzureADRequest) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *PostInstanceOAuthAzureADRequest) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


