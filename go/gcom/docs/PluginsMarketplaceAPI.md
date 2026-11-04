# \PluginsMarketplaceAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RenewMarketplaceLicense**](PluginsMarketplaceAPI.md#RenewMarketplaceLicense) | **Post** /plugins-marketplace/licenses/renew | Renew a marketplace license token.



## RenewMarketplaceLicense

> MarketplaceLicenseRenewalResponse RenewMarketplaceLicense(ctx).XRequestId(xRequestId).RenewMarketplaceLicenseRequest(renewMarketplaceLicenseRequest).Execute()

Renew a marketplace license token.

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
	xRequestId := "xRequestId_example" // string | 
	renewMarketplaceLicenseRequest := *openapiclient.NewRenewMarketplaceLicenseRequest("Token_example") // RenewMarketplaceLicenseRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PluginsMarketplaceAPI.RenewMarketplaceLicense(context.Background()).XRequestId(xRequestId).RenewMarketplaceLicenseRequest(renewMarketplaceLicenseRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PluginsMarketplaceAPI.RenewMarketplaceLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenewMarketplaceLicense`: MarketplaceLicenseRenewalResponse
	fmt.Fprintf(os.Stdout, "Response from `PluginsMarketplaceAPI.RenewMarketplaceLicense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRenewMarketplaceLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** |  | 
 **renewMarketplaceLicenseRequest** | [**RenewMarketplaceLicenseRequest**](RenewMarketplaceLicenseRequest.md) |  | 

### Return type

[**MarketplaceLicenseRenewalResponse**](MarketplaceLicenseRenewalResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

