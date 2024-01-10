# PublicRollingPropertyUpdatedOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | **string** |  | [default to "ROLLING_PROPERTY_UPDATED"]
**Operator** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**NumberOfDays** | **int32** |  | 

## Methods

### NewPublicRollingPropertyUpdatedOperation

`func NewPublicRollingPropertyUpdatedOperation(operationType string, operator string, includeObjectsWithNoValueSet bool, numberOfDays int32, ) *PublicRollingPropertyUpdatedOperation`

NewPublicRollingPropertyUpdatedOperation instantiates a new PublicRollingPropertyUpdatedOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicRollingPropertyUpdatedOperationWithDefaults

`func NewPublicRollingPropertyUpdatedOperationWithDefaults() *PublicRollingPropertyUpdatedOperation`

NewPublicRollingPropertyUpdatedOperationWithDefaults instantiates a new PublicRollingPropertyUpdatedOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *PublicRollingPropertyUpdatedOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicRollingPropertyUpdatedOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicRollingPropertyUpdatedOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperator

`func (o *PublicRollingPropertyUpdatedOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicRollingPropertyUpdatedOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicRollingPropertyUpdatedOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *PublicRollingPropertyUpdatedOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicRollingPropertyUpdatedOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicRollingPropertyUpdatedOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetNumberOfDays

`func (o *PublicRollingPropertyUpdatedOperation) GetNumberOfDays() int32`

GetNumberOfDays returns the NumberOfDays field if non-nil, zero value otherwise.

### GetNumberOfDaysOk

`func (o *PublicRollingPropertyUpdatedOperation) GetNumberOfDaysOk() (*int32, bool)`

GetNumberOfDaysOk returns a tuple with the NumberOfDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDays

`func (o *PublicRollingPropertyUpdatedOperation) SetNumberOfDays(v int32)`

SetNumberOfDays sets NumberOfDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


