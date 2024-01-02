# \StatusAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCommunicationPreferencesV3StatusEmailEmailAddress**](StatusAPI.md#GetCommunicationPreferencesV3StatusEmailEmailAddress) | **Get** /communication-preferences/v3/status/email/{emailAddress} | Get subscription statuses for a contact
[**PostCommunicationPreferencesV3Subscribe**](StatusAPI.md#PostCommunicationPreferencesV3Subscribe) | **Post** /communication-preferences/v3/subscribe | Subscribe a contact
[**PostCommunicationPreferencesV3Unsubscribe**](StatusAPI.md#PostCommunicationPreferencesV3Unsubscribe) | **Post** /communication-preferences/v3/unsubscribe | Unsubscribe a contact



## GetCommunicationPreferencesV3StatusEmailEmailAddress

> PublicSubscriptionStatusesResponse GetCommunicationPreferencesV3StatusEmailEmailAddress(ctx, emailAddress).Execute()

Get subscription statuses for a contact



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
	emailAddress := "emailAddress_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.GetCommunicationPreferencesV3StatusEmailEmailAddress(context.Background(), emailAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.GetCommunicationPreferencesV3StatusEmailEmailAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommunicationPreferencesV3StatusEmailEmailAddress`: PublicSubscriptionStatusesResponse
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.GetCommunicationPreferencesV3StatusEmailEmailAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommunicationPreferencesV3StatusEmailEmailAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PublicSubscriptionStatusesResponse**](PublicSubscriptionStatusesResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommunicationPreferencesV3Subscribe

> PublicSubscriptionStatus PostCommunicationPreferencesV3Subscribe(ctx).PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest).Execute()

Subscribe a contact



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
	publicUpdateSubscriptionStatusRequest := *openapiclient.NewPublicUpdateSubscriptionStatusRequest("EmailAddress_example", "SubscriptionId_example") // PublicUpdateSubscriptionStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.PostCommunicationPreferencesV3Subscribe(context.Background()).PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.PostCommunicationPreferencesV3Subscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCommunicationPreferencesV3Subscribe`: PublicSubscriptionStatus
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.PostCommunicationPreferencesV3Subscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCommunicationPreferencesV3SubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicUpdateSubscriptionStatusRequest** | [**PublicUpdateSubscriptionStatusRequest**](PublicUpdateSubscriptionStatusRequest.md) |  | 

### Return type

[**PublicSubscriptionStatus**](PublicSubscriptionStatus.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCommunicationPreferencesV3Unsubscribe

> PublicSubscriptionStatus PostCommunicationPreferencesV3Unsubscribe(ctx).PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest).Execute()

Unsubscribe a contact



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
	publicUpdateSubscriptionStatusRequest := *openapiclient.NewPublicUpdateSubscriptionStatusRequest("EmailAddress_example", "SubscriptionId_example") // PublicUpdateSubscriptionStatusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.PostCommunicationPreferencesV3Unsubscribe(context.Background()).PublicUpdateSubscriptionStatusRequest(publicUpdateSubscriptionStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.PostCommunicationPreferencesV3Unsubscribe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCommunicationPreferencesV3Unsubscribe`: PublicSubscriptionStatus
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.PostCommunicationPreferencesV3Unsubscribe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCommunicationPreferencesV3UnsubscribeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicUpdateSubscriptionStatusRequest** | [**PublicUpdateSubscriptionStatusRequest**](PublicUpdateSubscriptionStatusRequest.md) |  | 

### Return type

[**PublicSubscriptionStatus**](PublicSubscriptionStatus.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

