# \BlogAuthorsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3BlogsAuthorsObjectId**](BlogAuthorsAPI.md#DeleteCmsV3BlogsAuthorsObjectId) | **Delete** /cms/v3/blogs/authors/{objectId} | Delete a Blog Author
[**GetCmsV3BlogsAuthors**](BlogAuthorsAPI.md#GetCmsV3BlogsAuthors) | **Get** /cms/v3/blogs/authors | Get all Blog Authors
[**GetCmsV3BlogsAuthorsObjectId**](BlogAuthorsAPI.md#GetCmsV3BlogsAuthorsObjectId) | **Get** /cms/v3/blogs/authors/{objectId} | Retrieve a Blog Author
[**PatchCmsV3BlogsAuthorsObjectId**](BlogAuthorsAPI.md#PatchCmsV3BlogsAuthorsObjectId) | **Patch** /cms/v3/blogs/authors/{objectId} | Update a Blog Author
[**PostCmsV3BlogsAuthors**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthors) | **Post** /cms/v3/blogs/authors | Create a new Blog Author
[**PostCmsV3BlogsAuthorsBatchArchive**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsBatchArchive) | **Post** /cms/v3/blogs/authors/batch/archive | Delete a batch of Blog Authors
[**PostCmsV3BlogsAuthorsBatchCreate**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsBatchCreate) | **Post** /cms/v3/blogs/authors/batch/create | Create a batch of Blog Authors
[**PostCmsV3BlogsAuthorsBatchRead**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsBatchRead) | **Post** /cms/v3/blogs/authors/batch/read | Retrieve a batch of Blog Authors
[**PostCmsV3BlogsAuthorsBatchUpdate**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsBatchUpdate) | **Post** /cms/v3/blogs/authors/batch/update | Update a batch of Blog Authors
[**PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroup**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroup) | **Post** /cms/v3/blogs/authors/multi-language/attach-to-lang-group | Attach a Blog Author to a multi-language group
[**PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariation**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariation) | **Post** /cms/v3/blogs/authors/multi-language/create-language-variation | Create a new language variation
[**PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroup**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroup) | **Post** /cms/v3/blogs/authors/multi-language/detach-from-lang-group | Detach a Blog Author from a multi-language group
[**PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguages**](BlogAuthorsAPI.md#PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguages) | **Post** /cms/v3/blogs/authors/multi-language/update-languages | Update languages of multi-language group
[**PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimary**](BlogAuthorsAPI.md#PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimary) | **Put** /cms/v3/blogs/authors/multi-language/set-new-lang-primary | Set a new primary language



## DeleteCmsV3BlogsAuthorsObjectId

> DeleteCmsV3BlogsAuthorsObjectId(ctx, objectId).Archived(archived).Execute()

Delete a Blog Author



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
	objectId := "objectId_example" // string | The Blog Author id.
	archived := true // bool | Whether to return only results that have been archived. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogAuthorsAPI.DeleteCmsV3BlogsAuthorsObjectId(context.Background(), objectId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.DeleteCmsV3BlogsAuthorsObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Author id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3BlogsAuthorsObjectIdRequest struct via the builder pattern


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


## GetCmsV3BlogsAuthors

> CollectionResponseWithTotalBlogAuthorForwardPaging GetCmsV3BlogsAuthors(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()

Get all Blog Authors



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
	createdAt := time.Now() // time.Time | Only return Blog Authors created at exactly the specified time. (optional)
	createdAfter := time.Now() // time.Time | Only return Blog Authors created after the specified time. (optional)
	createdBefore := time.Now() // time.Time | Only return Blog Authors created before the specified time. (optional)
	updatedAt := time.Now() // time.Time | Only return Blog Authors last updated at exactly the specified time. (optional)
	updatedAfter := time.Now() // time.Time | Only return Blog Authors last updated after the specified time. (optional)
	updatedBefore := time.Now() // time.Time | Only return Blog Authors last updated before the specified time. (optional)
	sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)
	archived := true // bool | Specifies whether to return deleted Blog Authors. Defaults to `false`. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.GetCmsV3BlogsAuthors(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.GetCmsV3BlogsAuthors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsAuthors`: CollectionResponseWithTotalBlogAuthorForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.GetCmsV3BlogsAuthors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsAuthorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return Blog Authors created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return Blog Authors created after the specified time. | 
 **createdBefore** | **time.Time** | Only return Blog Authors created before the specified time. | 
 **updatedAt** | **time.Time** | Only return Blog Authors last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return Blog Authors last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return Blog Authors last updated before the specified time. | 
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Authors. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**CollectionResponseWithTotalBlogAuthorForwardPaging**](CollectionResponseWithTotalBlogAuthorForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsAuthorsObjectId

> BlogAuthor GetCmsV3BlogsAuthorsObjectId(ctx, objectId).Archived(archived).Property(property).Execute()

Retrieve a Blog Author



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
	objectId := "objectId_example" // string | The Blog Author id.
	archived := true // bool | Specifies whether to return deleted Blog Authors. Defaults to `false`. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.GetCmsV3BlogsAuthorsObjectId(context.Background(), objectId).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.GetCmsV3BlogsAuthorsObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsAuthorsObjectId`: BlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.GetCmsV3BlogsAuthorsObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Author id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsAuthorsObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Specifies whether to return deleted Blog Authors. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3BlogsAuthorsObjectId

> BlogAuthor PatchCmsV3BlogsAuthorsObjectId(ctx, objectId).BlogAuthor(blogAuthor).Archived(archived).Execute()

Update a Blog Author



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
	objectId := "objectId_example" // string | The Blog Author id.
	blogAuthor := *openapiclient.NewBlogAuthor("Website_example", "DisplayName_example", time.Now(), "Facebook_example", "FullName_example", "Bio_example", "Language_example", "Linkedin_example", "Avatar_example", int64(123), "Twitter_example", time.Now(), "Name_example", "Id_example", time.Now(), "Email_example", "Slug_example") // BlogAuthor | The JSON representation of the updated Blog Author.
	archived := true // bool | Specifies whether to update deleted Blog Authors. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.PatchCmsV3BlogsAuthorsObjectId(context.Background(), objectId).BlogAuthor(blogAuthor).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PatchCmsV3BlogsAuthorsObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCmsV3BlogsAuthorsObjectId`: BlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.PatchCmsV3BlogsAuthorsObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Author id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3BlogsAuthorsObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogAuthor** | [**BlogAuthor**](BlogAuthor.md) | The JSON representation of the updated Blog Author. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthors

> BlogAuthor PostCmsV3BlogsAuthors(ctx).BlogAuthor(blogAuthor).Execute()

Create a new Blog Author



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
	blogAuthor := *openapiclient.NewBlogAuthor("Website_example", "DisplayName_example", time.Now(), "Facebook_example", "FullName_example", "Bio_example", "Language_example", "Linkedin_example", "Avatar_example", int64(123), "Twitter_example", time.Now(), "Name_example", "Id_example", time.Now(), "Email_example", "Slug_example") // BlogAuthor | The JSON representation of a new Blog Author.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthors(context.Background()).BlogAuthor(blogAuthor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsAuthors`: BlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.PostCmsV3BlogsAuthors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogAuthor** | [**BlogAuthor**](BlogAuthor.md) | The JSON representation of a new Blog Author. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsBatchArchive

> PostCmsV3BlogsAuthorsBatchArchive(ctx).BatchInputString(batchInputString).Execute()

Delete a batch of Blog Authors



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Author ids.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchArchive(context.Background()).BatchInputString(batchInputString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Author ids. | 

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


## PostCmsV3BlogsAuthorsBatchCreate

> BatchResponseBlogAuthor PostCmsV3BlogsAuthorsBatchCreate(ctx).BatchInputBlogAuthor(batchInputBlogAuthor).Execute()

Create a batch of Blog Authors



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
	batchInputBlogAuthor := *openapiclient.NewBatchInputBlogAuthor([]openapiclient.BlogAuthor{*openapiclient.NewBlogAuthor("Website_example", "DisplayName_example", time.Now(), "Facebook_example", "FullName_example", "Bio_example", "Language_example", "Linkedin_example", "Avatar_example", int64(123), "Twitter_example", time.Now(), "Name_example", "Id_example", time.Now(), "Email_example", "Slug_example")}) // BatchInputBlogAuthor | The JSON array of new Blog Authors to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchCreate(context.Background()).BatchInputBlogAuthor(batchInputBlogAuthor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsAuthorsBatchCreate`: BatchResponseBlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputBlogAuthor** | [**BatchInputBlogAuthor**](BatchInputBlogAuthor.md) | The JSON array of new Blog Authors to create. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsBatchRead

> BatchResponseBlogAuthor PostCmsV3BlogsAuthorsBatchRead(ctx).BatchInputString(batchInputString).Archived(archived).Execute()

Retrieve a batch of Blog Authors



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Author ids.
	archived := true // bool | Specifies whether to return deleted Blog Authors. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchRead(context.Background()).BatchInputString(batchInputString).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsAuthorsBatchRead`: BatchResponseBlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Author ids. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsBatchUpdate

> BatchResponseBlogAuthor PostCmsV3BlogsAuthorsBatchUpdate(ctx).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()

Update a batch of Blog Authors



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
	batchInputJsonNode := *openapiclient.NewBatchInputJsonNode([]map[string]interface{}{map[string]interface{}(123)}) // BatchInputJsonNode | A JSON array of the JSON representations of the updated Blog Authors.
	archived := true // bool | Specifies whether to update deleted Blog Authors. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchUpdate(context.Background()).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsAuthorsBatchUpdate`: BatchResponseBlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.PostCmsV3BlogsAuthorsBatchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsBatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | A JSON array of the JSON representations of the updated Blog Authors. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Authors. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogAuthor**](BatchResponseBlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroup

> PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroup(ctx).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()

Attach a Blog Author to a multi-language group



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
	r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroup(context.Background()).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsMultiLanguageAttachToLangGroupRequest struct via the builder pattern


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


## PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariation

> BlogAuthor PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariation(ctx).BlogAuthorCloneRequestVNext(blogAuthorCloneRequestVNext).Execute()

Create a new language variation



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
	blogAuthorCloneRequestVNext := *openapiclient.NewBlogAuthorCloneRequestVNext("Id_example", *openapiclient.NewBlogAuthor("Website_example", "DisplayName_example", time.Now(), "Facebook_example", "FullName_example", "Bio_example", "Language_example", "Linkedin_example", "Avatar_example", int64(123), "Twitter_example", time.Now(), "Name_example", "Id_example", time.Now(), "Email_example", "Slug_example")) // BlogAuthorCloneRequestVNext | The JSON representation of the ContentLanguageCloneRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariation(context.Background()).BlogAuthorCloneRequestVNext(blogAuthorCloneRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariation`: BlogAuthor
	fmt.Fprintf(os.Stdout, "Response from `BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsMultiLanguageCreateLanguageVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogAuthorCloneRequestVNext** | [**BlogAuthorCloneRequestVNext**](BlogAuthorCloneRequestVNext.md) | The JSON representation of the ContentLanguageCloneRequest object. | 

### Return type

[**BlogAuthor**](BlogAuthor.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroup

> PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroup(ctx).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()

Detach a Blog Author from a multi-language group



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
	r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroup(context.Background()).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsMultiLanguageDetachFromLangGroupRequest struct via the builder pattern


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


## PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguages

> PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguages(ctx).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()

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
	r, err := apiClient.BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguages(context.Background()).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PostCmsV3BlogsAuthorsMultiLanguageUpdateLanguages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsAuthorsMultiLanguageUpdateLanguagesRequest struct via the builder pattern


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


## PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimary

> PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimary(ctx).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()

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
	r, err := apiClient.BlogAuthorsAPI.PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimary(context.Background()).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogAuthorsAPI.PutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCmsV3BlogsAuthorsMultiLanguageSetNewLangPrimaryRequest struct via the builder pattern


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

