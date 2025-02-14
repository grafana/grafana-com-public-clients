# PrivateConnectivityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenants** | [**[]TenantsInner**](TenantsInner.md) |  | 
**Otlp** | Pointer to [**Otlp**](Otlp.md) |  | [optional] 
**Pdc** | Pointer to [**Pdc**](Pdc.md) |  | [optional] 

## Methods

### NewPrivateConnectivityInfo

`func NewPrivateConnectivityInfo(tenants []TenantsInner, ) *PrivateConnectivityInfo`

NewPrivateConnectivityInfo instantiates a new PrivateConnectivityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateConnectivityInfoWithDefaults

`func NewPrivateConnectivityInfoWithDefaults() *PrivateConnectivityInfo`

NewPrivateConnectivityInfoWithDefaults instantiates a new PrivateConnectivityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenants

`func (o *PrivateConnectivityInfo) GetTenants() []TenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *PrivateConnectivityInfo) GetTenantsOk() (*[]TenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *PrivateConnectivityInfo) SetTenants(v []TenantsInner)`

SetTenants sets Tenants field to given value.


### GetOtlp

`func (o *PrivateConnectivityInfo) GetOtlp() Otlp`

GetOtlp returns the Otlp field if non-nil, zero value otherwise.

### GetOtlpOk

`func (o *PrivateConnectivityInfo) GetOtlpOk() (*Otlp, bool)`

GetOtlpOk returns a tuple with the Otlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlp

`func (o *PrivateConnectivityInfo) SetOtlp(v Otlp)`

SetOtlp sets Otlp field to given value.

### HasOtlp

`func (o *PrivateConnectivityInfo) HasOtlp() bool`

HasOtlp returns a boolean if a field has been set.

### GetPdc

`func (o *PrivateConnectivityInfo) GetPdc() Pdc`

GetPdc returns the Pdc field if non-nil, zero value otherwise.

### GetPdcOk

`func (o *PrivateConnectivityInfo) GetPdcOk() (*Pdc, bool)`

GetPdcOk returns a tuple with the Pdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdc

`func (o *PrivateConnectivityInfo) SetPdc(v Pdc)`

SetPdc sets Pdc field to given value.

### HasPdc

`func (o *PrivateConnectivityInfo) HasPdc() bool`

HasPdc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


