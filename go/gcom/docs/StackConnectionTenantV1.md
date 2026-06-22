# StackConnectionTenantV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowlistUrl** | Pointer to **string** | Allowlist API endpoint | [optional] 
**Id** | **int64** | tenant id | 
**PrivateConnectivityInfo** | Pointer to [**BasicPrivateConnectivityInfo**](BasicPrivateConnectivityInfo.md) |  | [optional] 
**Type** | **string** | tenant type | 
**Url** | Pointer to **string** | URL for the tenant | [optional] 

## Methods

### NewStackConnectionTenantV1

`func NewStackConnectionTenantV1(id int64, type_ string, ) *StackConnectionTenantV1`

NewStackConnectionTenantV1 instantiates a new StackConnectionTenantV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackConnectionTenantV1WithDefaults

`func NewStackConnectionTenantV1WithDefaults() *StackConnectionTenantV1`

NewStackConnectionTenantV1WithDefaults instantiates a new StackConnectionTenantV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowlistUrl

`func (o *StackConnectionTenantV1) GetAllowlistUrl() string`

GetAllowlistUrl returns the AllowlistUrl field if non-nil, zero value otherwise.

### GetAllowlistUrlOk

`func (o *StackConnectionTenantV1) GetAllowlistUrlOk() (*string, bool)`

GetAllowlistUrlOk returns a tuple with the AllowlistUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistUrl

`func (o *StackConnectionTenantV1) SetAllowlistUrl(v string)`

SetAllowlistUrl sets AllowlistUrl field to given value.

### HasAllowlistUrl

`func (o *StackConnectionTenantV1) HasAllowlistUrl() bool`

HasAllowlistUrl returns a boolean if a field has been set.

### GetId

`func (o *StackConnectionTenantV1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StackConnectionTenantV1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StackConnectionTenantV1) SetId(v int64)`

SetId sets Id field to given value.


### GetPrivateConnectivityInfo

`func (o *StackConnectionTenantV1) GetPrivateConnectivityInfo() BasicPrivateConnectivityInfo`

GetPrivateConnectivityInfo returns the PrivateConnectivityInfo field if non-nil, zero value otherwise.

### GetPrivateConnectivityInfoOk

`func (o *StackConnectionTenantV1) GetPrivateConnectivityInfoOk() (*BasicPrivateConnectivityInfo, bool)`

GetPrivateConnectivityInfoOk returns a tuple with the PrivateConnectivityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateConnectivityInfo

`func (o *StackConnectionTenantV1) SetPrivateConnectivityInfo(v BasicPrivateConnectivityInfo)`

SetPrivateConnectivityInfo sets PrivateConnectivityInfo field to given value.

### HasPrivateConnectivityInfo

`func (o *StackConnectionTenantV1) HasPrivateConnectivityInfo() bool`

HasPrivateConnectivityInfo returns a boolean if a field has been set.

### GetType

`func (o *StackConnectionTenantV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StackConnectionTenantV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StackConnectionTenantV1) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *StackConnectionTenantV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StackConnectionTenantV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StackConnectionTenantV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StackConnectionTenantV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


