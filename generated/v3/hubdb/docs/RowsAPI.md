# \RowsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft**](RowsAPI.md#DeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft) | **Delete** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Permanently deletes a row
[**GetCmsV3HubdbTablesTableIdOrNameRows**](RowsAPI.md#GetCmsV3HubdbTablesTableIdOrNameRows) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows | Get rows for a table
[**GetCmsV3HubdbTablesTableIdOrNameRowsDraft**](RowsAPI.md#GetCmsV3HubdbTablesTableIdOrNameRowsDraft) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/draft | Get rows from draft table
[**GetCmsV3HubdbTablesTableIdOrNameRowsRowId**](RowsAPI.md#GetCmsV3HubdbTablesTableIdOrNameRowsRowId) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId} | Get a table row
[**GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft**](RowsAPI.md#GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Get a row from the draft table
[**PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft**](RowsAPI.md#PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft) | **Patch** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Updates an existing row
[**PostCmsV3HubdbTablesTableIdOrNameRows**](RowsAPI.md#PostCmsV3HubdbTablesTableIdOrNameRows) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows | Add a new row to a table
[**PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftClone**](RowsAPI.md#PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftClone) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft/clone | Clone a row
[**PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft**](RowsAPI.md#PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft) | **Put** /cms/v3/hubdb/tables/{tableIdOrName}/rows/{rowId}/draft | Replaces an existing row



## DeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft

> DeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft(ctx, tableIdOrName, rowId).Execute()

Permanently deletes a row



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
	rowId := "rowId_example" // string | The ID of the row

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RowsAPI.DeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft(context.Background(), tableIdOrName, rowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.DeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftRequest struct via the builder pattern


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


## GetCmsV3HubdbTablesTableIdOrNameRows

> CollectionResponseWithTotalHubDbTableRowV3ForwardPaging GetCmsV3HubdbTablesTableIdOrNameRows(ctx, tableIdOrName).Sort(sort).After(after).Limit(limit).Properties(properties).Execute()

Get rows for a table



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
	sort := []string{"Inner_example"} // []string | Specifies the column names to sort the results by. See the above description for more details. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is `1000`. (optional)
	properties := []string{"Inner_example"} // []string | Specify the column names to get results containing only the required columns instead of all column details. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRows(context.Background(), tableIdOrName).Sort(sort).After(after).Limit(limit).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTablesTableIdOrNameRows`: CollectionResponseWithTotalHubDbTableRowV3ForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **[]string** | Specifies the column names to sort the results by. See the above description for more details. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is &#x60;1000&#x60;. | 
 **properties** | **[]string** | Specify the column names to get results containing only the required columns instead of all column details. | 

### Return type

[**CollectionResponseWithTotalHubDbTableRowV3ForwardPaging**](CollectionResponseWithTotalHubDbTableRowV3ForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrNameRowsDraft

> CollectionResponseWithTotalHubDbTableRowV3ForwardPaging GetCmsV3HubdbTablesTableIdOrNameRowsDraft(ctx, tableIdOrName).Sort(sort).After(after).Limit(limit).Properties(properties).Execute()

Get rows from draft table



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
	sort := []string{"Inner_example"} // []string | Specifies the column names to sort the results by. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is `1000`. (optional)
	properties := []string{"Inner_example"} // []string | Specify the column names to get results containing only the required columns instead of all column details. If you want to include multiple columns in the result, use this query param as many times.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsDraft(context.Background(), tableIdOrName).Sort(sort).After(after).Limit(limit).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTablesTableIdOrNameRowsDraft`: CollectionResponseWithTotalHubDbTableRowV3ForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to query. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameRowsDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **[]string** | Specifies the column names to sort the results by. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is &#x60;1000&#x60;. | 
 **properties** | **[]string** | Specify the column names to get results containing only the required columns instead of all column details. If you want to include multiple columns in the result, use this query param as many times.  | 

### Return type

[**CollectionResponseWithTotalHubDbTableRowV3ForwardPaging**](CollectionResponseWithTotalHubDbTableRowV3ForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrNameRowsRowId

> HubDbTableRowV3 GetCmsV3HubdbTablesTableIdOrNameRowsRowId(ctx, tableIdOrName, rowId).Execute()

Get a table row



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
	rowId := "rowId_example" // string | The ID of the row

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsRowId(context.Background(), tableIdOrName, rowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsRowId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTablesTableIdOrNameRowsRowId`: HubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsRowId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameRowsRowIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft

> HubDbTableRowV3 GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft(ctx, tableIdOrName, rowId).Execute()

Get a row from the draft table



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
	rowId := "rowId_example" // string | The ID of the row

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft(context.Background(), tableIdOrName, rowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft`: HubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.GetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft

> HubDbTableRowV3 PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft(ctx, tableIdOrName, rowId).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()

Updates an existing row



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
	rowId := "rowId_example" // string | The ID of the row
	hubDbTableRowV3Request := *openapiclient.NewHubDbTableRowV3Request(map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // HubDbTableRowV3Request | The JSON object of the row with necessary fields that needs to be updated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft(context.Background(), tableIdOrName, rowId).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft`: HubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.PatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hubDbTableRowV3Request** | [**HubDbTableRowV3Request**](HubDbTableRowV3Request.md) | The JSON object of the row with necessary fields that needs to be updated. | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRows

> HubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRows(ctx, tableIdOrName).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()

Add a new row to a table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the target table.
	hubDbTableRowV3Request := *openapiclient.NewHubDbTableRowV3Request(map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // HubDbTableRowV3Request | The row definition JSON, formatted as described above.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.PostCmsV3HubdbTablesTableIdOrNameRows(context.Background(), tableIdOrName).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.PostCmsV3HubdbTablesTableIdOrNameRows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameRows`: HubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.PostCmsV3HubdbTablesTableIdOrNameRows`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the target table. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hubDbTableRowV3Request** | [**HubDbTableRowV3Request**](HubDbTableRowV3Request.md) | The row definition JSON, formatted as described above. | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftClone

> HubDbTableRowV3 PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftClone(ctx, tableIdOrName, rowId).Execute()

Clone a row



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
	rowId := "rowId_example" // string | The ID of the row

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftClone(context.Background(), tableIdOrName, rowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftClone`: HubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.PostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftClone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft

> HubDbTableRowV3 PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft(ctx, tableIdOrName, rowId).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()

Replaces an existing row



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
	rowId := "rowId_example" // string | The ID of the row
	hubDbTableRowV3Request := *openapiclient.NewHubDbTableRowV3Request(map[string]map[string]interface{}{"key": map[string]interface{}(123)}) // HubDbTableRowV3Request | The JSON object of the row

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RowsAPI.PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft(context.Background(), tableIdOrName, rowId).HubDbTableRowV3Request(hubDbTableRowV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RowsAPI.PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft`: HubDbTableRowV3
	fmt.Fprintf(os.Stdout, "Response from `RowsAPI.PutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table | 
**rowId** | **string** | The ID of the row | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCmsV3HubdbTablesTableIdOrNameRowsRowIdDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **hubDbTableRowV3Request** | [**HubDbTableRowV3Request**](HubDbTableRowV3Request.md) | The JSON object of the row | 

### Return type

[**HubDbTableRowV3**](HubDbTableRowV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

