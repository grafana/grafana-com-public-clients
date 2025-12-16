# \StacksAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckStackReadinessV1**](StacksAPI.md#CheckStackReadinessV1) | **Get** /v1/stacks/{idOrSlug}/readiness | 
[**CheckStackSlugV1**](StacksAPI.md#CheckStackSlugV1) | **Post** /v1/stacks/checkUrl | 
[**CreateStackV1**](StacksAPI.md#CreateStackV1) | **Post** /v1/stacks | 
[**DeleteStackV1**](StacksAPI.md#DeleteStackV1) | **Delete** /v1/stacks/{idOrSlug} | 
[**GetStackV1**](StacksAPI.md#GetStackV1) | **Get** /v1/stacks/{idOrSlug} | 
[**ListStacksV1**](StacksAPI.md#ListStacksV1) | **Get** /v1/stacks | 
[**UpdateStackV1**](StacksAPI.md#UpdateStackV1) | **Post** /v1/stacks/{idOrSlug} | 



## CheckStackReadinessV1

> StackCheckReadinessV1 CheckStackReadinessV1(ctx, idOrSlug).Execute()



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
	idOrSlug := "idOrSlug_example" // string | id of the stack to check readiness

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.CheckStackReadinessV1(context.Background(), idOrSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.CheckStackReadinessV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckStackReadinessV1`: StackCheckReadinessV1
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.CheckStackReadinessV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrSlug** | **string** | id of the stack to check readiness | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckStackReadinessV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StackCheckReadinessV1**](StackCheckReadinessV1.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckStackSlugV1

> StackCheckV1 CheckStackSlugV1(ctx).StackCheckRequestV1(stackCheckRequestV1).Execute()



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
	stackCheckRequestV1 := *openapiclient.NewStackCheckRequestV1("Url_example") // StackCheckRequestV1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.CheckStackSlugV1(context.Background()).StackCheckRequestV1(stackCheckRequestV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.CheckStackSlugV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckStackSlugV1`: StackCheckV1
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.CheckStackSlugV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckStackSlugV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stackCheckRequestV1** | [**StackCheckRequestV1**](StackCheckRequestV1.md) |  | 

### Return type

[**StackCheckV1**](StackCheckV1.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateStackV1

> StackV1 CreateStackV1(ctx).StackCreateRequestV1(stackCreateRequestV1).Execute()



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
	stackCreateRequestV1 := *openapiclient.NewStackCreateRequestV1("Name_example", "Org_example", "Region_example", "Slug_example") // StackCreateRequestV1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.CreateStackV1(context.Background()).StackCreateRequestV1(stackCreateRequestV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.CreateStackV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStackV1`: StackV1
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.CreateStackV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStackV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stackCreateRequestV1** | [**StackCreateRequestV1**](StackCreateRequestV1.md) |  | 

### Return type

[**StackV1**](StackV1.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStackV1

> StackV1 DeleteStackV1(ctx, idOrSlug).Execute()



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
	idOrSlug := "idOrSlug_example" // string | id of the stack to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.DeleteStackV1(context.Background(), idOrSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.DeleteStackV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStackV1`: StackV1
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.DeleteStackV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrSlug** | **string** | id of the stack to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StackV1**](StackV1.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackV1

> StackV1 GetStackV1(ctx, idOrSlug).Execute()



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
	idOrSlug := "idOrSlug_example" // string | id of the stack to get

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.GetStackV1(context.Background(), idOrSlug).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.GetStackV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackV1`: StackV1
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.GetStackV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrSlug** | **string** | id of the stack to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StackV1**](StackV1.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListStacksV1

> PaginatedResponseStackV1 ListStacksV1(ctx).Org(org).Page(page).PageSize(pageSize).Execute()



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
	org := "org_example" // string | slug of the org to restrict the search to
	page := int64(789) // int64 | page to retrieve (default to 1)
	pageSize := int64(789) // int64 | page size to retrieve (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.ListStacksV1(context.Background()).Org(org).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.ListStacksV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListStacksV1`: PaginatedResponseStackV1
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.ListStacksV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListStacksV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **org** | **string** | slug of the org to restrict the search to | 
 **page** | **int64** | page to retrieve | [default to 1]
 **pageSize** | **int64** | page size to retrieve | [default to 10]

### Return type

[**PaginatedResponseStackV1**](PaginatedResponseStackV1.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStackV1

> StackV1 UpdateStackV1(ctx, idOrSlug).StackUpdateRequestV1(stackUpdateRequestV1).Execute()



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
	idOrSlug := "idOrSlug_example" // string | id of the stack to update
	stackUpdateRequestV1 := *openapiclient.NewStackUpdateRequestV1() // StackUpdateRequestV1 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StacksAPI.UpdateStackV1(context.Background(), idOrSlug).StackUpdateRequestV1(stackUpdateRequestV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StacksAPI.UpdateStackV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStackV1`: StackV1
	fmt.Fprintf(os.Stdout, "Response from `StacksAPI.UpdateStackV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrSlug** | **string** | id of the stack to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStackV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stackUpdateRequestV1** | [**StackUpdateRequestV1**](StackUpdateRequestV1.md) |  | 

### Return type

[**StackV1**](StackV1.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

