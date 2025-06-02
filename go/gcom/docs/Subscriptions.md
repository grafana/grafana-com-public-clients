# Subscriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | [**Current2**](Current2.md) |  | 
**Next** | **interface{}** |  | 

## Methods

### NewSubscriptions

`func NewSubscriptions(current Current2, next interface{}, ) *Subscriptions`

NewSubscriptions instantiates a new Subscriptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionsWithDefaults

`func NewSubscriptionsWithDefaults() *Subscriptions`

NewSubscriptionsWithDefaults instantiates a new Subscriptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *Subscriptions) GetCurrent() Current2`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *Subscriptions) GetCurrentOk() (*Current2, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *Subscriptions) SetCurrent(v Current2)`

SetCurrent sets Current field to given value.


### GetNext

`func (o *Subscriptions) GetNext() interface{}`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *Subscriptions) GetNextOk() (*interface{}, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *Subscriptions) SetNext(v interface{})`

SetNext sets Next field to given value.


### SetNextNil

`func (o *Subscriptions) SetNextNil(b bool)`

 SetNextNil sets the value for Next to be an explicit nil

### UnsetNext
`func (o *Subscriptions) UnsetNext()`

UnsetNext ensures that no value is present for Next, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


