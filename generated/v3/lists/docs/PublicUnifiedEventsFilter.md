# PublicUnifiedEventsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "UNIFIED_EVENTS"]
**PruningRefineBy** | Pointer to [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | [optional] 
**CoalescingRefineBy** | Pointer to [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | [optional] 
**EventTypeId** | Pointer to **string** |  | [optional] 
**FilterLines** | [**[]PublicEventFilterMetadata**](PublicEventFilterMetadata.md) |  | 

## Methods

### NewPublicUnifiedEventsFilter

`func NewPublicUnifiedEventsFilter(filterType string, filterLines []PublicEventFilterMetadata, ) *PublicUnifiedEventsFilter`

NewPublicUnifiedEventsFilter instantiates a new PublicUnifiedEventsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicUnifiedEventsFilterWithDefaults

`func NewPublicUnifiedEventsFilterWithDefaults() *PublicUnifiedEventsFilter`

NewPublicUnifiedEventsFilterWithDefaults instantiates a new PublicUnifiedEventsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicUnifiedEventsFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicUnifiedEventsFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicUnifiedEventsFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetPruningRefineBy

`func (o *PublicUnifiedEventsFilter) GetPruningRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetPruningRefineBy returns the PruningRefineBy field if non-nil, zero value otherwise.

### GetPruningRefineByOk

`func (o *PublicUnifiedEventsFilter) GetPruningRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetPruningRefineByOk returns a tuple with the PruningRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningRefineBy

`func (o *PublicUnifiedEventsFilter) SetPruningRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetPruningRefineBy sets PruningRefineBy field to given value.

### HasPruningRefineBy

`func (o *PublicUnifiedEventsFilter) HasPruningRefineBy() bool`

HasPruningRefineBy returns a boolean if a field has been set.

### GetCoalescingRefineBy

`func (o *PublicUnifiedEventsFilter) GetCoalescingRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *PublicUnifiedEventsFilter) GetCoalescingRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *PublicUnifiedEventsFilter) SetCoalescingRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.

### HasCoalescingRefineBy

`func (o *PublicUnifiedEventsFilter) HasCoalescingRefineBy() bool`

HasCoalescingRefineBy returns a boolean if a field has been set.

### GetEventTypeId

`func (o *PublicUnifiedEventsFilter) GetEventTypeId() string`

GetEventTypeId returns the EventTypeId field if non-nil, zero value otherwise.

### GetEventTypeIdOk

`func (o *PublicUnifiedEventsFilter) GetEventTypeIdOk() (*string, bool)`

GetEventTypeIdOk returns a tuple with the EventTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeId

`func (o *PublicUnifiedEventsFilter) SetEventTypeId(v string)`

SetEventTypeId sets EventTypeId field to given value.

### HasEventTypeId

`func (o *PublicUnifiedEventsFilter) HasEventTypeId() bool`

HasEventTypeId returns a boolean if a field has been set.

### GetFilterLines

`func (o *PublicUnifiedEventsFilter) GetFilterLines() []PublicEventFilterMetadata`

GetFilterLines returns the FilterLines field if non-nil, zero value otherwise.

### GetFilterLinesOk

`func (o *PublicUnifiedEventsFilter) GetFilterLinesOk() (*[]PublicEventFilterMetadata, bool)`

GetFilterLinesOk returns a tuple with the FilterLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLines

`func (o *PublicUnifiedEventsFilter) SetFilterLines(v []PublicEventFilterMetadata)`

SetFilterLines sets FilterLines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


