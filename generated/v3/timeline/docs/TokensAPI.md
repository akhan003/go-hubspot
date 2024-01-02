# \TokensAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName**](TokensAPI.md#DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName) | **Delete** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName} | Removes a token from the event template
[**PostCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokens**](TokensAPI.md#PostCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokens) | **Post** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens | Adds a token to an existing event template
[**PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName**](TokensAPI.md#PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName) | **Put** /crm/v3/timeline/{appId}/event-templates/{eventTemplateId}/tokens/{tokenName} | Updates an existing token on an event template



## DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName

> DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName(ctx, eventTemplateId, tokenName, appId).Execute()

Removes a token from the event template



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
	tokenName := "tokenName_example" // string | The token name.
	appId := int32(56) // int32 | The ID of the target app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TokensAPI.DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName(context.Background(), eventTemplateId, tokenName, appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.DeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**tokenName** | **string** | The token name. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenNameRequest struct via the builder pattern


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


## PostCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokens

> TimelineEventTemplateToken PostCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokens(ctx, eventTemplateId, appId).TimelineEventTemplateToken(timelineEventTemplateToken).Execute()

Adds a token to an existing event template



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
	timelineEventTemplateToken := *openapiclient.NewTimelineEventTemplateToken("petType", "Pet Type", "enumeration") // TimelineEventTemplateToken | The new token definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.PostCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokens(context.Background(), eventTemplateId, appId).TimelineEventTemplateToken(timelineEventTemplateToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.PostCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokens`: TimelineEventTemplateToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.PostCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokens`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timelineEventTemplateToken** | [**TimelineEventTemplateToken**](TimelineEventTemplateToken.md) | The new token definition. | 

### Return type

[**TimelineEventTemplateToken**](TimelineEventTemplateToken.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName

> TimelineEventTemplateToken PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName(ctx, eventTemplateId, tokenName, appId).TimelineEventTemplateTokenUpdateRequest(timelineEventTemplateTokenUpdateRequest).Execute()

Updates an existing token on an event template



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
	tokenName := "tokenName_example" // string | The token name.
	appId := int32(56) // int32 | The ID of the target app.
	timelineEventTemplateTokenUpdateRequest := *openapiclient.NewTimelineEventTemplateTokenUpdateRequest("petType edit") // TimelineEventTemplateTokenUpdateRequest | The updated token definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName(context.Background(), eventTemplateId, tokenName, appId).TimelineEventTemplateTokenUpdateRequest(timelineEventTemplateTokenUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName`: TimelineEventTemplateToken
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.PutCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**tokenName** | **string** | The token name. | 
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3TimelineAppIdEventTemplatesEventTemplateIdTokensTokenNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **timelineEventTemplateTokenUpdateRequest** | [**TimelineEventTemplateTokenUpdateRequest**](TimelineEventTemplateTokenUpdateRequest.md) | The updated token definition. | 

### Return type

[**TimelineEventTemplateToken**](TimelineEventTemplateToken.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

