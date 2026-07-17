# UpdateDatasourceBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **string** |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] 
**BasicAuth** | Pointer to [**UpdateDatasourceBodyBasicAuth**](UpdateDatasourceBodyBasicAuth.md) |  | [optional] 
**BasicAuthPassword** | Pointer to **string** |  | [optional] 
**BasicAuthUser** | Pointer to **string** |  | [optional] 
**Database** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to [**UpdateDatasourceBodyBasicAuth**](UpdateDatasourceBodyBasicAuth.md) |  | [optional] 
**IsDefault** | Pointer to [**UpdateDatasourceBodyBasicAuth**](UpdateDatasourceBodyBasicAuth.md) |  | [optional] 
**JsonData** | Pointer to **map[string]interface{}** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SecureJsonData** | Pointer to **map[string]interface{}** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**WithCredentials** | Pointer to [**UpdateDatasourceBodyBasicAuth**](UpdateDatasourceBodyBasicAuth.md) |  | [optional] 

## Methods

### NewUpdateDatasourceBody

`func NewUpdateDatasourceBody() *UpdateDatasourceBody`

NewUpdateDatasourceBody instantiates a new UpdateDatasourceBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDatasourceBodyWithDefaults

`func NewUpdateDatasourceBodyWithDefaults() *UpdateDatasourceBody`

NewUpdateDatasourceBodyWithDefaults instantiates a new UpdateDatasourceBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *UpdateDatasourceBody) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateDatasourceBody) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateDatasourceBody) SetAccess(v string)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *UpdateDatasourceBody) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetAliases

`func (o *UpdateDatasourceBody) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *UpdateDatasourceBody) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *UpdateDatasourceBody) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *UpdateDatasourceBody) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetBasicAuth

`func (o *UpdateDatasourceBody) GetBasicAuth() UpdateDatasourceBodyBasicAuth`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *UpdateDatasourceBody) GetBasicAuthOk() (*UpdateDatasourceBodyBasicAuth, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *UpdateDatasourceBody) SetBasicAuth(v UpdateDatasourceBodyBasicAuth)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *UpdateDatasourceBody) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetBasicAuthPassword

`func (o *UpdateDatasourceBody) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *UpdateDatasourceBody) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *UpdateDatasourceBody) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.

### HasBasicAuthPassword

`func (o *UpdateDatasourceBody) HasBasicAuthPassword() bool`

HasBasicAuthPassword returns a boolean if a field has been set.

### GetBasicAuthUser

`func (o *UpdateDatasourceBody) GetBasicAuthUser() string`

GetBasicAuthUser returns the BasicAuthUser field if non-nil, zero value otherwise.

### GetBasicAuthUserOk

`func (o *UpdateDatasourceBody) GetBasicAuthUserOk() (*string, bool)`

GetBasicAuthUserOk returns a tuple with the BasicAuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUser

`func (o *UpdateDatasourceBody) SetBasicAuthUser(v string)`

SetBasicAuthUser sets BasicAuthUser field to given value.

### HasBasicAuthUser

`func (o *UpdateDatasourceBody) HasBasicAuthUser() bool`

HasBasicAuthUser returns a boolean if a field has been set.

### GetDatabase

`func (o *UpdateDatasourceBody) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *UpdateDatasourceBody) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *UpdateDatasourceBody) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *UpdateDatasourceBody) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetEditable

`func (o *UpdateDatasourceBody) GetEditable() UpdateDatasourceBodyBasicAuth`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *UpdateDatasourceBody) GetEditableOk() (*UpdateDatasourceBodyBasicAuth, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *UpdateDatasourceBody) SetEditable(v UpdateDatasourceBodyBasicAuth)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *UpdateDatasourceBody) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetIsDefault

`func (o *UpdateDatasourceBody) GetIsDefault() UpdateDatasourceBodyBasicAuth`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UpdateDatasourceBody) GetIsDefaultOk() (*UpdateDatasourceBodyBasicAuth, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UpdateDatasourceBody) SetIsDefault(v UpdateDatasourceBodyBasicAuth)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *UpdateDatasourceBody) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetJsonData

`func (o *UpdateDatasourceBody) GetJsonData() map[string]interface{}`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *UpdateDatasourceBody) GetJsonDataOk() (*map[string]interface{}, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *UpdateDatasourceBody) SetJsonData(v map[string]interface{})`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *UpdateDatasourceBody) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateDatasourceBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateDatasourceBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateDatasourceBody) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateDatasourceBody) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSecureJsonData

`func (o *UpdateDatasourceBody) GetSecureJsonData() map[string]interface{}`

GetSecureJsonData returns the SecureJsonData field if non-nil, zero value otherwise.

### GetSecureJsonDataOk

`func (o *UpdateDatasourceBody) GetSecureJsonDataOk() (*map[string]interface{}, bool)`

GetSecureJsonDataOk returns a tuple with the SecureJsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureJsonData

`func (o *UpdateDatasourceBody) SetSecureJsonData(v map[string]interface{})`

SetSecureJsonData sets SecureJsonData field to given value.

### HasSecureJsonData

`func (o *UpdateDatasourceBody) HasSecureJsonData() bool`

HasSecureJsonData returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateDatasourceBody) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateDatasourceBody) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateDatasourceBody) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateDatasourceBody) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUser

`func (o *UpdateDatasourceBody) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdateDatasourceBody) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdateDatasourceBody) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdateDatasourceBody) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWithCredentials

`func (o *UpdateDatasourceBody) GetWithCredentials() UpdateDatasourceBodyBasicAuth`

GetWithCredentials returns the WithCredentials field if non-nil, zero value otherwise.

### GetWithCredentialsOk

`func (o *UpdateDatasourceBody) GetWithCredentialsOk() (*UpdateDatasourceBodyBasicAuth, bool)`

GetWithCredentialsOk returns a tuple with the WithCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCredentials

`func (o *UpdateDatasourceBody) SetWithCredentials(v UpdateDatasourceBodyBasicAuth)`

SetWithCredentials sets WithCredentials field to given value.

### HasWithCredentials

`func (o *UpdateDatasourceBody) HasWithCredentials() bool`

HasWithCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


