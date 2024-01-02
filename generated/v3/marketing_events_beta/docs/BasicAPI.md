# \BasicAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingV3MarketingEventsEventsExternalEventId**](BasicAPI.md#DeleteMarketingV3MarketingEventsEventsExternalEventId) | **Delete** /marketing/v3/marketing-events/events/{externalEventId} | Delete a marketing event
[**GetMarketingV3MarketingEventsEventsExternalEventId**](BasicAPI.md#GetMarketingV3MarketingEventsEventsExternalEventId) | **Get** /marketing/v3/marketing-events/events/{externalEventId} | Get a marketing event
[**PatchMarketingV3MarketingEventsEventsExternalEventId**](BasicAPI.md#PatchMarketingV3MarketingEventsEventsExternalEventId) | **Patch** /marketing/v3/marketing-events/events/{externalEventId} | Update a marketing event
[**PostMarketingV3MarketingEventsEvents**](BasicAPI.md#PostMarketingV3MarketingEventsEvents) | **Post** /marketing/v3/marketing-events/events | Create a marketing event
[**PostMarketingV3MarketingEventsEventsExternalEventIdCancel**](BasicAPI.md#PostMarketingV3MarketingEventsEventsExternalEventIdCancel) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/cancel | Mark a marketing event as cancelled
[**PutMarketingV3MarketingEventsEventsExternalEventId**](BasicAPI.md#PutMarketingV3MarketingEventsEventsExternalEventId) | **Put** /marketing/v3/marketing-events/events/{externalEventId} | Create or update a marketing event



## DeleteMarketingV3MarketingEventsEventsExternalEventId

> DeleteMarketingV3MarketingEventsEventsExternalEventId(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()

Delete a marketing event



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
	externalEventId := "externalEventId_example" // string | The id of the marketing event to delete
	externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BasicAPI.DeleteMarketingV3MarketingEventsEventsExternalEventId(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.DeleteMarketingV3MarketingEventsEventsExternalEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingV3MarketingEventsEventsExternalEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The account id associated with the marketing event | 

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


## GetMarketingV3MarketingEventsEventsExternalEventId

> MarketingEventPublicReadResponse GetMarketingV3MarketingEventsEventsExternalEventId(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()

Get a marketing event



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
	externalEventId := "externalEventId_example" // string | The id of the marketing event to return
	externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.GetMarketingV3MarketingEventsEventsExternalEventId(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.GetMarketingV3MarketingEventsEventsExternalEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsEventsExternalEventId`: MarketingEventPublicReadResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.GetMarketingV3MarketingEventsEventsExternalEventId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsEventsExternalEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

[**MarketingEventPublicReadResponse**](MarketingEventPublicReadResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchMarketingV3MarketingEventsEventsExternalEventId

> MarketingEventPublicDefaultResponse PatchMarketingV3MarketingEventsEventsExternalEventId(ctx, externalEventId).ExternalAccountId(externalAccountId).MarketingEventUpdateRequestParams(marketingEventUpdateRequestParams).Execute()

Update a marketing event



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
	externalEventId := "externalEventId_example" // string | The id of the marketing event to update
	externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event
	marketingEventUpdateRequestParams := *openapiclient.NewMarketingEventUpdateRequestParams() // MarketingEventUpdateRequestParams | The details of the marketing event to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PatchMarketingV3MarketingEventsEventsExternalEventId(context.Background(), externalEventId).ExternalAccountId(externalAccountId).MarketingEventUpdateRequestParams(marketingEventUpdateRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PatchMarketingV3MarketingEventsEventsExternalEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchMarketingV3MarketingEventsEventsExternalEventId`: MarketingEventPublicDefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PatchMarketingV3MarketingEventsEventsExternalEventId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchMarketingV3MarketingEventsEventsExternalEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The account id associated with the marketing event | 
 **marketingEventUpdateRequestParams** | [**MarketingEventUpdateRequestParams**](MarketingEventUpdateRequestParams.md) | The details of the marketing event to update | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEvents

> MarketingEventDefaultResponse PostMarketingV3MarketingEventsEvents(ctx).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()

Create a marketing event



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
	marketingEventCreateRequestParams := *openapiclient.NewMarketingEventCreateRequestParams("ExternalAccountId_example", "EventOrganizer_example", "ExternalEventId_example", "EventName_example") // MarketingEventCreateRequestParams | The details of the marketing event to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostMarketingV3MarketingEventsEvents(context.Background()).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostMarketingV3MarketingEventsEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsEvents`: MarketingEventDefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostMarketingV3MarketingEventsEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketingEventCreateRequestParams** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md) | The details of the marketing event to create | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsExternalEventIdCancel

> MarketingEventDefaultResponse PostMarketingV3MarketingEventsEventsExternalEventIdCancel(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()

Mark a marketing event as cancelled



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
	externalEventId := "externalEventId_example" // string | The id of the marketing event to mark as cancelled
	externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PostMarketingV3MarketingEventsEventsExternalEventIdCancel(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PostMarketingV3MarketingEventsEventsExternalEventIdCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsEventsExternalEventIdCancel`: MarketingEventDefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PostMarketingV3MarketingEventsEventsExternalEventIdCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to mark as cancelled | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsExternalEventIdCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** | The account id associated with the marketing event | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutMarketingV3MarketingEventsEventsExternalEventId

> MarketingEventPublicDefaultResponse PutMarketingV3MarketingEventsEventsExternalEventId(ctx, externalEventId).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()

Create or update a marketing event



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
	externalEventId := "externalEventId_example" // string | The id of the marketing event to upsert
	marketingEventCreateRequestParams := *openapiclient.NewMarketingEventCreateRequestParams("ExternalAccountId_example", "EventOrganizer_example", "ExternalEventId_example", "EventName_example") // MarketingEventCreateRequestParams | The details of the marketing event to upsert

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BasicAPI.PutMarketingV3MarketingEventsEventsExternalEventId(context.Background(), externalEventId).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BasicAPI.PutMarketingV3MarketingEventsEventsExternalEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutMarketingV3MarketingEventsEventsExternalEventId`: MarketingEventPublicDefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `BasicAPI.PutMarketingV3MarketingEventsEventsExternalEventId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event to upsert | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutMarketingV3MarketingEventsEventsExternalEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketingEventCreateRequestParams** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md) | The details of the marketing event to upsert | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

