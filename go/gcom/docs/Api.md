# Api

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateDNS** | **string** |  | 
**ServiceName** | **string** |  | 

## Methods

### NewApi

`func NewApi(privateDNS string, serviceName string, ) *Api`

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


