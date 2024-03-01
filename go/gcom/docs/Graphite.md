# Graphite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateDNS** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 

## Methods

### NewGraphite

`func NewGraphite() *Graphite`

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

### HasPrivateDNS

`func (o *Graphite) HasPrivateDNS() bool`

HasPrivateDNS returns a boolean if a field has been set.

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

### HasServiceName

`func (o *Graphite) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


