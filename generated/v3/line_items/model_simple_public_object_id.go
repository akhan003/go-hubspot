/*
Line Items

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package line_items

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SimplePublicObjectId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplePublicObjectId{}

// SimplePublicObjectId struct for SimplePublicObjectId
type SimplePublicObjectId struct {
	Id string `json:"id"`
}

type _SimplePublicObjectId SimplePublicObjectId

// NewSimplePublicObjectId instantiates a new SimplePublicObjectId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplePublicObjectId(id string) *SimplePublicObjectId {
	this := SimplePublicObjectId{}
	this.Id = id
	return &this
}

// NewSimplePublicObjectIdWithDefaults instantiates a new SimplePublicObjectId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplePublicObjectIdWithDefaults() *SimplePublicObjectId {
	this := SimplePublicObjectId{}
	return &this
}

// GetId returns the Id field value
func (o *SimplePublicObjectId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimplePublicObjectId) SetId(v string) {
	o.Id = v
}

func (o SimplePublicObjectId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplePublicObjectId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *SimplePublicObjectId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varSimplePublicObjectId := _SimplePublicObjectId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSimplePublicObjectId)

	if err != nil {
		return err
	}

	*o = SimplePublicObjectId(varSimplePublicObjectId)

	return err
}

type NullableSimplePublicObjectId struct {
	value *SimplePublicObjectId
	isSet bool
}

func (v NullableSimplePublicObjectId) Get() *SimplePublicObjectId {
	return v.value
}

func (v *NullableSimplePublicObjectId) Set(val *SimplePublicObjectId) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplePublicObjectId) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplePublicObjectId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplePublicObjectId(val *SimplePublicObjectId) *NullableSimplePublicObjectId {
	return &NullableSimplePublicObjectId{value: val, isSet: true}
}

func (v NullableSimplePublicObjectId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplePublicObjectId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
