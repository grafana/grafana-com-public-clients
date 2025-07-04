# BilledUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Notes** | **string** |  | 
**Description** | **string** |  | 
**Unit** | **string** |  | 
**Overage** | **float32** |  | 
**IncludedUsage** | **float32** |  | 
**AmountDue** | **float32** |  | 
**DimensionId** | **string** |  | 
**DimensionName** | **string** |  | 
**OrgRates** | **interface{}** |  | 
**PeriodEnd** | **string** |  | 
**PeriodStart** | **string** |  | 
**TotalUsage** | **float32** |  | 
**Usages** | [**[]UsagesInner**](UsagesInner.md) |  | 

## Methods

### NewBilledUsage

`func NewBilledUsage(id float32, notes string, description string, unit string, overage float32, includedUsage float32, amountDue float32, dimensionId string, dimensionName string, orgRates interface{}, periodEnd string, periodStart string, totalUsage float32, usages []UsagesInner, ) *BilledUsage`

NewBilledUsage instantiates a new BilledUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBilledUsageWithDefaults

`func NewBilledUsageWithDefaults() *BilledUsage`

NewBilledUsageWithDefaults instantiates a new BilledUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BilledUsage) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BilledUsage) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BilledUsage) SetId(v float32)`

SetId sets Id field to given value.


### GetNotes

`func (o *BilledUsage) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *BilledUsage) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *BilledUsage) SetNotes(v string)`

SetNotes sets Notes field to given value.


### GetDescription

`func (o *BilledUsage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BilledUsage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BilledUsage) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUnit

`func (o *BilledUsage) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *BilledUsage) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *BilledUsage) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetOverage

`func (o *BilledUsage) GetOverage() float32`

GetOverage returns the Overage field if non-nil, zero value otherwise.

### GetOverageOk

`func (o *BilledUsage) GetOverageOk() (*float32, bool)`

GetOverageOk returns a tuple with the Overage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverage

`func (o *BilledUsage) SetOverage(v float32)`

SetOverage sets Overage field to given value.


### GetIncludedUsage

`func (o *BilledUsage) GetIncludedUsage() float32`

GetIncludedUsage returns the IncludedUsage field if non-nil, zero value otherwise.

### GetIncludedUsageOk

`func (o *BilledUsage) GetIncludedUsageOk() (*float32, bool)`

GetIncludedUsageOk returns a tuple with the IncludedUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedUsage

`func (o *BilledUsage) SetIncludedUsage(v float32)`

SetIncludedUsage sets IncludedUsage field to given value.


### GetAmountDue

`func (o *BilledUsage) GetAmountDue() float32`

GetAmountDue returns the AmountDue field if non-nil, zero value otherwise.

### GetAmountDueOk

`func (o *BilledUsage) GetAmountDueOk() (*float32, bool)`

GetAmountDueOk returns a tuple with the AmountDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountDue

`func (o *BilledUsage) SetAmountDue(v float32)`

SetAmountDue sets AmountDue field to given value.


### GetDimensionId

`func (o *BilledUsage) GetDimensionId() string`

GetDimensionId returns the DimensionId field if non-nil, zero value otherwise.

### GetDimensionIdOk

`func (o *BilledUsage) GetDimensionIdOk() (*string, bool)`

GetDimensionIdOk returns a tuple with the DimensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionId

`func (o *BilledUsage) SetDimensionId(v string)`

SetDimensionId sets DimensionId field to given value.


### GetDimensionName

`func (o *BilledUsage) GetDimensionName() string`

GetDimensionName returns the DimensionName field if non-nil, zero value otherwise.

### GetDimensionNameOk

`func (o *BilledUsage) GetDimensionNameOk() (*string, bool)`

GetDimensionNameOk returns a tuple with the DimensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionName

`func (o *BilledUsage) SetDimensionName(v string)`

SetDimensionName sets DimensionName field to given value.


### GetOrgRates

`func (o *BilledUsage) GetOrgRates() interface{}`

GetOrgRates returns the OrgRates field if non-nil, zero value otherwise.

### GetOrgRatesOk

`func (o *BilledUsage) GetOrgRatesOk() (*interface{}, bool)`

GetOrgRatesOk returns a tuple with the OrgRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgRates

`func (o *BilledUsage) SetOrgRates(v interface{})`

SetOrgRates sets OrgRates field to given value.


### SetOrgRatesNil

`func (o *BilledUsage) SetOrgRatesNil(b bool)`

 SetOrgRatesNil sets the value for OrgRates to be an explicit nil

### UnsetOrgRates
`func (o *BilledUsage) UnsetOrgRates()`

UnsetOrgRates ensures that no value is present for OrgRates, not even an explicit nil
### GetPeriodEnd

`func (o *BilledUsage) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *BilledUsage) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *BilledUsage) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetPeriodStart

`func (o *BilledUsage) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *BilledUsage) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *BilledUsage) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetTotalUsage

`func (o *BilledUsage) GetTotalUsage() float32`

GetTotalUsage returns the TotalUsage field if non-nil, zero value otherwise.

### GetTotalUsageOk

`func (o *BilledUsage) GetTotalUsageOk() (*float32, bool)`

GetTotalUsageOk returns a tuple with the TotalUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUsage

`func (o *BilledUsage) SetTotalUsage(v float32)`

SetTotalUsage sets TotalUsage field to given value.


### GetUsages

`func (o *BilledUsage) GetUsages() []UsagesInner`

GetUsages returns the Usages field if non-nil, zero value otherwise.

### GetUsagesOk

`func (o *BilledUsage) GetUsagesOk() (*[]UsagesInner, bool)`

GetUsagesOk returns a tuple with the Usages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsages

`func (o *BilledUsage) SetUsages(v []UsagesInner)`

SetUsages sets Usages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


