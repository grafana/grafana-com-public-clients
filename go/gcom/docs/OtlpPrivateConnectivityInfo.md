# OtlpPrivateConnectivityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateDNS** | **string** |  | 
**ServiceName** | **string** |  | 
**Mimir** | [**Mimir**](Mimir.md) |  | 
**Graphite** | Pointer to [**Graphite**](Graphite.md) |  | [optional] 

## Methods

### NewOtlpPrivateConnectivityInfo

`func NewOtlpPrivateConnectivityInfo(privateDNS string, serviceName string, mimir Mimir, ) *OtlpPrivateConnectivityInfo`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


