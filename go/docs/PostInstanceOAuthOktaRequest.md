# PostInstanceOAuthOktaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedDomains** | Pointer to **[]string** |  | [optional] 
**AllowedGroups** | Pointer to **[]string** |  | [optional] 
**ApiUrl** | **string** |  | 
**AuthUrl** | **string** |  | 
**ClientId** | **string** |  | 
**ClientSecret** | Pointer to **string** |  | [optional] 
**RoleAttributePath** | Pointer to **string** |  | [optional] 
**TokenUrl** | **string** |  | 

## Methods

### NewPostInstanceOAuthOktaRequest

`func NewPostInstanceOAuthOktaRequest(apiUrl string, authUrl string, clientId string, tokenUrl string, ) *PostInstanceOAuthOktaRequest`

NewPostInstanceOAuthOktaRequest instantiates a new PostInstanceOAuthOktaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostInstanceOAuthOktaRequestWithDefaults

`func NewPostInstanceOAuthOktaRequestWithDefaults() *PostInstanceOAuthOktaRequest`

NewPostInstanceOAuthOktaRequestWithDefaults instantiates a new PostInstanceOAuthOktaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedDomains

`func (o *PostInstanceOAuthOktaRequest) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *PostInstanceOAuthOktaRequest) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *PostInstanceOAuthOktaRequest) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *PostInstanceOAuthOktaRequest) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetAllowedGroups

`func (o *PostInstanceOAuthOktaRequest) GetAllowedGroups() []string`

GetAllowedGroups returns the AllowedGroups field if non-nil, zero value otherwise.

### GetAllowedGroupsOk

`func (o *PostInstanceOAuthOktaRequest) GetAllowedGroupsOk() (*[]string, bool)`

GetAllowedGroupsOk returns a tuple with the AllowedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedGroups

`func (o *PostInstanceOAuthOktaRequest) SetAllowedGroups(v []string)`

SetAllowedGroups sets AllowedGroups field to given value.

### HasAllowedGroups

`func (o *PostInstanceOAuthOktaRequest) HasAllowedGroups() bool`

HasAllowedGroups returns a boolean if a field has been set.

### GetApiUrl

`func (o *PostInstanceOAuthOktaRequest) GetApiUrl() string`

GetApiUrl returns the ApiUrl field if non-nil, zero value otherwise.

### GetApiUrlOk

`func (o *PostInstanceOAuthOktaRequest) GetApiUrlOk() (*string, bool)`

GetApiUrlOk returns a tuple with the ApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiUrl

`func (o *PostInstanceOAuthOktaRequest) SetApiUrl(v string)`

SetApiUrl sets ApiUrl field to given value.


### GetAuthUrl

`func (o *PostInstanceOAuthOktaRequest) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *PostInstanceOAuthOktaRequest) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *PostInstanceOAuthOktaRequest) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.


### GetClientId

`func (o *PostInstanceOAuthOktaRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PostInstanceOAuthOktaRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PostInstanceOAuthOktaRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *PostInstanceOAuthOktaRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PostInstanceOAuthOktaRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PostInstanceOAuthOktaRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PostInstanceOAuthOktaRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetRoleAttributePath

`func (o *PostInstanceOAuthOktaRequest) GetRoleAttributePath() string`

GetRoleAttributePath returns the RoleAttributePath field if non-nil, zero value otherwise.

### GetRoleAttributePathOk

`func (o *PostInstanceOAuthOktaRequest) GetRoleAttributePathOk() (*string, bool)`

GetRoleAttributePathOk returns a tuple with the RoleAttributePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributePath

`func (o *PostInstanceOAuthOktaRequest) SetRoleAttributePath(v string)`

SetRoleAttributePath sets RoleAttributePath field to given value.

### HasRoleAttributePath

`func (o *PostInstanceOAuthOktaRequest) HasRoleAttributePath() bool`

HasRoleAttributePath returns a boolean if a field has been set.

### GetTokenUrl

`func (o *PostInstanceOAuthOktaRequest) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *PostInstanceOAuthOktaRequest) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *PostInstanceOAuthOktaRequest) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


