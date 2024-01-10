# PublicFormSubmissionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "FORM_SUBMISSION"]
**FormId** | Pointer to **string** |  | [optional] 
**PruningRefineBy** | Pointer to [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | [optional] 
**CoalescingRefineBy** | Pointer to [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | [optional] 
**Operator** | **string** |  | 

## Methods

### NewPublicFormSubmissionFilter

`func NewPublicFormSubmissionFilter(filterType string, operator string, ) *PublicFormSubmissionFilter`

NewPublicFormSubmissionFilter instantiates a new PublicFormSubmissionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicFormSubmissionFilterWithDefaults

`func NewPublicFormSubmissionFilterWithDefaults() *PublicFormSubmissionFilter`

NewPublicFormSubmissionFilterWithDefaults instantiates a new PublicFormSubmissionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicFormSubmissionFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicFormSubmissionFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicFormSubmissionFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetFormId

`func (o *PublicFormSubmissionFilter) GetFormId() string`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *PublicFormSubmissionFilter) GetFormIdOk() (*string, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *PublicFormSubmissionFilter) SetFormId(v string)`

SetFormId sets FormId field to given value.

### HasFormId

`func (o *PublicFormSubmissionFilter) HasFormId() bool`

HasFormId returns a boolean if a field has been set.

### GetPruningRefineBy

`func (o *PublicFormSubmissionFilter) GetPruningRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetPruningRefineBy returns the PruningRefineBy field if non-nil, zero value otherwise.

### GetPruningRefineByOk

`func (o *PublicFormSubmissionFilter) GetPruningRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetPruningRefineByOk returns a tuple with the PruningRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningRefineBy

`func (o *PublicFormSubmissionFilter) SetPruningRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetPruningRefineBy sets PruningRefineBy field to given value.

### HasPruningRefineBy

`func (o *PublicFormSubmissionFilter) HasPruningRefineBy() bool`

HasPruningRefineBy returns a boolean if a field has been set.

### GetCoalescingRefineBy

`func (o *PublicFormSubmissionFilter) GetCoalescingRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *PublicFormSubmissionFilter) GetCoalescingRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *PublicFormSubmissionFilter) SetCoalescingRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.

### HasCoalescingRefineBy

`func (o *PublicFormSubmissionFilter) HasCoalescingRefineBy() bool`

HasCoalescingRefineBy returns a boolean if a field has been set.

### GetOperator

`func (o *PublicFormSubmissionFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicFormSubmissionFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicFormSubmissionFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


