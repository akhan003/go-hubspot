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

// checks if the DetachFromLangGroupRequestVNext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetachFromLangGroupRequestVNext{}

// DetachFromLangGroupRequestVNext Request body object for detaching objects from multi-language groups.
type DetachFromLangGroupRequestVNext struct {
	// ID of the object to remove from a multi-language group.
	Id string `json:"id"`
}

type _DetachFromLangGroupRequestVNext DetachFromLangGroupRequestVNext

// NewDetachFromLangGroupRequestVNext instantiates a new DetachFromLangGroupRequestVNext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetachFromLangGroupRequestVNext(id string) *DetachFromLangGroupRequestVNext {
	this := DetachFromLangGroupRequestVNext{}
	this.Id = id
	return &this
}

// NewDetachFromLangGroupRequestVNextWithDefaults instantiates a new DetachFromLangGroupRequestVNext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetachFromLangGroupRequestVNextWithDefaults() *DetachFromLangGroupRequestVNext {
	this := DetachFromLangGroupRequestVNext{}
	return &this
}

// GetId returns the Id field value
func (o *DetachFromLangGroupRequestVNext) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DetachFromLangGroupRequestVNext) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DetachFromLangGroupRequestVNext) SetId(v string) {
	o.Id = v
}

func (o DetachFromLangGroupRequestVNext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetachFromLangGroupRequestVNext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *DetachFromLangGroupRequestVNext) UnmarshalJSON(data []byte) (err error) {
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

	varDetachFromLangGroupRequestVNext := _DetachFromLangGroupRequestVNext{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDetachFromLangGroupRequestVNext)

	if err != nil {
		return err
	}

	*o = DetachFromLangGroupRequestVNext(varDetachFromLangGroupRequestVNext)

	return err
}

type NullableDetachFromLangGroupRequestVNext struct {
	value *DetachFromLangGroupRequestVNext
	isSet bool
}

func (v NullableDetachFromLangGroupRequestVNext) Get() *DetachFromLangGroupRequestVNext {
	return v.value
}

func (v *NullableDetachFromLangGroupRequestVNext) Set(val *DetachFromLangGroupRequestVNext) {
	v.value = val
	v.isSet = true
}

func (v NullableDetachFromLangGroupRequestVNext) IsSet() bool {
	return v.isSet
}

func (v *NullableDetachFromLangGroupRequestVNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetachFromLangGroupRequestVNext(val *DetachFromLangGroupRequestVNext) *NullableDetachFromLangGroupRequestVNext {
	return &NullableDetachFromLangGroupRequestVNext{value: val, isSet: true}
}

func (v NullableDetachFromLangGroupRequestVNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetachFromLangGroupRequestVNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
