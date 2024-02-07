# PostInstanceOAuthGoogleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedDomains** | **[]string** |  | 
**ClientId** | **string** |  | 
**ClientSecret** | Pointer to **string** |  | [optional] 

## Methods

### NewPostInstanceOAuthGoogleRequest

`func NewPostInstanceOAuthGoogleRequest(allowedDomains []string, clientId string, ) *PostInstanceOAuthGoogleRequest`

NewPostInstanceOAuthGoogleRequest instantiates a new PostInstanceOAuthGoogleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostInstanceOAuthGoogleRequestWithDefaults

`func NewPostInstanceOAuthGoogleRequestWithDefaults() *PostInstanceOAuthGoogleRequest`

NewPostInstanceOAuthGoogleRequestWithDefaults instantiates a new PostInstanceOAuthGoogleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedDomains

`func (o *PostInstanceOAuthGoogleRequest) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *PostInstanceOAuthGoogleRequest) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *PostInstanceOAuthGoogleRequest) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.


### GetClientId

`func (o *PostInstanceOAuthGoogleRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PostInstanceOAuthGoogleRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PostInstanceOAuthGoogleRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *PostInstanceOAuthGoogleRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *PostInstanceOAuthGoogleRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *PostInstanceOAuthGoogleRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *PostInstanceOAuthGoogleRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


