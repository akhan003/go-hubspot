# \RecordingSettingsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCrmV3ExtensionsCallingAppIdSettingsRecording**](RecordingSettingsAPI.md#GetCrmV3ExtensionsCallingAppIdSettingsRecording) | **Get** /crm/v3/extensions/calling/{appId}/settings/recording | 
[**PatchCrmV3ExtensionsCallingAppIdSettingsRecording**](RecordingSettingsAPI.md#PatchCrmV3ExtensionsCallingAppIdSettingsRecording) | **Patch** /crm/v3/extensions/calling/{appId}/settings/recording | 
[**PostCrmV3ExtensionsCallingAppIdSettingsRecording**](RecordingSettingsAPI.md#PostCrmV3ExtensionsCallingAppIdSettingsRecording) | **Post** /crm/v3/extensions/calling/{appId}/settings/recording | 



## GetCrmV3ExtensionsCallingAppIdSettingsRecording

> RecordingSettingsResponse GetCrmV3ExtensionsCallingAppIdSettingsRecording(ctx, appId).Execute()



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
	appId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordingSettingsAPI.GetCrmV3ExtensionsCallingAppIdSettingsRecording(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingSettingsAPI.GetCrmV3ExtensionsCallingAppIdSettingsRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ExtensionsCallingAppIdSettingsRecording`: RecordingSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordingSettingsAPI.GetCrmV3ExtensionsCallingAppIdSettingsRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsCallingAppIdSettingsRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecordingSettingsResponse**](RecordingSettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3ExtensionsCallingAppIdSettingsRecording

> RecordingSettingsResponse PatchCrmV3ExtensionsCallingAppIdSettingsRecording(ctx, appId).RecordingSettingsPatchRequest(recordingSettingsPatchRequest).Execute()



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
	appId := int32(56) // int32 | 
	recordingSettingsPatchRequest := *openapiclient.NewRecordingSettingsPatchRequest() // RecordingSettingsPatchRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordingSettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettingsRecording(context.Background(), appId).RecordingSettingsPatchRequest(recordingSettingsPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingSettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettingsRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCrmV3ExtensionsCallingAppIdSettingsRecording`: RecordingSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordingSettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettingsRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3ExtensionsCallingAppIdSettingsRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordingSettingsPatchRequest** | [**RecordingSettingsPatchRequest**](RecordingSettingsPatchRequest.md) |  | 

### Return type

[**RecordingSettingsResponse**](RecordingSettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsCallingAppIdSettingsRecording

> RecordingSettingsResponse PostCrmV3ExtensionsCallingAppIdSettingsRecording(ctx, appId).RecordingSettingsRequest(recordingSettingsRequest).Execute()



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
	appId := int32(56) // int32 | 
	recordingSettingsRequest := *openapiclient.NewRecordingSettingsRequest("UrlToRetrieveAuthedRecording_example") // RecordingSettingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RecordingSettingsAPI.PostCrmV3ExtensionsCallingAppIdSettingsRecording(context.Background(), appId).RecordingSettingsRequest(recordingSettingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordingSettingsAPI.PostCrmV3ExtensionsCallingAppIdSettingsRecording``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3ExtensionsCallingAppIdSettingsRecording`: RecordingSettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordingSettingsAPI.PostCrmV3ExtensionsCallingAppIdSettingsRecording`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsCallingAppIdSettingsRecordingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **recordingSettingsRequest** | [**RecordingSettingsRequest**](RecordingSettingsRequest.md) |  | 

### Return type

[**RecordingSettingsResponse**](RecordingSettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

