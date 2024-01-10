# PublicObjectListSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcessingType** | **string** | The processing type of the list. | 
**ObjectTypeId** | **string** | The object type of the list. | 
**UpdatedById** | Pointer to **int32** | The ID of the user that last updated the list. | [optional] 
**FiltersUpdatedAt** | Pointer to **time.Time** | The time when the filters for this list were last updated. | [optional] 
**ListId** | **int32** | The **ILS ID** of the list. | 
**CreatedAt** | Pointer to **time.Time** | The time when the list was created. | [optional] 
**ProcessingStatus** | **string** | The processing status of the list. | 
**DeletedAt** | Pointer to **time.Time** | The time when the list was deleted. | [optional] 
**ListVersion** | **int32** | The version of the list. | 
**Name** | **string** | The name of the list. | 
**AdditionalPropertiesField** | **map[string]string** | The name and value of any additional properties that exist for this list and that were included in the search request. | 
**CreatedById** | Pointer to **int32** | The ID of the user that created the list. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the list was last updated. | [optional] 

## Methods

### NewPublicObjectListSearchResult

`func NewPublicObjectListSearchResult(processingType string, objectTypeId string, listId int32, processingStatus string, listVersion int32, name string, additionalPropertiesField map[string]string, ) *PublicObjectListSearchResult`

NewPublicObjectListSearchResult instantiates a new PublicObjectListSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicObjectListSearchResultWithDefaults

`func NewPublicObjectListSearchResultWithDefaults() *PublicObjectListSearchResult`

NewPublicObjectListSearchResultWithDefaults instantiates a new PublicObjectListSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcessingType

`func (o *PublicObjectListSearchResult) GetProcessingType() string`

GetProcessingType returns the ProcessingType field if non-nil, zero value otherwise.

### GetProcessingTypeOk

`func (o *PublicObjectListSearchResult) GetProcessingTypeOk() (*string, bool)`

GetProcessingTypeOk returns a tuple with the ProcessingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingType

`func (o *PublicObjectListSearchResult) SetProcessingType(v string)`

SetProcessingType sets ProcessingType field to given value.


### GetObjectTypeId

`func (o *PublicObjectListSearchResult) GetObjectTypeId() string`

GetObjectTypeId returns the ObjectTypeId field if non-nil, zero value otherwise.

### GetObjectTypeIdOk

`func (o *PublicObjectListSearchResult) GetObjectTypeIdOk() (*string, bool)`

GetObjectTypeIdOk returns a tuple with the ObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectTypeId

`func (o *PublicObjectListSearchResult) SetObjectTypeId(v string)`

SetObjectTypeId sets ObjectTypeId field to given value.


### GetUpdatedById

`func (o *PublicObjectListSearchResult) GetUpdatedById() int32`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *PublicObjectListSearchResult) GetUpdatedByIdOk() (*int32, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *PublicObjectListSearchResult) SetUpdatedById(v int32)`

SetUpdatedById sets UpdatedById field to given value.

### HasUpdatedById

`func (o *PublicObjectListSearchResult) HasUpdatedById() bool`

HasUpdatedById returns a boolean if a field has been set.

### GetFiltersUpdatedAt

`func (o *PublicObjectListSearchResult) GetFiltersUpdatedAt() time.Time`

GetFiltersUpdatedAt returns the FiltersUpdatedAt field if non-nil, zero value otherwise.

### GetFiltersUpdatedAtOk

`func (o *PublicObjectListSearchResult) GetFiltersUpdatedAtOk() (*time.Time, bool)`

GetFiltersUpdatedAtOk returns a tuple with the FiltersUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersUpdatedAt

`func (o *PublicObjectListSearchResult) SetFiltersUpdatedAt(v time.Time)`

SetFiltersUpdatedAt sets FiltersUpdatedAt field to given value.

### HasFiltersUpdatedAt

`func (o *PublicObjectListSearchResult) HasFiltersUpdatedAt() bool`

HasFiltersUpdatedAt returns a boolean if a field has been set.

### GetListId

`func (o *PublicObjectListSearchResult) GetListId() int32`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *PublicObjectListSearchResult) GetListIdOk() (*int32, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *PublicObjectListSearchResult) SetListId(v int32)`

SetListId sets ListId field to given value.


### GetCreatedAt

`func (o *PublicObjectListSearchResult) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublicObjectListSearchResult) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublicObjectListSearchResult) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PublicObjectListSearchResult) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetProcessingStatus

`func (o *PublicObjectListSearchResult) GetProcessingStatus() string`

GetProcessingStatus returns the ProcessingStatus field if non-nil, zero value otherwise.

### GetProcessingStatusOk

`func (o *PublicObjectListSearchResult) GetProcessingStatusOk() (*string, bool)`

GetProcessingStatusOk returns a tuple with the ProcessingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingStatus

`func (o *PublicObjectListSearchResult) SetProcessingStatus(v string)`

SetProcessingStatus sets ProcessingStatus field to given value.


### GetDeletedAt

`func (o *PublicObjectListSearchResult) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *PublicObjectListSearchResult) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *PublicObjectListSearchResult) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *PublicObjectListSearchResult) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetListVersion

`func (o *PublicObjectListSearchResult) GetListVersion() int32`

GetListVersion returns the ListVersion field if non-nil, zero value otherwise.

### GetListVersionOk

`func (o *PublicObjectListSearchResult) GetListVersionOk() (*int32, bool)`

GetListVersionOk returns a tuple with the ListVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListVersion

`func (o *PublicObjectListSearchResult) SetListVersion(v int32)`

SetListVersion sets ListVersion field to given value.


### GetName

`func (o *PublicObjectListSearchResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublicObjectListSearchResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublicObjectListSearchResult) SetName(v string)`

SetName sets Name field to given value.


### GetAdditionalPropertiesField

`func (o *PublicObjectListSearchResult) GetAdditionalPropertiesField() map[string]string`

GetAdditionalPropertiesField returns the AdditionalPropertiesField field if non-nil, zero value otherwise.

### GetAdditionalPropertiesFieldOk

`func (o *PublicObjectListSearchResult) GetAdditionalPropertiesFieldOk() (*map[string]string, bool)`

GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPropertiesField

`func (o *PublicObjectListSearchResult) SetAdditionalPropertiesField(v map[string]string)`

SetAdditionalPropertiesField sets AdditionalPropertiesField field to given value.


### GetCreatedById

`func (o *PublicObjectListSearchResult) GetCreatedById() int32`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *PublicObjectListSearchResult) GetCreatedByIdOk() (*int32, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *PublicObjectListSearchResult) SetCreatedById(v int32)`

SetCreatedById sets CreatedById field to given value.

### HasCreatedById

`func (o *PublicObjectListSearchResult) HasCreatedById() bool`

HasCreatedById returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PublicObjectListSearchResult) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublicObjectListSearchResult) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublicObjectListSearchResult) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PublicObjectListSearchResult) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


