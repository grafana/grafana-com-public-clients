# PostAccessPoliciesRequestAttributesLokiQueryPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxEntriesLimitPerQuery** | Pointer to **NullableInt32** |  | [optional] 
**MaxQueryBytesRead** | Pointer to **NullableString** |  | [optional] 
**MaxQueryInterval** | Pointer to **NullableString** |  | [optional] 
**MaxQueryLength** | Pointer to **NullableString** |  | [optional] 
**MaxQueryLookback** | Pointer to **NullableString** |  | [optional] 
**MaxQueryTime** | Pointer to **NullableString** |  | [optional] 
**MinimumLabelsNumber** | Pointer to **NullableInt32** |  | [optional] 
**RequiredLabels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPostAccessPoliciesRequestAttributesLokiQueryPolicy

`func NewPostAccessPoliciesRequestAttributesLokiQueryPolicy() *PostAccessPoliciesRequestAttributesLokiQueryPolicy`

NewPostAccessPoliciesRequestAttributesLokiQueryPolicy instantiates a new PostAccessPoliciesRequestAttributesLokiQueryPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAccessPoliciesRequestAttributesLokiQueryPolicyWithDefaults

`func NewPostAccessPoliciesRequestAttributesLokiQueryPolicyWithDefaults() *PostAccessPoliciesRequestAttributesLokiQueryPolicy`

NewPostAccessPoliciesRequestAttributesLokiQueryPolicyWithDefaults instantiates a new PostAccessPoliciesRequestAttributesLokiQueryPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxEntriesLimitPerQuery

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMaxEntriesLimitPerQuery() int32`

GetMaxEntriesLimitPerQuery returns the MaxEntriesLimitPerQuery field if non-nil, zero value otherwise.

### GetMaxEntriesLimitPerQueryOk

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMaxEntriesLimitPerQueryOk() (*int32, bool)`

GetMaxEntriesLimitPerQueryOk returns a tuple with the MaxEntriesLimitPerQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxEntriesLimitPerQuery

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMaxEntriesLimitPerQuery(v int32)`

SetMaxEntriesLimitPerQuery sets MaxEntriesLimitPerQuery field to given value.

### HasMaxEntriesLimitPerQuery

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) HasMaxEntriesLimitPerQuery() bool`

HasMaxEntriesLimitPerQuery returns a boolean if a field has been set.

### SetMaxEntriesLimitPerQueryNil

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMaxEntriesLimitPerQueryNil(b bool)`

 SetMaxEntriesLimitPerQueryNil sets the value for MaxEntriesLimitPerQuery to be an explicit nil

### UnsetMaxEntriesLimitPerQuery
`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) UnsetMaxEntriesLimitPerQuery()`

UnsetMaxEntriesLimitPerQuery ensures that no value is present for MaxEntriesLimitPerQuery, not even an explicit nil
### GetMaxQueryBytesRead

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMaxQueryBytesRead() string`

GetMaxQueryBytesRead returns the MaxQueryBytesRead field if non-nil, zero value otherwise.

### GetMaxQueryBytesReadOk

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMaxQueryBytesReadOk() (*string, bool)`

GetMaxQueryBytesReadOk returns a tuple with the MaxQueryBytesRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueryBytesRead

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMaxQueryBytesRead(v string)`

SetMaxQueryBytesRead sets MaxQueryBytesRead field to given value.

### HasMaxQueryBytesRead

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) HasMaxQueryBytesRead() bool`

HasMaxQueryBytesRead returns a boolean if a field has been set.

### SetMaxQueryBytesReadNil

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMaxQueryBytesReadNil(b bool)`

 SetMaxQueryBytesReadNil sets the value for MaxQueryBytesRead to be an explicit nil

### UnsetMaxQueryBytesRead
`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) UnsetMaxQueryBytesRead()`

UnsetMaxQueryBytesRead ensures that no value is present for MaxQueryBytesRead, not even an explicit nil
### GetMaxQueryInterval

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMaxQueryInterval() string`

GetMaxQueryInterval returns the MaxQueryInterval field if non-nil, zero value otherwise.

### GetMaxQueryIntervalOk

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMaxQueryIntervalOk() (*string, bool)`

GetMaxQueryIntervalOk returns a tuple with the MaxQueryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueryInterval

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMaxQueryInterval(v string)`

SetMaxQueryInterval sets MaxQueryInterval field to given value.

### HasMaxQueryInterval

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) HasMaxQueryInterval() bool`

HasMaxQueryInterval returns a boolean if a field has been set.

### SetMaxQueryIntervalNil

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMaxQueryIntervalNil(b bool)`

 SetMaxQueryIntervalNil sets the value for MaxQueryInterval to be an explicit nil

### UnsetMaxQueryInterval
`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) UnsetMaxQueryInterval()`

UnsetMaxQueryInterval ensures that no value is present for MaxQueryInterval, not even an explicit nil
### GetMaxQueryLength

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMaxQueryLength() string`

