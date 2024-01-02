# \DefinitionAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCommunicationPreferencesV3Definitions**](DefinitionAPI.md#GetCommunicationPreferencesV3Definitions) | **Get** /communication-preferences/v3/definitions | Get subscription definitions



## GetCommunicationPreferencesV3Definitions

> SubscriptionDefinitionsResponse GetCommunicationPreferencesV3Definitions(ctx).Execute()

Get subscription definitions



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefinitionAPI.GetCommunicationPreferencesV3Definitions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefinitionAPI.GetCommunicationPreferencesV3Definitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommunicationPreferencesV3Definitions`: SubscriptionDefinitionsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefinitionAPI.GetCommunicationPreferencesV3Definitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommunicationPreferencesV3DefinitionsRequest struct via the builder pattern


### Return type

[**SubscriptionDefinitionsResponse**](SubscriptionDefinitionsResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
