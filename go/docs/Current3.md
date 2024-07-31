# Current3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | **interface{}** |  | 
**IsTrial** | **bool** |  | 
**StartDate** | **interface{}** |  | 
**EndDate** | **interface{}** |  | 
**Payload** | **map[string]interface{}** |  | 
**Plan** | **interface{}** |  | 
**PublicName** | **interface{}** |  | 
**EnterprisePluginsAdded** | **bool** |  | 
**PlanBillingCycle** | **interface{}** |  | 

## Methods

### NewCurrent3

`func NewCurrent3(product interface{}, isTrial bool, startDate interface{}, endDate interface{}, payload map[string]interface{}, plan interface{}, publicName interface{}, enterprisePluginsAdded bool, planBillingCycle interface{}, ) *Current3`

NewCurrent3 instantiates a new Current3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrent3WithDefaults

`func NewCurrent3WithDefaults() *Current3`

NewCurrent3WithDefaults instantiates a new Current3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *Current3) GetProduct() interface{}`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Current3) GetProductOk() (*interface{}, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Current3) SetProduct(v interface{})`

SetProduct sets Product field to given value.


### SetProductNil

`func (o *Current3) SetProductNil(b bool)`

 SetProductNil sets the value for Product to be an explicit nil

### UnsetProduct
`func (o *Current3) UnsetProduct()`

UnsetProduct ensures that no value is present for Product, not even an explicit nil
### GetIsTrial

`func (o *Current3) GetIsTrial() bool`

GetIsTrial returns the IsTrial field if non-nil, zero value otherwise.

### GetIsTrialOk

`func (o *Current3) GetIsTrialOk() (*bool, bool)`

GetIsTrialOk returns a tuple with the IsTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrial

`func (o *Current3) SetIsTrial(v bool)`

SetIsTrial sets IsTrial field to given value.


### GetStartDate

`func (o *Current3) GetStartDate() interface{}`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Current3) GetStartDateOk() (*interface{}, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Current3) SetStartDate(v interface{})`

SetStartDate sets StartDate field to given value.


### SetStartDateNil

`func (o *Current3) SetStartDateNil(b bool)`

 SetStartDateNil sets the value for StartDate to be an explicit nil

### UnsetStartDate
`func (o *Current3) UnsetStartDate()`

UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
### GetEndDate

`func (o *Current3) GetEndDate() interface{}`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Current3) GetEndDateOk() (*interface{}, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Current3) SetEndDate(v interface{})`

SetEndDate sets EndDate field to given value.


### SetEndDateNil

`func (o *Current3) SetEndDateNil(b bool)`

 SetEndDateNil sets the value for EndDate to be an explicit nil

### UnsetEndDate
`func (o *Current3) UnsetEndDate()`

UnsetEndDate ensures that no value is present for EndDate, not even an explicit nil
### GetPayload

`func (o *Current3) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Current3) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Current3) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.


### GetPlan

`func (o *Current3) GetPlan() interface{}`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Current3) GetPlanOk() (*interface{}, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Current3) SetPlan(v interface{})`

SetPlan sets Plan field to given value.


### SetPlanNil

`func (o *Current3) SetPlanNil(b bool)`

 SetPlanNil sets the value for Plan to be an explicit nil

### UnsetPlan
`func (o *Current3) UnsetPlan()`

UnsetPlan ensures that no value is present for Plan, not even an explicit nil
### GetPublicName

`func (o *Current3) GetPublicName() interface{}`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *Current3) GetPublicNameOk() (*interface{}, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *Current3) SetPublicName(v interface{})`

SetPublicName sets PublicName field to given value.


### SetPublicNameNil

`func (o *Current3) SetPublicNameNil(b bool)`

 SetPublicNameNil sets the value for PublicName to be an explicit nil

### UnsetPublicName
`func (o *Current3) UnsetPublicName()`

UnsetPublicName ensures that no value is present for PublicName, not even an explicit nil
### GetEnterprisePluginsAdded

`func (o *Current3) GetEnterprisePluginsAdded() bool`

GetEnterprisePluginsAdded returns the EnterprisePluginsAdded field if non-nil, zero value otherwise.

### GetEnterprisePluginsAddedOk

`func (o *Current3) GetEnterprisePluginsAddedOk() (*bool, bool)`

GetEnterprisePluginsAddedOk returns a tuple with the EnterprisePluginsAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprisePluginsAdded

`func (o *Current3) SetEnterprisePluginsAdded(v bool)`

SetEnterprisePluginsAdded sets EnterprisePluginsAdded field to given value.


### GetPlanBillingCycle

`func (o *Current3) GetPlanBillingCycle() interface{}`

GetPlanBillingCycle returns the PlanBillingCycle field if non-nil, zero value otherwise.

### GetPlanBillingCycleOk

`func (o *Current3) GetPlanBillingCycleOk() (*interface{}, bool)`

GetPlanBillingCycleOk returns a tuple with the PlanBillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanBillingCycle

`func (o *Current3) SetPlanBillingCycle(v interface{})`

SetPlanBillingCycle sets PlanBillingCycle field to given value.


### SetPlanBillingCycleNil

`func (o *Current3) SetPlanBillingCycleNil(b bool)`

 SetPlanBillingCycleNil sets the value for PlanBillingCycle to be an explicit nil

### UnsetPlanBillingCycle
`func (o *Current3) UnsetPlanBillingCycle()`

UnsetPlanBillingCycle ensures that no value is present for PlanBillingCycle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


