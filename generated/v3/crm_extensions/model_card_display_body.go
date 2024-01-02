/*
CRM cards

Allows an app to extend the CRM UI by surfacing custom cards in the sidebar of record pages. These cards are defined up-front as part of app configuration, then populated by external data fetch requests when the record page is accessed by a user.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm_extensions

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CardDisplayBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardDisplayBody{}

// CardDisplayBody Configuration for displayed info on a card
type CardDisplayBody struct {
	// Card display properties. These will will be rendered as \"label : value\" pairs in the card UI. See the [example card](#) in the overview docs for more details.
	Properties []CardDisplayProperty `json:"properties"`
}

type _CardDisplayBody CardDisplayBody

// NewCardDisplayBody instantiates a new CardDisplayBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardDisplayBody(properties []CardDisplayProperty) *CardDisplayBody {
	this := CardDisplayBody{}
	this.Properties = properties
	return &this
}

// NewCardDisplayBodyWithDefaults instantiates a new CardDisplayBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardDisplayBodyWithDefaults() *CardDisplayBody {
	this := CardDisplayBody{}
	return &this
}

// GetProperties returns the Properties field value
func (o *CardDisplayBody) GetProperties() []CardDisplayProperty {
	if o == nil {
		var ret []CardDisplayProperty
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *CardDisplayBody) GetPropertiesOk() ([]CardDisplayProperty, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *CardDisplayBody) SetProperties(v []CardDisplayProperty) {
	o.Properties = v
}

func (o CardDisplayBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardDisplayBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

func (o *CardDisplayBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varCardDisplayBody := _CardDisplayBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCardDisplayBody)

	if err != nil {
		return err
	}

	*o = CardDisplayBody(varCardDisplayBody)

	return err
}

type NullableCardDisplayBody struct {
	value *CardDisplayBody
	isSet bool
}

func (v NullableCardDisplayBody) Get() *CardDisplayBody {
	return v.value
}

func (v *NullableCardDisplayBody) Set(val *CardDisplayBody) {
	v.value = val
	v.isSet = true
}

func (v NullableCardDisplayBody) IsSet() bool {
	return v.isSet
}

func (v *NullableCardDisplayBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardDisplayBody(val *CardDisplayBody) *NullableCardDisplayBody {
	return &NullableCardDisplayBody{value: val, isSet: true}
}

func (v NullableCardDisplayBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardDisplayBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
