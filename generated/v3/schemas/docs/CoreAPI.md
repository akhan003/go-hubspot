# \CoreAPI

All URIs are relative to *https://api.hubapi.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrmV3SchemasObjectType**](CoreAPI.md#DeleteCrmV3SchemasObjectType) | **Delete** /crm/v3/schemas/{objectType} | Delete a schema
[**DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifier**](CoreAPI.md#DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifier) | **Delete** /crm/v3/schemas/{objectType}/associations/{associationIdentifier} | Remove an association
[**GetCrmV3Schemas**](CoreAPI.md#GetCrmV3Schemas) | **Get** /crm/v3/schemas | Get all schemas
[**GetCrmV3SchemasObjectType**](CoreAPI.md#GetCrmV3SchemasObjectType) | **Get** /crm/v3/schemas/{objectType} | Get an existing schema
[**PatchCrmV3SchemasObjectType**](CoreAPI.md#PatchCrmV3SchemasObjectType) | **Patch** /crm/v3/schemas/{objectType} | Update a schema
[**PostCrmV3Schemas**](CoreAPI.md#PostCrmV3Schemas) | **Post** /crm/v3/schemas | Create a new schema
[**PostCrmV3SchemasObjectTypeAssociations**](CoreAPI.md#PostCrmV3SchemasObjectTypeAssociations) | **Post** /crm/v3/schemas/{objectType}/associations | Create an association



## DeleteCrmV3SchemasObjectType

> DeleteCrmV3SchemasObjectType(ctx, objectType).Archived(archived).Execute()

Delete a schema



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
	objectType := "objectType_example" // string | Fully qualified name or object type ID of your schema.
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.DeleteCrmV3SchemasObjectType(context.Background(), objectType).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.DeleteCrmV3SchemasObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Fully qualified name or object type ID of your schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3SchemasObjectTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

 (empty response body)

### Authorization

[private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifier

> DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifier(ctx, objectType, associationIdentifier).Execute()

Remove an association



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
	objectType := "objectType_example" // string | Fully qualified name or object type ID of your schema.
	associationIdentifier := "associationIdentifier_example" // string | Unique ID of the association to remove.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CoreAPI.DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifier(context.Background(), objectType, associationIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.DeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifier``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Fully qualified name or object type ID of your schema. | 
**associationIdentifier** | **string** | Unique ID of the association to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrmV3SchemasObjectTypeAssociationsAssociationIdentifierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3Schemas

> CollectionResponseObjectSchemaNoPaging GetCrmV3Schemas(ctx).Archived(archived).Execute()

Get all schemas



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
	archived := true // bool | Whether to return only results that have been archived. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.GetCrmV3Schemas(context.Background()).Archived(archived).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.GetCrmV3Schemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3Schemas`: CollectionResponseObjectSchemaNoPaging
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.GetCrmV3Schemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3SchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archived** | **bool** | Whether to return only results that have been archived. | [default to false]

### Return type

[**CollectionResponseObjectSchemaNoPaging**](CollectionResponseObjectSchemaNoPaging.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrmV3SchemasObjectType

> ObjectSchema GetCrmV3SchemasObjectType(ctx, objectType).Execute()

Get an existing schema



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
	objectType := "objectType_example" // string | Fully qualified name or object type ID of your schema.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.GetCrmV3SchemasObjectType(context.Background(), objectType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.GetCrmV3SchemasObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrmV3SchemasObjectType`: ObjectSchema
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.GetCrmV3SchemasObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Fully qualified name or object type ID of your schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrmV3SchemasObjectTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ObjectSchema**](ObjectSchema.md)

### Authorization

[oauth2_legacy](../README.md#oauth2_legacy), [private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchCrmV3SchemasObjectType

> ObjectTypeDefinition PatchCrmV3SchemasObjectType(ctx, objectType).ObjectTypeDefinitionPatch(objectTypeDefinitionPatch).Execute()

Update a schema



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
	objectType := "objectType_example" // string | Fully qualified name or object type ID of your schema.
	objectTypeDefinitionPatch := *openapiclient.NewObjectTypeDefinitionPatch() // ObjectTypeDefinitionPatch | Attributes to update in your schema.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.PatchCrmV3SchemasObjectType(context.Background(), objectType).ObjectTypeDefinitionPatch(objectTypeDefinitionPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.PatchCrmV3SchemasObjectType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchCrmV3SchemasObjectType`: ObjectTypeDefinition
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.PatchCrmV3SchemasObjectType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Fully qualified name or object type ID of your schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchCrmV3SchemasObjectTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectTypeDefinitionPatch** | [**ObjectTypeDefinitionPatch**](ObjectTypeDefinitionPatch.md) | Attributes to update in your schema. | 

### Return type

[**ObjectTypeDefinition**](ObjectTypeDefinition.md)

### Authorization

[private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3Schemas

> ObjectSchema PostCrmV3Schemas(ctx).ObjectSchemaEgg(objectSchemaEgg).Execute()

Create a new schema



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
	objectSchemaEgg := *openapiclient.NewObjectSchemaEgg([]string{"RequiredProperties_example"}, "my_object", []string{"AssociatedObjects_example"}, []openapiclient.ObjectTypePropertyCreate{*openapiclient.NewObjectTypePropertyCreate("My object property", "enumeration", "Name_example", "select")}, *openapiclient.NewObjectTypeDefinitionLabels()) // ObjectSchemaEgg | Object schema definition, including properties and associations.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.PostCrmV3Schemas(context.Background()).ObjectSchemaEgg(objectSchemaEgg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.PostCrmV3Schemas``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3Schemas`: ObjectSchema
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.PostCrmV3Schemas`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3SchemasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectSchemaEgg** | [**ObjectSchemaEgg**](ObjectSchemaEgg.md) | Object schema definition, including properties and associations. | 

### Return type

[**ObjectSchema**](ObjectSchema.md)

### Authorization

[private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCrmV3SchemasObjectTypeAssociations

> AssociationDefinition PostCrmV3SchemasObjectTypeAssociations(ctx, objectType).AssociationDefinitionEgg(associationDefinitionEgg).Execute()

Create an association



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
	objectType := "objectType_example" // string | Fully qualified name or object type ID of your schema.
	associationDefinitionEgg := *openapiclient.NewAssociationDefinitionEgg("2-123456", "contact") // AssociationDefinitionEgg | Attributes that define the association.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoreAPI.PostCrmV3SchemasObjectTypeAssociations(context.Background(), objectType).AssociationDefinitionEgg(associationDefinitionEgg).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoreAPI.PostCrmV3SchemasObjectTypeAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCrmV3SchemasObjectTypeAssociations`: AssociationDefinition
	fmt.Fprintf(os.Stdout, "Response from `CoreAPI.PostCrmV3SchemasObjectTypeAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectType** | **string** | Fully qualified name or object type ID of your schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCrmV3SchemasObjectTypeAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **associationDefinitionEgg** | [**AssociationDefinitionEgg**](AssociationDefinitionEgg.md) | Attributes that define the association. | 

### Return type

[**AssociationDefinition**](AssociationDefinition.md)

### Authorization

[private_apps_legacy](../README.md#private_apps_legacy)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

