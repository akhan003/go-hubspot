# PublicRangedDatePropertyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | **string** |  | [default to "RANGED_DATE"]
**Operator** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**UpperBound** | **int32** |  | 
**LowerBound** | **int32** |  | 
**RequiresTimeZoneConversion** | **bool** |  | 

## Methods

### NewPublicRangedDatePropertyOperation

`func NewPublicRangedDatePropertyOperation(operationType string, operator string, includeObjectsWithNoValueSet bool, upperBound int32, lowerBound int32, requiresTimeZoneConversion bool, ) *PublicRangedDatePropertyOperation`

NewPublicRangedDatePropertyOperation instantiates a new PublicRangedDatePropertyOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicRangedDatePropertyOperationWithDefaults

`func NewPublicRangedDatePropertyOperationWithDefaults() *PublicRangedDatePropertyOperation`

NewPublicRangedDatePropertyOperationWithDefaults instantiates a new PublicRangedDatePropertyOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *PublicRangedDatePropertyOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicRangedDatePropertyOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicRangedDatePropertyOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperator

`func (o *PublicRangedDatePropertyOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicRangedDatePropertyOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicRangedDatePropertyOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *PublicRangedDatePropertyOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicRangedDatePropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicRangedDatePropertyOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetUpperBound

`func (o *PublicRangedDatePropertyOperation) GetUpperBound() int32`

GetUpperBound returns the UpperBound field if non-nil, zero value otherwise.

### GetUpperBoundOk

`func (o *PublicRangedDatePropertyOperation) GetUpperBoundOk() (*int32, bool)`

GetUpperBoundOk returns a tuple with the UpperBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBound

`func (o *PublicRangedDatePropertyOperation) SetUpperBound(v int32)`

SetUpperBound sets UpperBound field to given value.


### GetLowerBound

`func (o *PublicRangedDatePropertyOperation) GetLowerBound() int32`

GetLowerBound returns the LowerBound field if non-nil, zero value otherwise.

### GetLowerBoundOk

`func (o *PublicRangedDatePropertyOperation) GetLowerBoundOk() (*int32, bool)`

GetLowerBoundOk returns a tuple with the LowerBound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBound

`func (o *PublicRangedDatePropertyOperation) SetLowerBound(v int32)`

SetLowerBound sets LowerBound field to given value.


### GetRequiresTimeZoneConversion

`func (o *PublicRangedDatePropertyOperation) GetRequiresTimeZoneConversion() bool`

GetRequiresTimeZoneConversion returns the RequiresTimeZoneConversion field if non-nil, zero value otherwise.

### GetRequiresTimeZoneConversionOk

`func (o *PublicRangedDatePropertyOperation) GetRequiresTimeZoneConversionOk() (*bool, bool)`

GetRequiresTimeZoneConversionOk returns a tuple with the RequiresTimeZoneConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresTimeZoneConversion

`func (o *PublicRangedDatePropertyOperation) SetRequiresTimeZoneConversion(v bool)`

SetRequiresTimeZoneConversion sets RequiresTimeZoneConversion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


