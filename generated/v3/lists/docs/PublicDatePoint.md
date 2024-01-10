# PublicDatePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeType** | **string** |  | [default to "DATE"]
**TimezoneSource** | Pointer to **string** |  | [optional] 
**ZoneId** | **string** |  | 
**Year** | **int32** |  | 
**Month** | **int32** |  | 
**Day** | **int32** |  | 
**Hour** | Pointer to **int32** |  | [optional] 
**Minute** | Pointer to **int32** |  | [optional] 
**Second** | Pointer to **int32** |  | [optional] 
**Millisecond** | Pointer to **int32** |  | [optional] 

## Methods

### NewPublicDatePoint

`func NewPublicDatePoint(timeType string, zoneId string, year int32, month int32, day int32, ) *PublicDatePoint`

NewPublicDatePoint instantiates a new PublicDatePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicDatePointWithDefaults

`func NewPublicDatePointWithDefaults() *PublicDatePoint`

NewPublicDatePointWithDefaults instantiates a new PublicDatePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeType

`func (o *PublicDatePoint) GetTimeType() string`

GetTimeType returns the TimeType field if non-nil, zero value otherwise.

### GetTimeTypeOk

`func (o *PublicDatePoint) GetTimeTypeOk() (*string, bool)`

GetTimeTypeOk returns a tuple with the TimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeType

`func (o *PublicDatePoint) SetTimeType(v string)`

SetTimeType sets TimeType field to given value.


### GetTimezoneSource

`func (o *PublicDatePoint) GetTimezoneSource() string`

GetTimezoneSource returns the TimezoneSource field if non-nil, zero value otherwise.

### GetTimezoneSourceOk

`func (o *PublicDatePoint) GetTimezoneSourceOk() (*string, bool)`

GetTimezoneSourceOk returns a tuple with the TimezoneSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneSource

`func (o *PublicDatePoint) SetTimezoneSource(v string)`

SetTimezoneSource sets TimezoneSource field to given value.

### HasTimezoneSource

`func (o *PublicDatePoint) HasTimezoneSource() bool`

HasTimezoneSource returns a boolean if a field has been set.

### GetZoneId

`func (o *PublicDatePoint) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *PublicDatePoint) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *PublicDatePoint) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.


### GetYear

`func (o *PublicDatePoint) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *PublicDatePoint) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *PublicDatePoint) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *PublicDatePoint) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *PublicDatePoint) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *PublicDatePoint) SetMonth(v int32)`

SetMonth sets Month field to given value.


### GetDay

`func (o *PublicDatePoint) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *PublicDatePoint) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *PublicDatePoint) SetDay(v int32)`

SetDay sets Day field to given value.


### GetHour

`func (o *PublicDatePoint) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *PublicDatePoint) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *PublicDatePoint) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *PublicDatePoint) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *PublicDatePoint) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *PublicDatePoint) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *PublicDatePoint) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *PublicDatePoint) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetSecond

`func (o *PublicDatePoint) GetSecond() int32`

GetSecond returns the Second field if non-nil, zero value otherwise.

### GetSecondOk

`func (o *PublicDatePoint) GetSecondOk() (*int32, bool)`

GetSecondOk returns a tuple with the Second field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecond

`func (o *PublicDatePoint) SetSecond(v int32)`

SetSecond sets Second field to given value.

### HasSecond

`func (o *PublicDatePoint) HasSecond() bool`

HasSecond returns a boolean if a field has been set.

### GetMillisecond

`func (o *PublicDatePoint) GetMillisecond() int32`

GetMillisecond returns the Millisecond field if non-nil, zero value otherwise.

### GetMillisecondOk

`func (o *PublicDatePoint) GetMillisecondOk() (*int32, bool)`

GetMillisecondOk returns a tuple with the Millisecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillisecond

`func (o *PublicDatePoint) SetMillisecond(v int32)`

SetMillisecond sets Millisecond field to given value.

### HasMillisecond

`func (o *PublicDatePoint) HasMillisecond() bool`

HasMillisecond returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


