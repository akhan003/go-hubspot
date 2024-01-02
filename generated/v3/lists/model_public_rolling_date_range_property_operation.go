/*
Lists

CRUD operations to manage lists and list memberships

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lists

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PublicRollingDateRangePropertyOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicRollingDateRangePropertyOperation{}

// PublicRollingDateRangePropertyOperation struct for PublicRollingDateRangePropertyOperation
type PublicRollingDateRangePropertyOperation struct {
	OperationType                string `json:"operationType"`
	Operator                     string `json:"operator"`
	IncludeObjectsWithNoValueSet bool   `json:"includeObjectsWithNoValueSet"`
	NumberOfDays                 int32  `json:"numberOfDays"`
	RequiresTimeZoneConversion   bool   `json:"requiresTimeZoneConversion"`
}

type _PublicRollingDateRangePropertyOperation PublicRollingDateRangePropertyOperation

// NewPublicRollingDateRangePropertyOperation instantiates a new PublicRollingDateRangePropertyOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicRollingDateRangePropertyOperation(operationType string, operator string, includeObjectsWithNoValueSet bool, numberOfDays int32, requiresTimeZoneConversion bool) *PublicRollingDateRangePropertyOperation {
	this := PublicRollingDateRangePropertyOperation{}
	this.OperationType = operationType
	this.Operator = operator
	this.IncludeObjectsWithNoValueSet = includeObjectsWithNoValueSet
	this.NumberOfDays = numberOfDays
	this.RequiresTimeZoneConversion = requiresTimeZoneConversion
	return &this
}

// NewPublicRollingDateRangePropertyOperationWithDefaults instantiates a new PublicRollingDateRangePropertyOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicRollingDateRangePropertyOperationWithDefaults() *PublicRollingDateRangePropertyOperation {
	this := PublicRollingDateRangePropertyOperation{}
	var operationType string = "ROLLING_DATE_RANGE"
	this.OperationType = operationType
	return &this
}

// GetOperationType returns the OperationType field value
func (o *PublicRollingDateRangePropertyOperation) GetOperationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationType
}

// GetOperationTypeOk returns a tuple with the OperationType field value
// and a boolean to check if the value has been set.
func (o *PublicRollingDateRangePropertyOperation) GetOperationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationType, true
}

// SetOperationType sets field value
func (o *PublicRollingDateRangePropertyOperation) SetOperationType(v string) {
	o.OperationType = v
}

// GetOperator returns the Operator field value
func (o *PublicRollingDateRangePropertyOperation) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *PublicRollingDateRangePropertyOperation) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *PublicRollingDateRangePropertyOperation) SetOperator(v string) {
	o.Operator = v
}

// GetIncludeObjectsWithNoValueSet returns the IncludeObjectsWithNoValueSet field value
func (o *PublicRollingDateRangePropertyOperation) GetIncludeObjectsWithNoValueSet() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IncludeObjectsWithNoValueSet
}

// GetIncludeObjectsWithNoValueSetOk returns a tuple with the IncludeObjectsWithNoValueSet field value
// and a boolean to check if the value has been set.
func (o *PublicRollingDateRangePropertyOperation) GetIncludeObjectsWithNoValueSetOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludeObjectsWithNoValueSet, true
}

// SetIncludeObjectsWithNoValueSet sets field value
func (o *PublicRollingDateRangePropertyOperation) SetIncludeObjectsWithNoValueSet(v bool) {
	o.IncludeObjectsWithNoValueSet = v
}

// GetNumberOfDays returns the NumberOfDays field value
func (o *PublicRollingDateRangePropertyOperation) GetNumberOfDays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfDays
}

// GetNumberOfDaysOk returns a tuple with the NumberOfDays field value
// and a boolean to check if the value has been set.
func (o *PublicRollingDateRangePropertyOperation) GetNumberOfDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfDays, true
}

// SetNumberOfDays sets field value
func (o *PublicRollingDateRangePropertyOperation) SetNumberOfDays(v int32) {
	o.NumberOfDays = v
}

// GetRequiresTimeZoneConversion returns the RequiresTimeZoneConversion field value
func (o *PublicRollingDateRangePropertyOperation) GetRequiresTimeZoneConversion() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequiresTimeZoneConversion
}

// GetRequiresTimeZoneConversionOk returns a tuple with the RequiresTimeZoneConversion field value
// and a boolean to check if the value has been set.
func (o *PublicRollingDateRangePropertyOperation) GetRequiresTimeZoneConversionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequiresTimeZoneConversion, true
}

// SetRequiresTimeZoneConversion sets field value
func (o *PublicRollingDateRangePropertyOperation) SetRequiresTimeZoneConversion(v bool) {
	o.RequiresTimeZoneConversion = v
}

func (o PublicRollingDateRangePropertyOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicRollingDateRangePropertyOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationType"] = o.OperationType
	toSerialize["operator"] = o.Operator
	toSerialize["includeObjectsWithNoValueSet"] = o.IncludeObjectsWithNoValueSet
	toSerialize["numberOfDays"] = o.NumberOfDays
	toSerialize["requiresTimeZoneConversion"] = o.RequiresTimeZoneConversion
	return toSerialize, nil
}

func (o *PublicRollingDateRangePropertyOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"operationType",
		"operator",
		"includeObjectsWithNoValueSet",
		"numberOfDays",
		"requiresTimeZoneConversion",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPublicRollingDateRangePropertyOperation := _PublicRollingDateRangePropertyOperation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicRollingDateRangePropertyOperation)

	if err != nil {
		return err
	}

	*o = PublicRollingDateRangePropertyOperation(varPublicRollingDateRangePropertyOperation)

	return err
}

type NullablePublicRollingDateRangePropertyOperation struct {
	value *PublicRollingDateRangePropertyOperation
	isSet bool
}

func (v NullablePublicRollingDateRangePropertyOperation) Get() *PublicRollingDateRangePropertyOperation {
	return v.value
}

func (v *NullablePublicRollingDateRangePropertyOperation) Set(val *PublicRollingDateRangePropertyOperation) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicRollingDateRangePropertyOperation) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicRollingDateRangePropertyOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicRollingDateRangePropertyOperation(val *PublicRollingDateRangePropertyOperation) *NullablePublicRollingDateRangePropertyOperation {
	return &NullablePublicRollingDateRangePropertyOperation{value: val, isSet: true}
}

func (v NullablePublicRollingDateRangePropertyOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicRollingDateRangePropertyOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}