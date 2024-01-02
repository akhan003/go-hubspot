# \PipelinesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Archive**](PipelinesAPI.md#Archive) | **Delete** /crm/v3/pipelines/{objectType}/{pipelineId} | Delete a pipeline
[**Create**](PipelinesAPI.md#Create) | **Post** /crm/v3/pipelines/{objectType} | Create a pipeline
[**GetAll**](PipelinesAPI.md#GetAll) | **Get** /crm/v3/pipelines/{objectType} | Retrieve all pipelines
[**GetByID**](PipelinesAPI.md#GetByID) | **Get** /crm/v3/pipelines/{objectType}/{pipelineId} | Return a pipeline by ID
[**Replace**](PipelinesAPI.md#Replace) | **Put** /crm/v3/pipelines/{objectType}/{pipelineId} | Replace a pipeline
[**Update**](PipelinesAPI.md#Update) | **Patch** /crm/v3/pipelines/{objectType}/{pipelineId} | Update a pipeline



## Archive

> Archive(ctx, objectType, pipelineId).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()

Delete a pipeline



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
	pipelineId := "pipelineId_example" // string | 
	validateReferencesBeforeDelete := true // bool |  (optional) (default to false)
	validateDealStageUsagesBeforeDelete := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PipelinesAPI.Archive(context.Background(), objectType, pipelineId).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.Archive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **validateReferencesBeforeDelete** | **bool** |  | [default to false]
 **validateDealStageUsagesBeforeDelete** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> Pipeline Create(ctx, objectType).PipelineInput(pipelineInput).Execute()

Create a pipeline



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
	pipelineInput := *openapiclient.NewPipelineInput(int32(0), []openapiclient.PipelineStageInput{*openapiclient.NewPipelineStageInput(map[string]string{"key": "Inner_example"}, int32(1), "Done")}, "My replaced pipeline") // PipelineInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.Create(context.Background(), objectType).PipelineInput(pipelineInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pipelineInput** | [**PipelineInput**](PipelineInput.md) |  | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> CollectionResponsePipelineNoPaging GetAll(ctx, objectType).Execute()

Retrieve all pipelines



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
	resp, r, err := apiClient.PipelinesAPI.GetAll(context.Background(), objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.GetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAll`: CollectionResponsePipelineNoPaging
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.GetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionResponsePipelineNoPaging**](CollectionResponsePipelineNoPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByID

> Pipeline GetByID(ctx, objectType, pipelineId).Execute()

Return a pipeline by ID



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
	pipelineId := "pipelineId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.GetByID(context.Background(), objectType, pipelineId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.GetByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByID`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.GetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Replace

> Pipeline Replace(ctx, objectType, pipelineId).PipelineInput(pipelineInput).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()

Replace a pipeline



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
	pipelineId := "pipelineId_example" // string | 
	pipelineInput := *openapiclient.NewPipelineInput(int32(0), []openapiclient.PipelineStageInput{*openapiclient.NewPipelineStageInput(map[string]string{"key": "Inner_example"}, int32(1), "Done")}, "My replaced pipeline") // PipelineInput | 
	validateReferencesBeforeDelete := true // bool |  (optional) (default to false)
	validateDealStageUsagesBeforeDelete := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.Replace(context.Background(), objectType, pipelineId).PipelineInput(pipelineInput).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.Replace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Replace`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.Replace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelineInput** | [**PipelineInput**](PipelineInput.md) |  | 
 **validateReferencesBeforeDelete** | **bool** |  | [default to false]
 **validateDealStageUsagesBeforeDelete** | **bool** |  | [default to false]

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Pipeline Update(ctx, objectType, pipelineId).PipelinePatchInput(pipelinePatchInput).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()

Update a pipeline



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
	pipelineId := "pipelineId_example" // string | 
	pipelinePatchInput := *openapiclient.NewPipelinePatchInput() // PipelinePatchInput | 
	validateReferencesBeforeDelete := true // bool |  (optional) (default to false)
	validateDealStageUsagesBeforeDelete := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PipelinesAPI.Update(context.Background(), objectType, pipelineId).PipelinePatchInput(pipelinePatchInput).ValidateReferencesBeforeDelete(validateReferencesBeforeDelete).ValidateDealStageUsagesBeforeDelete(validateDealStageUsagesBeforeDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PipelinesAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: Pipeline
	fmt.Fprintf(os.Stdout, "Response from `PipelinesAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**pipelineId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelinePatchInput** | [**PipelinePatchInput**](PipelinePatchInput.md) |  | 
 **validateReferencesBeforeDelete** | **bool** |  | [default to false]
 **validateDealStageUsagesBeforeDelete** | **bool** |  | [default to false]

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

