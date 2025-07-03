# UsagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSeries** | Pointer to **float32** |  | [optional] 
**ActiveUsers** | Pointer to **float32** |  | [optional] 
**AnonymousUsers** | Pointer to **float32** |  | [optional] 
**BrowserUsage** | Pointer to **float32** |  | [optional] 
**Dpm** | Pointer to **float32** |  | [optional] 
**GrafanaUsage** | Pointer to **float32** |  | [optional] 
**Id** | **float32** |  | 
**IncidentUsage** | Pointer to **float32** |  | [optional] 
**IngestUsage** | Pointer to **float32** |  | [optional] 
**IsProrated** | **bool** |  | 
**OnCallUsage** | Pointer to **float32** |  | [optional] 
**PeriodEnd** | **string** |  | 
**PeriodStart** | **string** |  | 
**ProtocolUsage** | Pointer to **float32** |  | [optional] 
**QueryUsage** | Pointer to **float32** |  | [optional] 
**StackId** | **float32** |  | 
**TotalUsage** | **float32** |  | 
**StackName** | **string** |  | 
**StackLabels** | **map[string]interface{}** |  | 

## Methods

### NewUsagesInner

`func NewUsagesInner(id float32, isProrated bool, periodEnd string, periodStart string, stackId float32, totalUsage float32, stackName string, stackLabels map[string]interface{}, ) *UsagesInner`

NewUsagesInner instantiates a new UsagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsagesInnerWithDefaults

`func NewUsagesInnerWithDefaults() *UsagesInner`

NewUsagesInnerWithDefaults instantiates a new UsagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSeries

`func (o *UsagesInner) GetActiveSeries() float32`

GetActiveSeries returns the ActiveSeries field if non-nil, zero value otherwise.

### GetActiveSeriesOk

`func (o *UsagesInner) GetActiveSeriesOk() (*float32, bool)`

GetActiveSeriesOk returns a tuple with the ActiveSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSeries

`func (o *UsagesInner) SetActiveSeries(v float32)`

SetActiveSeries sets ActiveSeries field to given value.

### HasActiveSeries

`func (o *UsagesInner) HasActiveSeries() bool`

HasActiveSeries returns a boolean if a field has been set.

### GetActiveUsers

`func (o *UsagesInner) GetActiveUsers() float32`

GetActiveUsers returns the ActiveUsers field if non-nil, zero value otherwise.

### GetActiveUsersOk

`func (o *UsagesInner) GetActiveUsersOk() (*float32, bool)`

GetActiveUsersOk returns a tuple with the ActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUsers

`func (o *UsagesInner) SetActiveUsers(v float32)`

SetActiveUsers sets ActiveUsers field to given value.

### HasActiveUsers

`func (o *UsagesInner) HasActiveUsers() bool`

HasActiveUsers returns a boolean if a field has been set.

### GetAnonymousUsers

`func (o *UsagesInner) GetAnonymousUsers() float32`

GetAnonymousUsers returns the AnonymousUsers field if non-nil, zero value otherwise.

### GetAnonymousUsersOk

`func (o *UsagesInner) GetAnonymousUsersOk() (*float32, bool)`

GetAnonymousUsersOk returns a tuple with the AnonymousUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousUsers

`func (o *UsagesInner) SetAnonymousUsers(v float32)`

SetAnonymousUsers sets AnonymousUsers field to given value.

### HasAnonymousUsers

`func (o *UsagesInner) HasAnonymousUsers() bool`

HasAnonymousUsers returns a boolean if a field has been set.

### GetBrowserUsage

`func (o *UsagesInner) GetBrowserUsage() float32`

GetBrowserUsage returns the BrowserUsage field if non-nil, zero value otherwise.

### GetBrowserUsageOk

`func (o *UsagesInner) GetBrowserUsageOk() (*float32, bool)`

GetBrowserUsageOk returns a tuple with the BrowserUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserUsage

`func (o *UsagesInner) SetBrowserUsage(v float32)`

SetBrowserUsage sets BrowserUsage field to given value.

### HasBrowserUsage

`func (o *UsagesInner) HasBrowserUsage() bool`

HasBrowserUsage returns a boolean if a field has been set.

### GetDpm

`func (o *UsagesInner) GetDpm() float32`

GetDpm returns the Dpm field if non-nil, zero value otherwise.

### GetDpmOk

`func (o *UsagesInner) GetDpmOk() (*float32, bool)`

GetDpmOk returns a tuple with the Dpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpm

`func (o *UsagesInner) SetDpm(v float32)`

SetDpm sets Dpm field to given value.

### HasDpm

`func (o *UsagesInner) HasDpm() bool`

HasDpm returns a boolean if a field has been set.

### GetGrafanaUsage

`func (o *UsagesInner) GetGrafanaUsage() float32`

GetGrafanaUsage returns the GrafanaUsage field if non-nil, zero value otherwise.

### GetGrafanaUsageOk

`func (o *UsagesInner) GetGrafanaUsageOk() (*float32, bool)`

GetGrafanaUsageOk returns a tuple with the GrafanaUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaUsage

`func (o *UsagesInner) SetGrafanaUsage(v float32)`

SetGrafanaUsage sets GrafanaUsage field to given value.

### HasGrafanaUsage

`func (o *UsagesInner) HasGrafanaUsage() bool`

HasGrafanaUsage returns a boolean if a field has been set.

