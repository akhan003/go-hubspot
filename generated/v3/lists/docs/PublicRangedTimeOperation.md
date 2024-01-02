# PublicRangedTimeOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | [default to "TIME_RANGED"]
**Operator** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**LowerBoundEndpointBehavior** | Pointer to **string** |  | [optional] 
**UpperBoundEndpointBehavior** | Pointer to **string** |  | [optional] 
**PropertyParser** | Pointer to **string** |  | [optional] 
**LowerBoundTimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 
**UpperBoundTimePoint** | [**PublicTimePointOperationTimePoint**](PublicTimePointOperationTimePoint.md) |  | 
**OperationType** | **string** |  | 

## Methods

### NewPublicRangedTimeOperation

`func NewPublicRangedTimeOperation(type_ string, operator string, includeObjectsWithNoValueSet bool, lowerBoundTimePoint PublicTimePointOperationTimePoint, upperBoundTimePoint PublicTimePointOperationTimePoint, operationType string, ) *PublicRangedTimeOperation`

NewPublicRangedTimeOperation instantiates a new PublicRangedTimeOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicRangedTimeOperationWithDefaults

`func NewPublicRangedTimeOperationWithDefaults() *PublicRangedTimeOperation`

NewPublicRangedTimeOperationWithDefaults instantiates a new PublicRangedTimeOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PublicRangedTimeOperation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PublicRangedTimeOperation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PublicRangedTimeOperation) SetType(v string)`

SetType sets Type field to given value.


### GetOperator

`func (o *PublicRangedTimeOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicRangedTimeOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicRangedTimeOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *PublicRangedTimeOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicRangedTimeOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicRangedTimeOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetLowerBoundEndpointBehavior

`func (o *PublicRangedTimeOperation) GetLowerBoundEndpointBehavior() string`

GetLowerBoundEndpointBehavior returns the LowerBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetLowerBoundEndpointBehaviorOk

`func (o *PublicRangedTimeOperation) GetLowerBoundEndpointBehaviorOk() (*string, bool)`

GetLowerBoundEndpointBehaviorOk returns a tuple with the LowerBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundEndpointBehavior

`func (o *PublicRangedTimeOperation) SetLowerBoundEndpointBehavior(v string)`

SetLowerBoundEndpointBehavior sets LowerBoundEndpointBehavior field to given value.

### HasLowerBoundEndpointBehavior

`func (o *PublicRangedTimeOperation) HasLowerBoundEndpointBehavior() bool`

HasLowerBoundEndpointBehavior returns a boolean if a field has been set.

### GetUpperBoundEndpointBehavior

`func (o *PublicRangedTimeOperation) GetUpperBoundEndpointBehavior() string`

GetUpperBoundEndpointBehavior returns the UpperBoundEndpointBehavior field if non-nil, zero value otherwise.

### GetUpperBoundEndpointBehaviorOk

`func (o *PublicRangedTimeOperation) GetUpperBoundEndpointBehaviorOk() (*string, bool)`

GetUpperBoundEndpointBehaviorOk returns a tuple with the UpperBoundEndpointBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundEndpointBehavior

`func (o *PublicRangedTimeOperation) SetUpperBoundEndpointBehavior(v string)`

SetUpperBoundEndpointBehavior sets UpperBoundEndpointBehavior field to given value.

### HasUpperBoundEndpointBehavior

`func (o *PublicRangedTimeOperation) HasUpperBoundEndpointBehavior() bool`

HasUpperBoundEndpointBehavior returns a boolean if a field has been set.

### GetPropertyParser

`func (o *PublicRangedTimeOperation) GetPropertyParser() string`

GetPropertyParser returns the PropertyParser field if non-nil, zero value otherwise.

### GetPropertyParserOk

`func (o *PublicRangedTimeOperation) GetPropertyParserOk() (*string, bool)`

GetPropertyParserOk returns a tuple with the PropertyParser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyParser

`func (o *PublicRangedTimeOperation) SetPropertyParser(v string)`

SetPropertyParser sets PropertyParser field to given value.

### HasPropertyParser

`func (o *PublicRangedTimeOperation) HasPropertyParser() bool`

HasPropertyParser returns a boolean if a field has been set.

### GetLowerBoundTimePoint

`func (o *PublicRangedTimeOperation) GetLowerBoundTimePoint() PublicTimePointOperationTimePoint`

GetLowerBoundTimePoint returns the LowerBoundTimePoint field if non-nil, zero value otherwise.

### GetLowerBoundTimePointOk

`func (o *PublicRangedTimeOperation) GetLowerBoundTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetLowerBoundTimePointOk returns a tuple with the LowerBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerBoundTimePoint

`func (o *PublicRangedTimeOperation) SetLowerBoundTimePoint(v PublicTimePointOperationTimePoint)`

SetLowerBoundTimePoint sets LowerBoundTimePoint field to given value.


### GetUpperBoundTimePoint

`func (o *PublicRangedTimeOperation) GetUpperBoundTimePoint() PublicTimePointOperationTimePoint`

GetUpperBoundTimePoint returns the UpperBoundTimePoint field if non-nil, zero value otherwise.

### GetUpperBoundTimePointOk

`func (o *PublicRangedTimeOperation) GetUpperBoundTimePointOk() (*PublicTimePointOperationTimePoint, bool)`

GetUpperBoundTimePointOk returns a tuple with the UpperBoundTimePoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperBoundTimePoint

`func (o *PublicRangedTimeOperation) SetUpperBoundTimePoint(v PublicTimePointOperationTimePoint)`

SetUpperBoundTimePoint sets UpperBoundTimePoint field to given value.


### GetOperationType

`func (o *PublicRangedTimeOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicRangedTimeOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicRangedTimeOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


