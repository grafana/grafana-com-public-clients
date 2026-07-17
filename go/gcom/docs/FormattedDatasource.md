# FormattedDatasource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **string** |  | 
**BasicAuth** | **int32** |  | 
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
**Type** | **string** |  | 
**UpdatedAt** | **NullableTime** |  | 
**Url** | **string** |  | 
**User** | **string** |  | 
**Version** | **int32** |  | 
**WithCredentials** | **int32** |  | 

## Methods

### NewFormattedDatasource

`func NewFormattedDatasource(access string, basicAuth int32, basicAuthUser string, createdAt time.Time, database string, delete int32, editable int32, grafanaOrgId int32, id int32, instanceId int32, isDefault int32, jsonData interface{}, name string, type_ string, updatedAt NullableTime, url string, user string, version int32, withCredentials int32, ) *FormattedDatasource`

NewFormattedDatasource instantiates a new FormattedDatasource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedDatasourceWithDefaults

`func NewFormattedDatasourceWithDefaults() *FormattedDatasource`

NewFormattedDatasourceWithDefaults instantiates a new FormattedDatasource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *FormattedDatasource) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *FormattedDatasource) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *FormattedDatasource) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetBasicAuth

`func (o *FormattedDatasource) GetBasicAuth() int32`

GetBasicAuth returns the BasicAuth field if non-nil, zero value otherwise.

### GetBasicAuthOk

`func (o *FormattedDatasource) GetBasicAuthOk() (*int32, bool)`

GetBasicAuthOk returns a tuple with the BasicAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuth

`func (o *FormattedDatasource) SetBasicAuth(v int32)`

SetBasicAuth sets BasicAuth field to given value.


### GetBasicAuthUser

`func (o *FormattedDatasource) GetBasicAuthUser() string`

GetBasicAuthUser returns the BasicAuthUser field if non-nil, zero value otherwise.

### GetBasicAuthUserOk

`func (o *FormattedDatasource) GetBasicAuthUserOk() (*string, bool)`

GetBasicAuthUserOk returns a tuple with the BasicAuthUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasicAuthUser

`func (o *FormattedDatasource) SetBasicAuthUser(v string)`

SetBasicAuthUser sets BasicAuthUser field to given value.


### GetCreatedAt

`func (o *FormattedDatasource) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FormattedDatasource) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FormattedDatasource) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDatabase

`func (o *FormattedDatasource) GetDatabase() string`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *FormattedDatasource) GetDatabaseOk() (*string, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *FormattedDatasource) SetDatabase(v string)`

SetDatabase sets Database field to given value.


### GetDelete

`func (o *FormattedDatasource) GetDelete() int32`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *FormattedDatasource) GetDeleteOk() (*int32, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *FormattedDatasource) SetDelete(v int32)`

SetDelete sets Delete field to given value.


### GetEditable

`func (o *FormattedDatasource) GetEditable() int32`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *FormattedDatasource) GetEditableOk() (*int32, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *FormattedDatasource) SetEditable(v int32)`

SetEditable sets Editable field to given value.


### GetGrafanaOrgId

`func (o *FormattedDatasource) GetGrafanaOrgId() int32`

GetGrafanaOrgId returns the GrafanaOrgId field if non-nil, zero value otherwise.

### GetGrafanaOrgIdOk

`func (o *FormattedDatasource) GetGrafanaOrgIdOk() (*int32, bool)`

GetGrafanaOrgIdOk returns a tuple with the GrafanaOrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaOrgId

`func (o *FormattedDatasource) SetGrafanaOrgId(v int32)`

SetGrafanaOrgId sets GrafanaOrgId field to given value.


### GetId

`func (o *FormattedDatasource) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormattedDatasource) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormattedDatasource) SetId(v int32)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *FormattedDatasource) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *FormattedDatasource) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *FormattedDatasource) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.


### GetIsDefault

`func (o *FormattedDatasource) GetIsDefault() int32`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *FormattedDatasource) GetIsDefaultOk() (*int32, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *FormattedDatasource) SetIsDefault(v int32)`

SetIsDefault sets IsDefault field to given value.


### GetJsonData

`func (o *FormattedDatasource) GetJsonData() interface{}`

GetJsonData returns the JsonData field if non-nil, zero value otherwise.

### GetJsonDataOk

`func (o *FormattedDatasource) GetJsonDataOk() (*interface{}, bool)`

GetJsonDataOk returns a tuple with the JsonData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonData

`func (o *FormattedDatasource) SetJsonData(v interface{})`

SetJsonData sets JsonData field to given value.


### SetJsonDataNil

`func (o *FormattedDatasource) SetJsonDataNil(b bool)`

 SetJsonDataNil sets the value for JsonData to be an explicit nil

### UnsetJsonData
`func (o *FormattedDatasource) UnsetJsonData()`

UnsetJsonData ensures that no value is present for JsonData, not even an explicit nil
### GetName

`func (o *FormattedDatasource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FormattedDatasource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FormattedDatasource) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FormattedDatasource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormattedDatasource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormattedDatasource) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *FormattedDatasource) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FormattedDatasource) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FormattedDatasource) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *FormattedDatasource) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *FormattedDatasource) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUrl

`func (o *FormattedDatasource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FormattedDatasource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FormattedDatasource) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUser

`func (o *FormattedDatasource) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FormattedDatasource) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FormattedDatasource) SetUser(v string)`

SetUser sets User field to given value.


### GetVersion

`func (o *FormattedDatasource) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FormattedDatasource) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FormattedDatasource) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetWithCredentials

`func (o *FormattedDatasource) GetWithCredentials() int32`

GetWithCredentials returns the WithCredentials field if non-nil, zero value otherwise.

### GetWithCredentialsOk

`func (o *FormattedDatasource) GetWithCredentialsOk() (*int32, bool)`

GetWithCredentialsOk returns a tuple with the WithCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithCredentials

`func (o *FormattedDatasource) SetWithCredentials(v int32)`

SetWithCredentials sets WithCredentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


