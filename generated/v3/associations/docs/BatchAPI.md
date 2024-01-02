# \BatchAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchArchive**](BatchAPI.md#BatchArchive) | **Post** /crm/v3/associations/{fromObjectType}/{toObjectType}/batch/archive | Archive a batch of associations
[**BatchCreate**](BatchAPI.md#BatchCreate) | **Post** /crm/v3/associations/{fromObjectType}/{toObjectType}/batch/create | Create a batch of associations
[**BatchRead**](BatchAPI.md#BatchRead) | **Post** /crm/v3/associations/{fromObjectType}/{toObjectType}/batch/read | Read a batch of associations



## BatchArchive

> BatchArchive(ctx, fromObjectType, toObjectType).BatchInputPublicAssociation(batchInputPublicAssociation).Execute()

Archive a batch of associations



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
	fromObjectType := "fromObjectType_example" // string | 
	toObjectType := "toObjectType_example" // string | 
	batchInputPublicAssociation := *openapiclient.NewBatchInputPublicAssociation([]openapiclient.PublicAssociation{*openapiclient.NewPublicAssociation(*openapiclient.NewPublicObjectId("Id_example"), *openapiclient.NewPublicObjectId("Id_example"), "Type_example")}) // BatchInputPublicAssociation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BatchAPI.BatchArchive(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociation(batchInputPublicAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociation** | [**BatchInputPublicAssociation**](BatchInputPublicAssociation.md) |  | 

### Return type

 (empty response body)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchCreate

> BatchResponsePublicAssociation BatchCreate(ctx, fromObjectType, toObjectType).BatchInputPublicAssociation(batchInputPublicAssociation).Execute()

Create a batch of associations



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
	fromObjectType := "fromObjectType_example" // string | 
	toObjectType := "toObjectType_example" // string | 
	batchInputPublicAssociation := *openapiclient.NewBatchInputPublicAssociation([]openapiclient.PublicAssociation{*openapiclient.NewPublicAssociation(*openapiclient.NewPublicObjectId("Id_example"), *openapiclient.NewPublicObjectId("Id_example"), "Type_example")}) // BatchInputPublicAssociation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchCreate(context.Background(), fromObjectType, toObjectType).BatchInputPublicAssociation(batchInputPublicAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchCreate`: BatchResponsePublicAssociation
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicAssociation** | [**BatchInputPublicAssociation**](BatchInputPublicAssociation.md) |  | 

### Return type

[**BatchResponsePublicAssociation**](BatchResponsePublicAssociation.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BatchRead

> BatchResponsePublicAssociationMulti BatchRead(ctx, fromObjectType, toObjectType).BatchInputPublicObjectId(batchInputPublicObjectId).Execute()

Read a batch of associations



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
	fromObjectType := "fromObjectType_example" // string | 
	toObjectType := "toObjectType_example" // string | 
	batchInputPublicObjectId := *openapiclient.NewBatchInputPublicObjectId([]openapiclient.PublicObjectId{*openapiclient.NewPublicObjectId("Id_example")}) // BatchInputPublicObjectId | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BatchRead(context.Background(), fromObjectType, toObjectType).BatchInputPublicObjectId(batchInputPublicObjectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BatchRead`: BatchResponsePublicAssociationMulti
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BatchRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fromObjectType** | **string** |  | 
**toObjectType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **batchInputPublicObjectId** | [**BatchInputPublicObjectId**](BatchInputPublicObjectId.md) |  | 

### Return type

[**BatchResponsePublicAssociationMulti**](BatchResponsePublicAssociationMulti.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [oauth2](../README.md#oauth2), [private_apps_legacy](../README.md#private_apps_legacy), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

