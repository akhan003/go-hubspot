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

// checks if the ActionConfirmationBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionConfirmationBody{}

// ActionConfirmationBody struct for ActionConfirmationBody
type ActionConfirmationBody struct {
	Prompt             string `json:"prompt"`
	ConfirmButtonLabel string `json:"confirmButtonLabel"`
	CancelButtonLabel  string `json:"cancelButtonLabel"`
}

type _ActionConfirmationBody ActionConfirmationBody

// NewActionConfirmationBody instantiates a new ActionConfirmationBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionConfirmationBody(prompt string, confirmButtonLabel string, cancelButtonLabel string) *ActionConfirmationBody {
	this := ActionConfirmationBody{}
	this.Prompt = prompt
	this.ConfirmButtonLabel = confirmButtonLabel
	this.CancelButtonLabel = cancelButtonLabel
	return &this
}

// NewActionConfirmationBodyWithDefaults instantiates a new ActionConfirmationBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionConfirmationBodyWithDefaults() *ActionConfirmationBody {
	this := ActionConfirmationBody{}
	return &this
}

// GetPrompt returns the Prompt field value
func (o *ActionConfirmationBody) GetPrompt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prompt
}

// GetPromptOk returns a tuple with the Prompt field value
// and a boolean to check if the value has been set.
func (o *ActionConfirmationBody) GetPromptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prompt, true
}

// SetPrompt sets field value
func (o *ActionConfirmationBody) SetPrompt(v string) {
	o.Prompt = v
}

// GetConfirmButtonLabel returns the ConfirmButtonLabel field value
func (o *ActionConfirmationBody) GetConfirmButtonLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfirmButtonLabel
}

// GetConfirmButtonLabelOk returns a tuple with the ConfirmButtonLabel field value
// and a boolean to check if the value has been set.
func (o *ActionConfirmationBody) GetConfirmButtonLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfirmButtonLabel, true
}

// SetConfirmButtonLabel sets field value
func (o *ActionConfirmationBody) SetConfirmButtonLabel(v string) {
	o.ConfirmButtonLabel = v
}

// GetCancelButtonLabel returns the CancelButtonLabel field value
func (o *ActionConfirmationBody) GetCancelButtonLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CancelButtonLabel
}

// GetCancelButtonLabelOk returns a tuple with the CancelButtonLabel field value
// and a boolean to check if the value has been set.
func (o *ActionConfirmationBody) GetCancelButtonLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelButtonLabel, true
}

// SetCancelButtonLabel sets field value
func (o *ActionConfirmationBody) SetCancelButtonLabel(v string) {
	o.CancelButtonLabel = v
}

func (o ActionConfirmationBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionConfirmationBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prompt"] = o.Prompt
	toSerialize["confirmButtonLabel"] = o.ConfirmButtonLabel
	toSerialize["cancelButtonLabel"] = o.CancelButtonLabel
	return toSerialize, nil
}

func (o *ActionConfirmationBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prompt",
		"confirmButtonLabel",
		"cancelButtonLabel",
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

	varActionConfirmationBody := _ActionConfirmationBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActionConfirmationBody)

	if err != nil {
		return err
	}

	*o = ActionConfirmationBody(varActionConfirmationBody)

	return err
}

type NullableActionConfirmationBody struct {
	value *ActionConfirmationBody
	isSet bool
}

func (v NullableActionConfirmationBody) Get() *ActionConfirmationBody {
	return v.value
}

func (v *NullableActionConfirmationBody) Set(val *ActionConfirmationBody) {
	v.value = val
	v.isSet = true
}

func (v NullableActionConfirmationBody) IsSet() bool {
	return v.isSet
}

func (v *NullableActionConfirmationBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionConfirmationBody(val *ActionConfirmationBody) *NullableActionConfirmationBody {
	return &NullableActionConfirmationBody{value: val, isSet: true}
}

func (v NullableActionConfirmationBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionConfirmationBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
