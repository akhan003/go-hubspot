# \TemplatesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateId**](TemplatesAPI.md#DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateId) | **Delete** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Deletes an event template for the app
[**GetCrmV3TimelineAppIdEventTemplates**](TemplatesAPI.md#GetCrmV3TimelineAppIdEventTemplates) | **Get** /crm/v3/timeline/{appId}/event-templates | List all event templates for your app
[**GetCrmV3TimelineAppIdEventTemplatesEventTemplateId**](TemplatesAPI.md#GetCrmV3TimelineAppIdEventTemplatesEventTemplateId) | **Get** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Gets a specific event template for your app
[**PostCrmV3TimelineAppIdEventTemplates**](TemplatesAPI.md#PostCrmV3TimelineAppIdEventTemplates) | **Post** /crm/v3/timeline/{appId}/event-templates | Create an event template for your app
[**PutCrmV3TimelineAppIdEventTemplatesEventTemplateId**](TemplatesAPI.md#PutCrmV3TimelineAppIdEventTemplatesEventTemplateId) | **Put** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId} | Update an existing event template



## DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateId

> DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateId(ctx, eventTemplateId, appId).Execute()

Deletes an event template for the app



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
	eventTemplateId := "eventTemplateId_example" // string | The event template ID.
	appId := int32(56) // int32 | The ID of the target app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TemplatesAPI.DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateId(context.Background(), eventTemplateId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdRequest struct via the builder pattern


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


## GetCrmV3TimelineAppIdEventTemplates

> CollectionResponseTimelineEventTemplateNoPaging GetCrmV3TimelineAppIdEventTemplates(ctx, appId).Execute()

List all event templates for your app



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
	appId := int32(56) // int32 | The ID of the target app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetCrmV3TimelineAppIdEventTemplates(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetCrmV3TimelineAppIdEventTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3TimelineAppIdEventTemplates`: CollectionResponseTimelineEventTemplateNoPaging
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetCrmV3TimelineAppIdEventTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3TimelineAppIdEventTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionResponseTimelineEventTemplateNoPaging**](CollectionResponseTimelineEventTemplateNoPaging.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3TimelineAppIdEventTemplatesEventTemplateId

> TimelineEventTemplate GetCrmV3TimelineAppIdEventTemplatesEventTemplateId(ctx, eventTemplateId, appId).Execute()

Gets a specific event template for your app



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
	eventTemplateId := "eventTemplateId_example" // string | The event template ID.
	appId := int32(56) // int32 | The ID of the target app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.GetCrmV3TimelineAppIdEventTemplatesEventTemplateId(context.Background(), eventTemplateId, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.GetCrmV3TimelineAppIdEventTemplatesEventTemplateId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3TimelineAppIdEventTemplatesEventTemplateId`: TimelineEventTemplate
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.GetCrmV3TimelineAppIdEventTemplatesEventTemplateId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3TimelineAppIdEventTemplatesEventTemplateIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TimelineEventTemplate**](TimelineEventTemplate.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3TimelineAppIdEventTemplates

> TimelineEventTemplate PostCrmV3TimelineAppIdEventTemplates(ctx, appId).TimelineEventTemplateCreateRequest(timelineEventTemplateCreateRequest).Execute()

Create an event template for your app



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
	appId := int32(56) // int32 | The ID of the target app.
	timelineEventTemplateCreateRequest := *openapiclient.NewTimelineEventTemplateCreateRequest("PetSpot Registration", []openapiclient.TimelineEventTemplateToken{*openapiclient.NewTimelineEventTemplateToken("petType", "Pet Type", "enumeration")}, "contacts") // TimelineEventTemplateCreateRequest | The new event template definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.PostCrmV3TimelineAppIdEventTemplates(context.Background(), appId).TimelineEventTemplateCreateRequest(timelineEventTemplateCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.PostCrmV3TimelineAppIdEventTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3TimelineAppIdEventTemplates`: TimelineEventTemplate
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.PostCrmV3TimelineAppIdEventTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3TimelineAppIdEventTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timelineEventTemplateCreateRequest** | [**TimelineEventTemplateCreateRequest**](TimelineEventTemplateCreateRequest.md) | The new event template definition. | 

### Return type

[**TimelineEventTemplate**](TimelineEventTemplate.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3TimelineAppIdEventTemplatesEventTemplateId

> TimelineEventTemplate PutCrmV3TimelineAppIdEventTemplatesEventTemplateId(ctx, eventTemplateId, appId).TimelineEventTemplateUpdateRequest(timelineEventTemplateUpdateRequest).Execute()

Update an existing event template



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
	eventTemplateId := "eventTemplateId_example" // string | The event template ID.
	appId := int32(56) // int32 | The ID of the target app.
	timelineEventTemplateUpdateRequest := *openapiclient.NewTimelineEventTemplateUpdateRequest("PetSpot Registration", []openapiclient.TimelineEventTemplateToken{*openapiclient.NewTimelineEventTemplateToken("petType", "Pet Type", "enumeration")}, "1001298") // TimelineEventTemplateUpdateRequest | The updated event template definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TemplatesAPI.PutCrmV3TimelineAppIdEventTemplatesEventTemplateId(context.Background(), eventTemplateId, appId).TimelineEventTemplateUpdateRequest(timelineEventTemplateUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TemplatesAPI.PutCrmV3TimelineAppIdEventTemplatesEventTemplateId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmV3TimelineAppIdEventTemplatesEventTemplateId`: TimelineEventTemplate
	fmt.Fprintf(os.Stdout, "Response from `TemplatesAPI.PutCrmV3TimelineAppIdEventTemplatesEventTemplateId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3TimelineAppIdEventTemplatesEventTemplateIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timelineEventTemplateUpdateRequest** | [**TimelineEventTemplateUpdateRequest**](TimelineEventTemplateUpdateRequest.md) | The updated event template definition. | 

### Return type

[**TimelineEventTemplate**](TimelineEventTemplate.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

