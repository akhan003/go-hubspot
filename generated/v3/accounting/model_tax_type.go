/*
Accounting Extension

These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accounting

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the TaxType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaxType{}

// TaxType Represents a tax in the external accounting system.
type TaxType struct {
	// The code/ID of the tax in the external accounting system.
	Code string `json:"code"`
	// The display name of the tax.
	Name *string `json:"name,omitempty"`
}

type _TaxType TaxType

// NewTaxType instantiates a new TaxType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxType(code string) *TaxType {
	this := TaxType{}
	this.Code = code
	return &this
}

// NewTaxTypeWithDefaults instantiates a new TaxType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxTypeWithDefaults() *TaxType {
	this := TaxType{}
	return &this
}

// GetCode returns the Code field value
func (o *TaxType) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *TaxType) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *TaxType) SetCode(v string) {
	o.Code = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TaxType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaxType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TaxType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TaxType) SetName(v string) {
	o.Name = &v
}

func (o TaxType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaxType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

func (o *TaxType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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

	varTaxType := _TaxType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTaxType)

	if err != nil {
		return err
	}

	*o = TaxType(varTaxType)

	return err
}

type NullableTaxType struct {
	value *TaxType
	isSet bool
}

func (v NullableTaxType) Get() *TaxType {
	return v.value
}

func (v *NullableTaxType) Set(val *TaxType) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxType) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxType(val *TaxType) *NullableTaxType {
	return &NullableTaxType{value: val, isSet: true}
}

func (v NullableTaxType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
