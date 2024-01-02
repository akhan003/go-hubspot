# PublicCtaAnalyticsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "CTA"]
**CoalescingRefineBy** | Pointer to [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | [optional] 
**PruningRefineBy** | Pointer to [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | [optional] 
**Operator** | **string** |  | 
**CtaName** | **string** |  | 

## Methods

### NewPublicCtaAnalyticsFilter

`func NewPublicCtaAnalyticsFilter(filterType string, operator string, ctaName string, ) *PublicCtaAnalyticsFilter`

NewPublicCtaAnalyticsFilter instantiates a new PublicCtaAnalyticsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCtaAnalyticsFilterWithDefaults

`func NewPublicCtaAnalyticsFilterWithDefaults() *PublicCtaAnalyticsFilter`

NewPublicCtaAnalyticsFilterWithDefaults instantiates a new PublicCtaAnalyticsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicCtaAnalyticsFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicCtaAnalyticsFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicCtaAnalyticsFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetCoalescingRefineBy

`func (o *PublicCtaAnalyticsFilter) GetCoalescingRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *PublicCtaAnalyticsFilter) GetCoalescingRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *PublicCtaAnalyticsFilter) SetCoalescingRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.

### HasCoalescingRefineBy

`func (o *PublicCtaAnalyticsFilter) HasCoalescingRefineBy() bool`

HasCoalescingRefineBy returns a boolean if a field has been set.

### GetPruningRefineBy

`func (o *PublicCtaAnalyticsFilter) GetPruningRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetPruningRefineBy returns the PruningRefineBy field if non-nil, zero value otherwise.

### GetPruningRefineByOk

`func (o *PublicCtaAnalyticsFilter) GetPruningRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetPruningRefineByOk returns a tuple with the PruningRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningRefineBy

`func (o *PublicCtaAnalyticsFilter) SetPruningRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetPruningRefineBy sets PruningRefineBy field to given value.

### HasPruningRefineBy

`func (o *PublicCtaAnalyticsFilter) HasPruningRefineBy() bool`

HasPruningRefineBy returns a boolean if a field has been set.

### GetOperator

`func (o *PublicCtaAnalyticsFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicCtaAnalyticsFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicCtaAnalyticsFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetCtaName

`func (o *PublicCtaAnalyticsFilter) GetCtaName() string`

GetCtaName returns the CtaName field if non-nil, zero value otherwise.

### GetCtaNameOk

`func (o *PublicCtaAnalyticsFilter) GetCtaNameOk() (*string, bool)`

GetCtaNameOk returns a tuple with the CtaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtaName

`func (o *PublicCtaAnalyticsFilter) SetCtaName(v string)`

SetCtaName sets CtaName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


