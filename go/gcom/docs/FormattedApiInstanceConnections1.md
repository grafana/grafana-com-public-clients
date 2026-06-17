# FormattedApiInstanceConnections1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppPlatform** | Pointer to [**StackConnectionAppPlatformV1**](StackConnectionAppPlatformV1.md) |  | [optional] 
**InfluxUrl** | Pointer to **string** |  | [optional] 
**OncallApiUrl** | Pointer to **string** |  | [optional] 
**OtlpHttpUrl** | **NullableString** |  | 
**OtlpMultiAZ** | **bool** |  | 
**PrivateConnectivityInfo** | [**FormattedApiInstancePrivateConnectivityInfo**](FormattedApiInstancePrivateConnectivityInfo.md) |  | 
**SyntheticMonitoringApiUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewFormattedApiInstanceConnections1

`func NewFormattedApiInstanceConnections1(otlpHttpUrl NullableString, otlpMultiAZ bool, privateConnectivityInfo FormattedApiInstancePrivateConnectivityInfo, ) *FormattedApiInstanceConnections1`

NewFormattedApiInstanceConnections1 instantiates a new FormattedApiInstanceConnections1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedApiInstanceConnections1WithDefaults

`func NewFormattedApiInstanceConnections1WithDefaults() *FormattedApiInstanceConnections1`

NewFormattedApiInstanceConnections1WithDefaults instantiates a new FormattedApiInstanceConnections1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppPlatform

`func (o *FormattedApiInstanceConnections1) GetAppPlatform() StackConnectionAppPlatformV1`

GetAppPlatform returns the AppPlatform field if non-nil, zero value otherwise.

### GetAppPlatformOk

`func (o *FormattedApiInstanceConnections1) GetAppPlatformOk() (*StackConnectionAppPlatformV1, bool)`

GetAppPlatformOk returns a tuple with the AppPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPlatform

`func (o *FormattedApiInstanceConnections1) SetAppPlatform(v StackConnectionAppPlatformV1)`

SetAppPlatform sets AppPlatform field to given value.

### HasAppPlatform

`func (o *FormattedApiInstanceConnections1) HasAppPlatform() bool`

HasAppPlatform returns a boolean if a field has been set.

### GetInfluxUrl

`func (o *FormattedApiInstanceConnections1) GetInfluxUrl() string`

GetInfluxUrl returns the InfluxUrl field if non-nil, zero value otherwise.

### GetInfluxUrlOk

`func (o *FormattedApiInstanceConnections1) GetInfluxUrlOk() (*string, bool)`

GetInfluxUrlOk returns a tuple with the InfluxUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluxUrl

`func (o *FormattedApiInstanceConnections1) SetInfluxUrl(v string)`

SetInfluxUrl sets InfluxUrl field to given value.

### HasInfluxUrl

`func (o *FormattedApiInstanceConnections1) HasInfluxUrl() bool`

HasInfluxUrl returns a boolean if a field has been set.

### GetOncallApiUrl

`func (o *FormattedApiInstanceConnections1) GetOncallApiUrl() string`

GetOncallApiUrl returns the OncallApiUrl field if non-nil, zero value otherwise.

### GetOncallApiUrlOk

`func (o *FormattedApiInstanceConnections1) GetOncallApiUrlOk() (*string, bool)`

GetOncallApiUrlOk returns a tuple with the OncallApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOncallApiUrl

`func (o *FormattedApiInstanceConnections1) SetOncallApiUrl(v string)`

SetOncallApiUrl sets OncallApiUrl field to given value.

### HasOncallApiUrl

`func (o *FormattedApiInstanceConnections1) HasOncallApiUrl() bool`

HasOncallApiUrl returns a boolean if a field has been set.

### GetOtlpHttpUrl

`func (o *FormattedApiInstanceConnections1) GetOtlpHttpUrl() string`

GetOtlpHttpUrl returns the OtlpHttpUrl field if non-nil, zero value otherwise.

### GetOtlpHttpUrlOk

`func (o *FormattedApiInstanceConnections1) GetOtlpHttpUrlOk() (*string, bool)`

GetOtlpHttpUrlOk returns a tuple with the OtlpHttpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlpHttpUrl

`func (o *FormattedApiInstanceConnections1) SetOtlpHttpUrl(v string)`

SetOtlpHttpUrl sets OtlpHttpUrl field to given value.


### SetOtlpHttpUrlNil

`func (o *FormattedApiInstanceConnections1) SetOtlpHttpUrlNil(b bool)`

 SetOtlpHttpUrlNil sets the value for OtlpHttpUrl to be an explicit nil

### UnsetOtlpHttpUrl
`func (o *FormattedApiInstanceConnections1) UnsetOtlpHttpUrl()`

UnsetOtlpHttpUrl ensures that no value is present for OtlpHttpUrl, not even an explicit nil
### GetOtlpMultiAZ

`func (o *FormattedApiInstanceConnections1) GetOtlpMultiAZ() bool`

GetOtlpMultiAZ returns the OtlpMultiAZ field if non-nil, zero value otherwise.

### GetOtlpMultiAZOk

`func (o *FormattedApiInstanceConnections1) GetOtlpMultiAZOk() (*bool, bool)`

GetOtlpMultiAZOk returns a tuple with the OtlpMultiAZ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlpMultiAZ

`func (o *FormattedApiInstanceConnections1) SetOtlpMultiAZ(v bool)`

SetOtlpMultiAZ sets OtlpMultiAZ field to given value.


### GetPrivateConnectivityInfo

`func (o *FormattedApiInstanceConnections1) GetPrivateConnectivityInfo() FormattedApiInstancePrivateConnectivityInfo`

GetPrivateConnectivityInfo returns the PrivateConnectivityInfo field if non-nil, zero value otherwise.

### GetPrivateConnectivityInfoOk

`func (o *FormattedApiInstanceConnections1) GetPrivateConnectivityInfoOk() (*FormattedApiInstancePrivateConnectivityInfo, bool)`

GetPrivateConnectivityInfoOk returns a tuple with the PrivateConnectivityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateConnectivityInfo

`func (o *FormattedApiInstanceConnections1) SetPrivateConnectivityInfo(v FormattedApiInstancePrivateConnectivityInfo)`

SetPrivateConnectivityInfo sets PrivateConnectivityInfo field to given value.


### GetSyntheticMonitoringApiUrl

`func (o *FormattedApiInstanceConnections1) GetSyntheticMonitoringApiUrl() string`

GetSyntheticMonitoringApiUrl returns the SyntheticMonitoringApiUrl field if non-nil, zero value otherwise.

### GetSyntheticMonitoringApiUrlOk

`func (o *FormattedApiInstanceConnections1) GetSyntheticMonitoringApiUrlOk() (*string, bool)`

GetSyntheticMonitoringApiUrlOk returns a tuple with the SyntheticMonitoringApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntheticMonitoringApiUrl

`func (o *FormattedApiInstanceConnections1) SetSyntheticMonitoringApiUrl(v string)`

SetSyntheticMonitoringApiUrl sets SyntheticMonitoringApiUrl field to given value.

### HasSyntheticMonitoringApiUrl

`func (o *FormattedApiInstanceConnections1) HasSyntheticMonitoringApiUrl() bool`

HasSyntheticMonitoringApiUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


