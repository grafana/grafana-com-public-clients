# Current5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | **string** |  | 
**IsTrial** | **bool** |  | 
**StartDate** | **string** |  | 
**EndDate** | **interface{}** |  | 
**Payload** | [**Payload**](Payload.md) |  | 
**Plan** | **NullableString** |  | 
**PublicName** | **NullableString** |  | 
**EnterprisePluginsAdded** | **bool** |  | 
**PlanBillingCycle** | **string** |  | 

## Methods

### NewCurrent5

`func NewCurrent5(product string, isTrial bool, startDate string, endDate interface{}, payload Payload, plan NullableString, publicName NullableString, enterprisePluginsAdded bool, planBillingCycle string, ) *Current5`

NewCurrent5 instantiates a new Current5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrent5WithDefaults

`func NewCurrent5WithDefaults() *Current5`

NewCurrent5WithDefaults instantiates a new Current5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *Current5) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Current5) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Current5) SetProduct(v string)`

SetProduct sets Product field to given value.


### GetIsTrial

`func (o *Current5) GetIsTrial() bool`

GetIsTrial returns the IsTrial field if non-nil, zero value otherwise.

### GetIsTrialOk

`func (o *Current5) GetIsTrialOk() (*bool, bool)`

GetIsTrialOk returns a tuple with the IsTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrial

`func (o *Current5) SetIsTrial(v bool)`

SetIsTrial sets IsTrial field to given value.


### GetStartDate

`func (o *Current5) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Current5) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Current5) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *Current5) GetEndDate() interface{}`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Current5) GetEndDateOk() (*interface{}, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Current5) SetEndDate(v interface{})`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *Current5) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *Current5) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetPayload

`func (o *Current5) GetPayload() Payload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Current5) GetPayloadOk() (*Payload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Current5) SetPayload(v Payload)`

SetPayload sets Payload field to given value.


### GetPlan

`func (o *Current5) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Current5) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Current5) SetPlan(v string)`

SetPlan sets Plan field to given value.


### SetPlanNil

`func (o *Current5) SetPlanNil(b bool)`

 SetPlanNil sets the value for Plan to be an explicit nil

### UnsetPlan
`func (o *Current5) UnsetPlan()`

UnsetPlan ensures that no value is present for Plan, not even an explicit nil
### GetPublicName

`func (o *Current5) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *Current5) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *Current5) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.


### SetPublicNameNil

`func (o *Current5) SetPublicNameNil(b bool)`

 SetPublicNameNil sets the value for PublicName to be an explicit nil

### UnsetPublicName
`func (o *Current5) UnsetPublicName()`

UnsetPublicName ensures that no value is present for PublicName, not even an explicit nil
### GetEnterprisePluginsAdded

`func (o *Current5) GetEnterprisePluginsAdded() bool`

GetEnterprisePluginsAdded returns the EnterprisePluginsAdded field if non-nil, zero value otherwise.

### GetEnterprisePluginsAddedOk

`func (o *Current5) GetEnterprisePluginsAddedOk() (*bool, bool)`

GetEnterprisePluginsAddedOk returns a tuple with the EnterprisePluginsAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprisePluginsAdded

`func (o *Current5) SetEnterprisePluginsAdded(v bool)`

SetEnterprisePluginsAdded sets EnterprisePluginsAdded field to given value.


### GetPlanBillingCycle

`func (o *Current5) GetPlanBillingCycle() string`

GetPlanBillingCycle returns the PlanBillingCycle field if non-nil, zero value otherwise.

### GetPlanBillingCycleOk

`func (o *Current5) GetPlanBillingCycleOk() (*string, bool)`

GetPlanBillingCycleOk returns a tuple with the PlanBillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanBillingCycle

`func (o *Current5) SetPlanBillingCycle(v string)`

SetPlanBillingCycle sets PlanBillingCycle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


