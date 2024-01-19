# FormattedApiApiKeyListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ItemsInner**](ItemsInner.md) |  | 
**OrderBy** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **float32** |  | [optional] 
**Pages** | Pointer to **float32** |  | [optional] 
**PageSize** | Pointer to **float32** |  | [optional] 
**Page** | Pointer to **float32** |  | [optional] 
**Cursor** | Pointer to **float32** |  | [optional] 
**NextCursor** | Pointer to **float32** |  | [optional] 
**Links** | Pointer to [**[]LinksInner1**](LinksInner1.md) |  | [optional] 

## Methods

### NewFormattedApiApiKeyListResponse

`func NewFormattedApiApiKeyListResponse(items []ItemsInner, ) *FormattedApiApiKeyListResponse`

NewFormattedApiApiKeyListResponse instantiates a new FormattedApiApiKeyListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormattedApiApiKeyListResponseWithDefaults

`func NewFormattedApiApiKeyListResponseWithDefaults() *FormattedApiApiKeyListResponse`

NewFormattedApiApiKeyListResponseWithDefaults instantiates a new FormattedApiApiKeyListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *FormattedApiApiKeyListResponse) GetItems() []ItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FormattedApiApiKeyListResponse) GetItemsOk() (*[]ItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FormattedApiApiKeyListResponse) SetItems(v []ItemsInner)`

SetItems sets Items field to given value.


### GetOrderBy

`func (o *FormattedApiApiKeyListResponse) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *FormattedApiApiKeyListResponse) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *FormattedApiApiKeyListResponse) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *FormattedApiApiKeyListResponse) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetDirection

`func (o *FormattedApiApiKeyListResponse) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FormattedApiApiKeyListResponse) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FormattedApiApiKeyListResponse) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *FormattedApiApiKeyListResponse) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTotal

`func (o *FormattedApiApiKeyListResponse) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FormattedApiApiKeyListResponse) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FormattedApiApiKeyListResponse) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FormattedApiApiKeyListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPages

`func (o *FormattedApiApiKeyListResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *FormattedApiApiKeyListResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *FormattedApiApiKeyListResponse) SetPages(v float32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *FormattedApiApiKeyListResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetPageSize

`func (o *FormattedApiApiKeyListResponse) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *FormattedApiApiKeyListResponse) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *FormattedApiApiKeyListResponse) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *FormattedApiApiKeyListResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPage

`func (o *FormattedApiApiKeyListResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *FormattedApiApiKeyListResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *FormattedApiApiKeyListResponse) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *FormattedApiApiKeyListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetCursor

`func (o *FormattedApiApiKeyListResponse) GetCursor() float32`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *FormattedApiApiKeyListResponse) GetCursorOk() (*float32, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *FormattedApiApiKeyListResponse) SetCursor(v float32)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *FormattedApiApiKeyListResponse) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetNextCursor

`func (o *FormattedApiApiKeyListResponse) GetNextCursor() float32`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *FormattedApiApiKeyListResponse) GetNextCursorOk() (*float32, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *FormattedApiApiKeyListResponse) SetNextCursor(v float32)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *FormattedApiApiKeyListResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetLinks

`func (o *FormattedApiApiKeyListResponse) GetLinks() []LinksInner1`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FormattedApiApiKeyListResponse) GetLinksOk() (*[]LinksInner1, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FormattedApiApiKeyListResponse) SetLinks(v []LinksInner1)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FormattedApiApiKeyListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


