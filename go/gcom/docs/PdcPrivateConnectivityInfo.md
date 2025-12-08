# PdcPrivateConnectivityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateDNS** | **string** |  | 
**ServiceName** | **string** |  | 
**Regions** | Pointer to **[]string** |  | [optional] 
**EndpointName** | **string** |  | 
**ServiceId** | Pointer to **string** |  | [optional] 
**ServiceAttachment** | **string** |  | 
**DomainNames** | **[]string** |  | 
**Api** | [**Api1**](Api1.md) |  | 
**Gateway** | [**Gateway1**](Gateway1.md) |  | 
**Mimir** | [**Mimir**](Mimir.md) |  | 
**Graphite** | Pointer to [**Graphite**](Graphite.md) |  | [optional] 
**Otlp** | Pointer to [**Otlp1**](Otlp1.md) |  | [optional] 

## Methods

### NewPdcPrivateConnectivityInfo

`func NewPdcPrivateConnectivityInfo(privateDNS string, serviceName string, endpointName string, serviceAttachment string, domainNames []string, api Api1, gateway Gateway1, mimir Mimir, ) *PdcPrivateConnectivityInfo`

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


### GetRegions

`func (o *PdcPrivateConnectivityInfo) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *PdcPrivateConnectivityInfo) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *PdcPrivateConnectivityInfo) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *PdcPrivateConnectivityInfo) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetEndpointName

`func (o *PdcPrivateConnectivityInfo) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *PdcPrivateConnectivityInfo) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *PdcPrivateConnectivityInfo) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.


### GetServiceId

`func (o *PdcPrivateConnectivityInfo) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PdcPrivateConnectivityInfo) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PdcPrivateConnectivityInfo) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *PdcPrivateConnectivityInfo) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceAttachment

`func (o *PdcPrivateConnectivityInfo) GetServiceAttachment() string`

GetServiceAttachment returns the ServiceAttachment field if non-nil, zero value otherwise.

### GetServiceAttachmentOk

`func (o *PdcPrivateConnectivityInfo) GetServiceAttachmentOk() (*string, bool)`

GetServiceAttachmentOk returns a tuple with the ServiceAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttachment

`func (o *PdcPrivateConnectivityInfo) SetServiceAttachment(v string)`

SetServiceAttachment sets ServiceAttachment field to given value.


### GetDomainNames

`func (o *PdcPrivateConnectivityInfo) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *PdcPrivateConnectivityInfo) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *PdcPrivateConnectivityInfo) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.


### GetApi

`func (o *PdcPrivateConnectivityInfo) GetApi() Api1`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *PdcPrivateConnectivityInfo) GetApiOk() (*Api1, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *PdcPrivateConnectivityInfo) SetApi(v Api1)`

SetApi sets Api field to given value.


### GetGateway

`func (o *PdcPrivateConnectivityInfo) GetGateway() Gateway1`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *PdcPrivateConnectivityInfo) GetGatewayOk() (*Gateway1, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *PdcPrivateConnectivityInfo) SetGateway(v Gateway1)`

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

`func (o *PdcPrivateConnectivityInfo) GetOtlp() Otlp1`

GetOtlp returns the Otlp field if non-nil, zero value otherwise.

### GetOtlpOk

`func (o *PdcPrivateConnectivityInfo) GetOtlpOk() (*Otlp1, bool)`

GetOtlpOk returns a tuple with the Otlp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlp

`func (o *PdcPrivateConnectivityInfo) SetOtlp(v Otlp1)`

SetOtlp sets Otlp field to given value.

### HasOtlp

`func (o *PdcPrivateConnectivityInfo) HasOtlp() bool`

HasOtlp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


