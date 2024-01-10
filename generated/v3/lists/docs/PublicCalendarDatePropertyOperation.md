# PublicCalendarDatePropertyOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationType** | **string** |  | [default to "CALENDAR_DATE"]
**Operator** | **string** |  | 
**IncludeObjectsWithNoValueSet** | **bool** |  | 
**TimeUnit** | **string** |  | 
**FiscalYearStart** | Pointer to **string** |  | [optional] 
**UseFiscalYear** | Pointer to **bool** |  | [optional] 
**TimeUnitCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewPublicCalendarDatePropertyOperation

`func NewPublicCalendarDatePropertyOperation(operationType string, operator string, includeObjectsWithNoValueSet bool, timeUnit string, ) *PublicCalendarDatePropertyOperation`

NewPublicCalendarDatePropertyOperation instantiates a new PublicCalendarDatePropertyOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublicCalendarDatePropertyOperationWithDefaults

`func NewPublicCalendarDatePropertyOperationWithDefaults() *PublicCalendarDatePropertyOperation`

NewPublicCalendarDatePropertyOperationWithDefaults instantiates a new PublicCalendarDatePropertyOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationType

`func (o *PublicCalendarDatePropertyOperation) GetOperationType() string`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *PublicCalendarDatePropertyOperation) GetOperationTypeOk() (*string, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *PublicCalendarDatePropertyOperation) SetOperationType(v string)`

SetOperationType sets OperationType field to given value.


### GetOperator

`func (o *PublicCalendarDatePropertyOperation) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *PublicCalendarDatePropertyOperation) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *PublicCalendarDatePropertyOperation) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetIncludeObjectsWithNoValueSet

`func (o *PublicCalendarDatePropertyOperation) GetIncludeObjectsWithNoValueSet() bool`

GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field if non-nil, zero value otherwise.

### GetIncludeObjectsWithNoValueSetOk

`func (o *PublicCalendarDatePropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool)`

GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeObjectsWithNoValueSet

`func (o *PublicCalendarDatePropertyOperation) SetIncludeObjectsWithNoValueSet(v bool)`

SetIncludeObjectsWithNoValueSet sets IncludeObjectsWithNoValueSet field to given value.


### GetTimeUnit

`func (o *PublicCalendarDatePropertyOperation) GetTimeUnit() string`

GetTimeUnit returns the TimeUnit field if non-nil, zero value otherwise.

### GetTimeUnitOk

`func (o *PublicCalendarDatePropertyOperation) GetTimeUnitOk() (*string, bool)`

GetTimeUnitOk returns a tuple with the TimeUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnit

`func (o *PublicCalendarDatePropertyOperation) SetTimeUnit(v string)`

SetTimeUnit sets TimeUnit field to given value.


### GetFiscalYearStart

`func (o *PublicCalendarDatePropertyOperation) GetFiscalYearStart() string`

GetFiscalYearStart returns the FiscalYearStart field if non-nil, zero value otherwise.

### GetFiscalYearStartOk

`func (o *PublicCalendarDatePropertyOperation) GetFiscalYearStartOk() (*string, bool)`

GetFiscalYearStartOk returns a tuple with the FiscalYearStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearStart

`func (o *PublicCalendarDatePropertyOperation) SetFiscalYearStart(v string)`

SetFiscalYearStart sets FiscalYearStart field to given value.

### HasFiscalYearStart

`func (o *PublicCalendarDatePropertyOperation) HasFiscalYearStart() bool`

HasFiscalYearStart returns a boolean if a field has been set.

### GetUseFiscalYear

`func (o *PublicCalendarDatePropertyOperation) GetUseFiscalYear() bool`

GetUseFiscalYear returns the UseFiscalYear field if non-nil, zero value otherwise.

### GetUseFiscalYearOk

`func (o *PublicCalendarDatePropertyOperation) GetUseFiscalYearOk() (*bool, bool)`

GetUseFiscalYearOk returns a tuple with the UseFiscalYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFiscalYear

`func (o *PublicCalendarDatePropertyOperation) SetUseFiscalYear(v bool)`

SetUseFiscalYear sets UseFiscalYear field to given value.

### HasUseFiscalYear

`func (o *PublicCalendarDatePropertyOperation) HasUseFiscalYear() bool`

HasUseFiscalYear returns a boolean if a field has been set.

### GetTimeUnitCount

`func (o *PublicCalendarDatePropertyOperation) GetTimeUnitCount() int32`

GetTimeUnitCount returns the TimeUnitCount field if non-nil, zero value otherwise.

### GetTimeUnitCountOk

`func (o *PublicCalendarDatePropertyOperation) GetTimeUnitCountOk() (*int32, bool)`

GetTimeUnitCountOk returns a tuple with the TimeUnitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeUnitCount

`func (o *PublicCalendarDatePropertyOperation) SetTimeUnitCount(v int32)`

SetTimeUnitCount sets TimeUnitCount field to given value.

### HasTimeUnitCount

`func (o *PublicCalendarDatePropertyOperation) HasTimeUnitCount() bool`

HasTimeUnitCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


