# Graphite

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

## Methods

### NewGraphite

`func NewGraphite(privateDNS string, serviceName string, endpointName string, serviceAttachment string, domainNames []string, ) *Graphite`

NewGraphite instantiates a new Graphite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphiteWithDefaults

`func NewGraphiteWithDefaults() *Graphite`

NewGraphiteWithDefaults instantiates a new Graphite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateDNS

`func (o *Graphite) GetPrivateDNS() string`

GetPrivateDNS returns the PrivateDNS field if non-nil, zero value otherwise.

### GetPrivateDNSOk

`func (o *Graphite) GetPrivateDNSOk() (*string, bool)`

GetPrivateDNSOk returns a tuple with the PrivateDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDNS

`func (o *Graphite) SetPrivateDNS(v string)`

SetPrivateDNS sets PrivateDNS field to given value.


### GetServiceName

`func (o *Graphite) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Graphite) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Graphite) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetRegions

`func (o *Graphite) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Graphite) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Graphite) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Graphite) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetEndpointName

`func (o *Graphite) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *Graphite) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *Graphite) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.


### GetServiceId

`func (o *Graphite) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Graphite) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Graphite) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *Graphite) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceAttachment

`func (o *Graphite) GetServiceAttachment() string`

GetServiceAttachment returns the ServiceAttachment field if non-nil, zero value otherwise.

### GetServiceAttachmentOk

`func (o *Graphite) GetServiceAttachmentOk() (*string, bool)`

GetServiceAttachmentOk returns a tuple with the ServiceAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttachment

`func (o *Graphite) SetServiceAttachment(v string)`

SetServiceAttachment sets ServiceAttachment field to given value.


### GetDomainNames

`func (o *Graphite) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *Graphite) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *Graphite) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


