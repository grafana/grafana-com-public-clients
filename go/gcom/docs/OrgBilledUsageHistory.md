# OrgBilledUsageHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] 
**Items** | [**[]ItemsInner3**](ItemsInner3.md) |  | 

## Methods

### NewOrgBilledUsageHistory

`func NewOrgBilledUsageHistory(items []ItemsInner3, ) *OrgBilledUsageHistory`

NewOrgBilledUsageHistory instantiates a new OrgBilledUsageHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgBilledUsageHistoryWithDefaults

`func NewOrgBilledUsageHistoryWithDefaults() *OrgBilledUsageHistory`

NewOrgBilledUsageHistoryWithDefaults instantiates a new OrgBilledUsageHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *OrgBilledUsageHistory) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *OrgBilledUsageHistory) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *OrgBilledUsageHistory) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *OrgBilledUsageHistory) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetItems

`func (o *OrgBilledUsageHistory) GetItems() []ItemsInner3`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrgBilledUsageHistory) GetItemsOk() (*[]ItemsInner3, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrgBilledUsageHistory) SetItems(v []ItemsInner3)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


