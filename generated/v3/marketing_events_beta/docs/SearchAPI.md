# \SearchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarketingV3MarketingEventsEventsSearch**](SearchAPI.md#GetMarketingV3MarketingEventsEventsSearch) | **Get** /marketing/v3/marketing-events/events/search | Search for marketing events



## GetMarketingV3MarketingEventsEventsSearch

> CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging GetMarketingV3MarketingEventsEventsSearch(ctx).Q(q).Execute()

Search for marketing events



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
	q := "q_example" // string | The id of the marketing event in the external event application

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.GetMarketingV3MarketingEventsEventsSearch(context.Background()).Q(q).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetMarketingV3MarketingEventsEventsSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3MarketingEventsEventsSearch`: CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.GetMarketingV3MarketingEventsEventsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3MarketingEventsEventsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The id of the marketing event in the external event application | 

### Return type

[**CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging**](CollectionResponseMarketingEventExternalUniqueIdentifierNoPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

