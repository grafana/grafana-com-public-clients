# Current2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** |  | [optional] 
**IsTrial** | Pointer to **bool** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **interface{}** |  | [optional] 
**Payload** | Pointer to [**Payload**](Payload.md) |  | [optional] 
**Plan** | Pointer to **NullableString** |  | [optional] 
**PublicName** | Pointer to **NullableString** |  | [optional] 
**EnterprisePluginsAdded** | Pointer to **bool** |  | [optional] 
**PlanBillingCycle** | Pointer to **string** |  | [optional] 

## Methods

### NewCurrent2

`func NewCurrent2() *Current2`

NewCurrent2 instantiates a new Current2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrent2WithDefaults

`func NewCurrent2WithDefaults() *Current2`

NewCurrent2WithDefaults instantiates a new Current2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *Current2) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Current2) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Current2) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Current2) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetIsTrial

`func (o *Current2) GetIsTrial() bool`

GetIsTrial returns the IsTrial field if non-nil, zero value otherwise.

### GetIsTrialOk

`func (o *Current2) GetIsTrialOk() (*bool, bool)`

GetIsTrialOk returns a tuple with the IsTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrial

`func (o *Current2) SetIsTrial(v bool)`

SetIsTrial sets IsTrial field to given value.

### HasIsTrial

`func (o *Current2) HasIsTrial() bool`

HasIsTrial returns a boolean if a field has been set.

### GetStartDate

`func (o *Current2) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Current2) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Current2) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Current2) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Current2) GetEndDate() interface{}`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Current2) GetEndDateOk() (*interface{}, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Current2) SetEndDate(v interface{})`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Current2) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### SetEndDateNil

`func (o *Current2) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *Current2) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetPayload

`func (o *Current2) GetPayload() Payload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Current2) GetPayloadOk() (*Payload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Current2) SetPayload(v Payload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Current2) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPlan

`func (o *Current2) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Current2) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Current2) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Current2) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### SetPlanNil

`func (o *Current2) SetPlanNil(b bool)`

 SetPlanNil sets the value for Plan to be an explicit nil

### UnsetPlan
`func (o *Current2) UnsetPlan()`

UnsetPlan ensures that no value is present for Plan, not even an explicit nil
### GetPublicName

`func (o *Current2) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *Current2) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *Current2) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *Current2) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.

### SetPublicNameNil

`func (o *Current2) SetPublicNameNil(b bool)`

 SetPublicNameNil sets the value for PublicName to be an explicit nil

### UnsetPublicName
`func (o *Current2) UnsetPublicName()`

UnsetPublicName ensures that no value is present for PublicName, not even an explicit nil
### GetEnterprisePluginsAdded

`func (o *Current2) GetEnterprisePluginsAdded() bool`

GetEnterprisePluginsAdded returns the EnterprisePluginsAdded field if non-nil, zero value otherwise.

### GetEnterprisePluginsAddedOk

`func (o *Current2) GetEnterprisePluginsAddedOk() (*bool, bool)`

GetEnterprisePluginsAddedOk returns a tuple with the EnterprisePluginsAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprisePluginsAdded

`func (o *Current2) SetEnterprisePluginsAdded(v bool)`

SetEnterprisePluginsAdded sets EnterprisePluginsAdded field to given value.

### HasEnterprisePluginsAdded

`func (o *Current2) HasEnterprisePluginsAdded() bool`

HasEnterprisePluginsAdded returns a boolean if a field has been set.

### GetPlanBillingCycle

`func (o *Current2) GetPlanBillingCycle() string`

GetPlanBillingCycle returns the PlanBillingCycle field if non-nil, zero value otherwise.

### GetPlanBillingCycleOk

`func (o *Current2) GetPlanBillingCycleOk() (*string, bool)`

GetPlanBillingCycleOk returns a tuple with the PlanBillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanBillingCycle

`func (o *Current2) SetPlanBillingCycle(v string)`

SetPlanBillingCycle sets PlanBillingCycle field to given value.

### HasPlanBillingCycle

`func (o *Current2) HasPlanBillingCycle() bool`

HasPlanBillingCycle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


