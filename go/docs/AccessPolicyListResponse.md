# AccessPolicyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | **[]map[string]interface{}** |  | 
**Metadata** | [**Metadata**](Metadata.md) |  | 

## Methods

### NewAccessPolicyListResponse

`func NewAccessPolicyListResponse(items []map[string]interface{}, metadata Metadata, ) *AccessPolicyListResponse`

NewAccessPolicyListResponse instantiates a new AccessPolicyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPolicyListResponseWithDefaults

`func NewAccessPolicyListResponseWithDefaults() *AccessPolicyListResponse`

NewAccessPolicyListResponseWithDefaults instantiates a new AccessPolicyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AccessPolicyListResponse) GetItems() []map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AccessPolicyListResponse) GetItemsOk() (*[]map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AccessPolicyListResponse) SetItems(v []map[string]interface{})`

SetItems sets Items field to given value.


### GetMetadata

`func (o *AccessPolicyListResponse) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccessPolicyListResponse) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccessPolicyListResponse) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


