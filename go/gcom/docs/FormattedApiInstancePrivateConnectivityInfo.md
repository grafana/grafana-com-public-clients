# FormattedApiInstancePrivateConnectivityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Otlp** | Pointer to [**BasicPrivateConnectivityInfo**](BasicPrivateConnectivityInfo.md) |  | [optional] 
**Pdc** | Pointer to [**StackConnectionPdcV1**](StackConnectionPdcV1.md) |  | [optional] 
**Tenants** | [**[]InstanceConnectionTenant**](InstanceConnectionTenant.md) |  | 

## Methods

### NewFormattedApiInstancePrivateConnectivityInfo

`func NewFormattedApiInstancePrivateConnectivityInfo(tenants []InstanceConnectionTenant, ) *FormattedApiInstancePrivateConnectivityInfo`

NewFormattedApiInstancePrivateConnectivityInfo instantiates a new FormattedApiInstancePrivateConnectivityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedApiInstancePrivateConnectivityInfoWithDefaults

`func NewFormattedApiInstancePrivateConnectivityInfoWithDefaults() *FormattedApiInstancePrivateConnectivityInfo`

NewFormattedApiInstancePrivateConnectivityInfoWithDefaults instantiates a new FormattedApiInstancePrivateConnectivityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtlp

`func (o *FormattedApiInstancePrivateConnectivityInfo) GetOtlp() BasicPrivateConnectivityInfo`

GetOtlp returns the Otlp field if non-nil, zero value otherwise.

### GetOtlpOk

`func (o *FormattedApiInstancePrivateConnectivityInfo) GetOtlpOk() (*BasicPrivateConnectivityInfo, bool)`

GetOtlpOk returns a tuple with the Otlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlp

`func (o *FormattedApiInstancePrivateConnectivityInfo) SetOtlp(v BasicPrivateConnectivityInfo)`

SetOtlp sets Otlp field to given value.

### HasOtlp

`func (o *FormattedApiInstancePrivateConnectivityInfo) HasOtlp() bool`

HasOtlp returns a boolean if a field has been set.

### GetPdc

`func (o *FormattedApiInstancePrivateConnectivityInfo) GetPdc() StackConnectionPdcV1`

GetPdc returns the Pdc field if non-nil, zero value otherwise.

### GetPdcOk

`func (o *FormattedApiInstancePrivateConnectivityInfo) GetPdcOk() (*StackConnectionPdcV1, bool)`

GetPdcOk returns a tuple with the Pdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdc

`func (o *FormattedApiInstancePrivateConnectivityInfo) SetPdc(v StackConnectionPdcV1)`

SetPdc sets Pdc field to given value.

### HasPdc

`func (o *FormattedApiInstancePrivateConnectivityInfo) HasPdc() bool`

HasPdc returns a boolean if a field has been set.

### GetTenants

`func (o *FormattedApiInstancePrivateConnectivityInfo) GetTenants() []InstanceConnectionTenant`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *FormattedApiInstancePrivateConnectivityInfo) GetTenantsOk() (*[]InstanceConnectionTenant, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *FormattedApiInstancePrivateConnectivityInfo) SetTenants(v []InstanceConnectionTenant)`

SetTenants sets Tenants field to given value.


### SetTenantsNil

`func (o *FormattedApiInstancePrivateConnectivityInfo) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *FormattedApiInstancePrivateConnectivityInfo) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


