# \MembershipsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](MembershipsAPI.md#Add) | **Put** /crm/v3/lists/{listId}/memberships/add | Add Records to a List
[**AddAllFromList**](MembershipsAPI.md#AddAllFromList) | **Put** /crm/v3/lists/{listId}/memberships/add-from/{sourceListId} | Add All Records from a Source List to a Destination List
[**AddAndRemove**](MembershipsAPI.md#AddAndRemove) | **Put** /crm/v3/lists/{listId}/memberships/add-and-remove | Add and/or Remove Records from a List
[**GetPage**](MembershipsAPI.md#GetPage) | **Get** /crm/v3/lists/{listId}/memberships | Fetch List Memberships Ordered by ID
[**Remove**](MembershipsAPI.md#Remove) | **Put** /crm/v3/lists/{listId}/memberships/remove | Remove Records from a List
[**RemoveAll**](MembershipsAPI.md#RemoveAll) | **Delete** /crm/v3/lists/{listId}/memberships | Delete All Records from a List



## Add

> MembershipsUpdateResponse Add(ctx, listId).RequestBody(requestBody).Execute()

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
	resp, r, err := apiClient.MembershipsAPI.Add(context.Background(), listId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.Add``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Add`: MembershipsUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.Add`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the &#x60;MANUAL&#x60; or &#x60;SNAPSHOT&#x60; list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRequest struct via the builder pattern


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


## AddAllFromList

> AddAllFromList(ctx, listId, sourceListId).Execute()

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
	r, err := apiClient.MembershipsAPI.AddAllFromList(context.Background(), listId, sourceListId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.AddAllFromList``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAddAllFromListRequest struct via the builder pattern


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


## AddAndRemove

> MembershipsUpdateResponse AddAndRemove(ctx, listId).MembershipChangeRequest(membershipChangeRequest).Execute()

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
	resp, r, err := apiClient.MembershipsAPI.AddAndRemove(context.Background(), listId).MembershipChangeRequest(membershipChangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.AddAndRemove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAndRemove`: MembershipsUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.AddAndRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the &#x60;MANUAL&#x60; or &#x60;SNAPSHOT&#x60; list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAndRemoveRequest struct via the builder pattern


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


## GetPage

> CollectionResponseLong GetPage(ctx, listId).After(after).Before(before).Limit(limit).Execute()

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
	resp, r, err := apiClient.MembershipsAPI.GetPage(context.Background(), listId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.GetPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPage`: CollectionResponseLong
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.GetPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPageRequest struct via the builder pattern


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


## Remove

> MembershipsUpdateResponse Remove(ctx, listId).RequestBody(requestBody).Execute()

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
	resp, r, err := apiClient.MembershipsAPI.Remove(context.Background(), listId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.Remove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Remove`: MembershipsUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `MembershipsAPI.Remove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the &#x60;MANUAL&#x60; or &#x60;SNAPSHOT&#x60; list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRequest struct via the builder pattern


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


## RemoveAll

> RemoveAll(ctx, listId).Execute()

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
	r, err := apiClient.MembershipsAPI.RemoveAll(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembershipsAPI.RemoveAll``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveAllRequest struct via the builder pattern


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

