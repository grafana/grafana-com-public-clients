# PostInstanceOAuthGithubRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedOrganizations** | Pointer to **[]string** |  | [optional] 
**ClientId** | **string** |  | 
**ClientSecret** | Pointer to **string** |  | [optional] 
**TeamIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPostInstanceOAuthGithubRequest

`func NewPostInstanceOAuthGithubRequest(clientId string, ) *PostInstanceOAuthGithubRequest`

NewPostInstanceOAuthGithubRequest instantiates a new PostInstanceOAuthGithubRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostInstanceOAuthGithubRequestWithDefaults

`func NewPostInstanceOAuthGithubRequestWithDefaults() *PostInstanceOAuthGithubRequest`

NewPostInstanceOAuthGithubRequestWithDefaults instantiates a new PostInstanceOAuthGithubRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedOrganizations

`func (o *PostInstanceOAuthGithubRequest) GetAllowedOrganizations() []string`

GetAllowedOrganizations returns the AllowedOrganizations field if non-nil, zero value otherwise.

### GetAllowedOrganizationsOk

`func (o *PostInstanceOAuthGithubRequest) GetAllowedOrganizationsOk() (*[]string, bool)`

GetAllowedOrganizationsOk returns a tuple with the AllowedOrganizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrganizations

`func (o *PostInstanceOAuthGithubRequest) SetAllowedOrganizations(v []string)`

SetAllowedOrganizations sets AllowedOrganizations field to given value.

### HasAllowedOrganizations

`func (o *PostInstanceOAuthGithubRequest) HasAllowedOrganizations() bool`

HasAllowedOrganizations returns a boolean if a field has been set.

### GetClientId

`func (o *PostInstanceOAuthGithubRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PostInstanceOAuthGithubRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PostInstanceOAuthGithubRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *PostInstanceOAuthGithubRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PostInstanceOAuthGithubRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PostInstanceOAuthGithubRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PostInstanceOAuthGithubRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetTeamIds

`func (o *PostInstanceOAuthGithubRequest) GetTeamIds() []string`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *PostInstanceOAuthGithubRequest) GetTeamIdsOk() (*[]string, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *PostInstanceOAuthGithubRequest) SetTeamIds(v []string)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *PostInstanceOAuthGithubRequest) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


