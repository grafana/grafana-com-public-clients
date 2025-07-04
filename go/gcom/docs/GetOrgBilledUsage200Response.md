# GetOrgBilledUsage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]BilledUsage**](BilledUsage.md) |  | 
**OrderBy** | **string** |  | 
**Direction** | **string** |  | 
**Total** | **float32** |  | 
**Pages** | **float32** |  | 
**PageSize** | **float32** |  | 
**Page** | **float32** |  | 

## Methods

### NewGetOrgBilledUsage200Response

`func NewGetOrgBilledUsage200Response(items []BilledUsage, orderBy string, direction string, total float32, pages float32, pageSize float32, page float32, ) *GetOrgBilledUsage200Response`

NewGetOrgBilledUsage200Response instantiates a new GetOrgBilledUsage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrgBilledUsage200ResponseWithDefaults

`func NewGetOrgBilledUsage200ResponseWithDefaults() *GetOrgBilledUsage200Response`

NewGetOrgBilledUsage200ResponseWithDefaults instantiates a new GetOrgBilledUsage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetOrgBilledUsage200Response) GetItems() []BilledUsage`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetOrgBilledUsage200Response) GetItemsOk() (*[]BilledUsage, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetOrgBilledUsage200Response) SetItems(v []BilledUsage)`

SetItems sets Items field to given value.


### GetOrderBy

`func (o *GetOrgBilledUsage200Response) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *GetOrgBilledUsage200Response) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *GetOrgBilledUsage200Response) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.


### GetDirection

`func (o *GetOrgBilledUsage200Response) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetOrgBilledUsage200Response) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetOrgBilledUsage200Response) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetTotal

`func (o *GetOrgBilledUsage200Response) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetOrgBilledUsage200Response) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetOrgBilledUsage200Response) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetPages

`func (o *GetOrgBilledUsage200Response) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GetOrgBilledUsage200Response) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GetOrgBilledUsage200Response) SetPages(v float32)`

SetPages sets Pages field to given value.


### GetPageSize

`func (o *GetOrgBilledUsage200Response) GetPageSize() float32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GetOrgBilledUsage200Response) GetPageSizeOk() (*float32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GetOrgBilledUsage200Response) SetPageSize(v float32)`

SetPageSize sets PageSize field to given value.


### GetPage

`func (o *GetOrgBilledUsage200Response) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetOrgBilledUsage200Response) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetOrgBilledUsage200Response) SetPage(v float32)`

SetPage sets Page field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


