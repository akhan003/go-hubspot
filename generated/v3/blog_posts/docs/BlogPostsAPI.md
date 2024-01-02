# \BlogPostsAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCmsV3BlogsPostsObjectId**](BlogPostsAPI.md#DeleteCmsV3BlogsPostsObjectId) | **Delete** /cms/v3/blogs/posts/{objectId} | Delete a Blog Post
[**GetCmsV3BlogsPosts**](BlogPostsAPI.md#GetCmsV3BlogsPosts) | **Get** /cms/v3/blogs/posts | Get all Blog Posts
[**GetCmsV3BlogsPostsObjectId**](BlogPostsAPI.md#GetCmsV3BlogsPostsObjectId) | **Get** /cms/v3/blogs/posts/{objectId} | Retrieve a Blog Post
[**GetCmsV3BlogsPostsObjectIdDraft**](BlogPostsAPI.md#GetCmsV3BlogsPostsObjectIdDraft) | **Get** /cms/v3/blogs/posts/{objectId}/draft | Retrieve the full draft version of the Blog Post
[**GetCmsV3BlogsPostsObjectIdRevisions**](BlogPostsAPI.md#GetCmsV3BlogsPostsObjectIdRevisions) | **Get** /cms/v3/blogs/posts/{objectId}/revisions | Retrieves all the previous versions of a blog post
[**GetCmsV3BlogsPostsObjectIdRevisionsRevisionId**](BlogPostsAPI.md#GetCmsV3BlogsPostsObjectIdRevisionsRevisionId) | **Get** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId} | Retrieves a previous version of a blog post
[**PatchCmsV3BlogsPostsObjectId**](BlogPostsAPI.md#PatchCmsV3BlogsPostsObjectId) | **Patch** /cms/v3/blogs/posts/{objectId} | Update a Blog Post
[**PatchCmsV3BlogsPostsObjectIdDraft**](BlogPostsAPI.md#PatchCmsV3BlogsPostsObjectIdDraft) | **Patch** /cms/v3/blogs/posts/{objectId}/draft | Update a Blog Post draft
[**PostCmsV3BlogsPosts**](BlogPostsAPI.md#PostCmsV3BlogsPosts) | **Post** /cms/v3/blogs/posts | Create a new Blog Post
[**PostCmsV3BlogsPostsBatchArchive**](BlogPostsAPI.md#PostCmsV3BlogsPostsBatchArchive) | **Post** /cms/v3/blogs/posts/batch/archive | Delete a batch of Blog Posts
[**PostCmsV3BlogsPostsBatchCreate**](BlogPostsAPI.md#PostCmsV3BlogsPostsBatchCreate) | **Post** /cms/v3/blogs/posts/batch/create | Create a batch of Blog Posts
[**PostCmsV3BlogsPostsBatchRead**](BlogPostsAPI.md#PostCmsV3BlogsPostsBatchRead) | **Post** /cms/v3/blogs/posts/batch/read | Retrieve a batch of Blog Posts
[**PostCmsV3BlogsPostsBatchUpdate**](BlogPostsAPI.md#PostCmsV3BlogsPostsBatchUpdate) | **Post** /cms/v3/blogs/posts/batch/update | Update a batch of Blog Posts
[**PostCmsV3BlogsPostsClone**](BlogPostsAPI.md#PostCmsV3BlogsPostsClone) | **Post** /cms/v3/blogs/posts/clone | Clone a Blog Post
[**PostCmsV3BlogsPostsMultiLanguageAttachToLangGroup**](BlogPostsAPI.md#PostCmsV3BlogsPostsMultiLanguageAttachToLangGroup) | **Post** /cms/v3/blogs/posts/multi-language/attach-to-lang-group | Attach a Blog Post to a multi-language group
[**PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariation**](BlogPostsAPI.md#PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariation) | **Post** /cms/v3/blogs/posts/multi-language/create-language-variation | Create a new language variation
[**PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroup**](BlogPostsAPI.md#PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroup) | **Post** /cms/v3/blogs/posts/multi-language/detach-from-lang-group | Detach a Blog Post from a multi-language group
[**PostCmsV3BlogsPostsMultiLanguageUpdateLanguages**](BlogPostsAPI.md#PostCmsV3BlogsPostsMultiLanguageUpdateLanguages) | **Post** /cms/v3/blogs/posts/multi-language/update-languages | Update languages of multi-language group
[**PostCmsV3BlogsPostsObjectIdDraftPushLive**](BlogPostsAPI.md#PostCmsV3BlogsPostsObjectIdDraftPushLive) | **Post** /cms/v3/blogs/posts/{objectId}/draft/push-live | Push Blog Post draft edits live
[**PostCmsV3BlogsPostsObjectIdDraftReset**](BlogPostsAPI.md#PostCmsV3BlogsPostsObjectIdDraftReset) | **Post** /cms/v3/blogs/posts/{objectId}/draft/reset | Reset the Blog Post draft to the live version
[**PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestore**](BlogPostsAPI.md#PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestore) | **Post** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId}/restore | Restore a previous version of a blog post
[**PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraft**](BlogPostsAPI.md#PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraft) | **Post** /cms/v3/blogs/posts/{objectId}/revisions/{revisionId}/restore-to-draft | Restore a previous version of a blog post, to the draft version of the blog post
[**PostCmsV3BlogsPostsSchedule**](BlogPostsAPI.md#PostCmsV3BlogsPostsSchedule) | **Post** /cms/v3/blogs/posts/schedule | Schedule a Blog Post to be Published
[**PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimary**](BlogPostsAPI.md#PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimary) | **Put** /cms/v3/blogs/posts/multi-language/set-new-lang-primary | Set a new primary language



## DeleteCmsV3BlogsPostsObjectId

> DeleteCmsV3BlogsPostsObjectId(ctx, objectId).Archived(archived).Execute()

Delete a Blog Post



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
	objectId := "objectId_example" // string | The Blog Post id.
	archived := true // bool | Whether to return only results that have been archived. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.DeleteCmsV3BlogsPostsObjectId(context.Background(), objectId).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.DeleteCmsV3BlogsPostsObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCmsV3BlogsPostsObjectIdRequest struct via the builder pattern


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


## GetCmsV3BlogsPosts

> CollectionResponseWithTotalBlogPostForwardPaging GetCmsV3BlogsPosts(ctx).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()

Get all Blog Posts



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
	createdAt := time.Now() // time.Time | Only return Blog Posts created at exactly the specified time. (optional)
	createdAfter := time.Now() // time.Time | Only return Blog Posts created after the specified time. (optional)
	createdBefore := time.Now() // time.Time | Only return Blog Posts created before the specified time. (optional)
	updatedAt := time.Now() // time.Time | Only return Blog Posts last updated at exactly the specified time. (optional)
	updatedAfter := time.Now() // time.Time | Only return Blog Posts last updated after the specified time. (optional)
	updatedBefore := time.Now() // time.Time | Only return Blog Posts last updated before the specified time. (optional)
	sort := []string{"Inner_example"} // []string | Specifies which fields to use for sorting results. Valid fields are `name`, `createdAt`, `updatedAt`, `createdBy`, `updatedBy`. `createdAt` will be used by default. (optional)
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 20. (optional)
	archived := true // bool | Specifies whether to return deleted Blog Posts. Defaults to `false`. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetCmsV3BlogsPosts(context.Background()).CreatedAt(createdAt).CreatedAfter(createdAfter).CreatedBefore(createdBefore).UpdatedAt(updatedAt).UpdatedAfter(updatedAfter).UpdatedBefore(updatedBefore).Sort(sort).After(after).Limit(limit).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetCmsV3BlogsPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsPosts`: CollectionResponseWithTotalBlogPostForwardPaging
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetCmsV3BlogsPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAt** | **time.Time** | Only return Blog Posts created at exactly the specified time. | 
 **createdAfter** | **time.Time** | Only return Blog Posts created after the specified time. | 
 **createdBefore** | **time.Time** | Only return Blog Posts created before the specified time. | 
 **updatedAt** | **time.Time** | Only return Blog Posts last updated at exactly the specified time. | 
 **updatedAfter** | **time.Time** | Only return Blog Posts last updated after the specified time. | 
 **updatedBefore** | **time.Time** | Only return Blog Posts last updated before the specified time. | 
 **sort** | **[]string** | Specifies which fields to use for sorting results. Valid fields are &#x60;name&#x60;, &#x60;createdAt&#x60;, &#x60;updatedAt&#x60;, &#x60;createdBy&#x60;, &#x60;updatedBy&#x60;. &#x60;createdAt&#x60; will be used by default. | 
 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **limit** | **int32** | The maximum number of results to return. Default is 20. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Posts. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**CollectionResponseWithTotalBlogPostForwardPaging**](CollectionResponseWithTotalBlogPostForwardPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsPostsObjectId

> BlogPost GetCmsV3BlogsPostsObjectId(ctx, objectId).Archived(archived).Property(property).Execute()

Retrieve a Blog Post



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
	objectId := "objectId_example" // string | The Blog Post id.
	archived := true // bool | Specifies whether to return deleted Blog Posts. Defaults to `false`. (optional)
	property := "property_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetCmsV3BlogsPostsObjectId(context.Background(), objectId).Archived(archived).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetCmsV3BlogsPostsObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsPostsObjectId`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetCmsV3BlogsPostsObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Specifies whether to return deleted Blog Posts. Defaults to &#x60;false&#x60;. | 
 **property** | **string** |  | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsPostsObjectIdDraft

> BlogPost GetCmsV3BlogsPostsObjectIdDraft(ctx, objectId).Execute()

Retrieve the full draft version of the Blog Post



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
	objectId := "objectId_example" // string | The Blog Post id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetCmsV3BlogsPostsObjectIdDraft(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetCmsV3BlogsPostsObjectIdDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsPostsObjectIdDraft`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetCmsV3BlogsPostsObjectIdDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsObjectIdDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsPostsObjectIdRevisions

> CollectionResponseWithTotalVersionBlogPost GetCmsV3BlogsPostsObjectIdRevisions(ctx, objectId).After(after).Before(before).Limit(limit).Execute()

Retrieves all the previous versions of a blog post



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
	objectId := "objectId_example" // string | The Blog Post id.
	after := "after_example" // string | The cursor token value to get the next set of results. You can get this from the `paging.next.after` JSON property of a paged response containing more results. (optional)
	before := "before_example" // string |  (optional)
	limit := int32(56) // int32 | The maximum number of results to return. Default is 100. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetCmsV3BlogsPostsObjectIdRevisions(context.Background(), objectId).After(after).Before(before).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetCmsV3BlogsPostsObjectIdRevisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsPostsObjectIdRevisions`: CollectionResponseWithTotalVersionBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetCmsV3BlogsPostsObjectIdRevisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsObjectIdRevisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The cursor token value to get the next set of results. You can get this from the &#x60;paging.next.after&#x60; JSON property of a paged response containing more results. | 
 **before** | **string** |  | 
 **limit** | **int32** | The maximum number of results to return. Default is 100. | 

### Return type

[**CollectionResponseWithTotalVersionBlogPost**](CollectionResponseWithTotalVersionBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCmsV3BlogsPostsObjectIdRevisionsRevisionId

> VersionBlogPost GetCmsV3BlogsPostsObjectIdRevisionsRevisionId(ctx, objectId, revisionId).Execute()

Retrieves a previous version of a blog post



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
	objectId := "objectId_example" // string | The Blog Post id.
	revisionId := "revisionId_example" // string | The Blog Post version id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.GetCmsV3BlogsPostsObjectIdRevisionsRevisionId(context.Background(), objectId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.GetCmsV3BlogsPostsObjectIdRevisionsRevisionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCmsV3BlogsPostsObjectIdRevisionsRevisionId`: VersionBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.GetCmsV3BlogsPostsObjectIdRevisionsRevisionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 
**revisionId** | **string** | The Blog Post version id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCmsV3BlogsPostsObjectIdRevisionsRevisionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**VersionBlogPost**](VersionBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3BlogsPostsObjectId

> BlogPost PatchCmsV3BlogsPostsObjectId(ctx, objectId).BlogPost(blogPost).Archived(archived).Execute()

Update a Blog Post



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
	objectId := "objectId_example" // string | The Blog Post id.
	blogPost := *openapiclient.NewBlogPost(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", "HtmlTitle_example", false, map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", "RssBody_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", []int64{int64(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "PostSummary_example", "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "BlogAuthorId_example", "ContentGroupId_example", "RssSummary_example", false, "Url_example", []map[string]interface{}{map[string]interface{}(123)}, false, int64(123), "PostBody_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example") // BlogPost | The JSON representation of the updated Blog Post.
	archived := true // bool | Specifies whether to update deleted Blog Posts. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.PatchCmsV3BlogsPostsObjectId(context.Background(), objectId).BlogPost(blogPost).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PatchCmsV3BlogsPostsObjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCmsV3BlogsPostsObjectId`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.PatchCmsV3BlogsPostsObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3BlogsPostsObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of the updated Blog Post. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCmsV3BlogsPostsObjectIdDraft

> BlogPost PatchCmsV3BlogsPostsObjectIdDraft(ctx, objectId).BlogPost(blogPost).Execute()

Update a Blog Post draft



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
	objectId := "objectId_example" // string | The Blog Post id.
	blogPost := *openapiclient.NewBlogPost(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", "HtmlTitle_example", false, map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", "RssBody_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", []int64{int64(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "PostSummary_example", "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "BlogAuthorId_example", "ContentGroupId_example", "RssSummary_example", false, "Url_example", []map[string]interface{}{map[string]interface{}(123)}, false, int64(123), "PostBody_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example") // BlogPost | The JSON representation of the updated Blog Post to be applied to the draft.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.PatchCmsV3BlogsPostsObjectIdDraft(context.Background(), objectId).BlogPost(blogPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PatchCmsV3BlogsPostsObjectIdDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCmsV3BlogsPostsObjectIdDraft`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.PatchCmsV3BlogsPostsObjectIdDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCmsV3BlogsPostsObjectIdDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of the updated Blog Post to be applied to the draft. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPosts

> BlogPost PostCmsV3BlogsPosts(ctx).BlogPost(blogPost).Execute()

Create a new Blog Post



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
	blogPost := *openapiclient.NewBlogPost(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", "HtmlTitle_example", false, map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", "RssBody_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", []int64{int64(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "PostSummary_example", "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "BlogAuthorId_example", "ContentGroupId_example", "RssSummary_example", false, "Url_example", []map[string]interface{}{map[string]interface{}(123)}, false, int64(123), "PostBody_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example") // BlogPost | The JSON representation of a new Blog Post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPosts(context.Background()).BlogPost(blogPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPosts`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.PostCmsV3BlogsPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogPost** | [**BlogPost**](BlogPost.md) | The JSON representation of a new Blog Post. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsBatchArchive

> PostCmsV3BlogsPostsBatchArchive(ctx).BatchInputString(batchInputString).Execute()

Delete a batch of Blog Posts



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Post ids.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsBatchArchive(context.Background()).BatchInputString(batchInputString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsBatchArchive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsBatchArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Post ids. | 

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


## PostCmsV3BlogsPostsBatchCreate

> BatchResponseBlogPost PostCmsV3BlogsPostsBatchCreate(ctx).BatchInputBlogPost(batchInputBlogPost).Execute()

Create a batch of Blog Posts



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
	batchInputBlogPost := *openapiclient.NewBatchInputBlogPost([]openapiclient.BlogPost{*openapiclient.NewBlogPost(time.Now(), "Language_example", false, "MetaDescription_example", []map[string]map[string]interface{}{map[string]map[string]interface{}{"key": map[string]interface{}(123)}}, "Password_example", "HtmlTitle_example", false, map[string]ContentLanguageVariation{"key": *openapiclient.NewContentLanguageVariation(false, time.Now(), time.Now(), []map[string]interface{}{map[string]interface{}(123)}, "Password_example", "AuthorName_example", false, "Name_example", "Campaign_example", int64(123), "State_example", time.Now(), "Slug_example")}, "Id_example", "State_example", "Slug_example", "CreatedById_example", "RssBody_example", false, false, time.Now(), "ContentTypeCategory_example", "MabExperimentId_example", "UpdatedById_example", "TranslatedFromId_example", "FolderId_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), int32(123), "FeaturedImage_example", "AuthorName_example", "Domain_example", "Name_example", "DynamicPageHubDbTableId_example", "Campaign_example", "DynamicPageDataSourceId_example", false, false, map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": *openapiclient.NewLayoutSection("CssStyle_example", "Label_example", "Type_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, []map[string]LayoutSection{map[string]LayoutSection{"key": }}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", *openapiclient.NewStyles(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)), "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop(*openapiclient.NewRGBAColor(float32(123), int32(123), int32(123), int32(123)))})))}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", *openapiclient.NewStyles(, "FlexboxPositioning_example", *openapiclient.NewBackgroundImage("ImageUrl_example", "BackgroundSize_example", "BackgroundPosition_example"), false, "VerticalAlignment_example", int32(123), *openapiclient.NewGradient(*openapiclient.NewAngle("Units_example", float32(123)), *openapiclient.NewSideOrCorner("HorizontalSide_example", "VerticalSide_example"), []openapiclient.ColorStop{*openapiclient.NewColorStop()})))}}, []openapiclient.RowMetaData{*openapiclient.NewRowMetaData("CssClass_example", )}, []openapiclient.LayoutSection{}, "CssClass_example", int32(123), "CssId_example", int32(123), "Name_example", )}, time.Now(), "FooterHtml_example", []int64{int64(123)}, map[string]map[string]interface{}{"key": map[string]interface{}(123)}, "PostSummary_example", "HeadHtml_example", "PageExpiryRedirectUrl_example", "AbStatus_example", false, "AbTestId_example", "FeaturedImageAltText_example", "BlogAuthorId_example", "ContentGroupId_example", "RssSummary_example", false, "Url_example", []map[string]interface{}{map[string]interface{}(123)}, false, int64(123), "PostBody_example", map[string]map[string]interface{}{"key": map[string]interface{}(123)}, int64(123), false, "CurrentState_example", int32(123), "LinkRelCanonicalUrl_example")}) // BatchInputBlogPost | The JSON array of new Blog Posts to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsBatchCreate(context.Background()).BatchInputBlogPost(batchInputBlogPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsBatchCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsBatchCreate`: BatchResponseBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.PostCmsV3BlogsPostsBatchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsBatchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputBlogPost** | [**BatchInputBlogPost**](BatchInputBlogPost.md) | The JSON array of new Blog Posts to create. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsBatchRead

> BatchResponseBlogPost PostCmsV3BlogsPostsBatchRead(ctx).BatchInputString(batchInputString).Archived(archived).Execute()

Retrieve a batch of Blog Posts



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
	batchInputString := *openapiclient.NewBatchInputString([]string{"Inputs_example"}) // BatchInputString | The JSON array of Blog Post ids.
	archived := true // bool | Specifies whether to return deleted Blog Posts. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsBatchRead(context.Background()).BatchInputString(batchInputString).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsBatchRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsBatchRead`: BatchResponseBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.PostCmsV3BlogsPostsBatchRead`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsBatchReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputString** | [**BatchInputString**](BatchInputString.md) | The JSON array of Blog Post ids. | 
 **archived** | **bool** | Specifies whether to return deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsBatchUpdate

> BatchResponseBlogPost PostCmsV3BlogsPostsBatchUpdate(ctx).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()

Update a batch of Blog Posts



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
	batchInputJsonNode := *openapiclient.NewBatchInputJsonNode([]map[string]interface{}{map[string]interface{}(123)}) // BatchInputJsonNode | A JSON array of the JSON representations of the updated Blog Posts.
	archived := true // bool | Specifies whether to update deleted Blog Posts. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsBatchUpdate(context.Background()).BatchInputJsonNode(batchInputJsonNode).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsBatchUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsBatchUpdate`: BatchResponseBlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.PostCmsV3BlogsPostsBatchUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsBatchUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputJsonNode** | [**BatchInputJsonNode**](BatchInputJsonNode.md) | A JSON array of the JSON representations of the updated Blog Posts. | 
 **archived** | **bool** | Specifies whether to update deleted Blog Posts. Defaults to &#x60;false&#x60;. | 

### Return type

[**BatchResponseBlogPost**](BatchResponseBlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsClone

> BlogPost PostCmsV3BlogsPostsClone(ctx).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()

Clone a Blog Post



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
	contentCloneRequestVNext := *openapiclient.NewContentCloneRequestVNext("Id_example") // ContentCloneRequestVNext | The JSON representation of the ContentCloneRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsClone(context.Background()).ContentCloneRequestVNext(contentCloneRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsClone``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsClone`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.PostCmsV3BlogsPostsClone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsCloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentCloneRequestVNext** | [**ContentCloneRequestVNext**](ContentCloneRequestVNext.md) | The JSON representation of the ContentCloneRequest object. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsMultiLanguageAttachToLangGroup

> PostCmsV3BlogsPostsMultiLanguageAttachToLangGroup(ctx).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()

Attach a Blog Post to a multi-language group



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
	r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsMultiLanguageAttachToLangGroup(context.Background()).AttachToLangPrimaryRequestVNext(attachToLangPrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsMultiLanguageAttachToLangGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsMultiLanguageAttachToLangGroupRequest struct via the builder pattern


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


## PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariation

> BlogPost PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariation(ctx).BlogPostLanguageCloneRequestVNext(blogPostLanguageCloneRequestVNext).Execute()

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
	blogPostLanguageCloneRequestVNext := *openapiclient.NewBlogPostLanguageCloneRequestVNext("Id_example") // BlogPostLanguageCloneRequestVNext | The JSON representation of the BlogPostLanguageCloneRequestVNext object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariation(context.Background()).BlogPostLanguageCloneRequestVNext(blogPostLanguageCloneRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariation`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.PostCmsV3BlogsPostsMultiLanguageCreateLanguageVariation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsMultiLanguageCreateLanguageVariationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blogPostLanguageCloneRequestVNext** | [**BlogPostLanguageCloneRequestVNext**](BlogPostLanguageCloneRequestVNext.md) | The JSON representation of the BlogPostLanguageCloneRequestVNext object. | 

### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroup

> PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroup(ctx).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()

Detach a Blog Post from a multi-language group



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
	r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroup(context.Background()).DetachFromLangGroupRequestVNext(detachFromLangGroupRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsMultiLanguageDetachFromLangGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsMultiLanguageDetachFromLangGroupRequest struct via the builder pattern


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


## PostCmsV3BlogsPostsMultiLanguageUpdateLanguages

> PostCmsV3BlogsPostsMultiLanguageUpdateLanguages(ctx).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()

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
	updateLanguagesRequestVNext := *openapiclient.NewUpdateLanguagesRequestVNext(map[string]string{"key": "Inner_example"}, "PrimaryId_example") // UpdateLanguagesRequestVNext | The JSON representation of the SetNewLanguagePrimaryRequest object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsMultiLanguageUpdateLanguages(context.Background()).UpdateLanguagesRequestVNext(updateLanguagesRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsMultiLanguageUpdateLanguages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsMultiLanguageUpdateLanguagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateLanguagesRequestVNext** | [**UpdateLanguagesRequestVNext**](UpdateLanguagesRequestVNext.md) | The JSON representation of the SetNewLanguagePrimaryRequest object. | 

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


## PostCmsV3BlogsPostsObjectIdDraftPushLive

> PostCmsV3BlogsPostsObjectIdDraftPushLive(ctx, objectId).Execute()

Push Blog Post draft edits live



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
	objectId := "objectId_example" // string | The id of the Blog Post for which it's draft will be pushed live.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsObjectIdDraftPushLive(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsObjectIdDraftPushLive``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The id of the Blog Post for which it&#39;s draft will be pushed live. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsObjectIdDraftPushLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostCmsV3BlogsPostsObjectIdDraftReset

> PostCmsV3BlogsPostsObjectIdDraftReset(ctx, objectId).Execute()

Reset the Blog Post draft to the live version



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
	objectId := "objectId_example" // string | The id of the Blog Post for which it's draft will be reset.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsObjectIdDraftReset(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsObjectIdDraftReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The id of the Blog Post for which it&#39;s draft will be reset. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsObjectIdDraftResetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestore

> BlogPost PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestore(ctx, objectId, revisionId).Execute()

Restore a previous version of a blog post



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
	objectId := "objectId_example" // string | The Blog Post id.
	revisionId := "revisionId_example" // string | The Blog Post version id to restore.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestore(context.Background(), objectId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestore`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 
**revisionId** | **string** | The Blog Post version id to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraft

> BlogPost PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraft(ctx, objectId, revisionId).Execute()

Restore a previous version of a blog post, to the draft version of the blog post



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
	objectId := "objectId_example" // string | The Blog Post id.
	revisionId := int64(789) // int64 | The Blog Post version id to restore.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraft(context.Background(), objectId, revisionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraft`: BlogPost
	fmt.Fprintf(os.Stdout, "Response from `BlogPostsAPI.PostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The Blog Post id. | 
**revisionId** | **int64** | The Blog Post version id to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsObjectIdRevisionsRevisionIdRestoreToDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BlogPost**](BlogPost.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCmsV3BlogsPostsSchedule

> PostCmsV3BlogsPostsSchedule(ctx).ContentScheduleRequestVNext(contentScheduleRequestVNext).Execute()

Schedule a Blog Post to be Published



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
	contentScheduleRequestVNext := *openapiclient.NewContentScheduleRequestVNext(time.Now(), "Id_example") // ContentScheduleRequestVNext | The JSON representation of the ContentScheduleRequestVNext object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlogPostsAPI.PostCmsV3BlogsPostsSchedule(context.Background()).ContentScheduleRequestVNext(contentScheduleRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PostCmsV3BlogsPostsSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCmsV3BlogsPostsScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentScheduleRequestVNext** | [**ContentScheduleRequestVNext**](ContentScheduleRequestVNext.md) | The JSON representation of the ContentScheduleRequestVNext object. | 

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


## PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimary

> PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimary(ctx).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()

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
	r, err := apiClient.BlogPostsAPI.PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimary(context.Background()).SetNewLanguagePrimaryRequestVNext(setNewLanguagePrimaryRequestVNext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlogPostsAPI.PutCmsV3BlogsPostsMultiLanguageSetNewLangPrimary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutCmsV3BlogsPostsMultiLanguageSetNewLangPrimaryRequest struct via the builder pattern


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

