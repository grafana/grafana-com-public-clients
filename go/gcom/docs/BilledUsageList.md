# BilledUsageList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **float32** | The total number of items. | 
**Page** | **float32** | The page number of the items. | 
**PageSize** | **float32** | The page size of the items. | 
**Schema** | Pointer to **string** | A URL to the JSON Schema for this object. | [optional] 
**Items** | [**[]ItemsInner3**](ItemsInner3.md) |  | 

## Methods

### NewBilledUsageList

`func NewBilledUsageList(total float32, page float32, pageSize float32, items []ItemsInner3, ) *BilledUsageList`

NewBilledUsageList instantiates a new BilledUsageList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBilledUsageListWithDefaults

`func NewBilledUsageListWithDefaults() *BilledUsageList`

NewBilledUsageListWithDefaults instantiates a new BilledUsageList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *BilledUsageList) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BilledUsageList) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BilledUsageList) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetPage

`func (o *BilledUsageList) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *BilledUsageList) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *BilledUsageList) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *BilledUsageList) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *BilledUsageList) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *BilledUsageList) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.


### GetSchema

`func (o *BilledUsageList) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *BilledUsageList) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *BilledUsageList) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *BilledUsageList) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetItems

`func (o *BilledUsageList) GetItems() []ItemsInner3`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BilledUsageList) GetItemsOk() (*[]ItemsInner3, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BilledUsageList) SetItems(v []ItemsInner3)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


