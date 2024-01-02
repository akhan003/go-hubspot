# \PublicObjectSchemasAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3SchemasObjectTypePurge**](PublicObjectSchemasAPI.md#DeleteCrmV3SchemasObjectTypePurge) | **Delete** /crm/v3/schemas/{objectType}/purge | 



## DeleteCrmV3SchemasObjectTypePurge

> DeleteCrmV3SchemasObjectTypePurge(ctx, objectType).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objectType := "objectType_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicObjectSchemasAPI.DeleteCrmV3SchemasObjectTypePurge(context.Background(), objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicObjectSchemasAPI.DeleteCrmV3SchemasObjectTypePurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3SchemasObjectTypePurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

