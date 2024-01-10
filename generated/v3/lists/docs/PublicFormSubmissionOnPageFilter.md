# PublicFormSubmissionOnPageFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "FORM_SUBMISSION_ON_PAGE"]
**FormId** | Pointer to **string** |  | [optional] 
**PruningRefineBy** | Pointer to [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | [optional] 
**CoalescingRefineBy** | Pointer to [**PublicEventAnalyticsFilterCoalescingRefineBy**](PublicEventAnalyticsFilterCoalescingRefineBy.md) |  | [optional] 
**Operator** | **string** |  | 
**PageId** | **string** |  | 

## Methods

### NewPublicFormSubmissionOnPageFilter

`func NewPublicFormSubmissionOnPageFilter(filterType string, operator string, pageId string, ) *PublicFormSubmissionOnPageFilter`

NewPublicFormSubmissionOnPageFilter instantiates a new PublicFormSubmissionOnPageFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicFormSubmissionOnPageFilterWithDefaults

`func NewPublicFormSubmissionOnPageFilterWithDefaults() *PublicFormSubmissionOnPageFilter`

NewPublicFormSubmissionOnPageFilterWithDefaults instantiates a new PublicFormSubmissionOnPageFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicFormSubmissionOnPageFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicFormSubmissionOnPageFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicFormSubmissionOnPageFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetFormId

`func (o *PublicFormSubmissionOnPageFilter) GetFormId() string`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *PublicFormSubmissionOnPageFilter) GetFormIdOk() (*string, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *PublicFormSubmissionOnPageFilter) SetFormId(v string)`

SetFormId sets FormId field to given value.

### HasFormId

`func (o *PublicFormSubmissionOnPageFilter) HasFormId() bool`

HasFormId returns a boolean if a field has been set.

### GetPruningRefineBy

`func (o *PublicFormSubmissionOnPageFilter) GetPruningRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetPruningRefineBy returns the PruningRefineBy field if non-nil, zero value otherwise.

### GetPruningRefineByOk

`func (o *PublicFormSubmissionOnPageFilter) GetPruningRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetPruningRefineByOk returns a tuple with the PruningRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPruningRefineBy

`func (o *PublicFormSubmissionOnPageFilter) SetPruningRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetPruningRefineBy sets PruningRefineBy field to given value.

### HasPruningRefineBy

`func (o *PublicFormSubmissionOnPageFilter) HasPruningRefineBy() bool`

HasPruningRefineBy returns a boolean if a field has been set.

### GetCoalescingRefineBy

`func (o *PublicFormSubmissionOnPageFilter) GetCoalescingRefineBy() PublicEventAnalyticsFilterCoalescingRefineBy`

GetCoalescingRefineBy returns the CoalescingRefineBy field if non-nil, zero value otherwise.

### GetCoalescingRefineByOk

`func (o *PublicFormSubmissionOnPageFilter) GetCoalescingRefineByOk() (*PublicEventAnalyticsFilterCoalescingRefineBy, bool)`

GetCoalescingRefineByOk returns a tuple with the CoalescingRefineBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoalescingRefineBy

`func (o *PublicFormSubmissionOnPageFilter) SetCoalescingRefineBy(v PublicEventAnalyticsFilterCoalescingRefineBy)`

SetCoalescingRefineBy sets CoalescingRefineBy field to given value.

### HasCoalescingRefineBy

`func (o *PublicFormSubmissionOnPageFilter) HasCoalescingRefineBy() bool`

HasCoalescingRefineBy returns a boolean if a field has been set.

### GetOperator

`func (o *PublicFormSubmissionOnPageFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicFormSubmissionOnPageFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicFormSubmissionOnPageFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetPageId

`func (o *PublicFormSubmissionOnPageFilter) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *PublicFormSubmissionOnPageFilter) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *PublicFormSubmissionOnPageFilter) SetPageId(v string)`

SetPageId sets PageId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


