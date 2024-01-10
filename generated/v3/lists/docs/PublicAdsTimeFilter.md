# PublicAdsTimeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "ADS_TIME"]
**PruningRefineBy** | [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | 

## Methods

### NewPublicAdsTimeFilter

`func NewPublicAdsTimeFilter(filterType string, pruningRefineBy PublicEventAnalyticsFilterCoalescingRefineBy, ) *PublicAdsTimeFilter`

NewPublicAdsTimeFilter instantiates a new PublicAdsTimeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicAdsTimeFilterWithDefaults

`func NewPublicAdsTimeFilterWithDefaults() *PublicAdsTimeFilter`

NewPublicAdsTimeFilterWithDefaults instantiates a new PublicAdsTimeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicAdsTimeFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicAdsTimeFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicAdsTimeFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetPruningRefineBy

`func (o *PublicAdsTimeFilter) GetPruningRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetPruningRefineBy returns the PruningRefineBy field if non-nil, zero value otherwise.

### GetPruningRefineByOk

`func (o *PublicAdsTimeFilter) GetPruningRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetPruningRefineByOk returns a tuple with the PruningRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningRefineBy

`func (o *PublicAdsTimeFilter) SetPruningRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetPruningRefineBy sets PruningRefineBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


