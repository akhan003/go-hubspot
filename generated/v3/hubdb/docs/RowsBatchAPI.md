# \RowsBatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCmsV3HubdbTablesTableIdOrNameRowsBatchRead**](RowsBatchAPI.md#PostCmsV3HubdbTablesTableIdOrNameRowsBatchRead) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/batch/read | Get a set of rows
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchClone**](RowsBatchAPI.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchClone) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/clone | Clone rows in batch
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreate**](RowsBatchAPI.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreate) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/create | Create rows in batch
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurge**](RowsBatchAPI.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurge) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/purge | Permanently deletes rows
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchRead**](RowsBatchAPI.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchRead) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/read | Get a set of rows from draft table
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplace**](RowsBatchAPI.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplace) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/replace | Replace rows in batch in draft table
[**PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdate**](RowsBatchAPI.md#PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdate) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft/batch/update | Update rows in batch in draft table



## PostCmsV3HubdbTablesTableIdOrNameRowsBatchRead

> BatchResponseHubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsBatchRead(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

Get a set of rows



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to query.
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of row ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsBatchRead(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameRowsBatchRead`: BatchResponseHubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsBatchRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of row ids | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchClone

> BatchResponseHubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchClone(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

Clone rows in batch



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of row ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchClone(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchClone`: BatchResponseHubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchClone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of row ids | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreate

> BatchResponseHubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreate(ctx, tableIdOrName).BatchInputHubDbTableRowV3Request(batchInputHubDbTableRowV3Request).Execute()

Create rows in batch



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
	batchInputHubDbTableRowV3Request := *openapiclient.NewBatchInputHubDbTableRowV3Request([]openapiclient.HubDbTableRowV3Request{*openapiclient.NewHubDbTableRowV3Request(map[string]map[string]interface{}{"key": map[string]interface{}(123)})}) // BatchInputHubDbTableRowV3Request | JSON array of row objects

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreate(context.Background(), tableIdOrName).BatchInputHubDbTableRowV3Request(batchInputHubDbTableRowV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreate`: BatchResponseHubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputHubDbTableRowV3Request** | [**BatchInputHubDbTableRowV3Request**](BatchInputHubDbTableRowV3Request.md) | JSON array of row objects | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurge

> PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurge(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

Permanently deletes rows



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | JSON array of row ids.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurge(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchPurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | JSON array of row ids. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchRead

> BatchResponseHubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchRead(ctx, tableIdOrName).BatchInputString(batchInputString).Execute()

Get a set of rows from draft table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | JSON array of row ids.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchRead(context.Background(), tableIdOrName).BatchInputString(batchInputString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchRead`: BatchResponseHubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputString** | [**BatchInputString**](BatchInputString.md) | JSON array of row ids. | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplace

> BatchResponseHubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplace(ctx, tableIdOrName).BatchInputHubDbTableRowV3BatchUpdateRequest(batchInputHubDbTableRowV3BatchUpdateRequest).Execute()

Replace rows in batch in draft table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
	batchInputHubDbTableRowV3BatchUpdateRequest := *openapiclient.NewBatchInputHubDbTableRowV3BatchUpdateRequest([]openapiclient.HubDbTableRowV3BatchUpdateRequest{*openapiclient.NewHubDbTableRowV3BatchUpdateRequest(map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "4099275931")}) // BatchInputHubDbTableRowV3BatchUpdateRequest | JSON array of row objects.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplace(context.Background(), tableIdOrName).BatchInputHubDbTableRowV3BatchUpdateRequest(batchInputHubDbTableRowV3BatchUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplace`: BatchResponseHubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputHubDbTableRowV3BatchUpdateRequest** | [**BatchInputHubDbTableRowV3BatchUpdateRequest**](BatchInputHubDbTableRowV3BatchUpdateRequest.md) | JSON array of row objects. | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdate

> BatchResponseHubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdate(ctx, tableIdOrName).BatchInputHubDbTableRowV3BatchUpdateRequest(batchInputHubDbTableRowV3BatchUpdateRequest).Execute()

Update rows in batch in draft table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table
	batchInputHubDbTableRowV3BatchUpdateRequest := *openapiclient.NewBatchInputHubDbTableRowV3BatchUpdateRequest([]openapiclient.HubDbTableRowV3BatchUpdateRequest{*openapiclient.NewHubDbTableRowV3BatchUpdateRequest(map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "4099275931")}) // BatchInputHubDbTableRowV3BatchUpdateRequest | JSON array of row objects.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdate(context.Background(), tableIdOrName).BatchInputHubDbTableRowV3BatchUpdateRequest(batchInputHubDbTableRowV3BatchUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdate`: BatchResponseHubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsBatchAPI.PostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsDraftBatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchInputHubDbTableRowV3BatchUpdateRequest** | [**BatchInputHubDbTableRowV3BatchUpdateRequest**](BatchInputHubDbTableRowV3BatchUpdateRequest.md) | JSON array of row objects. | 

### Return type

[**BatchResponseHubDbTableRowV3**](BatchResponseHubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

