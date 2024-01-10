# ListCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectTypeId** | **string** | The object type ID of the type of objects that the list will store. | 
**ProcessingType** | **string** | The processing type of the list. One of: &#x60;SNAPSHOT&#x60;, &#x60;MANUAL&#x60;, or &#x60;DYNAMIC&#x60;. | 
**ListFolderId** | Pointer to **int32** | The ID of the folder that the list should be created in. If left blank, then the list will be created in the root of the list folder structure. | [optional] 
**Name** | **string** | The name of the list, which must be globally unique across all public lists in the portal. | 
**FilterBranch** | Pointer to [**PublicPropertyAssociationFilterBranchFilterBranchesInner**](PublicPropertyAssociationFilterBranchFilterBranchesInner.md) |  | [optional] 

## Methods

### NewListCreateRequest

`func NewListCreateRequest(objectTypeId string, processingType string, name string, ) *ListCreateRequest`

NewListCreateRequest instantiates a new ListCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCreateRequestWithDefaults

`func NewListCreateRequestWithDefaults() *ListCreateRequest`

NewListCreateRequestWithDefaults instantiates a new ListCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectTypeId

`func (o *ListCreateRequest) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *ListCreateRequest) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *ListCreateRequest) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetProcessingType

`func (o *ListCreateRequest) GetProcessingType() string`

GetProcessingType returns the ProcessingType field if non-nil, zero value otherwise.

### GetProcessingTypeOk

`func (o *ListCreateRequest) GetProcessingTypeOk() (*string, bool)`

GetProcessingTypeOk returns a tuple with the ProcessingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingType

`func (o *ListCreateRequest) SetProcessingType(v string)`

SetProcessingType sets ProcessingType field to given value.


### GetListFolderId

`func (o *ListCreateRequest) GetListFolderId() int32`

GetListFolderId returns the ListFolderId field if non-nil, zero value otherwise.

### GetListFolderIdOk

`func (o *ListCreateRequest) GetListFolderIdOk() (*int32, bool)`

GetListFolderIdOk returns a tuple with the ListFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListFolderId

`func (o *ListCreateRequest) SetListFolderId(v int32)`

SetListFolderId sets ListFolderId field to given value.

### HasListFolderId

`func (o *ListCreateRequest) HasListFolderId() bool`

HasListFolderId returns a boolean if a field has been set.

### GetName

`func (o *ListCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFilterBranch

`func (o *ListCreateRequest) GetFilterBranch() PublicPropertyAssociationFilterBranchFilterBranchesInner`

GetFilterBranch returns the FilterBranch field if non-nil, zero value otherwise.

### GetFilterBranchOk

`func (o *ListCreateRequest) GetFilterBranchOk() (*PublicPropertyAssociationFilterBranchFilterBranchesInner, bool)`

GetFilterBranchOk returns a tuple with the FilterBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBranch

`func (o *ListCreateRequest) SetFilterBranch(v PublicPropertyAssociationFilterBranchFilterBranchesInner)`

SetFilterBranch sets FilterBranch field to given value.

### HasFilterBranch

`func (o *ListCreateRequest) HasFilterBranch() bool`

HasFilterBranch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


