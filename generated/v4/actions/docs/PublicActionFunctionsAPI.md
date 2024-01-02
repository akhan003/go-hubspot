# \PublicActionFunctionsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType**](PublicActionFunctionsAPI.md#DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | 
[**DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId**](PublicActionFunctionsAPI.md#DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId) | **Delete** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | 
[**GetAutomationV4ActionsAppIdDefinitionIdFunctions**](PublicActionFunctionsAPI.md#GetAutomationV4ActionsAppIdDefinitionIdFunctions) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions | 
[**GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType**](PublicActionFunctionsAPI.md#GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | 
[**GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId**](PublicActionFunctionsAPI.md#GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId) | **Get** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | 
[**PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType**](PublicActionFunctionsAPI.md#PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType} | 
[**PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId**](PublicActionFunctionsAPI.md#PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId) | **Put** /automation/v4/actions/{appId}/{definitionId}/functions/{functionType}/{functionId} | 



## DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType

> DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType(ctx, definitionId, functionType, appId).Execute()



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
	functionType := "functionType_example" // string | 
	appId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicActionFunctionsAPI.DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType(context.Background(), definitionId, functionType, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionFunctionsAPI.DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**functionType** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeRequest struct via the builder pattern


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


## DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId

> DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId(ctx, definitionId, functionType, functionId, appId).Execute()



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
	functionType := "functionType_example" // string | 
	functionId := "functionId_example" // string | 
	appId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicActionFunctionsAPI.DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId(context.Background(), definitionId, functionType, functionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionFunctionsAPI.DeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**functionType** | **string** |  | 
**functionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdRequest struct via the builder pattern


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


## GetAutomationV4ActionsAppIdDefinitionIdFunctions

> CollectionResponsePublicActionFunctionIdentifierNoPaging GetAutomationV4ActionsAppIdDefinitionIdFunctions(ctx, definitionId, appId).Execute()



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
	resp, r, err := apiClient.PublicActionFunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctions(context.Background(), definitionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionFunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppIdDefinitionIdFunctions`: CollectionResponsePublicActionFunctionIdentifierNoPaging
	fmt.Fprintf(os.Stdout, "Response from `PublicActionFunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResponsePublicActionFunctionIdentifierNoPaging**](CollectionResponsePublicActionFunctionIdentifierNoPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType

> PublicActionFunction GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType(ctx, definitionId, functionType, appId).Execute()



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
	functionType := "functionType_example" // string | 
	appId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicActionFunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType(context.Background(), definitionId, functionType, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionFunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType`: PublicActionFunction
	fmt.Fprintf(os.Stdout, "Response from `PublicActionFunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**functionType** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**PublicActionFunction**](PublicActionFunction.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId

> PublicActionFunction GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId(ctx, definitionId, functionType, functionId, appId).Execute()



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
	functionType := "functionType_example" // string | 
	functionId := "functionId_example" // string | 
	appId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicActionFunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId(context.Background(), definitionId, functionType, functionId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionFunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId`: PublicActionFunction
	fmt.Fprintf(os.Stdout, "Response from `PublicActionFunctionsAPI.GetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**functionType** | **string** |  | 
**functionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PublicActionFunction**](PublicActionFunction.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType

> PublicActionFunctionIdentifier PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType(ctx, definitionId, functionType, appId).Body(body).Execute()



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
	functionType := "functionType_example" // string | 
	appId := int32(56) // int32 | 
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicActionFunctionsAPI.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType(context.Background(), definitionId, functionType, appId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionFunctionsAPI.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType`: PublicActionFunctionIdentifier
	fmt.Fprintf(os.Stdout, "Response from `PublicActionFunctionsAPI.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**functionType** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **string** |  | 

### Return type

[**PublicActionFunctionIdentifier**](PublicActionFunctionIdentifier.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId

> PublicActionFunctionIdentifier PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId(ctx, definitionId, functionType, functionId, appId).Body(body).Execute()



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
	functionType := "functionType_example" // string | 
	functionId := "functionId_example" // string | 
	appId := int32(56) // int32 | 
	body := "body_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicActionFunctionsAPI.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId(context.Background(), definitionId, functionType, functionId, appId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicActionFunctionsAPI.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId`: PublicActionFunctionIdentifier
	fmt.Fprintf(os.Stdout, "Response from `PublicActionFunctionsAPI.PutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**definitionId** | **string** |  | 
**functionType** | **string** |  | 
**functionId** | **string** |  | 
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAutomationV4ActionsAppIdDefinitionIdFunctionsFunctionTypeFunctionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | **string** |  | 

### Return type

[**PublicActionFunctionIdentifier**](PublicActionFunctionIdentifier.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

