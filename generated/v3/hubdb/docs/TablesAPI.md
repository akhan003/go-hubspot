# \TablesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3HubdbTablesTableIdOrName**](TablesAPI.md#DeleteCmsV3HubdbTablesTableIdOrName) | **Delete** /cms/v3/hubdb/tables/{tableIdOrName} | Archive a table
[**GetCmsV3HubdbTables**](TablesAPI.md#GetCmsV3HubdbTables) | **Get** /cms/v3/hubdb/tables | Get all published tables
[**GetCmsV3HubdbTablesDraft**](TablesAPI.md#GetCmsV3HubdbTablesDraft) | **Get** /cms/v3/hubdb/tables/draft | Return all draft tables
[**GetCmsV3HubdbTablesTableIdOrName**](TablesAPI.md#GetCmsV3HubdbTablesTableIdOrName) | **Get** /cms/v3/hubdb/tables/{tableIdOrName} | Get details for a published table
[**GetCmsV3HubdbTablesTableIdOrNameDraft**](TablesAPI.md#GetCmsV3HubdbTablesTableIdOrNameDraft) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/draft | Get details for a draft table
[**GetCmsV3HubdbTablesTableIdOrNameDraftExport**](TablesAPI.md#GetCmsV3HubdbTablesTableIdOrNameDraftExport) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/draft/export | Export a draft table
[**GetCmsV3HubdbTablesTableIdOrNameExport**](TablesAPI.md#GetCmsV3HubdbTablesTableIdOrNameExport) | **Get** /cms/v3/hubdb/tables/{tableIdOrName}/export | Export a published version of a table
[**PatchCmsV3HubdbTablesTableIdOrNameDraft**](TablesAPI.md#PatchCmsV3HubdbTablesTableIdOrNameDraft) | **Patch** /cms/v3/hubdb/tables/{tableIdOrName}/draft | Update an existing table
[**PostCmsV3HubdbTables**](TablesAPI.md#PostCmsV3HubdbTables) | **Post** /cms/v3/hubdb/tables | Create a new table
[**PostCmsV3HubdbTablesTableIdOrNameDraftClone**](TablesAPI.md#PostCmsV3HubdbTablesTableIdOrNameDraftClone) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/clone | Clone a table
[**PostCmsV3HubdbTablesTableIdOrNameDraftImport**](TablesAPI.md#PostCmsV3HubdbTablesTableIdOrNameDraftImport) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/import | Import data into draft table
[**PostCmsV3HubdbTablesTableIdOrNameDraftPublish**](TablesAPI.md#PostCmsV3HubdbTablesTableIdOrNameDraftPublish) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/publish | Publish a table from draft
[**PostCmsV3HubdbTablesTableIdOrNameDraftReset**](TablesAPI.md#PostCmsV3HubdbTablesTableIdOrNameDraftReset) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/draft/reset | Reset a draft table
[**PostCmsV3HubdbTablesTableIdOrNameUnpublish**](TablesAPI.md#PostCmsV3HubdbTablesTableIdOrNameUnpublish) | **Post** /cms/v3/hubdb/tables/{tableIdOrName}/unpublish | Unpublish a table



## DeleteCmsV3HubdbTablesTableIdOrName

> DeleteCmsV3HubdbTablesTableIdOrName(ctx, tableIdOrName).Execute()

Archive a table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to archive.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TablesAPI.DeleteCmsV3HubdbTablesTableIdOrName(context.Background(), tableIdOrName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.DeleteCmsV3HubdbTablesTableIdOrName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3HubdbTablesTableIdOrNameRequest struct via the builder pattern


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


## GetCmsV3HubdbTables

> CollectionResponseWithTotalHubDbTableV3ForwardPaging GetCmsV3HubdbTables(ctx).Sort(sort).After(after).Limit(limit).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Archived(archived).Execute()

Get all published tables



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 1000. (optional)
	createdAt := time.Now() // time.Time | Only return tables created at exactly the specified time. (optional)
	createdAfter := time.Now() // time.Time | Only return tables created after the specified time. (optional)
	createdBefore := time.Now() // time.Time | Only return tables created before the specified time. (optional)
	updatedAt := time.Now() // time.Time | Only return tables last updated at exactly the specified time. (optional)
	updatedAfter := time.Now() // time.Time | Only return tables last updated after the specified time. (optional)
	updatedBefore := time.Now() // time.Time | Only return tables last updated before the specified time. (optional)
	archived := true // bool | Specifies whether to return archived tables. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablesAPI.GetCmsV3HubdbTables(context.Background()).Sort(sort).After(after).Limit(limit).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.GetCmsV3HubdbTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTables`: CollectionResponseWithTotalHubDbTableV3ForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `TablesAPI.GetCmsV3HubdbTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 1000. | 
 **createdAt** | **time.Time** | Only return tables created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return tables created after the specified time. | 
 **createdBefore** | **time.Time** | Only return tables created before the specified time. | 
 **updatedAt** | **time.Time** | Only return tables last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return tables last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return tables last updated before the specified time. | 
 **archived** | **bool** | Specifies whether to return archived tables. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalHubDbTableV3ForwardPaging**](CollectionResponseWithTotalHubDbTableV3ForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesDraft

> CollectionResponseWithTotalHubDbTableV3ForwardPaging GetCmsV3HubdbTablesDraft(ctx).Sort(sort).After(after).Limit(limit).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Archived(archived).Execute()

Return all draft tables



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 1000. (optional)
	createdAt := time.Now() // time.Time | Only return tables created at exactly the specified time. (optional)
	createdAfter := time.Now() // time.Time | Only return tables created after the specified time. (optional)
	createdBefore := time.Now() // time.Time | Only return tables created before the specified time. (optional)
	updatedAt := time.Now() // time.Time | Only return tables last updated at exactly the specified time. (optional)
	updatedAfter := time.Now() // time.Time | Only return tables last updated after the specified time. (optional)
	updatedBefore := time.Now() // time.Time | Only return tables last updated before the specified time. (optional)
	archived := true // bool | Specifies whether to return archived tables. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablesAPI.GetCmsV3HubdbTablesDraft(context.Background()).Sort(sort).After(after).Limit(limit).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.GetCmsV3HubdbTablesDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTablesDraft`: CollectionResponseWithTotalHubDbTableV3ForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `TablesAPI.GetCmsV3HubdbTablesDraft`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 1000. | 
 **createdAt** | **time.Time** | Only return tables created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return tables created after the specified time. | 
 **createdBefore** | **time.Time** | Only return tables created before the specified time. | 
 **updatedAt** | **time.Time** | Only return tables last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return tables last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return tables last updated before the specified time. | 
 **archived** | **bool** | Specifies whether to return archived tables. Defaults to &#x60;false&#x60;. | 

### Return type

[**CollectionResponseWithTotalHubDbTableV3ForwardPaging**](CollectionResponseWithTotalHubDbTableV3ForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrName

> HubDbTableV3 GetCmsV3HubdbTablesTableIdOrName(ctx, tableIdOrName).IncludeForeignIds(includeForeignIds).Archived(archived).Execute()

Get details for a published table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to return.
	includeForeignIds := true // bool | Set this to `true` to populate foreign ID values in the result. (optional)
	archived := true // bool | Set this to `true` to return details for an archived table. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablesAPI.GetCmsV3HubdbTablesTableIdOrName(context.Background(), tableIdOrName).IncludeForeignIds(includeForeignIds).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.GetCmsV3HubdbTablesTableIdOrName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTablesTableIdOrName`: HubDbTableV3
	fmt.Fprintf(os.Stdout, "Response from `TablesAPI.GetCmsV3HubdbTablesTableIdOrName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the result. | 
 **archived** | **bool** | Set this to &#x60;true&#x60; to return details for an archived table. Defaults to &#x60;false&#x60;. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrNameDraft

> HubDbTableV3 GetCmsV3HubdbTablesTableIdOrNameDraft(ctx, tableIdOrName).IncludeForeignIds(includeForeignIds).Archived(archived).Execute()

Get details for a draft table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to return.
	includeForeignIds := true // bool | Set this to `true` to populate foreign ID values in the result. (optional)
	archived := true // bool | Set this to `true` to return an archived table. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablesAPI.GetCmsV3HubdbTablesTableIdOrNameDraft(context.Background(), tableIdOrName).IncludeForeignIds(includeForeignIds).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.GetCmsV3HubdbTablesTableIdOrNameDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTablesTableIdOrNameDraft`: HubDbTableV3
	fmt.Fprintf(os.Stdout, "Response from `TablesAPI.GetCmsV3HubdbTablesTableIdOrNameDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the result. | 
 **archived** | **bool** | Set this to &#x60;true&#x60; to return an archived table. Defaults to &#x60;false&#x60;. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrNameDraftExport

> *os.File GetCmsV3HubdbTablesTableIdOrNameDraftExport(ctx, tableIdOrName).Format(format).Execute()

Export a draft table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to export.
	format := "format_example" // string | The file format to export. Possible values include `CSV`, `XLSX`, and `XLS`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablesAPI.GetCmsV3HubdbTablesTableIdOrNameDraftExport(context.Background(), tableIdOrName).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.GetCmsV3HubdbTablesTableIdOrNameDraftExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTablesTableIdOrNameDraftExport`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TablesAPI.GetCmsV3HubdbTablesTableIdOrNameDraftExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameDraftExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** | The file format to export. Possible values include &#x60;CSV&#x60;, &#x60;XLSX&#x60;, and &#x60;XLS&#x60;. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ms-excel, text/csv, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3HubdbTablesTableIdOrNameExport

> *os.File GetCmsV3HubdbTablesTableIdOrNameExport(ctx, tableIdOrName).Format(format).Execute()

Export a published version of a table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to export.
	format := "format_example" // string | The file format to export. Possible values include `CSV`, `XLSX`, and `XLS`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablesAPI.GetCmsV3HubdbTablesTableIdOrNameExport(context.Background(), tableIdOrName).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.GetCmsV3HubdbTablesTableIdOrNameExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3HubdbTablesTableIdOrNameExport`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `TablesAPI.GetCmsV3HubdbTablesTableIdOrNameExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to export. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3HubdbTablesTableIdOrNameExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **string** | The file format to export. Possible values include &#x60;CSV&#x60;, &#x60;XLSX&#x60;, and &#x60;XLS&#x60;. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ms-excel, text/csv, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3HubdbTablesTableIdOrNameDraft

> HubDbTableV3 PatchCmsV3HubdbTablesTableIdOrNameDraft(ctx, tableIdOrName).HubDbTableV3Request(hubDbTableV3Request).IncludeForeignIds(includeForeignIds).Archived(archived).Execute()

Update an existing table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to update.
	hubDbTableV3Request := *openapiclient.NewHubDbTableV3Request("test_table", "Test Table") // HubDbTableV3Request | The JSON schema for the table being updated.
	includeForeignIds := true // bool | Set this to `true` to populate foreign ID values in the result. (optional)
	archived := true // bool | Specifies whether to return archived tables. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablesAPI.PatchCmsV3HubdbTablesTableIdOrNameDraft(context.Background(), tableIdOrName).HubDbTableV3Request(hubDbTableV3Request).IncludeForeignIds(includeForeignIds).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.PatchCmsV3HubdbTablesTableIdOrNameDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCmsV3HubdbTablesTableIdOrNameDraft`: HubDbTableV3
	fmt.Fprintf(os.Stdout, "Response from `TablesAPI.PatchCmsV3HubdbTablesTableIdOrNameDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3HubdbTablesTableIdOrNameDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hubDbTableV3Request** | [**HubDbTableV3Request**](HubDbTableV3Request.md) | The JSON schema for the table being updated. | 
 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the result. | 
 **archived** | **bool** | Specifies whether to return archived tables. Defaults to &#x60;false&#x60;. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTables

> HubDbTableV3 PostCmsV3HubdbTables(ctx).HubDbTableV3Request(hubDbTableV3Request).Execute()

Create a new table



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
	hubDbTableV3Request := *openapiclient.NewHubDbTableV3Request("test_table", "Test Table") // HubDbTableV3Request | The JSON schema for the table being created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablesAPI.PostCmsV3HubdbTables(context.Background()).HubDbTableV3Request(hubDbTableV3Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.PostCmsV3HubdbTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTables`: HubDbTableV3
	fmt.Fprintf(os.Stdout, "Response from `TablesAPI.PostCmsV3HubdbTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **hubDbTableV3Request** | [**HubDbTableV3Request**](HubDbTableV3Request.md) | The JSON schema for the table being created. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameDraftClone

> HubDbTableV3 PostCmsV3HubdbTablesTableIdOrNameDraftClone(ctx, tableIdOrName).HubDbTableCloneRequest(hubDbTableCloneRequest).Execute()

Clone a table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to clone.
	hubDbTableCloneRequest := *openapiclient.NewHubDbTableCloneRequest(true) // HubDbTableCloneRequest | JSON object with the properties newName and newLabel. You can set copyRows to false to clone the table with copying rows and default is true.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablesAPI.PostCmsV3HubdbTablesTableIdOrNameDraftClone(context.Background(), tableIdOrName).HubDbTableCloneRequest(hubDbTableCloneRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.PostCmsV3HubdbTablesTableIdOrNameDraftClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameDraftClone`: HubDbTableV3
	fmt.Fprintf(os.Stdout, "Response from `TablesAPI.PostCmsV3HubdbTablesTableIdOrNameDraftClone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to clone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameDraftCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hubDbTableCloneRequest** | [**HubDbTableCloneRequest**](HubDbTableCloneRequest.md) | JSON object with the properties newName and newLabel. You can set copyRows to false to clone the table with copying rows and default is true. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameDraftImport

> ImportResult PostCmsV3HubdbTablesTableIdOrNameDraftImport(ctx, tableIdOrName).Config(config).File(file).Execute()

Import data into draft table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID of the destination table where data will be imported.
	config := "config_example" // string | Configuration for the import in JSON format as described above. (optional)
	file := os.NewFile(1234, "some_file") // *os.File | The source CSV file to be imported. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablesAPI.PostCmsV3HubdbTablesTableIdOrNameDraftImport(context.Background(), tableIdOrName).Config(config).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.PostCmsV3HubdbTablesTableIdOrNameDraftImport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameDraftImport`: ImportResult
	fmt.Fprintf(os.Stdout, "Response from `TablesAPI.PostCmsV3HubdbTablesTableIdOrNameDraftImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID of the destination table where data will be imported. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameDraftImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **config** | **string** | Configuration for the import in JSON format as described above. | 
 **file** | ***os.File** | The source CSV file to be imported. | 

### Return type

[**ImportResult**](ImportResult.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameDraftPublish

> HubDbTableV3 PostCmsV3HubdbTablesTableIdOrNameDraftPublish(ctx, tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()

Publish a table from draft



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to publish.
	includeForeignIds := true // bool | Set this to `true` to populate foreign ID values in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablesAPI.PostCmsV3HubdbTablesTableIdOrNameDraftPublish(context.Background(), tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.PostCmsV3HubdbTablesTableIdOrNameDraftPublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameDraftPublish`: HubDbTableV3
	fmt.Fprintf(os.Stdout, "Response from `TablesAPI.PostCmsV3HubdbTablesTableIdOrNameDraftPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to publish. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameDraftPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the response. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameDraftReset

> HubDbTableV3 PostCmsV3HubdbTablesTableIdOrNameDraftReset(ctx, tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()

Reset a draft table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to reset.
	includeForeignIds := true // bool | Set this to `true` to populate foreign ID values in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablesAPI.PostCmsV3HubdbTablesTableIdOrNameDraftReset(context.Background(), tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.PostCmsV3HubdbTablesTableIdOrNameDraftReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameDraftReset`: HubDbTableV3
	fmt.Fprintf(os.Stdout, "Response from `TablesAPI.PostCmsV3HubdbTablesTableIdOrNameDraftReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to reset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameDraftResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the response. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3HubdbTablesTableIdOrNameUnpublish

> HubDbTableV3 PostCmsV3HubdbTablesTableIdOrNameUnpublish(ctx, tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()

Unpublish a table



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
	tableIdOrName := "tableIdOrName_example" // string | The ID or name of the table to publish.
	includeForeignIds := true // bool | Set this to `true` to populate foreign ID values in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TablesAPI.PostCmsV3HubdbTablesTableIdOrNameUnpublish(context.Background(), tableIdOrName).IncludeForeignIds(includeForeignIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TablesAPI.PostCmsV3HubdbTablesTableIdOrNameUnpublish``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3HubdbTablesTableIdOrNameUnpublish`: HubDbTableV3
	fmt.Fprintf(os.Stdout, "Response from `TablesAPI.PostCmsV3HubdbTablesTableIdOrNameUnpublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tableIdOrName** | **string** | The ID or name of the table to publish. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3HubdbTablesTableIdOrNameUnpublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeForeignIds** | **bool** | Set this to &#x60;true&#x60; to populate foreign ID values in the response. | 

### Return type

[**HubDbTableV3**](HubDbTableV3.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

