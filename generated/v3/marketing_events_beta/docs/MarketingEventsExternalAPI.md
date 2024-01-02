# \MarketingEventsExternalAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExternalComplete**](MarketingEventsExternalAPI.md#ExternalComplete) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/complete | 



## ExternalComplete

> MarketingEventDefaultResponse ExternalComplete(ctx, externalEventId).ExternalAccountId(externalAccountId).MarketingEventCompleteRequestParams(marketingEventCompleteRequestParams).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	externalEventId := "externalEventId_example" // string | 
	externalAccountId := "externalAccountId_example" // string | 
	marketingEventCompleteRequestParams := *openapiclient.NewMarketingEventCompleteRequestParams(time.Now(), time.Now()) // MarketingEventCompleteRequestParams | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketingEventsExternalAPI.ExternalComplete(context.Background(), externalEventId).ExternalAccountId(externalAccountId).MarketingEventCompleteRequestParams(marketingEventCompleteRequestParams).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalAPI.ExternalComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExternalComplete`: MarketingEventDefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalAPI.ExternalComplete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** |  | 
 **marketingEventCompleteRequestParams** | [**MarketingEventCompleteRequestParams**](MarketingEventCompleteRequestParams.md) |  | 

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

