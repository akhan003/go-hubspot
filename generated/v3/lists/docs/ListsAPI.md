# \ListsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ListsAPI.md#Create) | **Post** /crm/v3/lists/ | Create List
[**DoSearch**](ListsAPI.md#DoSearch) | **Post** /crm/v3/lists/search | Search Lists
[**GetAll**](ListsAPI.md#GetAll) | **Get** /crm/v3/lists/ | Fetch Multiple Lists
[**GetById**](ListsAPI.md#GetById) | **Get** /crm/v3/lists/{listId} | Fetch List by ID
[**GetByName**](ListsAPI.md#GetByName) | **Get** /crm/v3/lists/object-type-id/{objectTypeId}/name/{listName} | Fetch List by Name
[**Remove**](ListsAPI.md#Remove) | **Delete** /crm/v3/lists/{listId} | Delete a List
[**Restore**](ListsAPI.md#Restore) | **Put** /crm/v3/lists/{listId}/restore | Restore a List
[**UpdateListFilters**](ListsAPI.md#UpdateListFilters) | **Put** /crm/v3/lists/{listId}/update-list-filters | Update List Filter Definition
[**UpdateName**](ListsAPI.md#UpdateName) | **Put** /crm/v3/lists/{listId}/update-list-name | Update List Name



## Create

> ListCreateResponse Create(ctx).ListCreateRequest(listCreateRequest).Execute()

Create List



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
	listCreateRequest := *openapiclient.NewListCreateRequest("ObjectTypeId_example", "ProcessingType_example", "Name_example") // ListCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.Create(context.Background()).ListCreateRequest(listCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: ListCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listCreateRequest** | [**ListCreateRequest**](ListCreateRequest.md) |  | 

### Return type

[**ListCreateResponse**](ListCreateResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DoSearch

> ListSearchResponse DoSearch(ctx).ListSearchRequest(listSearchRequest).Execute()

Search Lists



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
	listSearchRequest := *openapiclient.NewListSearchRequest(int32(123), []string{"AdditionalPropertiesField_example"}) // ListSearchRequest | The IDs of the records to add and/or remove from the list.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.DoSearch(context.Background()).ListSearchRequest(listSearchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.DoSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoSearch`: ListSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.DoSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDoSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listSearchRequest** | [**ListSearchRequest**](ListSearchRequest.md) | The IDs of the records to add and/or remove from the list. | 

### Return type

[**ListSearchResponse**](ListSearchResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> ListsByIdResponse GetAll(ctx).ListIds(listIds).IncludeFilters(includeFilters).Execute()

Fetch Multiple Lists



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
	listIds := []int32{int32(123)} // []int32 | The **ILS IDs** of the lists to fetch. (optional)
	includeFilters := true // bool | A flag indicating whether or not the response object list definitions should include a filter branch definition. By default, object list definitions will not have their filter branch definitions included in the response. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetAll(context.Background()).ListIds(listIds).IncludeFilters(includeFilters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAll`: ListsByIdResponse
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listIds** | **[]int32** | The **ILS IDs** of the lists to fetch. | 
 **includeFilters** | **bool** | A flag indicating whether or not the response object list definitions should include a filter branch definition. By default, object list definitions will not have their filter branch definitions included in the response. | [default to false]

### Return type

[**ListsByIdResponse**](ListsByIdResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetById

> ListFetchResponse GetById(ctx, listId).IncludeFilters(includeFilters).Execute()

Fetch List by ID



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
	listId := int32(56) // int32 | The **ILS ID** of the list to fetch.
	includeFilters := true // bool | A flag indicating whether or not the response object list definition should include a filter branch definition. By default, object list definitions will not have their filter branch definitions included in the response. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetById(context.Background(), listId).IncludeFilters(includeFilters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetById`: ListFetchResponse
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the list to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeFilters** | **bool** | A flag indicating whether or not the response object list definition should include a filter branch definition. By default, object list definitions will not have their filter branch definitions included in the response. | [default to false]

### Return type

[**ListFetchResponse**](ListFetchResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByName

> ListFetchResponse GetByName(ctx, listName, objectTypeId).IncludeFilters(includeFilters).Execute()

Fetch List by Name



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
	listName := "listName_example" // string | The name of the list to fetch. This is **not** case sensitive.
	objectTypeId := "objectTypeId_example" // string | The object type ID of the object types stored by the list to fetch. For example, `0-1` for a `CONTACT` list.
	includeFilters := true // bool | A flag indicating whether or not the response object list definition should include a filter branch definition. By default, object list definitions will not have their filter branch definitions included in the response. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.GetByName(context.Background(), listName, objectTypeId).IncludeFilters(includeFilters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.GetByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByName`: ListFetchResponse
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.GetByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listName** | **string** | The name of the list to fetch. This is **not** case sensitive. | 
**objectTypeId** | **string** | The object type ID of the object types stored by the list to fetch. For example, &#x60;0-1&#x60; for a &#x60;CONTACT&#x60; list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeFilters** | **bool** | A flag indicating whether or not the response object list definition should include a filter branch definition. By default, object list definitions will not have their filter branch definitions included in the response. | [default to false]

### Return type

[**ListFetchResponse**](ListFetchResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Remove

> Remove(ctx, listId).Execute()

Delete a List



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
	listId := int32(56) // int32 | The **ILS ID** of the list to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.Remove(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.Remove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the list to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRequest struct via the builder pattern


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


## Restore

> Restore(ctx, listId).Execute()

Restore a List



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
	listId := int32(56) // int32 | The **ILS ID** of the list to restore.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ListsAPI.Restore(context.Background(), listId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.Restore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the list to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreRequest struct via the builder pattern


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


## UpdateListFilters

> ListUpdateResponse UpdateListFilters(ctx, listId).ListFilterUpdateRequest(listFilterUpdateRequest).EnrollObjectsInWorkflows(enrollObjectsInWorkflows).Execute()

Update List Filter Definition



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
	listId := int32(56) // int32 | The **ILS ID** of the list to update.
	listFilterUpdateRequest := *openapiclient.NewListFilterUpdateRequest(openapiclient.PublicPropertyAssociationFilterBranch_filterBranches_inner{PublicAndFilterBranch: openapiclient.NewPublicAndFilterBranch("FilterBranchType_example", []openapiclient.PublicPropertyAssociationFilterBranchFilterBranchesInner{openapiclient.PublicPropertyAssociationFilterBranch_filterBranches_inner{PublicAndFilterBranch: openapiclient.NewPublicAndFilterBranch("FilterBranchType_example", []openapiclient.PublicPropertyAssociationFilterBranchFilterBranchesInner{openapiclient.PublicPropertyAssociationFilterBranch_filterBranches_inner{PublicAndFilterBranch: }}, []openapiclient.PublicPropertyAssociationFilterBranchFiltersInner{openapiclient.PublicPropertyAssociationFilterBranch_filters_inner{PublicAdsSearchFilter: openapiclient.NewPublicAdsSearchFilter("FilterType_example", "EntityType_example", "SearchTermType_example", []string{"SearchTerms_example"}, "AdNetwork_example", "Operator_example")}}, "FilterBranchOperator_example")}}, []openapiclient.PublicPropertyAssociationFilterBranchFiltersInner{openapiclient.PublicPropertyAssociationFilterBranch_filters_inner{PublicAdsSearchFilter: openapiclient.NewPublicAdsSearchFilter("FilterType_example", "EntityType_example", "SearchTermType_example", []string{"SearchTerms_example"}, "AdNetwork_example", "Operator_example")}}, "FilterBranchOperator_example")}) // ListFilterUpdateRequest | 
	enrollObjectsInWorkflows := true // bool | A flag indicating whether or not the memberships added to the list as a result of the filter change should be enrolled in workflows that are relevant to this list. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.UpdateListFilters(context.Background(), listId).ListFilterUpdateRequest(listFilterUpdateRequest).EnrollObjectsInWorkflows(enrollObjectsInWorkflows).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.UpdateListFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateListFilters`: ListUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.UpdateListFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the list to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listFilterUpdateRequest** | [**ListFilterUpdateRequest**](ListFilterUpdateRequest.md) |  | 
 **enrollObjectsInWorkflows** | **bool** | A flag indicating whether or not the memberships added to the list as a result of the filter change should be enrolled in workflows that are relevant to this list. | [default to false]

### Return type

[**ListUpdateResponse**](ListUpdateResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateName

> ListUpdateResponse UpdateName(ctx, listId).ListName(listName).IncludeFilters(includeFilters).Execute()

Update List Name



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
	listId := int32(56) // int32 | The **ILS ID** of the list to update.
	listName := "listName_example" // string | The name to update the list to. (optional)
	includeFilters := true // bool | A flag indicating whether or not the response object list definition should include a filter branch definition. By default, object list definitions will not have their filter branch definitions included in the response. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ListsAPI.UpdateName(context.Background(), listId).ListName(listName).IncludeFilters(includeFilters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ListsAPI.UpdateName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateName`: ListUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `ListsAPI.UpdateName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listId** | **int32** | The **ILS ID** of the list to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listName** | **string** | The name to update the list to. | 
 **includeFilters** | **bool** | A flag indicating whether or not the response object list definition should include a filter branch definition. By default, object list definitions will not have their filter branch definitions included in the response. | [default to false]

### Return type

[**ListUpdateResponse**](ListUpdateResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [private_apps](../README.md#private_apps)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

