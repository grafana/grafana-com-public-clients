# ServiceAccountSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int64** | the current page number | 
**PerPage** | **int64** | the number of service accounts per page | 
**ServiceAccounts** | [**[]ServiceAccountDTO**](ServiceAccountDTO.md) | service accounts on this page | 
**TotalCount** | **int64** | total number of service accounts matching the query | 

## Methods

### NewServiceAccountSearchResponse

`func NewServiceAccountSearchResponse(page int64, perPage int64, serviceAccounts []ServiceAccountDTO, totalCount int64, ) *ServiceAccountSearchResponse`

NewServiceAccountSearchResponse instantiates a new ServiceAccountSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceAccountSearchResponseWithDefaults

`func NewServiceAccountSearchResponseWithDefaults() *ServiceAccountSearchResponse`

NewServiceAccountSearchResponseWithDefaults instantiates a new ServiceAccountSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ServiceAccountSearchResponse) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ServiceAccountSearchResponse) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ServiceAccountSearchResponse) SetPage(v int64)`

SetPage sets Page field to given value.


### GetPerPage

`func (o *ServiceAccountSearchResponse) GetPerPage() int64`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ServiceAccountSearchResponse) GetPerPageOk() (*int64, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ServiceAccountSearchResponse) SetPerPage(v int64)`

SetPerPage sets PerPage field to given value.


### GetServiceAccounts

`func (o *ServiceAccountSearchResponse) GetServiceAccounts() []ServiceAccountDTO`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *ServiceAccountSearchResponse) GetServiceAccountsOk() (*[]ServiceAccountDTO, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *ServiceAccountSearchResponse) SetServiceAccounts(v []ServiceAccountDTO)`

SetServiceAccounts sets ServiceAccounts field to given value.


### GetTotalCount

`func (o *ServiceAccountSearchResponse) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ServiceAccountSearchResponse) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ServiceAccountSearchResponse) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


