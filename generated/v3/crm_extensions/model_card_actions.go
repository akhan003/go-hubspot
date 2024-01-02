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

// checks if the CardActions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardActions{}

// CardActions Configuration for custom user actions on cards.
type CardActions struct {
	// A list of URL prefixes that will be accepted for card action URLs. If your data fetch response includes an action URL that doesn't begin with one of these values, it will result in an error and the card will not be displayed.
	BaseUrls []string `json:"baseUrls"`
}

type _CardActions CardActions

// NewCardActions instantiates a new CardActions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardActions(baseUrls []string) *CardActions {
	this := CardActions{}
	this.BaseUrls = baseUrls
	return &this
}

// NewCardActionsWithDefaults instantiates a new CardActions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardActionsWithDefaults() *CardActions {
	this := CardActions{}
	return &this
}

// GetBaseUrls returns the BaseUrls field value
func (o *CardActions) GetBaseUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BaseUrls
}

// GetBaseUrlsOk returns a tuple with the BaseUrls field value
// and a boolean to check if the value has been set.
func (o *CardActions) GetBaseUrlsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseUrls, true
}

// SetBaseUrls sets field value
func (o *CardActions) SetBaseUrls(v []string) {
	o.BaseUrls = v
}

func (o CardActions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardActions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["baseUrls"] = o.BaseUrls
	return toSerialize, nil
}

func (o *CardActions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"baseUrls",
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

	varCardActions := _CardActions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCardActions)

	if err != nil {
		return err
	}

	*o = CardActions(varCardActions)

	return err
}

type NullableCardActions struct {
	value *CardActions
	isSet bool
}

func (v NullableCardActions) Get() *CardActions {
	return v.value
}

func (v *NullableCardActions) Set(val *CardActions) {
	v.value = val
	v.isSet = true
}

func (v NullableCardActions) IsSet() bool {
	return v.isSet
}

func (v *NullableCardActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardActions(val *CardActions) *NullableCardActions {
	return &NullableCardActions{value: val, isSet: true}
}

func (v NullableCardActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
