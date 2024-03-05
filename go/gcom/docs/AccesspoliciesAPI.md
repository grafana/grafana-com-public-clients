# \AccesspoliciesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccessPolicy**](AccesspoliciesAPI.md#DeleteAccessPolicy) | **Delete** /v1/accesspolicies/{id} | Delete an access policy
[**GetAccessPolicies**](AccesspoliciesAPI.md#GetAccessPolicies) | **Get** /v1/accesspolicies | Get a list of access policies
[**GetAccessPolicy**](AccesspoliciesAPI.md#GetAccessPolicy) | **Get** /v1/accesspolicies/{id} | Get an access policy
[**GetConfig**](AccesspoliciesAPI.md#GetConfig) | **Get** /v1/accesspolicies/config | Get details about the Cloud Access Policy API
[**PostAccessPolicies**](AccesspoliciesAPI.md#PostAccessPolicies) | **Post** /v1/accesspolicies | Create a new access policy
[**PostAccessPolicy**](AccesspoliciesAPI.md#PostAccessPolicy) | **Post** /v1/accesspolicies/{id} | Update an access policy



## DeleteAccessPolicy

> map[string]interface{} DeleteAccessPolicy(ctx, id).Region(region).XRequestId(xRequestId).OrgId(orgId).Execute()

Delete an access policy

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
	resp, r, err := apiClient.AccesspoliciesAPI.DeleteAccessPolicy(context.Background(), id).Region(region).XRequestId(xRequestId).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccesspoliciesAPI.DeleteAccessPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccessPolicy`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccesspoliciesAPI.DeleteAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessPolicyRequest struct via the builder pattern


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


## GetAccessPolicies

> GetAccessPolicies200Response GetAccessPolicies(ctx).Region(region).Name(name).OrgId(orgId).PageCursor(pageCursor).PageSize(pageSize).RealmIdentifier(realmIdentifier).RealmType(realmType).Status(status).Execute()

Get a list of access policies

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
	name := "name_example" // string |  (optional)
	orgId := int32(56) // int32 |  (optional)
	pageCursor := "pageCursor_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	realmIdentifier := "realmIdentifier_example" // string |  (optional)
	realmType := "realmType_example" // string |  (optional)
	status := "status_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccesspoliciesAPI.GetAccessPolicies(context.Background()).Region(region).Name(name).OrgId(orgId).PageCursor(pageCursor).PageSize(pageSize).RealmIdentifier(realmIdentifier).RealmType(realmType).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccesspoliciesAPI.GetAccessPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessPolicies`: GetAccessPolicies200Response
	fmt.Fprintf(os.Stdout, "Response from `AccesspoliciesAPI.GetAccessPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** |  | 
 **name** | **string** |  | 
 **orgId** | **int32** |  | 
 **pageCursor** | **string** |  | 
 **pageSize** | **int32** |  | 
 **realmIdentifier** | **string** |  | 
 **realmType** | **string** |  | 
 **status** | **string** |  | 

### Return type

[**GetAccessPolicies200Response**](GetAccessPolicies200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessPolicy

> AuthAccessPolicy GetAccessPolicy(ctx, id).Region(region).OrgId(orgId).Execute()

Get an access policy

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
	resp, r, err := apiClient.AccesspoliciesAPI.GetAccessPolicy(context.Background(), id).Region(region).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccesspoliciesAPI.GetAccessPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessPolicy`: AuthAccessPolicy
	fmt.Fprintf(os.Stdout, "Response from `AccesspoliciesAPI.GetAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** |  | 
 **orgId** | **int32** |  | 

### Return type

[**AuthAccessPolicy**](AuthAccessPolicy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfig

> map[string]interface{} GetConfig(ctx).Region(region).Execute()

Get details about the Cloud Access Policy API

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccesspoliciesAPI.GetConfig(context.Background()).Region(region).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccesspoliciesAPI.GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccesspoliciesAPI.GetConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** |  | 

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


## PostAccessPolicies

> AuthAccessPolicy PostAccessPolicies(ctx).Region(region).XRequestId(xRequestId).PostAccessPoliciesRequest(postAccessPoliciesRequest).OrgId(orgId).Execute()

Create a new access policy

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
	postAccessPoliciesRequest := *openapiclient.NewPostAccessPoliciesRequest("Name_example", []openapiclient.PostAccessPoliciesRequestRealmsInner{*openapiclient.NewPostAccessPoliciesRequestRealmsInner("Identifier_example", "Type_example")}, []string{"Scopes_example"}) // PostAccessPoliciesRequest | 
	orgId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccesspoliciesAPI.PostAccessPolicies(context.Background()).Region(region).XRequestId(xRequestId).PostAccessPoliciesRequest(postAccessPoliciesRequest).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccesspoliciesAPI.PostAccessPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAccessPolicies`: AuthAccessPolicy
	fmt.Fprintf(os.Stdout, "Response from `AccesspoliciesAPI.PostAccessPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAccessPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **region** | **string** |  | 
 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postAccessPoliciesRequest** | [**PostAccessPoliciesRequest**](PostAccessPoliciesRequest.md) |  | 
 **orgId** | **int32** |  | 

### Return type

[**AuthAccessPolicy**](AuthAccessPolicy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAccessPolicy

> AuthAccessPolicy PostAccessPolicy(ctx, id).Region(region).XRequestId(xRequestId).PostAccessPolicyRequest(postAccessPolicyRequest).OrgId(orgId).Execute()

Update an access policy

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
	postAccessPolicyRequest := *openapiclient.NewPostAccessPolicyRequest() // PostAccessPolicyRequest | 
	orgId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccesspoliciesAPI.PostAccessPolicy(context.Background(), id).Region(region).XRequestId(xRequestId).PostAccessPolicyRequest(postAccessPolicyRequest).OrgId(orgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccesspoliciesAPI.PostAccessPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAccessPolicy`: AuthAccessPolicy
	fmt.Fprintf(os.Stdout, "Response from `AccesspoliciesAPI.PostAccessPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAccessPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **region** | **string** |  | 
 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postAccessPolicyRequest** | [**PostAccessPolicyRequest**](PostAccessPolicyRequest.md) |  | 
 **orgId** | **int32** |  | 

### Return type

[**AuthAccessPolicy**](AuthAccessPolicy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

