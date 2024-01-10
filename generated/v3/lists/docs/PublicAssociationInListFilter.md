# PublicAssociationInListFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "ASSOCIATION"]
**CoalescingRefineBy** | [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | 
**Operator** | **string** |  | 
**ListId** | **int32** |  | 
**ToObjectTypeId** | Pointer to **string** |  | [optional] 
**AssociationTypeId** | **int32** |  | 
**AssociationCategory** | **string** |  | 
**ToObjectType** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicAssociationInListFilter

`func NewPublicAssociationInListFilter(filterType string, coalescingRefineBy PublicEventAnalyticsFilterCoalescingRefineBy, operator string, listId int32, associationTypeId int32, associationCategory string, ) *PublicAssociationInListFilter`

NewPublicAssociationInListFilter instantiates a new PublicAssociationInListFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAssociationInListFilterWithDefaults

`func NewPublicAssociationInListFilterWithDefaults() *PublicAssociationInListFilter`

NewPublicAssociationInListFilterWithDefaults instantiates a new PublicAssociationInListFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicAssociationInListFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicAssociationInListFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicAssociationInListFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetCoalescingRefineBy

`func (o *PublicAssociationInListFilter) GetCoalescingRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *PublicAssociationInListFilter) GetCoalescingRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *PublicAssociationInListFilter) SetCoalescingRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.


### GetOperator

`func (o *PublicAssociationInListFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicAssociationInListFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicAssociationInListFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetListId

`func (o *PublicAssociationInListFilter) GetListId() int32`

GetListId returns the ListId field if non-nil, zero value otherwise.

### GetListIdOk

`func (o *PublicAssociationInListFilter) GetListIdOk() (*int32, bool)`

GetListIdOk returns a tuple with the ListId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListId

`func (o *PublicAssociationInListFilter) SetListId(v int32)`

SetListId sets ListId field to given value.


### GetToObjectTypeId

`func (o *PublicAssociationInListFilter) GetToObjectTypeId() string`

GetToObjectTypeId returns the ToObjectTypeId field if non-nil, zero value otherwise.

### GetToObjectTypeIdOk

`func (o *PublicAssociationInListFilter) GetToObjectTypeIdOk() (*string, bool)`

GetToObjectTypeIdOk returns a tuple with the ToObjectTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectTypeId

`func (o *PublicAssociationInListFilter) SetToObjectTypeId(v string)`

SetToObjectTypeId sets ToObjectTypeId field to given value.

### HasToObjectTypeId

`func (o *PublicAssociationInListFilter) HasToObjectTypeId() bool`

HasToObjectTypeId returns a boolean if a field has been set.

### GetAssociationTypeId

`func (o *PublicAssociationInListFilter) GetAssociationTypeId() int32`

GetAssociationTypeId returns the AssociationTypeId field if non-nil, zero value otherwise.

### GetAssociationTypeIdOk

`func (o *PublicAssociationInListFilter) GetAssociationTypeIdOk() (*int32, bool)`

GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypeId

`func (o *PublicAssociationInListFilter) SetAssociationTypeId(v int32)`

SetAssociationTypeId sets AssociationTypeId field to given value.


### GetAssociationCategory

`func (o *PublicAssociationInListFilter) GetAssociationCategory() string`

GetAssociationCategory returns the AssociationCategory field if non-nil, zero value otherwise.

### GetAssociationCategoryOk

`func (o *PublicAssociationInListFilter) GetAssociationCategoryOk() (*string, bool)`

GetAssociationCategoryOk returns a tuple with the AssociationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCategory

`func (o *PublicAssociationInListFilter) SetAssociationCategory(v string)`

SetAssociationCategory sets AssociationCategory field to given value.


### GetToObjectType

`func (o *PublicAssociationInListFilter) GetToObjectType() string`

GetToObjectType returns the ToObjectType field if non-nil, zero value otherwise.

### GetToObjectTypeOk

`func (o *PublicAssociationInListFilter) GetToObjectTypeOk() (*string, bool)`

GetToObjectTypeOk returns a tuple with the ToObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToObjectType

`func (o *PublicAssociationInListFilter) SetToObjectType(v string)`

SetToObjectType sets ToObjectType field to given value.

### HasToObjectType

`func (o *PublicAssociationInListFilter) HasToObjectType() bool`

HasToObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


