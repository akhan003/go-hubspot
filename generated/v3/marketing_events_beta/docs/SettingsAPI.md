# \SettingsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketingV3MarketingEventsAppIdSettings**](SettingsAPI.md#GetMarketingV3MarketingEventsAppIdSettings) | **Get** /marketing/v3/marketing-events/{appId}/settings | Retrieve the application settings
[**PostMarketingV3MarketingEventsAppIdSettings**](SettingsAPI.md#PostMarketingV3MarketingEventsAppIdSettings) | **Post** /marketing/v3/marketing-events/{appId}/settings | Update the application settings



## GetMarketingV3MarketingEventsAppIdSettings

> EventDetailSettings GetMarketingV3MarketingEventsAppIdSettings(ctx, appId).Execute()

Retrieve the application settings



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
	appId := int32(56) // int32 | The id of the application to retrieve the settings for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.GetMarketingV3MarketingEventsAppIdSettings(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetMarketingV3MarketingEventsAppIdSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsAppIdSettings`: EventDetailSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetMarketingV3MarketingEventsAppIdSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The id of the application to retrieve the settings for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsAppIdSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventDetailSettings**](EventDetailSettings.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsAppIdSettings

> EventDetailSettings PostMarketingV3MarketingEventsAppIdSettings(ctx, appId).EventDetailSettingsUrl(eventDetailSettingsUrl).Execute()

Update the application settings



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
	appId := int32(56) // int32 | The id of the application to update the settings for.
	eventDetailSettingsUrl := *openapiclient.NewEventDetailSettingsUrl("EventDetailsUrl_example") // EventDetailSettingsUrl | The new application settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.PostMarketingV3MarketingEventsAppIdSettings(context.Background(), appId).EventDetailSettingsUrl(eventDetailSettingsUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PostMarketingV3MarketingEventsAppIdSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsAppIdSettings`: EventDetailSettings
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.PostMarketingV3MarketingEventsAppIdSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The id of the application to update the settings for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsAppIdSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eventDetailSettingsUrl** | [**EventDetailSettingsUrl**](EventDetailSettingsUrl.md) | The new application settings | 

### Return type

[**EventDetailSettings**](EventDetailSettings.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

