# StackConnectionsV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppPlatform** | Pointer to [**StackConnectionAppPlatformV1**](StackConnectionAppPlatformV1.md) |  | [optional] 
**Otlp** | Pointer to [**StackConnectionOtlpV1**](StackConnectionOtlpV1.md) |  | [optional] 
**Pdc** | Pointer to [**StackConnectionPdcV1**](StackConnectionPdcV1.md) |  | [optional] 
**Services** | Pointer to [**StackConnectionServicesV1**](StackConnectionServicesV1.md) |  | [optional] 
**Tenants** | [**[]StackConnectionTenantV1**](StackConnectionTenantV1.md) |  | 

## Methods

### NewStackConnectionsV1

`func NewStackConnectionsV1(tenants []StackConnectionTenantV1, ) *StackConnectionsV1`

NewStackConnectionsV1 instantiates a new StackConnectionsV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackConnectionsV1WithDefaults

`func NewStackConnectionsV1WithDefaults() *StackConnectionsV1`

NewStackConnectionsV1WithDefaults instantiates a new StackConnectionsV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppPlatform

`func (o *StackConnectionsV1) GetAppPlatform() StackConnectionAppPlatformV1`

GetAppPlatform returns the AppPlatform field if non-nil, zero value otherwise.

### GetAppPlatformOk

`func (o *StackConnectionsV1) GetAppPlatformOk() (*StackConnectionAppPlatformV1, bool)`

GetAppPlatformOk returns a tuple with the AppPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPlatform

`func (o *StackConnectionsV1) SetAppPlatform(v StackConnectionAppPlatformV1)`

SetAppPlatform sets AppPlatform field to given value.

### HasAppPlatform

`func (o *StackConnectionsV1) HasAppPlatform() bool`

HasAppPlatform returns a boolean if a field has been set.

### GetOtlp

`func (o *StackConnectionsV1) GetOtlp() StackConnectionOtlpV1`

GetOtlp returns the Otlp field if non-nil, zero value otherwise.

### GetOtlpOk

`func (o *StackConnectionsV1) GetOtlpOk() (*StackConnectionOtlpV1, bool)`

GetOtlpOk returns a tuple with the Otlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlp

`func (o *StackConnectionsV1) SetOtlp(v StackConnectionOtlpV1)`

SetOtlp sets Otlp field to given value.

### HasOtlp

`func (o *StackConnectionsV1) HasOtlp() bool`

HasOtlp returns a boolean if a field has been set.

### GetPdc

`func (o *StackConnectionsV1) GetPdc() StackConnectionPdcV1`

GetPdc returns the Pdc field if non-nil, zero value otherwise.

### GetPdcOk

`func (o *StackConnectionsV1) GetPdcOk() (*StackConnectionPdcV1, bool)`

GetPdcOk returns a tuple with the Pdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdc

`func (o *StackConnectionsV1) SetPdc(v StackConnectionPdcV1)`

SetPdc sets Pdc field to given value.

### HasPdc

`func (o *StackConnectionsV1) HasPdc() bool`

HasPdc returns a boolean if a field has been set.

### GetServices

`func (o *StackConnectionsV1) GetServices() StackConnectionServicesV1`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *StackConnectionsV1) GetServicesOk() (*StackConnectionServicesV1, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *StackConnectionsV1) SetServices(v StackConnectionServicesV1)`

SetServices sets Services field to given value.

### HasServices

`func (o *StackConnectionsV1) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTenants

`func (o *StackConnectionsV1) GetTenants() []StackConnectionTenantV1`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *StackConnectionsV1) GetTenantsOk() (*[]StackConnectionTenantV1, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *StackConnectionsV1) SetTenants(v []StackConnectionTenantV1)`

SetTenants sets Tenants field to given value.


### SetTenantsNil

`func (o *StackConnectionsV1) SetTenantsNil(b bool)`

 SetTenantsNil sets the value for Tenants to be an explicit nil

### UnsetTenants
`func (o *StackConnectionsV1) UnsetTenants()`

UnsetTenants ensures that no value is present for Tenants, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


