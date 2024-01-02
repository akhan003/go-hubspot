# \BlogTagsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3BlogsTagsObjectId**](BlogTagsAPI.md#DeleteCmsV3BlogsTagsObjectId) | **Delete** /cms/v3/blogs/tags/{objectId} | Delete a Blog Tag
[**GetCmsV3BlogsTags**](BlogTagsAPI.md#GetCmsV3BlogsTags) | **Get** /cms/v3/blogs/tags | Get all Blog Tags
[**GetCmsV3BlogsTagsObjectId**](BlogTagsAPI.md#GetCmsV3BlogsTagsObjectId) | **Get** /cms/v3/blogs/tags/{objectId} | Retrieve a Blog Tag
[**PatchCmsV3BlogsTagsObjectId**](BlogTagsAPI.md#PatchCmsV3BlogsTagsObjectId) | **Patch** /cms/v3/blogs/tags/{objectId} | Update a Blog Tag
[**PostCmsV3BlogsTags**](BlogTagsAPI.md#PostCmsV3BlogsTags) | **Post** /cms/v3/blogs/tags | Create a new Blog Tag
[**PostCmsV3BlogsTagsBatchArchive**](BlogTagsAPI.md#PostCmsV3BlogsTagsBatchArchive) | **Post** /cms/v3/blogs/tags/batch/archive | Delete a batch of Blog Tags
[**PostCmsV3BlogsTagsBatchCreate**](BlogTagsAPI.md#PostCmsV3BlogsTagsBatchCreate) | **Post** /cms/v3/blogs/tags/batch/create | Create a batch of Blog Tags
[**PostCmsV3BlogsTagsBatchRead**](BlogTagsAPI.md#PostCmsV3BlogsTagsBatchRead) | **Post** /cms/v3/blogs/tags/batch/read | Retrieve a batch of Blog Tags
[**PostCmsV3BlogsTagsBatchUpdate**](BlogTagsAPI.md#PostCmsV3BlogsTagsBatchUpdate) | **Post** /cms/v3/blogs/tags/batch/update | Update a batch of Blog Tags
[**PostCmsV3BlogsTagsMultiLanguageAttachToLangGroup**](BlogTagsAPI.md#PostCmsV3BlogsTagsMultiLanguageAttachToLangGroup) | **Post** /cms/v3/blogs/tags/multi-language/attach-to-lang-group | Attach a Blog Tag to a multi-language group
[**PostCmsV3BlogsTagsMultiLanguageCreateLanguageVariation**](BlogTagsAPI.md#PostCmsV3BlogsTagsMultiLanguageCreateLanguageVariation) | **Post** /cms/v3/blogs/tags/multi-language/create-language-variation | Create a new language variation
[**PostCmsV3BlogsTagsMultiLanguageDetachFromLangGroup**](BlogTagsAPI.md#PostCmsV3BlogsTagsMultiLanguageDetachFromLangGroup) | **Post** /cms/v3/blogs/tags/multi-language/detach-from-lang-group | Detach a Blog Tag from a multi-language group
[**PostCmsV3BlogsTagsMultiLanguageUpdateLanguages**](BlogTagsAPI.md#PostCmsV3BlogsTagsMultiLanguageUpdateLanguages) | **Post** /cms/v3/blogs/tags/multi-language/update-languages | Update languages of multi-language group
[**PutCmsV3BlogsTagsMultiLanguageSetNewLangPrimary**](BlogTagsAPI.md#PutCmsV3BlogsTagsMultiLanguageSetNewLangPrimary) | **Put** /cms/v3/blogs/tags/multi-language/set-new-lang-primary | Set a new primary language



## DeleteCmsV3BlogsTagsObjectId

> DeleteCmsV3BlogsTagsObjectId(ctx, objectId).Archived(archived).Execute()

Delete a Blog Tag



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
	objectId := "objectId_example" // string | The Blog Tag id.
	archived := true // bool | Whether to return only results that have been archived. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogTagsAPI.DeleteCmsV3BlogsTagsObjectId(context.Background(), objectId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.DeleteCmsV3BlogsTagsObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Tag id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3BlogsTagsObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | 

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


## GetCmsV3BlogsTags

> CollectionResponseWithTotalTagForwardPaging GetCmsV3BlogsTags(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()

Get all Blog Tags



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createdAt := time.Now() // time.Time | Only return Blog Tags created at exactly the specified time. (optional)
	createdAfter := time.Now() // time.Time | Only return Blog Tags created after the specified time. (optional)
	createdBefore := time.Now() // time.Time | Only return Blog Tags created before the specified time. (optional)
	updatedAt := time.Now() // time.Time | Only return Blog Tags last updated at exactly the specified time. (optional)
	updatedAfter := time.Now() // time.Time | Only return Blog Tags last updated after the specified time. (optional)
	updatedBefore := time.Now() // time.Time | Only return Blog Tags last updated before the specified time. (optional)
	sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)
	archived := true // bool | Specifies whether to return deleted Blog Tags. Defaults to `false`. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogTagsAPI.GetCmsV3BlogsTags(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.GetCmsV3BlogsTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsTags`: CollectionResponseWithTotalTagForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BlogTagsAPI.GetCmsV3BlogsTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return Blog Tags created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return Blog Tags created after the specified time. | 
 **createdBefore** | **time.Time** | Only return Blog Tags created before the specified time. | 
 **updatedAt** | **time.Time** | Only return Blog Tags last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return Blog Tags last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return Blog Tags last updated before the specified time. | 
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Tags. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**CollectionResponseWithTotalTagForwardPaging**](CollectionResponseWithTotalTagForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsTagsObjectId

> Tag GetCmsV3BlogsTagsObjectId(ctx, objectId).Archived(archived).Property(property).Execute()

Retrieve a Blog Tag



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
	objectId := "objectId_example" // string | The Blog Tag id.
	archived := true // bool | Specifies whether to return deleted Blog Tags. Defaults to `false`. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogTagsAPI.GetCmsV3BlogsTagsObjectId(context.Background(), objectId).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.GetCmsV3BlogsTagsObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsTagsObjectId`: Tag
	fmt.Fprintf(os.Stdout, "Response from `BlogTagsAPI.GetCmsV3BlogsTagsObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Tag id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsTagsObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Specifies whether to return deleted Blog Tags. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**Tag**](Tag.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3BlogsTagsObjectId

> Tag PatchCmsV3BlogsTagsObjectId(ctx, objectId).Tag(tag).Archived(archived).Execute()

Update a Blog Tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	objectId := "objectId_example" // string | The Blog Tag id.
	tag := *openapiclient.NewTag(time.Now(), time.Now(), "Name_example", "Language_example", "Id_example", int64(123), time.Now()) // Tag | The JSON representation of the updated Blog Tag.
	archived := true // bool | Specifies whether to update deleted Blog Tags. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogTagsAPI.PatchCmsV3BlogsTagsObjectId(context.Background(), objectId).Tag(tag).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.PatchCmsV3BlogsTagsObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCmsV3BlogsTagsObjectId`: Tag
	fmt.Fprintf(os.Stdout, "Response from `BlogTagsAPI.PatchCmsV3BlogsTagsObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Tag id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3BlogsTagsObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | [**Tag**](Tag.md) | The JSON representation of the updated Blog Tag. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Tags. Defaults to &#x60;false&#x60;. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsTags

> Tag PostCmsV3BlogsTags(ctx).Tag(tag).Execute()

Create a new Blog Tag



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tag := *openapiclient.NewTag(time.Now(), time.Now(), "Name_example", "Language_example", "Id_example", int64(123), time.Now()) // Tag | The JSON representation of a new Blog Tag.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogTagsAPI.PostCmsV3BlogsTags(context.Background()).Tag(tag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.PostCmsV3BlogsTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsTags`: Tag
	fmt.Fprintf(os.Stdout, "Response from `BlogTagsAPI.PostCmsV3BlogsTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag** | [**Tag**](Tag.md) | The JSON representation of a new Blog Tag. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsTagsBatchArchive

> PostCmsV3BlogsTagsBatchArchive(ctx).BatchInputString(batchInputString).Execute()

Delete a batch of Blog Tags



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Tag ids.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogTagsAPI.PostCmsV3BlogsTagsBatchArchive(context.Background()).BatchInputString(batchInputString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.PostCmsV3BlogsTagsBatchArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Tag ids. | 

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


## PostCmsV3BlogsTagsBatchCreate

> BatchResponseTag PostCmsV3BlogsTagsBatchCreate(ctx).BatchInputTag(batchInputTag).Execute()

Create a batch of Blog Tags



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	batchInputTag := *openapiclient.NewBatchInputTag([]openapiclient.Tag{*openapiclient.NewTag(time.Now(), time.Now(), "Name_example", "Language_example", "Id_example", int64(123), time.Now())}) // BatchInputTag | The JSON array of new Blog Tags to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogTagsAPI.PostCmsV3BlogsTagsBatchCreate(context.Background()).BatchInputTag(batchInputTag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.PostCmsV3BlogsTagsBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsTagsBatchCreate`: BatchResponseTag
	fmt.Fprintf(os.Stdout, "Response from `BlogTagsAPI.PostCmsV3BlogsTagsBatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputTag** | [**BatchInputTag**](BatchInputTag.md) | The JSON array of new Blog Tags to create. | 

### Return type

[**BatchResponseTag**](BatchResponseTag.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsTagsBatchRead

> BatchResponseTag PostCmsV3BlogsTagsBatchRead(ctx).BatchInputString(batchInputString).Archived(archived).Execute()

Retrieve a batch of Blog Tags



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Tag ids.
	archived := true // bool | Specifies whether to return deleted Blog Tags. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogTagsAPI.PostCmsV3BlogsTagsBatchRead(context.Background()).BatchInputString(batchInputString).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.PostCmsV3BlogsTagsBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsTagsBatchRead`: BatchResponseTag
	fmt.Fprintf(os.Stdout, "Response from `BlogTagsAPI.PostCmsV3BlogsTagsBatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Tag ids. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Tags. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseTag**](BatchResponseTag.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsTagsBatchUpdate

> BatchResponseTag PostCmsV3BlogsTagsBatchUpdate(ctx).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()

Update a batch of Blog Tags



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
	batchInputJsonNode := *openapiclient.NewBatchInputJsonNode([]map[string]interface{}{map[string]interface{}(123)}) // BatchInputJsonNode | A JSON array of the JSON representations of the updated Blog Tags.
	archived := true // bool | Specifies whether to update deleted Blog Tags. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogTagsAPI.PostCmsV3BlogsTagsBatchUpdate(context.Background()).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.PostCmsV3BlogsTagsBatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsTagsBatchUpdate`: BatchResponseTag
	fmt.Fprintf(os.Stdout, "Response from `BlogTagsAPI.PostCmsV3BlogsTagsBatchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsBatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | A JSON array of the JSON representations of the updated Blog Tags. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Tags. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseTag**](BatchResponseTag.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsTagsMultiLanguageAttachToLangGroup

> PostCmsV3BlogsTagsMultiLanguageAttachToLangGroup(ctx).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()

Attach a Blog Tag to a multi-language group



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
	attachToLangPrimaryRequestVNext := *openapiclient.NewAttachToLangPrimaryRequestVNext("Language_example", "Id_example", "PrimaryId_example") // AttachToLangPrimaryRequestVNext | The JSON representation of the AttachToLangPrimaryRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogTagsAPI.PostCmsV3BlogsTagsMultiLanguageAttachToLangGroup(context.Background()).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.PostCmsV3BlogsTagsMultiLanguageAttachToLangGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsMultiLanguageAttachToLangGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachToLangPrimaryRequestVNext** | [**AttachToLangPrimaryRequestVNext**](AttachToLangPrimaryRequestVNext.md) | The JSON representation of the AttachToLangPrimaryRequest object. | 

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


## PostCmsV3BlogsTagsMultiLanguageCreateLanguageVariation

> Tag PostCmsV3BlogsTagsMultiLanguageCreateLanguageVariation(ctx).TagCloneRequestVNext(tagCloneRequestVNext).Execute()

Create a new language variation



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
	tagCloneRequestVNext := *openapiclient.NewTagCloneRequestVNext("Name_example", "Id_example") // TagCloneRequestVNext | The JSON representation of the ContentLanguageCloneRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogTagsAPI.PostCmsV3BlogsTagsMultiLanguageCreateLanguageVariation(context.Background()).TagCloneRequestVNext(tagCloneRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.PostCmsV3BlogsTagsMultiLanguageCreateLanguageVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsTagsMultiLanguageCreateLanguageVariation`: Tag
	fmt.Fprintf(os.Stdout, "Response from `BlogTagsAPI.PostCmsV3BlogsTagsMultiLanguageCreateLanguageVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsMultiLanguageCreateLanguageVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagCloneRequestVNext** | [**TagCloneRequestVNext**](TagCloneRequestVNext.md) | The JSON representation of the ContentLanguageCloneRequest object. | 

### Return type

[**Tag**](Tag.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsTagsMultiLanguageDetachFromLangGroup

> PostCmsV3BlogsTagsMultiLanguageDetachFromLangGroup(ctx).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()

Detach a Blog Tag from a multi-language group



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
	detachFromLangGroupRequestVNext := *openapiclient.NewDetachFromLangGroupRequestVNext("Id_example") // DetachFromLangGroupRequestVNext | The JSON representation of the DetachFromLangGroupRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogTagsAPI.PostCmsV3BlogsTagsMultiLanguageDetachFromLangGroup(context.Background()).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.PostCmsV3BlogsTagsMultiLanguageDetachFromLangGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsMultiLanguageDetachFromLangGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **detachFromLangGroupRequestVNext** | [**DetachFromLangGroupRequestVNext**](DetachFromLangGroupRequestVNext.md) | The JSON representation of the DetachFromLangGroupRequest object. | 

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


## PostCmsV3BlogsTagsMultiLanguageUpdateLanguages

> PostCmsV3BlogsTagsMultiLanguageUpdateLanguages(ctx).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()

Update languages of multi-language group



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
	updateLanguagesRequestVNext := *openapiclient.NewUpdateLanguagesRequestVNext(map[string]string{"key": "Inner_example"}, "PrimaryId_example") // UpdateLanguagesRequestVNext | The JSON representation of the UpdateLanguagesRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogTagsAPI.PostCmsV3BlogsTagsMultiLanguageUpdateLanguages(context.Background()).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.PostCmsV3BlogsTagsMultiLanguageUpdateLanguages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsTagsMultiLanguageUpdateLanguagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLanguagesRequestVNext** | [**UpdateLanguagesRequestVNext**](UpdateLanguagesRequestVNext.md) | The JSON representation of the UpdateLanguagesRequest object. | 

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


## PutCmsV3BlogsTagsMultiLanguageSetNewLangPrimary

> PutCmsV3BlogsTagsMultiLanguageSetNewLangPrimary(ctx).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()

Set a new primary language



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
	setNewLanguagePrimaryRequestVNext := *openapiclient.NewSetNewLanguagePrimaryRequestVNext("Id_example") // SetNewLanguagePrimaryRequestVNext | The JSON representation of the SetNewLanguagePrimaryRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogTagsAPI.PutCmsV3BlogsTagsMultiLanguageSetNewLangPrimary(context.Background()).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogTagsAPI.PutCmsV3BlogsTagsMultiLanguageSetNewLangPrimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCmsV3BlogsTagsMultiLanguageSetNewLangPrimaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setNewLanguagePrimaryRequestVNext** | [**SetNewLanguagePrimaryRequestVNext**](SetNewLanguagePrimaryRequestVNext.md) | The JSON representation of the SetNewLanguagePrimaryRequest object. | 

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

