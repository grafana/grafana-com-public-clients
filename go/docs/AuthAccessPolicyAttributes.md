# AuthAccessPolicyAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LokiQueryPolicy** | Pointer to [**AuthAccessPolicyAttributesLokiQueryPolicy**](AuthAccessPolicyAttributesLokiQueryPolicy.md) |  | [optional] 
**AllowedScopes** | Pointer to **[]string** | List of scopes allowed to be signed by an access policy (required if the access policy contains &#x60;grafana-id-token:sign&#x60;).  | [optional] 
**PdcConfiguration** | Pointer to [**AuthAccessPolicyAttributesPdcConfiguration**](AuthAccessPolicyAttributesPdcConfiguration.md) |  | [optional] 
**AllowedAudiences** | Pointer to **[]string** | List of audience claims allowed to be included when signing access tokens.  | [optional] 

## Methods

### NewAuthAccessPolicyAttributes

`func NewAuthAccessPolicyAttributes() *AuthAccessPolicyAttributes`

NewAuthAccessPolicyAttributes instantiates a new AuthAccessPolicyAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAccessPolicyAttributesWithDefaults

`func NewAuthAccessPolicyAttributesWithDefaults() *AuthAccessPolicyAttributes`

NewAuthAccessPolicyAttributesWithDefaults instantiates a new AuthAccessPolicyAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLokiQueryPolicy

`func (o *AuthAccessPolicyAttributes) GetLokiQueryPolicy() AuthAccessPolicyAttributesLokiQueryPolicy`

GetLokiQueryPolicy returns the LokiQueryPolicy field if non-nil, zero value otherwise.

### GetLokiQueryPolicyOk

`func (o *AuthAccessPolicyAttributes) GetLokiQueryPolicyOk() (*AuthAccessPolicyAttributesLokiQueryPolicy, bool)`

GetLokiQueryPolicyOk returns a tuple with the LokiQueryPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLokiQueryPolicy

`func (o *AuthAccessPolicyAttributes) SetLokiQueryPolicy(v AuthAccessPolicyAttributesLokiQueryPolicy)`

SetLokiQueryPolicy sets LokiQueryPolicy field to given value.

### HasLokiQueryPolicy

`func (o *AuthAccessPolicyAttributes) HasLokiQueryPolicy() bool`

HasLokiQueryPolicy returns a boolean if a field has been set.

### GetAllowedScopes

`func (o *AuthAccessPolicyAttributes) GetAllowedScopes() []string`

GetAllowedScopes returns the AllowedScopes field if non-nil, zero value otherwise.

### GetAllowedScopesOk

`func (o *AuthAccessPolicyAttributes) GetAllowedScopesOk() (*[]string, bool)`

GetAllowedScopesOk returns a tuple with the AllowedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedScopes

`func (o *AuthAccessPolicyAttributes) SetAllowedScopes(v []string)`

SetAllowedScopes sets AllowedScopes field to given value.

### HasAllowedScopes

`func (o *AuthAccessPolicyAttributes) HasAllowedScopes() bool`

HasAllowedScopes returns a boolean if a field has been set.

### GetPdcConfiguration

`func (o *AuthAccessPolicyAttributes) GetPdcConfiguration() AuthAccessPolicyAttributesPdcConfiguration`

GetPdcConfiguration returns the PdcConfiguration field if non-nil, zero value otherwise.

### GetPdcConfigurationOk

`func (o *AuthAccessPolicyAttributes) GetPdcConfigurationOk() (*AuthAccessPolicyAttributesPdcConfiguration, bool)`

GetPdcConfigurationOk returns a tuple with the PdcConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdcConfiguration

`func (o *AuthAccessPolicyAttributes) SetPdcConfiguration(v AuthAccessPolicyAttributesPdcConfiguration)`

SetPdcConfiguration sets PdcConfiguration field to given value.

### HasPdcConfiguration

`func (o *AuthAccessPolicyAttributes) HasPdcConfiguration() bool`

HasPdcConfiguration returns a boolean if a field has been set.

### GetAllowedAudiences

`func (o *AuthAccessPolicyAttributes) GetAllowedAudiences() []string`

GetAllowedAudiences returns the AllowedAudiences field if non-nil, zero value otherwise.

### GetAllowedAudiencesOk

`func (o *AuthAccessPolicyAttributes) GetAllowedAudiencesOk() (*[]string, bool)`

GetAllowedAudiencesOk returns a tuple with the AllowedAudiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAudiences

`func (o *AuthAccessPolicyAttributes) SetAllowedAudiences(v []string)`

SetAllowedAudiences sets AllowedAudiences field to given value.

### HasAllowedAudiences

`func (o *AuthAccessPolicyAttributes) HasAllowedAudiences() bool`

HasAllowedAudiences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


