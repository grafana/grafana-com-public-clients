# Managed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**Strategy** | Pointer to **string** |  | [optional] 

## Methods

### NewManaged

`func NewManaged(enabled bool, ) *Managed`

NewManaged instantiates a new Managed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedWithDefaults

`func NewManagedWithDefaults() *Managed`

NewManagedWithDefaults instantiates a new Managed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *Managed) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Managed) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Managed) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetStrategy

`func (o *Managed) GetStrategy() string`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *Managed) GetStrategyOk() (*string, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *Managed) SetStrategy(v string)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *Managed) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


