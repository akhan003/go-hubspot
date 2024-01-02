# \GDPRAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostCrmV3ObjectsFeedbackSubmissionsGdprDeletePurge**](GDPRAPI.md#PostCrmV3ObjectsFeedbackSubmissionsGdprDeletePurge) | **Post** /crm/v3/objects/feedback_submissions/gdpr-delete | GDPR DELETE



## PostCrmV3ObjectsFeedbackSubmissionsGdprDeletePurge

> PostCrmV3ObjectsFeedbackSubmissionsGdprDeletePurge(ctx).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()

GDPR DELETE



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
	publicGdprDeleteInput := *openapiclient.NewPublicGdprDeleteInput("ObjectId_example") // PublicGdprDeleteInput | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GDPRAPI.PostCrmV3ObjectsFeedbackSubmissionsGdprDeletePurge(context.Background()).PublicGdprDeleteInput(publicGdprDeleteInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GDPRAPI.PostCrmV3ObjectsFeedbackSubmissionsGdprDeletePurge``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ObjectsFeedbackSubmissionsGdprDeletePurgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicGdprDeleteInput** | [**PublicGdprDeleteInput**](PublicGdprDeleteInput.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

