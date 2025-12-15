# PaginatedResponseStackV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]StackV1**](StackV1.md) | items requested for the current page | 
**Page** | **int64** | index of the current page (starting on 1) | 
**PageSize** | **int64** | applied page size | 
**Pages** | **int64** | number of pages | 
**Total** | **int64** | total of amount of items matching the request | 

## Methods

### NewPaginatedResponseStackV1

`func NewPaginatedResponseStackV1(items []StackV1, page int64, pageSize int64, pages int64, total int64, ) *PaginatedResponseStackV1`

NewPaginatedResponseStackV1 instantiates a new PaginatedResponseStackV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResponseStackV1WithDefaults

`func NewPaginatedResponseStackV1WithDefaults() *PaginatedResponseStackV1`

NewPaginatedResponseStackV1WithDefaults instantiates a new PaginatedResponseStackV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *PaginatedResponseStackV1) GetItems() []StackV1`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PaginatedResponseStackV1) GetItemsOk() (*[]StackV1, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PaginatedResponseStackV1) SetItems(v []StackV1)`

SetItems sets Items field to given value.


### SetItemsNil

`func (o *PaginatedResponseStackV1) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *PaginatedResponseStackV1) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetPage

`func (o *PaginatedResponseStackV1) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PaginatedResponseStackV1) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PaginatedResponseStackV1) SetPage(v int64)`

SetPage sets Page field to given value.


### GetPageSize

`func (o *PaginatedResponseStackV1) GetPageSize() int64`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PaginatedResponseStackV1) GetPageSizeOk() (*int64, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PaginatedResponseStackV1) SetPageSize(v int64)`

SetPageSize sets PageSize field to given value.


### GetPages

`func (o *PaginatedResponseStackV1) GetPages() int64`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *PaginatedResponseStackV1) GetPagesOk() (*int64, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *PaginatedResponseStackV1) SetPages(v int64)`

SetPages sets Pages field to given value.


### GetTotal

`func (o *PaginatedResponseStackV1) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginatedResponseStackV1) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginatedResponseStackV1) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


