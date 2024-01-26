# PostInstanceServiceAccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDisabled** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**Role** | **string** |  | 

## Methods

### NewPostInstanceServiceAccountsRequest

`func NewPostInstanceServiceAccountsRequest(name string, role string, ) *PostInstanceServiceAccountsRequest`

NewPostInstanceServiceAccountsRequest instantiates a new PostInstanceServiceAccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostInstanceServiceAccountsRequestWithDefaults

`func NewPostInstanceServiceAccountsRequestWithDefaults() *PostInstanceServiceAccountsRequest`

NewPostInstanceServiceAccountsRequestWithDefaults instantiates a new PostInstanceServiceAccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDisabled

`func (o *PostInstanceServiceAccountsRequest) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *PostInstanceServiceAccountsRequest) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *PostInstanceServiceAccountsRequest) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *PostInstanceServiceAccountsRequest) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetName

`func (o *PostInstanceServiceAccountsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostInstanceServiceAccountsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostInstanceServiceAccountsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *PostInstanceServiceAccountsRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PostInstanceServiceAccountsRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PostInstanceServiceAccountsRequest) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


