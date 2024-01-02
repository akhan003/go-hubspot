# \CoreAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Archive**](CoreAPI.md#Archive) | **Delete** /crm/v3/properties/{objectType}/{propertyName} | Archive a property
[**Create**](CoreAPI.md#Create) | **Post** /crm/v3/properties/{objectType} | Create a property
[**GetAll**](CoreAPI.md#GetAll) | **Get** /crm/v3/properties/{objectType} | Read all properties
[**GetByName**](CoreAPI.md#GetByName) | **Get** /crm/v3/properties/{objectType}/{propertyName} | Read a property
[**Update**](CoreAPI.md#Update) | **Patch** /crm/v3/properties/{objectType}/{propertyName} | Update a property



## Archive

> Archive(ctx, objectType, propertyName).Execute()

Archive a property



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
	objectType := "objectType_example" // string | 
	propertyName := "propertyName_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.Archive(context.Background(), objectType, propertyName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.Archive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**propertyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> Property Create(ctx, objectType).PropertyCreate(propertyCreate).Execute()

Create a property



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
	objectType := "objectType_example" // string | 
	propertyCreate := *openapiclient.NewPropertyCreate("My Contact Property", "enumeration", "contactinformation", "Name_example", "select") // PropertyCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.Create(context.Background(), objectType).PropertyCreate(propertyCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: Property
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **propertyCreate** | [**PropertyCreate**](PropertyCreate.md) |  | 

### Return type

[**Property**](Property.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> CollectionResponsePropertyNoPaging GetAll(ctx, objectType).Archived(archived).Properties(properties).Execute()

Read all properties



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
	objectType := "objectType_example" // string | 
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	properties := "properties_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.GetAll(context.Background(), objectType).Archived(archived).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.GetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAll`: CollectionResponsePropertyNoPaging
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.GetAll`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **properties** | **string** |  | 

### Return type

[**CollectionResponsePropertyNoPaging**](CollectionResponsePropertyNoPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByName

> Property GetByName(ctx, objectType, propertyName).Archived(archived).Properties(properties).Execute()

Read a property



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
	objectType := "objectType_example" // string | 
	propertyName := "propertyName_example" // string | 
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)
	properties := "properties_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.GetByName(context.Background(), objectType, propertyName).Archived(archived).Properties(properties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.GetByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByName`: Property
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.GetByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**propertyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]
 **properties** | **string** |  | 

### Return type

[**Property**](Property.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Property Update(ctx, objectType, propertyName).PropertyUpdate(propertyUpdate).Execute()

Update a property



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
	objectType := "objectType_example" // string | 
	propertyName := "propertyName_example" // string | 
	propertyUpdate := *openapiclient.NewPropertyUpdate() // PropertyUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.Update(context.Background(), objectType, propertyName).PropertyUpdate(propertyUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: Property
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** |  | 
**propertyName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **propertyUpdate** | [**PropertyUpdate**](PropertyUpdate.md) |  | 

### Return type

[**Property**](Property.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

