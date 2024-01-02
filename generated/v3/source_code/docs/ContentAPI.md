# \ContentAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3SourceCodeEnvironmentContentPath**](ContentAPI.md#DeleteCmsV3SourceCodeEnvironmentContentPath) | **Delete** /cms/v3/source-code/{environment}/content/{path} | Delete a file
[**GetCmsV3SourceCodeEnvironmentContentPath**](ContentAPI.md#GetCmsV3SourceCodeEnvironmentContentPath) | **Get** /cms/v3/source-code/{environment}/content/{path} | Download a file
[**PostCmsV3SourceCodeEnvironmentContentPath**](ContentAPI.md#PostCmsV3SourceCodeEnvironmentContentPath) | **Post** /cms/v3/source-code/{environment}/content/{path} | Create a file
[**PutCmsV3SourceCodeEnvironmentContentPath**](ContentAPI.md#PutCmsV3SourceCodeEnvironmentContentPath) | **Put** /cms/v3/source-code/{environment}/content/{path} | Create or update a file



## DeleteCmsV3SourceCodeEnvironmentContentPath

> DeleteCmsV3SourceCodeEnvironmentContentPath(ctx, environment, path).Execute()

Delete a file



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
	environment := "environment_example" // string | The environment of the file (\"draft\" or \"published\").
	path := "path_example" // string | The file system location of the file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ContentAPI.DeleteCmsV3SourceCodeEnvironmentContentPath(context.Background(), environment, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.DeleteCmsV3SourceCodeEnvironmentContentPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3SourceCodeEnvironmentContentPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3SourceCodeEnvironmentContentPath

> Error GetCmsV3SourceCodeEnvironmentContentPath(ctx, environment, path).Execute()

Download a file



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
	environment := "environment_example" // string | The environment of the file (\"draft\" or \"published\").
	path := "path_example" // string | The file system location of the file.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.GetCmsV3SourceCodeEnvironmentContentPath(context.Background(), environment, path).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.GetCmsV3SourceCodeEnvironmentContentPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3SourceCodeEnvironmentContentPath`: Error
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.GetCmsV3SourceCodeEnvironmentContentPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3SourceCodeEnvironmentContentPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3SourceCodeEnvironmentContentPath

> AssetFileMetadata PostCmsV3SourceCodeEnvironmentContentPath(ctx, environment, path).File(file).Execute()

Create a file



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
	environment := "environment_example" // string | The environment of the file (\"draft\" or \"published\").
	path := "path_example" // string | The file system location of the file.
	file := os.NewFile(1234, "some_file") // *os.File | The file to upload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.PostCmsV3SourceCodeEnvironmentContentPath(context.Background(), environment, path).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.PostCmsV3SourceCodeEnvironmentContentPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3SourceCodeEnvironmentContentPath`: AssetFileMetadata
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.PostCmsV3SourceCodeEnvironmentContentPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3SourceCodeEnvironmentContentPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | The file to upload. | 

### Return type

[**AssetFileMetadata**](AssetFileMetadata.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCmsV3SourceCodeEnvironmentContentPath

> AssetFileMetadata PutCmsV3SourceCodeEnvironmentContentPath(ctx, environment, path).File(file).Execute()

Create or update a file



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
	environment := "environment_example" // string | The environment of the file (\"draft\" or \"published\").
	path := "path_example" // string | The file system location of the file.
	file := os.NewFile(1234, "some_file") // *os.File | The file to upload. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContentAPI.PutCmsV3SourceCodeEnvironmentContentPath(context.Background(), environment, path).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContentAPI.PutCmsV3SourceCodeEnvironmentContentPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCmsV3SourceCodeEnvironmentContentPath`: AssetFileMetadata
	fmt.Fprintf(os.Stdout, "Response from `ContentAPI.PutCmsV3SourceCodeEnvironmentContentPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environment** | **string** | The environment of the file (\&quot;draft\&quot; or \&quot;published\&quot;). | 
**path** | **string** | The file system location of the file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCmsV3SourceCodeEnvironmentContentPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | The file to upload. | 

### Return type

[**AssetFileMetadata**](AssetFileMetadata.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

