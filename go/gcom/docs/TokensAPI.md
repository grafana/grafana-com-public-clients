# \TokensAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteToken**](TokensAPI.md#DeleteToken) | **Delete** /v1/tokens/{id} | Delete a token
[**GetToken**](TokensAPI.md#GetToken) | **Get** /v1/tokens/{id} | Get info for a specific token
[**GetTokens**](TokensAPI.md#GetTokens) | **Get** /v1/tokens | Get a list of tokens
[**PostToken**](TokensAPI.md#PostToken) | **Post** /v1/tokens/{id} | Update a token
[**PostTokens**](TokensAPI.md#PostTokens) | **Post** /v1/tokens | Create a new token



## DeleteToken

> map[string]interface{} DeleteToken(ctx, id).Region(region).XRequestId(xRequestId).OrgId(orgId).Execute()

Delete a token

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
	id := "id_example" // string | 
	region := "region_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	orgId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.DeleteToken(context.Background(), id).Region(region).XRequestId(xRequestId).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.DeleteToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteToken`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.DeleteToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** |  | 
 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **orgId** | **int32** |  | 

### Return type

**map[string]interface{}**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToken

> AuthToken GetToken(ctx, id).Region(region).OrgId(orgId).Execute()

Get info for a specific token

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
	id := "id_example" // string | 
	region := "region_example" // string | 
	orgId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.GetToken(context.Background(), id).Region(region).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.GetToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetToken`: AuthToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.GetToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** |  | 
 **orgId** | **int32** |  | 

### Return type

[**AuthToken**](AuthToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokens

> AccessPolicyListResponse GetTokens(ctx).Region(region).AccessPolicyId(accessPolicyId).AccessPolicyIds(accessPolicyIds).AccessPolicyName(accessPolicyName).AccessPolicyRealmIdentifier(accessPolicyRealmIdentifier).AccessPolicyRealmType(accessPolicyRealmType).AccessPolicyStatus(accessPolicyStatus).ExpiresAfter(expiresAfter).ExpiresBefore(expiresBefore).IncludeExpired(includeExpired).Name(name).OrgId(orgId).PageCursor(pageCursor).PageSize(pageSize).Execute()

Get a list of tokens

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
	region := "region_example" // string | 
	accessPolicyId := "accessPolicyId_example" // string |  (optional)
	accessPolicyIds := "accessPolicyIds_example" // string |  (optional)
	accessPolicyName := "accessPolicyName_example" // string |  (optional)
	accessPolicyRealmIdentifier := "accessPolicyRealmIdentifier_example" // string |  (optional)
	accessPolicyRealmType := "accessPolicyRealmType_example" // string |  (optional)
	accessPolicyStatus := "accessPolicyStatus_example" // string |  (optional)
	expiresAfter := "expiresAfter_example" // string |  (optional)
	expiresBefore := "expiresBefore_example" // string |  (optional)
	includeExpired := true // bool |  (optional)
	name := "name_example" // string |  (optional)
	orgId := int32(56) // int32 |  (optional)
	pageCursor := "pageCursor_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.GetTokens(context.Background()).Region(region).AccessPolicyId(accessPolicyId).AccessPolicyIds(accessPolicyIds).AccessPolicyName(accessPolicyName).AccessPolicyRealmIdentifier(accessPolicyRealmIdentifier).AccessPolicyRealmType(accessPolicyRealmType).AccessPolicyStatus(accessPolicyStatus).ExpiresAfter(expiresAfter).ExpiresBefore(expiresBefore).IncludeExpired(includeExpired).Name(name).OrgId(orgId).PageCursor(pageCursor).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.GetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokens`: AccessPolicyListResponse
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.GetTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** |  | 
 **accessPolicyId** | **string** |  | 
 **accessPolicyIds** | **string** |  | 
 **accessPolicyName** | **string** |  | 
 **accessPolicyRealmIdentifier** | **string** |  | 
 **accessPolicyRealmType** | **string** |  | 
 **accessPolicyStatus** | **string** |  | 
 **expiresAfter** | **string** |  | 
 **expiresBefore** | **string** |  | 
 **includeExpired** | **bool** |  | 
 **name** | **string** |  | 
 **orgId** | **int32** |  | 
 **pageCursor** | **string** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**AccessPolicyListResponse**](AccessPolicyListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostToken

> AuthToken PostToken(ctx, id).Region(region).XRequestId(xRequestId).PostTokenRequest(postTokenRequest).OrgId(orgId).Execute()

Update a token

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
	id := "id_example" // string | 
	region := "region_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postTokenRequest := *openapiclient.NewPostTokenRequest() // PostTokenRequest | 
	orgId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.PostToken(context.Background(), id).Region(region).XRequestId(xRequestId).PostTokenRequest(postTokenRequest).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.PostToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostToken`: AuthToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.PostToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** |  | 
 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postTokenRequest** | [**PostTokenRequest**](PostTokenRequest.md) |  | 
 **orgId** | **int32** |  | 

### Return type

[**AuthToken**](AuthToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTokens

> AuthTokenWithSecret PostTokens(ctx).Region(region).XRequestId(xRequestId).PostTokensRequest(postTokensRequest).OrgId(orgId).Execute()

Create a new token

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
	region := "region_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postTokensRequest := *openapiclient.NewPostTokensRequest("AccessPolicyId_example", "Name_example") // PostTokensRequest | 
	orgId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.PostTokens(context.Background()).Region(region).XRequestId(xRequestId).PostTokensRequest(postTokensRequest).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.PostTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostTokens`: AuthTokenWithSecret
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.PostTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** |  | 
 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postTokensRequest** | [**PostTokensRequest**](PostTokensRequest.md) |  | 
 **orgId** | **int32** |  | 

### Return type

[**AuthTokenWithSecret**](AuthTokenWithSecret.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

