# PublicPropertyFilterOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | **string** |  | 
**Operator** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**Value** | **string** |  | 
**RequiresTimeZoneConversion** | **bool** |  | 
**Timestamp** | **int32** |  | 
**UpperBound** | **int32** |  | 
**LowerBound** | **int32** |  | 
**ComparisonPropertyName** | **string** |  | 
**DefaultComparisonValue** | Pointer to **string** |  | [optional] 
**NumberOfDays** | **int32** |  | 
**Values** | **[]string** |  | 
**Year** | **int32** |  | 
**Month** | **string** |  | 
**Day** | **int32** |  | 
**TimeUnit** | **string** |  | 
**FiscalYearStart** | Pointer to **string** |  | [optional] 
**UseFiscalYear** | Pointer to **bool** |  | [optional] 
**TimeUnitCount** | Pointer to **int32** |  | [optional] 
**TimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 
**EndpointBehavior** | Pointer to **string** |  | [optional] 
**PropertyParser** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**LowerBoundEndpointBehavior** | Pointer to **string** |  | [optional] 
**UpperBoundEndpointBehavior** | Pointer to **string** |  | [optional] 
**LowerBoundTimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 
**UpperBoundTimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 

## Methods

### NewPublicPropertyFilterOperation

`func NewPublicPropertyFilterOperation(operationType string, operator string, includeObjectsWithNoValueSet bool, value string, requiresTimeZoneConversion bool, timestamp int32, upperBound int32, lowerBound int32, comparisonPropertyName string, numberOfDays int32, values []string, year int32, month string, day int32, timeUnit string, timePoint PublicTimePointOperationTimePoint, type_ string, lowerBoundTimePoint PublicTimePointOperationTimePoint, upperBoundTimePoint PublicTimePointOperationTimePoint, ) *PublicPropertyFilterOperation`

NewPublicPropertyFilterOperation instantiates a new PublicPropertyFilterOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicPropertyFilterOperationWithDefaults

`func NewPublicPropertyFilterOperationWithDefaults() *PublicPropertyFilterOperation`

NewPublicPropertyFilterOperationWithDefaults instantiates a new PublicPropertyFilterOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *PublicPropertyFilterOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicPropertyFilterOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicPropertyFilterOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperator

`func (o *PublicPropertyFilterOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicPropertyFilterOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicPropertyFilterOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *PublicPropertyFilterOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicPropertyFilterOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicPropertyFilterOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetValue

`func (o *PublicPropertyFilterOperation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PublicPropertyFilterOperation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PublicPropertyFilterOperation) SetValue(v string)`

SetValue sets Value field to given value.


### GetRequiresTimeZoneConversion

`func (o *PublicPropertyFilterOperation) GetRequiresTimeZoneConversion() bool`

GetRequiresTimeZoneConversion returns the RequiresTimeZoneConversion field if non-nil, zero value otherwise.

### GetRequiresTimeZoneConversionOk

`func (o *PublicPropertyFilterOperation) GetRequiresTimeZoneConversionOk() (*bool, bool)`

GetRequiresTimeZoneConversionOk returns a tuple with the RequiresTimeZoneConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTimeZoneConversion

`func (o *PublicPropertyFilterOperation) SetRequiresTimeZoneConversion(v bool)`

SetRequiresTimeZoneConversion sets RequiresTimeZoneConversion field to given value.


### GetTimestamp

`func (o *PublicPropertyFilterOperation) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PublicPropertyFilterOperation) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PublicPropertyFilterOperation) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetUpperBound

`func (o *PublicPropertyFilterOperation) GetUpperBound() int32`

GetUpperBound returns the UpperBound field if non-nil, zero value otherwise.

### GetUpperBoundOk

`func (o *PublicPropertyFilterOperation) GetUpperBoundOk() (*int32, bool)`

GetUpperBoundOk returns a tuple with the UpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBound

`func (o *PublicPropertyFilterOperation) SetUpperBound(v int32)`

SetUpperBound sets UpperBound field to given value.


### GetLowerBound

`func (o *PublicPropertyFilterOperation) GetLowerBound() int32`

GetLowerBound returns the LowerBound field if non-nil, zero value otherwise.

### GetLowerBoundOk

`func (o *PublicPropertyFilterOperation) GetLowerBoundOk() (*int32, bool)`

GetLowerBoundOk returns a tuple with the LowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBound

`func (o *PublicPropertyFilterOperation) SetLowerBound(v int32)`

SetLowerBound sets LowerBound field to given value.


### GetComparisonPropertyName

`func (o *PublicPropertyFilterOperation) GetComparisonPropertyName() string`

GetComparisonPropertyName returns the ComparisonPropertyName field if non-nil, zero value otherwise.

### GetComparisonPropertyNameOk

`func (o *PublicPropertyFilterOperation) GetComparisonPropertyNameOk() (*string, bool)`

GetComparisonPropertyNameOk returns a tuple with the ComparisonPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonPropertyName

