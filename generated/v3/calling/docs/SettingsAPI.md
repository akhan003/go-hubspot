# \SettingsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3ExtensionsCallingAppIdSettings**](SettingsAPI.md#DeleteCrmV3ExtensionsCallingAppIdSettings) | **Delete** /crm/v3/extensions/calling/{appId}/settings | Delete calling settings
[**GetCrmV3ExtensionsCallingAppIdSettings**](SettingsAPI.md#GetCrmV3ExtensionsCallingAppIdSettings) | **Get** /crm/v3/extensions/calling/{appId}/settings | Get calling settings
[**PatchCrmV3ExtensionsCallingAppIdSettings**](SettingsAPI.md#PatchCrmV3ExtensionsCallingAppIdSettings) | **Patch** /crm/v3/extensions/calling/{appId}/settings | Update settings
[**PostCrmV3ExtensionsCallingAppIdSettings**](SettingsAPI.md#PostCrmV3ExtensionsCallingAppIdSettings) | **Post** /crm/v3/extensions/calling/{appId}/settings | Configure a calling extension



## DeleteCrmV3ExtensionsCallingAppIdSettings

> DeleteCrmV3ExtensionsCallingAppIdSettings(ctx, appId).Execute()

Delete calling settings



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
	appId := int32(56) // int32 | The ID of the target app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SettingsAPI.DeleteCrmV3ExtensionsCallingAppIdSettings(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.DeleteCrmV3ExtensionsCallingAppIdSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ExtensionsCallingAppIdSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ExtensionsCallingAppIdSettings

> SettingsResponse GetCrmV3ExtensionsCallingAppIdSettings(ctx, appId).Execute()

Get calling settings



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
	appId := int32(56) // int32 | The ID of the target app.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.GetCrmV3ExtensionsCallingAppIdSettings(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.GetCrmV3ExtensionsCallingAppIdSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ExtensionsCallingAppIdSettings`: SettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.GetCrmV3ExtensionsCallingAppIdSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ExtensionsCallingAppIdSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3ExtensionsCallingAppIdSettings

> SettingsResponse PatchCrmV3ExtensionsCallingAppIdSettings(ctx, appId).SettingsPatchRequest(settingsPatchRequest).Execute()

Update settings



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
	appId := int32(56) // int32 | The ID of the target app.
	settingsPatchRequest := *openapiclient.NewSettingsPatchRequest() // SettingsPatchRequest | Updated details for the settings.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettings(context.Background(), appId).SettingsPatchRequest(settingsPatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCrmV3ExtensionsCallingAppIdSettings`: SettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.PatchCrmV3ExtensionsCallingAppIdSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3ExtensionsCallingAppIdSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settingsPatchRequest** | [**SettingsPatchRequest**](SettingsPatchRequest.md) | Updated details for the settings. | 

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3ExtensionsCallingAppIdSettings

> SettingsResponse PostCrmV3ExtensionsCallingAppIdSettings(ctx, appId).SettingsRequest(settingsRequest).Execute()

Configure a calling extension



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
	appId := int32(56) // int32 | The ID of the target app.
	settingsRequest := *openapiclient.NewSettingsRequest("HubPhone", "https://www.example.com/hubspot/iframe") // SettingsRequest | Settings state to create with.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SettingsAPI.PostCrmV3ExtensionsCallingAppIdSettings(context.Background(), appId).SettingsRequest(settingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SettingsAPI.PostCrmV3ExtensionsCallingAppIdSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3ExtensionsCallingAppIdSettings`: SettingsResponse
	fmt.Fprintf(os.Stdout, "Response from `SettingsAPI.PostCrmV3ExtensionsCallingAppIdSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **int32** | The ID of the target app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3ExtensionsCallingAppIdSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settingsRequest** | [**SettingsRequest**](SettingsRequest.md) | Settings state to create with. | 

### Return type

[**SettingsResponse**](SettingsResponse.md)

### Authorization

[developer_hapikey](../README.md#developer_hapikey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

