# NewApiKeyResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | service account token id | 
**Key** | **string** | the generated service account token secret | 
**Name** | **string** | service account token name | 

## Methods

### NewNewApiKeyResult

`func NewNewApiKeyResult(id int64, key string, name string, ) *NewApiKeyResult`

NewNewApiKeyResult instantiates a new NewApiKeyResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewApiKeyResultWithDefaults

`func NewNewApiKeyResultWithDefaults() *NewApiKeyResult`

NewNewApiKeyResultWithDefaults instantiates a new NewApiKeyResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NewApiKeyResult) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NewApiKeyResult) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NewApiKeyResult) SetId(v int64)`

SetId sets Id field to given value.


### GetKey

`func (o *NewApiKeyResult) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NewApiKeyResult) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NewApiKeyResult) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *NewApiKeyResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewApiKeyResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewApiKeyResult) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


