/*
Posts

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blog_posts

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BatchInputString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputString{}

// BatchInputString Wrapper for providing an array of strings as inputs.
type BatchInputString struct {
	// Strings to input.
	Inputs []string `json:"inputs"`
}

type _BatchInputString BatchInputString

// NewBatchInputString instantiates a new BatchInputString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputString(inputs []string) *BatchInputString {
	this := BatchInputString{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputStringWithDefaults instantiates a new BatchInputString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputStringWithDefaults() *BatchInputString {
	this := BatchInputString{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputString) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputString) GetInputsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputString) SetInputs(v []string) {
	o.Inputs = v
}

func (o BatchInputString) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

func (o *BatchInputString) UnmarshalJSON(data []byte) (err error) {
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

	varBatchInputString := _BatchInputString{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchInputString)

	if err != nil {
		return err
	}

	*o = BatchInputString(varBatchInputString)

	return err
}

type NullableBatchInputString struct {
	value *BatchInputString
	isSet bool
}

func (v NullableBatchInputString) Get() *BatchInputString {
	return v.value
}

func (v *NullableBatchInputString) Set(val *BatchInputString) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputString) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputString(val *BatchInputString) *NullableBatchInputString {
	return &NullableBatchInputString{value: val, isSet: true}
}

func (v NullableBatchInputString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
