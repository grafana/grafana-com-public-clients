# \OrgStackRegionsAPI

All URIs are relative to */api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgStackRegions**](OrgStackRegionsAPI.md#GetOrgStackRegions) | **Get** /org-stack-regions | 



## GetOrgStackRegions

> GetOrgStackRegions200Response GetOrgStackRegions(ctx).OrgId(orgId).RegionSlugOrId(regionSlugOrId).Execute()



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
	orgId := int32(56) // int32 |  (optional)
	regionSlugOrId := "regionSlugOrId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgStackRegionsAPI.GetOrgStackRegions(context.Background()).OrgId(orgId).RegionSlugOrId(regionSlugOrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgStackRegionsAPI.GetOrgStackRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgStackRegions`: GetOrgStackRegions200Response
	fmt.Fprintf(os.Stdout, "Response from `OrgStackRegionsAPI.GetOrgStackRegions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgStackRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgId** | **int32** |  | 
 **regionSlugOrId** | **string** |  | 

### Return type

[**GetOrgStackRegions200Response**](GetOrgStackRegions200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

