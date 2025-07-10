# OtlpPrivateConnectivityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateDNS** | **string** |  | 
**ServiceName** | **string** |  | 
**Regions** | Pointer to **[]string** |  | [optional] 
**EndpointName** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**Api** | [**Api1**](Api1.md) |  | 
**Gateway** | [**Gateway1**](Gateway1.md) |  | 
**Mimir** | [**Mimir**](Mimir.md) |  | 
**Graphite** | Pointer to [**Graphite**](Graphite.md) |  | [optional] 
**Otlp** | Pointer to [**NullablePdcPrivateConnectivityInfoAnyOf**](PdcPrivateConnectivityInfoAnyOf.md) |  | [optional] 
**Name** | **string** |  | 
**Network** | **string** |  | 
**IpCidrRange** | **string** |  | 

## Methods

### NewOtlpPrivateConnectivityInfo

`func NewOtlpPrivateConnectivityInfo(privateDNS string, serviceName string, api Api1, gateway Gateway1, mimir Mimir, name string, network string, ipCidrRange string, ) *OtlpPrivateConnectivityInfo`

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


### GetRegions

`func (o *OtlpPrivateConnectivityInfo) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *OtlpPrivateConnectivityInfo) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *OtlpPrivateConnectivityInfo) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *OtlpPrivateConnectivityInfo) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetEndpointName

`func (o *OtlpPrivateConnectivityInfo) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *OtlpPrivateConnectivityInfo) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *OtlpPrivateConnectivityInfo) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *OtlpPrivateConnectivityInfo) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetServiceId

`func (o *OtlpPrivateConnectivityInfo) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *OtlpPrivateConnectivityInfo) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *OtlpPrivateConnectivityInfo) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *OtlpPrivateConnectivityInfo) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetApi

`func (o *OtlpPrivateConnectivityInfo) GetApi() Api1`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *OtlpPrivateConnectivityInfo) GetApiOk() (*Api1, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *OtlpPrivateConnectivityInfo) SetApi(v Api1)`

SetApi sets Api field to given value.


### GetGateway

`func (o *OtlpPrivateConnectivityInfo) GetGateway() Gateway1`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *OtlpPrivateConnectivityInfo) GetGatewayOk() (*Gateway1, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *OtlpPrivateConnectivityInfo) SetGateway(v Gateway1)`

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
### GetName

`func (o *OtlpPrivateConnectivityInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OtlpPrivateConnectivityInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OtlpPrivateConnectivityInfo) SetName(v string)`

SetName sets Name field to given value.


### GetNetwork

`func (o *OtlpPrivateConnectivityInfo) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *OtlpPrivateConnectivityInfo) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *OtlpPrivateConnectivityInfo) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetIpCidrRange

`func (o *OtlpPrivateConnectivityInfo) GetIpCidrRange() string`

GetIpCidrRange returns the IpCidrRange field if non-nil, zero value otherwise.

### GetIpCidrRangeOk

`func (o *OtlpPrivateConnectivityInfo) GetIpCidrRangeOk() (*string, bool)`

GetIpCidrRangeOk returns a tuple with the IpCidrRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpCidrRange

`func (o *OtlpPrivateConnectivityInfo) SetIpCidrRange(v string)`

SetIpCidrRange sets IpCidrRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


