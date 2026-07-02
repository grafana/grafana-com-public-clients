# Gateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateDNS** | **string** |  | 
**ServiceName** | **string** |  | 
**Regions** | Pointer to **[]string** |  | [optional] 
**EndpointName** | **string** |  | 
**ServiceId** | Pointer to **string** |  | [optional] 
**AvailabilityZones** | Pointer to **[]string** |  | [optional] 
**AvailabilityZoneIds** | Pointer to **[]string** |  | [optional] 
**ServiceAttachment** | **string** |  | 
**DomainNames** | **[]string** |  | 

## Methods

### NewGateway

`func NewGateway(privateDNS string, serviceName string, endpointName string, serviceAttachment string, domainNames []string, ) *Gateway`

NewGateway instantiates a new Gateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayWithDefaults

`func NewGatewayWithDefaults() *Gateway`

NewGatewayWithDefaults instantiates a new Gateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateDNS

`func (o *Gateway) GetPrivateDNS() string`

GetPrivateDNS returns the PrivateDNS field if non-nil, zero value otherwise.

### GetPrivateDNSOk

`func (o *Gateway) GetPrivateDNSOk() (*string, bool)`

GetPrivateDNSOk returns a tuple with the PrivateDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDNS

`func (o *Gateway) SetPrivateDNS(v string)`

SetPrivateDNS sets PrivateDNS field to given value.


### GetServiceName

`func (o *Gateway) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Gateway) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Gateway) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetRegions

`func (o *Gateway) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Gateway) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Gateway) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Gateway) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetEndpointName

`func (o *Gateway) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *Gateway) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *Gateway) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.


### GetServiceId

`func (o *Gateway) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Gateway) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Gateway) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *Gateway) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetAvailabilityZones

`func (o *Gateway) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *Gateway) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *Gateway) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *Gateway) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### GetAvailabilityZoneIds

`func (o *Gateway) GetAvailabilityZoneIds() []string`

GetAvailabilityZoneIds returns the AvailabilityZoneIds field if non-nil, zero value otherwise.

### GetAvailabilityZoneIdsOk

`func (o *Gateway) GetAvailabilityZoneIdsOk() (*[]string, bool)`

GetAvailabilityZoneIdsOk returns a tuple with the AvailabilityZoneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneIds

`func (o *Gateway) SetAvailabilityZoneIds(v []string)`

SetAvailabilityZoneIds sets AvailabilityZoneIds field to given value.

### HasAvailabilityZoneIds

`func (o *Gateway) HasAvailabilityZoneIds() bool`

HasAvailabilityZoneIds returns a boolean if a field has been set.

### GetServiceAttachment

`func (o *Gateway) GetServiceAttachment() string`

GetServiceAttachment returns the ServiceAttachment field if non-nil, zero value otherwise.

### GetServiceAttachmentOk

`func (o *Gateway) GetServiceAttachmentOk() (*string, bool)`

GetServiceAttachmentOk returns a tuple with the ServiceAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttachment

`func (o *Gateway) SetServiceAttachment(v string)`

SetServiceAttachment sets ServiceAttachment field to given value.


### GetDomainNames

`func (o *Gateway) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *Gateway) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *Gateway) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