`func (o *PublicPropertyFilterOperation) SetComparisonPropertyName(v string)`

SetComparisonPropertyName sets ComparisonPropertyName field to given value.


### GetDefaultComparisonValue

`func (o *PublicPropertyFilterOperation) GetDefaultComparisonValue() string`

GetDefaultComparisonValue returns the DefaultComparisonValue field if non-nil, zero value otherwise.

### GetDefaultComparisonValueOk

`func (o *PublicPropertyFilterOperation) GetDefaultComparisonValueOk() (*string, bool)`

GetDefaultComparisonValueOk returns a tuple with the DefaultComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultComparisonValue

`func (o *PublicPropertyFilterOperation) SetDefaultComparisonValue(v string)`

SetDefaultComparisonValue sets DefaultComparisonValue field to given value.

### HasDefaultComparisonValue

`func (o *PublicPropertyFilterOperation) HasDefaultComparisonValue() bool`

HasDefaultComparisonValue returns a boolean if a field has been set.

### GetNumberOfDays

`func (o *PublicPropertyFilterOperation) GetNumberOfDays() int32`

GetNumberOfDays returns the NumberOfDays field if non-nil, zero value otherwise.

### GetNumberOfDaysOk

`func (o *PublicPropertyFilterOperation) GetNumberOfDaysOk() (*int32, bool)`

GetNumberOfDaysOk returns a tuple with the NumberOfDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDays

`func (o *PublicPropertyFilterOperation) SetNumberOfDays(v int32)`

SetNumberOfDays sets NumberOfDays field to given value.


### GetValues

`func (o *PublicPropertyFilterOperation) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PublicPropertyFilterOperation) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PublicPropertyFilterOperation) SetValues(v []string)`

SetValues sets Values field to given value.


### GetYear

`func (o *PublicPropertyFilterOperation) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *PublicPropertyFilterOperation) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *PublicPropertyFilterOperation) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *PublicPropertyFilterOperation) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *PublicPropertyFilterOperation) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *PublicPropertyFilterOperation) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetDay

`func (o *PublicPropertyFilterOperation) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *PublicPropertyFilterOperation) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *PublicPropertyFilterOperation) SetDay(v int32)`

SetDay sets Day field to given value.


### GetTimeUnit

`func (o *PublicPropertyFilterOperation) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *PublicPropertyFilterOperation) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *PublicPropertyFilterOperation) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.


### GetFiscalYearStart

`func (o *PublicPropertyFilterOperation) GetFiscalYearStart() string`

GetFiscalYearStart returns the FiscalYearStart field if non-nil, zero value otherwise.

### GetFiscalYearStartOk

`func (o *PublicPropertyFilterOperation) GetFiscalYearStartOk() (*string, bool)`

GetFiscalYearStartOk returns a tuple with the FiscalYearStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearStart

`func (o *PublicPropertyFilterOperation) SetFiscalYearStart(v string)`

SetFiscalYearStart sets FiscalYearStart field to given value.

### HasFiscalYearStart

`func (o *PublicPropertyFilterOperation) HasFiscalYearStart() bool`

HasFiscalYearStart returns a boolean if a field has been set.

### GetUseFiscalYear

`func (o *PublicPropertyFilterOperation) GetUseFiscalYear() bool`

GetUseFiscalYear returns the UseFiscalYear field if non-nil, zero value otherwise.

### GetUseFiscalYearOk

`func (o *PublicPropertyFilterOperation) GetUseFiscalYearOk() (*bool, bool)`

GetUseFiscalYearOk returns a tuple with the UseFiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFiscalYear

`func (o *PublicPropertyFilterOperation) SetUseFiscalYear(v bool)`

SetUseFiscalYear sets UseFiscalYear field to given value.

### HasUseFiscalYear

`func (o *PublicPropertyFilterOperation) HasUseFiscalYear() bool`

HasUseFiscalYear returns a boolean if a field has been set.

### GetTimeUnitCount

`func (o *PublicPropertyFilterOperation) GetTimeUnitCount() int32`

GetTimeUnitCount returns the TimeUnitCount field if non-nil, zero value otherwise.

### GetTimeUnitCountOk

`func (o *PublicPropertyFilterOperation) GetTimeUnitCountOk() (*int32, bool)`

GetTimeUnitCountOk returns a tuple with the TimeUnitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnitCount

`func (o *PublicPropertyFilterOperation) SetTimeUnitCount(v int32)`

SetTimeUnitCount sets TimeUnitCount field to given value.

### HasTimeUnitCount

`func (o *PublicPropertyFilterOperation) HasTimeUnitCount() bool`

HasTimeUnitCount returns a boolean if a field has been set.

### GetTimePoint

`func (o *PublicPropertyFilterOperation) GetTimePoint() PublicTimePointOperationTimePoint`

GetTimePoint returns the TimePoint field if non-nil, zero value otherwise.

### GetTimePointOk

