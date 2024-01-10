# PublicIntegrationEventFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "INTEGRATION_EVENT"]
**EventTypeId** | **int32** |  | 
**FilterLines** | [**[]PublicEventFilterMetadata**](PublicEventFilterMetadata.md) |  | 

## Methods

### NewPublicIntegrationEventFilter

`func NewPublicIntegrationEventFilter(filterType string, eventTypeId int32, filterLines []PublicEventFilterMetadata, ) *PublicIntegrationEventFilter`

NewPublicIntegrationEventFilter instantiates a new PublicIntegrationEventFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicIntegrationEventFilterWithDefaults

`func NewPublicIntegrationEventFilterWithDefaults() *PublicIntegrationEventFilter`

NewPublicIntegrationEventFilterWithDefaults instantiates a new PublicIntegrationEventFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicIntegrationEventFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicIntegrationEventFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicIntegrationEventFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetEventTypeId

`func (o *PublicIntegrationEventFilter) GetEventTypeId() int32`

GetEventTypeId returns the EventTypeId field if non-nil, zero value otherwise.

### GetEventTypeIdOk

`func (o *PublicIntegrationEventFilter) GetEventTypeIdOk() (*int32, bool)`

GetEventTypeIdOk returns a tuple with the EventTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeId

`func (o *PublicIntegrationEventFilter) SetEventTypeId(v int32)`

SetEventTypeId sets EventTypeId field to given value.


### GetFilterLines

`func (o *PublicIntegrationEventFilter) GetFilterLines() []PublicEventFilterMetadata`

GetFilterLines returns the FilterLines field if non-nil, zero value otherwise.

### GetFilterLinesOk

`func (o *PublicIntegrationEventFilter) GetFilterLinesOk() (*[]PublicEventFilterMetadata, bool)`

GetFilterLinesOk returns a tuple with the FilterLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterLines

`func (o *PublicIntegrationEventFilter) SetFilterLines(v []PublicEventFilterMetadata)`

SetFilterLines sets FilterLines field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


