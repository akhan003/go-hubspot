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

// checks if the BatchReadInputSimplePublicObjectId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchReadInputSimplePublicObjectId{}

// BatchReadInputSimplePublicObjectId struct for BatchReadInputSimplePublicObjectId
type BatchReadInputSimplePublicObjectId struct {
	PropertiesWithHistory []string               `json:"propertiesWithHistory"`
	IdProperty            *string                `json:"idProperty,omitempty"`
	Inputs                []SimplePublicObjectId `json:"inputs"`
	Properties            []string               `json:"properties"`
}

type _BatchReadInputSimplePublicObjectId BatchReadInputSimplePublicObjectId

// NewBatchReadInputSimplePublicObjectId instantiates a new BatchReadInputSimplePublicObjectId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchReadInputSimplePublicObjectId(propertiesWithHistory []string, inputs []SimplePublicObjectId, properties []string) *BatchReadInputSimplePublicObjectId {
	this := BatchReadInputSimplePublicObjectId{}
	this.PropertiesWithHistory = propertiesWithHistory
	this.Inputs = inputs
	this.Properties = properties
	return &this
}

// NewBatchReadInputSimplePublicObjectIdWithDefaults instantiates a new BatchReadInputSimplePublicObjectId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchReadInputSimplePublicObjectIdWithDefaults() *BatchReadInputSimplePublicObjectId {
	this := BatchReadInputSimplePublicObjectId{}
	return &this
}

// GetPropertiesWithHistory returns the PropertiesWithHistory field value
func (o *BatchReadInputSimplePublicObjectId) GetPropertiesWithHistory() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PropertiesWithHistory
}

// GetPropertiesWithHistoryOk returns a tuple with the PropertiesWithHistory field value
// and a boolean to check if the value has been set.
func (o *BatchReadInputSimplePublicObjectId) GetPropertiesWithHistoryOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PropertiesWithHistory, true
}

// SetPropertiesWithHistory sets field value
func (o *BatchReadInputSimplePublicObjectId) SetPropertiesWithHistory(v []string) {
	o.PropertiesWithHistory = v
}

// GetIdProperty returns the IdProperty field value if set, zero value otherwise.
func (o *BatchReadInputSimplePublicObjectId) GetIdProperty() string {
	if o == nil || IsNil(o.IdProperty) {
		var ret string
		return ret
	}
	return *o.IdProperty
}

// GetIdPropertyOk returns a tuple with the IdProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchReadInputSimplePublicObjectId) GetIdPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.IdProperty) {
		return nil, false
	}
	return o.IdProperty, true
}

// HasIdProperty returns a boolean if a field has been set.
func (o *BatchReadInputSimplePublicObjectId) HasIdProperty() bool {
	if o != nil && !IsNil(o.IdProperty) {
		return true
	}

	return false
}

// SetIdProperty gets a reference to the given string and assigns it to the IdProperty field.
func (o *BatchReadInputSimplePublicObjectId) SetIdProperty(v string) {
	o.IdProperty = &v
}

// GetInputs returns the Inputs field value
func (o *BatchReadInputSimplePublicObjectId) GetInputs() []SimplePublicObjectId {
	if o == nil {
		var ret []SimplePublicObjectId
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchReadInputSimplePublicObjectId) GetInputsOk() ([]SimplePublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchReadInputSimplePublicObjectId) SetInputs(v []SimplePublicObjectId) {
	o.Inputs = v
}

// GetProperties returns the Properties field value
func (o *BatchReadInputSimplePublicObjectId) GetProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *BatchReadInputSimplePublicObjectId) GetPropertiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *BatchReadInputSimplePublicObjectId) SetProperties(v []string) {
	o.Properties = v
}

func (o BatchReadInputSimplePublicObjectId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchReadInputSimplePublicObjectId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["propertiesWithHistory"] = o.PropertiesWithHistory
	if !IsNil(o.IdProperty) {
		toSerialize["idProperty"] = o.IdProperty
	}
	toSerialize["inputs"] = o.Inputs
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

func (o *BatchReadInputSimplePublicObjectId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"propertiesWithHistory",
		"inputs",
		"properties",
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

	varBatchReadInputSimplePublicObjectId := _BatchReadInputSimplePublicObjectId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchReadInputSimplePublicObjectId)

	if err != nil {
		return err
	}

	*o = BatchReadInputSimplePublicObjectId(varBatchReadInputSimplePublicObjectId)

	return err
}

type NullableBatchReadInputSimplePublicObjectId struct {
	value *BatchReadInputSimplePublicObjectId
	isSet bool
}

func (v NullableBatchReadInputSimplePublicObjectId) Get() *BatchReadInputSimplePublicObjectId {
	return v.value
}

func (v *NullableBatchReadInputSimplePublicObjectId) Set(val *BatchReadInputSimplePublicObjectId) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchReadInputSimplePublicObjectId) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchReadInputSimplePublicObjectId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchReadInputSimplePublicObjectId(val *BatchReadInputSimplePublicObjectId) *NullableBatchReadInputSimplePublicObjectId {
	return &NullableBatchReadInputSimplePublicObjectId{value: val, isSet: true}
}

func (v NullableBatchReadInputSimplePublicObjectId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchReadInputSimplePublicObjectId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
