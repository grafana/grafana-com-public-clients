# Api1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateDNS** | **string** |  | 
**ServiceName** | **string** |  | 
**Regions** | Pointer to **[]string** |  | [optional] 
**EndpointName** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 

## Methods

### NewApi1

`func NewApi1(privateDNS string, serviceName string, ) *Api1`

NewApi1 instantiates a new Api1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApi1WithDefaults

`func NewApi1WithDefaults() *Api1`

NewApi1WithDefaults instantiates a new Api1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateDNS

`func (o *Api1) GetPrivateDNS() string`

GetPrivateDNS returns the PrivateDNS field if non-nil, zero value otherwise.

### GetPrivateDNSOk

`func (o *Api1) GetPrivateDNSOk() (*string, bool)`

GetPrivateDNSOk returns a tuple with the PrivateDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDNS

`func (o *Api1) SetPrivateDNS(v string)`

SetPrivateDNS sets PrivateDNS field to given value.


### GetServiceName

`func (o *Api1) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Api1) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Api1) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetRegions

`func (o *Api1) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Api1) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Api1) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Api1) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetEndpointName

`func (o *Api1) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *Api1) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *Api1) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *Api1) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetServiceId

`func (o *Api1) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Api1) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Api1) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *Api1) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


