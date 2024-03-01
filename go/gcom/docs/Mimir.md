# Mimir

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateDNS** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 

## Methods

### NewMimir

`func NewMimir() *Mimir`

NewMimir instantiates a new Mimir object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMimirWithDefaults

`func NewMimirWithDefaults() *Mimir`

NewMimirWithDefaults instantiates a new Mimir object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateDNS

`func (o *Mimir) GetPrivateDNS() string`

GetPrivateDNS returns the PrivateDNS field if non-nil, zero value otherwise.

### GetPrivateDNSOk

`func (o *Mimir) GetPrivateDNSOk() (*string, bool)`

GetPrivateDNSOk returns a tuple with the PrivateDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateDNS

`func (o *Mimir) SetPrivateDNS(v string)`

SetPrivateDNS sets PrivateDNS field to given value.

### HasPrivateDNS

`func (o *Mimir) HasPrivateDNS() bool`

HasPrivateDNS returns a boolean if a field has been set.

### GetServiceName

`func (o *Mimir) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *Mimir) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *Mimir) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *Mimir) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


