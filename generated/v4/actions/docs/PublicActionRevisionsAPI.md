# \PublicActionRevisionsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutomationV4ActionsAppIdDefinitionIdRevisions**](PublicActionRevisionsAPI.md#GetAutomationV4ActionsAppIdDefinitionIdRevisions) | **Get** /automation/v4/actions/{appId}/{definitionId}/revisions | 
[**GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionId**](PublicActionRevisionsAPI.md#GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionId) | **Get** /automation/v4/actions/{appId}/{definitionId}/revisions/{revisionId} | 



## GetAutomationV4ActionsAppIdDefinitionIdRevisions

> CollectionResponsePublicActionRevisionForwardPaging GetAutomationV4ActionsAppIdDefinitionIdRevisions(ctx, definitionId, appId).Limit(limit).After(after).Execute()



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
	definitionId := "definitionId_example" // string | 
	appId := int32(56) // int32 | 
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicActionRevisionsAPI.GetAutomationV4ActionsAppIdDefinitionIdRevisions(context.Background(), definitionId, appId).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionRevisionsAPI.GetAutomationV4ActionsAppIdDefinitionIdRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppIdDefinitionIdRevisions`: CollectionResponsePublicActionRevisionForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `PublicActionRevisionsAPI.GetAutomationV4ActionsAppIdDefinitionIdRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The maximum number of results to display per page. | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 

### Return type

[**CollectionResponsePublicActionRevisionForwardPaging**](CollectionResponsePublicActionRevisionForwardPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionId

> PublicActionRevision GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionId(ctx, definitionId, revisionId, appId).Execute()



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
	definitionId := "definitionId_example" // string | 
	revisionId := "revisionId_example" // string | 
	appId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicActionRevisionsAPI.GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionId(context.Background(), definitionId, revisionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionRevisionsAPI.GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionId`: PublicActionRevision
	fmt.Fprintf(os.Stdout, "Response from `PublicActionRevisionsAPI.GetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**revisionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdRevisionsRevisionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PublicActionRevision**](PublicActionRevision.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

