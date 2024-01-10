# PublicRollingDateRangePropertyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | **string** |  | [default to "ROLLING_DATE_RANGE"]
**Operator** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**NumberOfDays** | **int32** |  | 
**RequiresTimeZoneConversion** | **bool** |  | 

## Methods

### NewPublicRollingDateRangePropertyOperation

`func NewPublicRollingDateRangePropertyOperation(operationType string, operator string, includeObjectsWithNoValueSet bool, numberOfDays int32, requiresTimeZoneConversion bool, ) *PublicRollingDateRangePropertyOperation`

NewPublicRollingDateRangePropertyOperation instantiates a new PublicRollingDateRangePropertyOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicRollingDateRangePropertyOperationWithDefaults

`func NewPublicRollingDateRangePropertyOperationWithDefaults() *PublicRollingDateRangePropertyOperation`

NewPublicRollingDateRangePropertyOperationWithDefaults instantiates a new PublicRollingDateRangePropertyOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *PublicRollingDateRangePropertyOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicRollingDateRangePropertyOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicRollingDateRangePropertyOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperator

`func (o *PublicRollingDateRangePropertyOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicRollingDateRangePropertyOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicRollingDateRangePropertyOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *PublicRollingDateRangePropertyOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicRollingDateRangePropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicRollingDateRangePropertyOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetNumberOfDays

`func (o *PublicRollingDateRangePropertyOperation) GetNumberOfDays() int32`

GetNumberOfDays returns the NumberOfDays field if non-nil, zero value otherwise.

### GetNumberOfDaysOk

`func (o *PublicRollingDateRangePropertyOperation) GetNumberOfDaysOk() (*int32, bool)`

GetNumberOfDaysOk returns a tuple with the NumberOfDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDays

`func (o *PublicRollingDateRangePropertyOperation) SetNumberOfDays(v int32)`

SetNumberOfDays sets NumberOfDays field to given value.


### GetRequiresTimeZoneConversion

`func (o *PublicRollingDateRangePropertyOperation) GetRequiresTimeZoneConversion() bool`

GetRequiresTimeZoneConversion returns the RequiresTimeZoneConversion field if non-nil, zero value otherwise.

### GetRequiresTimeZoneConversionOk

`func (o *PublicRollingDateRangePropertyOperation) GetRequiresTimeZoneConversionOk() (*bool, bool)`

GetRequiresTimeZoneConversionOk returns a tuple with the RequiresTimeZoneConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTimeZoneConversion

`func (o *PublicRollingDateRangePropertyOperation) SetRequiresTimeZoneConversion(v bool)`

SetRequiresTimeZoneConversion sets RequiresTimeZoneConversion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


