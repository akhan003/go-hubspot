# PublicInListFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "IN_LIST"]
**ListId** | **int32** |  | 
**Metadata** | Pointer to [**PublicInListFilterMetadata**](PublicInListFilterMetadata.md) |  | [optional] 
**Operator** | **string** |  | 

## Methods

### NewPublicInListFilter

`func NewPublicInListFilter(filterType string, listId int32, operator string, ) *PublicInListFilter`

NewPublicInListFilter instantiates a new PublicInListFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicInListFilterWithDefaults

`func NewPublicInListFilterWithDefaults() *PublicInListFilter`

NewPublicInListFilterWithDefaults instantiates a new PublicInListFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicInListFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicInListFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicInListFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetListId

`func (o *PublicInListFilter) GetListId() int32`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *PublicInListFilter) GetListIdOk() (*int32, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *PublicInListFilter) SetListId(v int32)`

SetListId sets ListId field to given value.


### GetMetadata

`func (o *PublicInListFilter) GetMetadata() PublicInListFilterMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PublicInListFilter) GetMetadataOk() (*PublicInListFilterMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PublicInListFilter) SetMetadata(v PublicInListFilterMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PublicInListFilter) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetOperator

`func (o *PublicInListFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicInListFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicInListFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


