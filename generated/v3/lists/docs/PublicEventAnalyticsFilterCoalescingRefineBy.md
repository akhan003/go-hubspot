# PublicEventAnalyticsFilterCoalescingRefineBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**MinOccurrences** | Pointer to **int32** |  | [optional] 
**MaxOccurrences** | Pointer to **int32** |  | [optional] 
**SetType_** | **string** |  | 
**Comparison** | **string** |  | 
**TimeOffset** | [**PublicTimeOffset**](PublicTimeOffset.md) |  | 
**RangeType** | **string** |  | 
**UpperBoundOffset** | [**PublicTimeOffset**](PublicTimeOffset.md) |  | 
**LowerBoundOffset** | [**PublicTimeOffset**](PublicTimeOffset.md) |  | 
**Timestamp** | **int64** |  | 
**LowerTimestamp** | **int64** |  | 
**UpperTimestamp** | **int64** |  | 
**OperationType** | **string** |  | 
**Operator** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**TimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 
**EndpointBehavior** | Pointer to **string** |  | [optional] 
**PropertyParser** | Pointer to **string** |  | [optional] 
**LowerBoundEndpointBehavior** | Pointer to **string** |  | [optional] 
**UpperBoundEndpointBehavior** | Pointer to **string** |  | [optional] 
**LowerBoundTimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 
**UpperBoundTimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 

## Methods

### NewPublicEventAnalyticsFilterCoalescingRefineBy

`func NewPublicEventAnalyticsFilterCoalescingRefineBy(type_ string, setType string, comparison string, timeOffset PublicTimeOffset, rangeType string, upperBoundOffset PublicTimeOffset, lowerBoundOffset PublicTimeOffset, timestamp int64, lowerTimestamp int64, upperTimestamp int64, operationType string, operator string, includeObjectsWithNoValueSet bool, timePoint PublicTimePointOperationTimePoint, lowerBoundTimePoint PublicTimePointOperationTimePoint, upperBoundTimePoint PublicTimePointOperationTimePoint, ) *PublicEventAnalyticsFilterCoalescingRefineBy`

NewPublicEventAnalyticsFilterCoalescingRefineBy instantiates a new PublicEventAnalyticsFilterCoalescingRefineBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicEventAnalyticsFilterCoalescingRefineByWithDefaults

`func NewPublicEventAnalyticsFilterCoalescingRefineByWithDefaults() *PublicEventAnalyticsFilterCoalescingRefineBy`

NewPublicEventAnalyticsFilterCoalescingRefineByWithDefaults instantiates a new PublicEventAnalyticsFilterCoalescingRefineBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetType(v string)`

SetType sets Type field to given value.


### GetMinOccurrences

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetMinOccurrences() int32`

GetMinOccurrences returns the MinOccurrences field if non-nil, zero value otherwise.

### GetMinOccurrencesOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetMinOccurrencesOk() (*int32, bool)`

GetMinOccurrencesOk returns a tuple with the MinOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOccurrences

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetMinOccurrences(v int32)`

SetMinOccurrences sets MinOccurrences field to given value.

### HasMinOccurrences

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) HasMinOccurrences() bool`

HasMinOccurrences returns a boolean if a field has been set.

### GetMaxOccurrences

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetMaxOccurrences() int32`

GetMaxOccurrences returns the MaxOccurrences field if non-nil, zero value otherwise.

### GetMaxOccurrencesOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetMaxOccurrencesOk() (*int32, bool)`

GetMaxOccurrencesOk returns a tuple with the MaxOccurrences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOccurrences

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetMaxOccurrences(v int32)`

SetMaxOccurrences sets MaxOccurrences field to given value.

### HasMaxOccurrences

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) HasMaxOccurrences() bool`

HasMaxOccurrences returns a boolean if a field has been set.

### GetSetType_

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetSetType_() string`

GetSetType_ returns the SetType_ field if non-nil, zero value otherwise.

### GetSetType_Ok

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetSetType_Ok() (*string, bool)`

GetSetType_Ok returns a tuple with the SetType_ field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetType_

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetSetType_(v string)`

SetSetType_ sets SetType_ field to given value.


### GetComparison

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetComparison() string`

GetComparison returns the Comparison field if non-nil, zero value otherwise.

### GetComparisonOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetComparisonOk() (*string, bool)`

