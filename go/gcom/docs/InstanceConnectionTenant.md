# InstanceConnectionTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowlistUrl** | Pointer to **string** | Allowlist API endpoint | [optional] 
**Id** | **int64** | tenant id | 
**Info** | Pointer to [**BasicPrivateConnectivityInfo**](BasicPrivateConnectivityInfo.md) |  | [optional] 
**IpAllowListCNAME** | **NullableString** | IP allow list CNAME | 
**Type** | **string** | tenant type | 

## Methods

### NewInstanceConnectionTenant

`func NewInstanceConnectionTenant(id int64, ipAllowListCNAME NullableString, type_ string, ) *InstanceConnectionTenant`

NewInstanceConnectionTenant instantiates a new InstanceConnectionTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceConnectionTenantWithDefaults

`func NewInstanceConnectionTenantWithDefaults() *InstanceConnectionTenant`

NewInstanceConnectionTenantWithDefaults instantiates a new InstanceConnectionTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowlistUrl

`func (o *InstanceConnectionTenant) GetAllowlistUrl() string`

GetAllowlistUrl returns the AllowlistUrl field if non-nil, zero value otherwise.

### GetAllowlistUrlOk

`func (o *InstanceConnectionTenant) GetAllowlistUrlOk() (*string, bool)`

GetAllowlistUrlOk returns a tuple with the AllowlistUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowlistUrl

`func (o *InstanceConnectionTenant) SetAllowlistUrl(v string)`

SetAllowlistUrl sets AllowlistUrl field to given value.

### HasAllowlistUrl

`func (o *InstanceConnectionTenant) HasAllowlistUrl() bool`

HasAllowlistUrl returns a boolean if a field has been set.

### GetId

`func (o *InstanceConnectionTenant) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceConnectionTenant) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceConnectionTenant) SetId(v int64)`

SetId sets Id field to given value.


### GetInfo

`func (o *InstanceConnectionTenant) GetInfo() BasicPrivateConnectivityInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *InstanceConnectionTenant) GetInfoOk() (*BasicPrivateConnectivityInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *InstanceConnectionTenant) SetInfo(v BasicPrivateConnectivityInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *InstanceConnectionTenant) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetIpAllowListCNAME

`func (o *InstanceConnectionTenant) GetIpAllowListCNAME() string`

GetIpAllowListCNAME returns the IpAllowListCNAME field if non-nil, zero value otherwise.

### GetIpAllowListCNAMEOk

`func (o *InstanceConnectionTenant) GetIpAllowListCNAMEOk() (*string, bool)`

GetIpAllowListCNAMEOk returns a tuple with the IpAllowListCNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowListCNAME

`func (o *InstanceConnectionTenant) SetIpAllowListCNAME(v string)`

SetIpAllowListCNAME sets IpAllowListCNAME field to given value.


### SetIpAllowListCNAMENil

`func (o *InstanceConnectionTenant) SetIpAllowListCNAMENil(b bool)`

 SetIpAllowListCNAMENil sets the value for IpAllowListCNAME to be an explicit nil

### UnsetIpAllowListCNAME
`func (o *InstanceConnectionTenant) UnsetIpAllowListCNAME()`

UnsetIpAllowListCNAME ensures that no value is present for IpAllowListCNAME, not even an explicit nil
### GetType

`func (o *InstanceConnectionTenant) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceConnectionTenant) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceConnectionTenant) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


