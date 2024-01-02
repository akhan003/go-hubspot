# \EventsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3TimelineEventsEventTemplateIdEventId**](EventsAPI.md#GetCrmV3TimelineEventsEventTemplateIdEventId) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId} | Gets the event
[**GetCrmV3TimelineEventsEventTemplateIdEventIdDetail**](EventsAPI.md#GetCrmV3TimelineEventsEventTemplateIdEventIdDetail) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId}/detail | Gets the detailTemplate as rendered
[**GetCrmV3TimelineEventsEventTemplateIdEventIdRender**](EventsAPI.md#GetCrmV3TimelineEventsEventTemplateIdEventIdRender) | **Get** /crm/v3/timeline/events/{eventTemplateId}/{eventId}/render | Renders the header or detail as HTML
[**PostCrmV3TimelineEvents**](EventsAPI.md#PostCrmV3TimelineEvents) | **Post** /crm/v3/timeline/events | Create a single event
[**PostCrmV3TimelineEventsBatchCreate**](EventsAPI.md#PostCrmV3TimelineEventsBatchCreate) | **Post** /crm/v3/timeline/events/batch/create | Creates multiple events



## GetCrmV3TimelineEventsEventTemplateIdEventId

> TimelineEventResponse GetCrmV3TimelineEventsEventTemplateIdEventId(ctx, eventTemplateId, eventId).Execute()

Gets the event



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
	eventId := "eventId_example" // string | The event ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventId(context.Background(), eventTemplateId, eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3TimelineEventsEventTemplateIdEventId`: TimelineEventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**eventId** | **string** | The event ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3TimelineEventsEventTemplateIdEventIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TimelineEventResponse**](TimelineEventResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3TimelineEventsEventTemplateIdEventIdDetail

> EventDetail GetCrmV3TimelineEventsEventTemplateIdEventIdDetail(ctx, eventTemplateId, eventId).Execute()

Gets the detailTemplate as rendered



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
	eventId := "eventId_example" // string | The event ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdDetail(context.Background(), eventTemplateId, eventId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3TimelineEventsEventTemplateIdEventIdDetail`: EventDetail
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**eventId** | **string** | The event ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3TimelineEventsEventTemplateIdEventIdDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EventDetail**](EventDetail.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3TimelineEventsEventTemplateIdEventIdRender

> string GetCrmV3TimelineEventsEventTemplateIdEventIdRender(ctx, eventTemplateId, eventId).Detail(detail).Execute()

Renders the header or detail as HTML



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
	eventId := "eventId_example" // string | The event ID.
	detail := true // bool | Set to 'true', we want to render the `detailTemplate` instead of the `headerTemplate`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdRender(context.Background(), eventTemplateId, eventId).Detail(detail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdRender``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3TimelineEventsEventTemplateIdEventIdRender`: string
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetCrmV3TimelineEventsEventTemplateIdEventIdRender`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventTemplateId** | **string** | The event template ID. | 
**eventId** | **string** | The event ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3TimelineEventsEventTemplateIdEventIdRenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **detail** | **bool** | Set to &#39;true&#39;, we want to render the &#x60;detailTemplate&#x60; instead of the &#x60;headerTemplate&#x60;. | 

### Return type

**string**

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3TimelineEvents

> TimelineEventResponse PostCrmV3TimelineEvents(ctx).TimelineEvent(timelineEvent).Execute()

Create a single event



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
	timelineEvent := *openapiclient.NewTimelineEvent("1001298", map[string]string{"key": "Inner_example"}) // TimelineEvent | The timeline event definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventsAPI.PostCrmV3TimelineEvents(context.Background()).TimelineEvent(timelineEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostCrmV3TimelineEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3TimelineEvents`: TimelineEventResponse
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.PostCrmV3TimelineEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3TimelineEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timelineEvent** | [**TimelineEvent**](TimelineEvent.md) | The timeline event definition. | 

### Return type

[**TimelineEventResponse**](TimelineEventResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3TimelineEventsBatchCreate

> PostCrmV3TimelineEventsBatchCreate(ctx).BatchInputTimelineEvent(batchInputTimelineEvent).Execute()

Creates multiple events



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
	batchInputTimelineEvent := *openapiclient.NewBatchInputTimelineEvent([]openapiclient.TimelineEvent{*openapiclient.NewTimelineEvent("1001298", map[string]string{"key": "Inner_example"})}) // BatchInputTimelineEvent | The timeline event definition.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EventsAPI.PostCrmV3TimelineEventsBatchCreate(context.Background()).BatchInputTimelineEvent(batchInputTimelineEvent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.PostCrmV3TimelineEventsBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3TimelineEventsBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputTimelineEvent** | [**BatchInputTimelineEvent**](BatchInputTimelineEvent.md) | The timeline event definition. | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

