# InfoAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** |  | 
**PrivateDNS** | **string** |  | 
**Regions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewInfoAnyOf

`func NewInfoAnyOf(serviceName string, privateDNS string, ) *InfoAnyOf`

NewInfoAnyOf instantiates a new InfoAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfoAnyOfWithDefaults

`func NewInfoAnyOfWithDefaults() *InfoAnyOf`

NewInfoAnyOfWithDefaults instantiates a new InfoAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *InfoAnyOf) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *InfoAnyOf) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *InfoAnyOf) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetPrivateDNS

`func (o *InfoAnyOf) GetPrivateDNS() string`

GetPrivateDNS returns the PrivateDNS field if non-nil, zero value otherwise.

### GetPrivateDNSOk

`func (o *InfoAnyOf) GetPrivateDNSOk() (*string, bool)`

GetPrivateDNSOk returns a tuple with the PrivateDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDNS

`func (o *InfoAnyOf) SetPrivateDNS(v string)`

SetPrivateDNS sets PrivateDNS field to given value.


### GetRegions

`func (o *InfoAnyOf) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *InfoAnyOf) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *InfoAnyOf) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *InfoAnyOf) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


