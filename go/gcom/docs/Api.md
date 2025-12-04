# Api

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** |  | 
**PrivateDNS** | **string** |  | 
**Regions** | Pointer to **[]string** |  | [optional] 
**ServiceAttachment** | **string** |  | 
**DomainNames** | **[]string** |  | 

## Methods

### NewApi

`func NewApi(serviceName string, privateDNS string, serviceAttachment string, domainNames []string, ) *Api`

NewApi instantiates a new Api object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWithDefaults

`func NewApiWithDefaults() *Api`

NewApiWithDefaults instantiates a new Api object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


