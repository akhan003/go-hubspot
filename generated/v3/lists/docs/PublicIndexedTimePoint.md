# PublicIndexedTimePoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeType** | **string** |  | [default to "INDEXED"]
**TimezoneSource** | Pointer to **string** |  | [optional] 
**ZoneId** | **string** |  | 
**IndexReference** | [**PublicIndexedTimePointIndexReference**](PublicIndexedTimePointIndexReference.md) |  | 
**Offset** | Pointer to [**PublicIndexOffset**](PublicIndexOffset.md) |  | [optional] 

## Methods

### NewPublicIndexedTimePoint

`func NewPublicIndexedTimePoint(timeType string, zoneId string, indexReference PublicIndexedTimePointIndexReference, ) *PublicIndexedTimePoint`

NewPublicIndexedTimePoint instantiates a new PublicIndexedTimePoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicIndexedTimePointWithDefaults

`func NewPublicIndexedTimePointWithDefaults() *PublicIndexedTimePoint`

NewPublicIndexedTimePointWithDefaults instantiates a new PublicIndexedTimePoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeType

`func (o *PublicIndexedTimePoint) GetTimeType() string`

GetTimeType returns the TimeType field if non-nil, zero value otherwise.

### GetTimeTypeOk

`func (o *PublicIndexedTimePoint) GetTimeTypeOk() (*string, bool)`

GetTimeTypeOk returns a tuple with the TimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeType

`func (o *PublicIndexedTimePoint) SetTimeType(v string)`

SetTimeType sets TimeType field to given value.


### GetTimezoneSource

`func (o *PublicIndexedTimePoint) GetTimezoneSource() string`

GetTimezoneSource returns the TimezoneSource field if non-nil, zero value otherwise.

### GetTimezoneSourceOk

`func (o *PublicIndexedTimePoint) GetTimezoneSourceOk() (*string, bool)`

GetTimezoneSourceOk returns a tuple with the TimezoneSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezoneSource

`func (o *PublicIndexedTimePoint) SetTimezoneSource(v string)`

SetTimezoneSource sets TimezoneSource field to given value.

### HasTimezoneSource

`func (o *PublicIndexedTimePoint) HasTimezoneSource() bool`

HasTimezoneSource returns a boolean if a field has been set.

### GetZoneId

`func (o *PublicIndexedTimePoint) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *PublicIndexedTimePoint) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *PublicIndexedTimePoint) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.


### GetIndexReference

`func (o *PublicIndexedTimePoint) GetIndexReference() PublicIndexedTimePointIndexReference`

GetIndexReference returns the IndexReference field if non-nil, zero value otherwise.

### GetIndexReferenceOk

`func (o *PublicIndexedTimePoint) GetIndexReferenceOk() (*PublicIndexedTimePointIndexReference, bool)`

GetIndexReferenceOk returns a tuple with the IndexReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexReference

`func (o *PublicIndexedTimePoint) SetIndexReference(v PublicIndexedTimePointIndexReference)`

SetIndexReference sets IndexReference field to given value.


### GetOffset

`func (o *PublicIndexedTimePoint) GetOffset() PublicIndexOffset`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PublicIndexedTimePoint) GetOffsetOk() (*PublicIndexOffset, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PublicIndexedTimePoint) SetOffset(v PublicIndexOffset)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PublicIndexedTimePoint) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


