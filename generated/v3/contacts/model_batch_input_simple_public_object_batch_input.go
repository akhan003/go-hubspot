/*
Contacts

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package contacts

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BatchInputSimplePublicObjectBatchInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputSimplePublicObjectBatchInput{}

// BatchInputSimplePublicObjectBatchInput struct for BatchInputSimplePublicObjectBatchInput
type BatchInputSimplePublicObjectBatchInput struct {
	Inputs []SimplePublicObjectBatchInput `json:"inputs"`
}

type _BatchInputSimplePublicObjectBatchInput BatchInputSimplePublicObjectBatchInput

// NewBatchInputSimplePublicObjectBatchInput instantiates a new BatchInputSimplePublicObjectBatchInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputSimplePublicObjectBatchInput(inputs []SimplePublicObjectBatchInput) *BatchInputSimplePublicObjectBatchInput {
	this := BatchInputSimplePublicObjectBatchInput{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputSimplePublicObjectBatchInputWithDefaults instantiates a new BatchInputSimplePublicObjectBatchInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputSimplePublicObjectBatchInputWithDefaults() *BatchInputSimplePublicObjectBatchInput {
	this := BatchInputSimplePublicObjectBatchInput{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputSimplePublicObjectBatchInput) GetInputs() []SimplePublicObjectBatchInput {
	if o == nil {
		var ret []SimplePublicObjectBatchInput
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputSimplePublicObjectBatchInput) GetInputsOk() ([]SimplePublicObjectBatchInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputSimplePublicObjectBatchInput) SetInputs(v []SimplePublicObjectBatchInput) {
	o.Inputs = v
}

func (o BatchInputSimplePublicObjectBatchInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputSimplePublicObjectBatchInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

func (o *BatchInputSimplePublicObjectBatchInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inputs",
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

	varBatchInputSimplePublicObjectBatchInput := _BatchInputSimplePublicObjectBatchInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchInputSimplePublicObjectBatchInput)

	if err != nil {
		return err
	}

	*o = BatchInputSimplePublicObjectBatchInput(varBatchInputSimplePublicObjectBatchInput)

	return err
}

type NullableBatchInputSimplePublicObjectBatchInput struct {
	value *BatchInputSimplePublicObjectBatchInput
	isSet bool
}

func (v NullableBatchInputSimplePublicObjectBatchInput) Get() *BatchInputSimplePublicObjectBatchInput {
	return v.value
}

func (v *NullableBatchInputSimplePublicObjectBatchInput) Set(val *BatchInputSimplePublicObjectBatchInput) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputSimplePublicObjectBatchInput) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputSimplePublicObjectBatchInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputSimplePublicObjectBatchInput(val *BatchInputSimplePublicObjectBatchInput) *NullableBatchInputSimplePublicObjectBatchInput {
	return &NullableBatchInputSimplePublicObjectBatchInput{value: val, isSet: true}
}

func (v NullableBatchInputSimplePublicObjectBatchInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputSimplePublicObjectBatchInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
