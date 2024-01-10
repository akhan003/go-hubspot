# PublicComparativeDatePropertyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | **string** |  | [default to "COMPARATIVE_DATE"]
**Operator** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**ComparisonPropertyName** | **string** |  | 
**DefaultComparisonValue** | Pointer to **string** |  | [optional] 

## Methods

### NewPublicComparativeDatePropertyOperation

`func NewPublicComparativeDatePropertyOperation(operationType string, operator string, includeObjectsWithNoValueSet bool, comparisonPropertyName string, ) *PublicComparativeDatePropertyOperation`

NewPublicComparativeDatePropertyOperation instantiates a new PublicComparativeDatePropertyOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicComparativeDatePropertyOperationWithDefaults

`func NewPublicComparativeDatePropertyOperationWithDefaults() *PublicComparativeDatePropertyOperation`

NewPublicComparativeDatePropertyOperationWithDefaults instantiates a new PublicComparativeDatePropertyOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *PublicComparativeDatePropertyOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicComparativeDatePropertyOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicComparativeDatePropertyOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperator

`func (o *PublicComparativeDatePropertyOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicComparativeDatePropertyOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicComparativeDatePropertyOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *PublicComparativeDatePropertyOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicComparativeDatePropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicComparativeDatePropertyOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetComparisonPropertyName

`func (o *PublicComparativeDatePropertyOperation) GetComparisonPropertyName() string`

GetComparisonPropertyName returns the ComparisonPropertyName field if non-nil, zero value otherwise.

### GetComparisonPropertyNameOk

`func (o *PublicComparativeDatePropertyOperation) GetComparisonPropertyNameOk() (*string, bool)`

GetComparisonPropertyNameOk returns a tuple with the ComparisonPropertyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonPropertyName

`func (o *PublicComparativeDatePropertyOperation) SetComparisonPropertyName(v string)`

SetComparisonPropertyName sets ComparisonPropertyName field to given value.


### GetDefaultComparisonValue

`func (o *PublicComparativeDatePropertyOperation) GetDefaultComparisonValue() string`

GetDefaultComparisonValue returns the DefaultComparisonValue field if non-nil, zero value otherwise.

### GetDefaultComparisonValueOk

`func (o *PublicComparativeDatePropertyOperation) GetDefaultComparisonValueOk() (*string, bool)`

GetDefaultComparisonValueOk returns a tuple with the DefaultComparisonValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultComparisonValue

`func (o *PublicComparativeDatePropertyOperation) SetDefaultComparisonValue(v string)`

SetDefaultComparisonValue sets DefaultComparisonValue field to given value.

### HasDefaultComparisonValue

`func (o *PublicComparativeDatePropertyOperation) HasDefaultComparisonValue() bool`

HasDefaultComparisonValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


