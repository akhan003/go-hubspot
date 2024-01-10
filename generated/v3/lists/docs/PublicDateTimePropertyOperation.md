# PublicDateTimePropertyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | **string** |  | [default to "DATETIME"]
**Operator** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**RequiresTimeZoneConversion** | **bool** |  | 
**Timestamp** | **int32** |  | 

## Methods

### NewPublicDateTimePropertyOperation

`func NewPublicDateTimePropertyOperation(operationType string, operator string, includeObjectsWithNoValueSet bool, requiresTimeZoneConversion bool, timestamp int32, ) *PublicDateTimePropertyOperation`

NewPublicDateTimePropertyOperation instantiates a new PublicDateTimePropertyOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicDateTimePropertyOperationWithDefaults

`func NewPublicDateTimePropertyOperationWithDefaults() *PublicDateTimePropertyOperation`

NewPublicDateTimePropertyOperationWithDefaults instantiates a new PublicDateTimePropertyOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *PublicDateTimePropertyOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicDateTimePropertyOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicDateTimePropertyOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperator

`func (o *PublicDateTimePropertyOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicDateTimePropertyOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicDateTimePropertyOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *PublicDateTimePropertyOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicDateTimePropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicDateTimePropertyOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetRequiresTimeZoneConversion

`func (o *PublicDateTimePropertyOperation) GetRequiresTimeZoneConversion() bool`

GetRequiresTimeZoneConversion returns the RequiresTimeZoneConversion field if non-nil, zero value otherwise.

### GetRequiresTimeZoneConversionOk

`func (o *PublicDateTimePropertyOperation) GetRequiresTimeZoneConversionOk() (*bool, bool)`

GetRequiresTimeZoneConversionOk returns a tuple with the RequiresTimeZoneConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTimeZoneConversion

`func (o *PublicDateTimePropertyOperation) SetRequiresTimeZoneConversion(v bool)`

SetRequiresTimeZoneConversion sets RequiresTimeZoneConversion field to given value.


### GetTimestamp

`func (o *PublicDateTimePropertyOperation) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PublicDateTimePropertyOperation) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PublicDateTimePropertyOperation) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


