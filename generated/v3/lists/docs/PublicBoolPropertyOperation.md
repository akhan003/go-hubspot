# PublicBoolPropertyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | **string** |  | [default to "BOOL"]
**Operator** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**Value** | **bool** |  | 

## Methods

### NewPublicBoolPropertyOperation

`func NewPublicBoolPropertyOperation(operationType string, operator string, includeObjectsWithNoValueSet bool, value bool, ) *PublicBoolPropertyOperation`

NewPublicBoolPropertyOperation instantiates a new PublicBoolPropertyOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicBoolPropertyOperationWithDefaults

`func NewPublicBoolPropertyOperationWithDefaults() *PublicBoolPropertyOperation`

NewPublicBoolPropertyOperationWithDefaults instantiates a new PublicBoolPropertyOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *PublicBoolPropertyOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicBoolPropertyOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicBoolPropertyOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperator

`func (o *PublicBoolPropertyOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicBoolPropertyOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicBoolPropertyOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *PublicBoolPropertyOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicBoolPropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicBoolPropertyOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetValue

`func (o *PublicBoolPropertyOperation) GetValue() bool`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PublicBoolPropertyOperation) GetValueOk() (*bool, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PublicBoolPropertyOperation) SetValue(v bool)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


