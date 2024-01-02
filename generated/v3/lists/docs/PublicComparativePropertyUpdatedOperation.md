# PublicComparativePropertyUpdatedOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | **string** |  | [default to "COMPARATIVE_PROPERTY_UPDATED"]
**Operator** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**ComparisonPropertyName** | **string** |  | 
**DefaultComparisonValue** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicComparativePropertyUpdatedOperation

`func NewPublicComparativePropertyUpdatedOperation(operationType string, operator string, includeObjectsWithNoValueSet bool, comparisonPropertyName string, ) *PublicComparativePropertyUpdatedOperation`

NewPublicComparativePropertyUpdatedOperation instantiates a new PublicComparativePropertyUpdatedOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicComparativePropertyUpdatedOperationWithDefaults

`func NewPublicComparativePropertyUpdatedOperationWithDefaults() *PublicComparativePropertyUpdatedOperation`

NewPublicComparativePropertyUpdatedOperationWithDefaults instantiates a new PublicComparativePropertyUpdatedOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *PublicComparativePropertyUpdatedOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicComparativePropertyUpdatedOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicComparativePropertyUpdatedOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperator

`func (o *PublicComparativePropertyUpdatedOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicComparativePropertyUpdatedOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicComparativePropertyUpdatedOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *PublicComparativePropertyUpdatedOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicComparativePropertyUpdatedOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicComparativePropertyUpdatedOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetComparisonPropertyName

`func (o *PublicComparativePropertyUpdatedOperation) GetComparisonPropertyName() string`

GetComparisonPropertyName returns the ComparisonPropertyName field if non-nil, zero value otherwise.

### GetComparisonPropertyNameOk

`func (o *PublicComparativePropertyUpdatedOperation) GetComparisonPropertyNameOk() (*string, bool)`

GetComparisonPropertyNameOk returns a tuple with the ComparisonPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonPropertyName

`func (o *PublicComparativePropertyUpdatedOperation) SetComparisonPropertyName(v string)`

SetComparisonPropertyName sets ComparisonPropertyName field to given value.


### GetDefaultComparisonValue

`func (o *PublicComparativePropertyUpdatedOperation) GetDefaultComparisonValue() string`

GetDefaultComparisonValue returns the DefaultComparisonValue field if non-nil, zero value otherwise.

### GetDefaultComparisonValueOk

`func (o *PublicComparativePropertyUpdatedOperation) GetDefaultComparisonValueOk() (*string, bool)`

GetDefaultComparisonValueOk returns a tuple with the DefaultComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultComparisonValue

`func (o *PublicComparativePropertyUpdatedOperation) SetDefaultComparisonValue(v string)`

SetDefaultComparisonValue sets DefaultComparisonValue field to given value.

### HasDefaultComparisonValue

`func (o *PublicComparativePropertyUpdatedOperation) HasDefaultComparisonValue() bool`

HasDefaultComparisonValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


