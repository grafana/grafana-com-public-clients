# PostInstanceServiceAccountTokensRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**SecondsToLive** | Pointer to **int32** |  | [optional] 

## Methods

### NewPostInstanceServiceAccountTokensRequest

`func NewPostInstanceServiceAccountTokensRequest(name string, ) *PostInstanceServiceAccountTokensRequest`

NewPostInstanceServiceAccountTokensRequest instantiates a new PostInstanceServiceAccountTokensRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostInstanceServiceAccountTokensRequestWithDefaults

`func NewPostInstanceServiceAccountTokensRequestWithDefaults() *PostInstanceServiceAccountTokensRequest`

NewPostInstanceServiceAccountTokensRequestWithDefaults instantiates a new PostInstanceServiceAccountTokensRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostInstanceServiceAccountTokensRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostInstanceServiceAccountTokensRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostInstanceServiceAccountTokensRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSecondsToLive

`func (o *PostInstanceServiceAccountTokensRequest) GetSecondsToLive() int32`

GetSecondsToLive returns the SecondsToLive field if non-nil, zero value otherwise.

### GetSecondsToLiveOk

`func (o *PostInstanceServiceAccountTokensRequest) GetSecondsToLiveOk() (*int32, bool)`

GetSecondsToLiveOk returns a tuple with the SecondsToLive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsToLive

`func (o *PostInstanceServiceAccountTokensRequest) SetSecondsToLive(v int32)`

SetSecondsToLive sets SecondsToLive field to given value.

### HasSecondsToLive

`func (o *PostInstanceServiceAccountTokensRequest) HasSecondsToLive() bool`

HasSecondsToLive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


