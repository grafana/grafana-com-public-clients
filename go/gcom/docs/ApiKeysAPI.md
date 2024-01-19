# \ApiKeysAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostAllApiKeys**](ApiKeysAPI.md#PostAllApiKeys) | **Post** /api-keys | Creates an API key



## PostAllApiKeys

> FormattedApiApiKey PostAllApiKeys(ctx).XRequestId(xRequestId).PostAllApiKeysRequest(postAllApiKeysRequest).Execute()

Creates an API key

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
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postAllApiKeysRequest := *openapiclient.NewPostAllApiKeysRequest("Name_example", "Org_example", "Role_example") // PostAllApiKeysRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApiKeysAPI.PostAllApiKeys(context.Background()).XRequestId(xRequestId).PostAllApiKeysRequest(postAllApiKeysRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysAPI.PostAllApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAllApiKeys`: FormattedApiApiKey
	fmt.Fprintf(os.Stdout, "Response from `ApiKeysAPI.PostAllApiKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAllApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postAllApiKeysRequest** | [**PostAllApiKeysRequest**](PostAllApiKeysRequest.md) |  | 

### Return type

[**FormattedApiApiKey**](FormattedApiApiKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

