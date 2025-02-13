# OtlpPrivateConnectivityInfo

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

### NewOtlpPrivateConnectivityInfo

`func NewOtlpPrivateConnectivityInfo(privateDNS string, serviceName string, api Api, gateway Gateway, mimir Mimir, ) *OtlpPrivateConnectivityInfo`

NewOtlpPrivateConnectivityInfo instantiates a new OtlpPrivateConnectivityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtlpPrivateConnectivityInfoWithDefaults

`func NewOtlpPrivateConnectivityInfoWithDefaults() *OtlpPrivateConnectivityInfo`

NewOtlpPrivateConnectivityInfoWithDefaults instantiates a new OtlpPrivateConnectivityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateDNS

`func (o *OtlpPrivateConnectivityInfo) GetPrivateDNS() string`

GetPrivateDNS returns the PrivateDNS field if non-nil, zero value otherwise.

### GetPrivateDNSOk

`func (o *OtlpPrivateConnectivityInfo) GetPrivateDNSOk() (*string, bool)`

GetPrivateDNSOk returns a tuple with the PrivateDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDNS

`func (o *OtlpPrivateConnectivityInfo) SetPrivateDNS(v string)`

SetPrivateDNS sets PrivateDNS field to given value.


### GetServiceName

`func (o *OtlpPrivateConnectivityInfo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *OtlpPrivateConnectivityInfo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *OtlpPrivateConnectivityInfo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetApi

`func (o *OtlpPrivateConnectivityInfo) GetApi() Api`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *OtlpPrivateConnectivityInfo) GetApiOk() (*Api, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *OtlpPrivateConnectivityInfo) SetApi(v Api)`

SetApi sets Api field to given value.


### GetGateway

`func (o *OtlpPrivateConnectivityInfo) GetGateway() Gateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *OtlpPrivateConnectivityInfo) GetGatewayOk() (*Gateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *OtlpPrivateConnectivityInfo) SetGateway(v Gateway)`

SetGateway sets Gateway field to given value.


### GetMimir

`func (o *OtlpPrivateConnectivityInfo) GetMimir() Mimir`

GetMimir returns the Mimir field if non-nil, zero value otherwise.

### GetMimirOk

`func (o *OtlpPrivateConnectivityInfo) GetMimirOk() (*Mimir, bool)`

GetMimirOk returns a tuple with the Mimir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimir

`func (o *OtlpPrivateConnectivityInfo) SetMimir(v Mimir)`

SetMimir sets Mimir field to given value.


### GetGraphite

`func (o *OtlpPrivateConnectivityInfo) GetGraphite() Graphite`

GetGraphite returns the Graphite field if non-nil, zero value otherwise.

### GetGraphiteOk

`func (o *OtlpPrivateConnectivityInfo) GetGraphiteOk() (*Graphite, bool)`

GetGraphiteOk returns a tuple with the Graphite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphite

`func (o *OtlpPrivateConnectivityInfo) SetGraphite(v Graphite)`

SetGraphite sets Graphite field to given value.

### HasGraphite

`func (o *OtlpPrivateConnectivityInfo) HasGraphite() bool`

HasGraphite returns a boolean if a field has been set.

### GetOtlp

`func (o *OtlpPrivateConnectivityInfo) GetOtlp() PdcPrivateConnectivityInfoAnyOf`

GetOtlp returns the Otlp field if non-nil, zero value otherwise.

### GetOtlpOk

`func (o *OtlpPrivateConnectivityInfo) GetOtlpOk() (*PdcPrivateConnectivityInfoAnyOf, bool)`

GetOtlpOk returns a tuple with the Otlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlp

`func (o *OtlpPrivateConnectivityInfo) SetOtlp(v PdcPrivateConnectivityInfoAnyOf)`

SetOtlp sets Otlp field to given value.

### HasOtlp

`func (o *OtlpPrivateConnectivityInfo) HasOtlp() bool`

HasOtlp returns a boolean if a field has been set.

### SetOtlpNil

`func (o *OtlpPrivateConnectivityInfo) SetOtlpNil(b bool)`

 SetOtlpNil sets the value for Otlp to be an explicit nil

### UnsetOtlp
`func (o *OtlpPrivateConnectivityInfo) UnsetOtlp()`

UnsetOtlp ensures that no value is present for Otlp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


