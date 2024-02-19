# \InstancesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DelInstanceOAuthAzureAD**](InstancesAPI.md#DelInstanceOAuthAzureAD) | **Delete** /instances/{instanceId}/oauth-azuread | Disable Azure OAuth in an instance
[**DelInstanceOAuthGithub**](InstancesAPI.md#DelInstanceOAuthGithub) | **Delete** /instances/{instanceId}/oauth-github | Disable Github OAuth in an instance
[**DelInstanceOAuthGoogle**](InstancesAPI.md#DelInstanceOAuthGoogle) | **Delete** /instances/{instanceId}/oauth-google | Disable Google OAuth in an instance
[**DelInstanceOAuthOkta**](InstancesAPI.md#DelInstanceOAuthOkta) | **Delete** /instances/{instanceId}/oauth-okta | Disable Okta OAuth in an instance
[**DeleteInstance**](InstancesAPI.md#DeleteInstance) | **Delete** /instances/{instanceId} | Deletes an instance
[**DeleteInstancePlugin**](InstancesAPI.md#DeleteInstancePlugin) | **Delete** /instances/{instanceId}/plugins/{pluginSlugOrId} | 
[**GetConnections**](InstancesAPI.md#GetConnections) | **Get** /instances/{instanceId}/connections | Gets an instance&#39;s connectivity information (InfluxDB, OTEL, AWS private link, etc.)
[**GetInstance**](InstancesAPI.md#GetInstance) | **Get** /instances/{instanceId} | Gets an instance
[**GetInstancePlugin**](InstancesAPI.md#GetInstancePlugin) | **Get** /instances/{instanceId}/plugins/{pluginSlugOrId} | 
[**GetInstancePlugins**](InstancesAPI.md#GetInstancePlugins) | **Get** /instances/{instanceId}/plugins | 
[**GetInstances**](InstancesAPI.md#GetInstances) | **Get** /instances | Get a list of instances
[**PostInstance**](InstancesAPI.md#PostInstance) | **Post** /instances/{instanceId} | Updates an instance
[**PostInstanceOAuthAzureAD**](InstancesAPI.md#PostInstanceOAuthAzureAD) | **Post** /instances/{instanceId}/oauth-azuread | Configure Azure OAuth in an instance
[**PostInstanceOAuthGithub**](InstancesAPI.md#PostInstanceOAuthGithub) | **Post** /instances/{instanceId}/oauth-github | Configure Github OAuth in an instance
[**PostInstanceOAuthGoogle**](InstancesAPI.md#PostInstanceOAuthGoogle) | **Post** /instances/{instanceId}/oauth-google | Configure Google OAuth in an instance
[**PostInstanceOAuthOkta**](InstancesAPI.md#PostInstanceOAuthOkta) | **Post** /instances/{instanceId}/oauth-okta | Configure Okta OAuth in an instance
[**PostInstancePlugin**](InstancesAPI.md#PostInstancePlugin) | **Post** /instances/{instanceId}/plugins/{pluginSlugOrId} | 
[**PostInstancePlugins**](InstancesAPI.md#PostInstancePlugins) | **Post** /instances/{instanceId}/plugins | 
[**PostInstanceServiceAccountTokens**](InstancesAPI.md#PostInstanceServiceAccountTokens) | **Post** /instances/{instanceId}/api/serviceaccounts/{serviceAccountId}/tokens | Creates a service account token on a Grafana instance
[**PostInstanceServiceAccounts**](InstancesAPI.md#PostInstanceServiceAccounts) | **Post** /instances/{instanceId}/api/serviceaccounts | Creates a service account on a Grafana instance
[**PostInstances**](InstancesAPI.md#PostInstances) | **Post** /instances | Create a new instance



## DelInstanceOAuthAzureAD

> FormattedApiInstance DelInstanceOAuthAzureAD(ctx, instanceId).XRequestId(xRequestId).Execute()

Disable Azure OAuth in an instance

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
	instanceId := "instanceId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.DelInstanceOAuthAzureAD(context.Background(), instanceId).XRequestId(xRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.DelInstanceOAuthAzureAD``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DelInstanceOAuthAzureAD`: FormattedApiInstance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.DelInstanceOAuthAzureAD`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelInstanceOAuthAzureADRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]

### Return type

[**FormattedApiInstance**](FormattedApiInstance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DelInstanceOAuthGithub

> FormattedApiInstance DelInstanceOAuthGithub(ctx, instanceId).XRequestId(xRequestId).Execute()

Disable Github OAuth in an instance

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
	instanceId := "instanceId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.DelInstanceOAuthGithub(context.Background(), instanceId).XRequestId(xRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.DelInstanceOAuthGithub``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DelInstanceOAuthGithub`: FormattedApiInstance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.DelInstanceOAuthGithub`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelInstanceOAuthGithubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]

### Return type

[**FormattedApiInstance**](FormattedApiInstance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DelInstanceOAuthGoogle

> FormattedApiInstance DelInstanceOAuthGoogle(ctx, instanceId).XRequestId(xRequestId).Execute()

Disable Google OAuth in an instance

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
	instanceId := "instanceId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.DelInstanceOAuthGoogle(context.Background(), instanceId).XRequestId(xRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.DelInstanceOAuthGoogle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DelInstanceOAuthGoogle`: FormattedApiInstance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.DelInstanceOAuthGoogle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelInstanceOAuthGoogleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]

### Return type

[**FormattedApiInstance**](FormattedApiInstance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DelInstanceOAuthOkta

> FormattedApiInstance DelInstanceOAuthOkta(ctx, instanceId).XRequestId(xRequestId).Execute()

Disable Okta OAuth in an instance

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
	instanceId := "instanceId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.DelInstanceOAuthOkta(context.Background(), instanceId).XRequestId(xRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.DelInstanceOAuthOkta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DelInstanceOAuthOkta`: FormattedApiInstance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.DelInstanceOAuthOkta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelInstanceOAuthOktaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]

### Return type

[**FormattedApiInstance**](FormattedApiInstance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> FormattedApiInstance DeleteInstance(ctx, instanceId).XRequestId(xRequestId).Execute()

Deletes an instance

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
	instanceId := "instanceId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.DeleteInstance(context.Background(), instanceId).XRequestId(xRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.DeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInstance`: FormattedApiInstance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.DeleteInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]

### Return type

[**FormattedApiInstance**](FormattedApiInstance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstancePlugin

> FormattedApiInstancePlugin DeleteInstancePlugin(ctx, instanceId, pluginSlugOrId).XRequestId(xRequestId).Execute()



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
	instanceId := "instanceId_example" // string | 
	pluginSlugOrId := "pluginSlugOrId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.DeleteInstancePlugin(context.Background(), instanceId, pluginSlugOrId).XRequestId(xRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.DeleteInstancePlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteInstancePlugin`: FormattedApiInstancePlugin
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.DeleteInstancePlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**pluginSlugOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstancePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]

### Return type

[**FormattedApiInstancePlugin**](FormattedApiInstancePlugin.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnections

> FormattedApiInstanceConnections GetConnections(ctx, instanceId).Config(config).Execute()

Gets an instance's connectivity information (InfluxDB, OTEL, AWS private link, etc.)

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
	instanceId := "instanceId_example" // string | 
	config := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.GetConnections(context.Background(), instanceId).Config(config).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.GetConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnections`: FormattedApiInstanceConnections
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.GetConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **config** | **bool** |  | 

### Return type

[**FormattedApiInstanceConnections**](FormattedApiInstanceConnections.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> FormattedApiInstance GetInstance(ctx, instanceId).Config(config).Execute()

Gets an instance

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
	instanceId := "instanceId_example" // string | 
	config := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.GetInstance(context.Background(), instanceId).Config(config).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance`: FormattedApiInstance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **config** | **bool** |  | 

### Return type

[**FormattedApiInstance**](FormattedApiInstance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstancePlugin

> FormattedApiInstancePlugin GetInstancePlugin(ctx, instanceId, pluginSlugOrId).Execute()



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
	instanceId := "instanceId_example" // string | 
	pluginSlugOrId := "pluginSlugOrId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.GetInstancePlugin(context.Background(), instanceId, pluginSlugOrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.GetInstancePlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstancePlugin`: FormattedApiInstancePlugin
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.GetInstancePlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**pluginSlugOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstancePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FormattedApiInstancePlugin**](FormattedApiInstancePlugin.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstancePlugins

> GetInstancePlugins200Response GetInstancePlugins(ctx, instanceId).Direction(direction).OrderBy(orderBy).Execute()



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
	instanceId := "instanceId_example" // string | 
	direction := "direction_example" // string |  (optional)
	orderBy := "orderBy_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.GetInstancePlugins(context.Background(), instanceId).Direction(direction).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.GetInstancePlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstancePlugins`: GetInstancePlugins200Response
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.GetInstancePlugins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstancePluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **direction** | **string** |  | 
 **orderBy** | **string** |  | 

### Return type

[**GetInstancePlugins200Response**](GetInstancePlugins200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstances

> GetInstances200Response GetInstances(ctx).Cluster(cluster).ClusterIdIn(clusterIdIn).Cursor(cursor).Direction(direction).Hosted(hosted).Id(id).IdIn(idIn).IdMin(idMin).IncludeLabels(includeLabels).IncludePromCurrentActiveSeries(includePromCurrentActiveSeries).Labels(labels).Name(name).NameIn(nameIn).OrderBy(orderBy).OrgId(orgId).OrgIdIn(orgIdIn).OrgSlug(orgSlug).OrgSlugIn(orgSlugIn).Page(page).PageSize(pageSize).Plan(plan).PlanIn(planIn).PlanNot(planNot).Slug(slug).SlugIn(slugIn).Status(status).UpdatedOrCreatedAtMin(updatedOrCreatedAtMin).Url(url).Version(version).VersionIn(versionIn).VersionNot(versionNot).VersionNotIn(versionNotIn).Execute()

Get a list of instances

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/grafana/grafana-com-public-clients/go/gcom"
)

func main() {
	cluster := "cluster_example" // string |  (optional)
	clusterIdIn := []int32{int32(123)} // []int32 |  (optional)
	cursor := int32(56) // int32 |  (optional)
	direction := "direction_example" // string |  (optional)
	hosted := true // bool |  (optional)
	id := "id_example" // string |  (optional)
	idIn := "idIn_example" // string |  (optional)
	idMin := int32(56) // int32 |  (optional)
	includeLabels := true // bool |  (optional)
	includePromCurrentActiveSeries := true // bool |  (optional)
	labels := []string{"Inner_example"} // []string |  (optional)
	name := "name_example" // string |  (optional)
	nameIn := "nameIn_example" // string |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	orgId := "orgId_example" // string |  (optional)
	orgIdIn := "orgIdIn_example" // string |  (optional)
	orgSlug := "orgSlug_example" // string |  (optional)
	orgSlugIn := "orgSlugIn_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	plan := "plan_example" // string |  (optional)
	planIn := "planIn_example" // string |  (optional)
	planNot := "planNot_example" // string |  (optional)
	slug := "slug_example" // string |  (optional)
	slugIn := "slugIn_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	updatedOrCreatedAtMin := time.Now() // time.Time |  (optional)
	url := "url_example" // string |  (optional)
	version := "version_example" // string |  (optional)
	versionIn := []string{"Inner_example"} // []string |  (optional)
	versionNot := "versionNot_example" // string |  (optional)
	versionNotIn := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.GetInstances(context.Background()).Cluster(cluster).ClusterIdIn(clusterIdIn).Cursor(cursor).Direction(direction).Hosted(hosted).Id(id).IdIn(idIn).IdMin(idMin).IncludeLabels(includeLabels).IncludePromCurrentActiveSeries(includePromCurrentActiveSeries).Labels(labels).Name(name).NameIn(nameIn).OrderBy(orderBy).OrgId(orgId).OrgIdIn(orgIdIn).OrgSlug(orgSlug).OrgSlugIn(orgSlugIn).Page(page).PageSize(pageSize).Plan(plan).PlanIn(planIn).PlanNot(planNot).Slug(slug).SlugIn(slugIn).Status(status).UpdatedOrCreatedAtMin(updatedOrCreatedAtMin).Url(url).Version(version).VersionIn(versionIn).VersionNot(versionNot).VersionNotIn(versionNotIn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.GetInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstances`: GetInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.GetInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** |  | 
 **clusterIdIn** | **[]int32** |  | 
 **cursor** | **int32** |  | 
 **direction** | **string** |  | 
 **hosted** | **bool** |  | 
 **id** | **string** |  | 
 **idIn** | **string** |  | 
 **idMin** | **int32** |  | 
 **includeLabels** | **bool** |  | 
 **includePromCurrentActiveSeries** | **bool** |  | 
 **labels** | **[]string** |  | 
 **name** | **string** |  | 
 **nameIn** | **string** |  | 
 **orderBy** | **string** |  | 
 **orgId** | **string** |  | 
 **orgIdIn** | **string** |  | 
 **orgSlug** | **string** |  | 
 **orgSlugIn** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **plan** | **string** |  | 
 **planIn** | **string** |  | 
 **planNot** | **string** |  | 
 **slug** | **string** |  | 
 **slugIn** | **string** |  | 
 **status** | **string** |  | 
 **updatedOrCreatedAtMin** | **time.Time** |  | 
 **url** | **string** |  | 
 **version** | **string** |  | 
 **versionIn** | **[]string** |  | 
 **versionNot** | **string** |  | 
 **versionNotIn** | **[]string** |  | 

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


## PostInstance

> FormattedApiInstance PostInstance(ctx, instanceId).XRequestId(xRequestId).PostInstanceRequest(postInstanceRequest).Execute()

Updates an instance

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
	instanceId := "instanceId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postInstanceRequest := *openapiclient.NewPostInstanceRequest() // PostInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.PostInstance(context.Background(), instanceId).XRequestId(xRequestId).PostInstanceRequest(postInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.PostInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstance`: FormattedApiInstance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.PostInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postInstanceRequest** | [**PostInstanceRequest**](PostInstanceRequest.md) |  | 

### Return type

[**FormattedApiInstance**](FormattedApiInstance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstanceOAuthAzureAD

> FormattedApiInstance PostInstanceOAuthAzureAD(ctx, instanceId).XRequestId(xRequestId).PostInstanceOAuthAzureADRequest(postInstanceOAuthAzureADRequest).Execute()

Configure Azure OAuth in an instance

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
	instanceId := "instanceId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postInstanceOAuthAzureADRequest := *openapiclient.NewPostInstanceOAuthAzureADRequest("AuthUrl_example", "ClientId_example", "TokenUrl_example") // PostInstanceOAuthAzureADRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.PostInstanceOAuthAzureAD(context.Background(), instanceId).XRequestId(xRequestId).PostInstanceOAuthAzureADRequest(postInstanceOAuthAzureADRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.PostInstanceOAuthAzureAD``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstanceOAuthAzureAD`: FormattedApiInstance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.PostInstanceOAuthAzureAD`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostInstanceOAuthAzureADRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postInstanceOAuthAzureADRequest** | [**PostInstanceOAuthAzureADRequest**](PostInstanceOAuthAzureADRequest.md) |  | 

### Return type

[**FormattedApiInstance**](FormattedApiInstance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstanceOAuthGithub

> FormattedApiInstance PostInstanceOAuthGithub(ctx, instanceId).XRequestId(xRequestId).PostInstanceOAuthGithubRequest(postInstanceOAuthGithubRequest).Execute()

Configure Github OAuth in an instance

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
	instanceId := "instanceId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postInstanceOAuthGithubRequest := *openapiclient.NewPostInstanceOAuthGithubRequest("ClientId_example") // PostInstanceOAuthGithubRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.PostInstanceOAuthGithub(context.Background(), instanceId).XRequestId(xRequestId).PostInstanceOAuthGithubRequest(postInstanceOAuthGithubRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.PostInstanceOAuthGithub``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstanceOAuthGithub`: FormattedApiInstance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.PostInstanceOAuthGithub`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostInstanceOAuthGithubRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postInstanceOAuthGithubRequest** | [**PostInstanceOAuthGithubRequest**](PostInstanceOAuthGithubRequest.md) |  | 

### Return type

[**FormattedApiInstance**](FormattedApiInstance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstanceOAuthGoogle

> FormattedApiInstance PostInstanceOAuthGoogle(ctx, instanceId).XRequestId(xRequestId).PostInstanceOAuthGoogleRequest(postInstanceOAuthGoogleRequest).Execute()

Configure Google OAuth in an instance

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
	instanceId := "instanceId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postInstanceOAuthGoogleRequest := *openapiclient.NewPostInstanceOAuthGoogleRequest([]string{"AllowedDomains_example"}, "ClientId_example") // PostInstanceOAuthGoogleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.PostInstanceOAuthGoogle(context.Background(), instanceId).XRequestId(xRequestId).PostInstanceOAuthGoogleRequest(postInstanceOAuthGoogleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.PostInstanceOAuthGoogle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstanceOAuthGoogle`: FormattedApiInstance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.PostInstanceOAuthGoogle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostInstanceOAuthGoogleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postInstanceOAuthGoogleRequest** | [**PostInstanceOAuthGoogleRequest**](PostInstanceOAuthGoogleRequest.md) |  | 

### Return type

[**FormattedApiInstance**](FormattedApiInstance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstanceOAuthOkta

> FormattedApiInstance PostInstanceOAuthOkta(ctx, instanceId).XRequestId(xRequestId).PostInstanceOAuthOktaRequest(postInstanceOAuthOktaRequest).Execute()

Configure Okta OAuth in an instance

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
	instanceId := "instanceId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postInstanceOAuthOktaRequest := *openapiclient.NewPostInstanceOAuthOktaRequest("ApiUrl_example", "AuthUrl_example", "ClientId_example", "TokenUrl_example") // PostInstanceOAuthOktaRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.PostInstanceOAuthOkta(context.Background(), instanceId).XRequestId(xRequestId).PostInstanceOAuthOktaRequest(postInstanceOAuthOktaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.PostInstanceOAuthOkta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstanceOAuthOkta`: FormattedApiInstance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.PostInstanceOAuthOkta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostInstanceOAuthOktaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postInstanceOAuthOktaRequest** | [**PostInstanceOAuthOktaRequest**](PostInstanceOAuthOktaRequest.md) |  | 

### Return type

[**FormattedApiInstance**](FormattedApiInstance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstancePlugin

> FormattedApiInstancePlugin PostInstancePlugin(ctx, instanceId, pluginSlugOrId).XRequestId(xRequestId).PostInstancePluginRequest(postInstancePluginRequest).Execute()



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
	instanceId := "instanceId_example" // string | 
	pluginSlugOrId := "pluginSlugOrId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postInstancePluginRequest := *openapiclient.NewPostInstancePluginRequest() // PostInstancePluginRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.PostInstancePlugin(context.Background(), instanceId, pluginSlugOrId).XRequestId(xRequestId).PostInstancePluginRequest(postInstancePluginRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.PostInstancePlugin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstancePlugin`: FormattedApiInstancePlugin
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.PostInstancePlugin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**pluginSlugOrId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostInstancePluginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postInstancePluginRequest** | [**PostInstancePluginRequest**](PostInstancePluginRequest.md) |  | 

### Return type

[**FormattedApiInstancePlugin**](FormattedApiInstancePlugin.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstancePlugins

> FormattedApiInstancePlugin PostInstancePlugins(ctx, instanceId).XRequestId(xRequestId).PostInstancePluginsRequest(postInstancePluginsRequest).Execute()



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
	instanceId := "instanceId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postInstancePluginsRequest := *openapiclient.NewPostInstancePluginsRequest("Plugin_example") // PostInstancePluginsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.PostInstancePlugins(context.Background(), instanceId).XRequestId(xRequestId).PostInstancePluginsRequest(postInstancePluginsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.PostInstancePlugins``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstancePlugins`: FormattedApiInstancePlugin
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.PostInstancePlugins`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostInstancePluginsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postInstancePluginsRequest** | [**PostInstancePluginsRequest**](PostInstancePluginsRequest.md) |  | 

### Return type

[**FormattedApiInstancePlugin**](FormattedApiInstancePlugin.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstanceServiceAccountTokens

> GrafanaNewApiKeyResult PostInstanceServiceAccountTokens(ctx, instanceId, serviceAccountId).XRequestId(xRequestId).PostInstanceServiceAccountTokensRequest(postInstanceServiceAccountTokensRequest).Execute()

Creates a service account token on a Grafana instance

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
	instanceId := "instanceId_example" // string | 
	serviceAccountId := "serviceAccountId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postInstanceServiceAccountTokensRequest := *openapiclient.NewPostInstanceServiceAccountTokensRequest("Name_example") // PostInstanceServiceAccountTokensRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.PostInstanceServiceAccountTokens(context.Background(), instanceId, serviceAccountId).XRequestId(xRequestId).PostInstanceServiceAccountTokensRequest(postInstanceServiceAccountTokensRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.PostInstanceServiceAccountTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstanceServiceAccountTokens`: GrafanaNewApiKeyResult
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.PostInstanceServiceAccountTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 
**serviceAccountId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostInstanceServiceAccountTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postInstanceServiceAccountTokensRequest** | [**PostInstanceServiceAccountTokensRequest**](PostInstanceServiceAccountTokensRequest.md) |  | 

### Return type

[**GrafanaNewApiKeyResult**](GrafanaNewApiKeyResult.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstanceServiceAccounts

> GrafanaServiceAccountDTO PostInstanceServiceAccounts(ctx, instanceId).XRequestId(xRequestId).PostInstanceServiceAccountsRequest(postInstanceServiceAccountsRequest).Execute()

Creates a service account on a Grafana instance

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
	instanceId := "instanceId_example" // string | 
	xRequestId := "xRequestId_example" // string |  (default to "openapi-x-request-id")
	postInstanceServiceAccountsRequest := *openapiclient.NewPostInstanceServiceAccountsRequest("Name_example", "Role_example") // PostInstanceServiceAccountsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.PostInstanceServiceAccounts(context.Background(), instanceId).XRequestId(xRequestId).PostInstanceServiceAccountsRequest(postInstanceServiceAccountsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.PostInstanceServiceAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstanceServiceAccounts`: GrafanaServiceAccountDTO
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.PostInstanceServiceAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostInstanceServiceAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postInstanceServiceAccountsRequest** | [**PostInstanceServiceAccountsRequest**](PostInstanceServiceAccountsRequest.md) |  | 

### Return type

[**GrafanaServiceAccountDTO**](GrafanaServiceAccountDTO.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostInstances

> FormattedApiInstance PostInstances(ctx).XRequestId(xRequestId).PostInstancesRequest(postInstancesRequest).Execute()

Create a new instance

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
	postInstancesRequest := *openapiclient.NewPostInstancesRequest("Name_example") // PostInstancesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.PostInstances(context.Background()).XRequestId(xRequestId).PostInstancesRequest(postInstancesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancesAPI.PostInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostInstances`: FormattedApiInstance
	fmt.Fprintf(os.Stdout, "Response from `InstancesAPI.PostInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xRequestId** | **string** |  | [default to &quot;openapi-x-request-id&quot;]
 **postInstancesRequest** | [**PostInstancesRequest**](PostInstancesRequest.md) |  | 

### Return type

[**FormattedApiInstance**](FormattedApiInstance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

