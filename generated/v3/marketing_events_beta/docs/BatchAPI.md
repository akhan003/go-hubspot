# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMarketingV3MarketingEventsEventsDelete**](BatchAPI.md#PostMarketingV3MarketingEventsEventsDelete) | **Post** /marketing/v3/marketing-events/events/delete | Delete multiple marketing events
[**PostMarketingV3MarketingEventsEventsUpsert**](BatchAPI.md#PostMarketingV3MarketingEventsEventsUpsert) | **Post** /marketing/v3/marketing-events/events/upsert | Create or update multiple marketing events



## PostMarketingV3MarketingEventsEventsDelete

> Error PostMarketingV3MarketingEventsEventsDelete(ctx).BatchInputMarketingEventExternalUniqueIdentifier(batchInputMarketingEventExternalUniqueIdentifier).Execute()

Delete multiple marketing events



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
	batchInputMarketingEventExternalUniqueIdentifier := *openapiclient.NewBatchInputMarketingEventExternalUniqueIdentifier([]openapiclient.MarketingEventExternalUniqueIdentifier{*openapiclient.NewMarketingEventExternalUniqueIdentifier("ExternalAccountId_example", "ExternalEventId_example", int32(123))}) // BatchInputMarketingEventExternalUniqueIdentifier | The details of the marketing events to delete

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostMarketingV3MarketingEventsEventsDelete(context.Background()).BatchInputMarketingEventExternalUniqueIdentifier(batchInputMarketingEventExternalUniqueIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostMarketingV3MarketingEventsEventsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsEventsDelete`: Error
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostMarketingV3MarketingEventsEventsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputMarketingEventExternalUniqueIdentifier** | [**BatchInputMarketingEventExternalUniqueIdentifier**](BatchInputMarketingEventExternalUniqueIdentifier.md) | The details of the marketing events to delete | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3MarketingEventsEventsUpsert

> BatchResponseMarketingEventPublicDefaultResponse PostMarketingV3MarketingEventsEventsUpsert(ctx).BatchInputMarketingEventCreateRequestParams(batchInputMarketingEventCreateRequestParams).Execute()

Create or update multiple marketing events



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
	batchInputMarketingEventCreateRequestParams := *openapiclient.NewBatchInputMarketingEventCreateRequestParams([]openapiclient.MarketingEventCreateRequestParams{*openapiclient.NewMarketingEventCreateRequestParams("ExternalAccountId_example", "EventOrganizer_example", "ExternalEventId_example", "EventName_example")}) // BatchInputMarketingEventCreateRequestParams | The details of the marketing events to upsert

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.PostMarketingV3MarketingEventsEventsUpsert(context.Background()).BatchInputMarketingEventCreateRequestParams(batchInputMarketingEventCreateRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.PostMarketingV3MarketingEventsEventsUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsEventsUpsert`: BatchResponseMarketingEventPublicDefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.PostMarketingV3MarketingEventsEventsUpsert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputMarketingEventCreateRequestParams** | [**BatchInputMarketingEventCreateRequestParams**](BatchInputMarketingEventCreateRequestParams.md) | The details of the marketing events to upsert | 

### Return type

[**BatchResponseMarketingEventPublicDefaultResponse**](BatchResponseMarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

