# \MarketingEventsExternalApi

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveBatch**](MarketingEventsExternalApi.md#ArchiveBatch) | **Post** /marketing/v3/marketing-events/events/delete | 
[**Create**](MarketingEventsExternalApi.md#Create) | **Post** /marketing/v3/marketing-events/events | 
[**ExternalArchive**](MarketingEventsExternalApi.md#ExternalArchive) | **Delete** /marketing/v3/marketing-events/events/{externalEventId} | 
[**ExternalCancel**](MarketingEventsExternalApi.md#ExternalCancel) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/cancel | 
[**ExternalCompleteComplete**](MarketingEventsExternalApi.md#ExternalCompleteComplete) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/complete | 
[**ExternalEmailUpsertByID**](MarketingEventsExternalApi.md#ExternalEmailUpsertByID) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/{subscriberState}/email-upsert | 
[**ExternalGetByID**](MarketingEventsExternalApi.md#ExternalGetByID) | **Get** /marketing/v3/marketing-events/events/{externalEventId} | 
[**ExternalReplace**](MarketingEventsExternalApi.md#ExternalReplace) | **Put** /marketing/v3/marketing-events/events/{externalEventId} | 
[**ExternalUpdate**](MarketingEventsExternalApi.md#ExternalUpdate) | **Patch** /marketing/v3/marketing-events/events/{externalEventId} | 
[**ExternalUpsertByID**](MarketingEventsExternalApi.md#ExternalUpsertByID) | **Post** /marketing/v3/marketing-events/events/{externalEventId}/{subscriberState}/upsert | 
[**Upsert**](MarketingEventsExternalApi.md#Upsert) | **Post** /marketing/v3/marketing-events/events/upsert | 



## ArchiveBatch

> Error ArchiveBatch(ctx).BatchInputMarketingEventExternalUniqueIdentifier(batchInputMarketingEventExternalUniqueIdentifier).Execute()



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
    batchInputMarketingEventExternalUniqueIdentifier := *openapiclient.NewBatchInputMarketingEventExternalUniqueIdentifier([]openapiclient.MarketingEventExternalUniqueIdentifier{*openapiclient.NewMarketingEventExternalUniqueIdentifier(int32(123), "ExternalAccountId_example", "ExternalEventId_example")}) // BatchInputMarketingEventExternalUniqueIdentifier | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.ArchiveBatch(context.Background()).BatchInputMarketingEventExternalUniqueIdentifier(batchInputMarketingEventExternalUniqueIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.ArchiveBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveBatch`: Error
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.ArchiveBatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputMarketingEventExternalUniqueIdentifier** | [**BatchInputMarketingEventExternalUniqueIdentifier**](BatchInputMarketingEventExternalUniqueIdentifier.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> MarketingEventDefaultResponse Create(ctx).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()



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
    marketingEventCreateRequestParams := *openapiclient.NewMarketingEventCreateRequestParams("EventName_example", "EventOrganizer_example", "ExternalAccountId_example", "ExternalEventId_example") // MarketingEventCreateRequestParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.Create(context.Background()).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: MarketingEventDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **marketingEventCreateRequestParams** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md) |  | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalArchive

> ExternalArchive(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()



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
    externalEventId := "externalEventId_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MarketingEventsExternalApi.ExternalArchive(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.ExternalArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** |  | 

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


## ExternalCancel

> MarketingEventDefaultResponse ExternalCancel(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()



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
    externalEventId := "externalEventId_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.ExternalCancel(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.ExternalCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalCancel`: MarketingEventDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.ExternalCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** |  | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalCompleteComplete

> MarketingEventDefaultResponse ExternalCompleteComplete(ctx, externalEventId).ExternalAccountId(externalAccountId).MarketingEventCompleteRequestParams(marketingEventCompleteRequestParams).Execute()



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
    externalEventId := "externalEventId_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 
    marketingEventCompleteRequestParams := *openapiclient.NewMarketingEventCompleteRequestParams(time.Now(), time.Now()) // MarketingEventCompleteRequestParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.ExternalCompleteComplete(context.Background(), externalEventId).ExternalAccountId(externalAccountId).MarketingEventCompleteRequestParams(marketingEventCompleteRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.ExternalCompleteComplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalCompleteComplete`: MarketingEventDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.ExternalCompleteComplete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalCompleteCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** |  | 
 **marketingEventCompleteRequestParams** | [**MarketingEventCompleteRequestParams**](MarketingEventCompleteRequestParams.md) |  | 

### Return type

[**MarketingEventDefaultResponse**](MarketingEventDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalEmailUpsertByID

> Error ExternalEmailUpsertByID(ctx, externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).Execute()



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
    externalEventId := "externalEventId_example" // string | 
    subscriberState := "subscriberState_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 
    batchInputMarketingEventEmailSubscriber := *openapiclient.NewBatchInputMarketingEventEmailSubscriber([]openapiclient.MarketingEventEmailSubscriber{*openapiclient.NewMarketingEventEmailSubscriber(int64(123), "Email_example")}) // BatchInputMarketingEventEmailSubscriber | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.ExternalEmailUpsertByID(context.Background(), externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventEmailSubscriber(batchInputMarketingEventEmailSubscriber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.ExternalEmailUpsertByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalEmailUpsertByID`: Error
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.ExternalEmailUpsertByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 
**subscriberState** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalEmailUpsertByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **externalAccountId** | **string** |  | 
 **batchInputMarketingEventEmailSubscriber** | [**BatchInputMarketingEventEmailSubscriber**](BatchInputMarketingEventEmailSubscriber.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalGetByID

> MarketingEventPublicReadResponse ExternalGetByID(ctx, externalEventId).ExternalAccountId(externalAccountId).Execute()



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
    externalEventId := "externalEventId_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.ExternalGetByID(context.Background(), externalEventId).ExternalAccountId(externalAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.ExternalGetByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalGetByID`: MarketingEventPublicReadResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.ExternalGetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** |  | 

### Return type

[**MarketingEventPublicReadResponse**](MarketingEventPublicReadResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalReplace

> MarketingEventPublicDefaultResponse ExternalReplace(ctx, externalEventId).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()



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
    externalEventId := "externalEventId_example" // string | 
    marketingEventCreateRequestParams := *openapiclient.NewMarketingEventCreateRequestParams("EventName_example", "EventOrganizer_example", "ExternalAccountId_example", "ExternalEventId_example") // MarketingEventCreateRequestParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.ExternalReplace(context.Background(), externalEventId).MarketingEventCreateRequestParams(marketingEventCreateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.ExternalReplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalReplace`: MarketingEventPublicDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.ExternalReplace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalReplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **marketingEventCreateRequestParams** | [**MarketingEventCreateRequestParams**](MarketingEventCreateRequestParams.md) |  | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalUpdate

> MarketingEventPublicDefaultResponse ExternalUpdate(ctx, externalEventId).ExternalAccountId(externalAccountId).MarketingEventUpdateRequestParams(marketingEventUpdateRequestParams).Execute()



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
    externalEventId := "externalEventId_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 
    marketingEventUpdateRequestParams := *openapiclient.NewMarketingEventUpdateRequestParams() // MarketingEventUpdateRequestParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.ExternalUpdate(context.Background(), externalEventId).ExternalAccountId(externalAccountId).MarketingEventUpdateRequestParams(marketingEventUpdateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.ExternalUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalUpdate`: MarketingEventPublicDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.ExternalUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **externalAccountId** | **string** |  | 
 **marketingEventUpdateRequestParams** | [**MarketingEventUpdateRequestParams**](MarketingEventUpdateRequestParams.md) |  | 

### Return type

[**MarketingEventPublicDefaultResponse**](MarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExternalUpsertByID

> Error ExternalUpsertByID(ctx, externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).Execute()



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
    externalEventId := "externalEventId_example" // string | 
    subscriberState := "subscriberState_example" // string | 
    externalAccountId := "externalAccountId_example" // string | 
    batchInputMarketingEventSubscriber := *openapiclient.NewBatchInputMarketingEventSubscriber([]openapiclient.MarketingEventSubscriber{*openapiclient.NewMarketingEventSubscriber(int64(123))}) // BatchInputMarketingEventSubscriber | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.ExternalUpsertByID(context.Background(), externalEventId, subscriberState).ExternalAccountId(externalAccountId).BatchInputMarketingEventSubscriber(batchInputMarketingEventSubscriber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.ExternalUpsertByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExternalUpsertByID`: Error
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.ExternalUpsertByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalEventId** | **string** |  | 
**subscriberState** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExternalUpsertByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **externalAccountId** | **string** |  | 
 **batchInputMarketingEventSubscriber** | [**BatchInputMarketingEventSubscriber**](BatchInputMarketingEventSubscriber.md) |  | 

### Return type

[**Error**](Error.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Upsert

> BatchResponseMarketingEventPublicDefaultResponse Upsert(ctx).BatchInputMarketingEventCreateRequestParams(batchInputMarketingEventCreateRequestParams).Execute()



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
    batchInputMarketingEventCreateRequestParams := *openapiclient.NewBatchInputMarketingEventCreateRequestParams([]openapiclient.MarketingEventCreateRequestParams{*openapiclient.NewMarketingEventCreateRequestParams("EventName_example", "EventOrganizer_example", "ExternalAccountId_example", "ExternalEventId_example")}) // BatchInputMarketingEventCreateRequestParams | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketingEventsExternalApi.Upsert(context.Background()).BatchInputMarketingEventCreateRequestParams(batchInputMarketingEventCreateRequestParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketingEventsExternalApi.Upsert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Upsert`: BatchResponseMarketingEventPublicDefaultResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketingEventsExternalApi.Upsert`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpsertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **batchInputMarketingEventCreateRequestParams** | [**BatchInputMarketingEventCreateRequestParams**](BatchInputMarketingEventCreateRequestParams.md) |  | 

### Return type

[**BatchResponseMarketingEventPublicDefaultResponse**](BatchResponseMarketingEventPublicDefaultResponse.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

