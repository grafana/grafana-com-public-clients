# Datasource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **string** |  | 
**BasicAuth** | **int32** |  | 
**BasicAuthPassword** | **string** |  | 
**BasicAuthUser** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Database** | **string** |  | 
**Delete** | **int32** |  | 
**Editable** | **int32** |  | 
**GrafanaOrgId** | **int32** |  | 
**Id** | **int32** |  | 
**InstanceId** | **int32** |  | 
**IsDefault** | **int32** |  | 
**JsonData** | **interface{}** |  | 
**Name** | **string** |  | 
**Password** | **string** |  | 
**SecureJsonData** | **map[string]interface{}** |  | 
**Type** | **string** |  | 
**UpdatedAt** | **NullableTime** |  | 
**Url** | **string** |  | 
**User** | **string** |  | 
**Version** | **int32** |  | 
**WithCredentials** | **int32** |  | 

## Methods

### NewDatasource

`func NewDatasource(access string, basicAuth int32, basicAuthPassword string, basicAuthUser string, createdAt time.Time, database string, delete int32, editable int32, grafanaOrgId int32, id int32, instanceId int32, isDefault int32, jsonData interface{}, name string, password string, secureJsonData map[string]interface{}, type_ string, updatedAt NullableTime, url string, user string, version int32, withCredentials int32, ) *Datasource`

NewDatasource instantiates a new Datasource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceWithDefaults

`func NewDatasourceWithDefaults() *Datasource`

NewDatasourceWithDefaults instantiates a new Datasource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *Datasource) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *Datasource) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *Datasource) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetBasicAuth

`func (o *Datasource) GetBasicAuth() int32`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *Datasource) GetBasicAuthOk() (*int32, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *Datasource) SetBasicAuth(v int32)`

SetBasicAuth sets BasicAuth field to given value.


### GetBasicAuthPassword

`func (o *Datasource) GetBasicAuthPassword() string`

GetBasicAuthPassword returns the BasicAuthPassword field if non-nil, zero value otherwise.

### GetBasicAuthPasswordOk

`func (o *Datasource) GetBasicAuthPasswordOk() (*string, bool)`

GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthPassword

`func (o *Datasource) SetBasicAuthPassword(v string)`

SetBasicAuthPassword sets BasicAuthPassword field to given value.


### GetBasicAuthUser

`func (o *Datasource) GetBasicAuthUser() string`

GetBasicAuthUser returns the BasicAuthUser field if non-nil, zero value otherwise.

### GetBasicAuthUserOk

`func (o *Datasource) GetBasicAuthUserOk() (*string, bool)`

GetBasicAuthUserOk returns a tuple with the BasicAuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUser

`func (o *Datasource) SetBasicAuthUser(v string)`

SetBasicAuthUser sets BasicAuthUser field to given value.


### GetCreatedAt

`func (o *Datasource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Datasource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Datasource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDatabase

`func (o *Datasource) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *Datasource) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *Datasource) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### GetDelete

`func (o *Datasource) GetDelete() int32`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *Datasource) GetDeleteOk() (*int32, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *Datasource) SetDelete(v int32)`

SetDelete sets Delete field to given value.


### GetEditable

`func (o *Datasource) GetEditable() int32`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *Datasource) GetEditableOk() (*int32, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *Datasource) SetEditable(v int32)`

SetEditable sets Editable field to given value.


### GetGrafanaOrgId

`func (o *Datasource) GetGrafanaOrgId() int32`

GetGrafanaOrgId returns the GrafanaOrgId field if non-nil, zero value otherwise.

### GetGrafanaOrgIdOk

`func (o *Datasource) GetGrafanaOrgIdOk() (*int32, bool)`

GetGrafanaOrgIdOk returns a tuple with the GrafanaOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaOrgId

`func (o *Datasource) SetGrafanaOrgId(v int32)`

SetGrafanaOrgId sets GrafanaOrgId field to given value.


### GetId

`func (o *Datasource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Datasource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Datasource) SetId(v int32)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *Datasource) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *Datasource) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *Datasource) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.


### GetIsDefault

`func (o *Datasource) GetIsDefault() int32`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Datasource) GetIsDefaultOk() (*int32, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Datasource) SetIsDefault(v int32)`

SetIsDefault sets IsDefault field to given value.


### GetJsonData

`func (o *Datasource) GetJsonData() interface{}`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *Datasource) GetJsonDataOk() (*interface{}, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *Datasource) SetJsonData(v interface{})`

SetJsonData sets JsonData field to given value.


### SetJsonDataNil

`func (o *Datasource) SetJsonDataNil(b bool)`

 SetJsonDataNil sets the value for JsonData to be an explicit nil

### UnsetJsonData
`func (o *Datasource) UnsetJsonData()`

UnsetJsonData ensures that no value is present for JsonData, not even an explicit nil
### GetName

`func (o *Datasource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Datasource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Datasource) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *Datasource) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Datasource) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Datasource) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSecureJsonData

`func (o *Datasource) GetSecureJsonData() map[string]interface{}`

GetSecureJsonData returns the SecureJsonData field if non-nil, zero value otherwise.

### GetSecureJsonDataOk

`func (o *Datasource) GetSecureJsonDataOk() (*map[string]interface{}, bool)`

GetSecureJsonDataOk returns a tuple with the SecureJsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureJsonData

`func (o *Datasource) SetSecureJsonData(v map[string]interface{})`

SetSecureJsonData sets SecureJsonData field to given value.


### GetType

`func (o *Datasource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Datasource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Datasource) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *Datasource) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Datasource) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Datasource) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Datasource) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Datasource) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUrl

`func (o *Datasource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Datasource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Datasource) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUser

`func (o *Datasource) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Datasource) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Datasource) SetUser(v string)`

SetUser sets User field to given value.


### GetVersion

`func (o *Datasource) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Datasource) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Datasource) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetWithCredentials

`func (o *Datasource) GetWithCredentials() int32`

GetWithCredentials returns the WithCredentials field if non-nil, zero value otherwise.

### GetWithCredentialsOk

`func (o *Datasource) GetWithCredentialsOk() (*int32, bool)`

GetWithCredentialsOk returns a tuple with the WithCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCredentials

`func (o *Datasource) SetWithCredentials(v int32)`

SetWithCredentials sets WithCredentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


