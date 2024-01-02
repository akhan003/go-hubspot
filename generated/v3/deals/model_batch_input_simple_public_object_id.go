/*
Deals

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deals

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BatchInputSimplePublicObjectId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputSimplePublicObjectId{}

// BatchInputSimplePublicObjectId struct for BatchInputSimplePublicObjectId
type BatchInputSimplePublicObjectId struct {
	Inputs []SimplePublicObjectId `json:"inputs"`
}

type _BatchInputSimplePublicObjectId BatchInputSimplePublicObjectId

// NewBatchInputSimplePublicObjectId instantiates a new BatchInputSimplePublicObjectId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputSimplePublicObjectId(inputs []SimplePublicObjectId) *BatchInputSimplePublicObjectId {
	this := BatchInputSimplePublicObjectId{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputSimplePublicObjectIdWithDefaults instantiates a new BatchInputSimplePublicObjectId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputSimplePublicObjectIdWithDefaults() *BatchInputSimplePublicObjectId {
	this := BatchInputSimplePublicObjectId{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputSimplePublicObjectId) GetInputs() []SimplePublicObjectId {
	if o == nil {
		var ret []SimplePublicObjectId
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputSimplePublicObjectId) GetInputsOk() ([]SimplePublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputSimplePublicObjectId) SetInputs(v []SimplePublicObjectId) {
	o.Inputs = v
}

func (o BatchInputSimplePublicObjectId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputSimplePublicObjectId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

func (o *BatchInputSimplePublicObjectId) UnmarshalJSON(data []byte) (err error) {
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

	varBatchInputSimplePublicObjectId := _BatchInputSimplePublicObjectId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchInputSimplePublicObjectId)

	if err != nil {
		return err
	}

	*o = BatchInputSimplePublicObjectId(varBatchInputSimplePublicObjectId)

	return err
}

type NullableBatchInputSimplePublicObjectId struct {
	value *BatchInputSimplePublicObjectId
	isSet bool
}

func (v NullableBatchInputSimplePublicObjectId) Get() *BatchInputSimplePublicObjectId {
	return v.value
}

func (v *NullableBatchInputSimplePublicObjectId) Set(val *BatchInputSimplePublicObjectId) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputSimplePublicObjectId) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputSimplePublicObjectId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputSimplePublicObjectId(val *BatchInputSimplePublicObjectId) *NullableBatchInputSimplePublicObjectId {
	return &NullableBatchInputSimplePublicObjectId{value: val, isSet: true}
}

func (v NullableBatchInputSimplePublicObjectId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputSimplePublicObjectId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
