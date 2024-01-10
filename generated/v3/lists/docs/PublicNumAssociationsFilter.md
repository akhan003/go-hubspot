# PublicNumAssociationsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "NUM_ASSOCIATIONS"]
**AssociationTypeId** | **int32** |  | 
**AssociationCategory** | **string** |  | 
**CoalescingRefineBy** | [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | 

## Methods

### NewPublicNumAssociationsFilter

`func NewPublicNumAssociationsFilter(filterType string, associationTypeId int32, associationCategory string, coalescingRefineBy PublicEventAnalyticsFilterCoalescingRefineBy, ) *PublicNumAssociationsFilter`

NewPublicNumAssociationsFilter instantiates a new PublicNumAssociationsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicNumAssociationsFilterWithDefaults

`func NewPublicNumAssociationsFilterWithDefaults() *PublicNumAssociationsFilter`

NewPublicNumAssociationsFilterWithDefaults instantiates a new PublicNumAssociationsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicNumAssociationsFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicNumAssociationsFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicNumAssociationsFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetAssociationTypeId

`func (o *PublicNumAssociationsFilter) GetAssociationTypeId() int32`

GetAssociationTypeId returns the AssociationTypeId field if non-nil, zero value otherwise.

### GetAssociationTypeIdOk

`func (o *PublicNumAssociationsFilter) GetAssociationTypeIdOk() (*int32, bool)`

GetAssociationTypeIdOk returns a tuple with the AssociationTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTypeId

`func (o *PublicNumAssociationsFilter) SetAssociationTypeId(v int32)`

SetAssociationTypeId sets AssociationTypeId field to given value.


### GetAssociationCategory

`func (o *PublicNumAssociationsFilter) GetAssociationCategory() string`

GetAssociationCategory returns the AssociationCategory field if non-nil, zero value otherwise.

### GetAssociationCategoryOk

`func (o *PublicNumAssociationsFilter) GetAssociationCategoryOk() (*string, bool)`

GetAssociationCategoryOk returns a tuple with the AssociationCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationCategory

`func (o *PublicNumAssociationsFilter) SetAssociationCategory(v string)`

SetAssociationCategory sets AssociationCategory field to given value.


### GetCoalescingRefineBy

`func (o *PublicNumAssociationsFilter) GetCoalescingRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *PublicNumAssociationsFilter) GetCoalescingRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *PublicNumAssociationsFilter) SetCoalescingRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


