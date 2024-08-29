# OrgMemberListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]ItemsInner2**](ItemsInner2.md) |  | 
**Links** | [**[]LinksInner1**](LinksInner1.md) |  | 
**Total** | Pointer to **float32** |  | [optional] 
**Pages** | Pointer to **float32** |  | [optional] 
**Page** | Pointer to **float32** |  | [optional] 
**PerPage** | Pointer to **float32** |  | [optional] 
**TotalCount** | Pointer to **float32** |  | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Cursor** | Pointer to **float32** |  | [optional] 
**NextCursor** | Pointer to **float32** |  | [optional] 

## Methods

### NewOrgMemberListResponse

`func NewOrgMemberListResponse(items []ItemsInner2, links []LinksInner1, ) *OrgMemberListResponse`

NewOrgMemberListResponse instantiates a new OrgMemberListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgMemberListResponseWithDefaults

`func NewOrgMemberListResponseWithDefaults() *OrgMemberListResponse`

NewOrgMemberListResponseWithDefaults instantiates a new OrgMemberListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *OrgMemberListResponse) GetItems() []ItemsInner2`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrgMemberListResponse) GetItemsOk() (*[]ItemsInner2, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrgMemberListResponse) SetItems(v []ItemsInner2)`

SetItems sets Items field to given value.


### GetLinks

`func (o *OrgMemberListResponse) GetLinks() []LinksInner1`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OrgMemberListResponse) GetLinksOk() (*[]LinksInner1, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OrgMemberListResponse) SetLinks(v []LinksInner1)`

SetLinks sets Links field to given value.


### GetTotal

`func (o *OrgMemberListResponse) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *OrgMemberListResponse) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *OrgMemberListResponse) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *OrgMemberListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetPages

`func (o *OrgMemberListResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *OrgMemberListResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *OrgMemberListResponse) SetPages(v float32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *OrgMemberListResponse) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetPage

`func (o *OrgMemberListResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *OrgMemberListResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *OrgMemberListResponse) SetPage(v float32)`

SetPage sets Page field to given value.

### HasPage

`func (o *OrgMemberListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPerPage

`func (o *OrgMemberListResponse) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *OrgMemberListResponse) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *OrgMemberListResponse) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.

### HasPerPage

`func (o *OrgMemberListResponse) HasPerPage() bool`

HasPerPage returns a boolean if a field has been set.

### GetTotalCount

`func (o *OrgMemberListResponse) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OrgMemberListResponse) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OrgMemberListResponse) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OrgMemberListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetOrderBy

`func (o *OrgMemberListResponse) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *OrgMemberListResponse) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *OrgMemberListResponse) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *OrgMemberListResponse) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetDirection

`func (o *OrgMemberListResponse) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *OrgMemberListResponse) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *OrgMemberListResponse) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *OrgMemberListResponse) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetCursor

`func (o *OrgMemberListResponse) GetCursor() float32`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *OrgMemberListResponse) GetCursorOk() (*float32, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *OrgMemberListResponse) SetCursor(v float32)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *OrgMemberListResponse) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetNextCursor

`func (o *OrgMemberListResponse) GetNextCursor() float32`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *OrgMemberListResponse) GetNextCursorOk() (*float32, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *OrgMemberListResponse) SetNextCursor(v float32)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *OrgMemberListResponse) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


