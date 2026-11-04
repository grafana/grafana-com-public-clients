# InstanceUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]InstanceUser**](InstanceUser.md) | the instance&#39;s users | 

## Methods

### NewInstanceUsersResponse

`func NewInstanceUsersResponse(items []InstanceUser, ) *InstanceUsersResponse`

NewInstanceUsersResponse instantiates a new InstanceUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceUsersResponseWithDefaults

`func NewInstanceUsersResponseWithDefaults() *InstanceUsersResponse`

NewInstanceUsersResponseWithDefaults instantiates a new InstanceUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *InstanceUsersResponse) GetItems() []InstanceUser`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *InstanceUsersResponse) GetItemsOk() (*[]InstanceUser, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *InstanceUsersResponse) SetItems(v []InstanceUser)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *InstanceUsersResponse) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *InstanceUsersResponse) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


