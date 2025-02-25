# FormattedApiInstanceConnections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateConnectivityInfo** | [**PrivateConnectivityInfo**](PrivateConnectivityInfo.md) |  | 
**AppPlatform** | Pointer to [**AppPlatform**](AppPlatform.md) |  | [optional] 
**InfluxUrl** | Pointer to **NullableString** |  | [optional] 
**OtlpHttpUrl** | Pointer to **NullableString** |  | [optional] 
**OncallApiUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFormattedApiInstanceConnections

`func NewFormattedApiInstanceConnections(privateConnectivityInfo PrivateConnectivityInfo, ) *FormattedApiInstanceConnections`

NewFormattedApiInstanceConnections instantiates a new FormattedApiInstanceConnections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedApiInstanceConnectionsWithDefaults

`func NewFormattedApiInstanceConnectionsWithDefaults() *FormattedApiInstanceConnections`

NewFormattedApiInstanceConnectionsWithDefaults instantiates a new FormattedApiInstanceConnections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateConnectivityInfo

`func (o *FormattedApiInstanceConnections) GetPrivateConnectivityInfo() PrivateConnectivityInfo`

GetPrivateConnectivityInfo returns the PrivateConnectivityInfo field if non-nil, zero value otherwise.

### GetPrivateConnectivityInfoOk

`func (o *FormattedApiInstanceConnections) GetPrivateConnectivityInfoOk() (*PrivateConnectivityInfo, bool)`

GetPrivateConnectivityInfoOk returns a tuple with the PrivateConnectivityInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateConnectivityInfo

`func (o *FormattedApiInstanceConnections) SetPrivateConnectivityInfo(v PrivateConnectivityInfo)`

SetPrivateConnectivityInfo sets PrivateConnectivityInfo field to given value.


### GetAppPlatform

`func (o *FormattedApiInstanceConnections) GetAppPlatform() AppPlatform`

GetAppPlatform returns the AppPlatform field if non-nil, zero value otherwise.

### GetAppPlatformOk

`func (o *FormattedApiInstanceConnections) GetAppPlatformOk() (*AppPlatform, bool)`

GetAppPlatformOk returns a tuple with the AppPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPlatform

`func (o *FormattedApiInstanceConnections) SetAppPlatform(v AppPlatform)`

SetAppPlatform sets AppPlatform field to given value.

### HasAppPlatform

`func (o *FormattedApiInstanceConnections) HasAppPlatform() bool`

HasAppPlatform returns a boolean if a field has been set.

### GetInfluxUrl

`func (o *FormattedApiInstanceConnections) GetInfluxUrl() string`

GetInfluxUrl returns the InfluxUrl field if non-nil, zero value otherwise.

### GetInfluxUrlOk

`func (o *FormattedApiInstanceConnections) GetInfluxUrlOk() (*string, bool)`

GetInfluxUrlOk returns a tuple with the InfluxUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfluxUrl

`func (o *FormattedApiInstanceConnections) SetInfluxUrl(v string)`

SetInfluxUrl sets InfluxUrl field to given value.

### HasInfluxUrl

`func (o *FormattedApiInstanceConnections) HasInfluxUrl() bool`

HasInfluxUrl returns a boolean if a field has been set.

### SetInfluxUrlNil

`func (o *FormattedApiInstanceConnections) SetInfluxUrlNil(b bool)`

 SetInfluxUrlNil sets the value for InfluxUrl to be an explicit nil

### UnsetInfluxUrl
`func (o *FormattedApiInstanceConnections) UnsetInfluxUrl()`

UnsetInfluxUrl ensures that no value is present for InfluxUrl, not even an explicit nil
### GetOtlpHttpUrl

`func (o *FormattedApiInstanceConnections) GetOtlpHttpUrl() string`

GetOtlpHttpUrl returns the OtlpHttpUrl field if non-nil, zero value otherwise.

### GetOtlpHttpUrlOk

`func (o *FormattedApiInstanceConnections) GetOtlpHttpUrlOk() (*string, bool)`

GetOtlpHttpUrlOk returns a tuple with the OtlpHttpUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtlpHttpUrl

`func (o *FormattedApiInstanceConnections) SetOtlpHttpUrl(v string)`

SetOtlpHttpUrl sets OtlpHttpUrl field to given value.

### HasOtlpHttpUrl

`func (o *FormattedApiInstanceConnections) HasOtlpHttpUrl() bool`

HasOtlpHttpUrl returns a boolean if a field has been set.

### SetOtlpHttpUrlNil

`func (o *FormattedApiInstanceConnections) SetOtlpHttpUrlNil(b bool)`

 SetOtlpHttpUrlNil sets the value for OtlpHttpUrl to be an explicit nil

### UnsetOtlpHttpUrl
`func (o *FormattedApiInstanceConnections) UnsetOtlpHttpUrl()`

UnsetOtlpHttpUrl ensures that no value is present for OtlpHttpUrl, not even an explicit nil
### GetOncallApiUrl

`func (o *FormattedApiInstanceConnections) GetOncallApiUrl() string`

GetOncallApiUrl returns the OncallApiUrl field if non-nil, zero value otherwise.

### GetOncallApiUrlOk

`func (o *FormattedApiInstanceConnections) GetOncallApiUrlOk() (*string, bool)`

GetOncallApiUrlOk returns a tuple with the OncallApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOncallApiUrl

`func (o *FormattedApiInstanceConnections) SetOncallApiUrl(v string)`

SetOncallApiUrl sets OncallApiUrl field to given value.

### HasOncallApiUrl

`func (o *FormattedApiInstanceConnections) HasOncallApiUrl() bool`

HasOncallApiUrl returns a boolean if a field has been set.

### SetOncallApiUrlNil

`func (o *FormattedApiInstanceConnections) SetOncallApiUrlNil(b bool)`

 SetOncallApiUrlNil sets the value for OncallApiUrl to be an explicit nil

### UnsetOncallApiUrl
`func (o *FormattedApiInstanceConnections) UnsetOncallApiUrl()`

UnsetOncallApiUrl ensures that no value is present for OncallApiUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


