# \SubscriberStateChangesAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsert**](SubscriberStateChangesAPI.md#PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsert) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/{subscriberState}/email-upsert | Record
[**PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsert**](SubscriberStateChangesAPI.md#PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsert) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/{subscriberState}/upsert | Record



## PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsert

> Error PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsert(ctx, externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).Execute()

Record



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
	externalEventId := "externalEventId_example" // string | The id of the marketing event
	subscriberState := "subscriberState_example" // string | The new subscriber state for the HubSpot contacts and the specified marketing event
	externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event
	batchInputMarketingEventEmailSubscriber := *openapiclient.NewBatchInputMarketingEventEmailSubscriber([]openapiclient.MarketingEventEmailSubscriber{*openapiclient.NewMarketingEventEmailSubscriber("Email_example", int64(123))}) // BatchInputMarketingEventEmailSubscriber | The details of the contacts to subscribe to the event

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriberStateChangesAPI.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsert(context.Background(), externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriberStateChangesAPI.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsert`: Error
	fmt.Fprintf(os.Stdout, "Response from `SubscriberStateChangesAPI.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event | 
**subscriberState** | **string** | The new subscriber state for the HubSpot contacts and the specified marketing event | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateEmailUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **externalAccountId** | **string** | The account id associated with the marketing event | 
 **batchInputMarketingEventEmailSubscriber** | [**BatchInputMarketingEventEmailSubscriber**](BatchInputMarketingEventEmailSubscriber.md) | The details of the contacts to subscribe to the event | 

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


## PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsert

> Error PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsert(ctx, externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).Execute()

Record



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
	externalEventId := "externalEventId_example" // string | The id of the marketing event
	subscriberState := "subscriberState_example" // string | The new subscriber state for the HubSpot contacts and the specified marketing event
	externalAccountId := "externalAccountId_example" // string | The account id associated with the marketing event
	batchInputMarketingEventSubscriber := *openapiclient.NewBatchInputMarketingEventSubscriber([]openapiclient.MarketingEventSubscriber{*openapiclient.NewMarketingEventSubscriber(int64(123))}) // BatchInputMarketingEventSubscriber | The details of the contacts to subscribe to the event

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriberStateChangesAPI.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsert(context.Background(), externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriberStateChangesAPI.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsert``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsert`: Error
	fmt.Fprintf(os.Stdout, "Response from `SubscriberStateChangesAPI.PostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** | The id of the marketing event | 
**subscriberState** | **string** | The new subscriber state for the HubSpot contacts and the specified marketing event | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3MarketingEventsEventsExternalEventIdSubscriberStateUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **externalAccountId** | **string** | The account id associated with the marketing event | 
 **batchInputMarketingEventSubscriber** | [**BatchInputMarketingEventSubscriber**](BatchInputMarketingEventSubscriber.md) | The details of the contacts to subscribe to the event | 

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

