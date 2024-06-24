# TenantsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **float32** |  | 
**Info** | Pointer to [**Info**](Info.md) |  | [optional] 
**IpAllowListCNAME** | Pointer to **string** |  | [optional] 

## Methods

### NewTenantsInner

`func NewTenantsInner(type_ string, id float32, ) *TenantsInner`

NewTenantsInner instantiates a new TenantsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantsInnerWithDefaults

`func NewTenantsInnerWithDefaults() *TenantsInner`

NewTenantsInnerWithDefaults instantiates a new TenantsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TenantsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TenantsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TenantsInner) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *TenantsInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantsInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantsInner) SetId(v float32)`

SetId sets Id field to given value.


### GetInfo

`func (o *TenantsInner) GetInfo() Info`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TenantsInner) GetInfoOk() (*Info, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TenantsInner) SetInfo(v Info)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TenantsInner) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetIpAllowListCNAME

`func (o *TenantsInner) GetIpAllowListCNAME() string`

GetIpAllowListCNAME returns the IpAllowListCNAME field if non-nil, zero value otherwise.

### GetIpAllowListCNAMEOk

`func (o *TenantsInner) GetIpAllowListCNAMEOk() (*string, bool)`

GetIpAllowListCNAMEOk returns a tuple with the IpAllowListCNAME field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAllowListCNAME

`func (o *TenantsInner) SetIpAllowListCNAME(v string)`

SetIpAllowListCNAME sets IpAllowListCNAME field to given value.

### HasIpAllowListCNAME

`func (o *TenantsInner) HasIpAllowListCNAME() bool`

HasIpAllowListCNAME returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


