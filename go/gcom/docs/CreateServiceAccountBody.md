# CreateServiceAccountBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDisabled** | Pointer to [**CreateServiceAccountBodyIsDisabled**](CreateServiceAccountBodyIsDisabled.md) |  | [optional] 
**Name** | **string** | service account name | 
**Role** | **string** | the service account&#39;s role | 

## Methods

### NewCreateServiceAccountBody

`func NewCreateServiceAccountBody(name string, role string, ) *CreateServiceAccountBody`

NewCreateServiceAccountBody instantiates a new CreateServiceAccountBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceAccountBodyWithDefaults

`func NewCreateServiceAccountBodyWithDefaults() *CreateServiceAccountBody`

NewCreateServiceAccountBodyWithDefaults instantiates a new CreateServiceAccountBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDisabled

`func (o *CreateServiceAccountBody) GetIsDisabled() CreateServiceAccountBodyIsDisabled`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *CreateServiceAccountBody) GetIsDisabledOk() (*CreateServiceAccountBodyIsDisabled, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *CreateServiceAccountBody) SetIsDisabled(v CreateServiceAccountBodyIsDisabled)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *CreateServiceAccountBody) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetName

`func (o *CreateServiceAccountBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceAccountBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceAccountBody) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *CreateServiceAccountBody) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateServiceAccountBody) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateServiceAccountBody) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


