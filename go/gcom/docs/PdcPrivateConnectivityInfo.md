# PdcPrivateConnectivityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateDNS** | **string** |  | 
**ServiceName** | **string** |  | 
**Api** | [**Api**](Api.md) |  | 
**Gateway** | [**Gateway**](Gateway.md) |  | 
**Mimir** | [**Mimir**](Mimir.md) |  | 
**Graphite** | Pointer to [**Graphite**](Graphite.md) |  | [optional] 
**Otlp** | Pointer to [**NullablePdcPrivateConnectivityInfoAnyOf**](PdcPrivateConnectivityInfoAnyOf.md) |  | [optional] 

## Methods

### NewPdcPrivateConnectivityInfo

`func NewPdcPrivateConnectivityInfo(privateDNS string, serviceName string, api Api, gateway Gateway, mimir Mimir, ) *PdcPrivateConnectivityInfo`

NewPdcPrivateConnectivityInfo instantiates a new PdcPrivateConnectivityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPdcPrivateConnectivityInfoWithDefaults

`func NewPdcPrivateConnectivityInfoWithDefaults() *PdcPrivateConnectivityInfo`

NewPdcPrivateConnectivityInfoWithDefaults instantiates a new PdcPrivateConnectivityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateDNS

`func (o *PdcPrivateConnectivityInfo) GetPrivateDNS() string`

GetPrivateDNS returns the PrivateDNS field if non-nil, zero value otherwise.

### GetPrivateDNSOk

`func (o *PdcPrivateConnectivityInfo) GetPrivateDNSOk() (*string, bool)`

GetPrivateDNSOk returns a tuple with the PrivateDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDNS

`func (o *PdcPrivateConnectivityInfo) SetPrivateDNS(v string)`

SetPrivateDNS sets PrivateDNS field to given value.


### GetServiceName

`func (o *PdcPrivateConnectivityInfo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *PdcPrivateConnectivityInfo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *PdcPrivateConnectivityInfo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetApi

`func (o *PdcPrivateConnectivityInfo) GetApi() Api`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *PdcPrivateConnectivityInfo) GetApiOk() (*Api, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *PdcPrivateConnectivityInfo) SetApi(v Api)`

SetApi sets Api field to given value.


### GetGateway

`func (o *PdcPrivateConnectivityInfo) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PdcPrivateConnectivityInfo) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PdcPrivateConnectivityInfo) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.


### GetMimir

`func (o *PdcPrivateConnectivityInfo) GetMimir() Mimir`

GetMimir returns the Mimir field if non-nil, zero value otherwise.

### GetMimirOk

`func (o *PdcPrivateConnectivityInfo) GetMimirOk() (*Mimir, bool)`

GetMimirOk returns a tuple with the Mimir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimir

`func (o *PdcPrivateConnectivityInfo) SetMimir(v Mimir)`

SetMimir sets Mimir field to given value.


### GetGraphite

`func (o *PdcPrivateConnectivityInfo) GetGraphite() Graphite`

GetGraphite returns the Graphite field if non-nil, zero value otherwise.

### GetGraphiteOk

`func (o *PdcPrivateConnectivityInfo) GetGraphiteOk() (*Graphite, bool)`

GetGraphiteOk returns a tuple with the Graphite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphite

`func (o *PdcPrivateConnectivityInfo) SetGraphite(v Graphite)`

SetGraphite sets Graphite field to given value.

### HasGraphite

`func (o *PdcPrivateConnectivityInfo) HasGraphite() bool`

HasGraphite returns a boolean if a field has been set.

### GetOtlp

`func (o *PdcPrivateConnectivityInfo) GetOtlp() PdcPrivateConnectivityInfoAnyOf`

GetOtlp returns the Otlp field if non-nil, zero value otherwise.

### GetOtlpOk

`func (o *PdcPrivateConnectivityInfo) GetOtlpOk() (*PdcPrivateConnectivityInfoAnyOf, bool)`

GetOtlpOk returns a tuple with the Otlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlp

`func (o *PdcPrivateConnectivityInfo) SetOtlp(v PdcPrivateConnectivityInfoAnyOf)`

SetOtlp sets Otlp field to given value.

### HasOtlp

`func (o *PdcPrivateConnectivityInfo) HasOtlp() bool`

HasOtlp returns a boolean if a field has been set.

### SetOtlpNil

`func (o *PdcPrivateConnectivityInfo) SetOtlpNil(b bool)`

 SetOtlpNil sets the value for Otlp to be an explicit nil

### UnsetOtlp
`func (o *PdcPrivateConnectivityInfo) UnsetOtlp()`

UnsetOtlp ensures that no value is present for Otlp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