GetComparisonOk returns a tuple with the Comparison field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparison

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetComparison(v string)`

SetComparison sets Comparison field to given value.


### GetTimeOffset

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetTimeOffset() PublicTimeOffset`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetTimeOffsetOk() (*PublicTimeOffset, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffset

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetTimeOffset(v PublicTimeOffset)`

SetTimeOffset sets TimeOffset field to given value.


### GetRangeType

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetRangeType() string`

GetRangeType returns the RangeType field if non-nil, zero value otherwise.

### GetRangeTypeOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetRangeTypeOk() (*string, bool)`

GetRangeTypeOk returns a tuple with the RangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeType

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetRangeType(v string)`

SetRangeType sets RangeType field to given value.


### GetUpperBoundOffset

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetUpperBoundOffset() PublicTimeOffset`

GetUpperBoundOffset returns the UpperBoundOffset field if non-nil, zero value otherwise.

### GetUpperBoundOffsetOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetUpperBoundOffsetOk() (*PublicTimeOffset, bool)`

GetUpperBoundOffsetOk returns a tuple with the UpperBoundOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundOffset

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetUpperBoundOffset(v PublicTimeOffset)`

SetUpperBoundOffset sets UpperBoundOffset field to given value.


### GetLowerBoundOffset

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetLowerBoundOffset() PublicTimeOffset`

GetLowerBoundOffset returns the LowerBoundOffset field if non-nil, zero value otherwise.

### GetLowerBoundOffsetOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetLowerBoundOffsetOk() (*PublicTimeOffset, bool)`

GetLowerBoundOffsetOk returns a tuple with the LowerBoundOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundOffset

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetLowerBoundOffset(v PublicTimeOffset)`

SetLowerBoundOffset sets LowerBoundOffset field to given value.


### GetTimestamp

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetLowerTimestamp

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetLowerTimestamp() int64`

GetLowerTimestamp returns the LowerTimestamp field if non-nil, zero value otherwise.

### GetLowerTimestampOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetLowerTimestampOk() (*int64, bool)`

GetLowerTimestampOk returns a tuple with the LowerTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerTimestamp

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetLowerTimestamp(v int64)`

SetLowerTimestamp sets LowerTimestamp field to given value.


### GetUpperTimestamp

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetUpperTimestamp() int64`

GetUpperTimestamp returns the UpperTimestamp field if non-nil, zero value otherwise.

### GetUpperTimestampOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetUpperTimestampOk() (*int64, bool)`

GetUpperTimestampOk returns a tuple with the UpperTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperTimestamp

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetUpperTimestamp(v int64)`

SetUpperTimestamp sets UpperTimestamp field to given value.


### GetOperationType

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperator

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetTimePoint

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetTimePoint() PublicTimePointOperationTimePoint`

GetTimePoint returns the TimePoint field if non-nil, zero value otherwise.

### GetTimePointOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetTimePointOk returns a tuple with the TimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePoint

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetTimePoint(v PublicTimePointOperationTimePoint)`

SetTimePoint sets TimePoint field to given value.


### GetEndpointBehavior

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetEndpointBehavior() string`

GetEndpointBehavior returns the EndpointBehavior field if non-nil, zero value otherwise.

### GetEndpointBehaviorOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetEndpointBehaviorOk() (*string, bool)`

GetEndpointBehaviorOk returns a tuple with the EndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointBehavior

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetEndpointBehavior(v string)`

SetEndpointBehavior sets EndpointBehavior field to given value.

### HasEndpointBehavior

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) HasEndpointBehavior() bool`

HasEndpointBehavior returns a boolean if a field has been set.

### GetPropertyParser

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetPropertyParser() string`

GetPropertyParser returns the PropertyParser field if non-nil, zero value otherwise.

### GetPropertyParserOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetPropertyParserOk() (*string, bool)`

GetPropertyParserOk returns a tuple with the PropertyParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyParser

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetPropertyParser(v string)`

SetPropertyParser sets PropertyParser field to given value.

### HasPropertyParser

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) HasPropertyParser() bool`

HasPropertyParser returns a boolean if a field has been set.

### GetLowerBoundEndpointBehavior

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetLowerBoundEndpointBehavior() string`

GetLowerBoundEndpointBehavior returns the LowerBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetLowerBoundEndpointBehaviorOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetLowerBoundEndpointBehaviorOk() (*string, bool)`

GetLowerBoundEndpointBehaviorOk returns a tuple with the LowerBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundEndpointBehavior

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetLowerBoundEndpointBehavior(v string)`

SetLowerBoundEndpointBehavior sets LowerBoundEndpointBehavior field to given value.

### HasLowerBoundEndpointBehavior

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) HasLowerBoundEndpointBehavior() bool`

HasLowerBoundEndpointBehavior returns a boolean if a field has been set.

### GetUpperBoundEndpointBehavior

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetUpperBoundEndpointBehavior() string`

GetUpperBoundEndpointBehavior returns the UpperBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetUpperBoundEndpointBehaviorOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetUpperBoundEndpointBehaviorOk() (*string, bool)`

GetUpperBoundEndpointBehaviorOk returns a tuple with the UpperBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundEndpointBehavior

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetUpperBoundEndpointBehavior(v string)`

SetUpperBoundEndpointBehavior sets UpperBoundEndpointBehavior field to given value.

### HasUpperBoundEndpointBehavior

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) HasUpperBoundEndpointBehavior() bool`

HasUpperBoundEndpointBehavior returns a boolean if a field has been set.

### GetLowerBoundTimePoint

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetLowerBoundTimePoint() PublicTimePointOperationTimePoint`

GetLowerBoundTimePoint returns the LowerBoundTimePoint field if non-nil, zero value otherwise.

### GetLowerBoundTimePointOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetLowerBoundTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetLowerBoundTimePointOk returns a tuple with the LowerBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundTimePoint

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetLowerBoundTimePoint(v PublicTimePointOperationTimePoint)`

SetLowerBoundTimePoint sets LowerBoundTimePoint field to given value.


### GetUpperBoundTimePoint

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetUpperBoundTimePoint() PublicTimePointOperationTimePoint`

GetUpperBoundTimePoint returns the UpperBoundTimePoint field if non-nil, zero value otherwise.

### GetUpperBoundTimePointOk

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) GetUpperBoundTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetUpperBoundTimePointOk returns a tuple with the UpperBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundTimePoint

`func (o *PublicEventAnalyticsFilterCoalescingRefineBy) SetUpperBoundTimePoint(v PublicTimePointOperationTimePoint)`

SetUpperBoundTimePoint sets UpperBoundTimePoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


