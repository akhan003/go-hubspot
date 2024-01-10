# PublicQuarterReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceType** | **string** |  | [default to "QUARTER"]
**Hour** | Pointer to **int32** |  | [optional] 
**Minute** | Pointer to **int32** |  | [optional] 
**Second** | Pointer to **int32** |  | [optional] 
**Millisecond** | Pointer to **int32** |  | [optional] 
**Month** | **int32** |  | 
**Day** | **int32** |  | 

## Methods

### NewPublicQuarterReference

`func NewPublicQuarterReference(referenceType string, month int32, day int32, ) *PublicQuarterReference`

NewPublicQuarterReference instantiates a new PublicQuarterReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicQuarterReferenceWithDefaults

`func NewPublicQuarterReferenceWithDefaults() *PublicQuarterReference`

NewPublicQuarterReferenceWithDefaults instantiates a new PublicQuarterReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceType

`func (o *PublicQuarterReference) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *PublicQuarterReference) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *PublicQuarterReference) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.


### GetHour

`func (o *PublicQuarterReference) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *PublicQuarterReference) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *PublicQuarterReference) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *PublicQuarterReference) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *PublicQuarterReference) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *PublicQuarterReference) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *PublicQuarterReference) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *PublicQuarterReference) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetSecond

`func (o *PublicQuarterReference) GetSecond() int32`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *PublicQuarterReference) GetSecondOk() (*int32, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecond

`func (o *PublicQuarterReference) SetSecond(v int32)`

SetSecond sets Second field to given value.

### HasSecond

`func (o *PublicQuarterReference) HasSecond() bool`

HasSecond returns a boolean if a field has been set.

### GetMillisecond

`func (o *PublicQuarterReference) GetMillisecond() int32`

GetMillisecond returns the Millisecond field if non-nil, zero value otherwise.

### GetMillisecondOk

`func (o *PublicQuarterReference) GetMillisecondOk() (*int32, bool)`

GetMillisecondOk returns a tuple with the Millisecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillisecond

`func (o *PublicQuarterReference) SetMillisecond(v int32)`

SetMillisecond sets Millisecond field to given value.

### HasMillisecond

`func (o *PublicQuarterReference) HasMillisecond() bool`

HasMillisecond returns a boolean if a field has been set.

### GetMonth

`func (o *PublicQuarterReference) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *PublicQuarterReference) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *PublicQuarterReference) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetDay

`func (o *PublicQuarterReference) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *PublicQuarterReference) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *PublicQuarterReference) SetDay(v int32)`

SetDay sets Day field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


