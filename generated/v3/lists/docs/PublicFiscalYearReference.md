# PublicFiscalYearReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceType** | **string** |  | [default to "FISCAL_YEAR"]
**Hour** | Pointer to **int32** |  | [optional] 
**Minute** | Pointer to **int32** |  | [optional] 
**Second** | Pointer to **int32** |  | [optional] 
**Millisecond** | Pointer to **int32** |  | [optional] 
**Month** | **int32** |  | 
**Day** | **int32** |  | 

## Methods

### NewPublicFiscalYearReference

`func NewPublicFiscalYearReference(referenceType string, month int32, day int32, ) *PublicFiscalYearReference`

NewPublicFiscalYearReference instantiates a new PublicFiscalYearReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicFiscalYearReferenceWithDefaults

`func NewPublicFiscalYearReferenceWithDefaults() *PublicFiscalYearReference`

NewPublicFiscalYearReferenceWithDefaults instantiates a new PublicFiscalYearReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceType

`func (o *PublicFiscalYearReference) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *PublicFiscalYearReference) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *PublicFiscalYearReference) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.


### GetHour

`func (o *PublicFiscalYearReference) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *PublicFiscalYearReference) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *PublicFiscalYearReference) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *PublicFiscalYearReference) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *PublicFiscalYearReference) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *PublicFiscalYearReference) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *PublicFiscalYearReference) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *PublicFiscalYearReference) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetSecond

`func (o *PublicFiscalYearReference) GetSecond() int32`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *PublicFiscalYearReference) GetSecondOk() (*int32, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecond

`func (o *PublicFiscalYearReference) SetSecond(v int32)`

SetSecond sets Second field to given value.

### HasSecond

`func (o *PublicFiscalYearReference) HasSecond() bool`

HasSecond returns a boolean if a field has been set.

### GetMillisecond

`func (o *PublicFiscalYearReference) GetMillisecond() int32`

GetMillisecond returns the Millisecond field if non-nil, zero value otherwise.

### GetMillisecondOk

`func (o *PublicFiscalYearReference) GetMillisecondOk() (*int32, bool)`

GetMillisecondOk returns a tuple with the Millisecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillisecond

`func (o *PublicFiscalYearReference) SetMillisecond(v int32)`

SetMillisecond sets Millisecond field to given value.

### HasMillisecond

`func (o *PublicFiscalYearReference) HasMillisecond() bool`

HasMillisecond returns a boolean if a field has been set.

### GetMonth

`func (o *PublicFiscalYearReference) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *PublicFiscalYearReference) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *PublicFiscalYearReference) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetDay

`func (o *PublicFiscalYearReference) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *PublicFiscalYearReference) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *PublicFiscalYearReference) SetDay(v int32)`

SetDay sets Day field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