`func (o *PublicPropertyFilterOperation) GetTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetTimePointOk returns a tuple with the TimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePoint

`func (o *PublicPropertyFilterOperation) SetTimePoint(v PublicTimePointOperationTimePoint)`

SetTimePoint sets TimePoint field to given value.


### GetEndpointBehavior

`func (o *PublicPropertyFilterOperation) GetEndpointBehavior() string`

GetEndpointBehavior returns the EndpointBehavior field if non-nil, zero value otherwise.

### GetEndpointBehaviorOk

`func (o *PublicPropertyFilterOperation) GetEndpointBehaviorOk() (*string, bool)`

GetEndpointBehaviorOk returns a tuple with the EndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointBehavior

`func (o *PublicPropertyFilterOperation) SetEndpointBehavior(v string)`

SetEndpointBehavior sets EndpointBehavior field to given value.

### HasEndpointBehavior

`func (o *PublicPropertyFilterOperation) HasEndpointBehavior() bool`

HasEndpointBehavior returns a boolean if a field has been set.

### GetPropertyParser

`func (o *PublicPropertyFilterOperation) GetPropertyParser() string`

GetPropertyParser returns the PropertyParser field if non-nil, zero value otherwise.

### GetPropertyParserOk

`func (o *PublicPropertyFilterOperation) GetPropertyParserOk() (*string, bool)`

GetPropertyParserOk returns a tuple with the PropertyParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyParser

`func (o *PublicPropertyFilterOperation) SetPropertyParser(v string)`

SetPropertyParser sets PropertyParser field to given value.

### HasPropertyParser

`func (o *PublicPropertyFilterOperation) HasPropertyParser() bool`

HasPropertyParser returns a boolean if a field has been set.

### GetType

`func (o *PublicPropertyFilterOperation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicPropertyFilterOperation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicPropertyFilterOperation) SetType(v string)`

SetType sets Type field to given value.


### GetLowerBoundEndpointBehavior

`func (o *PublicPropertyFilterOperation) GetLowerBoundEndpointBehavior() string`

GetLowerBoundEndpointBehavior returns the LowerBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetLowerBoundEndpointBehaviorOk

`func (o *PublicPropertyFilterOperation) GetLowerBoundEndpointBehaviorOk() (*string, bool)`

GetLowerBoundEndpointBehaviorOk returns a tuple with the LowerBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundEndpointBehavior

`func (o *PublicPropertyFilterOperation) SetLowerBoundEndpointBehavior(v string)`

SetLowerBoundEndpointBehavior sets LowerBoundEndpointBehavior field to given value.

### HasLowerBoundEndpointBehavior

`func (o *PublicPropertyFilterOperation) HasLowerBoundEndpointBehavior() bool`

HasLowerBoundEndpointBehavior returns a boolean if a field has been set.

### GetUpperBoundEndpointBehavior

`func (o *PublicPropertyFilterOperation) GetUpperBoundEndpointBehavior() string`

GetUpperBoundEndpointBehavior returns the UpperBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetUpperBoundEndpointBehaviorOk

`func (o *PublicPropertyFilterOperation) GetUpperBoundEndpointBehaviorOk() (*string, bool)`

GetUpperBoundEndpointBehaviorOk returns a tuple with the UpperBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundEndpointBehavior

`func (o *PublicPropertyFilterOperation) SetUpperBoundEndpointBehavior(v string)`

SetUpperBoundEndpointBehavior sets UpperBoundEndpointBehavior field to given value.

### HasUpperBoundEndpointBehavior

`func (o *PublicPropertyFilterOperation) HasUpperBoundEndpointBehavior() bool`

HasUpperBoundEndpointBehavior returns a boolean if a field has been set.

### GetLowerBoundTimePoint

`func (o *PublicPropertyFilterOperation) GetLowerBoundTimePoint() PublicTimePointOperationTimePoint`

GetLowerBoundTimePoint returns the LowerBoundTimePoint field if non-nil, zero value otherwise.

### GetLowerBoundTimePointOk

`func (o *PublicPropertyFilterOperation) GetLowerBoundTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetLowerBoundTimePointOk returns a tuple with the LowerBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundTimePoint

`func (o *PublicPropertyFilterOperation) SetLowerBoundTimePoint(v PublicTimePointOperationTimePoint)`

SetLowerBoundTimePoint sets LowerBoundTimePoint field to given value.


### GetUpperBoundTimePoint

`func (o *PublicPropertyFilterOperation) GetUpperBoundTimePoint() PublicTimePointOperationTimePoint`

GetUpperBoundTimePoint returns the UpperBoundTimePoint field if non-nil, zero value otherwise.

### GetUpperBoundTimePointOk

`func (o *PublicPropertyFilterOperation) GetUpperBoundTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetUpperBoundTimePointOk returns a tuple with the UpperBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundTimePoint

`func (o *PublicPropertyFilterOperation) SetUpperBoundTimePoint(v PublicTimePointOperationTimePoint)`

SetUpperBoundTimePoint sets UpperBoundTimePoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


