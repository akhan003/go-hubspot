# \CoreAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Cancel**](CoreAPI.md#Cancel) | **Post** /crm/v3/imports/{importId}/cancel | Cancel an active import
[**Create**](CoreAPI.md#Create) | **Post** /crm/v3/imports/ | Start a new import
[**GetByID**](CoreAPI.md#GetByID) | **Get** /crm/v3/imports/{importId} | Get the information on any import
[**GetPage**](CoreAPI.md#GetPage) | **Get** /crm/v3/imports/ | Get active imports



## Cancel

> ActionResponse Cancel(ctx, importId).Execute()

Cancel an active import



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
	importId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.Cancel(context.Background(), importId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.Cancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Cancel`: ActionResponse
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.Cancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActionResponse**](ActionResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> PublicImportResponse Create(ctx).Files(files).ImportRequest(importRequest).Execute()

Start a new import



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
	files := os.NewFile(1234, "some_file") // *os.File | A list of files containing the data to import (optional)
	importRequest := "importRequest_example" // string | JSON formatted metadata about the import. This includes a name for the import and the column mappings for each file. See the overview tab for more on the required format. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.Create(context.Background()).Files(files).ImportRequest(importRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: PublicImportResponse
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **files** | ***os.File** | A list of files containing the data to import | 
 **importRequest** | **string** | JSON formatted metadata about the import. This includes a name for the import and the column mappings for each file. See the overview tab for more on the required format. | 

### Return type

[**PublicImportResponse**](PublicImportResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> PublicImportResponse GetByID(ctx, importId).Execute()

Get the information on any import



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
	importId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.GetByID(context.Background(), importId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.GetByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByID`: PublicImportResponse
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.GetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**importId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicImportResponse**](PublicImportResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPage

> CollectionResponsePublicImportResponse GetPage(ctx).After(after).Before(before).Limit(limit).Execute()

Get active imports



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
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	before := "before_example" // string |  (optional)
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.GetPage(context.Background()).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.GetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPage`: CollectionResponsePublicImportResponse
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.GetPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | The maximum number of results to display per page. | 

### Return type

[**CollectionResponsePublicImportResponse**](CollectionResponsePublicImportResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