### GetId

`func (o *UsagesInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsagesInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsagesInner) SetId(v float32)`

SetId sets Id field to given value.


### GetIncidentUsage

`func (o *UsagesInner) GetIncidentUsage() float32`

GetIncidentUsage returns the IncidentUsage field if non-nil, zero value otherwise.

### GetIncidentUsageOk

`func (o *UsagesInner) GetIncidentUsageOk() (*float32, bool)`

GetIncidentUsageOk returns a tuple with the IncidentUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidentUsage

`func (o *UsagesInner) SetIncidentUsage(v float32)`

SetIncidentUsage sets IncidentUsage field to given value.

### HasIncidentUsage

`func (o *UsagesInner) HasIncidentUsage() bool`

HasIncidentUsage returns a boolean if a field has been set.

### GetIngestUsage

`func (o *UsagesInner) GetIngestUsage() float32`

GetIngestUsage returns the IngestUsage field if non-nil, zero value otherwise.

### GetIngestUsageOk

`func (o *UsagesInner) GetIngestUsageOk() (*float32, bool)`

GetIngestUsageOk returns a tuple with the IngestUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestUsage

`func (o *UsagesInner) SetIngestUsage(v float32)`

SetIngestUsage sets IngestUsage field to given value.

### HasIngestUsage

`func (o *UsagesInner) HasIngestUsage() bool`

HasIngestUsage returns a boolean if a field has been set.

### GetIsProrated

`func (o *UsagesInner) GetIsProrated() bool`

GetIsProrated returns the IsProrated field if non-nil, zero value otherwise.

### GetIsProratedOk

`func (o *UsagesInner) GetIsProratedOk() (*bool, bool)`

GetIsProratedOk returns a tuple with the IsProrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProrated

`func (o *UsagesInner) SetIsProrated(v bool)`

SetIsProrated sets IsProrated field to given value.


### GetOnCallUsage

`func (o *UsagesInner) GetOnCallUsage() float32`

GetOnCallUsage returns the OnCallUsage field if non-nil, zero value otherwise.

### GetOnCallUsageOk

`func (o *UsagesInner) GetOnCallUsageOk() (*float32, bool)`

GetOnCallUsageOk returns a tuple with the OnCallUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCallUsage

`func (o *UsagesInner) SetOnCallUsage(v float32)`

SetOnCallUsage sets OnCallUsage field to given value.

### HasOnCallUsage

`func (o *UsagesInner) HasOnCallUsage() bool`

HasOnCallUsage returns a boolean if a field has been set.

### GetPeriodEnd

`func (o *UsagesInner) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *UsagesInner) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *UsagesInner) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetPeriodStart

`func (o *UsagesInner) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *UsagesInner) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *UsagesInner) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetProtocolUsage

`func (o *UsagesInner) GetProtocolUsage() float32`

GetProtocolUsage returns the ProtocolUsage field if non-nil, zero value otherwise.

### GetProtocolUsageOk

`func (o *UsagesInner) GetProtocolUsageOk() (*float32, bool)`

GetProtocolUsageOk returns a tuple with the ProtocolUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolUsage

`func (o *UsagesInner) SetProtocolUsage(v float32)`

SetProtocolUsage sets ProtocolUsage field to given value.

### HasProtocolUsage

`func (o *UsagesInner) HasProtocolUsage() bool`

HasProtocolUsage returns a boolean if a field has been set.

### GetQueryUsage

`func (o *UsagesInner) GetQueryUsage() float32`

GetQueryUsage returns the QueryUsage field if non-nil, zero value otherwise.

### GetQueryUsageOk

`func (o *UsagesInner) GetQueryUsageOk() (*float32, bool)`

GetQueryUsageOk returns a tuple with the QueryUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryUsage

`func (o *UsagesInner) SetQueryUsage(v float32)`

SetQueryUsage sets QueryUsage field to given value.

### HasQueryUsage

`func (o *UsagesInner) HasQueryUsage() bool`

HasQueryUsage returns a boolean if a field has been set.

### GetStackId

`func (o *UsagesInner) GetStackId() float32`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *UsagesInner) GetStackIdOk() (*float32, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *UsagesInner) SetStackId(v float32)`

SetStackId sets StackId field to given value.


### GetTotalUsage

`func (o *UsagesInner) GetTotalUsage() float32`

GetTotalUsage returns the TotalUsage field if non-nil, zero value otherwise.

### GetTotalUsageOk

`func (o *UsagesInner) GetTotalUsageOk() (*float32, bool)`

GetTotalUsageOk returns a tuple with the TotalUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsage

`func (o *UsagesInner) SetTotalUsage(v float32)`

SetTotalUsage sets TotalUsage field to given value.


### GetStackName

`func (o *UsagesInner) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *UsagesInner) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *UsagesInner) SetStackName(v string)`

SetStackName sets StackName field to given value.


### GetStackLabels

`func (o *UsagesInner) GetStackLabels() map[string]interface{}`

GetStackLabels returns the StackLabels field if non-nil, zero value otherwise.

### GetStackLabelsOk

`func (o *UsagesInner) GetStackLabelsOk() (*map[string]interface{}, bool)`

GetStackLabelsOk returns a tuple with the StackLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackLabels

`func (o *UsagesInner) SetStackLabels(v map[string]interface{})`

SetStackLabels sets StackLabels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


