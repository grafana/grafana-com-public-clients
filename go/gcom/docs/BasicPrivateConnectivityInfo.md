# BasicPrivateConnectivityInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailabilityZoneIds** | Pointer to **[]string** |  | [optional] 
**AvailabilityZones** | Pointer to **[]string** |  | [optional] 
**DomainNames** | Pointer to **[]string** |  | [optional] 
**EndpointName** | Pointer to **string** |  | [optional] 
**PrivateDNS** | Pointer to **string** |  | [optional] 
**Regions** | Pointer to **[]string** |  | [optional] 
**ServiceAttachment** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 

## Methods

### NewBasicPrivateConnectivityInfo

`func NewBasicPrivateConnectivityInfo() *BasicPrivateConnectivityInfo`

NewBasicPrivateConnectivityInfo instantiates a new BasicPrivateConnectivityInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasicPrivateConnectivityInfoWithDefaults

`func NewBasicPrivateConnectivityInfoWithDefaults() *BasicPrivateConnectivityInfo`

NewBasicPrivateConnectivityInfoWithDefaults instantiates a new BasicPrivateConnectivityInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailabilityZoneIds

`func (o *BasicPrivateConnectivityInfo) GetAvailabilityZoneIds() []string`

GetAvailabilityZoneIds returns the AvailabilityZoneIds field if non-nil, zero value otherwise.

### GetAvailabilityZoneIdsOk

`func (o *BasicPrivateConnectivityInfo) GetAvailabilityZoneIdsOk() (*[]string, bool)`

GetAvailabilityZoneIdsOk returns a tuple with the AvailabilityZoneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZoneIds

`func (o *BasicPrivateConnectivityInfo) SetAvailabilityZoneIds(v []string)`

SetAvailabilityZoneIds sets AvailabilityZoneIds field to given value.

### HasAvailabilityZoneIds

`func (o *BasicPrivateConnectivityInfo) HasAvailabilityZoneIds() bool`

HasAvailabilityZoneIds returns a boolean if a field has been set.

### SetAvailabilityZoneIdsNil

`func (o *BasicPrivateConnectivityInfo) SetAvailabilityZoneIdsNil(b bool)`

 SetAvailabilityZoneIdsNil sets the value for AvailabilityZoneIds to be an explicit nil

### UnsetAvailabilityZoneIds
`func (o *BasicPrivateConnectivityInfo) UnsetAvailabilityZoneIds()`

UnsetAvailabilityZoneIds ensures that no value is present for AvailabilityZoneIds, not even an explicit nil
### GetAvailabilityZones

`func (o *BasicPrivateConnectivityInfo) GetAvailabilityZones() []string`

GetAvailabilityZones returns the AvailabilityZones field if non-nil, zero value otherwise.

### GetAvailabilityZonesOk

`func (o *BasicPrivateConnectivityInfo) GetAvailabilityZonesOk() (*[]string, bool)`

GetAvailabilityZonesOk returns a tuple with the AvailabilityZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZones

`func (o *BasicPrivateConnectivityInfo) SetAvailabilityZones(v []string)`

SetAvailabilityZones sets AvailabilityZones field to given value.

### HasAvailabilityZones

`func (o *BasicPrivateConnectivityInfo) HasAvailabilityZones() bool`

HasAvailabilityZones returns a boolean if a field has been set.

### SetAvailabilityZonesNil

`func (o *BasicPrivateConnectivityInfo) SetAvailabilityZonesNil(b bool)`

 SetAvailabilityZonesNil sets the value for AvailabilityZones to be an explicit nil

### UnsetAvailabilityZones
`func (o *BasicPrivateConnectivityInfo) UnsetAvailabilityZones()`

UnsetAvailabilityZones ensures that no value is present for AvailabilityZones, not even an explicit nil
### GetDomainNames

`func (o *BasicPrivateConnectivityInfo) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *BasicPrivateConnectivityInfo) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *BasicPrivateConnectivityInfo) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *BasicPrivateConnectivityInfo) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### SetDomainNamesNil

`func (o *BasicPrivateConnectivityInfo) SetDomainNamesNil(b bool)`

 SetDomainNamesNil sets the value for DomainNames to be an explicit nil

### UnsetDomainNames
`func (o *BasicPrivateConnectivityInfo) UnsetDomainNames()`

UnsetDomainNames ensures that no value is present for DomainNames, not even an explicit nil
### GetEndpointName

`func (o *BasicPrivateConnectivityInfo) GetEndpointName() string`

GetEndpointName returns the EndpointName field if non-nil, zero value otherwise.

### GetEndpointNameOk

`func (o *BasicPrivateConnectivityInfo) GetEndpointNameOk() (*string, bool)`

GetEndpointNameOk returns a tuple with the EndpointName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointName

`func (o *BasicPrivateConnectivityInfo) SetEndpointName(v string)`

SetEndpointName sets EndpointName field to given value.

### HasEndpointName

`func (o *BasicPrivateConnectivityInfo) HasEndpointName() bool`

HasEndpointName returns a boolean if a field has been set.

### GetPrivateDNS

`func (o *BasicPrivateConnectivityInfo) GetPrivateDNS() string`

GetPrivateDNS returns the PrivateDNS field if non-nil, zero value otherwise.

### GetPrivateDNSOk

`func (o *BasicPrivateConnectivityInfo) GetPrivateDNSOk() (*string, bool)`

GetPrivateDNSOk returns a tuple with the PrivateDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDNS

`func (o *BasicPrivateConnectivityInfo) SetPrivateDNS(v string)`

SetPrivateDNS sets PrivateDNS field to given value.

### HasPrivateDNS

`func (o *BasicPrivateConnectivityInfo) HasPrivateDNS() bool`

HasPrivateDNS returns a boolean if a field has been set.

### GetRegions

`func (o *BasicPrivateConnectivityInfo) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *BasicPrivateConnectivityInfo) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *BasicPrivateConnectivityInfo) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *BasicPrivateConnectivityInfo) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### SetRegionsNil

`func (o *BasicPrivateConnectivityInfo) SetRegionsNil(b bool)`

 SetRegionsNil sets the value for Regions to be an explicit nil

### UnsetRegions
`func (o *BasicPrivateConnectivityInfo) UnsetRegions()`

UnsetRegions ensures that no value is present for Regions, not even an explicit nil
### GetServiceAttachment

`func (o *BasicPrivateConnectivityInfo) GetServiceAttachment() string`

GetServiceAttachment returns the ServiceAttachment field if non-nil, zero value otherwise.

### GetServiceAttachmentOk

`func (o *BasicPrivateConnectivityInfo) GetServiceAttachmentOk() (*string, bool)`

GetServiceAttachmentOk returns a tuple with the ServiceAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAttachment

`func (o *BasicPrivateConnectivityInfo) SetServiceAttachment(v string)`

SetServiceAttachment sets ServiceAttachment field to given value.

### HasServiceAttachment

`func (o *BasicPrivateConnectivityInfo) HasServiceAttachment() bool`

HasServiceAttachment returns a boolean if a field has been set.

### GetServiceName

`func (o *BasicPrivateConnectivityInfo) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *BasicPrivateConnectivityInfo) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *BasicPrivateConnectivityInfo) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *BasicPrivateConnectivityInfo) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


