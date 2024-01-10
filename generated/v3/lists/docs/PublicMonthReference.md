# PublicMonthReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceType** | **string** |  | [default to "MONTH"]
**Hour** | Pointer to **int32** |  | [optional] 
**Minute** | Pointer to **int32** |  | [optional] 
**Second** | Pointer to **int32** |  | [optional] 
**Millisecond** | Pointer to **int32** |  | [optional] 
**Day** | **int32** |  | 

## Methods

### NewPublicMonthReference

`func NewPublicMonthReference(referenceType string, day int32, ) *PublicMonthReference`

NewPublicMonthReference instantiates a new PublicMonthReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicMonthReferenceWithDefaults

`func NewPublicMonthReferenceWithDefaults() *PublicMonthReference`

NewPublicMonthReferenceWithDefaults instantiates a new PublicMonthReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceType

`func (o *PublicMonthReference) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *PublicMonthReference) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *PublicMonthReference) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.


### GetHour

`func (o *PublicMonthReference) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *PublicMonthReference) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *PublicMonthReference) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *PublicMonthReference) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *PublicMonthReference) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *PublicMonthReference) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *PublicMonthReference) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *PublicMonthReference) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetSecond

`func (o *PublicMonthReference) GetSecond() int32`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *PublicMonthReference) GetSecondOk() (*int32, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecond

`func (o *PublicMonthReference) SetSecond(v int32)`

SetSecond sets Second field to given value.

### HasSecond

`func (o *PublicMonthReference) HasSecond() bool`

HasSecond returns a boolean if a field has been set.

### GetMillisecond

`func (o *PublicMonthReference) GetMillisecond() int32`

GetMillisecond returns the Millisecond field if non-nil, zero value otherwise.

### GetMillisecondOk

`func (o *PublicMonthReference) GetMillisecondOk() (*int32, bool)`

GetMillisecondOk returns a tuple with the Millisecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillisecond

`func (o *PublicMonthReference) SetMillisecond(v int32)`

SetMillisecond sets Millisecond field to given value.

### HasMillisecond

`func (o *PublicMonthReference) HasMillisecond() bool`

HasMillisecond returns a boolean if a field has been set.

### GetDay

`func (o *PublicMonthReference) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *PublicMonthReference) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *PublicMonthReference) SetDay(v int32)`

SetDay sets Day field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


