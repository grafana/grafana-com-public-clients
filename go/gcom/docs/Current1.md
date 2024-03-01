# Current1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to **string** |  | [optional] 
**IsTrial** | Pointer to **bool** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Payload** | Pointer to [**Payload**](Payload.md) |  | [optional] 
**Plan** | Pointer to **NullableString** |  | [optional] 
**PublicName** | Pointer to **NullableString** |  | [optional] 
**EnterprisePluginsAdded** | Pointer to **bool** |  | [optional] 
**PlanBillingCycle** | Pointer to **string** |  | [optional] 

## Methods

### NewCurrent1

`func NewCurrent1() *Current1`

NewCurrent1 instantiates a new Current1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrent1WithDefaults

`func NewCurrent1WithDefaults() *Current1`

NewCurrent1WithDefaults instantiates a new Current1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *Current1) GetProduct() string`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Current1) GetProductOk() (*string, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Current1) SetProduct(v string)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *Current1) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetIsTrial

`func (o *Current1) GetIsTrial() bool`

GetIsTrial returns the IsTrial field if non-nil, zero value otherwise.

### GetIsTrialOk

`func (o *Current1) GetIsTrialOk() (*bool, bool)`

GetIsTrialOk returns a tuple with the IsTrial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTrial

`func (o *Current1) SetIsTrial(v bool)`

SetIsTrial sets IsTrial field to given value.

### HasIsTrial

`func (o *Current1) HasIsTrial() bool`

HasIsTrial returns a boolean if a field has been set.

### GetStartDate

`func (o *Current1) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Current1) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Current1) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Current1) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *Current1) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Current1) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Current1) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Current1) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetPayload

`func (o *Current1) GetPayload() Payload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Current1) GetPayloadOk() (*Payload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Current1) SetPayload(v Payload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *Current1) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPlan

`func (o *Current1) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Current1) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Current1) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Current1) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### SetPlanNil

`func (o *Current1) SetPlanNil(b bool)`

 SetPlanNil sets the value for Plan to be an explicit nil

### UnsetPlan
`func (o *Current1) UnsetPlan()`

UnsetPlan ensures that no value is present for Plan, not even an explicit nil
### GetPublicName

`func (o *Current1) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *Current1) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *Current1) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.

### HasPublicName

`func (o *Current1) HasPublicName() bool`

HasPublicName returns a boolean if a field has been set.

### SetPublicNameNil

`func (o *Current1) SetPublicNameNil(b bool)`

 SetPublicNameNil sets the value for PublicName to be an explicit nil

### UnsetPublicName
`func (o *Current1) UnsetPublicName()`

UnsetPublicName ensures that no value is present for PublicName, not even an explicit nil
### GetEnterprisePluginsAdded

`func (o *Current1) GetEnterprisePluginsAdded() bool`

GetEnterprisePluginsAdded returns the EnterprisePluginsAdded field if non-nil, zero value otherwise.

### GetEnterprisePluginsAddedOk

`func (o *Current1) GetEnterprisePluginsAddedOk() (*bool, bool)`

GetEnterprisePluginsAddedOk returns a tuple with the EnterprisePluginsAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterprisePluginsAdded

`func (o *Current1) SetEnterprisePluginsAdded(v bool)`

SetEnterprisePluginsAdded sets EnterprisePluginsAdded field to given value.

### HasEnterprisePluginsAdded

`func (o *Current1) HasEnterprisePluginsAdded() bool`

HasEnterprisePluginsAdded returns a boolean if a field has been set.

### GetPlanBillingCycle

`func (o *Current1) GetPlanBillingCycle() string`

GetPlanBillingCycle returns the PlanBillingCycle field if non-nil, zero value otherwise.

### GetPlanBillingCycleOk

`func (o *Current1) GetPlanBillingCycleOk() (*string, bool)`

GetPlanBillingCycleOk returns a tuple with the PlanBillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanBillingCycle

`func (o *Current1) SetPlanBillingCycle(v string)`

SetPlanBillingCycle sets PlanBillingCycle field to given value.

### HasPlanBillingCycle

`func (o *Current1) HasPlanBillingCycle() bool`

HasPlanBillingCycle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


