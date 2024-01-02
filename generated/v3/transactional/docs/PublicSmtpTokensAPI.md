# \PublicSmtpTokensAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMarketingV3TransactionalSmtpTokensTokenId**](PublicSmtpTokensAPI.md#DeleteMarketingV3TransactionalSmtpTokensTokenId) | **Delete** /marketing/v3/transactional/smtp-tokens/{tokenId} | Delete a single token by ID.
[**GetMarketingV3TransactionalSmtpTokens**](PublicSmtpTokensAPI.md#GetMarketingV3TransactionalSmtpTokens) | **Get** /marketing/v3/transactional/smtp-tokens | Query SMTP API tokens by campaign name or an emailCampaignId.
[**GetMarketingV3TransactionalSmtpTokensTokenId**](PublicSmtpTokensAPI.md#GetMarketingV3TransactionalSmtpTokensTokenId) | **Get** /marketing/v3/transactional/smtp-tokens/{tokenId} | Query a single token by ID.
[**PostMarketingV3TransactionalSmtpTokens**](PublicSmtpTokensAPI.md#PostMarketingV3TransactionalSmtpTokens) | **Post** /marketing/v3/transactional/smtp-tokens | Create a SMTP API token.
[**PostMarketingV3TransactionalSmtpTokensTokenIdPasswordReset**](PublicSmtpTokensAPI.md#PostMarketingV3TransactionalSmtpTokensTokenIdPasswordReset) | **Post** /marketing/v3/transactional/smtp-tokens/{tokenId}/password-reset | Reset the password of an existing token.



## DeleteMarketingV3TransactionalSmtpTokensTokenId

> DeleteMarketingV3TransactionalSmtpTokensTokenId(ctx, tokenId).Execute()

Delete a single token by ID.



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
	tokenId := "tokenId_example" // string | Identifier generated when a token is created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PublicSmtpTokensAPI.DeleteMarketingV3TransactionalSmtpTokensTokenId(context.Background(), tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensAPI.DeleteMarketingV3TransactionalSmtpTokensTokenId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Identifier generated when a token is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMarketingV3TransactionalSmtpTokensTokenIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetMarketingV3TransactionalSmtpTokens

> CollectionResponseSmtpApiTokenViewForwardPaging GetMarketingV3TransactionalSmtpTokens(ctx).CampaignName(campaignName).EmailCampaignId(emailCampaignId).After(after).Limit(limit).Execute()

Query SMTP API tokens by campaign name or an emailCampaignId.



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
	campaignName := "campaignName_example" // string | A name for the campaign tied to the SMTP API token. (optional)
	emailCampaignId := "emailCampaignId_example" // string | Identifier assigned to the campaign provided during the token creation. (optional)
	after := "after_example" // string | Starting point to get the next set of results. (optional)
	limit := int32(56) // int32 | Maximum number of tokens to return. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicSmtpTokensAPI.GetMarketingV3TransactionalSmtpTokens(context.Background()).CampaignName(campaignName).EmailCampaignId(emailCampaignId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensAPI.GetMarketingV3TransactionalSmtpTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3TransactionalSmtpTokens`: CollectionResponseSmtpApiTokenViewForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `PublicSmtpTokensAPI.GetMarketingV3TransactionalSmtpTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3TransactionalSmtpTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignName** | **string** | A name for the campaign tied to the SMTP API token. | 
 **emailCampaignId** | **string** | Identifier assigned to the campaign provided during the token creation. | 
 **after** | **string** | Starting point to get the next set of results. | 
 **limit** | **int32** | Maximum number of tokens to return. | 

### Return type

[**CollectionResponseSmtpApiTokenViewForwardPaging**](CollectionResponseSmtpApiTokenViewForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketingV3TransactionalSmtpTokensTokenId

> SmtpApiTokenView GetMarketingV3TransactionalSmtpTokensTokenId(ctx, tokenId).Execute()

Query a single token by ID.



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
	tokenId := "tokenId_example" // string | Identifier generated when a token is created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicSmtpTokensAPI.GetMarketingV3TransactionalSmtpTokensTokenId(context.Background(), tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensAPI.GetMarketingV3TransactionalSmtpTokensTokenId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketingV3TransactionalSmtpTokensTokenId`: SmtpApiTokenView
	fmt.Fprintf(os.Stdout, "Response from `PublicSmtpTokensAPI.GetMarketingV3TransactionalSmtpTokensTokenId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Identifier generated when a token is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketingV3TransactionalSmtpTokensTokenIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmtpApiTokenView**](SmtpApiTokenView.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3TransactionalSmtpTokens

> SmtpApiTokenView PostMarketingV3TransactionalSmtpTokens(ctx).SmtpApiTokenRequestEgg(smtpApiTokenRequestEgg).Execute()

Create a SMTP API token.



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
	smtpApiTokenRequestEgg := *openapiclient.NewSmtpApiTokenRequestEgg(false, "CampaignName_example") // SmtpApiTokenRequestEgg | A request object that includes the campaign name tied to the token and whether contacts should be created for email recipients.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicSmtpTokensAPI.PostMarketingV3TransactionalSmtpTokens(context.Background()).SmtpApiTokenRequestEgg(smtpApiTokenRequestEgg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensAPI.PostMarketingV3TransactionalSmtpTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3TransactionalSmtpTokens`: SmtpApiTokenView
	fmt.Fprintf(os.Stdout, "Response from `PublicSmtpTokensAPI.PostMarketingV3TransactionalSmtpTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3TransactionalSmtpTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **smtpApiTokenRequestEgg** | [**SmtpApiTokenRequestEgg**](SmtpApiTokenRequestEgg.md) | A request object that includes the campaign name tied to the token and whether contacts should be created for email recipients. | 

### Return type

[**SmtpApiTokenView**](SmtpApiTokenView.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMarketingV3TransactionalSmtpTokensTokenIdPasswordReset

> SmtpApiTokenView PostMarketingV3TransactionalSmtpTokensTokenIdPasswordReset(ctx, tokenId).Execute()

Reset the password of an existing token.



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
	tokenId := "tokenId_example" // string | Identifier generated when a token is created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicSmtpTokensAPI.PostMarketingV3TransactionalSmtpTokensTokenIdPasswordReset(context.Background(), tokenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicSmtpTokensAPI.PostMarketingV3TransactionalSmtpTokensTokenIdPasswordReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMarketingV3TransactionalSmtpTokensTokenIdPasswordReset`: SmtpApiTokenView
	fmt.Fprintf(os.Stdout, "Response from `PublicSmtpTokensAPI.PostMarketingV3TransactionalSmtpTokensTokenIdPasswordReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Identifier generated when a token is created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMarketingV3TransactionalSmtpTokensTokenIdPasswordResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SmtpApiTokenView**](SmtpApiTokenView.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

