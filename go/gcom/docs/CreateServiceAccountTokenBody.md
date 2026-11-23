# CreateServiceAccountTokenBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | service account token name | 
**SecondsToLive** | Pointer to **int64** | seconds until expiry; omit for no expiry | [optional] 

## Methods

### NewCreateServiceAccountTokenBody

`func NewCreateServiceAccountTokenBody(name string, ) *CreateServiceAccountTokenBody`

NewCreateServiceAccountTokenBody instantiates a new CreateServiceAccountTokenBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceAccountTokenBodyWithDefaults

`func NewCreateServiceAccountTokenBodyWithDefaults() *CreateServiceAccountTokenBody`

NewCreateServiceAccountTokenBodyWithDefaults instantiates a new CreateServiceAccountTokenBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateServiceAccountTokenBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceAccountTokenBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceAccountTokenBody) SetName(v string)`

SetName sets Name field to given value.


### GetSecondsToLive

`func (o *CreateServiceAccountTokenBody) GetSecondsToLive() int64`

GetSecondsToLive returns the SecondsToLive field if non-nil, zero value otherwise.

### GetSecondsToLiveOk

`func (o *CreateServiceAccountTokenBody) GetSecondsToLiveOk() (*int64, bool)`

GetSecondsToLiveOk returns a tuple with the SecondsToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsToLive

`func (o *CreateServiceAccountTokenBody) SetSecondsToLive(v int64)`

SetSecondsToLive sets SecondsToLive field to given value.

### HasSecondsToLive

`func (o *CreateServiceAccountTokenBody) HasSecondsToLive() bool`

HasSecondsToLive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


