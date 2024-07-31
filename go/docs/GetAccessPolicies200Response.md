# GetAccessPolicies200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AuthAccessPolicy**](AuthAccessPolicy.md) |  | [optional] 

## Methods

### NewGetAccessPolicies200Response

`func NewGetAccessPolicies200Response() *GetAccessPolicies200Response`

NewGetAccessPolicies200Response instantiates a new GetAccessPolicies200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccessPolicies200ResponseWithDefaults

`func NewGetAccessPolicies200ResponseWithDefaults() *GetAccessPolicies200Response`

NewGetAccessPolicies200ResponseWithDefaults instantiates a new GetAccessPolicies200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetAccessPolicies200Response) GetItems() []AuthAccessPolicy`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetAccessPolicies200Response) GetItemsOk() (*[]AuthAccessPolicy, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetAccessPolicies200Response) SetItems(v []AuthAccessPolicy)`

SetItems sets Items field to given value.

### HasItems

`func (o *GetAccessPolicies200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


