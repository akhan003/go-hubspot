# PublicTimePointOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | **string** |  | [default to "TIME_POINT"]
**Operator** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**TimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 
**EndpointBehavior** | Pointer to **string** |  | [optional] 
**PropertyParser** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewPublicTimePointOperation

`func NewPublicTimePointOperation(operationType string, operator string, includeObjectsWithNoValueSet bool, timePoint PublicTimePointOperationTimePoint, type_ string, ) *PublicTimePointOperation`

NewPublicTimePointOperation instantiates a new PublicTimePointOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicTimePointOperationWithDefaults

`func NewPublicTimePointOperationWithDefaults() *PublicTimePointOperation`

NewPublicTimePointOperationWithDefaults instantiates a new PublicTimePointOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *PublicTimePointOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicTimePointOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicTimePointOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperator

`func (o *PublicTimePointOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicTimePointOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicTimePointOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *PublicTimePointOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicTimePointOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicTimePointOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetTimePoint

`func (o *PublicTimePointOperation) GetTimePoint() PublicTimePointOperationTimePoint`

GetTimePoint returns the TimePoint field if non-nil, zero value otherwise.

### GetTimePointOk

`func (o *PublicTimePointOperation) GetTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetTimePointOk returns a tuple with the TimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePoint

`func (o *PublicTimePointOperation) SetTimePoint(v PublicTimePointOperationTimePoint)`

SetTimePoint sets TimePoint field to given value.


### GetEndpointBehavior

`func (o *PublicTimePointOperation) GetEndpointBehavior() string`

GetEndpointBehavior returns the EndpointBehavior field if non-nil, zero value otherwise.

### GetEndpointBehaviorOk

`func (o *PublicTimePointOperation) GetEndpointBehaviorOk() (*string, bool)`

GetEndpointBehaviorOk returns a tuple with the EndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointBehavior

`func (o *PublicTimePointOperation) SetEndpointBehavior(v string)`

SetEndpointBehavior sets EndpointBehavior field to given value.

### HasEndpointBehavior

`func (o *PublicTimePointOperation) HasEndpointBehavior() bool`

HasEndpointBehavior returns a boolean if a field has been set.

### GetPropertyParser

`func (o *PublicTimePointOperation) GetPropertyParser() string`

GetPropertyParser returns the PropertyParser field if non-nil, zero value otherwise.

### GetPropertyParserOk

`func (o *PublicTimePointOperation) GetPropertyParserOk() (*string, bool)`

GetPropertyParserOk returns a tuple with the PropertyParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyParser

`func (o *PublicTimePointOperation) SetPropertyParser(v string)`

SetPropertyParser sets PropertyParser field to given value.

### HasPropertyParser

`func (o *PublicTimePointOperation) HasPropertyParser() bool`

HasPropertyParser returns a boolean if a field has been set.

### GetType

`func (o *PublicTimePointOperation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicTimePointOperation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicTimePointOperation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


