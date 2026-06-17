# StackConnectionOtlpV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MultiAZ** | **bool** | whether OTLP is multi-AZ | 
**PrivateConnectivityInfo** | Pointer to [**BasicPrivateConnectivityInfo**](BasicPrivateConnectivityInfo.md) |  | [optional] 
**Url** | Pointer to **string** | OTLP HTTP endpoint URL | [optional] 

## Methods

### NewStackConnectionOtlpV1

`func NewStackConnectionOtlpV1(multiAZ bool, ) *StackConnectionOtlpV1`

NewStackConnectionOtlpV1 instantiates a new StackConnectionOtlpV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStackConnectionOtlpV1WithDefaults

`func NewStackConnectionOtlpV1WithDefaults() *StackConnectionOtlpV1`

NewStackConnectionOtlpV1WithDefaults instantiates a new StackConnectionOtlpV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMultiAZ

`func (o *StackConnectionOtlpV1) GetMultiAZ() bool`

GetMultiAZ returns the MultiAZ field if non-nil, zero value otherwise.

### GetMultiAZOk

`func (o *StackConnectionOtlpV1) GetMultiAZOk() (*bool, bool)`

GetMultiAZOk returns a tuple with the MultiAZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiAZ

`func (o *StackConnectionOtlpV1) SetMultiAZ(v bool)`

SetMultiAZ sets MultiAZ field to given value.


### GetPrivateConnectivityInfo

`func (o *StackConnectionOtlpV1) GetPrivateConnectivityInfo() BasicPrivateConnectivityInfo`

GetPrivateConnectivityInfo returns the PrivateConnectivityInfo field if non-nil, zero value otherwise.

### GetPrivateConnectivityInfoOk

`func (o *StackConnectionOtlpV1) GetPrivateConnectivityInfoOk() (*BasicPrivateConnectivityInfo, bool)`

GetPrivateConnectivityInfoOk returns a tuple with the PrivateConnectivityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateConnectivityInfo

`func (o *StackConnectionOtlpV1) SetPrivateConnectivityInfo(v BasicPrivateConnectivityInfo)`

SetPrivateConnectivityInfo sets PrivateConnectivityInfo field to given value.

### HasPrivateConnectivityInfo

`func (o *StackConnectionOtlpV1) HasPrivateConnectivityInfo() bool`

HasPrivateConnectivityInfo returns a boolean if a field has been set.

### GetUrl

`func (o *StackConnectionOtlpV1) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StackConnectionOtlpV1) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StackConnectionOtlpV1) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StackConnectionOtlpV1) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


