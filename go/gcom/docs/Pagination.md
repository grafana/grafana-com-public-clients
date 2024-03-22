# Pagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageCursor** | Pointer to **NullableString** |  | [optional] 
**NextPage** | Pointer to **NullableString** |  | [optional] 
**PageSize** | **float32** |  | 

## Methods

### NewPagination

`func NewPagination(pageSize float32, ) *Pagination`

NewPagination instantiates a new Pagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationWithDefaults

`func NewPaginationWithDefaults() *Pagination`

NewPaginationWithDefaults instantiates a new Pagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageCursor

`func (o *Pagination) GetPageCursor() string`

GetPageCursor returns the PageCursor field if non-nil, zero value otherwise.

### GetPageCursorOk

`func (o *Pagination) GetPageCursorOk() (*string, bool)`

GetPageCursorOk returns a tuple with the PageCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCursor

`func (o *Pagination) SetPageCursor(v string)`

SetPageCursor sets PageCursor field to given value.

### HasPageCursor

`func (o *Pagination) HasPageCursor() bool`

HasPageCursor returns a boolean if a field has been set.

### SetPageCursorNil

`func (o *Pagination) SetPageCursorNil(b bool)`

 SetPageCursorNil sets the value for PageCursor to be an explicit nil

### UnsetPageCursor
`func (o *Pagination) UnsetPageCursor()`

UnsetPageCursor ensures that no value is present for PageCursor, not even an explicit nil
### GetNextPage

`func (o *Pagination) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *Pagination) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *Pagination) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *Pagination) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### SetNextPageNil

`func (o *Pagination) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *Pagination) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil
### GetPageSize

`func (o *Pagination) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Pagination) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Pagination) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


