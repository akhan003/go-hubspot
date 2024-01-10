# PublicTimePointOperationTimePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeType** | **string** |  | [default to "PROPERTY_REFERENCED"]
**TimezoneSource** | Pointer to **string** |  | [optional] 
**ZoneId** | **string** |  | 
**Year** | **int32** |  | 
**Month** | **int32** |  | 
**Day** | **int32** |  | 
**Hour** | Pointer to **int32** |  | [optional] 
**Minute** | Pointer to **int32** |  | [optional] 
**Second** | Pointer to **int32** |  | [optional] 
**Millisecond** | Pointer to **int32** |  | [optional] 
**IndexReference** | [**PublicIndexedTimePointIndexReference**](PublicIndexedTimePointIndexReference.md) |  | 
**Offset** | Pointer to [**PublicIndexOffset**](PublicIndexOffset.md) |  | [optional] 
**Property** | **string** |  | 
**ReferenceType** | **string** |  | 

## Methods

### NewPublicTimePointOperationTimePoint

`func NewPublicTimePointOperationTimePoint(timeType string, zoneId string, year int32, month int32, day int32, indexReference PublicIndexedTimePointIndexReference, property string, referenceType string, ) *PublicTimePointOperationTimePoint`

NewPublicTimePointOperationTimePoint instantiates a new PublicTimePointOperationTimePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTimePointOperationTimePointWithDefaults

`func NewPublicTimePointOperationTimePointWithDefaults() *PublicTimePointOperationTimePoint`

NewPublicTimePointOperationTimePointWithDefaults instantiates a new PublicTimePointOperationTimePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeType

`func (o *PublicTimePointOperationTimePoint) GetTimeType() string`

GetTimeType returns the TimeType field if non-nil, zero value otherwise.

### GetTimeTypeOk

`func (o *PublicTimePointOperationTimePoint) GetTimeTypeOk() (*string, bool)`

GetTimeTypeOk returns a tuple with the TimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeType

`func (o *PublicTimePointOperationTimePoint) SetTimeType(v string)`

SetTimeType sets TimeType field to given value.


### GetTimezoneSource

`func (o *PublicTimePointOperationTimePoint) GetTimezoneSource() string`

GetTimezoneSource returns the TimezoneSource field if non-nil, zero value otherwise.

### GetTimezoneSourceOk

`func (o *PublicTimePointOperationTimePoint) GetTimezoneSourceOk() (*string, bool)`

GetTimezoneSourceOk returns a tuple with the TimezoneSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneSource

`func (o *PublicTimePointOperationTimePoint) SetTimezoneSource(v string)`

SetTimezoneSource sets TimezoneSource field to given value.

### HasTimezoneSource

`func (o *PublicTimePointOperationTimePoint) HasTimezoneSource() bool`

HasTimezoneSource returns a boolean if a field has been set.

### GetZoneId

`func (o *PublicTimePointOperationTimePoint) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *PublicTimePointOperationTimePoint) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *PublicTimePointOperationTimePoint) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.


### GetYear

`func (o *PublicTimePointOperationTimePoint) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *PublicTimePointOperationTimePoint) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *PublicTimePointOperationTimePoint) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *PublicTimePointOperationTimePoint) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *PublicTimePointOperationTimePoint) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *PublicTimePointOperationTimePoint) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetDay

`func (o *PublicTimePointOperationTimePoint) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *PublicTimePointOperationTimePoint) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *PublicTimePointOperationTimePoint) SetDay(v int32)`

SetDay sets Day field to given value.


### GetHour

`func (o *PublicTimePointOperationTimePoint) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *PublicTimePointOperationTimePoint) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *PublicTimePointOperationTimePoint) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *PublicTimePointOperationTimePoint) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *PublicTimePointOperationTimePoint) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *PublicTimePointOperationTimePoint) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *PublicTimePointOperationTimePoint) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *PublicTimePointOperationTimePoint) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetSecond

`func (o *PublicTimePointOperationTimePoint) GetSecond() int32`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *PublicTimePointOperationTimePoint) GetSecondOk() (*int32, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecond

`func (o *PublicTimePointOperationTimePoint) SetSecond(v int32)`

SetSecond sets Second field to given value.

### HasSecond

`func (o *PublicTimePointOperationTimePoint) HasSecond() bool`

HasSecond returns a boolean if a field has been set.

### GetMillisecond

`func (o *PublicTimePointOperationTimePoint) GetMillisecond() int32`

GetMillisecond returns the Millisecond field if non-nil, zero value otherwise.

### GetMillisecondOk

`func (o *PublicTimePointOperationTimePoint) GetMillisecondOk() (*int32, bool)`

GetMillisecondOk returns a tuple with the Millisecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillisecond

`func (o *PublicTimePointOperationTimePoint) SetMillisecond(v int32)`

SetMillisecond sets Millisecond field to given value.

### HasMillisecond

`func (o *PublicTimePointOperationTimePoint) HasMillisecond() bool`

HasMillisecond returns a boolean if a field has been set.

### GetIndexReference

`func (o *PublicTimePointOperationTimePoint) GetIndexReference() PublicIndexedTimePointIndexReference`

GetIndexReference returns the IndexReference field if non-nil, zero value otherwise.

### GetIndexReferenceOk

`func (o *PublicTimePointOperationTimePoint) GetIndexReferenceOk() (*PublicIndexedTimePointIndexReference, bool)`

GetIndexReferenceOk returns a tuple with the IndexReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexReference

`func (o *PublicTimePointOperationTimePoint) SetIndexReference(v PublicIndexedTimePointIndexReference)`

SetIndexReference sets IndexReference field to given value.


### GetOffset

`func (o *PublicTimePointOperationTimePoint) GetOffset() PublicIndexOffset`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PublicTimePointOperationTimePoint) GetOffsetOk() (*PublicIndexOffset, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PublicTimePointOperationTimePoint) SetOffset(v PublicIndexOffset)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PublicTimePointOperationTimePoint) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetProperty

`func (o *PublicTimePointOperationTimePoint) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *PublicTimePointOperationTimePoint) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *PublicTimePointOperationTimePoint) SetProperty(v string)`

SetProperty sets Property field to given value.


### GetReferenceType

`func (o *PublicTimePointOperationTimePoint) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *PublicTimePointOperationTimePoint) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *PublicTimePointOperationTimePoint) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


