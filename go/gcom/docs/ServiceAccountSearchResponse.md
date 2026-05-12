# ServiceAccountSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **float32** |  | 
**ServiceAccounts** | [**[]ServiceAccountsInner**](ServiceAccountsInner.md) |  | 
**Page** | **float32** |  | 
**PerPage** | **float32** |  | 

## Methods

### NewServiceAccountSearchResponse

`func NewServiceAccountSearchResponse(totalCount float32, serviceAccounts []ServiceAccountsInner, page float32, perPage float32, ) *ServiceAccountSearchResponse`

NewServiceAccountSearchResponse instantiates a new ServiceAccountSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountSearchResponseWithDefaults

`func NewServiceAccountSearchResponseWithDefaults() *ServiceAccountSearchResponse`

NewServiceAccountSearchResponseWithDefaults instantiates a new ServiceAccountSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *ServiceAccountSearchResponse) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ServiceAccountSearchResponse) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ServiceAccountSearchResponse) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.


### GetServiceAccounts

`func (o *ServiceAccountSearchResponse) GetServiceAccounts() []ServiceAccountsInner`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *ServiceAccountSearchResponse) GetServiceAccountsOk() (*[]ServiceAccountsInner, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *ServiceAccountSearchResponse) SetServiceAccounts(v []ServiceAccountsInner)`

SetServiceAccounts sets ServiceAccounts field to given value.


### GetPage

`func (o *ServiceAccountSearchResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ServiceAccountSearchResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ServiceAccountSearchResponse) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPerPage

`func (o *ServiceAccountSearchResponse) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ServiceAccountSearchResponse) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ServiceAccountSearchResponse) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


