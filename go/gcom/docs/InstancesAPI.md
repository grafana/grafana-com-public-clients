# \InstancesAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteInstance**](InstancesAPI.md#DeleteInstance) | **Delete** /instances/{instanceId} | Deletes an instance
[**GetInstance**](InstancesAPI.md#GetInstance) | **Get** /instances/{instanceId} | Gets an instance
[**GetInstances**](InstancesAPI.md#GetInstances) | **Get** /instances | Get a list of instances
[**PostInstance**](InstancesAPI.md#PostInstance) | **Post** /instances/{instanceId} | Updates an instance
[**PostInstances**](InstancesAPI.md#PostInstances) | **Post** /instances | Create a new instance



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


## GetInstance

> FormattedApiInstance GetInstance(ctx, instanceId).Config(config).IncludeDeleted(includeDeleted).Execute()

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
	includeDeleted := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.GetInstance(context.Background(), instanceId).Config(config).IncludeDeleted(includeDeleted).Execute()
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
 **includeDeleted** | **bool** |  | 

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


## GetInstances

> GetInstances200Response GetInstances(ctx).AmInstanceId(amInstanceId).Cluster(cluster).ClusterIdIn(clusterIdIn).Cursor(cursor).Direction(direction).HlInstanceId(hlInstanceId).HmInstanceGraphiteId(hmInstanceGraphiteId).HmInstancePromId(hmInstancePromId).Hosted(hosted).Id(id).IdIn(idIn).IdMin(idMin).Incident(incident).IncludeDeleted(includeDeleted).IncludePromCurrentActiveSeries(includePromCurrentActiveSeries).IncludeVersionIssueLink(includeVersionIssueLink).MachineLearning(machineLearning).MachineLearningLogsToken(machineLearningLogsToken).Name(name).NameIn(nameIn).OrderBy(orderBy).OrgAccountManagerId(orgAccountManagerId).OrgAccountOwnerId(orgAccountOwnerId).OrgId(orgId).OrgIdIn(orgIdIn).OrgSlug(orgSlug).OrgSlugIn(orgSlugIn).OrgType(orgType).Page(page).PageSize(pageSize).Plan(plan).PlanIn(planIn).PlanNot(planNot).Search(search).Slug(slug).SlugIn(slugIn).Status(status).Trial(trial).TrialExpiresAtMax(trialExpiresAtMax).TrialExpiresAtMin(trialExpiresAtMin).UpdatedOrCreatedAtMin(updatedOrCreatedAtMin).Url(url).UsageStatsId(usageStatsId).Version(version).VersionIn(versionIn).VersionNot(versionNot).VersionNotIn(versionNotIn).Execute()

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
	amInstanceId := int32(56) // int32 |  (optional)
	cluster := "cluster_example" // string |  (optional)
	clusterIdIn := []int32{int32(123)} // []int32 |  (optional)
	cursor := int32(56) // int32 |  (optional)
	direction := "direction_example" // string |  (optional)
	hlInstanceId := int32(56) // int32 |  (optional)
	hmInstanceGraphiteId := int32(56) // int32 |  (optional)
	hmInstancePromId := int32(56) // int32 |  (optional)
	hosted := true // bool |  (optional)
	id := "id_example" // string |  (optional)
	idIn := "idIn_example" // string |  (optional)
	idMin := int32(56) // int32 |  (optional)
	incident := true // bool |  (optional)
	includeDeleted := true // bool |  (optional)
	includePromCurrentActiveSeries := true // bool |  (optional)
	includeVersionIssueLink := true // bool |  (optional)
	machineLearning := true // bool |  (optional)
	machineLearningLogsToken := "machineLearningLogsToken_example" // string |  (optional)
	name := "name_example" // string |  (optional)
	nameIn := "nameIn_example" // string |  (optional)
	orderBy := "orderBy_example" // string |  (optional)
	orgAccountManagerId := int32(56) // int32 |  (optional)
	orgAccountOwnerId := int32(56) // int32 |  (optional)
	orgId := "orgId_example" // string |  (optional)
	orgIdIn := "orgIdIn_example" // string |  (optional)
	orgSlug := "orgSlug_example" // string |  (optional)
	orgSlugIn := "orgSlugIn_example" // string |  (optional)
	orgType := "orgType_example" // string |  (optional)
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)
	plan := "plan_example" // string |  (optional)
	planIn := "planIn_example" // string |  (optional)
	planNot := "planNot_example" // string |  (optional)
	search := "search_example" // string |  (optional)
	slug := "slug_example" // string |  (optional)
	slugIn := "slugIn_example" // string |  (optional)
	status := "status_example" // string |  (optional)
	trial := true // bool |  (optional)
	trialExpiresAtMax := time.Now() // time.Time |  (optional)
	trialExpiresAtMin := time.Now() // time.Time |  (optional)
	updatedOrCreatedAtMin := time.Now() // time.Time |  (optional)
	url := "url_example" // string |  (optional)
	usageStatsId := "usageStatsId_example" // string |  (optional)
	version := "version_example" // string |  (optional)
	versionIn := []string{"Inner_example"} // []string |  (optional)
	versionNot := "versionNot_example" // string |  (optional)
	versionNotIn := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancesAPI.GetInstances(context.Background()).AmInstanceId(amInstanceId).Cluster(cluster).ClusterIdIn(clusterIdIn).Cursor(cursor).Direction(direction).HlInstanceId(hlInstanceId).HmInstanceGraphiteId(hmInstanceGraphiteId).HmInstancePromId(hmInstancePromId).Hosted(hosted).Id(id).IdIn(idIn).IdMin(idMin).Incident(incident).IncludeDeleted(includeDeleted).IncludePromCurrentActiveSeries(includePromCurrentActiveSeries).IncludeVersionIssueLink(includeVersionIssueLink).MachineLearning(machineLearning).MachineLearningLogsToken(machineLearningLogsToken).Name(name).NameIn(nameIn).OrderBy(orderBy).OrgAccountManagerId(orgAccountManagerId).OrgAccountOwnerId(orgAccountOwnerId).OrgId(orgId).OrgIdIn(orgIdIn).OrgSlug(orgSlug).OrgSlugIn(orgSlugIn).OrgType(orgType).Page(page).PageSize(pageSize).Plan(plan).PlanIn(planIn).PlanNot(planNot).Search(search).Slug(slug).SlugIn(slugIn).Status(status).Trial(trial).TrialExpiresAtMax(trialExpiresAtMax).TrialExpiresAtMin(trialExpiresAtMin).UpdatedOrCreatedAtMin(updatedOrCreatedAtMin).Url(url).UsageStatsId(usageStatsId).Version(version).VersionIn(versionIn).VersionNot(versionNot).VersionNotIn(versionNotIn).Execute()
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
 **amInstanceId** | **int32** |  | 
 **cluster** | **string** |  | 
 **clusterIdIn** | **[]int32** |  | 
 **cursor** | **int32** |  | 
 **direction** | **string** |  | 
 **hlInstanceId** | **int32** |  | 
 **hmInstanceGraphiteId** | **int32** |  | 
 **hmInstancePromId** | **int32** |  | 
 **hosted** | **bool** |  | 
 **id** | **string** |  | 
 **idIn** | **string** |  | 
 **idMin** | **int32** |  | 
 **incident** | **bool** |  | 
 **includeDeleted** | **bool** |  | 
 **includePromCurrentActiveSeries** | **bool** |  | 
 **includeVersionIssueLink** | **bool** |  | 
 **machineLearning** | **bool** |  | 
 **machineLearningLogsToken** | **string** |  | 
 **name** | **string** |  | 
 **nameIn** | **string** |  | 
 **orderBy** | **string** |  | 
 **orgAccountManagerId** | **int32** |  | 
 **orgAccountOwnerId** | **int32** |  | 
 **orgId** | **string** |  | 
 **orgIdIn** | **string** |  | 
 **orgSlug** | **string** |  | 
 **orgSlugIn** | **string** |  | 
 **orgType** | **string** |  | 
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 
 **plan** | **string** |  | 
 **planIn** | **string** |  | 
 **planNot** | **string** |  | 
 **search** | **string** |  | 
 **slug** | **string** |  | 
 **slugIn** | **string** |  | 
 **status** | **string** |  | 
 **trial** | **bool** |  | 
 **trialExpiresAtMax** | **time.Time** |  | 
 **trialExpiresAtMin** | **time.Time** |  | 
 **updatedOrCreatedAtMin** | **time.Time** |  | 
 **url** | **string** |  | 
 **usageStatsId** | **string** |  | 
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

