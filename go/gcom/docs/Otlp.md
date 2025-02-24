# Otlp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** |  | 
**PrivateDNS** | **string** |  | 
**Regions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOtlp

`func NewOtlp(serviceName string, privateDNS string, ) *Otlp`

NewOtlp instantiates a new Otlp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtlpWithDefaults

`func NewOtlpWithDefaults() *Otlp`

NewOtlpWithDefaults instantiates a new Otlp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *Otlp) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Otlp) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Otlp) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetPrivateDNS

`func (o *Otlp) GetPrivateDNS() string`

GetPrivateDNS returns the PrivateDNS field if non-nil, zero value otherwise.

### GetPrivateDNSOk

`func (o *Otlp) GetPrivateDNSOk() (*string, bool)`

GetPrivateDNSOk returns a tuple with the PrivateDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDNS

`func (o *Otlp) SetPrivateDNS(v string)`

SetPrivateDNS sets PrivateDNS field to given value.


### GetRegions

`func (o *Otlp) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Otlp) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Otlp) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Otlp) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


