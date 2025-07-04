# \OrgsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckOrgVisibility**](OrgsAPI.md#CheckOrgVisibility) | **Get** /orgs/{slugOrId}/check-visibility | 
[**DelApiKey**](OrgsAPI.md#DelApiKey) | **Delete** /orgs/{slugOrId}/api-keys/{name} | Delete an API key by name
[**DeleteOrgMember**](OrgsAPI.md#DeleteOrgMember) | **Delete** /orgs/{slugOrId}/members/{usernameOrId} | 
[**GetApiKey**](OrgsAPI.md#GetApiKey) | **Get** /orgs/{slugOrId}/api-keys/{name} | Get an API key by name
[**GetApiKeys**](OrgsAPI.md#GetApiKeys) | **Get** /orgs/{slugOrId}/api-keys | Get an organization&#39;s API keys
[**GetOrg**](OrgsAPI.md#GetOrg) | **Get** /orgs/{slugOrId} | 
[**GetOrgBilledUsage**](OrgsAPI.md#GetOrgBilledUsage) | **Get** /orgs/{orgSlugOrId}/billed-usage | Retrieve an org&#39;s billed usage for a specific month
[**GetOrgBilledUsageHistory**](OrgsAPI.md#GetOrgBilledUsageHistory) | **Get** /orgs/{orgSlugOrId}/billed-usage-history | Retrieve an org&#39;s billed usage items for the last 12 months
[**GetOrgInstances**](OrgsAPI.md#GetOrgInstances) | **Get** /orgs/{orgSlug}/instances | Get the list of instances belonging to the org
[**GetOrgMember**](OrgsAPI.md#GetOrgMember) | **Get** /orgs/{slugOrId}/members/{usernameOrId} | 
[**GetOrgMembers**](OrgsAPI.md#GetOrgMembers) | **Get** /orgs/{slugOrId}/members | 
[**PostApiKeys**](OrgsAPI.md#PostApiKeys) | **Post** /orgs/{slugOrId}/api-keys | Create an API key.
[**PostOrgMember**](OrgsAPI.md#PostOrgMember) | **Post** /orgs/{slugOrId}/members/{usernameOrId} | 
[**PostOrgMembers**](OrgsAPI.md#PostOrgMembers) | **Post** /orgs/{slugOrId}/members | 



## CheckOrgVisibility

> CheckOrgVisibility(ctx, slugOrId).Execute()





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
	slugOrId := "slugOrId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAPI.CheckOrgVisibility(context.Background(), slugOrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.CheckOrgVisibility``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slugOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckOrgVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DelApiKey

> DelApiKey(ctx, name, slugOrId).XRequestId(xRequestId).Unlock(unlock).Execute()

Delete an API key by name

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
	name := "name_example" // string | 
	slugOrId := "slugOrId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	unlock := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAPI.DelApiKey(context.Background(), name, slugOrId).XRequestId(xRequestId).Unlock(unlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.DelApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**slugOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **unlock** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgMember

> DeleteOrgMember(ctx, slugOrId, usernameOrId).XRequestId(xRequestId).Execute()



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
	slugOrId := "slugOrId_example" // string | 
	usernameOrId := "usernameOrId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgsAPI.DeleteOrgMember(context.Background(), slugOrId, usernameOrId).XRequestId(xRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.DeleteOrgMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slugOrId** | **string** |  | 
**usernameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKey

> FormattedApiApiKey GetApiKey(ctx, name, slugOrId).Execute()

Get an API key by name

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
	name := "name_example" // string | 
	slugOrId := "slugOrId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.GetApiKey(context.Background(), name, slugOrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.GetApiKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKey`: FormattedApiApiKey
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.GetApiKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**slugOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FormattedApiApiKey**](FormattedApiApiKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeys

> FormattedApiApiKeyListResponse GetApiKeys(ctx, slugOrId).ExcludeProvisioned(excludeProvisioned).Execute()

Get an organization's API keys

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
	slugOrId := "slugOrId_example" // string | 
	excludeProvisioned := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.GetApiKeys(context.Background(), slugOrId).ExcludeProvisioned(excludeProvisioned).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.GetApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApiKeys`: FormattedApiApiKeyListResponse
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.GetApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slugOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **excludeProvisioned** | **bool** |  | 

### Return type

[**FormattedApiApiKeyListResponse**](FormattedApiApiKeyListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrg

> FormattedApiOrgPublic GetOrg(ctx, slugOrId).Execute()



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
	slugOrId := "slugOrId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.GetOrg(context.Background(), slugOrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.GetOrg``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrg`: FormattedApiOrgPublic
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.GetOrg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slugOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FormattedApiOrgPublic**](FormattedApiOrgPublic.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgBilledUsage

> GetOrgBilledUsage200Response GetOrgBilledUsage(ctx, orgSlugOrId).Month(month).Year(year).Format(format).Execute()

Retrieve an org's billed usage for a specific month

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
	orgSlugOrId := "orgSlugOrId_example" // string | 
	month := int32(56) // int32 | 
	year := int32(56) // int32 | 
	format := "format_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.GetOrgBilledUsage(context.Background(), orgSlugOrId).Month(month).Year(year).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.GetOrgBilledUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgBilledUsage`: GetOrgBilledUsage200Response
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.GetOrgBilledUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlugOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgBilledUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **month** | **int32** |  | 
 **year** | **int32** |  | 
 **format** | **string** |  | 

### Return type

[**GetOrgBilledUsage200Response**](GetOrgBilledUsage200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgBilledUsageHistory

> OrgBilledUsageHistory GetOrgBilledUsageHistory(ctx, orgSlugOrId).Execute()

Retrieve an org's billed usage items for the last 12 months

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
	orgSlugOrId := "orgSlugOrId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.GetOrgBilledUsageHistory(context.Background(), orgSlugOrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.GetOrgBilledUsageHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgBilledUsageHistory`: OrgBilledUsageHistory
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.GetOrgBilledUsageHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlugOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgBilledUsageHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OrgBilledUsageHistory**](OrgBilledUsageHistory.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgInstances

> GetInstances200Response GetOrgInstances(ctx, orgSlug).Cluster(cluster).Direction(direction).Id(id).IdIn(idIn).Name(name).NameIn(nameIn).OrderBy(orderBy).Plan(plan).PlanIn(planIn).PlanNot(planNot).Slug(slug).SlugIn(slugIn).Url(url).UrlIn(urlIn).Execute()

Get the list of instances belonging to the org

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
	orgSlug := "orgSlug_example" // string | 
	cluster := "cluster_example" // string |  (optional)
	direction := "direction_example" // string |  (optional)
	id := "id_example" // string |  (optional)
	idIn := "idIn_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	nameIn := "nameIn_example" // string |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	plan := "plan_example" // string |  (optional)
	planIn := "planIn_example" // string |  (optional)
	planNot := "planNot_example" // string |  (optional)
	slug := "slug_example" // string |  (optional)
	slugIn := "slugIn_example" // string |  (optional)
	url := "url_example" // string |  (optional)
	urlIn := "urlIn_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.GetOrgInstances(context.Background(), orgSlug).Cluster(cluster).Direction(direction).Id(id).IdIn(idIn).Name(name).NameIn(nameIn).OrderBy(orderBy).Plan(plan).PlanIn(planIn).PlanNot(planNot).Slug(slug).SlugIn(slugIn).Url(url).UrlIn(urlIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.GetOrgInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgInstances`: GetInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.GetOrgInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgSlug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cluster** | **string** |  | 
 **direction** | **string** |  | 
 **id** | **string** |  | 
 **idIn** | **string** |  | 
 **name** | **string** |  | 
 **nameIn** | **string** |  | 
 **orderBy** | **string** |  | 
 **plan** | **string** |  | 
 **planIn** | **string** |  | 
 **planNot** | **string** |  | 
 **slug** | **string** |  | 
 **slugIn** | **string** |  | 
 **url** | **string** |  | 
 **urlIn** | **string** |  | 

### Return type

[**GetInstances200Response**](GetInstances200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgMember

> FormattedOrgMembership GetOrgMember(ctx, slugOrId, usernameOrId).Execute()



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
	slugOrId := "slugOrId_example" // string | 
	usernameOrId := "usernameOrId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.GetOrgMember(context.Background(), slugOrId, usernameOrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.GetOrgMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgMember`: FormattedOrgMembership
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.GetOrgMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slugOrId** | **string** |  | 
**usernameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FormattedOrgMembership**](FormattedOrgMembership.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgMembers

> OrgMemberListResponse GetOrgMembers(ctx, slugOrId).Billing(billing).Direction(direction).OrderBy(orderBy).Privacy(privacy).PrivacyIn(privacyIn).Execute()



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
	slugOrId := "slugOrId_example" // string | 
	billing := "billing_example" // string |  (optional)
	direction := "direction_example" // string |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	privacy := "privacy_example" // string |  (optional)
	privacyIn := "privacyIn_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.GetOrgMembers(context.Background(), slugOrId).Billing(billing).Direction(direction).OrderBy(orderBy).Privacy(privacy).PrivacyIn(privacyIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.GetOrgMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgMembers`: OrgMemberListResponse
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.GetOrgMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slugOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billing** | **string** |  | 
 **direction** | **string** |  | 
 **orderBy** | **string** |  | 
 **privacy** | **string** |  | 
 **privacyIn** | **string** |  | 

### Return type

[**OrgMemberListResponse**](OrgMemberListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostApiKeys

> FormattedApiApiKey PostApiKeys(ctx, slugOrId).XRequestId(xRequestId).PostApiKeysRequest(postApiKeysRequest).Execute()

Create an API key.

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
	slugOrId := "slugOrId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postApiKeysRequest := *openapiclient.NewPostApiKeysRequest("Name_example", "Role_example") // PostApiKeysRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.PostApiKeys(context.Background(), slugOrId).XRequestId(xRequestId).PostApiKeysRequest(postApiKeysRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.PostApiKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApiKeys`: FormattedApiApiKey
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.PostApiKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slugOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApiKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postApiKeysRequest** | [**PostApiKeysRequest**](PostApiKeysRequest.md) |  | 

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


## PostOrgMember

> FormattedOrgMembership PostOrgMember(ctx, slugOrId, usernameOrId).XRequestId(xRequestId).PostOrgMemberRequest(postOrgMemberRequest).Execute()



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
	slugOrId := "slugOrId_example" // string | 
	usernameOrId := "usernameOrId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postOrgMemberRequest := *openapiclient.NewPostOrgMemberRequest() // PostOrgMemberRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.PostOrgMember(context.Background(), slugOrId, usernameOrId).XRequestId(xRequestId).PostOrgMemberRequest(postOrgMemberRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.PostOrgMember``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgMember`: FormattedOrgMembership
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.PostOrgMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slugOrId** | **string** |  | 
**usernameOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postOrgMemberRequest** | [**PostOrgMemberRequest**](PostOrgMemberRequest.md) |  | 

### Return type

[**FormattedOrgMembership**](FormattedOrgMembership.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostOrgMembers

> FormattedOrgMembership PostOrgMembers(ctx, slugOrId).XRequestId(xRequestId).PostOrgMembersRequest(postOrgMembersRequest).Execute()



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
	slugOrId := "slugOrId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postOrgMembersRequest := *openapiclient.NewPostOrgMembersRequest("Username_example") // PostOrgMembersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgsAPI.PostOrgMembers(context.Background(), slugOrId).XRequestId(xRequestId).PostOrgMembersRequest(postOrgMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgsAPI.PostOrgMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrgMembers`: FormattedOrgMembership
	fmt.Fprintf(os.Stdout, "Response from `OrgsAPI.PostOrgMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slugOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrgMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postOrgMembersRequest** | [**PostOrgMembersRequest**](PostOrgMembersRequest.md) |  | 

### Return type

[**FormattedOrgMembership**](FormattedOrgMembership.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

