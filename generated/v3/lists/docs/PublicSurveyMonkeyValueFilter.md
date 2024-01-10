# PublicSurveyMonkeyValueFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterType** | **string** |  | [default to "SURVEY_MONKEY_VALUE"]
**SurveyId** | **string** |  | 
**SurveyQuestion** | **string** |  | 
**SurveyAnswerRowId** | Pointer to **string** |  | [optional] 
**SurveyAnswerColId** | Pointer to **string** |  | [optional] 
**ValueComparison** | [**PublicPropertyFilterOperation**](PublicPropertyFilterOperation.md) |  | 
**Operator** | **string** |  | 

## Methods

### NewPublicSurveyMonkeyValueFilter

`func NewPublicSurveyMonkeyValueFilter(filterType string, surveyId string, surveyQuestion string, valueComparison PublicPropertyFilterOperation, operator string, ) *PublicSurveyMonkeyValueFilter`

NewPublicSurveyMonkeyValueFilter instantiates a new PublicSurveyMonkeyValueFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicSurveyMonkeyValueFilterWithDefaults

`func NewPublicSurveyMonkeyValueFilterWithDefaults() *PublicSurveyMonkeyValueFilter`

NewPublicSurveyMonkeyValueFilterWithDefaults instantiates a new PublicSurveyMonkeyValueFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterType

`func (o *PublicSurveyMonkeyValueFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *PublicSurveyMonkeyValueFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *PublicSurveyMonkeyValueFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetSurveyId

`func (o *PublicSurveyMonkeyValueFilter) GetSurveyId() string`

GetSurveyId returns the SurveyId field if non-nil, zero value otherwise.

### GetSurveyIdOk

`func (o *PublicSurveyMonkeyValueFilter) GetSurveyIdOk() (*string, bool)`

GetSurveyIdOk returns a tuple with the SurveyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyId

`func (o *PublicSurveyMonkeyValueFilter) SetSurveyId(v string)`

SetSurveyId sets SurveyId field to given value.


### GetSurveyQuestion

`func (o *PublicSurveyMonkeyValueFilter) GetSurveyQuestion() string`

GetSurveyQuestion returns the SurveyQuestion field if non-nil, zero value otherwise.

### GetSurveyQuestionOk

`func (o *PublicSurveyMonkeyValueFilter) GetSurveyQuestionOk() (*string, bool)`

GetSurveyQuestionOk returns a tuple with the SurveyQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyQuestion

`func (o *PublicSurveyMonkeyValueFilter) SetSurveyQuestion(v string)`

SetSurveyQuestion sets SurveyQuestion field to given value.


### GetSurveyAnswerRowId

`func (o *PublicSurveyMonkeyValueFilter) GetSurveyAnswerRowId() string`

GetSurveyAnswerRowId returns the SurveyAnswerRowId field if non-nil, zero value otherwise.

### GetSurveyAnswerRowIdOk

`func (o *PublicSurveyMonkeyValueFilter) GetSurveyAnswerRowIdOk() (*string, bool)`

GetSurveyAnswerRowIdOk returns a tuple with the SurveyAnswerRowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyAnswerRowId

`func (o *PublicSurveyMonkeyValueFilter) SetSurveyAnswerRowId(v string)`

SetSurveyAnswerRowId sets SurveyAnswerRowId field to given value.

### HasSurveyAnswerRowId

`func (o *PublicSurveyMonkeyValueFilter) HasSurveyAnswerRowId() bool`

HasSurveyAnswerRowId returns a boolean if a field has been set.

### GetSurveyAnswerColId

`func (o *PublicSurveyMonkeyValueFilter) GetSurveyAnswerColId() string`

GetSurveyAnswerColId returns the SurveyAnswerColId field if non-nil, zero value otherwise.

### GetSurveyAnswerColIdOk

`func (o *PublicSurveyMonkeyValueFilter) GetSurveyAnswerColIdOk() (*string, bool)`

GetSurveyAnswerColIdOk returns a tuple with the SurveyAnswerColId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurveyAnswerColId

`func (o *PublicSurveyMonkeyValueFilter) SetSurveyAnswerColId(v string)`

SetSurveyAnswerColId sets SurveyAnswerColId field to given value.

### HasSurveyAnswerColId

`func (o *PublicSurveyMonkeyValueFilter) HasSurveyAnswerColId() bool`

HasSurveyAnswerColId returns a boolean if a field has been set.

### GetValueComparison

`func (o *PublicSurveyMonkeyValueFilter) GetValueComparison() PublicPropertyFilterOperation`

GetValueComparison returns the ValueComparison field if non-nil, zero value otherwise.

### GetValueComparisonOk

`func (o *PublicSurveyMonkeyValueFilter) GetValueComparisonOk() (*PublicPropertyFilterOperation, bool)`

GetValueComparisonOk returns a tuple with the ValueComparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueComparison

`func (o *PublicSurveyMonkeyValueFilter) SetValueComparison(v PublicPropertyFilterOperation)`

SetValueComparison sets ValueComparison field to given value.


### GetOperator

`func (o *PublicSurveyMonkeyValueFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicSurveyMonkeyValueFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicSurveyMonkeyValueFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


