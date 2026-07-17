# CreateDatasourceRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **string** | datasource access mode | 
**BasicAuth** | Pointer to **int32** |  | [optional] 
**BasicAuthPassword** | Pointer to **string** |  | [optional] 
**BasicAuthUser** | Pointer to **string** |  | [optional] 
**Database** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **int32** |  | [optional] 
**GrafanaOrgId** | Pointer to **int32** | grafana org id | [optional] [default to 1]
**IsDefault** | Pointer to **int32** |  | [optional] 
**JsonData** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** | datasource name | 
**Password** | Pointer to **string** |  | [optional] 
**SecureJsonData** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | **string** | datasource type | 
**Url** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**WithCredentials** | Pointer to **int32** |  | [optional] 

## Methods

### NewCreateDatasourceRequestBody

`func NewCreateDatasourceRequestBody(access string, name string, type_ string, ) *CreateDatasourceRequestBody`

NewCreateDatasourceRequestBody instantiates a new CreateDatasourceRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatasourceRequestBodyWithDefaults

`func NewCreateDatasourceRequestBodyWithDefaults() *CreateDatasourceRequestBody`

NewCreateDatasourceRequestBodyWithDefaults instantiates a new CreateDatasourceRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *CreateDatasourceRequestBody) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *CreateDatasourceRequestBody) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *CreateDatasourceRequestBody) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetBasicAuth

`func (o *CreateDatasourceRequestBody) GetBasicAuth() int32`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *CreateDatasourceRequestBody) GetBasicAuthOk() (*int32, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *CreateDatasourceRequestBody) SetBasicAuth(v int32)`

SetBasicAuth sets BasicAuth field to given value.

### HasBasicAuth

`func (o *CreateDatasourceRequestBody) HasBasicAuth() bool`

HasBasicAuth returns a boolean if a field has been set.

### GetBasicAuthPassword

`func (o *CreateDatasourceRequestBody) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *CreateDatasourceRequestBody) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *CreateDatasourceRequestBody) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.

### HasBasicAuthPassword

`func (o *CreateDatasourceRequestBody) HasBasicAuthPassword() bool`

HasBasicAuthPassword returns a boolean if a field has been set.

### GetBasicAuthUser

`func (o *CreateDatasourceRequestBody) GetBasicAuthUser() string`

GetBasicAuthUser returns the BasicAuthUser field if non-nil, zero value otherwise.

### GetBasicAuthUserOk

`func (o *CreateDatasourceRequestBody) GetBasicAuthUserOk() (*string, bool)`

GetBasicAuthUserOk returns a tuple with the BasicAuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUser

`func (o *CreateDatasourceRequestBody) SetBasicAuthUser(v string)`

SetBasicAuthUser sets BasicAuthUser field to given value.

### HasBasicAuthUser

`func (o *CreateDatasourceRequestBody) HasBasicAuthUser() bool`

HasBasicAuthUser returns a boolean if a field has been set.

### GetDatabase

`func (o *CreateDatasourceRequestBody) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *CreateDatasourceRequestBody) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *CreateDatasourceRequestBody) SetDatabase(v string)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *CreateDatasourceRequestBody) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetEditable

`func (o *CreateDatasourceRequestBody) GetEditable() int32`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *CreateDatasourceRequestBody) GetEditableOk() (*int32, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *CreateDatasourceRequestBody) SetEditable(v int32)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *CreateDatasourceRequestBody) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetGrafanaOrgId

`func (o *CreateDatasourceRequestBody) GetGrafanaOrgId() int32`

GetGrafanaOrgId returns the GrafanaOrgId field if non-nil, zero value otherwise.

### GetGrafanaOrgIdOk

`func (o *CreateDatasourceRequestBody) GetGrafanaOrgIdOk() (*int32, bool)`

GetGrafanaOrgIdOk returns a tuple with the GrafanaOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaOrgId

`func (o *CreateDatasourceRequestBody) SetGrafanaOrgId(v int32)`

SetGrafanaOrgId sets GrafanaOrgId field to given value.

### HasGrafanaOrgId

`func (o *CreateDatasourceRequestBody) HasGrafanaOrgId() bool`

HasGrafanaOrgId returns a boolean if a field has been set.

### GetIsDefault

`func (o *CreateDatasourceRequestBody) GetIsDefault() int32`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreateDatasourceRequestBody) GetIsDefaultOk() (*int32, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreateDatasourceRequestBody) SetIsDefault(v int32)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CreateDatasourceRequestBody) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetJsonData

`func (o *CreateDatasourceRequestBody) GetJsonData() map[string]interface{}`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *CreateDatasourceRequestBody) GetJsonDataOk() (*map[string]interface{}, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *CreateDatasourceRequestBody) SetJsonData(v map[string]interface{})`

SetJsonData sets JsonData field to given value.

### HasJsonData

`func (o *CreateDatasourceRequestBody) HasJsonData() bool`

HasJsonData returns a boolean if a field has been set.

### GetName

`func (o *CreateDatasourceRequestBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDatasourceRequestBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDatasourceRequestBody) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateDatasourceRequestBody) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateDatasourceRequestBody) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateDatasourceRequestBody) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateDatasourceRequestBody) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSecureJsonData

`func (o *CreateDatasourceRequestBody) GetSecureJsonData() map[string]interface{}`

GetSecureJsonData returns the SecureJsonData field if non-nil, zero value otherwise.

### GetSecureJsonDataOk

`func (o *CreateDatasourceRequestBody) GetSecureJsonDataOk() (*map[string]interface{}, bool)`

GetSecureJsonDataOk returns a tuple with the SecureJsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureJsonData

`func (o *CreateDatasourceRequestBody) SetSecureJsonData(v map[string]interface{})`

SetSecureJsonData sets SecureJsonData field to given value.

### HasSecureJsonData

`func (o *CreateDatasourceRequestBody) HasSecureJsonData() bool`

HasSecureJsonData returns a boolean if a field has been set.

### GetType

`func (o *CreateDatasourceRequestBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDatasourceRequestBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDatasourceRequestBody) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *CreateDatasourceRequestBody) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateDatasourceRequestBody) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateDatasourceRequestBody) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateDatasourceRequestBody) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUser

`func (o *CreateDatasourceRequestBody) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CreateDatasourceRequestBody) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CreateDatasourceRequestBody) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CreateDatasourceRequestBody) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWithCredentials

`func (o *CreateDatasourceRequestBody) GetWithCredentials() int32`

GetWithCredentials returns the WithCredentials field if non-nil, zero value otherwise.

### GetWithCredentialsOk

`func (o *CreateDatasourceRequestBody) GetWithCredentialsOk() (*int32, bool)`

GetWithCredentialsOk returns a tuple with the WithCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCredentials

`func (o *CreateDatasourceRequestBody) SetWithCredentials(v int32)`

SetWithCredentials sets WithCredentials field to given value.

### HasWithCredentials

`func (o *CreateDatasourceRequestBody) HasWithCredentials() bool`

HasWithCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


