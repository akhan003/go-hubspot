# \GenerateAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostConversationsV3VisitorIdentificationTokensCreate**](GenerateAPI.md#PostConversationsV3VisitorIdentificationTokensCreate) | **Post** /conversations/v3/visitor-identification/tokens/create | Generate a token



## PostConversationsV3VisitorIdentificationTokensCreate

> IdentificationTokenResponse PostConversationsV3VisitorIdentificationTokensCreate(ctx).IdentificationTokenGenerationRequest(identificationTokenGenerationRequest).Execute()

Generate a token



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
	identificationTokenGenerationRequest := *openapiclient.NewIdentificationTokenGenerationRequest("visitor-email@example.com") // IdentificationTokenGenerationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GenerateAPI.PostConversationsV3VisitorIdentificationTokensCreate(context.Background()).IdentificationTokenGenerationRequest(identificationTokenGenerationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GenerateAPI.PostConversationsV3VisitorIdentificationTokensCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostConversationsV3VisitorIdentificationTokensCreate`: IdentificationTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `GenerateAPI.PostConversationsV3VisitorIdentificationTokensCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostConversationsV3VisitorIdentificationTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identificationTokenGenerationRequest** | [**IdentificationTokenGenerationRequest**](IdentificationTokenGenerationRequest.md) |  | 

### Return type

[**IdentificationTokenResponse**](IdentificationTokenResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