GetMaxQueryLength returns the MaxQueryLength field if non-nil, zero value otherwise.

### GetMaxQueryLengthOk

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMaxQueryLengthOk() (*string, bool)`

GetMaxQueryLengthOk returns a tuple with the MaxQueryLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueryLength

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMaxQueryLength(v string)`

SetMaxQueryLength sets MaxQueryLength field to given value.

### HasMaxQueryLength

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) HasMaxQueryLength() bool`

HasMaxQueryLength returns a boolean if a field has been set.

### SetMaxQueryLengthNil

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMaxQueryLengthNil(b bool)`

 SetMaxQueryLengthNil sets the value for MaxQueryLength to be an explicit nil

### UnsetMaxQueryLength
`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) UnsetMaxQueryLength()`

UnsetMaxQueryLength ensures that no value is present for MaxQueryLength, not even an explicit nil
### GetMaxQueryLookback

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMaxQueryLookback() string`

GetMaxQueryLookback returns the MaxQueryLookback field if non-nil, zero value otherwise.

### GetMaxQueryLookbackOk

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMaxQueryLookbackOk() (*string, bool)`

GetMaxQueryLookbackOk returns a tuple with the MaxQueryLookback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueryLookback

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMaxQueryLookback(v string)`

SetMaxQueryLookback sets MaxQueryLookback field to given value.

### HasMaxQueryLookback

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) HasMaxQueryLookback() bool`

HasMaxQueryLookback returns a boolean if a field has been set.

### SetMaxQueryLookbackNil

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMaxQueryLookbackNil(b bool)`

 SetMaxQueryLookbackNil sets the value for MaxQueryLookback to be an explicit nil

### UnsetMaxQueryLookback
`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) UnsetMaxQueryLookback()`

UnsetMaxQueryLookback ensures that no value is present for MaxQueryLookback, not even an explicit nil
### GetMaxQueryTime

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMaxQueryTime() string`

GetMaxQueryTime returns the MaxQueryTime field if non-nil, zero value otherwise.

### GetMaxQueryTimeOk

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMaxQueryTimeOk() (*string, bool)`

GetMaxQueryTimeOk returns a tuple with the MaxQueryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxQueryTime

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMaxQueryTime(v string)`

SetMaxQueryTime sets MaxQueryTime field to given value.

### HasMaxQueryTime

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) HasMaxQueryTime() bool`

HasMaxQueryTime returns a boolean if a field has been set.

### SetMaxQueryTimeNil

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMaxQueryTimeNil(b bool)`

 SetMaxQueryTimeNil sets the value for MaxQueryTime to be an explicit nil

### UnsetMaxQueryTime
`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) UnsetMaxQueryTime()`

UnsetMaxQueryTime ensures that no value is present for MaxQueryTime, not even an explicit nil
### GetMinimumLabelsNumber

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMinimumLabelsNumber() int32`

GetMinimumLabelsNumber returns the MinimumLabelsNumber field if non-nil, zero value otherwise.

### GetMinimumLabelsNumberOk

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetMinimumLabelsNumberOk() (*int32, bool)`

GetMinimumLabelsNumberOk returns a tuple with the MinimumLabelsNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumLabelsNumber

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMinimumLabelsNumber(v int32)`

SetMinimumLabelsNumber sets MinimumLabelsNumber field to given value.

### HasMinimumLabelsNumber

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) HasMinimumLabelsNumber() bool`

HasMinimumLabelsNumber returns a boolean if a field has been set.

### SetMinimumLabelsNumberNil

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetMinimumLabelsNumberNil(b bool)`

 SetMinimumLabelsNumberNil sets the value for MinimumLabelsNumber to be an explicit nil

### UnsetMinimumLabelsNumber
`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) UnsetMinimumLabelsNumber()`

UnsetMinimumLabelsNumber ensures that no value is present for MinimumLabelsNumber, not even an explicit nil
### GetRequiredLabels

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetRequiredLabels() []string`

GetRequiredLabels returns the RequiredLabels field if non-nil, zero value otherwise.

### GetRequiredLabelsOk

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) GetRequiredLabelsOk() (*[]string, bool)`

GetRequiredLabelsOk returns a tuple with the RequiredLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredLabels

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetRequiredLabels(v []string)`

SetRequiredLabels sets RequiredLabels field to given value.

### HasRequiredLabels

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) HasRequiredLabels() bool`

HasRequiredLabels returns a boolean if a field has been set.

### SetRequiredLabelsNil

`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) SetRequiredLabelsNil(b bool)`

 SetRequiredLabelsNil sets the value for RequiredLabels to be an explicit nil

### UnsetRequiredLabels
`func (o *PostAccessPoliciesRequestAttributesLokiQueryPolicy) UnsetRequiredLabels()`

UnsetRequiredLabels ensures that no value is present for RequiredLabels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


