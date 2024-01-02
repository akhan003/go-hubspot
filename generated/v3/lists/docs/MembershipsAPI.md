# \MembershipsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAndRemoveMemberships**](MembershipsAPI.md#AddAndRemoveMemberships) | **Put** /crm/v3/lists/{listId}/memberships/add-and-remove | Add and/or Remove Records from a List
[**AddMemberships**](MembershipsAPI.md#AddMemberships) | **Put** /crm/v3/lists/{listId}/memberships/add | Add Records to a List
[**DeleteCrmV3ListsListIdMembershipsRemoveAll**](MembershipsAPI.md#DeleteCrmV3ListsListIdMembershipsRemoveAll) | **Delete** /crm/v3/lists/{listId}/memberships | Delete All Records from a List
[**GetCrmV3ListsListIdMembershipsGetPage**](MembershipsAPI.md#GetCrmV3ListsListIdMembershipsGetPage) | **Get** /crm/v3/lists/{listId}/memberships | Fetch List Memberships Ordered by ID
[**PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList**](MembershipsAPI.md#PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList) | **Put** /crm/v3/lists/{listId}/memberships/add-from/{sourceListId} | Add All Records from a Source List to a Destination List
[**RemoveMemberships**](MembershipsAPI.md#RemoveMemberships) | **Put** /crm/v3/lists/{listId}/memberships/remove | Remove Records from a List



## AddAndRemoveMemberships

> MembershipsUpdateResponse AddAndRemoveMemberships(ctx, listId).MembershipChangeRequest(membershipChangeRequest).Execute()

Add and/or Remove Records from a List



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
	listId := int32(56) // int32 | The **ILS ID** of the `MANUAL` or `SNAPSHOT` list.
	membershipChangeRequest := *openapiclient.NewMembershipChangeRequest([]int64{int64(123)}, []int64{int64(123)}) // MembershipChangeRequest | The IDs of the records to add and/or remove from the list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.AddAndRemoveMemberships(context.Background(), listId).MembershipChangeRequest(membershipChangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.AddAndRemoveMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAndRemoveMemberships`: MembershipsUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.AddAndRemoveMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the &#x60;MANUAL&#x60; or &#x60;SNAPSHOT&#x60; list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAndRemoveMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **membershipChangeRequest** | [**MembershipChangeRequest**](MembershipChangeRequest.md) | The IDs of the records to add and/or remove from the list. | 

### Return type

[**MembershipsUpdateResponse**](MembershipsUpdateResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddMemberships

> MembershipsUpdateResponse AddMemberships(ctx, listId).RequestBody(requestBody).Execute()

Add Records to a List



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
	listId := int32(56) // int32 | The **ILS ID** of the `MANUAL` or `SNAPSHOT` list.
	requestBody := []int64{int64(123)} // []int64 | The IDs of the records to add to the list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.AddMemberships(context.Background(), listId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.AddMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddMemberships`: MembershipsUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.AddMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the &#x60;MANUAL&#x60; or &#x60;SNAPSHOT&#x60; list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]int64** | The IDs of the records to add to the list. | 

### Return type

[**MembershipsUpdateResponse**](MembershipsUpdateResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCrmV3ListsListIdMembershipsRemoveAll

> DeleteCrmV3ListsListIdMembershipsRemoveAll(ctx, listId).Execute()

Delete All Records from a List



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
	listId := int32(56) // int32 | The **ILS ID** of the `MANUAL` or `SNAPSHOT` list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MembershipsAPI.DeleteCrmV3ListsListIdMembershipsRemoveAll(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.DeleteCrmV3ListsListIdMembershipsRemoveAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the &#x60;MANUAL&#x60; or &#x60;SNAPSHOT&#x60; list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3ListsListIdMembershipsRemoveAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3ListsListIdMembershipsGetPage

> CollectionResponseLong GetCrmV3ListsListIdMembershipsGetPage(ctx, listId).After(after).Before(before).Limit(limit).Execute()

Fetch List Memberships Ordered by ID



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
	listId := int32(56) // int32 | The **ILS ID** of the list.
	after := "after_example" // string | The paging offset token for the page that comes `after` the previously requested records.  If provided, then the records in the response will be the records following the offset, sorted in *ascending* order. Takes precedence over the `before` offset. (optional)
	before := "before_example" // string | The paging offset token for the page that comes `before` the previously requested records.  If provided, then the records in the response will be the records preceding the offset, sorted in *descending* order. (optional)
	limit := int32(56) // int32 | The number of records to return in the response. The maximum `limit` is 250. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.GetCrmV3ListsListIdMembershipsGetPage(context.Background(), listId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.GetCrmV3ListsListIdMembershipsGetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3ListsListIdMembershipsGetPage`: CollectionResponseLong
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.GetCrmV3ListsListIdMembershipsGetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3ListsListIdMembershipsGetPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The paging offset token for the page that comes &#x60;after&#x60; the previously requested records.  If provided, then the records in the response will be the records following the offset, sorted in *ascending* order. Takes precedence over the &#x60;before&#x60; offset. | 
 **before** | **string** | The paging offset token for the page that comes &#x60;before&#x60; the previously requested records.  If provided, then the records in the response will be the records preceding the offset, sorted in *descending* order. | 
 **limit** | **int32** | The number of records to return in the response. The maximum &#x60;limit&#x60; is 250. | [default to 100]

### Return type

[**CollectionResponseLong**](CollectionResponseLong.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList

> PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList(ctx, listId, sourceListId).Execute()

Add All Records from a Source List to a Destination List



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
	listId := int32(56) // int32 | The **ILS ID** of the `MANUAL` or `SNAPSHOT` *destination list*, which the *source list* records are added to.
	sourceListId := int32(56) // int32 | The **ILS ID** of the *source list* to grab the records from, which are then added to the *destination list*.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MembershipsAPI.PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList(context.Background(), listId, sourceListId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.PutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the &#x60;MANUAL&#x60; or &#x60;SNAPSHOT&#x60; *destination list*, which the *source list* records are added to. | 
**sourceListId** | **int32** | The **ILS ID** of the *source list* to grab the records from, which are then added to the *destination list*. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCrmV3ListsListIdMembershipsAddFromSourceListIdAddAllFromListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMemberships

> MembershipsUpdateResponse RemoveMemberships(ctx, listId).RequestBody(requestBody).Execute()

Remove Records from a List



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
	listId := int32(56) // int32 | The **ILS ID** of the `MANUAL` or `SNAPSHOT` list.
	requestBody := []int64{int64(123)} // []int64 | The IDs of the records to remove from the list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MembershipsAPI.RemoveMemberships(context.Background(), listId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.RemoveMemberships``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveMemberships`: MembershipsUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.RemoveMemberships`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the &#x60;MANUAL&#x60; or &#x60;SNAPSHOT&#x60; list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMembershipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]int64** | The IDs of the records to remove from the list. | 

### Return type

[**MembershipsUpdateResponse**](MembershipsUpdateResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

