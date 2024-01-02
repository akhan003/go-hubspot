# \PublicActionDefinitionsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAutomationV4ActionsAppIdDefinitionId**](PublicActionDefinitionsAPI.md#DeleteAutomationV4ActionsAppIdDefinitionId) | **Delete** /automation/v4/actions/{appId}/{definitionId} | 
[**GetAutomationV4ActionsAppId**](PublicActionDefinitionsAPI.md#GetAutomationV4ActionsAppId) | **Get** /automation/v4/actions/{appId} | 
[**GetAutomationV4ActionsAppIdDefinitionId**](PublicActionDefinitionsAPI.md#GetAutomationV4ActionsAppIdDefinitionId) | **Get** /automation/v4/actions/{appId}/{definitionId} | 
[**PatchAutomationV4ActionsAppIdDefinitionId**](PublicActionDefinitionsAPI.md#PatchAutomationV4ActionsAppIdDefinitionId) | **Patch** /automation/v4/actions/{appId}/{definitionId} | 
[**PostAutomationV4ActionsAppId**](PublicActionDefinitionsAPI.md#PostAutomationV4ActionsAppId) | **Post** /automation/v4/actions/{appId} | 



## DeleteAutomationV4ActionsAppIdDefinitionId

> DeleteAutomationV4ActionsAppIdDefinitionId(ctx, definitionId, appId).Execute()



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicActionDefinitionsAPI.DeleteAutomationV4ActionsAppIdDefinitionId(context.Background(), definitionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionDefinitionsAPI.DeleteAutomationV4ActionsAppIdDefinitionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutomationV4ActionsAppIdDefinitionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationV4ActionsAppId

> CollectionResponsePublicActionDefinitionForwardPaging GetAutomationV4ActionsAppId(ctx, appId).Limit(limit).After(after).Archived(archived).Execute()



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
	appId := int32(56) // int32 | 
	limit := int32(56) // int32 | The maximum number of results to display per page. (optional)
	after := "after_example" // string | The paging cursor token of the last successfully read resource will be returned as the `paging.next.after` JSON property of a paged response containing more results. (optional)
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicActionDefinitionsAPI.GetAutomationV4ActionsAppId(context.Background(), appId).Limit(limit).After(after).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionDefinitionsAPI.GetAutomationV4ActionsAppId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppId`: CollectionResponsePublicActionDefinitionForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `PublicActionDefinitionsAPI.GetAutomationV4ActionsAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of results to display per page. | 
 **after** | **string** | The paging cursor token of the last successfully read resource will be returned as the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponsePublicActionDefinitionForwardPaging**](CollectionResponsePublicActionDefinitionForwardPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationV4ActionsAppIdDefinitionId

> PublicActionDefinition GetAutomationV4ActionsAppIdDefinitionId(ctx, definitionId, appId).Archived(archived).Execute()



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
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicActionDefinitionsAPI.GetAutomationV4ActionsAppIdDefinitionId(context.Background(), definitionId, appId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionDefinitionsAPI.GetAutomationV4ActionsAppIdDefinitionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppIdDefinitionId`: PublicActionDefinition
	fmt.Fprintf(os.Stdout, "Response from `PublicActionDefinitionsAPI.GetAutomationV4ActionsAppIdDefinitionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**PublicActionDefinition**](PublicActionDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchAutomationV4ActionsAppIdDefinitionId

> PublicActionDefinition PatchAutomationV4ActionsAppIdDefinitionId(ctx, definitionId, appId).PublicActionDefinitionPatch(publicActionDefinitionPatch).Execute()



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
	publicActionDefinitionPatch := *openapiclient.NewPublicActionDefinitionPatch() // PublicActionDefinitionPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicActionDefinitionsAPI.PatchAutomationV4ActionsAppIdDefinitionId(context.Background(), definitionId, appId).PublicActionDefinitionPatch(publicActionDefinitionPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionDefinitionsAPI.PatchAutomationV4ActionsAppIdDefinitionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchAutomationV4ActionsAppIdDefinitionId`: PublicActionDefinition
	fmt.Fprintf(os.Stdout, "Response from `PublicActionDefinitionsAPI.PatchAutomationV4ActionsAppIdDefinitionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchAutomationV4ActionsAppIdDefinitionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **publicActionDefinitionPatch** | [**PublicActionDefinitionPatch**](PublicActionDefinitionPatch.md) |  | 

### Return type

[**PublicActionDefinition**](PublicActionDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAutomationV4ActionsAppId

> PublicActionDefinition PostAutomationV4ActionsAppId(ctx, appId).PublicActionDefinitionEgg(publicActionDefinitionEgg).Execute()



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
	appId := int32(56) // int32 | 
	publicActionDefinitionEgg := *openapiclient.NewPublicActionDefinitionEgg([]openapiclient.InputFieldDefinition{*openapiclient.NewInputFieldDefinition(false, *openapiclient.NewFieldTypeDefinition("Name_example", []openapiclient.Option{*openapiclient.NewOption(false, int32(123), float32(123), "Description_example", false, "Label_example", "Value_example")}, "Type_example", false))}, []openapiclient.PublicActionFunction{*openapiclient.NewPublicActionFunction("FunctionSource_example", "FunctionType_example")}, "ActionUrl_example", false, []string{"ObjectTypes_example"}, map[string]PublicActionLabels{"key": *openapiclient.NewPublicActionLabels("ActionName_example")}) // PublicActionDefinitionEgg | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicActionDefinitionsAPI.PostAutomationV4ActionsAppId(context.Background(), appId).PublicActionDefinitionEgg(publicActionDefinitionEgg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionDefinitionsAPI.PostAutomationV4ActionsAppId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAutomationV4ActionsAppId`: PublicActionDefinition
	fmt.Fprintf(os.Stdout, "Response from `PublicActionDefinitionsAPI.PostAutomationV4ActionsAppId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAutomationV4ActionsAppIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicActionDefinitionEgg** | [**PublicActionDefinitionEgg**](PublicActionDefinitionEgg.md) |  | 

### Return type

[**PublicActionDefinition**](PublicActionDefinition.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

