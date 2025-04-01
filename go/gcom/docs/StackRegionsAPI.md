# \StackRegionsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetClosestStackRegion**](StackRegionsAPI.md#GetClosestStackRegion) | **Get** /stack-regions/closest | 
[**GetStackRegions**](StackRegionsAPI.md#GetStackRegions) | **Get** /stack-regions | 



## GetClosestStackRegion

> FormattedApiStackRegion GetClosestStackRegion(ctx).Provider(provider).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/grafana-com-public-clients/go/gcom"
)

func main() {
	provider := "provider_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackRegionsAPI.GetClosestStackRegion(context.Background()).Provider(provider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackRegionsAPI.GetClosestStackRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClosestStackRegion`: FormattedApiStackRegion
	fmt.Fprintf(os.Stdout, "Response from `StackRegionsAPI.GetClosestStackRegion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetClosestStackRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** |  | 

### Return type

[**FormattedApiStackRegion**](FormattedApiStackRegion.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackRegions

> GetStackRegions200Response GetStackRegions(ctx).CountryCode(countryCode).Direction(direction).Id(id).IdIn(idIn).OrderBy(orderBy).Provider(provider).ProviderRegion(providerRegion).Slug(slug).SlugIn(slugIn).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/grafana/grafana-com-public-clients/go/gcom"
)

func main() {
	countryCode := "countryCode_example" // string |  (optional)
	direction := "direction_example" // string |  (optional)
	id := int32(56) // int32 |  (optional)
	idIn := []int32{int32(123)} // []int32 |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	provider := "provider_example" // string |  (optional)
	providerRegion := "providerRegion_example" // string |  (optional)
	slug := "slug_example" // string |  (optional)
	slugIn := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StackRegionsAPI.GetStackRegions(context.Background()).CountryCode(countryCode).Direction(direction).Id(id).IdIn(idIn).OrderBy(orderBy).Provider(provider).ProviderRegion(providerRegion).Slug(slug).SlugIn(slugIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StackRegionsAPI.GetStackRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackRegions`: GetStackRegions200Response
	fmt.Fprintf(os.Stdout, "Response from `StackRegionsAPI.GetStackRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStackRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **countryCode** | **string** |  | 
 **direction** | **string** |  | 
 **id** | **int32** |  | 
 **idIn** | **[]int32** |  | 
 **orderBy** | **string** |  | 
 **provider** | **string** |  | 
 **providerRegion** | **string** |  | 
 **slug** | **string** |  | 
 **slugIn** | **[]string** |  | 

### Return type

[**GetStackRegions200Response**](GetStackRegions200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

