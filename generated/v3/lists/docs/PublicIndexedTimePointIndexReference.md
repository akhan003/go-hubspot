# PublicIndexedTimePointIndexReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceType** | **string** |  | [default to "MONTH"]
**Hour** | Pointer to **int32** |  | [optional] 
**Minute** | Pointer to **int32** |  | [optional] 
**Second** | Pointer to **int32** |  | [optional] 
**Millisecond** | Pointer to **int32** |  | [optional] 
**DayOfWeek** | **string** |  | 
**Month** | **int32** |  | 
**Day** | **int32** |  | 

## Methods

### NewPublicIndexedTimePointIndexReference

`func NewPublicIndexedTimePointIndexReference(referenceType string, dayOfWeek string, month int32, day int32, ) *PublicIndexedTimePointIndexReference`

NewPublicIndexedTimePointIndexReference instantiates a new PublicIndexedTimePointIndexReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicIndexedTimePointIndexReferenceWithDefaults

`func NewPublicIndexedTimePointIndexReferenceWithDefaults() *PublicIndexedTimePointIndexReference`

NewPublicIndexedTimePointIndexReferenceWithDefaults instantiates a new PublicIndexedTimePointIndexReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceType

`func (o *PublicIndexedTimePointIndexReference) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *PublicIndexedTimePointIndexReference) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *PublicIndexedTimePointIndexReference) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.


### GetHour

`func (o *PublicIndexedTimePointIndexReference) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *PublicIndexedTimePointIndexReference) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *PublicIndexedTimePointIndexReference) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *PublicIndexedTimePointIndexReference) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *PublicIndexedTimePointIndexReference) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *PublicIndexedTimePointIndexReference) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *PublicIndexedTimePointIndexReference) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *PublicIndexedTimePointIndexReference) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetSecond

`func (o *PublicIndexedTimePointIndexReference) GetSecond() int32`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *PublicIndexedTimePointIndexReference) GetSecondOk() (*int32, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecond

`func (o *PublicIndexedTimePointIndexReference) SetSecond(v int32)`

SetSecond sets Second field to given value.

### HasSecond

`func (o *PublicIndexedTimePointIndexReference) HasSecond() bool`

HasSecond returns a boolean if a field has been set.

### GetMillisecond

`func (o *PublicIndexedTimePointIndexReference) GetMillisecond() int32`

GetMillisecond returns the Millisecond field if non-nil, zero value otherwise.

### GetMillisecondOk

`func (o *PublicIndexedTimePointIndexReference) GetMillisecondOk() (*int32, bool)`

GetMillisecondOk returns a tuple with the Millisecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillisecond

`func (o *PublicIndexedTimePointIndexReference) SetMillisecond(v int32)`

SetMillisecond sets Millisecond field to given value.

### HasMillisecond

`func (o *PublicIndexedTimePointIndexReference) HasMillisecond() bool`

HasMillisecond returns a boolean if a field has been set.

### GetDayOfWeek

`func (o *PublicIndexedTimePointIndexReference) GetDayOfWeek() string`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *PublicIndexedTimePointIndexReference) GetDayOfWeekOk() (*string, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *PublicIndexedTimePointIndexReference) SetDayOfWeek(v string)`

SetDayOfWeek sets DayOfWeek field to given value.


### GetMonth

`func (o *PublicIndexedTimePointIndexReference) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *PublicIndexedTimePointIndexReference) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *PublicIndexedTimePointIndexReference) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetDay

`func (o *PublicIndexedTimePointIndexReference) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *PublicIndexedTimePointIndexReference) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *PublicIndexedTimePointIndexReference) SetDay(v int32)`

SetDay sets Day field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


