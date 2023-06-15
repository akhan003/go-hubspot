/*
Blog Post endpoints

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tags

import (
	"encoding/json"
)

// checks if the BatchInputTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputTag{}

// BatchInputTag Wrapper for providing an array of blog tags as inputs.
type BatchInputTag struct {
	// Blog tags to input.
	Inputs []Tag `json:"inputs"`
}

// NewBatchInputTag instantiates a new BatchInputTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputTag(inputs []Tag) *BatchInputTag {
	this := BatchInputTag{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputTagWithDefaults instantiates a new BatchInputTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputTagWithDefaults() *BatchInputTag {
	this := BatchInputTag{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputTag) GetInputs() []Tag {
	if o == nil {
		var ret []Tag
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputTag) GetInputsOk() ([]Tag, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputTag) SetInputs(v []Tag) {
	o.Inputs = v
}

func (o BatchInputTag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

type NullableBatchInputTag struct {
	value *BatchInputTag
	isSet bool
}

func (v NullableBatchInputTag) Get() *BatchInputTag {
	return v.value
}

func (v *NullableBatchInputTag) Set(val *BatchInputTag) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputTag) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputTag(val *BatchInputTag) *NullableBatchInputTag {
	return &NullableBatchInputTag{value: val, isSet: true}
}

func (v NullableBatchInputTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}