# Api

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

### NewApi

`func NewApi(privateDNS string, serviceName string, endpointName string, serviceAttachment string, domainNames []string, ) *Api`

NewApi instantiates a new Api object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWithDefaults

`func NewApiWithDefaults() *Api`

NewApiWithDefaults instantiates a new Api object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateDNS

`func (o *Api) GetPrivateDNS() string`

GetPrivateDNS returns the PrivateDNS field if non-nil, zero value otherwise.

### GetPrivateDNSOk

`func (o *Api) GetPrivateDNSOk() (*string, bool)`

GetPrivateDNSOk returns a tuple with the PrivateDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDNS

`func (o *Api) SetPrivateDNS(v string)`

SetPrivateDNS sets PrivateDNS field to given value.


### GetServiceName

`func (o *Api) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Api) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Api) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetRegions

`func (o *Api) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Api) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Api) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Api) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetEndpointName

`func (o *Api) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *Api) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *Api) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.


### GetServiceId

`func (o *Api) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *Api) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *Api) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *Api) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetAvailabilityZones

`func (o *Api) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *Api) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *Api) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *Api) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### GetAvailabilityZoneIds

`func (o *Api) GetAvailabilityZoneIds() []string`

GetAvailabilityZoneIds returns the AvailabilityZoneIds field if non-nil, zero value otherwise.

### GetAvailabilityZoneIdsOk

`func (o *Api) GetAvailabilityZoneIdsOk() (*[]string, bool)`

GetAvailabilityZoneIdsOk returns a tuple with the AvailabilityZoneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneIds

`func (o *Api) SetAvailabilityZoneIds(v []string)`

SetAvailabilityZoneIds sets AvailabilityZoneIds field to given value.

### HasAvailabilityZoneIds

`func (o *Api) HasAvailabilityZoneIds() bool`

HasAvailabilityZoneIds returns a boolean if a field has been set.

### GetServiceAttachment

`func (o *Api) GetServiceAttachment() string`

GetServiceAttachment returns the ServiceAttachment field if non-nil, zero value otherwise.

### GetServiceAttachmentOk

`func (o *Api) GetServiceAttachmentOk() (*string, bool)`

GetServiceAttachmentOk returns a tuple with the ServiceAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttachment

`func (o *Api) SetServiceAttachment(v string)`

SetServiceAttachment sets ServiceAttachment field to given value.


### GetDomainNames

`func (o *Api) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *Api) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *Api) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


