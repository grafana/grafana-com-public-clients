# QueryOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxDataPoints** | Pointer to **bool** | For data source plugins. If the &#x60;max data points&#x60; option should be shown in the query options section in the query editor. | [optional] 
**MinInterval** | Pointer to **bool** | For data source plugins. If the &#x60;min interval&#x60; option should be shown in the query options section in the query editor. | [optional] 
**CacheTimeout** | Pointer to **bool** | For data source plugins. If the &#x60;cache timeout&#x60; option should be shown in the query options section in the query editor. | [optional] 

## Methods

### NewQueryOptions

`func NewQueryOptions() *QueryOptions`

NewQueryOptions instantiates a new QueryOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryOptionsWithDefaults

`func NewQueryOptionsWithDefaults() *QueryOptions`

NewQueryOptionsWithDefaults instantiates a new QueryOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxDataPoints

`func (o *QueryOptions) GetMaxDataPoints() bool`

GetMaxDataPoints returns the MaxDataPoints field if non-nil, zero value otherwise.

### GetMaxDataPointsOk

`func (o *QueryOptions) GetMaxDataPointsOk() (*bool, bool)`

GetMaxDataPointsOk returns a tuple with the MaxDataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDataPoints

`func (o *QueryOptions) SetMaxDataPoints(v bool)`

SetMaxDataPoints sets MaxDataPoints field to given value.

### HasMaxDataPoints

`func (o *QueryOptions) HasMaxDataPoints() bool`

HasMaxDataPoints returns a boolean if a field has been set.

### GetMinInterval

`func (o *QueryOptions) GetMinInterval() bool`

GetMinInterval returns the MinInterval field if non-nil, zero value otherwise.

### GetMinIntervalOk

`func (o *QueryOptions) GetMinIntervalOk() (*bool, bool)`

GetMinIntervalOk returns a tuple with the MinInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinInterval

`func (o *QueryOptions) SetMinInterval(v bool)`

SetMinInterval sets MinInterval field to given value.

### HasMinInterval

`func (o *QueryOptions) HasMinInterval() bool`

HasMinInterval returns a boolean if a field has been set.

### GetCacheTimeout

`func (o *QueryOptions) GetCacheTimeout() bool`

GetCacheTimeout returns the CacheTimeout field if non-nil, zero value otherwise.

### GetCacheTimeoutOk

`func (o *QueryOptions) GetCacheTimeoutOk() (*bool, bool)`

GetCacheTimeoutOk returns a tuple with the CacheTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTimeout

`func (o *QueryOptions) SetCacheTimeout(v bool)`

SetCacheTimeout sets CacheTimeout field to given value.

### HasCacheTimeout

`func (o *QueryOptions) HasCacheTimeout() bool`

HasCacheTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


